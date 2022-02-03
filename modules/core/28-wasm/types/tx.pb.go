// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/wasm/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Message type to push new wasm code
type MsgPushNewWasmCode struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Code   []byte `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *MsgPushNewWasmCode) Reset()         { *m = MsgPushNewWasmCode{} }
func (m *MsgPushNewWasmCode) String() string { return proto.CompactTextString(m) }
func (*MsgPushNewWasmCode) ProtoMessage()    {}
func (*MsgPushNewWasmCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7804a9f49664df6, []int{0}
}
func (m *MsgPushNewWasmCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPushNewWasmCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPushNewWasmCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPushNewWasmCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPushNewWasmCode.Merge(m, src)
}
func (m *MsgPushNewWasmCode) XXX_Size() int {
	return m.Size()
}
func (m *MsgPushNewWasmCode) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPushNewWasmCode.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPushNewWasmCode proto.InternalMessageInfo

func (m *MsgPushNewWasmCode) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgPushNewWasmCode) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// Response in case of successful handling
type MsgPushNewWasmCodeResponse struct {
	CodeId []byte `protobuf:"bytes,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
}

func (m *MsgPushNewWasmCodeResponse) Reset()         { *m = MsgPushNewWasmCodeResponse{} }
func (m *MsgPushNewWasmCodeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPushNewWasmCodeResponse) ProtoMessage()    {}
func (*MsgPushNewWasmCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7804a9f49664df6, []int{1}
}
func (m *MsgPushNewWasmCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPushNewWasmCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPushNewWasmCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPushNewWasmCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPushNewWasmCodeResponse.Merge(m, src)
}
func (m *MsgPushNewWasmCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPushNewWasmCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPushNewWasmCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPushNewWasmCodeResponse proto.InternalMessageInfo

func (m *MsgPushNewWasmCodeResponse) GetCodeId() []byte {
	if m != nil {
		return m.CodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgPushNewWasmCode)(nil), "ibc.core.wasm.v1.MsgPushNewWasmCode")
	proto.RegisterType((*MsgPushNewWasmCodeResponse)(nil), "ibc.core.wasm.v1.MsgPushNewWasmCodeResponse")
}

func init() { proto.RegisterFile("ibc/core/wasm/v1/tx.proto", fileDescriptor_e7804a9f49664df6) }

var fileDescriptor_e7804a9f49664df6 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x4f, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0x4c, 0x4a, 0xd6, 0x03, 0x49, 0xe9, 0x81, 0xa4,
	0xf4, 0xca, 0x0c, 0x95, 0x1c, 0xb8, 0x84, 0x7c, 0x8b, 0xd3, 0x03, 0x4a, 0x8b, 0x33, 0xfc, 0x52,
	0xcb, 0xc3, 0x13, 0x8b, 0x73, 0x9d, 0xf3, 0x53, 0x52, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x33, 0xd3,
	0xf3, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x21, 0x2e, 0x96,
	0xe4, 0xfc, 0x94, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0xc9, 0x94, 0x4b,
	0x0a, 0xd3, 0x84, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x71, 0x2e, 0x76, 0x90,
	0xaa, 0xf8, 0xcc, 0x14, 0xb0, 0x51, 0x3c, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x51, 0x0e, 0x17,
	0xb3, 0x6f, 0x71, 0xba, 0x50, 0x2a, 0x17, 0x3f, 0xba, 0xe5, 0x2a, 0x7a, 0xe8, 0xae, 0xd4, 0xc3,
	0xb4, 0x40, 0x4a, 0x87, 0x18, 0x55, 0x30, 0x67, 0x38, 0x05, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x59, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x66, 0x52, 0xb2, 0x6e, 0x7a, 0xbe, 0x7e, 0x99,
	0xb1, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0x31, 0x24, 0x34, 0x8d, 0x2c, 0x74, 0xc1, 0x01,
	0x5a, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xff, 0xd0, 0x5a, 0xe8, 0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PushNewWasmCode defines a rpc handler method for PushNewWasmCode.
	PushNewWasmCode(ctx context.Context, in *MsgPushNewWasmCode, opts ...grpc.CallOption) (*MsgPushNewWasmCodeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PushNewWasmCode(ctx context.Context, in *MsgPushNewWasmCode, opts ...grpc.CallOption) (*MsgPushNewWasmCodeResponse, error) {
	out := new(MsgPushNewWasmCodeResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.wasm.v1.Msg/PushNewWasmCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PushNewWasmCode defines a rpc handler method for PushNewWasmCode.
	PushNewWasmCode(context.Context, *MsgPushNewWasmCode) (*MsgPushNewWasmCodeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PushNewWasmCode(ctx context.Context, req *MsgPushNewWasmCode) (*MsgPushNewWasmCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNewWasmCode not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PushNewWasmCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPushNewWasmCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PushNewWasmCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.wasm.v1.Msg/PushNewWasmCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PushNewWasmCode(ctx, req.(*MsgPushNewWasmCode))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.wasm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushNewWasmCode",
			Handler:    _Msg_PushNewWasmCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/wasm/v1/tx.proto",
}

func (m *MsgPushNewWasmCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPushNewWasmCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPushNewWasmCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPushNewWasmCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPushNewWasmCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPushNewWasmCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CodeId) > 0 {
		i -= len(m.CodeId)
		copy(dAtA[i:], m.CodeId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CodeId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPushNewWasmCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPushNewWasmCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CodeId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPushNewWasmCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPushNewWasmCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPushNewWasmCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPushNewWasmCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPushNewWasmCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPushNewWasmCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeId = append(m.CodeId[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeId == nil {
				m.CodeId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
