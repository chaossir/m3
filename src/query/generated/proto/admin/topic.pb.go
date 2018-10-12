// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/query/generated/proto/admin/topic.proto

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package admin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import topicpb "github.com/m3db/m3/src/msg/generated/proto/topicpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TopicGetResponse struct {
	Topic   *topicpb.Topic `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Version uint32         `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *TopicGetResponse) Reset()                    { *m = TopicGetResponse{} }
func (m *TopicGetResponse) String() string            { return proto.CompactTextString(m) }
func (*TopicGetResponse) ProtoMessage()               {}
func (*TopicGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorTopic, []int{0} }

func (m *TopicGetResponse) GetTopic() *topicpb.Topic {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *TopicGetResponse) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type TopicInitRequest struct {
	NumberOfShards uint32 `protobuf:"varint,1,opt,name=number_of_shards,json=numberOfShards,proto3" json:"number_of_shards,omitempty"`
}

func (m *TopicInitRequest) Reset()                    { *m = TopicInitRequest{} }
func (m *TopicInitRequest) String() string            { return proto.CompactTextString(m) }
func (*TopicInitRequest) ProtoMessage()               {}
func (*TopicInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorTopic, []int{1} }

func (m *TopicInitRequest) GetNumberOfShards() uint32 {
	if m != nil {
		return m.NumberOfShards
	}
	return 0
}

type TopicAddRequest struct {
	ConsumerService *topicpb.ConsumerService `protobuf:"bytes,1,opt,name=consumer_service,json=consumerService" json:"consumer_service,omitempty"`
}

func (m *TopicAddRequest) Reset()                    { *m = TopicAddRequest{} }
func (m *TopicAddRequest) String() string            { return proto.CompactTextString(m) }
func (*TopicAddRequest) ProtoMessage()               {}
func (*TopicAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorTopic, []int{2} }

func (m *TopicAddRequest) GetConsumerService() *topicpb.ConsumerService {
	if m != nil {
		return m.ConsumerService
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicGetResponse)(nil), "admin.TopicGetResponse")
	proto.RegisterType((*TopicInitRequest)(nil), "admin.TopicInitRequest")
	proto.RegisterType((*TopicAddRequest)(nil), "admin.TopicAddRequest")
}
func (m *TopicGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Topic != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTopic(dAtA, i, uint64(m.Topic.Size()))
		n1, err := m.Topic.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTopic(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func (m *TopicInitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicInitRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NumberOfShards != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTopic(dAtA, i, uint64(m.NumberOfShards))
	}
	return i, nil
}

func (m *TopicAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ConsumerService != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTopic(dAtA, i, uint64(m.ConsumerService.Size()))
		n2, err := m.ConsumerService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintTopic(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TopicGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Topic != nil {
		l = m.Topic.Size()
		n += 1 + l + sovTopic(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovTopic(uint64(m.Version))
	}
	return n
}

func (m *TopicInitRequest) Size() (n int) {
	var l int
	_ = l
	if m.NumberOfShards != 0 {
		n += 1 + sovTopic(uint64(m.NumberOfShards))
	}
	return n
}

func (m *TopicAddRequest) Size() (n int) {
	var l int
	_ = l
	if m.ConsumerService != nil {
		l = m.ConsumerService.Size()
		n += 1 + l + sovTopic(uint64(l))
	}
	return n
}

func sovTopic(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTopic(x uint64) (n int) {
	return sovTopic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TopicGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Topic == nil {
				m.Topic = &topicpb.Topic{}
			}
			if err := m.Topic.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopic
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
func (m *TopicInitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicInitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicInitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfShards", wireType)
			}
			m.NumberOfShards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumberOfShards |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopic
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
func (m *TopicAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsumerService == nil {
				m.ConsumerService = &topicpb.ConsumerService{}
			}
			if err := m.ConsumerService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopic
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
func skipTopic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
				return 0, ErrInvalidLengthTopic
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTopic
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
				next, err := skipTopic(dAtA[start:])
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
	ErrInvalidLengthTopic = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopic   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/query/generated/proto/admin/topic.proto", fileDescriptorTopic)
}

var fileDescriptorTopic = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x5d, 0xa1, 0x0a, 0x2b, 0x6d, 0x43, 0x4e, 0xc1, 0x43, 0x28, 0xc1, 0x43, 0x4e, 0x59,
	0x30, 0x27, 0x41, 0x04, 0xed, 0x41, 0x3c, 0x09, 0x5b, 0xf1, 0x1a, 0x92, 0xdd, 0x69, 0xba, 0x87,
	0xdd, 0x4d, 0xf7, 0x4f, 0xc1, 0xb7, 0xf0, 0xb1, 0x3c, 0xfa, 0x08, 0x12, 0x5f, 0x44, 0xdc, 0xa4,
	0x62, 0xe9, 0x71, 0x7e, 0xf3, 0xfd, 0x3e, 0x86, 0xc1, 0x77, 0xad, 0x70, 0x1b, 0xdf, 0x14, 0x4c,
	0x4b, 0x22, 0x4b, 0xde, 0x10, 0x59, 0x12, 0x6b, 0x18, 0xd9, 0x7a, 0x30, 0x6f, 0xa4, 0x05, 0x05,
	0xa6, 0x76, 0xc0, 0x49, 0x67, 0xb4, 0xd3, 0xa4, 0xe6, 0x52, 0x28, 0xe2, 0x74, 0x27, 0x58, 0x11,
	0x48, 0x3c, 0x09, 0xe8, 0xf2, 0xe6, 0xb8, 0x46, 0xda, 0xf6, 0xc8, 0x0f, 0x66, 0xd7, 0xfc, 0x6f,
	0xc8, 0x28, 0x8e, 0x5e, 0x7e, 0xc7, 0x47, 0x70, 0x14, 0x6c, 0xa7, 0x95, 0x85, 0xf8, 0x0a, 0x4f,
	0x42, 0x24, 0x41, 0x0b, 0x94, 0x5f, 0x5c, 0xcf, 0x8a, 0x51, 0x2c, 0x42, 0x92, 0x0e, 0xcb, 0x38,
	0xc1, 0xe7, 0x3b, 0x30, 0x56, 0x68, 0x95, 0x9c, 0x2e, 0x50, 0x3e, 0xa5, 0xfb, 0x31, 0xbb, 0x1d,
	0x3b, 0x9f, 0x94, 0x70, 0x14, 0xb6, 0x1e, 0xac, 0x8b, 0x73, 0x1c, 0x29, 0x2f, 0x1b, 0x30, 0x95,
	0x5e, 0x57, 0x76, 0x53, 0x1b, 0x6e, 0x43, 0xfd, 0x94, 0xce, 0x06, 0xfe, 0xbc, 0x5e, 0x05, 0x9a,
	0xbd, 0xe2, 0x79, 0xb0, 0xef, 0x39, 0xdf, 0xcb, 0x4b, 0x1c, 0x31, 0xad, 0xac, 0x97, 0x60, 0x2a,
	0x0b, 0x66, 0x27, 0x18, 0x8c, 0xb7, 0x25, 0x7f, 0xb7, 0x2d, 0xc7, 0xc0, 0x6a, 0xd8, 0xd3, 0x39,
	0x3b, 0x04, 0x0f, 0xd1, 0x47, 0x9f, 0xa2, 0xcf, 0x3e, 0x45, 0x5f, 0x7d, 0x8a, 0xde, 0xbf, 0xd3,
	0x93, 0xe6, 0x2c, 0xbc, 0xa0, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x22, 0xb0, 0x99, 0x86,
	0x01, 0x00, 0x00,
}
