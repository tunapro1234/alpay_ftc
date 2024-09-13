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

#include <geometry_msgs/msg/vector3_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Vector3Stamped", Vector3StampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/Vector3Stamped", Vector3StampedTypeSupport)
}

type Vector3Stamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Vector Vector3 `yaml:"vector"`
}

// NewVector3Stamped creates a new Vector3Stamped with default values.
func NewVector3Stamped() *Vector3Stamped {
	self := Vector3Stamped{}
	self.SetDefaults()
	return &self
}

func (t *Vector3Stamped) Clone() *Vector3Stamped {
	c := &Vector3Stamped{}
	c.Header = *t.Header.Clone()
	c.Vector = *t.Vector.Clone()
	return c
}

func (t *Vector3Stamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Vector3Stamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Vector.SetDefaults()
}

func (t *Vector3Stamped) GetTypeSupport() types.MessageTypeSupport {
	return Vector3StampedTypeSupport
}

// Vector3StampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Vector3StampedPublisher struct {
	*rclgo.Publisher
}

// NewVector3StampedPublisher creates and returns a new publisher for the
// Vector3Stamped
func NewVector3StampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Vector3StampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Vector3StampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Vector3StampedPublisher{pub}, nil
}

func (p *Vector3StampedPublisher) Publish(msg *Vector3Stamped) error {
	return p.Publisher.Publish(msg)
}

// Vector3StampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Vector3StampedSubscription struct {
	*rclgo.Subscription
}

// Vector3StampedSubscriptionCallback type is used to provide a subscription
// handler function for a Vector3StampedSubscription.
type Vector3StampedSubscriptionCallback func(msg *Vector3Stamped, info *rclgo.MessageInfo, err error)

// NewVector3StampedSubscription creates and returns a new subscription for the
// Vector3Stamped
func NewVector3StampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Vector3StampedSubscriptionCallback) (*Vector3StampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Vector3Stamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Vector3StampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Vector3StampedSubscription{sub}, nil
}

func (s *Vector3StampedSubscription) TakeMessage(out *Vector3Stamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneVector3StampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVector3StampedSlice(dst, src []Vector3Stamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Vector3StampedTypeSupport types.MessageTypeSupport = _Vector3StampedTypeSupport{}

type _Vector3StampedTypeSupport struct{}

func (t _Vector3StampedTypeSupport) New() types.Message {
	return NewVector3Stamped()
}

func (t _Vector3StampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Vector3Stamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__Vector3Stamped__create())
}

func (t _Vector3StampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Vector3Stamped__destroy((*C.geometry_msgs__msg__Vector3Stamped)(pointer_to_free))
}

func (t _Vector3StampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vector3Stamped)
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.vector), &m.Vector)
}

func (t _Vector3StampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vector3Stamped)
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	Vector3TypeSupport.AsGoStruct(&m.Vector, unsafe.Pointer(&mem.vector))
}

func (t _Vector3StampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Vector3Stamped())
}

type CVector3Stamped = C.geometry_msgs__msg__Vector3Stamped
type CVector3Stamped__Sequence = C.geometry_msgs__msg__Vector3Stamped__Sequence

func Vector3Stamped__Sequence_to_Go(goSlice *[]Vector3Stamped, cSlice CVector3Stamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vector3Stamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Vector3StampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Vector3Stamped__Sequence_to_C(cSlice *CVector3Stamped__Sequence, goSlice []Vector3Stamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Vector3Stamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Vector3StampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Vector3Stamped__Array_to_Go(goSlice []Vector3Stamped, cSlice []CVector3Stamped) {
	for i := 0; i < len(cSlice); i++ {
		Vector3StampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vector3Stamped__Array_to_C(cSlice []CVector3Stamped, goSlice []Vector3Stamped) {
	for i := 0; i < len(goSlice); i++ {
		Vector3StampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
