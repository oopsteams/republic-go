// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Nothing
	Address
	MultiAddress
	Query
	Order
	OrderId
	OrderFragment
	OrderFragmentId
	SyncRequest
	SyncBlock
	SignOrderFragmentRequest
	OpenOrderRequest
	CancelOrderRequest
	SmpcMessage
	GenerateRandomShares
	GenerateXiShares
	GenerateXiFragments
	RhoSigmaFragments
	DeltaFragments
	Delta
	DeltaFragment
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Nothing struct {
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Address struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MultiAddress struct {
	Signature    []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	MultiAddress string `protobuf:"bytes,2,opt,name=multiAddress" json:"multiAddress,omitempty"`
}

func (m *MultiAddress) Reset()                    { *m = MultiAddress{} }
func (m *MultiAddress) String() string            { return proto.CompactTextString(m) }
func (*MultiAddress) ProtoMessage()               {}
func (*MultiAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MultiAddress) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *MultiAddress) GetMultiAddress() string {
	if m != nil {
		return m.MultiAddress
	}
	return ""
}

type Query struct {
	From   *MultiAddress `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Target *Address      `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Query) GetFrom() *MultiAddress {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Query) GetTarget() *Address {
	if m != nil {
		return m.Target
	}
	return nil
}

type Order struct {
	Id     *OrderId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type   int64    `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Parity int64    `protobuf:"varint,3,opt,name=parity" json:"parity,omitempty"`
	Expiry int64    `protobuf:"varint,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Order) GetId() *OrderId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Order) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Order) GetParity() int64 {
	if m != nil {
		return m.Parity
	}
	return 0
}

func (m *Order) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type OrderId struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	OrderId   []byte `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (m *OrderId) Reset()                    { *m = OrderId{} }
func (m *OrderId) String() string            { return proto.CompactTextString(m) }
func (*OrderId) ProtoMessage()               {}
func (*OrderId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OrderId) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *OrderId) GetOrderId() []byte {
	if m != nil {
		return m.OrderId
	}
	return nil
}

type OrderFragment struct {
	Id             *OrderFragmentId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Order          *Order           `protobuf:"bytes,2,opt,name=order" json:"order,omitempty"`
	FstCodeShare   []byte           `protobuf:"bytes,3,opt,name=fstCodeShare,proto3" json:"fstCodeShare,omitempty"`
	SndCodeShare   []byte           `protobuf:"bytes,4,opt,name=sndCodeShare,proto3" json:"sndCodeShare,omitempty"`
	PriceShare     []byte           `protobuf:"bytes,5,opt,name=priceShare,proto3" json:"priceShare,omitempty"`
	MaxVolumeShare []byte           `protobuf:"bytes,6,opt,name=maxVolumeShare,proto3" json:"maxVolumeShare,omitempty"`
	MinVolumeShare []byte           `protobuf:"bytes,7,opt,name=minVolumeShare,proto3" json:"minVolumeShare,omitempty"`
}

func (m *OrderFragment) Reset()                    { *m = OrderFragment{} }
func (m *OrderFragment) String() string            { return proto.CompactTextString(m) }
func (*OrderFragment) ProtoMessage()               {}
func (*OrderFragment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OrderFragment) GetId() *OrderFragmentId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OrderFragment) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderFragment) GetFstCodeShare() []byte {
	if m != nil {
		return m.FstCodeShare
	}
	return nil
}

func (m *OrderFragment) GetSndCodeShare() []byte {
	if m != nil {
		return m.SndCodeShare
	}
	return nil
}

func (m *OrderFragment) GetPriceShare() []byte {
	if m != nil {
		return m.PriceShare
	}
	return nil
}

func (m *OrderFragment) GetMaxVolumeShare() []byte {
	if m != nil {
		return m.MaxVolumeShare
	}
	return nil
}

func (m *OrderFragment) GetMinVolumeShare() []byte {
	if m != nil {
		return m.MinVolumeShare
	}
	return nil
}

type OrderFragmentId struct {
	Signature       []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	OrderFragmentId []byte `protobuf:"bytes,2,opt,name=orderFragmentId,proto3" json:"orderFragmentId,omitempty"`
}

func (m *OrderFragmentId) Reset()                    { *m = OrderFragmentId{} }
func (m *OrderFragmentId) String() string            { return proto.CompactTextString(m) }
func (*OrderFragmentId) ProtoMessage()               {}
func (*OrderFragmentId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *OrderFragmentId) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *OrderFragmentId) GetOrderFragmentId() []byte {
	if m != nil {
		return m.OrderFragmentId
	}
	return nil
}

type SyncRequest struct {
	From *MultiAddress `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
}

func (m *SyncRequest) Reset()                    { *m = SyncRequest{} }
func (m *SyncRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()               {}
func (*SyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SyncRequest) GetFrom() *MultiAddress {
	if m != nil {
		return m.From
	}
	return nil
}

type SyncBlock struct {
}

func (m *SyncBlock) Reset()                    { *m = SyncBlock{} }
func (m *SyncBlock) String() string            { return proto.CompactTextString(m) }
func (*SyncBlock) ProtoMessage()               {}
func (*SyncBlock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SignOrderFragmentRequest struct {
	From            *MultiAddress    `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	OrderFragmentId *OrderFragmentId `protobuf:"bytes,2,opt,name=orderFragmentId" json:"orderFragmentId,omitempty"`
}

func (m *SignOrderFragmentRequest) Reset()                    { *m = SignOrderFragmentRequest{} }
func (m *SignOrderFragmentRequest) String() string            { return proto.CompactTextString(m) }
func (*SignOrderFragmentRequest) ProtoMessage()               {}
func (*SignOrderFragmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SignOrderFragmentRequest) GetFrom() *MultiAddress {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *SignOrderFragmentRequest) GetOrderFragmentId() *OrderFragmentId {
	if m != nil {
		return m.OrderFragmentId
	}
	return nil
}

type OpenOrderRequest struct {
	From          *MultiAddress  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	OrderFragment *OrderFragment `protobuf:"bytes,2,opt,name=orderFragment" json:"orderFragment,omitempty"`
}

func (m *OpenOrderRequest) Reset()                    { *m = OpenOrderRequest{} }
func (m *OpenOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenOrderRequest) ProtoMessage()               {}
func (*OpenOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *OpenOrderRequest) GetFrom() *MultiAddress {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *OpenOrderRequest) GetOrderFragment() *OrderFragment {
	if m != nil {
		return m.OrderFragment
	}
	return nil
}

type CancelOrderRequest struct {
	From            *MultiAddress    `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	OrderFragmentId *OrderFragmentId `protobuf:"bytes,2,opt,name=orderFragmentId" json:"orderFragmentId,omitempty"`
}

func (m *CancelOrderRequest) Reset()                    { *m = CancelOrderRequest{} }
func (m *CancelOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelOrderRequest) ProtoMessage()               {}
func (*CancelOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CancelOrderRequest) GetFrom() *MultiAddress {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *CancelOrderRequest) GetOrderFragmentId() *OrderFragmentId {
	if m != nil {
		return m.OrderFragmentId
	}
	return nil
}

type SmpcMessage struct {
	MultiAddress         *MultiAddress         `protobuf:"bytes,1,opt,name=multiAddress" json:"multiAddress,omitempty"`
	GenerateRandomShares *GenerateRandomShares `protobuf:"bytes,2,opt,name=generateRandomShares" json:"generateRandomShares,omitempty"`
	GenerateXiShares     *GenerateXiShares     `protobuf:"bytes,3,opt,name=generateXiShares" json:"generateXiShares,omitempty"`
	GenerateXiFragments  *GenerateXiFragments  `protobuf:"bytes,4,opt,name=generateXiFragments" json:"generateXiFragments,omitempty"`
	RhoSigmaFragments    *RhoSigmaFragments    `protobuf:"bytes,5,opt,name=rhoSigmaFragments" json:"rhoSigmaFragments,omitempty"`
	DeltaFragments       *DeltaFragments       `protobuf:"bytes,6,opt,name=deltaFragments" json:"deltaFragments,omitempty"`
}

func (m *SmpcMessage) Reset()                    { *m = SmpcMessage{} }
func (m *SmpcMessage) String() string            { return proto.CompactTextString(m) }
func (*SmpcMessage) ProtoMessage()               {}
func (*SmpcMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SmpcMessage) GetMultiAddress() *MultiAddress {
	if m != nil {
		return m.MultiAddress
	}
	return nil
}

func (m *SmpcMessage) GetGenerateRandomShares() *GenerateRandomShares {
	if m != nil {
		return m.GenerateRandomShares
	}
	return nil
}

func (m *SmpcMessage) GetGenerateXiShares() *GenerateXiShares {
	if m != nil {
		return m.GenerateXiShares
	}
	return nil
}

func (m *SmpcMessage) GetGenerateXiFragments() *GenerateXiFragments {
	if m != nil {
		return m.GenerateXiFragments
	}
	return nil
}

func (m *SmpcMessage) GetRhoSigmaFragments() *RhoSigmaFragments {
	if m != nil {
		return m.RhoSigmaFragments
	}
	return nil
}

func (m *SmpcMessage) GetDeltaFragments() *DeltaFragments {
	if m != nil {
		return m.DeltaFragments
	}
	return nil
}

type GenerateRandomShares struct {
}

func (m *GenerateRandomShares) Reset()                    { *m = GenerateRandomShares{} }
func (m *GenerateRandomShares) String() string            { return proto.CompactTextString(m) }
func (*GenerateRandomShares) ProtoMessage()               {}
func (*GenerateRandomShares) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type GenerateXiShares struct {
}

func (m *GenerateXiShares) Reset()                    { *m = GenerateXiShares{} }
func (m *GenerateXiShares) String() string            { return proto.CompactTextString(m) }
func (*GenerateXiShares) ProtoMessage()               {}
func (*GenerateXiShares) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type GenerateXiFragments struct {
}

func (m *GenerateXiFragments) Reset()                    { *m = GenerateXiFragments{} }
func (m *GenerateXiFragments) String() string            { return proto.CompactTextString(m) }
func (*GenerateXiFragments) ProtoMessage()               {}
func (*GenerateXiFragments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type RhoSigmaFragments struct {
}

func (m *RhoSigmaFragments) Reset()                    { *m = RhoSigmaFragments{} }
func (m *RhoSigmaFragments) String() string            { return proto.CompactTextString(m) }
func (*RhoSigmaFragments) ProtoMessage()               {}
func (*RhoSigmaFragments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type DeltaFragments struct {
	Signature      []byte           `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	DeltaFragments []*DeltaFragment `protobuf:"bytes,2,rep,name=deltaFragments" json:"deltaFragments,omitempty"`
}

func (m *DeltaFragments) Reset()                    { *m = DeltaFragments{} }
func (m *DeltaFragments) String() string            { return proto.CompactTextString(m) }
func (*DeltaFragments) ProtoMessage()               {}
func (*DeltaFragments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *DeltaFragments) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *DeltaFragments) GetDeltaFragments() []*DeltaFragment {
	if m != nil {
		return m.DeltaFragments
	}
	return nil
}

type Delta struct {
	Signature      []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Id             []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	BuyOrderId     []byte `protobuf:"bytes,3,opt,name=buyOrderId,proto3" json:"buyOrderId,omitempty"`
	SellOrderId    []byte `protobuf:"bytes,4,opt,name=sellOrderId,proto3" json:"sellOrderId,omitempty"`
	FstCode        []byte `protobuf:"bytes,5,opt,name=fstCode,proto3" json:"fstCode,omitempty"`
	SndCode        []byte `protobuf:"bytes,6,opt,name=sndCode,proto3" json:"sndCode,omitempty"`
	Price          []byte `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
	MaxVolumeShare []byte `protobuf:"bytes,8,opt,name=maxVolumeShare,proto3" json:"maxVolumeShare,omitempty"`
	MinVolumeShare []byte `protobuf:"bytes,9,opt,name=minVolumeShare,proto3" json:"minVolumeShare,omitempty"`
}

func (m *Delta) Reset()                    { *m = Delta{} }
func (m *Delta) String() string            { return proto.CompactTextString(m) }
func (*Delta) ProtoMessage()               {}
func (*Delta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Delta) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Delta) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Delta) GetBuyOrderId() []byte {
	if m != nil {
		return m.BuyOrderId
	}
	return nil
}

func (m *Delta) GetSellOrderId() []byte {
	if m != nil {
		return m.SellOrderId
	}
	return nil
}

func (m *Delta) GetFstCode() []byte {
	if m != nil {
		return m.FstCode
	}
	return nil
}

func (m *Delta) GetSndCode() []byte {
	if m != nil {
		return m.SndCode
	}
	return nil
}

func (m *Delta) GetPrice() []byte {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Delta) GetMaxVolumeShare() []byte {
	if m != nil {
		return m.MaxVolumeShare
	}
	return nil
}

func (m *Delta) GetMinVolumeShare() []byte {
	if m != nil {
		return m.MinVolumeShare
	}
	return nil
}

type DeltaFragment struct {
	Signature           []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Id                  []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	DeltaId             []byte `protobuf:"bytes,3,opt,name=deltaId,proto3" json:"deltaId,omitempty"`
	BuyOrderId          []byte `protobuf:"bytes,4,opt,name=buyOrderId,proto3" json:"buyOrderId,omitempty"`
	SellOrderId         []byte `protobuf:"bytes,5,opt,name=sellOrderId,proto3" json:"sellOrderId,omitempty"`
	BuyOrderFragmentId  []byte `protobuf:"bytes,6,opt,name=buyOrderFragmentId,proto3" json:"buyOrderFragmentId,omitempty"`
	SellOrderFragmentId []byte `protobuf:"bytes,7,opt,name=sellOrderFragmentId,proto3" json:"sellOrderFragmentId,omitempty"`
	FstCodeShare        []byte `protobuf:"bytes,8,opt,name=fstCodeShare,proto3" json:"fstCodeShare,omitempty"`
	SndCodeShare        []byte `protobuf:"bytes,9,opt,name=sndCodeShare,proto3" json:"sndCodeShare,omitempty"`
	PriceShare          []byte `protobuf:"bytes,10,opt,name=priceShare,proto3" json:"priceShare,omitempty"`
	MaxVolumeShare      []byte `protobuf:"bytes,11,opt,name=maxVolumeShare,proto3" json:"maxVolumeShare,omitempty"`
	MinVolumeShare      []byte `protobuf:"bytes,12,opt,name=minVolumeShare,proto3" json:"minVolumeShare,omitempty"`
}

func (m *DeltaFragment) Reset()                    { *m = DeltaFragment{} }
func (m *DeltaFragment) String() string            { return proto.CompactTextString(m) }
func (*DeltaFragment) ProtoMessage()               {}
func (*DeltaFragment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *DeltaFragment) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *DeltaFragment) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeltaFragment) GetDeltaId() []byte {
	if m != nil {
		return m.DeltaId
	}
	return nil
}

func (m *DeltaFragment) GetBuyOrderId() []byte {
	if m != nil {
		return m.BuyOrderId
	}
	return nil
}

func (m *DeltaFragment) GetSellOrderId() []byte {
	if m != nil {
		return m.SellOrderId
	}
	return nil
}

func (m *DeltaFragment) GetBuyOrderFragmentId() []byte {
	if m != nil {
		return m.BuyOrderFragmentId
	}
	return nil
}

func (m *DeltaFragment) GetSellOrderFragmentId() []byte {
	if m != nil {
		return m.SellOrderFragmentId
	}
	return nil
}

func (m *DeltaFragment) GetFstCodeShare() []byte {
	if m != nil {
		return m.FstCodeShare
	}
	return nil
}

func (m *DeltaFragment) GetSndCodeShare() []byte {
	if m != nil {
		return m.SndCodeShare
	}
	return nil
}

func (m *DeltaFragment) GetPriceShare() []byte {
	if m != nil {
		return m.PriceShare
	}
	return nil
}

func (m *DeltaFragment) GetMaxVolumeShare() []byte {
	if m != nil {
		return m.MaxVolumeShare
	}
	return nil
}

func (m *DeltaFragment) GetMinVolumeShare() []byte {
	if m != nil {
		return m.MinVolumeShare
	}
	return nil
}

func init() {
	proto.RegisterType((*Nothing)(nil), "rpc.Nothing")
	proto.RegisterType((*Address)(nil), "rpc.Address")
	proto.RegisterType((*MultiAddress)(nil), "rpc.MultiAddress")
	proto.RegisterType((*Query)(nil), "rpc.Query")
	proto.RegisterType((*Order)(nil), "rpc.Order")
	proto.RegisterType((*OrderId)(nil), "rpc.OrderId")
	proto.RegisterType((*OrderFragment)(nil), "rpc.OrderFragment")
	proto.RegisterType((*OrderFragmentId)(nil), "rpc.OrderFragmentId")
	proto.RegisterType((*SyncRequest)(nil), "rpc.SyncRequest")
	proto.RegisterType((*SyncBlock)(nil), "rpc.SyncBlock")
	proto.RegisterType((*SignOrderFragmentRequest)(nil), "rpc.SignOrderFragmentRequest")
	proto.RegisterType((*OpenOrderRequest)(nil), "rpc.OpenOrderRequest")
	proto.RegisterType((*CancelOrderRequest)(nil), "rpc.CancelOrderRequest")
	proto.RegisterType((*SmpcMessage)(nil), "rpc.SmpcMessage")
	proto.RegisterType((*GenerateRandomShares)(nil), "rpc.GenerateRandomShares")
	proto.RegisterType((*GenerateXiShares)(nil), "rpc.GenerateXiShares")
	proto.RegisterType((*GenerateXiFragments)(nil), "rpc.GenerateXiFragments")
	proto.RegisterType((*RhoSigmaFragments)(nil), "rpc.RhoSigmaFragments")
	proto.RegisterType((*DeltaFragments)(nil), "rpc.DeltaFragments")
	proto.RegisterType((*Delta)(nil), "rpc.Delta")
	proto.RegisterType((*DeltaFragment)(nil), "rpc.DeltaFragment")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Swarm service

type SwarmClient interface {
	Ping(ctx context.Context, in *MultiAddress, opts ...grpc.CallOption) (*MultiAddress, error)
	QueryPeers(ctx context.Context, in *Query, opts ...grpc.CallOption) (Swarm_QueryPeersClient, error)
	QueryPeersDeep(ctx context.Context, in *Query, opts ...grpc.CallOption) (Swarm_QueryPeersDeepClient, error)
}

type swarmClient struct {
	cc *grpc.ClientConn
}

func NewSwarmClient(cc *grpc.ClientConn) SwarmClient {
	return &swarmClient{cc}
}

func (c *swarmClient) Ping(ctx context.Context, in *MultiAddress, opts ...grpc.CallOption) (*MultiAddress, error) {
	out := new(MultiAddress)
	err := grpc.Invoke(ctx, "/rpc.Swarm/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmClient) QueryPeers(ctx context.Context, in *Query, opts ...grpc.CallOption) (Swarm_QueryPeersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Swarm_serviceDesc.Streams[0], c.cc, "/rpc.Swarm/QueryPeers", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmQueryPeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Swarm_QueryPeersClient interface {
	Recv() (*MultiAddress, error)
	grpc.ClientStream
}

type swarmQueryPeersClient struct {
	grpc.ClientStream
}

func (x *swarmQueryPeersClient) Recv() (*MultiAddress, error) {
	m := new(MultiAddress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmClient) QueryPeersDeep(ctx context.Context, in *Query, opts ...grpc.CallOption) (Swarm_QueryPeersDeepClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Swarm_serviceDesc.Streams[1], c.cc, "/rpc.Swarm/QueryPeersDeep", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmQueryPeersDeepClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Swarm_QueryPeersDeepClient interface {
	Recv() (*MultiAddress, error)
	grpc.ClientStream
}

type swarmQueryPeersDeepClient struct {
	grpc.ClientStream
}

func (x *swarmQueryPeersDeepClient) Recv() (*MultiAddress, error) {
	m := new(MultiAddress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Swarm service

type SwarmServer interface {
	Ping(context.Context, *MultiAddress) (*MultiAddress, error)
	QueryPeers(*Query, Swarm_QueryPeersServer) error
	QueryPeersDeep(*Query, Swarm_QueryPeersDeepServer) error
}

func RegisterSwarmServer(s *grpc.Server, srv SwarmServer) {
	s.RegisterService(&_Swarm_serviceDesc, srv)
}

func _Swarm_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Swarm/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServer).Ping(ctx, req.(*MultiAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Swarm_QueryPeers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).QueryPeers(m, &swarmQueryPeersServer{stream})
}

type Swarm_QueryPeersServer interface {
	Send(*MultiAddress) error
	grpc.ServerStream
}

type swarmQueryPeersServer struct {
	grpc.ServerStream
}

func (x *swarmQueryPeersServer) Send(m *MultiAddress) error {
	return x.ServerStream.SendMsg(m)
}

func _Swarm_QueryPeersDeep_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).QueryPeersDeep(m, &swarmQueryPeersDeepServer{stream})
}

type Swarm_QueryPeersDeepServer interface {
	Send(*MultiAddress) error
	grpc.ServerStream
}

type swarmQueryPeersDeepServer struct {
	grpc.ServerStream
}

func (x *swarmQueryPeersDeepServer) Send(m *MultiAddress) error {
	return x.ServerStream.SendMsg(m)
}

var _Swarm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Swarm",
	HandlerType: (*SwarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Swarm_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryPeers",
			Handler:       _Swarm_QueryPeers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryPeersDeep",
			Handler:       _Swarm_QueryPeersDeep_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// Client API for Orderbook service

type OrderbookClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Orderbook_SyncClient, error)
	SignOrderFragment(ctx context.Context, in *OrderFragmentId, opts ...grpc.CallOption) (*OrderFragmentId, error)
	OpenOrder(ctx context.Context, in *OpenOrderRequest, opts ...grpc.CallOption) (*Nothing, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type orderbookClient struct {
	cc *grpc.ClientConn
}

func NewOrderbookClient(cc *grpc.ClientConn) OrderbookClient {
	return &orderbookClient{cc}
}

func (c *orderbookClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Orderbook_SyncClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Orderbook_serviceDesc.Streams[0], c.cc, "/rpc.Orderbook/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderbookSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Orderbook_SyncClient interface {
	Recv() (*SyncBlock, error)
	grpc.ClientStream
}

type orderbookSyncClient struct {
	grpc.ClientStream
}

func (x *orderbookSyncClient) Recv() (*SyncBlock, error) {
	m := new(SyncBlock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderbookClient) SignOrderFragment(ctx context.Context, in *OrderFragmentId, opts ...grpc.CallOption) (*OrderFragmentId, error) {
	out := new(OrderFragmentId)
	err := grpc.Invoke(ctx, "/rpc.Orderbook/SignOrderFragment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderbookClient) OpenOrder(ctx context.Context, in *OpenOrderRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/rpc.Orderbook/OpenOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderbookClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/rpc.Orderbook/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orderbook service

type OrderbookServer interface {
	Sync(*SyncRequest, Orderbook_SyncServer) error
	SignOrderFragment(context.Context, *OrderFragmentId) (*OrderFragmentId, error)
	OpenOrder(context.Context, *OpenOrderRequest) (*Nothing, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*Nothing, error)
}

func RegisterOrderbookServer(s *grpc.Server, srv OrderbookServer) {
	s.RegisterService(&_Orderbook_serviceDesc, srv)
}

func _Orderbook_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderbookServer).Sync(m, &orderbookSyncServer{stream})
}

type Orderbook_SyncServer interface {
	Send(*SyncBlock) error
	grpc.ServerStream
}

type orderbookSyncServer struct {
	grpc.ServerStream
}

func (x *orderbookSyncServer) Send(m *SyncBlock) error {
	return x.ServerStream.SendMsg(m)
}

func _Orderbook_SignOrderFragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFragmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderbookServer).SignOrderFragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Orderbook/SignOrderFragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderbookServer).SignOrderFragment(ctx, req.(*OrderFragmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderbook_OpenOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderbookServer).OpenOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Orderbook/OpenOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderbookServer).OpenOrder(ctx, req.(*OpenOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderbook_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderbookServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Orderbook/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderbookServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orderbook_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Orderbook",
	HandlerType: (*OrderbookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignOrderFragment",
			Handler:    _Orderbook_SignOrderFragment_Handler,
		},
		{
			MethodName: "OpenOrder",
			Handler:    _Orderbook_OpenOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Orderbook_CancelOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Orderbook_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// Client API for Smpc service

type SmpcClient interface {
	Compute(ctx context.Context, opts ...grpc.CallOption) (Smpc_ComputeClient, error)
}

type smpcClient struct {
	cc *grpc.ClientConn
}

func NewSmpcClient(cc *grpc.ClientConn) SmpcClient {
	return &smpcClient{cc}
}

func (c *smpcClient) Compute(ctx context.Context, opts ...grpc.CallOption) (Smpc_ComputeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Smpc_serviceDesc.Streams[0], c.cc, "/rpc.Smpc/Compute", opts...)
	if err != nil {
		return nil, err
	}
	x := &smpcComputeClient{stream}
	return x, nil
}

type Smpc_ComputeClient interface {
	Send(*SmpcMessage) error
	Recv() (*SmpcMessage, error)
	grpc.ClientStream
}

type smpcComputeClient struct {
	grpc.ClientStream
}

func (x *smpcComputeClient) Send(m *SmpcMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smpcComputeClient) Recv() (*SmpcMessage, error) {
	m := new(SmpcMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Smpc service

type SmpcServer interface {
	Compute(Smpc_ComputeServer) error
}

func RegisterSmpcServer(s *grpc.Server, srv SmpcServer) {
	s.RegisterService(&_Smpc_serviceDesc, srv)
}

func _Smpc_Compute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmpcServer).Compute(&smpcComputeServer{stream})
}

type Smpc_ComputeServer interface {
	Send(*SmpcMessage) error
	Recv() (*SmpcMessage, error)
	grpc.ServerStream
}

type smpcComputeServer struct {
	grpc.ServerStream
}

func (x *smpcComputeServer) Send(m *SmpcMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smpcComputeServer) Recv() (*SmpcMessage, error) {
	m := new(SmpcMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Smpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Smpc",
	HandlerType: (*SmpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Compute",
			Handler:       _Smpc_Compute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0x96, 0x93, 0x38, 0x59, 0x8f, 0xd3, 0x6c, 0x3b, 0xe9, 0x16, 0x13, 0xad, 0x56, 0xd1, 0x61,
	0x41, 0x15, 0x82, 0x2a, 0x1b, 0x16, 0x89, 0x1f, 0x09, 0x54, 0x5a, 0x81, 0x40, 0x2a, 0x5b, 0x4e,
	0x10, 0x82, 0x4b, 0x37, 0x3e, 0xeb, 0x9a, 0x8d, 0x7f, 0x38, 0x76, 0xc4, 0x46, 0x5c, 0xf1, 0x04,
	0x88, 0x2b, 0x9e, 0x80, 0x27, 0xe0, 0x85, 0x78, 0x14, 0xe4, 0xf1, 0x71, 0xfc, 0xbb, 0x4a, 0xca,
	0xde, 0xe5, 0x7c, 0xf3, 0xcd, 0x77, 0x66, 0xc6, 0xce, 0x37, 0x06, 0x43, 0x46, 0xcb, 0xb3, 0x48,
	0x86, 0x49, 0x88, 0x5d, 0x19, 0x2d, 0x99, 0x01, 0x83, 0x6f, 0xc3, 0xe4, 0xd6, 0x0b, 0x5c, 0xf6,
	0x16, 0x0c, 0xce, 0x1d, 0x47, 0x8a, 0x38, 0x46, 0x0b, 0x06, 0x76, 0xf6, 0xd3, 0xd2, 0xa6, 0xda,
	0xa9, 0xc1, 0xf3, 0x23, 0xbb, 0x86, 0xe1, 0xd5, 0x7a, 0x95, 0x78, 0x39, 0xf3, 0x21, 0x18, 0xb1,
	0xe7, 0x06, 0x76, 0xb2, 0x96, 0x82, 0xb8, 0x43, 0x5e, 0x00, 0xc8, 0x60, 0xe8, 0x97, 0xd8, 0x56,
	0x87, 0xc4, 0x2a, 0x18, 0xfb, 0x1e, 0xf4, 0xef, 0xd6, 0x42, 0x6e, 0xf0, 0x6d, 0xe8, 0x3d, 0x97,
	0xa1, 0x4f, 0x2a, 0xe6, 0xfc, 0xe8, 0x2c, 0xad, 0xb4, 0x7c, 0x17, 0xa7, 0x30, 0x3e, 0x86, 0x7e,
	0x62, 0x4b, 0x57, 0x24, 0xa4, 0x66, 0xce, 0x87, 0x44, 0xcc, 0x39, 0x2a, 0xc6, 0x3c, 0xd0, 0x9f,
	0x49, 0x47, 0x48, 0x7c, 0x08, 0x1d, 0xcf, 0x51, 0x9a, 0x19, 0x95, 0xf0, 0xaf, 0x1d, 0xde, 0xf1,
	0x1c, 0x44, 0xe8, 0x25, 0x9b, 0x48, 0x90, 0x54, 0x97, 0xd3, 0x6f, 0x3c, 0x81, 0x7e, 0x64, 0x4b,
	0x2f, 0xd9, 0x58, 0x5d, 0x42, 0xd5, 0x29, 0xc5, 0xc5, 0xcb, 0xc8, 0x93, 0x1b, 0xab, 0x97, 0xe1,
	0xd9, 0x89, 0x9d, 0xc3, 0x40, 0x49, 0xee, 0x98, 0x86, 0x05, 0x83, 0x30, 0x23, 0xd2, 0x7d, 0x43,
	0x9e, 0x1f, 0xd9, 0x1f, 0x1d, 0x38, 0x20, 0x8d, 0x2f, 0xa5, 0xed, 0xfa, 0x22, 0x48, 0xf0, 0x71,
	0xa9, 0xec, 0xe3, 0xa2, 0xec, 0x3c, 0xae, 0xca, 0x9f, 0x82, 0x4e, 0x12, 0x6a, 0x14, 0x50, 0x10,
	0x79, 0x16, 0x48, 0x9f, 0xc0, 0xf3, 0x38, 0xb9, 0x08, 0x1d, 0xb1, 0xb8, 0xb5, 0xa5, 0xa0, 0x96,
	0x86, 0xbc, 0x82, 0xa5, 0x9c, 0x38, 0x70, 0x0a, 0x4e, 0x2f, 0xe3, 0x94, 0x31, 0x7c, 0x04, 0x10,
	0x49, 0x6f, 0xa9, 0x18, 0x3a, 0x31, 0x4a, 0x08, 0xbe, 0x03, 0x23, 0xdf, 0x7e, 0xf9, 0x43, 0xb8,
	0x5a, 0xfb, 0x8a, 0xd3, 0x27, 0x4e, 0x0d, 0x25, 0x9e, 0x17, 0x94, 0x79, 0x03, 0xc5, 0xab, 0xa0,
	0xec, 0x27, 0xb8, 0x5f, 0x6b, 0x78, 0xc7, 0x70, 0x4f, 0xe1, 0x7e, 0x58, 0x4d, 0x50, 0x43, 0xae,
	0xc3, 0xec, 0x29, 0x98, 0x8b, 0x4d, 0xb0, 0xe4, 0xe2, 0x97, 0xb5, 0x88, 0x93, 0x3d, 0x5f, 0x3b,
	0x66, 0x82, 0x91, 0x66, 0x7d, 0xb1, 0x0a, 0x97, 0x2f, 0xd8, 0xef, 0x1a, 0x58, 0x0b, 0xcf, 0x0d,
	0x2a, 0x25, 0xde, 0x4d, 0x10, 0x3f, 0x6b, 0x2f, 0xf8, 0x55, 0x8f, 0xbb, 0xd1, 0x46, 0x0c, 0x87,
	0xcf, 0x22, 0x91, 0x95, 0x70, 0xc7, 0xab, 0x3f, 0x82, 0x83, 0x8a, 0x9a, 0xba, 0x18, 0x9b, 0x17,
	0xf3, 0x2a, 0x91, 0xfd, 0x06, 0x78, 0x61, 0x07, 0x4b, 0xb1, 0xfa, 0x3f, 0xd7, 0xbe, 0x6e, 0xc7,
	0x7f, 0x77, 0xc1, 0x5c, 0xf8, 0xd1, 0xf2, 0x4a, 0xc4, 0xb1, 0xed, 0x0a, 0xfc, 0xb0, 0xe6, 0x2e,
	0xaf, 0xbc, 0xbe, 0x42, 0xc3, 0x2b, 0x38, 0x76, 0x45, 0x20, 0xa4, 0x9d, 0x08, 0x6e, 0x07, 0x4e,
	0xe8, 0xd3, 0x1b, 0x17, 0xab, 0x5a, 0xde, 0xa4, 0xf4, 0xaf, 0x5a, 0x08, 0xbc, 0x35, 0x0d, 0xcf,
	0xe1, 0x30, 0xc7, 0x7f, 0xf4, 0x94, 0x54, 0x97, 0xa4, 0x1e, 0x54, 0xa4, 0xf2, 0x20, 0x6f, 0xd0,
	0xf1, 0x1b, 0x18, 0x17, 0x58, 0xde, 0x70, 0x4c, 0xff, 0x43, 0x73, 0x6e, 0xd5, 0x54, 0xb6, 0x71,
	0xde, 0x96, 0x84, 0x97, 0x70, 0x24, 0x6f, 0xc3, 0x85, 0xe7, 0xfa, 0x76, 0xa1, 0xa4, 0x93, 0xd2,
	0x09, 0x29, 0xf1, 0x7a, 0x94, 0x37, 0x13, 0xf0, 0x53, 0x18, 0x39, 0x62, 0x95, 0x94, 0x24, 0xfa,
	0x24, 0x31, 0x26, 0x89, 0xcb, 0x4a, 0x88, 0xd7, 0xa8, 0xec, 0x04, 0x8e, 0xdb, 0xe6, 0xc7, 0x10,
	0x0e, 0xeb, 0xc3, 0x60, 0x0f, 0x60, 0xdc, 0xd2, 0x1a, 0x1b, 0xc3, 0x51, 0xa3, 0x4e, 0xf6, 0x33,
	0x8c, 0xaa, 0x37, 0xef, 0xb0, 0x84, 0x4f, 0x1a, 0x4d, 0x74, 0xa6, 0xdd, 0xed, 0x7b, 0x5e, 0x91,
	0x6a, 0xf4, 0xf0, 0x67, 0x07, 0x74, 0x62, 0xec, 0xb8, 0x63, 0x44, 0x3e, 0x9d, 0x39, 0x4d, 0xea,
	0xc8, 0x8f, 0x00, 0x6e, 0xd6, 0x1b, 0xb5, 0x0f, 0x94, 0xdb, 0x96, 0x10, 0x9c, 0x82, 0x19, 0x8b,
	0xd5, 0x2a, 0x27, 0x64, 0x56, 0x5b, 0x86, 0xd2, 0x2d, 0xa1, 0xdc, 0x59, 0xd9, 0x6c, 0x7e, 0x4c,
	0x23, 0xca, 0x93, 0x95, 0xb9, 0xe6, 0x47, 0x3c, 0x06, 0x9d, 0xbc, 0x58, 0x99, 0x69, 0x76, 0x68,
	0xf1, 0xe4, 0x7b, 0x7b, 0x7a, 0xb2, 0xd1, 0xea, 0xc9, 0xff, 0x74, 0xe1, 0xa0, 0x32, 0xb5, 0x3b,
	0xce, 0xc6, 0x82, 0x01, 0x4d, 0x79, 0x3b, 0x98, 0xfc, 0x58, 0x9b, 0x5a, 0x6f, 0xd7, 0xd4, 0xf4,
	0xe6, 0xd4, 0xce, 0x00, 0x73, 0x7e, 0xc9, 0x5e, 0xb2, 0x31, 0xb5, 0x44, 0x70, 0x06, 0xe3, 0x6d,
	0x7a, 0x29, 0x21, 0x9b, 0x5f, 0x5b, 0xa8, 0xb1, 0x49, 0xef, 0xed, 0xb1, 0x49, 0x8d, 0x9d, 0x9b,
	0x14, 0xf6, 0xd8, 0xa4, 0xe6, 0x9e, 0x4f, 0x6d, 0xd8, 0xf6, 0xd4, 0xe6, 0x7f, 0x69, 0xa0, 0x2f,
	0x7e, 0xb5, 0xa5, 0x8f, 0xef, 0x41, 0xef, 0xda, 0x0b, 0x5c, 0x6c, 0x3a, 0xe4, 0xa4, 0x09, 0xe1,
	0xfb, 0x00, 0xf4, 0x5d, 0x76, 0x2d, 0x84, 0x8c, 0x31, 0xfb, 0xb4, 0x20, 0xa0, 0x85, 0x3c, 0xd3,
	0xf0, 0x09, 0x8c, 0x0a, 0xfa, 0xa5, 0x10, 0xd1, 0xce, 0x94, 0xf9, 0xbf, 0x1a, 0x18, 0x34, 0xe5,
	0x9b, 0x30, 0x7c, 0x81, 0xef, 0x42, 0x2f, 0x5d, 0xb0, 0x78, 0x48, 0xd4, 0xd2, 0x86, 0x9e, 0x8c,
	0xb6, 0x08, 0x6d, 0xdf, 0x99, 0x86, 0x9f, 0xc3, 0x51, 0x63, 0xfd, 0x62, 0xeb, 0x16, 0x99, 0xb4,
	0xa2, 0x38, 0x03, 0x63, 0xbb, 0x3c, 0x31, 0xf3, 0xe9, 0xfa, 0x32, 0x9d, 0x64, 0x5f, 0x8b, 0xea,
	0xeb, 0x18, 0x9f, 0x82, 0x59, 0xda, 0x7c, 0xf8, 0x06, 0x05, 0x9b, 0xbb, 0xb0, 0x9a, 0x35, 0xff,
	0x18, 0x7a, 0xe9, 0xc6, 0xc2, 0x27, 0x30, 0xb8, 0x08, 0xfd, 0x68, 0x9d, 0x88, 0xbc, 0xbf, 0x62,
	0x8f, 0x4d, 0x1a, 0xc8, 0xa9, 0x36, 0xd3, 0x6e, 0xfa, 0xf4, 0x95, 0xfe, 0xc1, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0xb5, 0x0a, 0x1d, 0xb2, 0x0b, 0x00, 0x00,
}
