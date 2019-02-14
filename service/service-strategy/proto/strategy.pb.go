// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/strategy.proto

package bitesla_srv_strategy

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

type StrategyInfo struct {
	CurrentLoginUserID   int64    `protobuf:"varint,1,opt,name=currentLoginUserID,proto3" json:"currentLoginUserID,omitempty"`
	StrategyId           int64    `protobuf:"varint,2,opt,name=strategyId,proto3" json:"strategyId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Script               string   `protobuf:"bytes,5,opt,name=script,proto3" json:"script,omitempty"`
	Page                 int32    `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	CreateTime           int64    `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StrategyInfo) Reset()         { *m = StrategyInfo{} }
func (m *StrategyInfo) String() string { return proto.CompactTextString(m) }
func (*StrategyInfo) ProtoMessage()    {}
func (*StrategyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{0}
}

func (m *StrategyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyInfo.Unmarshal(m, b)
}
func (m *StrategyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyInfo.Marshal(b, m, deterministic)
}
func (m *StrategyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyInfo.Merge(m, src)
}
func (m *StrategyInfo) XXX_Size() int {
	return xxx_messageInfo_StrategyInfo.Size(m)
}
func (m *StrategyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyInfo proto.InternalMessageInfo

func (m *StrategyInfo) GetCurrentLoginUserID() int64 {
	if m != nil {
		return m.CurrentLoginUserID
	}
	return 0
}

func (m *StrategyInfo) GetStrategyId() int64 {
	if m != nil {
		return m.StrategyId
	}
	return 0
}

func (m *StrategyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StrategyInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StrategyInfo) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *StrategyInfo) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *StrategyInfo) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *StrategyInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *StrategyInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type StrategyInfos struct {
	StrategyInfos        []*StrategyInfo `protobuf:"bytes,1,rep,name=strategyInfos,proto3" json:"strategyInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StrategyInfos) Reset()         { *m = StrategyInfos{} }
func (m *StrategyInfos) String() string { return proto.CompactTextString(m) }
func (*StrategyInfos) ProtoMessage()    {}
func (*StrategyInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{1}
}

func (m *StrategyInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyInfos.Unmarshal(m, b)
}
func (m *StrategyInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyInfos.Marshal(b, m, deterministic)
}
func (m *StrategyInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyInfos.Merge(m, src)
}
func (m *StrategyInfos) XXX_Size() int {
	return xxx_messageInfo_StrategyInfos.Size(m)
}
func (m *StrategyInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyInfos.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyInfos proto.InternalMessageInfo

func (m *StrategyInfos) GetStrategyInfos() []*StrategyInfo {
	if m != nil {
		return m.StrategyInfos
	}
	return nil
}

type User struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UnionId              string   `protobuf:"bytes,2,opt,name=unionId,proto3" json:"unionId,omitempty"`
	LoginTime            string   `protobuf:"bytes,3,opt,name=loginTime,proto3" json:"loginTime,omitempty"`
	PhoneCode            string   `protobuf:"bytes,4,opt,name=phoneCode,proto3" json:"phoneCode,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Sex                  int32    `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Nickname             string   `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Birthday             string   `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	VerifyCode           string   `protobuf:"bytes,9,opt,name=verifyCode,proto3" json:"verifyCode,omitempty"`
	Phone                string   `protobuf:"bytes,10,opt,name=phone,proto3" json:"phone,omitempty"`
	RecommendCode        string   `protobuf:"bytes,11,opt,name=recommendCode,proto3" json:"recommendCode,omitempty"`
	Avatar               string   `protobuf:"bytes,12,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Hometown             string   `protobuf:"bytes,13,opt,name=hometown,proto3" json:"hometown,omitempty"`
	RealName             string   `protobuf:"bytes,14,opt,name=realName,proto3" json:"realName,omitempty"`
	Remark               string   `protobuf:"bytes,15,opt,name=remark,proto3" json:"remark,omitempty"`
	Email                string   `protobuf:"bytes,16,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,17,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,18,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func (m *User) GetLoginTime() string {
	if m != nil {
		return m.LoginTime
	}
	return ""
}

func (m *User) GetPhoneCode() string {
	if m != nil {
		return m.PhoneCode
	}
	return ""
}

func (m *User) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *User) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetRecommendCode() string {
	if m != nil {
		return m.RecommendCode
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetHometown() string {
	if m != nil {
		return m.Hometown
	}
	return ""
}

func (m *User) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *User) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*StrategyInfo)(nil), "bitesla.srv.strategy.StrategyInfo")
	proto.RegisterType((*StrategyInfos)(nil), "bitesla.srv.strategy.StrategyInfos")
	proto.RegisterType((*User)(nil), "bitesla.srv.strategy.User")
}

func init() { proto.RegisterFile("proto/strategy.proto", fileDescriptor_8a9cfd4cb1c623bb) }

var fileDescriptor_8a9cfd4cb1c623bb = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x6d, 0x9a, 0xcf, 0xbd, 0x49, 0x6a, 0x3b, 0x04, 0x19, 0x8a, 0x48, 0x58, 0x7d, 0xc8, 0xd3,
	0x0a, 0xf5, 0x27, 0x18, 0xd0, 0x40, 0x11, 0x59, 0x15, 0xa9, 0xe0, 0xc3, 0x24, 0x7b, 0x9b, 0x0c,
	0xc9, 0xce, 0x2c, 0x33, 0x93, 0x68, 0xfc, 0x1d, 0xfe, 0x3e, 0xff, 0x87, 0x6f, 0x72, 0x67, 0x76,
	0x93, 0x2d, 0x14, 0xac, 0xd0, 0xb7, 0x7b, 0xce, 0x99, 0x7b, 0x72, 0xf6, 0xcc, 0x10, 0x18, 0x15,
	0x46, 0x3b, 0xfd, 0xca, 0x3a, 0x23, 0x1c, 0x2e, 0xf7, 0x89, 0x87, 0x6c, 0x34, 0x97, 0x0e, 0xed,
	0x46, 0x24, 0xd6, 0xec, 0x92, 0x4a, 0x8b, 0x7f, 0x9d, 0xc2, 0xe0, 0x63, 0x09, 0x66, 0xea, 0x56,
	0xb3, 0x04, 0xd8, 0x62, 0x6b, 0x0c, 0x2a, 0x77, 0xad, 0x97, 0x52, 0x7d, 0xb6, 0x68, 0x66, 0x53,
	0xde, 0x18, 0x37, 0x26, 0xcd, 0xf4, 0x1e, 0x85, 0x3d, 0x07, 0xa8, 0xcc, 0x66, 0x19, 0x3f, 0xf5,
	0xe7, 0x6a, 0x0c, 0x63, 0xd0, 0x52, 0x22, 0x47, 0xde, 0x1c, 0x37, 0x26, 0x51, 0xea, 0x67, 0x36,
	0x86, 0x7e, 0x86, 0x76, 0x61, 0x64, 0xe1, 0xa4, 0x56, 0xbc, 0xe5, 0xa5, 0x3a, 0xc5, 0x9e, 0x42,
	0x27, 0x00, 0xde, 0xf6, 0x62, 0x89, 0xc8, 0xad, 0x10, 0x4b, 0xe4, 0x9d, 0x71, 0x63, 0xd2, 0x4e,
	0xfd, 0x4c, 0x9c, 0x95, 0x3f, 0x91, 0x77, 0x03, 0x47, 0x33, 0xa5, 0x5a, 0x18, 0x14, 0x0e, 0x3f,
	0xc9, 0x1c, 0x79, 0x2f, 0xa4, 0x3a, 0x32, 0xa4, 0x6f, 0x8b, 0xac, 0xd2, 0xa3, 0xa0, 0x1f, 0x99,
	0xf8, 0x06, 0x86, 0xf5, 0x56, 0x2c, 0x7b, 0x07, 0x43, 0x5b, 0x27, 0x78, 0x63, 0xdc, 0x9c, 0xf4,
	0xaf, 0xe2, 0xe4, 0xbe, 0x56, 0x93, 0xfa, 0x6e, 0x7a, 0x77, 0x31, 0xfe, 0xdd, 0x84, 0x16, 0x75,
	0x47, 0xdf, 0xb8, 0xa5, 0x0e, 0xb3, 0xb2, 0xdd, 0x12, 0x31, 0x0e, 0xdd, 0xad, 0x92, 0x5a, 0x95,
	0x75, 0x46, 0x69, 0x05, 0xd9, 0x33, 0x88, 0x36, 0x54, 0xbd, 0x0f, 0x1d, 0x0a, 0x3d, 0x12, 0xa4,
	0x16, 0x2b, 0xad, 0xf0, 0x8d, 0xce, 0xb0, 0xec, 0xf4, 0x48, 0x50, 0x4b, 0x6e, 0x5f, 0x60, 0xd9,
	0xa7, 0x9f, 0xd9, 0x39, 0x34, 0x2d, 0xfe, 0x28, 0xcb, 0xa4, 0x91, 0x5d, 0x42, 0x4f, 0xc9, 0xc5,
	0xda, 0xdf, 0x58, 0xd7, 0x9f, 0x3c, 0x60, 0xd2, 0xe6, 0xd2, 0xb8, 0x55, 0x26, 0xf6, 0xbe, 0xd1,
	0x28, 0x3d, 0x60, 0xea, 0x73, 0x87, 0x46, 0xde, 0xee, 0xfd, 0x8f, 0x47, 0x5e, 0xad, 0x31, 0x6c,
	0x04, 0x6d, 0x1f, 0x85, 0x83, 0x97, 0x02, 0x60, 0x2f, 0x61, 0x68, 0x70, 0xa1, 0xf3, 0x1c, 0x55,
	0xe6, 0x17, 0xfb, 0x5e, 0xbd, 0x4b, 0x52, 0x4f, 0x62, 0x27, 0x9c, 0x30, 0x7c, 0x10, 0xde, 0x42,
	0x40, 0x94, 0x67, 0xa5, 0x73, 0x74, 0xfa, 0xbb, 0xe2, 0xc3, 0x90, 0xa7, 0xc2, 0xa4, 0x19, 0x14,
	0x9b, 0xf7, 0xf4, 0x1d, 0x67, 0x41, 0xab, 0x30, 0xf9, 0x19, 0xcc, 0x85, 0x59, 0xf3, 0x27, 0xc1,
	0x2f, 0x20, 0xca, 0x88, 0xb9, 0x90, 0x1b, 0x7e, 0x1e, 0x32, 0x7a, 0x40, 0x4e, 0x74, 0x2f, 0xbe,
	0x91, 0x8b, 0xe0, 0x54, 0x61, 0xda, 0x70, 0x7a, 0x8d, 0x8a, 0xb3, 0xb0, 0xe1, 0xc1, 0xd5, 0x9f,
	0x53, 0xe8, 0x55, 0x0f, 0x80, 0xdd, 0xc0, 0xe0, 0x5a, 0x5a, 0x77, 0xc0, 0x0f, 0x78, 0x30, 0x97,
	0x2f, 0xfe, 0x7d, 0xc6, 0xc6, 0x27, 0xec, 0x0b, 0xf4, 0x3f, 0x6c, 0xff, 0xcf, 0xf9, 0x01, 0x67,
	0xe2, 0x13, 0xf6, 0x0d, 0x2e, 0xde, 0xe2, 0xc1, 0x78, 0x8a, 0x8e, 0x7a, 0x78, 0x3c, 0xfb, 0xaf,
	0x70, 0x36, 0xc5, 0x0d, 0x3a, 0x7c, 0xfc, 0xe8, 0xf3, 0x8e, 0xff, 0xaf, 0x7b, 0xfd, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x84, 0x44, 0x31, 0x8a, 0x03, 0x05, 0x00, 0x00,
}