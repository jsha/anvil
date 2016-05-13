// Code generated by protoc-gen-go.
// source: core/proto/core.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	core/proto/core.proto

It has these top-level messages:
	Challenge
	ValidationRecord
	ProblemDetails
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type Challenge struct {
	Id                *int64              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type              *string             `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Status            *string             `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Uri               *string             `protobuf:"bytes,9,opt,name=uri" json:"uri,omitempty"`
	Token             *string             `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	AccountKey        *string             `protobuf:"bytes,4,opt,name=accountKey" json:"accountKey,omitempty"`
	KeyAuthorization  *string             `protobuf:"bytes,5,opt,name=keyAuthorization" json:"keyAuthorization,omitempty"`
	Validationrecords []*ValidationRecord `protobuf:"bytes,10,rep,name=validationrecords" json:"validationrecords,omitempty"`
	Error             *ProblemDetails     `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *Challenge) Reset()                    { *m = Challenge{} }
func (m *Challenge) String() string            { return proto1.CompactTextString(m) }
func (*Challenge) ProtoMessage()               {}
func (*Challenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Challenge) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Challenge) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Challenge) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Challenge) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Challenge) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Challenge) GetAccountKey() string {
	if m != nil && m.AccountKey != nil {
		return *m.AccountKey
	}
	return ""
}

func (m *Challenge) GetKeyAuthorization() string {
	if m != nil && m.KeyAuthorization != nil {
		return *m.KeyAuthorization
	}
	return ""
}

func (m *Challenge) GetValidationrecords() []*ValidationRecord {
	if m != nil {
		return m.Validationrecords
	}
	return nil
}

func (m *Challenge) GetError() *ProblemDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

type ValidationRecord struct {
	Hostname          *string  `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port              *string  `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	AddressesResolved []string `protobuf:"bytes,3,rep,name=addressesResolved" json:"addressesResolved,omitempty"`
	AddressUsed       *string  `protobuf:"bytes,4,opt,name=addressUsed" json:"addressUsed,omitempty"`
	Authorities       []string `protobuf:"bytes,5,rep,name=authorities" json:"authorities,omitempty"`
	Url               *string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *ValidationRecord) Reset()                    { *m = ValidationRecord{} }
func (m *ValidationRecord) String() string            { return proto1.CompactTextString(m) }
func (*ValidationRecord) ProtoMessage()               {}
func (*ValidationRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidationRecord) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ValidationRecord) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *ValidationRecord) GetAddressesResolved() []string {
	if m != nil {
		return m.AddressesResolved
	}
	return nil
}

func (m *ValidationRecord) GetAddressUsed() string {
	if m != nil && m.AddressUsed != nil {
		return *m.AddressUsed
	}
	return ""
}

func (m *ValidationRecord) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *ValidationRecord) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type ProblemDetails struct {
	ProblemType      *string `protobuf:"bytes,1,opt,name=problemType" json:"problemType,omitempty"`
	Detail           *string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	HttpStatus       *int32  `protobuf:"varint,3,opt,name=httpStatus" json:"httpStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProblemDetails) Reset()                    { *m = ProblemDetails{} }
func (m *ProblemDetails) String() string            { return proto1.CompactTextString(m) }
func (*ProblemDetails) ProtoMessage()               {}
func (*ProblemDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProblemDetails) GetProblemType() string {
	if m != nil && m.ProblemType != nil {
		return *m.ProblemType
	}
	return ""
}

func (m *ProblemDetails) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *ProblemDetails) GetHttpStatus() int32 {
	if m != nil && m.HttpStatus != nil {
		return *m.HttpStatus
	}
	return 0
}

func init() {
	proto1.RegisterType((*Challenge)(nil), "core.Challenge")
	proto1.RegisterType((*ValidationRecord)(nil), "core.ValidationRecord")
	proto1.RegisterType((*ProblemDetails)(nil), "core.ProblemDetails")
}

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x51, 0x4b, 0x6f, 0xf2, 0x30,
	0x10, 0x54, 0x08, 0x81, 0x2f, 0x9b, 0xaf, 0x08, 0xd2, 0x87, 0xdc, 0x1b, 0xa2, 0x17, 0x4e, 0xa0,
	0xf2, 0x0f, 0xfa, 0xb8, 0x54, 0xbd, 0x54, 0xf4, 0x71, 0xe8, 0xcd, 0xc5, 0xab, 0xc6, 0xc2, 0xc4,
	0x91, 0xbd, 0x41, 0xa2, 0xe7, 0xfe, 0xc7, 0xfe, 0x9d, 0xfa, 0x01, 0x95, 0xaa, 0xde, 0x76, 0x66,
	0xbc, 0xab, 0x99, 0x31, 0x9c, 0xae, 0xb4, 0xc1, 0x79, 0x63, 0x34, 0xe9, 0xb9, 0x1f, 0x67, 0x61,
	0x2c, 0xbb, 0x7e, 0x9e, 0x7c, 0x25, 0x90, 0xdf, 0x54, 0x5c, 0x29, 0xac, 0xdf, 0xb1, 0x04, 0xe8,
	0x48, 0xc1, 0x92, 0x71, 0x32, 0x4d, 0xcb, 0xff, 0xd0, 0xa5, 0x5d, 0x83, 0xac, 0xe3, 0x50, 0x5e,
	0x0e, 0xa0, 0x67, 0x89, 0x53, 0x6b, 0x59, 0x2f, 0xe0, 0x02, 0xd2, 0xd6, 0x48, 0x96, 0x07, 0x70,
	0x04, 0x19, 0xe9, 0x35, 0xd6, 0x2c, 0x0d, 0xd0, 0x9d, 0xe1, 0xab, 0x95, 0x6e, 0x6b, 0xba, 0xc7,
	0x1d, 0xeb, 0x06, 0x8e, 0xc1, 0x70, 0x8d, 0xbb, 0xab, 0x96, 0x2a, 0x6d, 0xe4, 0x07, 0x27, 0xa9,
	0x6b, 0x96, 0x05, 0xe5, 0x12, 0x46, 0x5b, 0xae, 0xa4, 0x08, 0x9c, 0x41, 0xe7, 0x4a, 0x58, 0x06,
	0xe3, 0x74, 0x5a, 0x2c, 0xce, 0x66, 0xc1, 0xef, 0xcb, 0x8f, 0xbc, 0x0c, 0x72, 0x79, 0x01, 0x19,
	0x1a, 0xa3, 0x0d, 0xeb, 0xbb, 0x0b, 0xc5, 0xe2, 0x24, 0x3e, 0x7b, 0x30, 0xfa, 0x4d, 0xe1, 0xe6,
	0x16, 0x89, 0x4b, 0x65, 0x27, 0x9f, 0x09, 0x0c, 0xff, 0x6c, 0x0e, 0xe1, 0x5f, 0xa5, 0x2d, 0xd5,
	0x7c, 0x83, 0x21, 0x66, 0xee, 0x63, 0x36, 0xda, 0xd0, 0x3e, 0xe6, 0x39, 0x8c, 0xb8, 0x10, 0x06,
	0xad, 0x45, 0xbb, 0x44, 0xab, 0xd5, 0x16, 0x85, 0x4b, 0x95, 0x3a, 0xe9, 0x18, 0x8a, 0xbd, 0xf4,
	0x6c, 0x1d, 0x19, 0x63, 0x79, 0x32, 0x66, 0x22, 0x89, 0xd6, 0x25, 0x4a, 0x0f, 0xdd, 0xa8, 0x58,
	0xd4, 0xe4, 0x0e, 0x06, 0xbf, 0x8d, 0xf9, 0x9d, 0x26, 0x32, 0x4f, 0xbe, 0xdf, 0xe4, 0xd0, 0xaf,
	0x08, 0xfa, 0xde, 0x88, 0xeb, 0xb0, 0x22, 0x6a, 0x1e, 0x63, 0xe7, 0xbe, 0xd7, 0xec, 0xba, 0xff,
	0x9a, 0x85, 0xaf, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x33, 0x0d, 0xc9, 0xd2, 0x01, 0x00,
	0x00,
}
