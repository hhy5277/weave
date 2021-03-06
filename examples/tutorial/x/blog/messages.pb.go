// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/blog/messages.proto

/*
	Package blog is a generated protocol buffer package.

	It is generated from these files:
		x/blog/messages.proto
		x/blog/state.proto

	It has these top-level messages:
		CreateBlogMsg
		RenameBlogMsg
		ChangeBlogAuthorsMsg
		CreatePostMsg
		SetProfileMsg
		Blog
		Post
		Profile
*/
package blog

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CreateBlogMsg starts a new blog with a set of authors
type CreateBlogMsg struct {
	// slug is a short, unique string used as primary key
	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	// title is longer text used for display
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// initial set of authors (must be 1 - MaxAuthors)
	Authors [][]byte `protobuf:"bytes,3,rep,name=authors" json:"authors,omitempty"`
}

func (m *CreateBlogMsg) Reset()                    { *m = CreateBlogMsg{} }
func (m *CreateBlogMsg) String() string            { return proto.CompactTextString(m) }
func (*CreateBlogMsg) ProtoMessage()               {}
func (*CreateBlogMsg) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{0} }

func (m *CreateBlogMsg) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *CreateBlogMsg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateBlogMsg) GetAuthors() [][]byte {
	if m != nil {
		return m.Authors
	}
	return nil
}

// RenameBlogMsg updates the title of an existing blog
type RenameBlogMsg struct {
	// slug is a short, unique string used as primary key
	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	// title is longer text used for display
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (m *RenameBlogMsg) Reset()                    { *m = RenameBlogMsg{} }
func (m *RenameBlogMsg) String() string            { return proto.CompactTextString(m) }
func (*RenameBlogMsg) ProtoMessage()               {}
func (*RenameBlogMsg) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{1} }

func (m *RenameBlogMsg) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *RenameBlogMsg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

// ChangeBlogAuthorsMsg adds or removes an author from the blog's
// authorized author list
type ChangeBlogAuthorsMsg struct {
	// slug is a short, unique string used as primary key
	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	// whether we add or remove them
	Add bool `protobuf:"varint,2,opt,name=add,proto3" json:"add,omitempty"`
	// author to add or remove
	Author []byte `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (m *ChangeBlogAuthorsMsg) Reset()                    { *m = ChangeBlogAuthorsMsg{} }
func (m *ChangeBlogAuthorsMsg) String() string            { return proto.CompactTextString(m) }
func (*ChangeBlogAuthorsMsg) ProtoMessage()               {}
func (*ChangeBlogAuthorsMsg) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{2} }

func (m *ChangeBlogAuthorsMsg) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *ChangeBlogAuthorsMsg) GetAdd() bool {
	if m != nil {
		return m.Add
	}
	return false
}

func (m *ChangeBlogAuthorsMsg) GetAuthor() []byte {
	if m != nil {
		return m.Author
	}
	return nil
}

// CreatePostMsg adds a post to a blog
type CreatePostMsg struct {
	// blog is the slug of the blog this post belongs to
	Blog  string `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// author is optional, by default the first signer,
	// only needed if it is multisig
	Author []byte `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
}

func (m *CreatePostMsg) Reset()                    { *m = CreatePostMsg{} }
func (m *CreatePostMsg) String() string            { return proto.CompactTextString(m) }
func (*CreatePostMsg) ProtoMessage()               {}
func (*CreatePostMsg) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{3} }

func (m *CreatePostMsg) GetBlog() string {
	if m != nil {
		return m.Blog
	}
	return ""
}

func (m *CreatePostMsg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePostMsg) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreatePostMsg) GetAuthor() []byte {
	if m != nil {
		return m.Author
	}
	return nil
}

// SetProfileMsg will create or update a profile
type SetProfileMsg struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// author is optional, by default the first signer,
	// only needed if it is multisig
	Author []byte `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (m *SetProfileMsg) Reset()                    { *m = SetProfileMsg{} }
func (m *SetProfileMsg) String() string            { return proto.CompactTextString(m) }
func (*SetProfileMsg) ProtoMessage()               {}
func (*SetProfileMsg) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{4} }

func (m *SetProfileMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetProfileMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetProfileMsg) GetAuthor() []byte {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateBlogMsg)(nil), "blog.CreateBlogMsg")
	proto.RegisterType((*RenameBlogMsg)(nil), "blog.RenameBlogMsg")
	proto.RegisterType((*ChangeBlogAuthorsMsg)(nil), "blog.ChangeBlogAuthorsMsg")
	proto.RegisterType((*CreatePostMsg)(nil), "blog.CreatePostMsg")
	proto.RegisterType((*SetProfileMsg)(nil), "blog.SetProfileMsg")
}
func (m *CreateBlogMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateBlogMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Slug)))
		i += copy(dAtA[i:], m.Slug)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Authors) > 0 {
		for _, b := range m.Authors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintMessages(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *RenameBlogMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RenameBlogMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Slug)))
		i += copy(dAtA[i:], m.Slug)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	return i, nil
}

func (m *ChangeBlogAuthorsMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeBlogAuthorsMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Slug)))
		i += copy(dAtA[i:], m.Slug)
	}
	if m.Add {
		dAtA[i] = 0x10
		i++
		if m.Add {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Author) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	return i, nil
}

func (m *CreatePostMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePostMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Blog) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Blog)))
		i += copy(dAtA[i:], m.Blog)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Text) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if len(m.Author) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	return i, nil
}

func (m *SetProfileMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetProfileMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Author) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	return i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateBlogMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if len(m.Authors) > 0 {
		for _, b := range m.Authors {
			l = len(b)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	return n
}

func (m *RenameBlogMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *ChangeBlogAuthorsMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Add {
		n += 2
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *CreatePostMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Blog)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *SetProfileMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateBlogMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateBlogMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateBlogMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authors", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authors = append(m.Authors, make([]byte, postIndex-iNdEx))
			copy(m.Authors[len(m.Authors)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *RenameBlogMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RenameBlogMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RenameBlogMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *ChangeBlogAuthorsMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChangeBlogAuthorsMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeBlogAuthorsMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Add", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Add = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = append(m.Author[:0], dAtA[iNdEx:postIndex]...)
			if m.Author == nil {
				m.Author = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *CreatePostMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePostMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePostMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blog = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = append(m.Author[:0], dAtA[iNdEx:postIndex]...)
			if m.Author == nil {
				m.Author = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *SetProfileMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetProfileMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetProfileMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = append(m.Author[:0], dAtA[iNdEx:postIndex]...)
			if m.Author == nil {
				m.Author = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessages
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessages
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessages(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessages = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/blog/messages.proto", fileDescriptorMessages) }

var fileDescriptorMessages = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x31, 0x09, 0x85, 0xfe, 0x34, 0x52, 0x65, 0x15, 0x94, 0x29, 0x8a, 0x32, 0x65, 0xa2,
	0x03, 0x13, 0x23, 0xed, 0x8c, 0x54, 0xb9, 0xac, 0x0c, 0x2e, 0xf9, 0x71, 0x23, 0xa5, 0x71, 0x65,
	0xff, 0x95, 0xfa, 0x18, 0x3c, 0x16, 0x23, 0x8f, 0x80, 0xc2, 0x8b, 0x20, 0x3b, 0x89, 0x54, 0x86,
	0x0c, 0xdd, 0xee, 0x4e, 0xfa, 0xbf, 0xcb, 0xc5, 0x70, 0x77, 0x9c, 0x6f, 0x2a, 0xad, 0xe6, 0x3b,
	0xb4, 0x56, 0x2a, 0xb4, 0x0f, 0x7b, 0xa3, 0x49, 0xf3, 0xd0, 0x85, 0xd9, 0x1a, 0xa2, 0xa5, 0x41,
	0x49, 0xb8, 0xa8, 0xb4, 0x7a, 0xb1, 0x8a, 0x73, 0x08, 0x6d, 0x75, 0x50, 0x31, 0x4b, 0x59, 0x3e,
	0x16, 0x5e, 0xf3, 0x19, 0x5c, 0x51, 0x49, 0x15, 0xc6, 0x97, 0x3e, 0x6c, 0x0d, 0x8f, 0xe1, 0x5a,
	0x1e, 0x68, 0xab, 0x8d, 0x8d, 0x83, 0x34, 0xc8, 0x27, 0xa2, 0xb7, 0xd9, 0x13, 0x44, 0x02, 0x6b,
	0xb9, 0x3b, 0x1f, 0x9a, 0xbd, 0xc2, 0x6c, 0xb9, 0x95, 0xb5, 0xf2, 0xa7, 0xcf, 0x2d, 0x6f, 0x88,
	0x30, 0x85, 0x40, 0x16, 0x85, 0xbf, 0xbf, 0x11, 0x4e, 0xf2, 0x7b, 0x18, 0xb5, 0xdf, 0x10, 0x07,
	0x29, 0xcb, 0x27, 0xa2, 0x73, 0x19, 0xf6, 0x2b, 0x57, 0xda, 0x52, 0x87, 0x73, 0xf3, 0x7b, 0x9c,
	0xd3, 0x03, 0x2b, 0x39, 0x84, 0x84, 0x47, 0xf2, 0xc0, 0xb1, 0xf0, 0xfa, 0xa4, 0x26, 0xfc, 0x57,
	0xf3, 0x06, 0xd1, 0x1a, 0x69, 0x65, 0xf4, 0x47, 0x59, 0x61, 0x57, 0xe3, 0x7e, 0x43, 0x5f, 0xe3,
	0x34, 0x4f, 0xe1, 0xb6, 0x40, 0xfb, 0x6e, 0xca, 0x3d, 0x95, 0xba, 0xee, 0xca, 0x4e, 0xa3, 0xa1,
	0x15, 0x8b, 0xe9, 0x57, 0x93, 0xb0, 0xef, 0x26, 0x61, 0x3f, 0x4d, 0xc2, 0x3e, 0x7f, 0x93, 0x8b,
	0xcd, 0xc8, 0x3f, 0xe5, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xb3, 0x13, 0x2e, 0xe3,
	0x01, 0x00, 0x00,
}
