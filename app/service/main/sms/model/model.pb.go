// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ModelUserActionLog struct {
	MsgID                string   `protobuf:"bytes,1,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Desc                 string   `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Type                 int32    `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Action               int32    `protobuf:"varint,7,opt,name=action,proto3" json:"action,omitempty"`
	Ts                   int64    `protobuf:"varint,8,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelUserActionLog) Reset()         { *m = ModelUserActionLog{} }
func (m *ModelUserActionLog) String() string { return proto.CompactTextString(m) }
func (*ModelUserActionLog) ProtoMessage()    {}
func (*ModelUserActionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}
func (m *ModelUserActionLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModelUserActionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModelUserActionLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModelUserActionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelUserActionLog.Merge(m, src)
}
func (m *ModelUserActionLog) XXX_Size() int {
	return m.Size()
}
func (m *ModelUserActionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelUserActionLog.DiscardUnknown(m)
}

var xxx_messageInfo_ModelUserActionLog proto.InternalMessageInfo

func (m *ModelUserActionLog) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ModelUserActionLog) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ModelUserActionLog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ModelUserActionLog) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ModelUserActionLog) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ModelUserActionLog) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ModelUserActionLog) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ModelUserActionLog) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func init() {
	proto.RegisterType((*ModelUserActionLog)(nil), "sms.service.model.ModelUserActionLog")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x71, 0xdb, 0xa4, 0xd4, 0x48, 0x48, 0x78, 0x00, 0x8b, 0x21, 0x8d, 0x98, 0xb2, 0x90,
	0x0e, 0x3c, 0x01, 0x15, 0x0b, 0x12, 0x5d, 0x22, 0xb1, 0xb0, 0x35, 0x8e, 0x31, 0x96, 0xea, 0x5c,
	0x95, 0xbb, 0x20, 0xf1, 0x26, 0x3c, 0x12, 0x1b, 0x3c, 0x01, 0x42, 0xe1, 0x45, 0x90, 0x2f, 0x65,
	0xfb, 0xbf, 0xff, 0xee, 0xb3, 0x6c, 0xcb, 0x93, 0x00, 0x8d, 0xdd, 0x95, 0xfb, 0x0e, 0x08, 0xd4,
	0x19, 0x06, 0x2c, 0xd1, 0x76, 0xaf, 0xde, 0xd8, 0x92, 0x07, 0x97, 0xd7, 0xce, 0xd3, 0x4b, 0x5f,
	0x97, 0x06, 0xc2, 0xca, 0x81, 0x83, 0x15, 0x6f, 0xd6, 0xfd, 0x33, 0x13, 0x03, 0xa7, 0xf1, 0x84,
	0xab, 0x4f, 0x21, 0xd5, 0x26, 0x8a, 0x8f, 0x68, 0xbb, 0x5b, 0x43, 0x1e, 0xda, 0x07, 0x70, 0x6a,
	0x29, 0x93, 0x80, 0xce, 0x37, 0x5a, 0xe4, 0xa2, 0x58, 0xac, 0x17, 0xc3, 0xf7, 0x32, 0xd9, 0xa0,
	0xbb, 0xbf, 0xab, 0xc6, 0x5e, 0x9d, 0xcb, 0x34, 0x40, 0xed, 0x77, 0x56, 0x4f, 0xe2, 0x46, 0x75,
	0x20, 0xa5, 0xe5, 0xdc, 0x40, 0x4b, 0xb6, 0x25, 0x3d, 0xe5, 0xc1, 0x3f, 0x46, 0x03, 0x69, 0x4b,
	0x3d, 0xea, 0xd9, 0x68, 0x8c, 0xa4, 0x94, 0x9c, 0x35, 0x16, 0x8d, 0x4e, 0xb8, 0xe5, 0x1c, 0x3b,
	0x7a, 0xdb, 0x5b, 0x9d, 0xe6, 0xa2, 0x48, 0x2a, 0xce, 0xd1, 0xdf, 0xf2, 0xfd, 0xf4, 0x9c, 0xdb,
	0x03, 0xa9, 0x53, 0x39, 0x21, 0xd4, 0xc7, 0xb9, 0x28, 0xa6, 0xd5, 0x84, 0x70, 0x7d, 0xf1, 0x31,
	0x64, 0xe2, 0x6b, 0xc8, 0xc4, 0xcf, 0x90, 0x89, 0xf7, 0xdf, 0xec, 0xe8, 0x29, 0xe1, 0x9f, 0xa9,
	0x53, 0x7e, 0xf1, 0xcd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xc1, 0x41, 0xad, 0x42, 0x01,
	0x00, 0x00,
}

func (m *ModelUserActionLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModelUserActionLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModelUserActionLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Ts != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.Ts))
		i--
		dAtA[i] = 0x40
	}
	if m.Action != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x38
	}
	if m.Type != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Mobile) > 0 {
		i -= len(m.Mobile)
		copy(dAtA[i:], m.Mobile)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Mobile)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MsgID) > 0 {
		i -= len(m.MsgID)
		copy(dAtA[i:], m.MsgID)
		i = encodeVarintModel(dAtA, i, uint64(len(m.MsgID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ModelUserActionLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgID)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Mobile)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovModel(uint64(m.Type))
	}
	if m.Action != 0 {
		n += 1 + sovModel(uint64(m.Action))
	}
	if m.Ts != 0 {
		n += 1 + sovModel(uint64(m.Ts))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ModelUserActionLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: ModelUserActionLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModelUserActionLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mobile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mobile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ts", wireType)
			}
			m.Ts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ts |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
				return 0, ErrInvalidLengthModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModel = fmt.Errorf("proto: unexpected end of group")
)