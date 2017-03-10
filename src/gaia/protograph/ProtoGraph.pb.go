// Code generated by protoc-gen-go.
// source: src/gaia/protograph/ProtoGraph.proto
// DO NOT EDIT!

/*
Package protograph is a generated protocol buffer package.

It is generated from these files:
	src/gaia/protograph/ProtoGraph.proto

It has these top-level messages:
	SingleEdge
	RepeatedEdges
	EmbeddedEdges
	RenameProperty
	SerializeField
	SpliceMap
	InnerVertex
	JoinList
	StoreField
	FieldAction
	TransformMessage
*/
package protograph

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SingleEdge struct {
	EdgeLabel        string `protobuf:"bytes,1,opt,name=edge_label,json=edgeLabel" json:"edge_label,omitempty"`
	DestinationLabel string `protobuf:"bytes,2,opt,name=destination_label,json=destinationLabel" json:"destination_label,omitempty"`
}

func (m *SingleEdge) Reset()                    { *m = SingleEdge{} }
func (m *SingleEdge) String() string            { return proto.CompactTextString(m) }
func (*SingleEdge) ProtoMessage()               {}
func (*SingleEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SingleEdge) GetEdgeLabel() string {
	if m != nil {
		return m.EdgeLabel
	}
	return ""
}

func (m *SingleEdge) GetDestinationLabel() string {
	if m != nil {
		return m.DestinationLabel
	}
	return ""
}

type RepeatedEdges struct {
	EdgeLabel        string `protobuf:"bytes,1,opt,name=edge_label,json=edgeLabel" json:"edge_label,omitempty"`
	DestinationLabel string `protobuf:"bytes,2,opt,name=destination_label,json=destinationLabel" json:"destination_label,omitempty"`
}

func (m *RepeatedEdges) Reset()                    { *m = RepeatedEdges{} }
func (m *RepeatedEdges) String() string            { return proto.CompactTextString(m) }
func (*RepeatedEdges) ProtoMessage()               {}
func (*RepeatedEdges) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepeatedEdges) GetEdgeLabel() string {
	if m != nil {
		return m.EdgeLabel
	}
	return ""
}

func (m *RepeatedEdges) GetDestinationLabel() string {
	if m != nil {
		return m.DestinationLabel
	}
	return ""
}

type EmbeddedEdges struct {
	EdgeLabel        string `protobuf:"bytes,1,opt,name=edge_label,json=edgeLabel" json:"edge_label,omitempty"`
	DestinationLabel string `protobuf:"bytes,2,opt,name=destination_label,json=destinationLabel" json:"destination_label,omitempty"`
	EmbeddedIn       string `protobuf:"bytes,3,opt,name=embedded_in,json=embeddedIn" json:"embedded_in,omitempty"`
}

func (m *EmbeddedEdges) Reset()                    { *m = EmbeddedEdges{} }
func (m *EmbeddedEdges) String() string            { return proto.CompactTextString(m) }
func (*EmbeddedEdges) ProtoMessage()               {}
func (*EmbeddedEdges) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EmbeddedEdges) GetEdgeLabel() string {
	if m != nil {
		return m.EdgeLabel
	}
	return ""
}

func (m *EmbeddedEdges) GetDestinationLabel() string {
	if m != nil {
		return m.DestinationLabel
	}
	return ""
}

func (m *EmbeddedEdges) GetEmbeddedIn() string {
	if m != nil {
		return m.EmbeddedIn
	}
	return ""
}

type RenameProperty struct {
	Rename string `protobuf:"bytes,1,opt,name=rename" json:"rename,omitempty"`
}

func (m *RenameProperty) Reset()                    { *m = RenameProperty{} }
func (m *RenameProperty) String() string            { return proto.CompactTextString(m) }
func (*RenameProperty) ProtoMessage()               {}
func (*RenameProperty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RenameProperty) GetRename() string {
	if m != nil {
		return m.Rename
	}
	return ""
}

type SerializeField struct {
	SerializedName string `protobuf:"bytes,1,opt,name=serialized_name,json=serializedName" json:"serialized_name,omitempty"`
}

func (m *SerializeField) Reset()                    { *m = SerializeField{} }
func (m *SerializeField) String() string            { return proto.CompactTextString(m) }
func (*SerializeField) ProtoMessage()               {}
func (*SerializeField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SerializeField) GetSerializedName() string {
	if m != nil {
		return m.SerializedName
	}
	return ""
}

type SpliceMap struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *SpliceMap) Reset()                    { *m = SpliceMap{} }
func (m *SpliceMap) String() string            { return proto.CompactTextString(m) }
func (*SpliceMap) ProtoMessage()               {}
func (*SpliceMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SpliceMap) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type InnerVertex struct {
	EdgeLabel        string `protobuf:"bytes,1,opt,name=edge_label,json=edgeLabel" json:"edge_label,omitempty"`
	DestinationLabel string `protobuf:"bytes,2,opt,name=destination_label,json=destinationLabel" json:"destination_label,omitempty"`
	OuterId          string `protobuf:"bytes,3,opt,name=outer_id,json=outerId" json:"outer_id,omitempty"`
}

func (m *InnerVertex) Reset()                    { *m = InnerVertex{} }
func (m *InnerVertex) String() string            { return proto.CompactTextString(m) }
func (*InnerVertex) ProtoMessage()               {}
func (*InnerVertex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InnerVertex) GetEdgeLabel() string {
	if m != nil {
		return m.EdgeLabel
	}
	return ""
}

func (m *InnerVertex) GetDestinationLabel() string {
	if m != nil {
		return m.DestinationLabel
	}
	return ""
}

func (m *InnerVertex) GetOuterId() string {
	if m != nil {
		return m.OuterId
	}
	return ""
}

type JoinList struct {
	Delimiter string `protobuf:"bytes,1,opt,name=delimiter" json:"delimiter,omitempty"`
}

func (m *JoinList) Reset()                    { *m = JoinList{} }
func (m *JoinList) String() string            { return proto.CompactTextString(m) }
func (*JoinList) ProtoMessage()               {}
func (*JoinList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *JoinList) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

type StoreField struct {
	Store bool `protobuf:"varint,1,opt,name=store" json:"store,omitempty"`
}

func (m *StoreField) Reset()                    { *m = StoreField{} }
func (m *StoreField) String() string            { return proto.CompactTextString(m) }
func (*StoreField) ProtoMessage()               {}
func (*StoreField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StoreField) GetStore() bool {
	if m != nil {
		return m.Store
	}
	return false
}

type FieldAction struct {
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*FieldAction_SingleEdge
	//	*FieldAction_RepeatedEdges
	//	*FieldAction_EmbeddedEdges
	//	*FieldAction_RenameProperty
	//	*FieldAction_SerializeField
	//	*FieldAction_SpliceMap
	//	*FieldAction_InnerVertex
	//	*FieldAction_JoinList
	//	*FieldAction_StoreField
	Action isFieldAction_Action `protobuf_oneof:"action"`
}

func (m *FieldAction) Reset()                    { *m = FieldAction{} }
func (m *FieldAction) String() string            { return proto.CompactTextString(m) }
func (*FieldAction) ProtoMessage()               {}
func (*FieldAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type isFieldAction_Action interface {
	isFieldAction_Action()
}

type FieldAction_SingleEdge struct {
	SingleEdge *SingleEdge `protobuf:"bytes,3,opt,name=single_edge,json=singleEdge,oneof"`
}
type FieldAction_RepeatedEdges struct {
	RepeatedEdges *RepeatedEdges `protobuf:"bytes,4,opt,name=repeated_edges,json=repeatedEdges,oneof"`
}
type FieldAction_EmbeddedEdges struct {
	EmbeddedEdges *EmbeddedEdges `protobuf:"bytes,5,opt,name=embedded_edges,json=embeddedEdges,oneof"`
}
type FieldAction_RenameProperty struct {
	RenameProperty *RenameProperty `protobuf:"bytes,6,opt,name=rename_property,json=renameProperty,oneof"`
}
type FieldAction_SerializeField struct {
	SerializeField *SerializeField `protobuf:"bytes,7,opt,name=serialize_field,json=serializeField,oneof"`
}
type FieldAction_SpliceMap struct {
	SpliceMap *SpliceMap `protobuf:"bytes,8,opt,name=splice_map,json=spliceMap,oneof"`
}
type FieldAction_InnerVertex struct {
	InnerVertex *InnerVertex `protobuf:"bytes,9,opt,name=inner_vertex,json=innerVertex,oneof"`
}
type FieldAction_JoinList struct {
	JoinList *JoinList `protobuf:"bytes,10,opt,name=join_list,json=joinList,oneof"`
}
type FieldAction_StoreField struct {
	StoreField *StoreField `protobuf:"bytes,11,opt,name=store_field,json=storeField,oneof"`
}

func (*FieldAction_SingleEdge) isFieldAction_Action()     {}
func (*FieldAction_RepeatedEdges) isFieldAction_Action()  {}
func (*FieldAction_EmbeddedEdges) isFieldAction_Action()  {}
func (*FieldAction_RenameProperty) isFieldAction_Action() {}
func (*FieldAction_SerializeField) isFieldAction_Action() {}
func (*FieldAction_SpliceMap) isFieldAction_Action()      {}
func (*FieldAction_InnerVertex) isFieldAction_Action()    {}
func (*FieldAction_JoinList) isFieldAction_Action()       {}
func (*FieldAction_StoreField) isFieldAction_Action()     {}

func (m *FieldAction) GetAction() isFieldAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *FieldAction) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *FieldAction) GetSingleEdge() *SingleEdge {
	if x, ok := m.GetAction().(*FieldAction_SingleEdge); ok {
		return x.SingleEdge
	}
	return nil
}

func (m *FieldAction) GetRepeatedEdges() *RepeatedEdges {
	if x, ok := m.GetAction().(*FieldAction_RepeatedEdges); ok {
		return x.RepeatedEdges
	}
	return nil
}

func (m *FieldAction) GetEmbeddedEdges() *EmbeddedEdges {
	if x, ok := m.GetAction().(*FieldAction_EmbeddedEdges); ok {
		return x.EmbeddedEdges
	}
	return nil
}

func (m *FieldAction) GetRenameProperty() *RenameProperty {
	if x, ok := m.GetAction().(*FieldAction_RenameProperty); ok {
		return x.RenameProperty
	}
	return nil
}

func (m *FieldAction) GetSerializeField() *SerializeField {
	if x, ok := m.GetAction().(*FieldAction_SerializeField); ok {
		return x.SerializeField
	}
	return nil
}

func (m *FieldAction) GetSpliceMap() *SpliceMap {
	if x, ok := m.GetAction().(*FieldAction_SpliceMap); ok {
		return x.SpliceMap
	}
	return nil
}

func (m *FieldAction) GetInnerVertex() *InnerVertex {
	if x, ok := m.GetAction().(*FieldAction_InnerVertex); ok {
		return x.InnerVertex
	}
	return nil
}

func (m *FieldAction) GetJoinList() *JoinList {
	if x, ok := m.GetAction().(*FieldAction_JoinList); ok {
		return x.JoinList
	}
	return nil
}

func (m *FieldAction) GetStoreField() *StoreField {
	if x, ok := m.GetAction().(*FieldAction_StoreField); ok {
		return x.StoreField
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FieldAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FieldAction_OneofMarshaler, _FieldAction_OneofUnmarshaler, _FieldAction_OneofSizer, []interface{}{
		(*FieldAction_SingleEdge)(nil),
		(*FieldAction_RepeatedEdges)(nil),
		(*FieldAction_EmbeddedEdges)(nil),
		(*FieldAction_RenameProperty)(nil),
		(*FieldAction_SerializeField)(nil),
		(*FieldAction_SpliceMap)(nil),
		(*FieldAction_InnerVertex)(nil),
		(*FieldAction_JoinList)(nil),
		(*FieldAction_StoreField)(nil),
	}
}

func _FieldAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FieldAction)
	// action
	switch x := m.Action.(type) {
	case *FieldAction_SingleEdge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SingleEdge); err != nil {
			return err
		}
	case *FieldAction_RepeatedEdges:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RepeatedEdges); err != nil {
			return err
		}
	case *FieldAction_EmbeddedEdges:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EmbeddedEdges); err != nil {
			return err
		}
	case *FieldAction_RenameProperty:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RenameProperty); err != nil {
			return err
		}
	case *FieldAction_SerializeField:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SerializeField); err != nil {
			return err
		}
	case *FieldAction_SpliceMap:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpliceMap); err != nil {
			return err
		}
	case *FieldAction_InnerVertex:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InnerVertex); err != nil {
			return err
		}
	case *FieldAction_JoinList:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinList); err != nil {
			return err
		}
	case *FieldAction_StoreField:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StoreField); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FieldAction.Action has unexpected type %T", x)
	}
	return nil
}

func _FieldAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FieldAction)
	switch tag {
	case 3: // action.single_edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SingleEdge)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_SingleEdge{msg}
		return true, err
	case 4: // action.repeated_edges
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RepeatedEdges)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_RepeatedEdges{msg}
		return true, err
	case 5: // action.embedded_edges
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EmbeddedEdges)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_EmbeddedEdges{msg}
		return true, err
	case 6: // action.rename_property
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RenameProperty)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_RenameProperty{msg}
		return true, err
	case 7: // action.serialize_field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SerializeField)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_SerializeField{msg}
		return true, err
	case 8: // action.splice_map
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpliceMap)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_SpliceMap{msg}
		return true, err
	case 9: // action.inner_vertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InnerVertex)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_InnerVertex{msg}
		return true, err
	case 10: // action.join_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinList)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_JoinList{msg}
		return true, err
	case 11: // action.store_field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StoreField)
		err := b.DecodeMessage(msg)
		m.Action = &FieldAction_StoreField{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FieldAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FieldAction)
	// action
	switch x := m.Action.(type) {
	case *FieldAction_SingleEdge:
		s := proto.Size(x.SingleEdge)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_RepeatedEdges:
		s := proto.Size(x.RepeatedEdges)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_EmbeddedEdges:
		s := proto.Size(x.EmbeddedEdges)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_RenameProperty:
		s := proto.Size(x.RenameProperty)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_SerializeField:
		s := proto.Size(x.SerializeField)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_SpliceMap:
		s := proto.Size(x.SpliceMap)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_InnerVertex:
		s := proto.Size(x.InnerVertex)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_JoinList:
		s := proto.Size(x.JoinList)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldAction_StoreField:
		s := proto.Size(x.StoreField)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TransformMessage struct {
	Label   string         `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Gid     string         `protobuf:"bytes,2,opt,name=gid" json:"gid,omitempty"`
	Actions []*FieldAction `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *TransformMessage) Reset()                    { *m = TransformMessage{} }
func (m *TransformMessage) String() string            { return proto.CompactTextString(m) }
func (*TransformMessage) ProtoMessage()               {}
func (*TransformMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TransformMessage) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *TransformMessage) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *TransformMessage) GetActions() []*FieldAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*SingleEdge)(nil), "gaia.schema.SingleEdge")
	proto.RegisterType((*RepeatedEdges)(nil), "gaia.schema.RepeatedEdges")
	proto.RegisterType((*EmbeddedEdges)(nil), "gaia.schema.EmbeddedEdges")
	proto.RegisterType((*RenameProperty)(nil), "gaia.schema.RenameProperty")
	proto.RegisterType((*SerializeField)(nil), "gaia.schema.SerializeField")
	proto.RegisterType((*SpliceMap)(nil), "gaia.schema.SpliceMap")
	proto.RegisterType((*InnerVertex)(nil), "gaia.schema.InnerVertex")
	proto.RegisterType((*JoinList)(nil), "gaia.schema.JoinList")
	proto.RegisterType((*StoreField)(nil), "gaia.schema.StoreField")
	proto.RegisterType((*FieldAction)(nil), "gaia.schema.FieldAction")
	proto.RegisterType((*TransformMessage)(nil), "gaia.schema.TransformMessage")
}

func init() { proto.RegisterFile("src/gaia/protograph/ProtoGraph.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0x53, 0xb7, 0xdd, 0x26, 0x37, 0x6d, 0x5a, 0x83, 0xd6, 0xf8, 0x0f, 0x4b, 0x14, 0x2c,
	0x08, 0xbb, 0x50, 0x05, 0x51, 0xf0, 0xc1, 0x96, 0xae, 0x59, 0x69, 0xa5, 0x64, 0x45, 0x44, 0x1f,
	0xc2, 0xec, 0xce, 0xdd, 0x74, 0x4a, 0x32, 0x09, 0x33, 0xa3, 0x54, 0xf1, 0x2b, 0xfb, 0x1d, 0x64,
	0x26, 0xc9, 0x6e, 0xa2, 0xfb, 0x58, 0xdf, 0xe6, 0xdc, 0x9c, 0x73, 0x26, 0x33, 0xfc, 0x12, 0x78,
	0x22, 0xc5, 0x6c, 0x98, 0x12, 0x46, 0x86, 0xa5, 0x28, 0x54, 0x91, 0x0a, 0x52, 0x5e, 0x0c, 0xcf,
	0xf5, 0xf2, 0x9d, 0x5e, 0x0e, 0xcc, 0xd4, 0x77, 0xb5, 0x63, 0x20, 0x67, 0x17, 0x98, 0x93, 0xf0,
	0x33, 0xc0, 0x84, 0xf1, 0x34, 0xc3, 0x13, 0x9a, 0xa2, 0xff, 0x10, 0x00, 0x69, 0x8a, 0x49, 0x46,
	0xa6, 0x98, 0x05, 0x6b, 0xfb, 0x6b, 0x07, 0x4e, 0xec, 0xe8, 0xc9, 0xa9, 0x1e, 0xf8, 0xcf, 0xe0,
	0x26, 0x45, 0xa9, 0x18, 0x27, 0x8a, 0x15, 0xbc, 0x76, 0xdd, 0x30, 0xae, 0xdd, 0xd6, 0x03, 0x63,
	0x0e, 0xbf, 0xc2, 0x76, 0x8c, 0x25, 0x12, 0x85, 0x54, 0x77, 0xcb, 0x6b, 0x2d, 0xff, 0x05, 0xdb,
	0x27, 0xf9, 0x14, 0x29, 0xfd, 0x0f, 0xe5, 0xfe, 0x23, 0x70, 0xb1, 0x2e, 0x4f, 0x18, 0x0f, 0x7a,
	0xc6, 0x06, 0xcd, 0x68, 0xcc, 0xc3, 0x03, 0xf0, 0x62, 0xe4, 0x24, 0xc7, 0x73, 0x51, 0x94, 0x28,
	0xd4, 0x0f, 0x7f, 0x0f, 0xfa, 0xc2, 0x4c, 0xea, 0xad, 0x6b, 0x15, 0xbe, 0x02, 0x6f, 0x82, 0x82,
	0x91, 0x8c, 0xfd, 0xc4, 0x11, 0xc3, 0x8c, 0xfa, 0x4f, 0x61, 0x47, 0x36, 0x13, 0x9a, 0xb4, 0x22,
	0xde, 0x72, 0xfc, 0x41, 0x47, 0x1f, 0x83, 0x33, 0x29, 0x33, 0x36, 0xc3, 0x33, 0x52, 0xea, 0xfe,
	0x52, 0xe0, 0x9c, 0x5d, 0x35, 0xfd, 0x95, 0x0a, 0x15, 0xb8, 0x63, 0xce, 0x51, 0x7c, 0x42, 0xa1,
	0xf0, 0xea, 0x5a, 0x6f, 0xe1, 0x2e, 0xd8, 0xc5, 0x37, 0x85, 0x22, 0x61, 0xb4, 0xbe, 0x82, 0x4d,
	0xa3, 0xc7, 0x34, 0x3c, 0x00, 0xfb, 0x7d, 0xc1, 0xf8, 0x29, 0x93, 0xca, 0x7f, 0x00, 0x0e, 0xc5,
	0x8c, 0xe5, 0x4c, 0xa1, 0x68, 0x76, 0x5c, 0x0c, 0xc2, 0x10, 0x60, 0xa2, 0x0a, 0x51, 0x9f, 0xfd,
	0x16, 0x6c, 0x48, 0xad, 0x8c, 0xcf, 0x8e, 0x2b, 0x11, 0xfe, 0x5e, 0x07, 0xd7, 0x3c, 0x7f, 0x3b,
	0xd3, 0xbb, 0x6b, 0xd7, 0x5c, 0xcb, 0xba, 0xad, 0x12, 0xfe, 0x6b, 0x70, 0xa5, 0x01, 0x35, 0xd1,
	0xe7, 0x31, 0x6f, 0xe4, 0x1e, 0xde, 0x19, 0xb4, 0x58, 0x1e, 0x2c, 0x41, 0x8e, 0xac, 0x18, 0xe4,
	0x12, 0xeb, 0x63, 0xf0, 0x44, 0x8d, 0xa2, 0x49, 0xcb, 0x60, 0xdd, 0xc4, 0xef, 0x75, 0xe2, 0x1d,
	0x5a, 0x23, 0x2b, 0xde, 0x16, 0x1d, 0x7c, 0x8f, 0xc1, 0x5b, 0x50, 0x51, 0x95, 0x6c, 0xac, 0x28,
	0xe9, 0x50, 0xa9, 0x4b, 0xb0, 0x83, 0xe9, 0x08, 0x76, 0x2a, 0x32, 0x92, 0xb2, 0x46, 0x27, 0xe8,
	0x9b, 0x96, 0xfb, 0x7f, 0xbd, 0x4a, 0x9b, 0xae, 0xc8, 0x8a, 0x3d, 0xd1, 0xe5, 0x6d, 0xd4, 0xa2,
	0x28, 0xa9, 0x6e, 0x6b, 0x73, 0x45, 0x4f, 0x97, 0x3d, 0xdd, 0x23, 0xbb, 0x34, 0xbe, 0x04, 0x90,
	0x06, 0xb2, 0x24, 0x27, 0x65, 0x60, 0x9b, 0x8a, 0xbd, 0x6e, 0x45, 0xc3, 0x60, 0x64, 0xc5, 0x8e,
	0x5c, 0x00, 0xf9, 0x06, 0xb6, 0x98, 0x06, 0x2f, 0xf9, 0x6e, 0xc8, 0x0b, 0x1c, 0x13, 0x0d, 0x3a,
	0xd1, 0x16, 0x99, 0x91, 0x15, 0xbb, 0xac, 0x05, 0xea, 0x0b, 0x70, 0x2e, 0x0b, 0xc6, 0x93, 0x8c,
	0x49, 0x15, 0x80, 0xc9, 0xde, 0xee, 0x64, 0x1b, 0xbe, 0x22, 0x2b, 0xb6, 0x2f, 0x1b, 0xd6, 0x34,
	0x03, 0x1a, 0x99, 0xfa, 0xc4, 0xee, 0x2a, 0x06, 0x16, 0xb4, 0x19, 0x06, 0x16, 0xea, 0xc8, 0x86,
	0x3e, 0x31, 0x7c, 0x85, 0x1c, 0x76, 0x3f, 0x0a, 0xc2, 0xe5, 0xbc, 0x10, 0xf9, 0x19, 0x4a, 0x49,
	0x52, 0xd4, 0xcc, 0xb5, 0xbf, 0x99, 0x4a, 0xf8, 0xbb, 0xd0, 0x4b, 0x19, 0xad, 0xbf, 0x10, 0xbd,
	0xf4, 0x0f, 0x61, 0xb3, 0x6a, 0x91, 0x41, 0x6f, 0xbf, 0xf7, 0xcf, 0x89, 0x5b, 0x18, 0xc7, 0x8d,
	0xf1, 0x68, 0xeb, 0x0b, 0x2c, 0x7f, 0xc7, 0xd3, 0xbe, 0x59, 0x3f, 0xff, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x0b, 0x38, 0x4e, 0xac, 0x05, 0x00, 0x00,
}
