// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/genesis.proto

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

// GenesisState defines the pki module's genesis state.
type GenesisState struct {
	ApprovedCertificatesList          []ApprovedCertificates          `protobuf:"bytes,1,rep,name=approvedCertificatesList,proto3" json:"approvedCertificatesList"`
	ProposedCertificateList           []ProposedCertificate           `protobuf:"bytes,2,rep,name=proposedCertificateList,proto3" json:"proposedCertificateList"`
	ChildCertificatesList             []ChildCertificates             `protobuf:"bytes,3,rep,name=childCertificatesList,proto3" json:"childCertificatesList"`
	ProposedCertificateRevocationList []ProposedCertificateRevocation `protobuf:"bytes,4,rep,name=proposedCertificateRevocationList,proto3" json:"proposedCertificateRevocationList"`
	RevokedCertificatesList           []RevokedCertificates           `protobuf:"bytes,5,rep,name=revokedCertificatesList,proto3" json:"revokedCertificatesList"`
	UniqueCertificateList             []UniqueCertificate             `protobuf:"bytes,6,rep,name=uniqueCertificateList,proto3" json:"uniqueCertificateList"`
	ApprovedRootCertificates          *ApprovedRootCertificates       `protobuf:"bytes,7,opt,name=approvedRootCertificates,proto3" json:"approvedRootCertificates,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9478608499b59120, []int{0}
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

func (m *GenesisState) GetApprovedCertificatesList() []ApprovedCertificates {
	if m != nil {
		return m.ApprovedCertificatesList
	}
	return nil
}

func (m *GenesisState) GetProposedCertificateList() []ProposedCertificate {
	if m != nil {
		return m.ProposedCertificateList
	}
	return nil
}

func (m *GenesisState) GetChildCertificatesList() []ChildCertificates {
	if m != nil {
		return m.ChildCertificatesList
	}
	return nil
}

func (m *GenesisState) GetProposedCertificateRevocationList() []ProposedCertificateRevocation {
	if m != nil {
		return m.ProposedCertificateRevocationList
	}
	return nil
}

func (m *GenesisState) GetRevokedCertificatesList() []RevokedCertificates {
	if m != nil {
		return m.RevokedCertificatesList
	}
	return nil
}

func (m *GenesisState) GetUniqueCertificateList() []UniqueCertificate {
	if m != nil {
		return m.UniqueCertificateList
	}
	return nil
}

func (m *GenesisState) GetApprovedRootCertificates() *ApprovedRootCertificates {
	if m != nil {
		return m.ApprovedRootCertificates
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zigbeealliance.distributedcomplianceledger.pki.GenesisState")
}

func init() { proto.RegisterFile("pki/genesis.proto", fileDescriptor_9478608499b59120) }

var fileDescriptor_9478608499b59120 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xbf, 0x6a, 0xdb, 0x40,
	0x00, 0xc6, 0x75, 0xb5, 0x6b, 0x83, 0xdc, 0xa5, 0xa2, 0xa5, 0xc6, 0x14, 0xd9, 0x2d, 0x1d, 0xdc,
	0xc1, 0x12, 0xb8, 0x4f, 0xe0, 0x3f, 0xd0, 0x0e, 0x2d, 0x14, 0x95, 0x2e, 0x1d, 0x6a, 0xf4, 0xe7,
	0x22, 0x1f, 0x96, 0x75, 0x97, 0xd3, 0xc9, 0x24, 0x81, 0x4c, 0x09, 0x99, 0xf3, 0x0a, 0x79, 0x1b,
	0x0f, 0x19, 0x3c, 0x66, 0x0a, 0xc1, 0x7e, 0x91, 0xa0, 0x93, 0x14, 0xcb, 0xd2, 0x89, 0x60, 0x93,
	0x4d, 0xdc, 0x9f, 0xef, 0xfb, 0xe9, 0xfb, 0xee, 0x4e, 0x7e, 0x4b, 0x66, 0x48, 0x77, 0xa1, 0x0f,
	0x03, 0x14, 0x68, 0x84, 0x62, 0x86, 0x15, 0xed, 0x0c, 0xb9, 0x16, 0x84, 0xa6, 0xe7, 0x21, 0xd3,
	0xb7, 0xa1, 0xe6, 0xa0, 0x80, 0x51, 0x64, 0x85, 0x0c, 0x3a, 0x36, 0x9e, 0x93, 0x78, 0xd4, 0x83,
	0x8e, 0x0b, 0xa9, 0x46, 0x66, 0xa8, 0xd5, 0x8e, 0x24, 0x4c, 0x42, 0x28, 0x5e, 0x40, 0x67, 0x62,
	0x43, 0xca, 0xd0, 0x11, 0xb2, 0x4d, 0x06, 0x13, 0xc1, 0x96, 0x1a, 0x2d, 0x20, 0x14, 0x13, 0x1c,
	0xec, 0x2e, 0x48, 0xe6, 0x3f, 0x46, 0xf3, 0xf6, 0x14, 0x79, 0xc2, 0xdd, 0x5f, 0xcb, 0x76, 0x4f,
	0x28, 0x5c, 0x60, 0xdb, 0x64, 0x08, 0xfb, 0x59, 0xa3, 0x68, 0x74, 0x26, 0x06, 0xe1, 0x46, 0xa1,
	0x8f, 0x8e, 0x43, 0x28, 0xc0, 0xf8, 0xb2, 0xf3, 0x1f, 0x14, 0x63, 0x26, 0xd2, 0x78, 0xe7, 0x62,
	0x17, 0xf3, 0x4f, 0x3d, 0xfa, 0x8a, 0x47, 0x3f, 0xdf, 0xd6, 0xe5, 0x37, 0xdf, 0xe3, 0x14, 0xff,
	0x30, 0x93, 0x41, 0xe5, 0x0a, 0xc8, 0xcd, 0x54, 0x6b, 0x94, 0x51, 0xf9, 0x89, 0x02, 0xd6, 0x04,
	0x9d, 0x4a, 0xb7, 0xd1, 0x1f, 0xef, 0x19, 0xb4, 0x36, 0x10, 0xe8, 0x0d, 0xab, 0xcb, 0xfb, 0xb6,
	0x64, 0x94, 0x7a, 0x29, 0x17, 0x40, 0xfe, 0x90, 0xa6, 0x97, 0x99, 0xe4, 0x1c, 0xaf, 0x38, 0xc7,
	0x68, 0x5f, 0x8e, 0xdf, 0x45, 0xb9, 0x04, 0xa3, 0xcc, 0x49, 0x39, 0x97, 0xdf, 0xf3, 0x82, 0x0b,
	0x51, 0x54, 0x38, 0xc2, 0x60, 0x5f, 0x84, 0x51, 0x5e, 0x2c, 0x01, 0x10, 0xbb, 0x28, 0x37, 0x40,
	0xfe, 0x24, 0x40, 0x33, 0x9e, 0x0e, 0x10, 0x67, 0xa9, 0x72, 0x96, 0x5f, 0x2f, 0x10, 0xc7, 0x56,
	0x38, 0xe1, 0x7a, 0xde, 0x9d, 0x17, 0x95, 0x9c, 0xdd, 0x42, 0x4a, 0xaf, 0x0f, 0x2b, 0xca, 0x28,
	0xca, 0xa5, 0x45, 0x95, 0x38, 0x45, 0x45, 0xc5, 0x17, 0x24, 0x7f, 0x56, 0x6a, 0x87, 0x15, 0xf5,
	0x37, 0x2f, 0x96, 0x16, 0x25, 0x74, 0x51, 0x2e, 0x33, 0xd7, 0xc6, 0xc0, 0x98, 0x65, 0xf9, 0x9a,
	0xf5, 0x0e, 0xe8, 0x36, 0xfa, 0x3f, 0x0e, 0xbd, 0x36, 0x79, 0x3d, 0xa3, 0xd4, 0x69, 0xf8, 0x7f,
	0xb9, 0x56, 0xc1, 0x6a, 0xad, 0x82, 0x87, 0xb5, 0x0a, 0xae, 0x37, 0xaa, 0xb4, 0xda, 0xa8, 0xd2,
	0xdd, 0x46, 0x95, 0xfe, 0x8d, 0x5d, 0xc4, 0xa6, 0xa1, 0xa5, 0xd9, 0x78, 0xae, 0xc7, 0x1c, 0xbd,
	0x14, 0x44, 0xcf, 0x80, 0xf4, 0xb6, 0x24, 0xbd, 0x18, 0x45, 0x3f, 0xd1, 0xa3, 0xf7, 0x85, 0x9d,
	0x12, 0x18, 0x58, 0x35, 0xfe, 0x6a, 0x7c, 0x7b, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x73, 0x6f, 0x2e,
	0x97, 0x7e, 0x05, 0x00, 0x00,
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
	if m.ApprovedRootCertificates != nil {
		{
			size, err := m.ApprovedRootCertificates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.UniqueCertificateList) > 0 {
		for iNdEx := len(m.UniqueCertificateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UniqueCertificateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.RevokedCertificatesList) > 0 {
		for iNdEx := len(m.RevokedCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RevokedCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ProposedCertificateRevocationList) > 0 {
		for iNdEx := len(m.ProposedCertificateRevocationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposedCertificateRevocationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ChildCertificatesList) > 0 {
		for iNdEx := len(m.ChildCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChildCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ProposedCertificateList) > 0 {
		for iNdEx := len(m.ProposedCertificateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposedCertificateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ApprovedCertificatesList) > 0 {
		for iNdEx := len(m.ApprovedCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApprovedCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.ApprovedCertificatesList) > 0 {
		for _, e := range m.ApprovedCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposedCertificateList) > 0 {
		for _, e := range m.ProposedCertificateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChildCertificatesList) > 0 {
		for _, e := range m.ChildCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposedCertificateRevocationList) > 0 {
		for _, e := range m.ProposedCertificateRevocationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RevokedCertificatesList) > 0 {
		for _, e := range m.RevokedCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UniqueCertificateList) > 0 {
		for _, e := range m.UniqueCertificateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ApprovedRootCertificates != nil {
		l = m.ApprovedRootCertificates.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedCertificatesList", wireType)
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
			m.ApprovedCertificatesList = append(m.ApprovedCertificatesList, ApprovedCertificates{})
			if err := m.ApprovedCertificatesList[len(m.ApprovedCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedCertificateList", wireType)
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
			m.ProposedCertificateList = append(m.ProposedCertificateList, ProposedCertificate{})
			if err := m.ProposedCertificateList[len(m.ProposedCertificateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildCertificatesList", wireType)
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
			m.ChildCertificatesList = append(m.ChildCertificatesList, ChildCertificates{})
			if err := m.ChildCertificatesList[len(m.ChildCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedCertificateRevocationList", wireType)
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
			m.ProposedCertificateRevocationList = append(m.ProposedCertificateRevocationList, ProposedCertificateRevocation{})
			if err := m.ProposedCertificateRevocationList[len(m.ProposedCertificateRevocationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokedCertificatesList", wireType)
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
			m.RevokedCertificatesList = append(m.RevokedCertificatesList, RevokedCertificates{})
			if err := m.RevokedCertificatesList[len(m.RevokedCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniqueCertificateList", wireType)
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
			m.UniqueCertificateList = append(m.UniqueCertificateList, UniqueCertificate{})
			if err := m.UniqueCertificateList[len(m.UniqueCertificateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedRootCertificates", wireType)
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
			if m.ApprovedRootCertificates == nil {
				m.ApprovedRootCertificates = &ApprovedRootCertificates{}
			}
			if err := m.ApprovedRootCertificates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
