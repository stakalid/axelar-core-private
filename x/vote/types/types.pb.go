// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/vote/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/vote/exported"
	github_com_axelarnetwork_axelar_core_x_vote_exported "github.com/axelarnetwork/axelar-core/x/vote/exported"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "github.com/regen-network/cosmos-proto"
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

// TalliedVote represents a vote for a poll with the accumulated stake of all
// validators voting for the same data
type TalliedVote struct {
	Tally       github_com_cosmos_cosmos_sdk_types.Uint                     `protobuf:"bytes,1,opt,name=tally,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"tally"`
	Data        *types.Any                                                  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	PollID      github_com_axelarnetwork_axelar_core_x_vote_exported.PollID `protobuf:"varint,4,opt,name=poll_id,json=pollId,proto3,customtype=github.com/axelarnetwork/axelar-core/x/vote/exported.PollID" json:"poll_id"`
	IsVoterLate map[string]bool                                             `protobuf:"bytes,5,rep,name=is_voter_late,json=isVoterLate,proto3" json:"is_voter_late,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *TalliedVote) Reset()         { *m = TalliedVote{} }
func (m *TalliedVote) String() string { return proto.CompactTextString(m) }
func (*TalliedVote) ProtoMessage()    {}
func (*TalliedVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_584be12bf9f97fd2, []int{0}
}
func (m *TalliedVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalliedVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TalliedVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalliedVote.Merge(m, src)
}
func (m *TalliedVote) XXX_Size() int {
	return m.Size()
}
func (m *TalliedVote) XXX_DiscardUnknown() {
	xxx_messageInfo_TalliedVote.DiscardUnknown(m)
}

var xxx_messageInfo_TalliedVote proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TalliedVote)(nil), "axelar.vote.v1beta1.TalliedVote")
	proto.RegisterMapType((map[string]bool)(nil), "axelar.vote.v1beta1.TalliedVote.IsVoterLateEntry")
}

func init() { proto.RegisterFile("axelar/vote/v1beta1/types.proto", fileDescriptor_584be12bf9f97fd2) }

var fileDescriptor_584be12bf9f97fd2 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0xb1, 0x13, 0xca, 0x05, 0xa4, 0xc8, 0x64, 0x70, 0x33, 0xd8, 0x16, 0x03, 0x58,
	0x48, 0xb9, 0x53, 0xca, 0x82, 0x8a, 0x84, 0x44, 0x44, 0x87, 0x20, 0x2a, 0x2a, 0xab, 0x65, 0x60,
	0x89, 0x2e, 0xf1, 0xab, 0x6b, 0xe5, 0xea, 0x8b, 0xce, 0x97, 0x10, 0x7f, 0x0b, 0x46, 0x56, 0xbe,
	0x03, 0x7c, 0x87, 0x88, 0xa9, 0x23, 0x62, 0x88, 0xc0, 0xf9, 0x22, 0xc8, 0x3e, 0x23, 0x02, 0xca,
	0xd2, 0xe9, 0xfe, 0xef, 0xee, 0xe7, 0xbf, 0xdf, 0xfb, 0x3f, 0xec, 0xb1, 0x15, 0x70, 0x26, 0xe9,
	0x52, 0x28, 0xa0, 0xcb, 0xc1, 0x04, 0x14, 0x1b, 0x50, 0x95, 0xcf, 0x21, 0x23, 0x73, 0x29, 0x94,
	0xb0, 0x1f, 0x68, 0x80, 0x94, 0x00, 0xa9, 0x81, 0xde, 0x61, 0x2c, 0x44, 0xcc, 0x81, 0x56, 0xc8,
	0x64, 0x71, 0x49, 0x59, 0x9a, 0x6b, 0xbe, 0xd7, 0x8d, 0x45, 0x2c, 0x2a, 0x49, 0x4b, 0x55, 0xdf,
	0x1e, 0x4e, 0x45, 0x76, 0x2d, 0xb2, 0xb1, 0x7e, 0xd0, 0x45, 0xfd, 0x14, 0xec, 0x76, 0x00, 0xab,
	0xb9, 0x90, 0x0a, 0xa2, 0x7d, 0xad, 0x3c, 0xfc, 0x6a, 0xe2, 0xf6, 0x39, 0xe3, 0x3c, 0x81, 0xe8,
	0x9d, 0x50, 0x60, 0x9f, 0xe0, 0xa6, 0x62, 0x9c, 0xe7, 0x0e, 0xf2, 0x51, 0x70, 0x6f, 0x48, 0xd7,
	0x1b, 0xcf, 0xf8, 0xb1, 0xf1, 0x1e, 0xc7, 0x89, 0xba, 0x5a, 0x4c, 0xc8, 0x54, 0x5c, 0xd7, 0x7f,
	0xaa, 0x8f, 0x7e, 0x16, 0xcd, 0x6a, 0xc3, 0x8b, 0x24, 0x55, 0xa1, 0xfe, 0xda, 0x3e, 0xc7, 0x56,
	0xc4, 0x14, 0x73, 0x4c, 0x1f, 0x05, 0xed, 0xa3, 0x2e, 0xd1, 0xb3, 0x91, 0x3f, 0xb3, 0x91, 0x97,
	0x69, 0x3e, 0x7c, 0xf2, 0xed, 0x4b, 0xff, 0xd1, 0x3e, 0xdf, 0x08, 0xa6, 0xf4, 0xac, 0x24, 0x4f,
	0x99, 0xcc, 0xae, 0x18, 0x07, 0x19, 0x56, 0x6e, 0xf6, 0x25, 0xbe, 0x33, 0x17, 0x9c, 0x8f, 0x93,
	0xc8, 0xb1, 0x7c, 0x14, 0x58, 0xc3, 0xd3, 0xba, 0xbd, 0xe7, 0x3b, 0x36, 0x7a, 0xf4, 0x14, 0xd4,
	0x07, 0x21, 0x67, 0x75, 0xd5, 0x9f, 0x0a, 0x09, 0x74, 0xf5, 0x6f, 0x1e, 0xe4, 0x4c, 0x70, 0x3e,
	0x7a, 0x55, 0x6c, 0xbc, 0x96, 0x56, 0x61, 0xab, 0x74, 0x1f, 0x45, 0xf6, 0x05, 0xbe, 0x9f, 0x64,
	0xe3, 0x12, 0x96, 0x63, 0xce, 0x14, 0x38, 0x4d, 0xdf, 0x0c, 0xda, 0x47, 0x03, 0xb2, 0x67, 0x6f,
	0x64, 0x27, 0x3d, 0x32, 0xca, 0xca, 0x43, 0xbe, 0x61, 0x0a, 0x4e, 0x52, 0x25, 0xf3, 0xb0, 0x9d,
	0xfc, 0xbd, 0xe9, 0xbd, 0xc0, 0x9d, 0xff, 0x01, 0xbb, 0x83, 0xcd, 0x19, 0xe8, 0xb4, 0xef, 0x86,
	0xa5, 0xb4, 0xbb, 0xb8, 0xb9, 0x64, 0x7c, 0x01, 0x4e, 0xc3, 0x47, 0xc1, 0x41, 0xa8, 0x8b, 0xe3,
	0xc6, 0x33, 0x74, 0x6c, 0x7d, 0xfa, 0xec, 0xa1, 0xd7, 0xd6, 0x41, 0xa3, 0x63, 0x0e, 0xdf, 0xae,
	0x7f, 0xb9, 0xc6, 0xba, 0x70, 0xd1, 0x4d, 0xe1, 0xa2, 0x9f, 0x85, 0x8b, 0x3e, 0x6e, 0x5d, 0xe3,
	0x66, 0xeb, 0x1a, 0xdf, 0xb7, 0xae, 0xf1, 0x7e, 0x70, 0x9b, 0x3c, 0xaa, 0xed, 0x4d, 0x5a, 0xd5,
	0x6e, 0x9e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x3e, 0x19, 0x8d, 0xbd, 0x02, 0x00, 0x00,
}

func (m *TalliedVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalliedVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalliedVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IsVoterLate) > 0 {
		keysForIsVoterLate := make([]string, 0, len(m.IsVoterLate))
		for k := range m.IsVoterLate {
			keysForIsVoterLate = append(keysForIsVoterLate, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForIsVoterLate)
		for iNdEx := len(keysForIsVoterLate) - 1; iNdEx >= 0; iNdEx-- {
			v := m.IsVoterLate[string(keysForIsVoterLate[iNdEx])]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(keysForIsVoterLate[iNdEx])
			copy(dAtA[i:], keysForIsVoterLate[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(keysForIsVoterLate[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PollID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.PollID))
		i--
		dAtA[i] = 0x20
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Tally.Size()
		i -= size
		if _, err := m.Tally.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TalliedVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tally.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.PollID != 0 {
		n += 1 + sovTypes(uint64(m.PollID))
	}
	if len(m.IsVoterLate) > 0 {
		for k, v := range m.IsVoterLate {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTypes(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TalliedVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TalliedVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalliedVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tally", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollID", wireType)
			}
			m.PollID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PollID |= github_com_axelarnetwork_axelar_core_x_vote_exported.PollID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVoterLate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IsVoterLate == nil {
				m.IsVoterLate = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.IsVoterLate[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
