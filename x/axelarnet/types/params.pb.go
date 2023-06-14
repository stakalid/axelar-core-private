// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/axelarnet/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params represent the genesis parameters for the module
type Params struct {
	// IBC packet route timeout window
	RouteTimeoutWindow               uint64                          `protobuf:"varint,1,opt,name=route_timeout_window,json=routeTimeoutWindow,proto3" json:"route_timeout_window,omitempty"`
	TransferLimit                    uint64                          `protobuf:"varint,3,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	EndBlockerLimit                  uint64                          `protobuf:"varint,4,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	CallContractsProposalMinDeposits CallContractProposalMinDeposits `protobuf:"bytes,5,rep,name=call_contracts_proposal_min_deposits,json=callContractsProposalMinDeposits,proto3,castrepeated=CallContractProposalMinDeposits" json:"call_contracts_proposal_min_deposits"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9e88637795bf, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type CallContractProposalMinDeposit struct {
	Chain           github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ContractAddress string                                                          `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	MinDeposits     github_com_cosmos_cosmos_sdk_types.Coins                        `protobuf:"bytes,3,rep,name=min_deposits,json=minDeposits,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_deposits"`
}

func (m *CallContractProposalMinDeposit) Reset()         { *m = CallContractProposalMinDeposit{} }
func (m *CallContractProposalMinDeposit) String() string { return proto.CompactTextString(m) }
func (*CallContractProposalMinDeposit) ProtoMessage()    {}
func (*CallContractProposalMinDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9e88637795bf, []int{1}
}
func (m *CallContractProposalMinDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallContractProposalMinDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallContractProposalMinDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallContractProposalMinDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractProposalMinDeposit.Merge(m, src)
}
func (m *CallContractProposalMinDeposit) XXX_Size() int {
	return m.Size()
}
func (m *CallContractProposalMinDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractProposalMinDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractProposalMinDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.axelarnet.v1beta1.Params")
	proto.RegisterType((*CallContractProposalMinDeposit)(nil), "axelar.axelarnet.v1beta1.CallContractProposalMinDeposit")
}

func init() {
	proto.RegisterFile("axelar/axelarnet/v1beta1/params.proto", fileDescriptor_241d9e88637795bf)
}

var fileDescriptor_241d9e88637795bf = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0x93, 0xb6, 0x9b, 0xfe, 0x7f, 0x0f, 0xd8, 0x88, 0x76, 0x08, 0x3b, 0xa4, 0xd5, 0xc4,
	0x44, 0x41, 0x5a, 0xbc, 0x81, 0x84, 0xb8, 0x21, 0x52, 0x4e, 0x08, 0xd0, 0x14, 0x4d, 0x42, 0x70,
	0xb1, 0x1c, 0xc7, 0x74, 0x56, 0x13, 0xbf, 0x91, 0xed, 0xd2, 0xf2, 0x2d, 0x38, 0xf1, 0x21, 0xb8,
	0xf1, 0x29, 0xe8, 0x71, 0x47, 0x4e, 0x03, 0xda, 0x6f, 0xc1, 0x09, 0xc5, 0x4e, 0xdb, 0x21, 0xc1,
	0xc4, 0xc9, 0xc9, 0xf3, 0xfc, 0xf2, 0xbc, 0xd1, 0x93, 0x37, 0xe8, 0x80, 0x4e, 0x79, 0x41, 0x15,
	0x76, 0x87, 0xe4, 0x06, 0xbf, 0x3b, 0xce, 0xb8, 0xa1, 0xc7, 0xb8, 0xa2, 0x8a, 0x96, 0x3a, 0xae,
	0x14, 0x18, 0x08, 0x42, 0xe7, 0xc7, 0x2b, 0x2c, 0x6e, 0xb0, 0xbd, 0xdd, 0x21, 0x0c, 0xc1, 0x42,
	0xb8, 0xbe, 0x72, 0xfc, 0x5e, 0xc4, 0x40, 0x97, 0xa0, 0x71, 0x46, 0x35, 0x5f, 0x25, 0x32, 0x10,
	0xd2, 0xf9, 0xfb, 0x5f, 0x5a, 0x68, 0xf3, 0xc4, 0x0e, 0x08, 0x8e, 0xd0, 0xae, 0x82, 0xb1, 0xe1,
	0xc4, 0x88, 0x92, 0xc3, 0xd8, 0x90, 0x89, 0x90, 0x39, 0x4c, 0x42, 0xbf, 0xe7, 0xf7, 0x3b, 0x69,
	0x60, 0xbd, 0x53, 0x67, 0xbd, 0xb2, 0x4e, 0x70, 0x80, 0x6e, 0x18, 0x45, 0xa5, 0x7e, 0xcb, 0x15,
	0x29, 0x44, 0x29, 0x4c, 0xd8, 0xb6, 0xec, 0xf5, 0xa5, 0xfa, 0xbc, 0x16, 0x83, 0x7b, 0xe8, 0x26,
	0x97, 0x39, 0xc9, 0x0a, 0x60, 0xa3, 0x15, 0xd9, 0xb1, 0xe4, 0x36, 0x97, 0x79, 0xe2, 0x74, 0xc7,
	0x7e, 0xf6, 0xd1, 0x6d, 0x46, 0x8b, 0x82, 0x30, 0x90, 0x46, 0x51, 0x66, 0x34, 0xa9, 0x14, 0x54,
	0xa0, 0x69, 0x41, 0x4a, 0x21, 0x49, 0xce, 0x2b, 0xd0, 0xc2, 0xe8, 0x70, 0xa3, 0xd7, 0xee, 0x6f,
	0xdd, 0x7f, 0x14, 0xff, 0xad, 0x8f, 0x78, 0x40, 0x8b, 0x62, 0xd0, 0x84, 0x9c, 0x34, 0x11, 0x2f,
	0x84, 0x7c, 0xea, 0x02, 0x92, 0x3b, 0xb3, 0x8b, 0xae, 0xf7, 0xe9, 0x5b, 0xb7, 0x7b, 0x35, 0xa7,
	0xd3, 0x1e, 0xbb, 0x04, 0xe8, 0x3f, 0x10, 0xcf, 0x3a, 0xff, 0xb5, 0x76, 0xda, 0xfb, 0x1f, 0x5b,
	0x28, 0xba, 0x3a, 0x2b, 0x78, 0x8d, 0x36, 0xd8, 0x19, 0x15, 0xd2, 0x56, 0xfa, 0x7f, 0x32, 0xf8,
	0x79, 0xd1, 0x7d, 0x3c, 0x14, 0xe6, 0x6c, 0x9c, 0xc5, 0x0c, 0xca, 0xf5, 0xa7, 0x9f, 0x80, 0x1a,
	0x35, 0x77, 0x87, 0x0c, 0x14, 0xc7, 0x53, 0x2c, 0xf9, 0x74, 0xac, 0x31, 0x9f, 0x56, 0xa0, 0x0c,
	0xcf, 0xe3, 0x41, 0x1d, 0xf3, 0x92, 0x96, 0x3c, 0x75, 0x89, 0xc1, 0x5d, 0xb4, 0xb3, 0x6c, 0x8c,
	0xd0, 0x3c, 0x57, 0x5c, 0xeb, 0xb0, 0x55, 0x4f, 0x49, 0xb7, 0x97, 0xfa, 0x13, 0x27, 0x07, 0x12,
	0x5d, 0xfb, 0xad, 0xc9, 0xb6, 0x6d, 0xf2, 0x56, 0xec, 0x36, 0x25, 0xae, 0x37, 0x65, 0x5d, 0x22,
	0x08, 0x99, 0x1c, 0x35, 0x55, 0xf5, 0x2f, 0xbd, 0x6b, 0xb3, 0x56, 0xee, 0x38, 0xd4, 0xf9, 0x08,
	0x9b, 0xf7, 0x15, 0xd7, 0xf6, 0x01, 0x9d, 0x6e, 0x95, 0xeb, 0x7a, 0x92, 0xd3, 0xd9, 0x8f, 0xc8,
	0x9b, 0xcd, 0x23, 0xff, 0x7c, 0x1e, 0xf9, 0xdf, 0xe7, 0x91, 0xff, 0x61, 0x11, 0x79, 0xe7, 0x8b,
	0xc8, 0xfb, 0xba, 0x88, 0xbc, 0x37, 0x0f, 0xff, 0xb1, 0x80, 0xf5, 0x7f, 0x61, 0x07, 0x65, 0x9b,
	0x76, 0x7f, 0x1f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xba, 0xbf, 0x18, 0xc2, 0x38, 0x03, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for iNdEx := len(m.CallContractsProposalMinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallContractsProposalMinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.RouteTimeoutWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RouteTimeoutWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CallContractProposalMinDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallContractProposalMinDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallContractProposalMinDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinDeposits) > 0 {
		for iNdEx := len(m.MinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RouteTimeoutWindow != 0 {
		n += 1 + sovParams(uint64(m.RouteTimeoutWindow))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for _, e := range m.CallContractsProposalMinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *CallContractProposalMinDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.MinDeposits) > 0 {
		for _, e := range m.MinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteTimeoutWindow", wireType)
			}
			m.RouteTimeoutWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteTimeoutWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallContractsProposalMinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallContractsProposalMinDeposits = append(m.CallContractsProposalMinDeposits, CallContractProposalMinDeposit{})
			if err := m.CallContractsProposalMinDeposits[len(m.CallContractsProposalMinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CallContractProposalMinDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CallContractProposalMinDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallContractProposalMinDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposits = append(m.MinDeposits, types.Coin{})
			if err := m.MinDeposits[len(m.MinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
