// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gov/v1beta/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty" yaml:"starting_proposal_id"`
	// deposits defines all the deposits present at genesis.
	Deposits Deposits `protobuf:"bytes,2,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	// votes defines all the votes present at genesis.
	Votes Votes `protobuf:"bytes,3,rep,name=votes,proto3,castrepeated=Votes" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals Proposals `protobuf:"bytes,4,rep,name=proposals,proto3,castrepeated=Proposals" json:"proposals"`
	// params defines all the paramaters of related to deposit.
	DepositParams DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params" yaml:"deposit_params"`
	// params defines all the paramaters of related to voting.
	VotingParams VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params" yaml:"voting_params"`
	// params defines all the paramaters of related to tally.
	TallyParams TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params" yaml:"tally_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7615c42136362a1b, []int{0}
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

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() Deposits {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() Votes {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() Proposals {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return DepositParams{}
}

func (m *GenesisState) GetVotingParams() VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return VotingParams{}
}

func (m *GenesisState) GetTallyParams() TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return TallyParams{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "gov.v1beta.GenesisState")
}

func init() { proto.RegisterFile("gov/v1beta/genesis.proto", fileDescriptor_7615c42136362a1b) }

var fileDescriptor_7615c42136362a1b = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0xae, 0xd2, 0x40,
	0x18, 0x86, 0x5b, 0xf9, 0x11, 0x06, 0x30, 0x38, 0xd4, 0x58, 0x45, 0x5b, 0xd2, 0x15, 0x0b, 0xd3,
	0x46, 0x8c, 0x1b, 0x77, 0x36, 0xfe, 0xc4, 0x1d, 0x56, 0xa3, 0x89, 0x2e, 0x9a, 0x81, 0x4e, 0x86,
	0xc6, 0x96, 0x69, 0x3a, 0x63, 0x95, 0xbb, 0xf0, 0x3a, 0xbc, 0x03, 0xef, 0x80, 0x25, 0x4b, 0x57,
	0x9c, 0x13, 0xb8, 0x03, 0xae, 0xe0, 0xa4, 0x9d, 0x29, 0x94, 0x13, 0x76, 0xed, 0x3b, 0xef, 0xf3,
	0xe4, 0x9b, 0x2f, 0x03, 0x74, 0x42, 0x33, 0x27, 0x7b, 0x3e, 0xc3, 0x1c, 0x39, 0x04, 0x2f, 0x31,
	0x0b, 0x99, 0x9d, 0xa4, 0x94, 0x53, 0x08, 0x08, 0xcd, 0x6c, 0x71, 0xf2, 0x58, 0x23, 0x94, 0xd0,
	0x22, 0x76, 0xf2, 0x2f, 0xd1, 0xc8, 0xd3, 0x23, 0x1b, 0x33, 0x22, 0x52, 0xeb, 0x5f, 0x1d, 0x74,
	0xdf, 0x0b, 0xd3, 0x27, 0x8e, 0x38, 0x86, 0x1f, 0x81, 0xc6, 0x38, 0x4a, 0x79, 0xb8, 0x24, 0x7e,
	0x92, 0xd2, 0x84, 0x32, 0x14, 0xf9, 0x61, 0xa0, 0xab, 0x23, 0x75, 0x5c, 0x77, 0xcd, 0xc3, 0xd6,
	0x1c, 0xae, 0x50, 0x1c, 0xbd, 0xb2, 0x2e, 0xb5, 0x2c, 0x0f, 0x96, 0xf1, 0x54, 0xa6, 0x1f, 0x02,
	0xf8, 0x1a, 0xb4, 0x02, 0x9c, 0x50, 0x16, 0x72, 0xa6, 0xdf, 0x19, 0xd5, 0xc6, 0x9d, 0xc9, 0xc0,
	0x3e, 0x8d, 0x6b, 0xbf, 0x11, 0x67, 0x6e, 0x7f, 0xbd, 0x35, 0x95, 0xbf, 0x57, 0x66, 0x4b, 0x06,
	0xcc, 0x3b, 0x62, 0xf0, 0x25, 0x68, 0x64, 0x94, 0x63, 0xa6, 0xd7, 0x0a, 0xbe, 0x5f, 0xe5, 0xbf,
	0x50, 0x8e, 0xdd, 0x9e, 0x84, 0x1b, 0xf9, 0x1f, 0xf3, 0x44, 0x1b, 0xbe, 0x05, 0xed, 0x72, 0x3a,
	0xa6, 0xd7, 0x0b, 0x54, 0xab, 0xa2, 0xe5, 0x90, 0xee, 0x7d, 0x89, 0xb7, 0xcb, 0x84, 0x79, 0x27,
	0x12, 0xfa, 0xe0, 0x9e, 0x9c, 0xc4, 0x4f, 0x50, 0x8a, 0x62, 0xa6, 0x37, 0x46, 0xea, 0xb8, 0x33,
	0x79, 0x74, 0xe1, 0x1a, 0xd3, 0xa2, 0xe0, 0x3e, 0xcd, 0x85, 0x87, 0xad, 0xf9, 0x40, 0x2c, 0xeb,
	0x1c, 0xb7, 0xbc, 0x5e, 0x50, 0x6d, 0xc3, 0xef, 0xa0, 0x97, 0x51, 0xb1, 0x4c, 0xe1, 0x6f, 0x16,
	0x7e, 0xfd, 0xd6, 0x35, 0xf3, 0xb5, 0x0a, 0xfd, 0x13, 0xa9, 0xd7, 0x84, 0xfe, 0x0c, 0xb6, 0xbc,
	0x6e, 0x56, 0xe9, 0xc2, 0xaf, 0xa0, 0xcb, 0x51, 0x14, 0xad, 0x4a, 0xf7, 0xdd, 0xc2, 0xfd, 0xb0,
	0xea, 0xfe, 0x9c, 0x9f, 0x4b, 0xf5, 0x50, 0xaa, 0x07, 0x42, 0x5d, 0x45, 0x2d, 0xaf, 0xc3, 0x2b,
	0xcd, 0x77, 0xeb, 0x9d, 0xa1, 0x6e, 0x76, 0x86, 0x7a, 0xbd, 0x33, 0xd4, 0x3f, 0x7b, 0x43, 0xd9,
	0xec, 0x0d, 0xe5, 0xff, 0xde, 0x50, 0xbe, 0x3d, 0x23, 0x21, 0x5f, 0xfc, 0x9c, 0xd9, 0x73, 0x1a,
	0x3b, 0x31, 0xe2, 0xe1, 0x7c, 0x89, 0xf9, 0x2f, 0x9a, 0xfe, 0x70, 0x16, 0x38, 0x8c, 0x03, 0x14,
	0x45, 0xce, 0x6f, 0x27, 0x7f, 0x8e, 0x7c, 0x95, 0x60, 0x36, 0x6b, 0x16, 0x4f, 0xf1, 0xc5, 0x4d,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xe0, 0x48, 0x63, 0xde, 0x02, 0x00, 0x00,
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
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
	}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.DepositParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.VotingParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
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
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
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
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
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
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
