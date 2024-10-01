// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: svapi/comment.proto

package svapi

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty,string"` // 评论id
	// @gotags: json:"videoId,omitempty,string"
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"videoId,omitempty,string"` // 视频id
	// @gotags: json:"parentId,omitempty,string"
	ParentId   int64        `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parentId,omitempty,string"`      // 父评论id
	User       *CommentUser `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`                               // 评论用户
	ReplyUser  *CommentUser `protobuf:"bytes,5,opt,name=reply_user,json=replyUser,proto3" json:"reply_user,omitempty"`    // 回复用户
	Content    string       `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`                         // 评论内容
	Date       string       `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`                               // 评论日期
	LikeCount  string       `protobuf:"bytes,8,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`    // 点赞数
	ReplyCount string       `protobuf:"bytes,9,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"` // 回复数
	Comments   []*Comment   `protobuf:"bytes,10,rep,name=comments,proto3" json:"comments,omitempty"`                      // 子评论
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *Comment) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Comment) GetUser() *CommentUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetReplyUser() *CommentUser {
	if x != nil {
		return x.ReplyUser
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Comment) GetLikeCount() string {
	if x != nil {
		return x.LikeCount
	}
	return ""
}

func (x *Comment) GetReplyCount() string {
	if x != nil {
		return x.ReplyCount
	}
	return ""
}

func (x *Comment) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty,string"`                                      // 用户id
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // 用户名称
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                               // 用户头像
	IsFollowing bool   `protobuf:"varint,4,opt,name=is_following,json=isFollowing,proto3" json:"is_following,omitempty"` // 是否关注
}

func (x *CommentUser) Reset() {
	*x = CommentUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUser) ProtoMessage() {}

func (x *CommentUser) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUser.ProtoReflect.Descriptor instead.
func (*CommentUser) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommentUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CommentUser) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"videoId,omitempty,string"
	VideoId int64  `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"videoId,omitempty,string"` // 视频id
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                 // 评论内容
	// @gotags: json:"parentId,omitempty,string"
	ParentId int64 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parentId,omitempty,string"`
	// @gotags: json:"replyUserId,omitempty,string"
	ReplyUserId int64 `protobuf:"varint,4,opt,name=reply_user_id,json=replyUserId,proto3" json:"replyUserId,omitempty,string"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CreateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateCommentRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CreateCommentRequest) GetReplyUserId() int64 {
	if x != nil {
		return x.ReplyUserId
	}
	return 0
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type RemoveCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty,string"` // 评论id
}

func (x *RemoveCommentRequest) Reset() {
	*x = RemoveCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentRequest) ProtoMessage() {}

func (x *RemoveCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentRequest.ProtoReflect.Descriptor instead.
func (*RemoveCommentRequest) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveCommentResponse) Reset() {
	*x = RemoveCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentResponse) ProtoMessage() {}

func (x *RemoveCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentResponse.ProtoReflect.Descriptor instead.
func (*RemoveCommentResponse) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{5}
}

type ListComment4VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"videoId,omitempty,string"
	VideoId    int64              `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"videoId,omitempty,string"` // 视频id
	Pagination *PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListComment4VideoRequest) Reset() {
	*x = ListComment4VideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComment4VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComment4VideoRequest) ProtoMessage() {}

func (x *ListComment4VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComment4VideoRequest.ProtoReflect.Descriptor instead.
func (*ListComment4VideoRequest) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{6}
}

func (x *ListComment4VideoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *ListComment4VideoRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListComment4VideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*Comment          `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListComment4VideoResponse) Reset() {
	*x = ListComment4VideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComment4VideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComment4VideoResponse) ProtoMessage() {}

func (x *ListComment4VideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComment4VideoResponse.ProtoReflect.Descriptor instead.
func (*ListComment4VideoResponse) Descriptor() ([]byte, []int) {
	return file_svapi_comment_proto_rawDescGZIP(), []int{7}
}

func (x *ListComment4VideoResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListComment4VideoResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_svapi_comment_proto protoreflect.FileDescriptor

var file_svapi_comment_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x76, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x76,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x76, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x6c, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x22, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x76, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6f, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc2, 0x02, 0x0a, 0x0e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e,
	0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x76, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a,
	0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x76,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x2a, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6e,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x1f, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44, 0x6f, 0x75, 0x54, 0x6f, 0x6b, 0x2f,
	0x2e, 0x2e, 0x2e, 0x3b, 0x73, 0x76, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_svapi_comment_proto_rawDescOnce sync.Once
	file_svapi_comment_proto_rawDescData = file_svapi_comment_proto_rawDesc
)

func file_svapi_comment_proto_rawDescGZIP() []byte {
	file_svapi_comment_proto_rawDescOnce.Do(func() {
		file_svapi_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_svapi_comment_proto_rawDescData)
	})
	return file_svapi_comment_proto_rawDescData
}

var file_svapi_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_svapi_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),                   // 0: svapi.Comment
	(*CommentUser)(nil),               // 1: svapi.CommentUser
	(*CreateCommentRequest)(nil),      // 2: svapi.CreateCommentRequest
	(*CreateCommentResponse)(nil),     // 3: svapi.CreateCommentResponse
	(*RemoveCommentRequest)(nil),      // 4: svapi.RemoveCommentRequest
	(*RemoveCommentResponse)(nil),     // 5: svapi.RemoveCommentResponse
	(*ListComment4VideoRequest)(nil),  // 6: svapi.ListComment4VideoRequest
	(*ListComment4VideoResponse)(nil), // 7: svapi.ListComment4VideoResponse
	(*PaginationRequest)(nil),         // 8: svapi.PaginationRequest
	(*PaginationResponse)(nil),        // 9: svapi.PaginationResponse
}
var file_svapi_comment_proto_depIdxs = []int32{
	1,  // 0: svapi.Comment.user:type_name -> svapi.CommentUser
	1,  // 1: svapi.Comment.reply_user:type_name -> svapi.CommentUser
	0,  // 2: svapi.Comment.comments:type_name -> svapi.Comment
	0,  // 3: svapi.CreateCommentResponse.comment:type_name -> svapi.Comment
	8,  // 4: svapi.ListComment4VideoRequest.pagination:type_name -> svapi.PaginationRequest
	0,  // 5: svapi.ListComment4VideoResponse.comments:type_name -> svapi.Comment
	9,  // 6: svapi.ListComment4VideoResponse.pagination:type_name -> svapi.PaginationResponse
	2,  // 7: svapi.CommentService.CreateComment:input_type -> svapi.CreateCommentRequest
	4,  // 8: svapi.CommentService.RemoveComment:input_type -> svapi.RemoveCommentRequest
	6,  // 9: svapi.CommentService.ListComment4Video:input_type -> svapi.ListComment4VideoRequest
	3,  // 10: svapi.CommentService.CreateComment:output_type -> svapi.CreateCommentResponse
	5,  // 11: svapi.CommentService.RemoveComment:output_type -> svapi.RemoveCommentResponse
	7,  // 12: svapi.CommentService.ListComment4Video:output_type -> svapi.ListComment4VideoResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_svapi_comment_proto_init() }
func file_svapi_comment_proto_init() {
	if File_svapi_comment_proto != nil {
		return
	}
	file_svapi_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_svapi_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCommentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComment4VideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComment4VideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svapi_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svapi_comment_proto_goTypes,
		DependencyIndexes: file_svapi_comment_proto_depIdxs,
		MessageInfos:      file_svapi_comment_proto_msgTypes,
	}.Build()
	File_svapi_comment_proto = out.File
	file_svapi_comment_proto_rawDesc = nil
	file_svapi_comment_proto_goTypes = nil
	file_svapi_comment_proto_depIdxs = nil
}
