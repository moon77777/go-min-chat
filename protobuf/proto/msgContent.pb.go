// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgContent.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Content struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ParamId              int32    `protobuf:"varint,3,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	ParamString          string   `protobuf:"bytes,4,opt,name=param_string,json=paramString,proto3" json:"param_string,omitempty"`
	Nick                 string   `protobuf:"bytes,5,opt,name=nick,proto3" json:"nick,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{0}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Content) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Content) GetParamId() int32 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

func (m *Content) GetParamString() string {
	if m != nil {
		return m.ParamString
	}
	return ""
}

func (m *Content) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *Content) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type BackContent struct {
	Id                   int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 int32     `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string    `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Auth                 *Auth     `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	Showroom             *ShowRoom `protobuf:"bytes,5,opt,name=showroom,proto3" json:"showroom,omitempty"`
	Groupmsg             *GroupMsg `protobuf:"bytes,6,opt,name=groupmsg,proto3" json:"groupmsg,omitempty"`
	Room                 *Room     `protobuf:"bytes,7,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BackContent) Reset()         { *m = BackContent{} }
func (m *BackContent) String() string { return proto.CompactTextString(m) }
func (*BackContent) ProtoMessage()    {}
func (*BackContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{1}
}

func (m *BackContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackContent.Unmarshal(m, b)
}
func (m *BackContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackContent.Marshal(b, m, deterministic)
}
func (m *BackContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackContent.Merge(m, src)
}
func (m *BackContent) XXX_Size() int {
	return xxx_messageInfo_BackContent.Size(m)
}
func (m *BackContent) XXX_DiscardUnknown() {
	xxx_messageInfo_BackContent.DiscardUnknown(m)
}

var xxx_messageInfo_BackContent proto.InternalMessageInfo

func (m *BackContent) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BackContent) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BackContent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BackContent) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *BackContent) GetShowroom() *ShowRoom {
	if m != nil {
		return m.Showroom
	}
	return nil
}

func (m *BackContent) GetGroupmsg() *GroupMsg {
	if m != nil {
		return m.Groupmsg
	}
	return nil
}

func (m *BackContent) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type Auth struct {
	IsOk                 bool      `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UseInfo              *Userinfo `protobuf:"bytes,3,opt,name=useInfo,proto3" json:"useInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{2}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *Auth) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Auth) GetUseInfo() *Userinfo {
	if m != nil {
		return m.UseInfo
	}
	return nil
}

type GroupMsg struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Nick                 string   `protobuf:"bytes,3,opt,name=nick,proto3" json:"nick,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupMsg) Reset()         { *m = GroupMsg{} }
func (m *GroupMsg) String() string { return proto.CompactTextString(m) }
func (*GroupMsg) ProtoMessage()    {}
func (*GroupMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{3}
}

func (m *GroupMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupMsg.Unmarshal(m, b)
}
func (m *GroupMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupMsg.Marshal(b, m, deterministic)
}
func (m *GroupMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMsg.Merge(m, src)
}
func (m *GroupMsg) XXX_Size() int {
	return xxx_messageInfo_GroupMsg.Size(m)
}
func (m *GroupMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMsg proto.InternalMessageInfo

func (m *GroupMsg) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *GroupMsg) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GroupMsg) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *GroupMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ShowRoom struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	RoomsAndIds          string   `protobuf:"bytes,2,opt,name=roomsAndIds,proto3" json:"roomsAndIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowRoom) Reset()         { *m = ShowRoom{} }
func (m *ShowRoom) String() string { return proto.CompactTextString(m) }
func (*ShowRoom) ProtoMessage()    {}
func (*ShowRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{4}
}

func (m *ShowRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRoom.Unmarshal(m, b)
}
func (m *ShowRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRoom.Marshal(b, m, deterministic)
}
func (m *ShowRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRoom.Merge(m, src)
}
func (m *ShowRoom) XXX_Size() int {
	return xxx_messageInfo_ShowRoom.Size(m)
}
func (m *ShowRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRoom proto.InternalMessageInfo

func (m *ShowRoom) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ShowRoom) GetRoomsAndIds() string {
	if m != nil {
		return m.RoomsAndIds
	}
	return ""
}

type Room struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{5}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Room) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type Userinfo struct {
	Nick                 string   `protobuf:"bytes,1,opt,name=nick,proto3" json:"nick,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId               int32    `protobuf:"varint,3,opt,name=roomId,proto3" json:"roomId,omitempty"`
	RoomName             string   `protobuf:"bytes,4,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Userinfo) Reset()         { *m = Userinfo{} }
func (m *Userinfo) String() string { return proto.CompactTextString(m) }
func (*Userinfo) ProtoMessage()    {}
func (*Userinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7ebc832082b9679, []int{6}
}

func (m *Userinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Userinfo.Unmarshal(m, b)
}
func (m *Userinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Userinfo.Marshal(b, m, deterministic)
}
func (m *Userinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Userinfo.Merge(m, src)
}
func (m *Userinfo) XXX_Size() int {
	return xxx_messageInfo_Userinfo.Size(m)
}
func (m *Userinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Userinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Userinfo proto.InternalMessageInfo

func (m *Userinfo) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *Userinfo) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Userinfo) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Userinfo) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func init() {
	proto.RegisterType((*Content)(nil), "protobuf.Content")
	proto.RegisterType((*BackContent)(nil), "protobuf.BackContent")
	proto.RegisterType((*Auth)(nil), "protobuf.Auth")
	proto.RegisterType((*GroupMsg)(nil), "protobuf.GroupMsg")
	proto.RegisterType((*ShowRoom)(nil), "protobuf.ShowRoom")
	proto.RegisterType((*Room)(nil), "protobuf.Room")
	proto.RegisterType((*Userinfo)(nil), "protobuf.Userinfo")
}

func init() { proto.RegisterFile("msgContent.proto", fileDescriptor_d7ebc832082b9679) }

var fileDescriptor_d7ebc832082b9679 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0x6c, 0xd9, 0x62, 0x4e, 0x45, 0x10, 0x1c, 0x8a, 0x42, 0xcd, 0xe4, 0x72, 0xca, 0x50,
	0x78, 0x48, 0xb7, 0x6e, 0x49, 0x87, 0xc2, 0x43, 0x5b, 0x80, 0x41, 0x97, 0x2e, 0x85, 0x2c, 0xca,
	0x12, 0x61, 0x90, 0x67, 0x88, 0x12, 0xfc, 0x3d, 0xfd, 0xbe, 0xfe, 0x44, 0xc1, 0x8b, 0x28, 0x3b,
	0x29, 0x3c, 0xe9, 0x1e, 0xef, 0x91, 0xef, 0xbd, 0x3b, 0xc1, 0x8d, 0xf5, 0xcd, 0x17, 0x72, 0x7d,
	0xed, 0xfa, 0xf5, 0xa1, 0xa3, 0x9e, 0x50, 0xf0, 0x67, 0x3b, 0xec, 0xe4, 0x9f, 0x04, 0xb2, 0xb1,
	0x87, 0xd7, 0x30, 0x33, 0xba, 0x48, 0x56, 0xc9, 0xdd, 0x42, 0xcd, 0x8c, 0xc6, 0x02, 0xb2, 0x8a,
	0xac, 0x2d, 0x9d, 0x2e, 0x66, 0xab, 0xe4, 0xee, 0x4a, 0x45, 0x88, 0xef, 0x41, 0x1c, 0xca, 0xae,
	0xb4, 0xbf, 0x8d, 0x2e, 0xe6, 0xcc, 0xcf, 0x18, 0x6f, 0x34, 0x7e, 0x80, 0x37, 0xcf, 0x2d, 0xdf,
	0x77, 0xc6, 0x35, 0x45, 0xca, 0x37, 0x73, 0x3e, 0x7b, 0xe2, 0x23, 0x44, 0x48, 0x9d, 0xa9, 0xf6,
	0xc5, 0x82, 0x5b, 0x5c, 0xe3, 0x6d, 0x78, 0xd1, 0xfb, 0x23, 0x75, 0xba, 0x58, 0xf2, 0xf9, 0x84,
	0xe5, 0xdf, 0x04, 0xf2, 0xc7, 0xb2, 0xda, 0x5f, 0xf2, 0x89, 0x90, 0x56, 0xa4, 0x6b, 0x36, 0xb9,
	0x50, 0x5c, 0xe3, 0x0d, 0xcc, 0xad, 0x6f, 0xd8, 0xdc, 0x95, 0x0a, 0x25, 0x4a, 0x48, 0xcb, 0xa1,
	0x6f, 0xd9, 0x50, 0x7e, 0x7f, 0xbd, 0x8e, 0x23, 0x58, 0x3f, 0x0c, 0x7d, 0xab, 0xb8, 0x87, 0x6b,
	0x10, 0xbe, 0xa5, 0x63, 0x47, 0x64, 0xd9, 0x5d, 0x7e, 0x8f, 0x27, 0xde, 0x53, 0x4b, 0x47, 0x45,
	0x64, 0xd5, 0xc4, 0x09, 0xfc, 0xa6, 0xa3, 0xe1, 0x10, 0xa4, 0x96, 0xaf, 0xf9, 0x5f, 0x43, 0xe7,
	0x9b, 0x6f, 0xd4, 0xc4, 0x09, 0x1e, 0xf8, 0xed, 0xec, 0xb5, 0x07, 0x7e, 0x97, 0x7b, 0xf2, 0x17,
	0xa4, 0xc1, 0x51, 0x48, 0x65, 0xfc, 0x8f, 0x3d, 0xe7, 0x14, 0x8a, 0xeb, 0x98, 0x6a, 0x76, 0x4a,
	0xf5, 0x11, 0xb2, 0xc1, 0xd7, 0x1b, 0xb7, 0x23, 0xce, 0xfa, 0xc2, 0xc0, 0x4f, 0x5f, 0x77, 0xc6,
	0xed, 0x48, 0x45, 0x8a, 0xdc, 0x82, 0x88, 0xae, 0xf0, 0x1d, 0x2c, 0x83, 0xde, 0x26, 0x4e, 0x72,
	0x44, 0x41, 0x63, 0x30, 0x7a, 0x1c, 0x66, 0x28, 0xa7, 0x7d, 0xcd, 0xcf, 0xf6, 0xc5, 0xff, 0x06,
	0xaf, 0x63, 0xdc, 0x70, 0x84, 0xf2, 0x11, 0x44, 0x9c, 0x14, 0xbe, 0x85, 0x45, 0x45, 0x83, 0xeb,
	0x47, 0x89, 0x67, 0x80, 0x2b, 0xc8, 0x83, 0x96, 0x7f, 0x70, 0x7a, 0xa3, 0xfd, 0x98, 0xe6, 0xfc,
	0x48, 0x7e, 0x86, 0x94, 0xef, 0x5f, 0xf2, 0x78, 0x0b, 0x22, 0x54, 0xdf, 0x4b, 0x5b, 0x8f, 0xd7,
	0x27, 0x2c, 0x35, 0x88, 0x18, 0x7c, 0x72, 0x9e, 0x9c, 0x39, 0xff, 0x3f, 0xdf, 0x49, 0x65, 0x7e,
	0x51, 0x25, 0x7d, 0xa9, 0xb2, 0x5d, 0xf2, 0x94, 0x3f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0xd4, 0xb7, 0xb7, 0x5c, 0x03, 0x00, 0x00,
}
