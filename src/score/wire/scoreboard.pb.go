// Code generated by protoc-gen-go.
// source: scoreboard.proto
// DO NOT EDIT!

/*
Package wire is a generated protocol buffer package.

It is generated from these files:
	scoreboard.proto

It has these top-level messages:
	Message
	Event
	ContestSetup
	Problem
	Team
*/
package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SState int32

const (
	SState_CORRECT SState = 1
	SState_WRONG   SState = 2
	SState_PENDING SState = 3
	SState_FIRST   SState = 4
)

var SState_name = map[int32]string{
	1: "CORRECT",
	2: "WRONG",
	3: "PENDING",
	4: "FIRST",
}
var SState_value = map[string]int32{
	"CORRECT": 1,
	"WRONG":   2,
	"PENDING": 3,
	"FIRST":   4,
}

func (x SState) Enum() *SState {
	p := new(SState)
	*p = x
	return p
}
func (x SState) String() string {
	return proto.EnumName(SState_name, int32(x))
}
func (x *SState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SState_value, data, "SState")
	if err != nil {
		return err
	}
	*x = SState(value)
	return nil
}
func (SState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	// Types that are valid to be assigned to MessageType:
	//	*Message_Event
	//	*Message_Setup
	MessageType      isMessage_MessageType `protobuf_oneof:"MessageType"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMessage_MessageType interface {
	isMessage_MessageType()
}

type Message_Event struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,oneof"`
}
type Message_Setup struct {
	Setup *ContestSetup `protobuf:"bytes,2,opt,name=setup,oneof"`
}

func (*Message_Event) isMessage_MessageType() {}
func (*Message_Setup) isMessage_MessageType() {}

func (m *Message) GetMessageType() isMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *Message) GetEvent() *Event {
	if x, ok := m.GetMessageType().(*Message_Event); ok {
		return x.Event
	}
	return nil
}

func (m *Message) GetSetup() *ContestSetup {
	if x, ok := m.GetMessageType().(*Message_Setup); ok {
		return x.Setup
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Event)(nil),
		(*Message_Setup)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// MessageType
	switch x := m.MessageType.(type) {
	case *Message_Event:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Event); err != nil {
			return err
		}
	case *Message_Setup:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Setup); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.MessageType has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // MessageType.event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event)
		err := b.DecodeMessage(msg)
		m.MessageType = &Message_Event{msg}
		return true, err
	case 2: // MessageType.setup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContestSetup)
		err := b.DecodeMessage(msg)
		m.MessageType = &Message_Setup{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// MessageType
	switch x := m.MessageType.(type) {
	case *Message_Event:
		s := proto.Size(x.Event)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Setup:
		s := proto.Size(x.Setup)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Event struct {
	Team             *int64  `protobuf:"varint,6,req,name=Team" json:"Team,omitempty"`
	Problem          *int64  `protobuf:"varint,1,req,name=Problem" json:"Problem,omitempty"`
	SubmitCount      *int64  `protobuf:"varint,2,req,name=SubmitCount" json:"SubmitCount,omitempty"`
	Penalty          *int64  `protobuf:"varint,3,req,name=Penalty" json:"Penalty,omitempty"`
	ContestTime      *int64  `protobuf:"varint,7,opt,name=ContestTime" json:"ContestTime,omitempty"`
	State            *SState `protobuf:"varint,4,req,name=State,enum=wire.SState" json:"State,omitempty"`
	Unfrozen         *Event  `protobuf:"bytes,5,opt,name=Unfrozen" json:"Unfrozen,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetTeam() int64 {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return 0
}

func (m *Event) GetProblem() int64 {
	if m != nil && m.Problem != nil {
		return *m.Problem
	}
	return 0
}

func (m *Event) GetSubmitCount() int64 {
	if m != nil && m.SubmitCount != nil {
		return *m.SubmitCount
	}
	return 0
}

func (m *Event) GetPenalty() int64 {
	if m != nil && m.Penalty != nil {
		return *m.Penalty
	}
	return 0
}

func (m *Event) GetContestTime() int64 {
	if m != nil && m.ContestTime != nil {
		return *m.ContestTime
	}
	return 0
}

func (m *Event) GetState() SState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return SState_CORRECT
}

func (m *Event) GetUnfrozen() *Event {
	if m != nil {
		return m.Unfrozen
	}
	return nil
}

type ContestSetup struct {
	Name             *string    `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Teams            []*Team    `protobuf:"bytes,2,rep,name=Teams" json:"Teams,omitempty"`
	Problems         []*Problem `protobuf:"bytes,3,rep,name=Problems" json:"Problems,omitempty"`
	Start            *int64     `protobuf:"varint,4,req,name=Start" json:"Start,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ContestSetup) Reset()                    { *m = ContestSetup{} }
func (m *ContestSetup) String() string            { return proto.CompactTextString(m) }
func (*ContestSetup) ProtoMessage()               {}
func (*ContestSetup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContestSetup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ContestSetup) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *ContestSetup) GetProblems() []*Problem {
	if m != nil {
		return m.Problems
	}
	return nil
}

func (m *ContestSetup) GetStart() int64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

type Problem struct {
	Id               *int64  `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Label            *string `protobuf:"bytes,2,req,name=Label" json:"Label,omitempty"`
	Color            *string `protobuf:"bytes,3,opt,name=Color" json:"Color,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Problem) Reset()                    { *m = Problem{} }
func (m *Problem) String() string            { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()               {}
func (*Problem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Problem) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Problem) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Problem) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return ""
}

type Team struct {
	Id               *int64  `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=Name" json:"Name,omitempty"`
	Affiliation      *string `protobuf:"bytes,3,opt,name=Affiliation" json:"Affiliation,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Team) Reset()                    { *m = Team{} }
func (m *Team) String() string            { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()               {}
func (*Team) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Team) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Team) GetAffiliation() string {
	if m != nil && m.Affiliation != nil {
		return *m.Affiliation
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "wire.Message")
	proto.RegisterType((*Event)(nil), "wire.Event")
	proto.RegisterType((*ContestSetup)(nil), "wire.ContestSetup")
	proto.RegisterType((*Problem)(nil), "wire.Problem")
	proto.RegisterType((*Team)(nil), "wire.Team")
	proto.RegisterEnum("wire.SState", SState_name, SState_value)
}

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x6f, 0xda, 0x30,
	0x18, 0xc7, 0xc9, 0x1b, 0x21, 0x4f, 0x60, 0x8b, 0xbc, 0x4b, 0xa6, 0x6d, 0x1a, 0xca, 0x2e, 0x68,
	0x07, 0x0e, 0x4c, 0x93, 0x7a, 0x6d, 0x53, 0x5a, 0x90, 0x5a, 0x40, 0x49, 0xaa, 0x1e, 0x7a, 0x72,
	0xca, 0x03, 0x8a, 0x94, 0xc4, 0xc8, 0x31, 0xad, 0xe8, 0x77, 0xe9, 0x77, 0xad, 0x9d, 0x04, 0xd4,
	0xaa, 0x47, 0xff, 0x5f, 0xec, 0xdf, 0xdf, 0xe0, 0x55, 0x8f, 0x8c, 0x63, 0xca, 0x28, 0x5f, 0x8f,
	0x77, 0x9c, 0x09, 0x46, 0xcc, 0xe7, 0x8c, 0x63, 0xf0, 0x00, 0xf6, 0x2d, 0x56, 0x15, 0xdd, 0x22,
	0xf9, 0x09, 0x16, 0x3e, 0x61, 0x29, 0x7c, 0x6d, 0xa8, 0x8d, 0xdc, 0x89, 0x3b, 0x56, 0x81, 0xf1,
	0x54, 0x49, 0xb3, 0x0e, 0xf9, 0x03, 0x56, 0x85, 0x62, 0xbf, 0xf3, 0xf5, 0xda, 0x25, 0x8d, 0x1b,
	0xb2, 0x52, 0x60, 0x25, 0x62, 0xe5, 0xcc, 0x3a, 0x17, 0x03, 0x70, 0xdb, 0xdb, 0x92, 0xc3, 0x0e,
	0x83, 0x57, 0x0d, 0xac, 0xba, 0x4f, 0xfa, 0x60, 0x26, 0x48, 0x0b, 0xbf, 0x3b, 0xd4, 0x47, 0x06,
	0xf9, 0x0a, 0xf6, 0x8a, 0xb3, 0x34, 0xc7, 0x42, 0xbe, 0xa5, 0x84, 0x6f, 0xe0, 0xc6, 0xfb, 0xb4,
	0xc8, 0x44, 0xc8, 0xf6, 0x12, 0x40, 0x3f, 0xa5, 0xb0, 0xa4, 0xb9, 0x38, 0xf8, 0xc6, 0x31, 0xd5,
	0xbe, 0x97, 0x64, 0x05, 0xfa, 0xb6, 0x04, 0x31, 0xc8, 0x0f, 0xb0, 0x62, 0x41, 0x05, 0xfa, 0xa6,
	0xcc, 0x7c, 0x99, 0xf4, 0x1b, 0xae, 0xb8, 0xd6, 0xc8, 0x2f, 0xe8, 0xdd, 0x95, 0x1b, 0xce, 0x5e,
	0xb0, 0xf4, 0xad, 0x4f, 0xab, 0x82, 0x2d, 0xf4, 0xdf, 0x0f, 0x50, 0x94, 0x0b, 0x2a, 0x6f, 0x56,
	0x50, 0x0e, 0xf9, 0x0e, 0x96, 0x62, 0xae, 0x24, 0x8e, 0x21, 0x9b, 0xd0, 0x34, 0x95, 0x44, 0x7e,
	0x43, 0xaf, 0x1d, 0x50, 0x49, 0x36, 0xe5, 0x0e, 0x1a, 0xb7, 0x55, 0xc9, 0xa0, 0xa6, 0xe2, 0xa2,
	0xa6, 0x32, 0x82, 0x7f, 0xa7, 0xc1, 0x04, 0x40, 0x9f, 0xaf, 0xdb, 0xd9, 0x32, 0x75, 0x43, 0x53,
	0xcc, 0xeb, 0xc1, 0x8e, 0x3a, 0x86, 0x2c, 0x67, 0x5c, 0x5e, 0xa9, 0x8d, 0x9c, 0xe0, 0x7f, 0xf3,
	0x67, 0x1f, 0x1a, 0x47, 0xc2, 0xa6, 0x20, 0x3f, 0xe4, 0x7c, 0xb3, 0xc9, 0xf2, 0x8c, 0x8a, 0x8c,
	0x95, 0x4d, 0xed, 0xef, 0x19, 0x74, 0xdb, 0xf5, 0x2e, 0xd8, 0xe1, 0x32, 0x8a, 0xa6, 0x61, 0xe2,
	0x69, 0xc4, 0x01, 0xeb, 0x3e, 0x5a, 0x2e, 0xae, 0x3d, 0x5d, 0xe9, 0xab, 0xe9, 0xe2, 0x72, 0x2e,
	0x0f, 0x86, 0xd2, 0xaf, 0xe6, 0x51, 0x9c, 0x78, 0xe6, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x81, 0xc7, 0x85, 0x24, 0x02, 0x00, 0x00,
}
