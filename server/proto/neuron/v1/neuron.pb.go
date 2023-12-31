// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: neuron/v1/neuron.proto

package neuronv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NeuronCompartmentType int32

const (
	// 未定義
	NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_UNSPECIFIED NeuronCompartmentType = 0
	// 細胞体
	NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_SOMA NeuronCompartmentType = 1
	// 軸索
	NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_AXON NeuronCompartmentType = 2
	// 基底樹状突起
	NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_BASAL_DENDRITE NeuronCompartmentType = 3
	// 尖端樹状突起
	NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_APICAL_DENDRITE NeuronCompartmentType = 4
)

// Enum value maps for NeuronCompartmentType.
var (
	NeuronCompartmentType_name = map[int32]string{
		0: "NEURON_COMPARTMENT_TYPE_UNSPECIFIED",
		1: "NEURON_COMPARTMENT_TYPE_SOMA",
		2: "NEURON_COMPARTMENT_TYPE_AXON",
		3: "NEURON_COMPARTMENT_TYPE_BASAL_DENDRITE",
		4: "NEURON_COMPARTMENT_TYPE_APICAL_DENDRITE",
	}
	NeuronCompartmentType_value = map[string]int32{
		"NEURON_COMPARTMENT_TYPE_UNSPECIFIED":     0,
		"NEURON_COMPARTMENT_TYPE_SOMA":            1,
		"NEURON_COMPARTMENT_TYPE_AXON":            2,
		"NEURON_COMPARTMENT_TYPE_BASAL_DENDRITE":  3,
		"NEURON_COMPARTMENT_TYPE_APICAL_DENDRITE": 4,
	}
)

func (x NeuronCompartmentType) Enum() *NeuronCompartmentType {
	p := new(NeuronCompartmentType)
	*p = x
	return p
}

func (x NeuronCompartmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NeuronCompartmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_neuron_v1_neuron_proto_enumTypes[0].Descriptor()
}

func (NeuronCompartmentType) Type() protoreflect.EnumType {
	return &file_neuron_v1_neuron_proto_enumTypes[0]
}

func (x NeuronCompartmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NeuronCompartmentType.Descriptor instead.
func (NeuronCompartmentType) EnumDescriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{0}
}

type ListNeuronsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 利用可能なニューロン名の一覧
	NeuronNames []string `protobuf:"bytes,1,rep,name=neuron_names,json=neuronNames,proto3" json:"neuron_names,omitempty"`
}

func (x *ListNeuronsResponse) Reset() {
	*x = ListNeuronsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNeuronsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNeuronsResponse) ProtoMessage() {}

func (x *ListNeuronsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNeuronsResponse.ProtoReflect.Descriptor instead.
func (*ListNeuronsResponse) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{0}
}

func (x *ListNeuronsResponse) GetNeuronNames() []string {
	if x != nil {
		return x.NeuronNames
	}
	return nil
}

type GetNeuronShapeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取得したいニューロンの名前
	NeuronName string `protobuf:"bytes,1,opt,name=neuron_name,json=neuronName,proto3" json:"neuron_name,omitempty"`
}

func (x *GetNeuronShapeRequest) Reset() {
	*x = GetNeuronShapeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNeuronShapeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNeuronShapeRequest) ProtoMessage() {}

func (x *GetNeuronShapeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNeuronShapeRequest.ProtoReflect.Descriptor instead.
func (*GetNeuronShapeRequest) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{1}
}

func (x *GetNeuronShapeRequest) GetNeuronName() string {
	if x != nil {
		return x.NeuronName
	}
	return ""
}

type GetNeuronShapeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ニューロンの形状・空間情報
	NeuronCompartments []*NeuronCompartment `protobuf:"bytes,1,rep,name=neuron_compartments,json=neuronCompartments,proto3" json:"neuron_compartments,omitempty"`
}

func (x *GetNeuronShapeResponse) Reset() {
	*x = GetNeuronShapeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNeuronShapeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNeuronShapeResponse) ProtoMessage() {}

func (x *GetNeuronShapeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNeuronShapeResponse.ProtoReflect.Descriptor instead.
func (*GetNeuronShapeResponse) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{2}
}

func (x *GetNeuronShapeResponse) GetNeuronCompartments() []*NeuronCompartment {
	if x != nil {
		return x.NeuronCompartments
	}
	return nil
}

type GetMembranePotentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 各コンパートメントの膜電位を取得したいニューロンの名前
	NeuronName string `protobuf:"bytes,1,opt,name=neuron_name,json=neuronName,proto3" json:"neuron_name,omitempty"`
}

func (x *GetMembranePotentialsRequest) Reset() {
	*x = GetMembranePotentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembranePotentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembranePotentialsRequest) ProtoMessage() {}

func (x *GetMembranePotentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembranePotentialsRequest.ProtoReflect.Descriptor instead.
func (*GetMembranePotentialsRequest) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{3}
}

func (x *GetMembranePotentialsRequest) GetNeuronName() string {
	if x != nil {
		return x.NeuronName
	}
	return ""
}

type GetMembranePotentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 時間ステップ
	TimeStep float32 `protobuf:"fixed32,1,opt,name=time_step,json=timeStep,proto3" json:"time_step,omitempty"`
	// 各コンパートメントの膜電位
	MembranePotentials []float32 `protobuf:"fixed32,2,rep,packed,name=membrane_potentials,json=membranePotentials,proto3" json:"membrane_potentials,omitempty"`
}

func (x *GetMembranePotentialsResponse) Reset() {
	*x = GetMembranePotentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembranePotentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembranePotentialsResponse) ProtoMessage() {}

func (x *GetMembranePotentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembranePotentialsResponse.ProtoReflect.Descriptor instead.
func (*GetMembranePotentialsResponse) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{4}
}

func (x *GetMembranePotentialsResponse) GetTimeStep() float32 {
	if x != nil {
		return x.TimeStep
	}
	return 0
}

func (x *GetMembranePotentialsResponse) GetMembranePotentials() []float32 {
	if x != nil {
		return x.MembranePotentials
	}
	return nil
}

type NeuronCompartment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// （同一ニューロンで一意な）コンパートメントの id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// コンパートメントの種類
	Type NeuronCompartmentType `protobuf:"varint,2,opt,name=type,proto3,enum=neuron.v1.NeuronCompartmentType" json:"type,omitempty"`
	// コンパートメントの x 座標の位置
	PositionX float32 `protobuf:"fixed32,3,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	// コンパートメントの y 座標の位置
	PositionY float32 `protobuf:"fixed32,4,opt,name=position_y,json=positionY,proto3" json:"position_y,omitempty"`
	// コンパートメントの z 座標の位置
	PositionZ float32 `protobuf:"fixed32,5,opt,name=position_z,json=positionZ,proto3" json:"position_z,omitempty"`
	// コンパートメントの半径
	Radius float32 `protobuf:"fixed32,6,opt,name=radius,proto3" json:"radius,omitempty"`
	// 親コンパートメントの id
	ParentId int64 `protobuf:"varint,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *NeuronCompartment) Reset() {
	*x = NeuronCompartment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neuron_v1_neuron_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeuronCompartment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeuronCompartment) ProtoMessage() {}

func (x *NeuronCompartment) ProtoReflect() protoreflect.Message {
	mi := &file_neuron_v1_neuron_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeuronCompartment.ProtoReflect.Descriptor instead.
func (*NeuronCompartment) Descriptor() ([]byte, []int) {
	return file_neuron_v1_neuron_proto_rawDescGZIP(), []int{5}
}

func (x *NeuronCompartment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NeuronCompartment) GetType() NeuronCompartmentType {
	if x != nil {
		return x.Type
	}
	return NeuronCompartmentType_NEURON_COMPARTMENT_TYPE_UNSPECIFIED
}

func (x *NeuronCompartment) GetPositionX() float32 {
	if x != nil {
		return x.PositionX
	}
	return 0
}

func (x *NeuronCompartment) GetPositionY() float32 {
	if x != nil {
		return x.PositionY
	}
	return 0
}

func (x *NeuronCompartment) GetPositionZ() float32 {
	if x != nil {
		return x.PositionZ
	}
	return 0
}

func (x *NeuronCompartment) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *NeuronCompartment) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

var File_neuron_v1_neuron_proto protoreflect.FileDescriptor

var file_neuron_v1_neuron_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x75, 0x72,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x38, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x75, 0x72, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x75, 0x72, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f,
	0x6e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x13, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65,
	0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x6e, 0x65, 0x75, 0x72, 0x6f,
	0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3f, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6d,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x65, 0x70, 0x12, 0x2f, 0x0a, 0x13,
	0x6d, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x12, 0x6d, 0x65, 0x6d, 0x62, 0x72,
	0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0xeb, 0x01,
	0x0a, 0x11, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x75, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0xdd, 0x01, 0x0a, 0x15,
	0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x4e, 0x45, 0x55, 0x52, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20,
	0x0a, 0x1c, 0x4e, 0x45, 0x55, 0x52, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x4d, 0x41, 0x10, 0x01,
	0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x45, 0x55, 0x52, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x58, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x4e, 0x45, 0x55, 0x52, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41,
	0x53, 0x41, 0x4c, 0x5f, 0x44, 0x45, 0x4e, 0x44, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x2b,
	0x0a, 0x27, 0x4e, 0x45, 0x55, 0x52, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x43, 0x41, 0x4c,
	0x5f, 0x44, 0x45, 0x4e, 0x44, 0x52, 0x49, 0x54, 0x45, 0x10, 0x04, 0x32, 0x9b, 0x02, 0x0a, 0x0d,
	0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f,
	0x6e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x72, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xa6, 0x01, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4e, 0x65, 0x75,
	0x72, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x31, 0x31, 0x6b, 0x6d, 0x2f, 0x6e, 0x65,
	0x75, 0x72, 0x6f, 0x6e, 0x2d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x75,
	0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x6e, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neuron_v1_neuron_proto_rawDescOnce sync.Once
	file_neuron_v1_neuron_proto_rawDescData = file_neuron_v1_neuron_proto_rawDesc
)

func file_neuron_v1_neuron_proto_rawDescGZIP() []byte {
	file_neuron_v1_neuron_proto_rawDescOnce.Do(func() {
		file_neuron_v1_neuron_proto_rawDescData = protoimpl.X.CompressGZIP(file_neuron_v1_neuron_proto_rawDescData)
	})
	return file_neuron_v1_neuron_proto_rawDescData
}

var file_neuron_v1_neuron_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_neuron_v1_neuron_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_neuron_v1_neuron_proto_goTypes = []interface{}{
	(NeuronCompartmentType)(0),            // 0: neuron.v1.NeuronCompartmentType
	(*ListNeuronsResponse)(nil),           // 1: neuron.v1.ListNeuronsResponse
	(*GetNeuronShapeRequest)(nil),         // 2: neuron.v1.GetNeuronShapeRequest
	(*GetNeuronShapeResponse)(nil),        // 3: neuron.v1.GetNeuronShapeResponse
	(*GetMembranePotentialsRequest)(nil),  // 4: neuron.v1.GetMembranePotentialsRequest
	(*GetMembranePotentialsResponse)(nil), // 5: neuron.v1.GetMembranePotentialsResponse
	(*NeuronCompartment)(nil),             // 6: neuron.v1.NeuronCompartment
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
}
var file_neuron_v1_neuron_proto_depIdxs = []int32{
	6, // 0: neuron.v1.GetNeuronShapeResponse.neuron_compartments:type_name -> neuron.v1.NeuronCompartment
	0, // 1: neuron.v1.NeuronCompartment.type:type_name -> neuron.v1.NeuronCompartmentType
	7, // 2: neuron.v1.NeuronService.ListNeurons:input_type -> google.protobuf.Empty
	2, // 3: neuron.v1.NeuronService.GetNeuronShape:input_type -> neuron.v1.GetNeuronShapeRequest
	4, // 4: neuron.v1.NeuronService.GetMembranePotentials:input_type -> neuron.v1.GetMembranePotentialsRequest
	1, // 5: neuron.v1.NeuronService.ListNeurons:output_type -> neuron.v1.ListNeuronsResponse
	3, // 6: neuron.v1.NeuronService.GetNeuronShape:output_type -> neuron.v1.GetNeuronShapeResponse
	5, // 7: neuron.v1.NeuronService.GetMembranePotentials:output_type -> neuron.v1.GetMembranePotentialsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_neuron_v1_neuron_proto_init() }
func file_neuron_v1_neuron_proto_init() {
	if File_neuron_v1_neuron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neuron_v1_neuron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNeuronsResponse); i {
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
		file_neuron_v1_neuron_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNeuronShapeRequest); i {
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
		file_neuron_v1_neuron_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNeuronShapeResponse); i {
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
		file_neuron_v1_neuron_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembranePotentialsRequest); i {
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
		file_neuron_v1_neuron_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembranePotentialsResponse); i {
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
		file_neuron_v1_neuron_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeuronCompartment); i {
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
			RawDescriptor: file_neuron_v1_neuron_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neuron_v1_neuron_proto_goTypes,
		DependencyIndexes: file_neuron_v1_neuron_proto_depIdxs,
		EnumInfos:         file_neuron_v1_neuron_proto_enumTypes,
		MessageInfos:      file_neuron_v1_neuron_proto_msgTypes,
	}.Build()
	File_neuron_v1_neuron_proto = out.File
	file_neuron_v1_neuron_proto_rawDesc = nil
	file_neuron_v1_neuron_proto_goTypes = nil
	file_neuron_v1_neuron_proto_depIdxs = nil
}
