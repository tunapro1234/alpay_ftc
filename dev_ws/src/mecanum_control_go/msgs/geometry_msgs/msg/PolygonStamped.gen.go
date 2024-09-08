// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "mecanum_control_go/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/polygon_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PolygonStamped", PolygonStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/PolygonStamped", PolygonStampedTypeSupport)
}

type PolygonStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Polygon Polygon `yaml:"polygon"`
}

// NewPolygonStamped creates a new PolygonStamped with default values.
func NewPolygonStamped() *PolygonStamped {
	self := PolygonStamped{}
	self.SetDefaults()
	return &self
}

func (t *PolygonStamped) Clone() *PolygonStamped {
	c := &PolygonStamped{}
	c.Header = *t.Header.Clone()
	c.Polygon = *t.Polygon.Clone()
	return c
}

func (t *PolygonStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PolygonStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Polygon.SetDefaults()
}

func (t *PolygonStamped) GetTypeSupport() types.MessageTypeSupport {
	return PolygonStampedTypeSupport
}

// PolygonStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PolygonStampedPublisher struct {
	*rclgo.Publisher
}

// NewPolygonStampedPublisher creates and returns a new publisher for the
// PolygonStamped
func NewPolygonStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PolygonStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PolygonStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PolygonStampedPublisher{pub}, nil
}

func (p *PolygonStampedPublisher) Publish(msg *PolygonStamped) error {
	return p.Publisher.Publish(msg)
}

// PolygonStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PolygonStampedSubscription struct {
	*rclgo.Subscription
}

// PolygonStampedSubscriptionCallback type is used to provide a subscription
// handler function for a PolygonStampedSubscription.
type PolygonStampedSubscriptionCallback func(msg *PolygonStamped, info *rclgo.MessageInfo, err error)

// NewPolygonStampedSubscription creates and returns a new subscription for the
// PolygonStamped
func NewPolygonStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PolygonStampedSubscriptionCallback) (*PolygonStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PolygonStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PolygonStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PolygonStampedSubscription{sub}, nil
}

func (s *PolygonStampedSubscription) TakeMessage(out *PolygonStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePolygonStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePolygonStampedSlice(dst, src []PolygonStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PolygonStampedTypeSupport types.MessageTypeSupport = _PolygonStampedTypeSupport{}

type _PolygonStampedTypeSupport struct{}

func (t _PolygonStampedTypeSupport) New() types.Message {
	return NewPolygonStamped()
}

func (t _PolygonStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PolygonStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__PolygonStamped__create())
}

func (t _PolygonStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PolygonStamped__destroy((*C.geometry_msgs__msg__PolygonStamped)(pointer_to_free))
}

func (t _PolygonStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PolygonStamped)
	mem := (*C.geometry_msgs__msg__PolygonStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	PolygonTypeSupport.AsCStruct(unsafe.Pointer(&mem.polygon), &m.Polygon)
}

func (t _PolygonStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PolygonStamped)
	mem := (*C.geometry_msgs__msg__PolygonStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	PolygonTypeSupport.AsGoStruct(&m.Polygon, unsafe.Pointer(&mem.polygon))
}

func (t _PolygonStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PolygonStamped())
}

type CPolygonStamped = C.geometry_msgs__msg__PolygonStamped
type CPolygonStamped__Sequence = C.geometry_msgs__msg__PolygonStamped__Sequence

func PolygonStamped__Sequence_to_Go(goSlice *[]PolygonStamped, cSlice CPolygonStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PolygonStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PolygonStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PolygonStamped__Sequence_to_C(cSlice *CPolygonStamped__Sequence, goSlice []PolygonStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PolygonStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__PolygonStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PolygonStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PolygonStamped__Array_to_Go(goSlice []PolygonStamped, cSlice []CPolygonStamped) {
	for i := 0; i < len(cSlice); i++ {
		PolygonStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PolygonStamped__Array_to_C(cSlice []CPolygonStamped, goSlice []PolygonStamped) {
	for i := 0; i < len(goSlice); i++ {
		PolygonStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
