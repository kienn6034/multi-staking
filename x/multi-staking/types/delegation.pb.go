// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multistaking/v1/delegation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type MultiStakingLock struct {
	IntermediaryAccount string                                 `protobuf:"bytes,1,opt,name=intermediary_account,json=intermediaryAccount,proto3" json:"intermediary_account,omitempty"`
	ConversionRatio     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=conversion_ratio,json=conversionRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"conversion_ratio"`
	LockedAmount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=locked_amount,json=lockedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"locked_amount"`
}

func (m *MultiStakingLock) Reset()         { *m = MultiStakingLock{} }
func (m *MultiStakingLock) String() string { return proto.CompactTextString(m) }
func (*MultiStakingLock) ProtoMessage()    {}
func (*MultiStakingLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dade11e2944357, []int{0}
}
func (m *MultiStakingLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiStakingLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiStakingLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiStakingLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiStakingLock.Merge(m, src)
}
func (m *MultiStakingLock) XXX_Size() int {
	return m.Size()
}
func (m *MultiStakingLock) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiStakingLock.DiscardUnknown(m)
}

var xxx_messageInfo_MultiStakingLock proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MultiStakingLock)(nil), "multistaking.v1.MultiStakingLock")
}

func init() { proto.RegisterFile("multistaking/v1/delegation.proto", fileDescriptor_e9dade11e2944357) }

var fileDescriptor_e9dade11e2944357 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0x5e, 0xe9, 0xea, 0xde, 0x08, 0xd4, 0x2a, 0x74, 0x28, 0x1d, 0x92, 0x8a, 0x01,
	0xb1, 0x24, 0x51, 0xc5, 0x04, 0x62, 0x69, 0xd5, 0xa5, 0x02, 0x96, 0x74, 0x63, 0x89, 0x52, 0xc7,
	0x72, 0xad, 0x24, 0x3e, 0x95, 0xed, 0x56, 0xf4, 0x0d, 0x18, 0x79, 0x84, 0x3e, 0x44, 0x79, 0x87,
	0x8e, 0x55, 0x27, 0xc4, 0x50, 0xa1, 0x76, 0xe1, 0x31, 0x50, 0xec, 0x20, 0x22, 0x36, 0x26, 0xfb,
	0xe8, 0x3f, 0xe7, 0xfb, 0x7f, 0xe9, 0xb7, 0x3a, 0xf9, 0x2c, 0x93, 0x54, 0xc8, 0x38, 0xa5, 0x8c,
	0x04, 0xf3, 0x6e, 0x90, 0xe0, 0x0c, 0x93, 0x58, 0x52, 0x60, 0xfe, 0x94, 0x83, 0x04, 0xbb, 0x5e,
	0xdd, 0xf0, 0xe7, 0xdd, 0x76, 0x93, 0x00, 0x01, 0xa5, 0x05, 0xc5, 0x4f, 0xaf, 0xb5, 0x4f, 0x11,
	0x88, 0x1c, 0x44, 0xa4, 0x05, 0x3d, 0x68, 0xe9, 0xec, 0xa5, 0x66, 0x35, 0xee, 0x0b, 0xc8, 0x48,
	0x43, 0xee, 0x00, 0xa5, 0xf6, 0xad, 0xd5, 0xa4, 0x4c, 0x62, 0x9e, 0xe3, 0x84, 0xc6, 0x7c, 0x11,
	0xc5, 0x08, 0xc1, 0x8c, 0xc9, 0x96, 0xd9, 0x31, 0x2f, 0xfe, 0xf7, 0x5b, 0xdb, 0x95, 0xd7, 0x2c,
	0x21, 0xbd, 0x24, 0xe1, 0x58, 0x88, 0x91, 0xe4, 0x94, 0x91, 0xf0, 0xa4, 0x7a, 0xd5, 0xd3, 0x47,
	0x36, 0xb1, 0x1a, 0x08, 0xd8, 0x1c, 0x73, 0x41, 0x81, 0x45, 0xbc, 0x88, 0xdf, 0xaa, 0x29, 0xd0,
	0xcd, 0x7a, 0xe7, 0x1a, 0x6f, 0x3b, 0xf7, 0x9c, 0x50, 0x39, 0x99, 0x8d, 0x7d, 0x04, 0x79, 0x19,
	0xae, 0x7c, 0x3c, 0x91, 0xa4, 0x81, 0x5c, 0x4c, 0xb1, 0xf0, 0x07, 0x18, 0x6d, 0x57, 0x9e, 0x55,
	0xda, 0x0e, 0x30, 0x0a, 0xeb, 0xdf, 0xd4, 0xb0, 0x80, 0xda, 0xb1, 0x75, 0x9c, 0x01, 0x4a, 0x71,
	0x12, 0xc5, 0xb9, 0x8a, 0xfb, 0xe7, 0xd7, 0x2e, 0x43, 0x26, 0x2b, 0x2e, 0x43, 0x26, 0xc3, 0x23,
	0x8d, 0xec, 0x29, 0xe2, 0xf5, 0xbf, 0xa7, 0xa5, 0x6b, 0x7c, 0x2c, 0x5d, 0xa3, 0x3f, 0x5a, 0xef,
	0x1d, 0x73, 0xb3, 0x77, 0xcc, 0xf7, 0xbd, 0x63, 0x3e, 0x1f, 0x1c, 0x63, 0x73, 0x70, 0x8c, 0xd7,
	0x83, 0x63, 0x3c, 0x5c, 0x55, 0x7c, 0x38, 0x8e, 0x33, 0x0a, 0x12, 0xa3, 0x49, 0xa0, 0x9a, 0xf2,
	0xbe, 0xca, 0x7c, 0xfc, 0x31, 0x2b, 0xfb, 0xf1, 0x5f, 0xd5, 0xc9, 0xe5, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x6f, 0xed, 0xc5, 0xf9, 0x01, 0x00, 0x00,
}

func (m *MultiStakingLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiStakingLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiStakingLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LockedAmount.Size()
		i -= size
		if _, err := m.LockedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ConversionRatio.Size()
		i -= size
		if _, err := m.ConversionRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.IntermediaryAccount) > 0 {
		i -= len(m.IntermediaryAccount)
		copy(dAtA[i:], m.IntermediaryAccount)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.IntermediaryAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiStakingLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IntermediaryAccount)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	l = m.ConversionRatio.Size()
	n += 1 + l + sovDelegation(uint64(l))
	l = m.LockedAmount.Size()
	n += 1 + l + sovDelegation(uint64(l))
	return n
}

func sovDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegation(x uint64) (n int) {
	return sovDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiStakingLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegation
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
			return fmt.Errorf("proto: MultiStakingLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiStakingLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntermediaryAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntermediaryAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversionRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConversionRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegation
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
func skipDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
				return 0, ErrInvalidLengthDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegation = fmt.Errorf("proto: unexpected end of group")
)
