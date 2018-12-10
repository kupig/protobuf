// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/compiler/plugin.proto

package plugin_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	descriptor "github.com/golang/protobuf/v2/types/descriptor"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The version number of protocol compiler.
type Version struct {
	Major *int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor *int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch *int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               *string  `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Version struct{ m *Version }

func (m *Version) ProtoReflect() protoreflect.Message {
	return xxx_Version{m}
}
func (m xxx_Version) Type() protoreflect.MessageType {
	return xxx_Plugin_ProtoFile_MessageTypes[0].Type
}
func (m xxx_Version) KnownFields() protoreflect.KnownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_Version) UnknownFields() protoreflect.UnknownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_Version) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil && m.Minor != nil {
		return *m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil && m.Suffix != nil {
		return *m.Suffix
	}
	return ""
}

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate" json:"file_to_generate,omitempty"`
	// The generator parameter passed on the command-line.
	Parameter *string `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	// FileDescriptorProtos for all files in files_to_generate and everything
	// they import.  The files will appear in topological order, so each file
	// appears before any file that imports it.
	//
	// protoc guarantees that all proto_files will be written after
	// the fields above, even though this is not technically guaranteed by the
	// protobuf wire format.  This theoretically could allow a plugin to stream
	// in the FileDescriptorProtos and handle them one by one rather than read
	// the entire set into memory at once.  However, as of this writing, this
	// is not similarly optimized on protoc's end -- it will store all fields in
	// memory at once before sending them to the plugin.
	//
	// Type names of fields and extensions in the FileDescriptorProto are always
	// fully qualified.
	ProtoFile []*descriptor.FileDescriptorProto `protobuf:"bytes,15,rep,name=proto_file,json=protoFile" json:"proto_file,omitempty"`
	// The version number of protocol compiler.
	CompilerVersion      *Version `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_CodeGeneratorRequest struct{ m *CodeGeneratorRequest }

func (m *CodeGeneratorRequest) ProtoReflect() protoreflect.Message {
	return xxx_CodeGeneratorRequest{m}
}
func (m xxx_CodeGeneratorRequest) Type() protoreflect.MessageType {
	return xxx_Plugin_ProtoFile_MessageTypes[1].Type
}
func (m xxx_CodeGeneratorRequest) KnownFields() protoreflect.KnownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorRequest) UnknownFields() protoreflect.UnknownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorRequest) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *CodeGeneratorRequest) Reset()         { *m = CodeGeneratorRequest{} }
func (m *CodeGeneratorRequest) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorRequest) ProtoMessage()    {}
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{1}
}

func (m *CodeGeneratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorRequest.Unmarshal(m, b)
}
func (m *CodeGeneratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorRequest.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorRequest.Merge(m, src)
}
func (m *CodeGeneratorRequest) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorRequest.Size(m)
}
func (m *CodeGeneratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorRequest proto.InternalMessageInfo

func (m *CodeGeneratorRequest) GetFileToGenerate() []string {
	if m != nil {
		return m.FileToGenerate
	}
	return nil
}

func (m *CodeGeneratorRequest) GetParameter() string {
	if m != nil && m.Parameter != nil {
		return *m.Parameter
	}
	return ""
}

func (m *CodeGeneratorRequest) GetProtoFile() []*descriptor.FileDescriptorProto {
	if m != nil {
		return m.ProtoFile
	}
	return nil
}

func (m *CodeGeneratorRequest) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message.  If non-empty, code generation failed.  The plugin process
	// should exit with status code zero even if it reports an error in this way.
	//
	// This should be used to indicate errors in .proto files which prevent the
	// code generator from generating correct code.  Errors which indicate a
	// problem in protoc itself -- such as the input CodeGeneratorRequest being
	// unparseable -- should be reported by writing a message to stderr and
	// exiting with a non-zero status code.
	Error                *string                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	File                 []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

type xxx_CodeGeneratorResponse struct{ m *CodeGeneratorResponse }

func (m *CodeGeneratorResponse) ProtoReflect() protoreflect.Message {
	return xxx_CodeGeneratorResponse{m}
}
func (m xxx_CodeGeneratorResponse) Type() protoreflect.MessageType {
	return xxx_Plugin_ProtoFile_MessageTypes[2].Type
}
func (m xxx_CodeGeneratorResponse) KnownFields() protoreflect.KnownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[2].KnownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorResponse) UnknownFields() protoreflect.UnknownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[2].UnknownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorResponse) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *CodeGeneratorResponse) Reset()         { *m = CodeGeneratorResponse{} }
func (m *CodeGeneratorResponse) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse) ProtoMessage()    {}
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{2}
}

func (m *CodeGeneratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse.Merge(m, src)
}
func (m *CodeGeneratorResponse) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse.Size(m)
}
func (m *CodeGeneratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse proto.InternalMessageInfo

func (m *CodeGeneratorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist, and the
	// content here is to be inserted into that file at a defined insertion
	// point.  This feature allows a code generator to extend the output
	// produced by another code generator.  The original generator may provide
	// insertion points by placing special annotations in the file that look
	// like:
	//   @@protoc_insertion_point(NAME)
	// The annotation can have arbitrary text before and after it on the line,
	// which allows it to be placed in a comment.  NAME should be replaced with
	// an identifier naming the point -- this is what other generators will use
	// as the insertion_point.  Code inserted at this point will be placed
	// immediately above the line containing the insertion point (thus multiple
	// insertions to the same point will come out in the order they were added).
	// The double-@ is intended to make it unlikely that the generated code
	// could contain things that look like insertion points by accident.
	//
	// For example, the C++ code generator places the following line in the
	// .pb.h files that it generates:
	//   // @@protoc_insertion_point(namespace_scope)
	// This line appears within the scope of the file's package namespace, but
	// outside of any particular class.  Another plugin can then specify the
	// insertion_point "namespace_scope" to generate additional classes or
	// other declarations that should be placed in this scope.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace, the same whitespace will be added to every line of the
	// inserted text.  This is useful for languages like Python, where
	// indentation matters.  In these languages, the insertion point comment
	// should be indented the same amount as any inserted code will need to be
	// in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one which
	// inserts into it must both run as part of a single invocation of protoc.
	// Code generators are executed in the order in which they appear on the
	// command line.
	//
	// If |insertion_point| is present, |name| must also be present.
	InsertionPoint *string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint" json:"insertion_point,omitempty"`
	// The file contents.
	Content              *string  `protobuf:"bytes,15,opt,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_CodeGeneratorResponse_File struct{ m *CodeGeneratorResponse_File }

func (m *CodeGeneratorResponse_File) ProtoReflect() protoreflect.Message {
	return xxx_CodeGeneratorResponse_File{m}
}
func (m xxx_CodeGeneratorResponse_File) Type() protoreflect.MessageType {
	return xxx_Plugin_ProtoFile_MessageTypes[3].Type
}
func (m xxx_CodeGeneratorResponse_File) KnownFields() protoreflect.KnownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[3].KnownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorResponse_File) UnknownFields() protoreflect.UnknownFields {
	return xxx_Plugin_ProtoFile_MessageTypes[3].UnknownFieldsOf(m.m)
}
func (m xxx_CodeGeneratorResponse_File) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *CodeGeneratorResponse_File) Reset()         { *m = CodeGeneratorResponse_File{} }
func (m *CodeGeneratorResponse_File) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse_File) ProtoMessage()    {}
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{2, 0}
}

func (m *CodeGeneratorResponse_File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse_File.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse_File.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse_File.Merge(m, src)
}
func (m *CodeGeneratorResponse_File) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse_File.Size(m)
}
func (m *CodeGeneratorResponse_File) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse_File.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse_File proto.InternalMessageInfo

func (m *CodeGeneratorResponse_File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if m != nil && m.InsertionPoint != nil {
		return *m.InsertionPoint
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
	proto.RegisterFile("google/protobuf/compiler/plugin.proto", fileDescriptor_3562add825dafed5)
	proto.RegisterType((*Version)(nil), "google.protobuf.compiler.Version")
	proto.RegisterType((*CodeGeneratorRequest)(nil), "google.protobuf.compiler.CodeGeneratorRequest")
	proto.RegisterType((*CodeGeneratorResponse)(nil), "google.protobuf.compiler.CodeGeneratorResponse")
	proto.RegisterType((*CodeGeneratorResponse_File)(nil), "google.protobuf.compiler.CodeGeneratorResponse.File")
}

var fileDescriptor_3562add825dafed5 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6a, 0xd5, 0x40,
	0x14, 0xc6, 0x89, 0xbd, 0xb5, 0xe4, 0x54, 0x7a, 0xcb, 0x50, 0x65, 0x28, 0x5d, 0xc4, 0x8b, 0x62,
	0x56, 0x09, 0x5c, 0x04, 0x17, 0xee, 0x5a, 0x51, 0x17, 0x2e, 0x2e, 0x83, 0xb8, 0x10, 0x24, 0xa4,
	0xe9, 0xb9, 0xe9, 0x48, 0x32, 0x67, 0x9c, 0x99, 0x14, 0x7d, 0x51, 0xdf, 0xc3, 0x37, 0x90, 0xf9,
	0x93, 0x56, 0x2e, 0xde, 0x55, 0xf2, 0xfd, 0xbe, 0x33, 0x33, 0xe7, 0x7c, 0x1c, 0x78, 0xd9, 0x13,
	0xf5, 0x03, 0xd6, 0xda, 0x90, 0xa3, 0xeb, 0x69, 0x5b, 0x77, 0x34, 0x6a, 0x39, 0xa0, 0xa9, 0xf5,
	0x30, 0xf5, 0x52, 0x55, 0xc1, 0x60, 0x3c, 0x96, 0x55, 0x73, 0x59, 0x35, 0x97, 0x9d, 0x17, 0xbb,
	0x17, 0xdc, 0xa0, 0xed, 0x8c, 0xd4, 0x8e, 0x4c, 0xac, 0x5e, 0x75, 0x70, 0xf4, 0x05, 0x8d, 0x95,
	0xa4, 0xd8, 0x19, 0x1c, 0x8e, 0xed, 0x77, 0x32, 0x3c, 0x2b, 0xb2, 0xf2, 0x50, 0x44, 0x11, 0xa8,
	0x54, 0x64, 0xf8, 0xa3, 0x44, 0xbd, 0xf0, 0x54, 0xb7, 0xae, 0xbb, 0xe5, 0x07, 0x91, 0x06, 0xc1,
	0x9e, 0xc1, 0x63, 0x3b, 0x6d, 0xb7, 0xf2, 0x27, 0x5f, 0x14, 0x59, 0x99, 0x8b, 0xa4, 0x56, 0x7f,
	0x32, 0x38, 0xbb, 0xa2, 0x1b, 0xfc, 0x80, 0x0a, 0x4d, 0xeb, 0xc8, 0x08, 0xfc, 0x31, 0xa1, 0x75,
	0xac, 0x84, 0xd3, 0xad, 0x1c, 0xb0, 0x71, 0xd4, 0xf4, 0xd1, 0x43, 0x9e, 0x15, 0x07, 0x65, 0x2e,
	0x4e, 0x3c, 0xff, 0x4c, 0xe9, 0x04, 0xb2, 0x0b, 0xc8, 0x75, 0x6b, 0xda, 0x11, 0x1d, 0xc6, 0x56,
	0x72, 0xf1, 0x00, 0xd8, 0x15, 0x40, 0x18, 0xa7, 0xf1, 0xa7, 0xf8, 0xb2, 0x38, 0x28, 0x8f, 0xd7,
	0x2f, 0xaa, 0xdd, 0x58, 0xde, 0xcb, 0x01, 0xdf, 0xdd, 0x07, 0xb0, 0xf1, 0x58, 0xe4, 0xc1, 0xf5,
	0x0e, 0xfb, 0x04, 0xa7, 0x73, 0x70, 0xcd, 0x5d, 0xcc, 0x24, 0x8c, 0x77, 0xbc, 0x7e, 0x5e, 0xed,
	0x4b, 0xb8, 0x4a, 0xe1, 0x89, 0xe5, 0x4c, 0x12, 0x58, 0xfd, 0xce, 0xe0, 0xe9, 0xce, 0xcc, 0x56,
	0x93, 0xb2, 0xe8, 0xb3, 0x43, 0x63, 0x52, 0xce, 0xb9, 0x88, 0x82, 0x7d, 0x84, 0xc5, 0x3f, 0xcd,
	0xbf, 0xde, 0xff, 0xe2, 0x7f, 0x2f, 0x0d, 0xb3, 0x89, 0x70, 0xc3, 0xf9, 0x37, 0x58, 0x84, 0x79,
	0x18, 0x2c, 0x54, 0x3b, 0x62, 0x7a, 0x26, 0xfc, 0xb3, 0x57, 0xb0, 0x94, 0xca, 0xa2, 0x71, 0x92,
	0x54, 0xa3, 0x49, 0x2a, 0x97, 0xc2, 0x3c, 0xb9, 0xc7, 0x1b, 0x4f, 0x19, 0x87, 0xa3, 0x8e, 0x94,
	0x43, 0xe5, 0xf8, 0x32, 0x14, 0xcc, 0xf2, 0x12, 0xe1, 0xa2, 0xa3, 0x71, 0x6f, 0x7f, 0x97, 0x4f,
	0x36, 0x61, 0x37, 0x43, 0xbc, 0xf6, 0xeb, 0x9b, 0x5e, 0xba, 0xdb, 0xe9, 0xda, 0xdb, 0x75, 0x4f,
	0x43, 0xab, 0xfa, 0x87, 0x65, 0xbc, 0x5b, 0xd7, 0xee, 0x97, 0x46, 0x9b, 0xb6, 0xf9, 0x6d, 0xfc,
	0x34, 0xc1, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xfc, 0xb6, 0x26, 0xfc, 0x02, 0x00, 0x00,
}

func init() {
	xxx_Plugin_ProtoFile_FileDesc.Messages = xxx_Plugin_ProtoFile_MessageDescs[0:3]
	xxx_Plugin_ProtoFile_MessageDescs[2].Messages = xxx_Plugin_ProtoFile_MessageDescs[3:4]
	xxx_Plugin_ProtoFile_MessageDescs[1].Fields[2].MessageType = protoimpl.X.MessageTypeOf((*descriptor.FileDescriptorProto)(nil))
	xxx_Plugin_ProtoFile_MessageDescs[1].Fields[3].MessageType = xxx_Plugin_ProtoFile_MessageTypes[0].Type
	xxx_Plugin_ProtoFile_MessageDescs[2].Fields[1].MessageType = xxx_Plugin_ProtoFile_MessageTypes[3].Type
	var err error
	Plugin_ProtoFile, err = prototype.NewFile(&xxx_Plugin_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Plugin_ProtoFile protoreflect.FileDescriptor

var xxx_Plugin_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "google/protobuf/compiler/plugin.proto",
	Package: "google.protobuf.compiler",
	Imports: []protoreflect.FileImport{
		{FileDescriptor: prototype.PlaceholderFile("google/protobuf/descriptor.proto", "google.protobuf")},
	},
}
var xxx_Plugin_ProtoFile_MessageTypes = [4]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Plugin_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Version)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Plugin_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(CodeGeneratorRequest)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Plugin_ProtoFile_MessageDescs[2].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(CodeGeneratorResponse)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Plugin_ProtoFile_MessageDescs[3].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(CodeGeneratorResponse_File)
		},
	)},
}
var xxx_Plugin_ProtoFile_MessageDescs = [4]prototype.Message{
	{
		Name: "Version",
		Fields: []prototype.Field{
			{
				Name:        "major",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int32Kind,
				JSONName:    "major",
				IsPacked:    prototype.False,
			},
			{
				Name:        "minor",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int32Kind,
				JSONName:    "minor",
				IsPacked:    prototype.False,
			},
			{
				Name:        "patch",
				Number:      3,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int32Kind,
				JSONName:    "patch",
				IsPacked:    prototype.False,
			},
			{
				Name:        "suffix",
				Number:      4,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "suffix",
				IsPacked:    prototype.False,
			},
		},
	},
	{
		Name: "CodeGeneratorRequest",
		Fields: []prototype.Field{
			{
				Name:        "file_to_generate",
				Number:      1,
				Cardinality: protoreflect.Repeated,
				Kind:        protoreflect.StringKind,
				JSONName:    "fileToGenerate",
				IsPacked:    prototype.False,
			},
			{
				Name:        "parameter",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "parameter",
				IsPacked:    prototype.False,
			},
			{
				Name:        "proto_file",
				Number:      15,
				Cardinality: protoreflect.Repeated,
				Kind:        protoreflect.MessageKind,
				JSONName:    "protoFile",
				IsPacked:    prototype.False,
			},
			{
				Name:        "compiler_version",
				Number:      3,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "compilerVersion",
				IsPacked:    prototype.False,
			},
		},
	},
	{
		Name: "CodeGeneratorResponse",
		Fields: []prototype.Field{
			{
				Name:        "error",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "error",
				IsPacked:    prototype.False,
			},
			{
				Name:        "file",
				Number:      15,
				Cardinality: protoreflect.Repeated,
				Kind:        protoreflect.MessageKind,
				JSONName:    "file",
				IsPacked:    prototype.False,
			},
		},
	},
	{
		Name: "File",
		Fields: []prototype.Field{
			{
				Name:        "name",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "name",
				IsPacked:    prototype.False,
			},
			{
				Name:        "insertion_point",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "insertionPoint",
				IsPacked:    prototype.False,
			},
			{
				Name:        "content",
				Number:      15,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "content",
				IsPacked:    prototype.False,
			},
		},
	},
}
