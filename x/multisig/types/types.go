package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/multisig/exported"
	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	"github.com/axelarnetwork/utils/funcs"
	"github.com/axelarnetwork/utils/slices"
)

// Signature is an alias for signature in raw bytes
type Signature []byte

// ValidateBasic returns an error if the signature is not a valid S256 elliptic curve signature
func (sig Signature) ValidateBasic() error {
	_, err := btcec.ParseDERSignature(sig, btcec.S256())
	if err != nil {
		return err
	}

	return nil
}

// Verify checks if the signature matches the payload and public key
func (sig Signature) Verify(payload []byte, pk PublicKey) bool {
	s, err := btcec.ParseDERSignature(sig, btcec.S256())
	if err != nil {
		return false
	}

	parsedKey, err := btcec.ParsePubKey(pk, btcec.S256())
	if err != nil {
		return false
	}

	return s.Verify(payload, parsedKey)
}

// PublicKey is an alias for public key in raw bytes
type PublicKey []byte

// ValidateBasic returns an error if the given public key is invalid; nil otherwise
func (pk PublicKey) ValidateBasic() error {
	if _, err := btcec.ParsePubKey(pk, btcec.S256()); err != nil {
		return err
	}

	return nil
}

// String returns the hex encoding of the given public key
func (pk PublicKey) String() string {
	return hex.EncodeToString(pk)
}

func validateBasicPendingKey(key Key) error {
	if err := key.ID.ValidateBasic(); err != nil {
		return err
	}

	if err := key.Snapshot.ValidateBasic(); err != nil {
		return err
	}

	pubKeySeen := make(map[string]bool, len(key.PubKeys))
	for address, pubkey := range key.PubKeys {
		pubkeyStr := pubkey.String()
		if pubKeySeen[pubkeyStr] {
			return fmt.Errorf("duplicate public key seen")
		}
		pubKeySeen[pubkeyStr] = true

		p, err := sdk.ValAddressFromBech32(address)
		if err != nil {
			return err
		}

		if err := pubkey.ValidateBasic(); err != nil {
			return err
		}

		if key.Snapshot.GetParticipantWeight(p).IsZero() {
			return fmt.Errorf("invalid participant with public key submitted")
		}
	}

	return nil
}

// NewKeygenSession is the contructor for keygen session
func NewKeygenSession(id exported.KeyID, keygenThreshold utils.Threshold, signingThreshold utils.Threshold, snapshot snapshot.Snapshot, expiresAt int64, gracePeriod int64) KeygenSession {
	return KeygenSession{
		Key: Key{
			ID:               id,
			Snapshot:         snapshot,
			SigningThreshold: signingThreshold,
		},
		State:           exported.Pending,
		KeygenThreshold: keygenThreshold,
		ExpiresAt:       expiresAt,
		GracePeriod:     gracePeriod,
	}
}

// ValidateBasic returns an error if the given keygen session is invalid; nil otherwise
func (m KeygenSession) ValidateBasic() error {
	var keyErr error
	if m.State == exported.Completed {
		keyErr = m.Key.ValidateBasic()
	} else {
		keyErr = validateBasicPendingKey(m.Key)
	}

	if keyErr != nil {
		return keyErr
	}

	if m.KeygenThreshold.LT(m.Key.SigningThreshold) {
		return fmt.Errorf("keygen threshold must be >=signing threshold")
	}

	if m.ExpiresAt <= 0 {
		return fmt.Errorf("expires at must be >0")
	}

	if m.State == exported.Completed && m.CompletedAt == 0 {
		return fmt.Errorf("completed keygen session must have completed at set")
	}

	if m.State != exported.Completed && m.CompletedAt != 0 {
		return fmt.Errorf("pending keygen session must not have completed at set")
	}

	return nil
}

// GetKeyID returns the key ID of the given keygen session
func (m KeygenSession) GetKeyID() exported.KeyID {
	return m.Key.ID
}

// AddKey adds a new public key for the given participant into the keygen session
func (m *KeygenSession) AddKey(blockHeight int64, participant sdk.ValAddress, pubKey PublicKey) error {
	if m.Key.PubKeys == nil {
		m.Key.PubKeys = make(map[string]PublicKey)
		m.IsPubKeyReceived = make(map[string]bool)
	}

	if m.isExpired(blockHeight) {
		return fmt.Errorf("keygen session %s has expired", m.GetKeyID())
	}

	if m.Key.Snapshot.GetParticipantWeight(participant).IsZero() {
		return fmt.Errorf("%s is not a participant of keygen %s", participant.String(), m.GetKeyID())
	}

	if _, ok := m.Key.PubKeys[participant.String()]; ok {
		return fmt.Errorf("participant %s already submitted its public key for keygen %s", participant.String(), m.GetKeyID())
	}

	if m.IsPubKeyReceived[pubKey.String()] {
		return fmt.Errorf("duplicate public key received")
	}

	if m.State == exported.Completed && !m.isWithinGracePeriod(blockHeight) {
		return fmt.Errorf("keygen session %s has closed", m.GetKeyID())
	}

	m.addKey(participant, pubKey)

	if m.State != exported.Completed && Key(m.Key).GetParticipantsWeight().GTE(m.Key.Snapshot.CalculateMinPassingWeight(m.KeygenThreshold)) {
		m.CompletedAt = blockHeight
		m.State = exported.Completed
	}

	return nil
}

// GetMissingParticipants returns all participants who failed to submit their public keys
func (m KeygenSession) GetMissingParticipants() []sdk.ValAddress {
	participants := m.Key.Snapshot.GetParticipantAddresses()

	return slices.Filter(participants, func(p sdk.ValAddress) bool {
		_, ok := m.Key.PubKeys[p.String()]

		return !ok
	})
}

// Result returns the generated key if the session is completed and the key is valid
func (m KeygenSession) Result() (Key, error) {
	if m.GetState() != exported.Completed {
		return Key{}, fmt.Errorf("keygen %s is not completed yet", m.GetKeyID())
	}

	key := Key(m.Key)
	funcs.MustNoErr(key.ValidateBasic())

	return key, nil
}

func (m KeygenSession) isWithinGracePeriod(blockHeight int64) bool {
	return blockHeight <= m.CompletedAt+m.GracePeriod
}

func (m KeygenSession) isExpired(blockHeight int64) bool {
	return blockHeight >= m.ExpiresAt
}

func (m *KeygenSession) addKey(participant sdk.ValAddress, pubKey PublicKey) {
	m.Key.PubKeys[participant.String()] = pubKey
	m.IsPubKeyReceived[pubKey.String()] = true
}

// GetParticipants returns the participants of the given key
func (m Key) GetParticipants() []sdk.ValAddress {
	participants := make([]sdk.ValAddress, 0, len(m.PubKeys))
	for address := range m.PubKeys {
		p := funcs.Must(sdk.ValAddressFromBech32(address))
		participants = append(participants, p)
	}

	sort.SliceStable(participants, func(i, j int) bool { return bytes.Compare(participants[i], participants[j]) < 0 })

	return participants
}

// GetParticipantsWeight returns the total weight of all participants who have submitted their public keys
func (m Key) GetParticipantsWeight() sdk.Uint {
	totalWeight := sdk.ZeroUint()
	for address := range m.PubKeys {
		p := funcs.Must(sdk.ValAddressFromBech32(address))
		totalWeight = totalWeight.Add(m.Snapshot.GetParticipantWeight(p))
	}

	return totalWeight
}

// ValidateBasic returns an error if the given key is invalid; nil otherwise
func (m Key) ValidateBasic() error {
	if err := validateBasicPendingKey(m); err != nil {
		return err
	}

	if m.GetParticipantsWeight().LT(m.Snapshot.CalculateMinPassingWeight(m.SigningThreshold)) {
		return fmt.Errorf("invalid signing threshold")
	}

	return nil
}
