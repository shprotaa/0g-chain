// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Params struct {
	TokensPerVote string `protobuf:"bytes,1,opt,name=tokens_per_vote,json=tokensPerVote,proto3" json:"tokens_per_vote,omitempty"`
	MaxQuorums    uint64 `protobuf:"varint,2,opt,name=max_quorums,json=maxQuorums,proto3" json:"max_quorums,omitempty"`
	EpochBlocks   uint64 `protobuf:"varint,3,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty"`
	EncodedSlices uint64 `protobuf:"varint,4,opt,name=encoded_slices,json=encodedSlices,proto3" json:"encoded_slices,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{0}
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

func (m *Params) GetTokensPerVote() string {
	if m != nil {
		return m.TokensPerVote
	}
	return ""
}

func (m *Params) GetMaxQuorums() uint64 {
	if m != nil {
		return m.MaxQuorums
	}
	return 0
}

func (m *Params) GetEpochBlocks() uint64 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

func (m *Params) GetEncodedSlices() uint64 {
	if m != nil {
		return m.EncodedSlices
	}
	return 0
}

// GenesisState defines the dasigners module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// params epoch_number the epoch number
	EpochNumber uint64 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	// signers defines all signers information
	Signers []*Signer `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
	// quorums_by_epoch defines chosen quorums by epoch
	QuorumsByEpoch []*Quorums `protobuf:"bytes,4,rep,name=quorums_by_epoch,json=quorumsByEpoch,proto3" json:"quorums_by_epoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *GenesisState) GetSigners() []*Signer {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *GenesisState) GetQuorumsByEpoch() []*Quorums {
	if m != nil {
		return m.QuorumsByEpoch
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "zgc.dasigners.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "zgc.dasigners.v1.GenesisState")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/genesis.proto", fileDescriptor_896efa766aaca3be) }

var fileDescriptor_896efa766aaca3be = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb3, 0x24, 0x0a, 0xc2, 0x69, 0x4b, 0xb5, 0xe2, 0xb0, 0xe9, 0x61, 0x13, 0x2a, 0x81,
	0x7a, 0x61, 0xdd, 0x16, 0x89, 0x07, 0x08, 0x42, 0x88, 0x0b, 0x2a, 0x1b, 0x89, 0x03, 0x17, 0xcb,
	0xeb, 0x0c, 0xce, 0xaa, 0xf1, 0xce, 0xb2, 0xf6, 0x46, 0x49, 0x9f, 0x82, 0x3b, 0x2f, 0xd4, 0x63,
	0x8f, 0x9c, 0x10, 0xda, 0xbc, 0x08, 0x62, 0x6c, 0xa8, 0x68, 0xb9, 0xd9, 0xdf, 0xff, 0xcf, 0xe8,
	0xf7, 0x6f, 0x96, 0x5e, 0x69, 0xc5, 0x17, 0xd2, 0x96, 0xba, 0x82, 0xc6, 0xf2, 0xf5, 0x19, 0xd7,
	0x50, 0x81, 0x2d, 0x6d, 0x56, 0x37, 0xe8, 0x30, 0x3e, 0xbc, 0xd2, 0x2a, 0xfb, 0xab, 0x67, 0xeb,
	0xb3, 0xa3, 0xb1, 0x42, 0x6b, 0xd0, 0x0a, 0xd2, 0xb9, 0xbf, 0x78, 0xf3, 0xd1, 0x13, 0x8d, 0x1a,
	0x3d, 0xff, 0x7d, 0x0a, 0x74, 0xac, 0x11, 0xf5, 0x0a, 0x38, 0xdd, 0x8a, 0xf6, 0x33, 0x97, 0xd5,
	0x36, 0x48, 0x93, 0xbb, 0x92, 0x2b, 0x0d, 0x58, 0x27, 0x4d, 0x1d, 0x0c, 0xd3, 0x7b, 0xf1, 0x6e,
	0xb3, 0x90, 0xe3, 0xf8, 0x5b, 0xc4, 0x86, 0x17, 0xb2, 0x91, 0xc6, 0xc6, 0xcf, 0xd9, 0x63, 0x87,
	0x97, 0x50, 0x59, 0x51, 0x43, 0x23, 0xd6, 0xe8, 0x20, 0x89, 0xa6, 0xd1, 0xc9, 0xa3, 0x7c, 0xdf,
	0xe3, 0x0b, 0x68, 0x3e, 0xa2, 0x83, 0x78, 0xc2, 0x46, 0x46, 0x6e, 0xc4, 0x97, 0x16, 0x9b, 0xd6,
	0xd8, 0xe4, 0xc1, 0x34, 0x3a, 0x19, 0xe4, 0xcc, 0xc8, 0xcd, 0x07, 0x4f, 0xe2, 0xa7, 0x6c, 0x0f,
	0x6a, 0x54, 0x4b, 0x51, 0xac, 0x50, 0x5d, 0xda, 0xa4, 0x4f, 0x8e, 0x11, 0xb1, 0x19, 0xa1, 0xf8,
	0x19, 0x3b, 0x80, 0x4a, 0xe1, 0x02, 0x16, 0xc2, 0xae, 0x4a, 0x05, 0x36, 0x19, 0x90, 0x69, 0x3f,
	0xd0, 0x39, 0xc1, 0xe3, 0x2e, 0x62, 0x7b, 0x6f, 0x7d, 0xa1, 0x73, 0x27, 0x1d, 0xc4, 0xaf, 0xd8,
	0xb0, 0xa6, 0xb4, 0x14, 0x6d, 0x74, 0x9e, 0x64, 0x77, 0x0b, 0xce, 0xfc, 0x6b, 0x66, 0x83, 0xeb,
	0x1f, 0x93, 0x5e, 0x1e, 0xdc, 0xb7, 0x91, 0xaa, 0xd6, 0x14, 0xd0, 0x84, 0xd0, 0x3e, 0xd2, 0x7b,
	0x42, 0xf1, 0x39, 0x7b, 0x18, 0xb6, 0x24, 0xfd, 0x69, 0xff, 0xff, 0xbb, 0xe7, 0x74, 0xcc, 0xff,
	0x18, 0xe3, 0xd7, 0xec, 0x30, 0xd4, 0x20, 0x8a, 0xad, 0xa0, 0x6d, 0xc9, 0x80, 0x86, 0xc7, 0xf7,
	0x87, 0x43, 0x3d, 0xf9, 0x41, 0x18, 0x99, 0x6d, 0xdf, 0x50, 0x23, 0xef, 0xae, 0xbb, 0x34, 0xba,
	0xe9, 0xd2, 0xe8, 0x67, 0x97, 0x46, 0x5f, 0x77, 0x69, 0xef, 0x66, 0x97, 0xf6, 0xbe, 0xef, 0xd2,
	0xde, 0x27, 0xae, 0x4b, 0xb7, 0x6c, 0x8b, 0x4c, 0xa1, 0xe1, 0xa7, 0x7a, 0x25, 0x0b, 0xcb, 0x4f,
	0xf5, 0x0b, 0xb5, 0x94, 0x65, 0xc5, 0x37, 0xff, 0xfe, 0xab, 0xdb, 0xd6, 0x60, 0x8b, 0x21, 0x7d,
	0xea, 0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0xe1, 0x73, 0x5d, 0x97, 0x02, 0x00, 0x00,
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
	if m.EncodedSlices != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EncodedSlices))
		i--
		dAtA[i] = 0x20
	}
	if m.EpochBlocks != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxQuorums != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxQuorums))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TokensPerVote) > 0 {
		i -= len(m.TokensPerVote)
		copy(dAtA[i:], m.TokensPerVote)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokensPerVote)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QuorumsByEpoch) > 0 {
		for iNdEx := len(m.QuorumsByEpoch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QuorumsByEpoch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EpochNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	l = len(m.TokensPerVote)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxQuorums != 0 {
		n += 1 + sovGenesis(uint64(m.MaxQuorums))
	}
	if m.EpochBlocks != 0 {
		n += 1 + sovGenesis(uint64(m.EpochBlocks))
	}
	if m.EncodedSlices != 0 {
		n += 1 + sovGenesis(uint64(m.EncodedSlices))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovGenesis(uint64(m.EpochNumber))
	}
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.QuorumsByEpoch) > 0 {
		for _, e := range m.QuorumsByEpoch {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerVote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokensPerVote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxQuorums", wireType)
			}
			m.MaxQuorums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxQuorums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedSlices", wireType)
			}
			m.EncodedSlices = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncodedSlices |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &Signer{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumsByEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuorumsByEpoch = append(m.QuorumsByEpoch, &Quorums{})
			if err := m.QuorumsByEpoch[len(m.QuorumsByEpoch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
