// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/pose_with_covariance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PoseWithCovariance", PoseWithCovarianceTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/PoseWithCovariance", PoseWithCovarianceTypeSupport)
}

type PoseWithCovariance struct {
	Pose Pose `yaml:"pose"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewPoseWithCovariance creates a new PoseWithCovariance with default values.
func NewPoseWithCovariance() *PoseWithCovariance {
	self := PoseWithCovariance{}
	self.SetDefaults()
	return &self
}

func (t *PoseWithCovariance) Clone() *PoseWithCovariance {
	c := &PoseWithCovariance{}
	c.Pose = *t.Pose.Clone()
	c.Covariance = t.Covariance
	return c
}

func (t *PoseWithCovariance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PoseWithCovariance) SetDefaults() {
	t.Pose.SetDefaults()
	t.Covariance = [36]float64{}
}

func (t *PoseWithCovariance) GetTypeSupport() types.MessageTypeSupport {
	return PoseWithCovarianceTypeSupport
}

// PoseWithCovariancePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PoseWithCovariancePublisher struct {
	*rclgo.Publisher
}

// NewPoseWithCovariancePublisher creates and returns a new publisher for the
// PoseWithCovariance
func NewPoseWithCovariancePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PoseWithCovariancePublisher, error) {
	pub, err := node.NewPublisher(topic_name, PoseWithCovarianceTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PoseWithCovariancePublisher{pub}, nil
}

func (p *PoseWithCovariancePublisher) Publish(msg *PoseWithCovariance) error {
	return p.Publisher.Publish(msg)
}

// PoseWithCovarianceSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PoseWithCovarianceSubscription struct {
	*rclgo.Subscription
}

// PoseWithCovarianceSubscriptionCallback type is used to provide a subscription
// handler function for a PoseWithCovarianceSubscription.
type PoseWithCovarianceSubscriptionCallback func(msg *PoseWithCovariance, info *rclgo.MessageInfo, err error)

// NewPoseWithCovarianceSubscription creates and returns a new subscription for the
// PoseWithCovariance
func NewPoseWithCovarianceSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PoseWithCovarianceSubscriptionCallback) (*PoseWithCovarianceSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PoseWithCovariance
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PoseWithCovarianceTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PoseWithCovarianceSubscription{sub}, nil
}

func (s *PoseWithCovarianceSubscription) TakeMessage(out *PoseWithCovariance) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePoseWithCovarianceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseWithCovarianceSlice(dst, src []PoseWithCovariance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseWithCovarianceTypeSupport types.MessageTypeSupport = _PoseWithCovarianceTypeSupport{}

type _PoseWithCovarianceTypeSupport struct{}

func (t _PoseWithCovarianceTypeSupport) New() types.Message {
	return NewPoseWithCovariance()
}

func (t _PoseWithCovarianceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseWithCovariance__create())
}

func (t _PoseWithCovarianceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseWithCovariance__destroy((*C.geometry_msgs__msg__PoseWithCovariance)(pointer_to_free))
}

func (t _PoseWithCovarianceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PoseWithCovariance)
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(dst)
	PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)), m.Covariance[:])
}

func (t _PoseWithCovarianceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PoseWithCovariance)
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(ros2_message_buffer)
	PoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_Go(m.Covariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}

func (t _PoseWithCovarianceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseWithCovariance())
}

type CPoseWithCovariance = C.geometry_msgs__msg__PoseWithCovariance
type CPoseWithCovariance__Sequence = C.geometry_msgs__msg__PoseWithCovariance__Sequence

func PoseWithCovariance__Sequence_to_Go(goSlice *[]PoseWithCovariance, cSlice CPoseWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseWithCovariance, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PoseWithCovarianceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PoseWithCovariance__Sequence_to_C(cSlice *CPoseWithCovariance__Sequence, goSlice []PoseWithCovariance) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseWithCovariance)(C.malloc(C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PoseWithCovariance__Array_to_Go(goSlice []PoseWithCovariance, cSlice []CPoseWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		PoseWithCovarianceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PoseWithCovariance__Array_to_C(cSlice []CPoseWithCovariance, goSlice []PoseWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
