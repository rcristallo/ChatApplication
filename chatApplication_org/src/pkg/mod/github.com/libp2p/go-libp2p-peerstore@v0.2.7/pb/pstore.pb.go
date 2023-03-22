// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pstore.proto

package pstore_pb

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

// AddrBookRecord represents a record for a peer in the address book.
type AddrBookRecord struct {
	// The peer ID.
	Id *ProtoPeerID `protobuf:"bytes,1,opt,name=id,proto3,customtype=ProtoPeerID" json:"id,omitempty"`
	// The multiaddresses. This is a sorted list where element 0 expires the soonest.
	Addrs []*AddrBookRecord_AddrEntry `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	// The most recently received signed PeerRecord.
	CertifiedRecord *AddrBookRecord_CertifiedRecord `protobuf:"bytes,3,opt,name=certified_record,json=certifiedRecord,proto3" json:"certified_record,omitempty"`
}

func (m *AddrBookRecord) Reset()         { *m = AddrBookRecord{} }
func (m *AddrBookRecord) String() string { return proto.CompactTextString(m) }
func (*AddrBookRecord) ProtoMessage()    {}
func (*AddrBookRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96873690e08a98f, []int{0}
}
func (m *AddrBookRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddrBookRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddrBookRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddrBookRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrBookRecord.Merge(m, src)
}
func (m *AddrBookRecord) XXX_Size() int {
	return m.Size()
}
func (m *AddrBookRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrBookRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AddrBookRecord proto.InternalMessageInfo

func (m *AddrBookRecord) GetAddrs() []*AddrBookRecord_AddrEntry {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *AddrBookRecord) GetCertifiedRecord() *AddrBookRecord_CertifiedRecord {
	if m != nil {
		return m.CertifiedRecord
	}
	return nil
}

// AddrEntry represents a single multiaddress.
type AddrBookRecord_AddrEntry struct {
	Addr *ProtoAddr `protobuf:"bytes,1,opt,name=addr,proto3,customtype=ProtoAddr" json:"addr,omitempty"`
	// The point in time when this address expires.
	Expiry int64 `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// The original TTL of this address.
	Ttl int64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (m *AddrBookRecord_AddrEntry) Reset()         { *m = AddrBookRecord_AddrEntry{} }
func (m *AddrBookRecord_AddrEntry) String() string { return proto.CompactTextString(m) }
func (*AddrBookRecord_AddrEntry) ProtoMessage()    {}
func (*AddrBookRecord_AddrEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96873690e08a98f, []int{0, 0}
}
func (m *AddrBookRecord_AddrEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddrBookRecord_AddrEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddrBookRecord_AddrEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddrBookRecord_AddrEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrBookRecord_AddrEntry.Merge(m, src)
}
func (m *AddrBookRecord_AddrEntry) XXX_Size() int {
	return m.Size()
}
func (m *AddrBookRecord_AddrEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrBookRecord_AddrEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AddrBookRecord_AddrEntry proto.InternalMessageInfo

func (m *AddrBookRecord_AddrEntry) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *AddrBookRecord_AddrEntry) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// CertifiedRecord contains a serialized signed PeerRecord used to
// populate the signedAddrs list.
type AddrBookRecord_CertifiedRecord struct {
	// The Seq counter from the signed PeerRecord envelope
	Seq uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// The serialized bytes of the SignedEnvelope containing the PeerRecord.
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (m *AddrBookRecord_CertifiedRecord) Reset()         { *m = AddrBookRecord_CertifiedRecord{} }
func (m *AddrBookRecord_CertifiedRecord) String() string { return proto.CompactTextString(m) }
func (*AddrBookRecord_CertifiedRecord) ProtoMessage()    {}
func (*AddrBookRecord_CertifiedRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96873690e08a98f, []int{0, 1}
}
func (m *AddrBookRecord_CertifiedRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddrBookRecord_CertifiedRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddrBookRecord_CertifiedRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddrBookRecord_CertifiedRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrBookRecord_CertifiedRecord.Merge(m, src)
}
func (m *AddrBookRecord_CertifiedRecord) XXX_Size() int {
	return m.Size()
}
func (m *AddrBookRecord_CertifiedRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrBookRecord_CertifiedRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AddrBookRecord_CertifiedRecord proto.InternalMessageInfo

func (m *AddrBookRecord_CertifiedRecord) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *AddrBookRecord_CertifiedRecord) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*AddrBookRecord)(nil), "pstore.pb.AddrBookRecord")
	proto.RegisterType((*AddrBookRecord_AddrEntry)(nil), "pstore.pb.AddrBookRecord.AddrEntry")
	proto.RegisterType((*AddrBookRecord_CertifiedRecord)(nil), "pstore.pb.AddrBookRecord.CertifiedRecord")
}

func init() { proto.RegisterFile("pstore.proto", fileDescriptor_f96873690e08a98f) }

var fileDescriptor_f96873690e08a98f = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0x2e, 0xc9,
	0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xf1, 0x92, 0xa4, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0x2a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xe8, 0x54, 0xba, 0xcc, 0xc4, 0xc5,
	0xe7, 0x98, 0x92, 0x52, 0xe4, 0x94, 0x9f, 0x9f, 0x1d, 0x94, 0x9a, 0x9c, 0x5f, 0x94, 0x22, 0x24,
	0xcf, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe3, 0xc4, 0x7f, 0xeb, 0x9e, 0x3c,
	0x77, 0x00, 0x48, 0x65, 0x40, 0x6a, 0x6a, 0x91, 0xa7, 0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x25,
	0x17, 0x6b, 0x62, 0x4a, 0x4a, 0x51, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xb2, 0x1e,
	0xdc, 0x76, 0x3d, 0x54, 0xa3, 0xc0, 0x5c, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x88, 0x0e, 0xa1,
	0x10, 0x2e, 0x81, 0xe4, 0xd4, 0xa2, 0x92, 0xcc, 0xb4, 0xcc, 0xd4, 0x94, 0xf8, 0x22, 0xb0, 0x22,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x4d, 0xdc, 0xa6, 0x38, 0xc3, 0x74, 0x40, 0xf8, 0x41,
	0xfc, 0xc9, 0xa8, 0x02, 0x52, 0x11, 0x5c, 0x9c, 0x70, 0x9b, 0x84, 0x14, 0xb9, 0x58, 0x40, 0x76,
	0x41, 0x3d, 0xc0, 0x7b, 0xeb, 0x9e, 0x3c, 0x27, 0xd8, 0x03, 0x20, 0x15, 0x41, 0x60, 0x29, 0x21,
	0x31, 0x2e, 0xb6, 0xd4, 0x8a, 0x82, 0xcc, 0xa2, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x28, 0x4f, 0x48, 0x80, 0x8b, 0xb9, 0xa4, 0x24, 0x07, 0xec, 0x20, 0xe6, 0x20, 0x10, 0x53, 0xca,
	0x94, 0x8b, 0x1f, 0xcd, 0x76, 0x90, 0xa2, 0xe2, 0xd4, 0x42, 0xb0, 0xf1, 0x2c, 0x41, 0x20, 0x26,
	0x48, 0xa4, 0x28, 0xb1, 0x1c, 0x6c, 0x16, 0x4f, 0x10, 0x88, 0xe9, 0xa4, 0xf0, 0xe3, 0xa1, 0x1c,
	0xe3, 0x81, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4,
	0x06, 0x0e, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xcd, 0xe8, 0xd4, 0xc8, 0x01,
	0x00, 0x00,
}

func (m *AddrBookRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddrBookRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddrBookRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertifiedRecord != nil {
		{
			size, err := m.CertifiedRecord.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPstore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Addrs) > 0 {
		for iNdEx := len(m.Addrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPstore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != nil {
		{
			size := m.Id.Size()
			i -= size
			if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPstore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddrBookRecord_AddrEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddrBookRecord_AddrEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddrBookRecord_AddrEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ttl != 0 {
		i = encodeVarintPstore(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x18
	}
	if m.Expiry != 0 {
		i = encodeVarintPstore(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x10
	}
	if m.Addr != nil {
		{
			size := m.Addr.Size()
			i -= size
			if _, err := m.Addr.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPstore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddrBookRecord_CertifiedRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddrBookRecord_CertifiedRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddrBookRecord_CertifiedRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Raw) > 0 {
		i -= len(m.Raw)
		copy(dAtA[i:], m.Raw)
		i = encodeVarintPstore(dAtA, i, uint64(len(m.Raw)))
		i--
		dAtA[i] = 0x12
	}
	if m.Seq != 0 {
		i = encodeVarintPstore(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPstore(dAtA []byte, offset int, v uint64) int {
	offset -= sovPstore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedAddrBookRecord(r randyPstore, easy bool) *AddrBookRecord {
	this := &AddrBookRecord{}
	this.Id = NewPopulatedProtoPeerID(r)
	if r.Intn(5) != 0 {
		v1 := r.Intn(5)
		this.Addrs = make([]*AddrBookRecord_AddrEntry, v1)
		for i := 0; i < v1; i++ {
			this.Addrs[i] = NewPopulatedAddrBookRecord_AddrEntry(r, easy)
		}
	}
	if r.Intn(5) != 0 {
		this.CertifiedRecord = NewPopulatedAddrBookRecord_CertifiedRecord(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAddrBookRecord_AddrEntry(r randyPstore, easy bool) *AddrBookRecord_AddrEntry {
	this := &AddrBookRecord_AddrEntry{}
	this.Addr = NewPopulatedProtoAddr(r)
	this.Expiry = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Expiry *= -1
	}
	this.Ttl = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Ttl *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAddrBookRecord_CertifiedRecord(r randyPstore, easy bool) *AddrBookRecord_CertifiedRecord {
	this := &AddrBookRecord_CertifiedRecord{}
	this.Seq = uint64(uint64(r.Uint32()))
	v2 := r.Intn(100)
	this.Raw = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Raw[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPstore interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePstore(r randyPstore) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPstore(r randyPstore) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RunePstore(r)
	}
	return string(tmps)
}
func randUnrecognizedPstore(r randyPstore, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPstore(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPstore(dAtA []byte, r randyPstore, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePstore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePstore(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AddrBookRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovPstore(uint64(l))
	}
	if len(m.Addrs) > 0 {
		for _, e := range m.Addrs {
			l = e.Size()
			n += 1 + l + sovPstore(uint64(l))
		}
	}
	if m.CertifiedRecord != nil {
		l = m.CertifiedRecord.Size()
		n += 1 + l + sovPstore(uint64(l))
	}
	return n
}

func (m *AddrBookRecord_AddrEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Addr != nil {
		l = m.Addr.Size()
		n += 1 + l + sovPstore(uint64(l))
	}
	if m.Expiry != 0 {
		n += 1 + sovPstore(uint64(m.Expiry))
	}
	if m.Ttl != 0 {
		n += 1 + sovPstore(uint64(m.Ttl))
	}
	return n
}

func (m *AddrBookRecord_CertifiedRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovPstore(uint64(m.Seq))
	}
	l = len(m.Raw)
	if l > 0 {
		n += 1 + l + sovPstore(uint64(l))
	}
	return n
}

func sovPstore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPstore(x uint64) (n int) {
	return sovPstore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddrBookRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPstore
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
			return fmt.Errorf("proto: AddrBookRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddrBookRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
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
				return ErrInvalidLengthPstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoPeerID
			m.Id = &v
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
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
				return ErrInvalidLengthPstore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, &AddrBookRecord_AddrEntry{})
			if err := m.Addrs[len(m.Addrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifiedRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
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
				return ErrInvalidLengthPstore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertifiedRecord == nil {
				m.CertifiedRecord = &AddrBookRecord_CertifiedRecord{}
			}
			if err := m.CertifiedRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPstore
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
func (m *AddrBookRecord_AddrEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPstore
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
			return fmt.Errorf("proto: AddrEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddrEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
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
				return ErrInvalidLengthPstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoAddr
			m.Addr = &v
			if err := m.Addr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPstore
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
func (m *AddrBookRecord_CertifiedRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPstore
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
			return fmt.Errorf("proto: CertifiedRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertifiedRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPstore
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
				return ErrInvalidLengthPstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPstore
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
func skipPstore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPstore
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
					return 0, ErrIntOverflowPstore
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
					return 0, ErrIntOverflowPstore
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
				return 0, ErrInvalidLengthPstore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPstore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPstore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPstore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPstore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPstore = fmt.Errorf("proto: unexpected end of group")
)
