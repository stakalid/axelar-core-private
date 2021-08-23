// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: snapshot/v1beta1/query.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
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

type QueryValidatorsResponse struct {
	Validators []*QueryValidatorsResponse_Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *QueryValidatorsResponse) Reset()         { *m = QueryValidatorsResponse{} }
func (m *QueryValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse) ProtoMessage()    {}
func (*QueryValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12258278c48b9b0b, []int{0}
}
func (m *QueryValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse.Merge(m, src)
}
func (m *QueryValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse proto.InternalMessageInfo

type QueryValidatorsResponse_Validator struct {
	OperatorAddress string                 `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	Moniker         string                 `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Info            exported.ValidatorInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info"`
}

func (m *QueryValidatorsResponse_Validator) Reset()         { *m = QueryValidatorsResponse_Validator{} }
func (m *QueryValidatorsResponse_Validator) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse_Validator) ProtoMessage()    {}
func (*QueryValidatorsResponse_Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_12258278c48b9b0b, []int{0, 0}
}
func (m *QueryValidatorsResponse_Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse_Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse_Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse_Validator.Merge(m, src)
}
func (m *QueryValidatorsResponse_Validator) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse_Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse_Validator proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryValidatorsResponse)(nil), "snapshot.v1beta1.QueryValidatorsResponse")
	proto.RegisterType((*QueryValidatorsResponse_Validator)(nil), "snapshot.v1beta1.QueryValidatorsResponse.Validator")
}

func init() { proto.RegisterFile("snapshot/v1beta1/query.proto", fileDescriptor_12258278c48b9b0b) }

var fileDescriptor_12258278c48b9b0b = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0xec, 0x02, 0xd1, 0xb0, 0x1c, 0x24, 0x8d, 0x89, 0x0d, 0x31, 0x2b, 0x31, 0x31, 0xa9, 0x07,
	0xdb, 0x00, 0xf1, 0x03, 0xe4, 0xe6, 0xd1, 0x92, 0x78, 0xf0, 0x62, 0x16, 0xfa, 0x80, 0x06, 0xd8,
	0xb7, 0xee, 0x2e, 0x08, 0x1f, 0xe0, 0xdd, 0x83, 0x1f, 0xc5, 0x91, 0xa3, 0x27, 0xa3, 0xf0, 0x23,
	0xa6, 0x2d, 0x5d, 0x8d, 0x89, 0xb7, 0x7d, 0x33, 0xf3, 0x66, 0xde, 0x0e, 0x3d, 0xd5, 0x82, 0x4b,
	0x3d, 0x46, 0x13, 0x2e, 0x5a, 0x7d, 0x30, 0xbc, 0x15, 0x3e, 0xcd, 0x41, 0xad, 0x02, 0xa9, 0xd0,
	0xa0, 0x5b, 0x2f, 0xd8, 0x60, 0xcf, 0x36, 0x8e, 0x47, 0x38, 0xc2, 0x8c, 0x0c, 0xd3, 0x57, 0xae,
	0x6b, 0x5c, 0x58, 0x17, 0x58, 0x4a, 0x54, 0x06, 0x62, 0x6b, 0x67, 0x56, 0x12, 0x74, 0x2e, 0x3b,
	0x7f, 0x29, 0xd1, 0x93, 0xbb, 0xd4, 0xfe, 0x9e, 0x4f, 0x93, 0x98, 0x1b, 0x54, 0x3a, 0x02, 0x2d,
	0x51, 0x68, 0x70, 0x7b, 0x94, 0x2e, 0x2c, 0xea, 0x91, 0x66, 0xd9, 0xaf, 0xb5, 0x3b, 0xc1, 0xdf,
	0xfc, 0xe0, 0x9f, 0xf5, 0xc0, 0x42, 0xd1, 0x2f, 0x9b, 0xc6, 0x1b, 0xa1, 0x55, 0xcb, 0xb8, 0x97,
	0xb4, 0x8e, 0x12, 0x54, 0xfa, 0x7e, 0xe4, 0x71, 0xac, 0x40, 0xa7, 0x41, 0xc4, 0xaf, 0x46, 0x47,
	0x05, 0x7e, 0x93, 0xc3, 0xae, 0x47, 0x0f, 0x67, 0x28, 0x92, 0x09, 0x28, 0xaf, 0x94, 0x29, 0x8a,
	0xd1, 0xed, 0xd2, 0x4a, 0x22, 0x86, 0xe8, 0x95, 0x9b, 0xc4, 0xaf, 0xb5, 0xfd, 0x9f, 0x0b, 0x8b,
	0x9f, 0xdb, 0x53, 0x6d, 0xf0, 0xad, 0x18, 0x62, 0xb7, 0xb2, 0xfe, 0x38, 0x73, 0xa2, 0x6c, 0xb7,
	0xdb, 0x5b, 0x7f, 0x31, 0x67, 0xbd, 0x65, 0x64, 0xb3, 0x65, 0xe4, 0x73, 0xcb, 0xc8, 0xeb, 0x8e,
	0x39, 0x9b, 0x1d, 0x73, 0xde, 0x77, 0xcc, 0x79, 0xb8, 0x1e, 0x25, 0x66, 0x3c, 0xef, 0x07, 0x03,
	0x9c, 0x85, 0x7c, 0x09, 0x53, 0xae, 0x04, 0x98, 0x67, 0x54, 0x93, 0xfd, 0x74, 0x35, 0x40, 0x05,
	0xe1, 0x32, 0xb4, 0x9d, 0x67, 0x15, 0xf7, 0x0f, 0xb2, 0x8e, 0x3b, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x62, 0x08, 0x27, 0xd2, 0x01, 0x00, 0x00,
}

func (m *QueryValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsResponse_Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse_Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse_Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryValidatorsResponse_Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.Info.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &QueryValidatorsResponse_Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorsResponse_Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)