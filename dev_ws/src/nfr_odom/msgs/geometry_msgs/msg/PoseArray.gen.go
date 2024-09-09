// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "nfr_odom/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/pose_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PoseArray", PoseArrayTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/PoseArray", PoseArrayTypeSupport)
}

type PoseArray struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Poses []Pose `yaml:"poses"`
}

// NewPoseArray creates a new PoseArray with default values.
func NewPoseArray() *PoseArray {
	self := PoseArray{}
	self.SetDefaults()
	return &self
}

func (t *PoseArray) Clone() *PoseArray {
	c := &PoseArray{}
	c.Header = *t.Header.Clone()
	if t.Poses != nil {
		c.Poses = make([]Pose, len(t.Poses))
		ClonePoseSlice(c.Poses, t.Poses)
	}
	return c
}

func (t *PoseArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PoseArray) SetDefaults() {
	t.Header.SetDefaults()
	t.Poses = nil
}

func (t *PoseArray) GetTypeSupport() types.MessageTypeSupport {
	return PoseArrayTypeSupport
}

// PoseArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PoseArrayPublisher struct {
	*rclgo.Publisher
}

// NewPoseArrayPublisher creates and returns a new publisher for the
// PoseArray
func NewPoseArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PoseArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PoseArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PoseArrayPublisher{pub}, nil
}

func (p *PoseArrayPublisher) Publish(msg *PoseArray) error {
	return p.Publisher.Publish(msg)
}

// PoseArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PoseArraySubscription struct {
	*rclgo.Subscription
}

// PoseArraySubscriptionCallback type is used to provide a subscription
// handler function for a PoseArraySubscription.
type PoseArraySubscriptionCallback func(msg *PoseArray, info *rclgo.MessageInfo, err error)

// NewPoseArraySubscription creates and returns a new subscription for the
// PoseArray
func NewPoseArraySubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PoseArraySubscriptionCallback) (*PoseArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PoseArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PoseArrayTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PoseArraySubscription{sub}, nil
}

func (s *PoseArraySubscription) TakeMessage(out *PoseArray) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePoseArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseArraySlice(dst, src []PoseArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseArrayTypeSupport types.MessageTypeSupport = _PoseArrayTypeSupport{}

type _PoseArrayTypeSupport struct{}

func (t _PoseArrayTypeSupport) New() types.Message {
	return NewPoseArray()
}

func (t _PoseArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseArray
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseArray__create())
}

func (t _PoseArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseArray__destroy((*C.geometry_msgs__msg__PoseArray)(pointer_to_free))
}

func (t _PoseArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PoseArray)
	mem := (*C.geometry_msgs__msg__PoseArray)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	Pose__Sequence_to_C(&mem.poses, m.Poses)
}

func (t _PoseArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PoseArray)
	mem := (*C.geometry_msgs__msg__PoseArray)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	Pose__Sequence_to_Go(&m.Poses, mem.poses)
}

func (t _PoseArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseArray())
}

type CPoseArray = C.geometry_msgs__msg__PoseArray
type CPoseArray__Sequence = C.geometry_msgs__msg__PoseArray__Sequence

func PoseArray__Sequence_to_Go(goSlice *[]PoseArray, cSlice CPoseArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseArray, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PoseArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PoseArray__Sequence_to_C(cSlice *CPoseArray__Sequence, goSlice []PoseArray) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseArray)(C.malloc(C.sizeof_struct_geometry_msgs__msg__PoseArray * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PoseArrayTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PoseArray__Array_to_Go(goSlice []PoseArray, cSlice []CPoseArray) {
	for i := 0; i < len(cSlice); i++ {
		PoseArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PoseArray__Array_to_C(cSlice []CPoseArray, goSlice []PoseArray) {
	for i := 0; i < len(goSlice); i++ {
		PoseArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
