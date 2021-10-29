package vulkan

//#include <stdlib.h>
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

func TEXT(str string) *byte {
	return StringToChar(str)
}
func StringToChar(str string) *byte {
	if len(str) != 0 {
		if str[len(str)-1] != byte('\x00') {
			str = str + "\x00"
		}
	} else {
		str = "\x00"
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*byte)(unsafe.Pointer(header.Data))
}
func CharToString(cstr *byte) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}

type VkStringVector struct {
	proto []string
	free  []func()
}

func (this *VkStringVector) Push(strs ...string) {
	for _, str := range strs {
		this.proto = append(this.proto, str)
	}
}
func (this *VkStringVector) PushList(strs []string) {
	for _, str := range strs {
		this.proto = append(this.proto, str)
	}
}
func (this *VkStringVector) Append(strs ...string) {
	for _, str := range strs {
		this.proto = append([]string{str}, this.proto[:]...)
	}
}
func (this *VkStringVector) AppendList(strs []string) {
	this.proto = append(strs, this.proto[:]...)
}
func (this *VkStringVector) Pop() (str string) {
	str = this.proto[len(this.proto)-1]
	this.proto = this.proto[:len(this.proto)-1]
	return
}
func (this *VkStringVector) Shift() (str string) {
	str = this.proto[0]
	this.proto = this.proto[1:]
	return
}
func (this *VkStringVector) Export() (cstrs **byte) {
	if len(this.proto) == 0 {
		return
	}
	var list []string
	for _, str := range this.proto {
		if len(str) != 0 {
			if str[len(str)-1] != byte('\x00') {
				str = str + "\x00"
			}
		} else {
			str = "\x00"
		}
		list = append(list, str)
	}
	n := 0
	for i := range list {
		n += len(list[i])
	}
	data := C.malloc(C.size_t(n))
	// Copy all the strings into data.
	dataSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(data),
		Len:  n,
		Cap:  n,
	}))
	css := make([]*byte, len(list)) // Populated with pointers to each string.
	offset := 0
	for i := range list {
		copy(dataSlice[offset:offset+len(list[i])], list[i][:]) // Copy strs[i] into proper data location.
		css[i] = (*byte)(unsafe.Pointer(&dataSlice[offset]))    // Set a pointer to it.
		offset += len(list[i])
	}
	cstrs = (**byte)(&css[0])
	this.free = append(this.free, func() { C.free(data) })

	return
}
func (this *VkStringVector) Lenght() int {
	return len(this.proto)
}
func (this *VkStringVector) Relese() {
	for _, fun := range this.free {
		fun()
	}
	this.free = make([]func(), 0)
}

func GetError(result VkResult) (err error) {
	switch result {
	case VK_SUCCESS:
	case VK_NOT_READY:
		err = errors.New("VK_NOT_READY")
	case VK_TIMEOUT:
		err = errors.New("VK_TIMEOUT")
	case VK_EVENT_SET:
		err = errors.New("VK_EVENT_SET")
	case VK_EVENT_RESET:
		err = errors.New("VK_EVENT_RESET")
	case VK_INCOMPLETE:
		err = errors.New("VK_INCOMPLETE")
	case VK_ERROR_OUT_OF_HOST_MEMORY:
		err = errors.New("VK_ERROR_OUT_OF_HOST_MEMORY")
	case VK_ERROR_OUT_OF_DEVICE_MEMORY:
		err = errors.New("VK_ERROR_OUT_OF_DEVICE_MEMORY")
	case VK_ERROR_INITIALIZATION_FAILED:
		err = errors.New("VK_ERROR_INITIALIZATION_FAILED")
	case VK_ERROR_DEVICE_LOST:
		err = errors.New("VK_ERROR_DEVICE_LOST")
	case VK_ERROR_MEMORY_MAP_FAILED:
		err = errors.New("VK_ERROR_MEMORY_MAP_FAILED")
	case VK_ERROR_LAYER_NOT_PRESENT:
		err = errors.New("VK_ERROR_LAYER_NOT_PRESENT")
	case VK_ERROR_EXTENSION_NOT_PRESENT:
		err = errors.New("VK_ERROR_EXTENSION_NOT_PRESENT")
	case VK_ERROR_FEATURE_NOT_PRESENT:
		err = errors.New("VK_ERROR_FEATURE_NOT_PRESENT")
	case VK_ERROR_INCOMPATIBLE_DRIVER:
		err = errors.New("VK_ERROR_INCOMPATIBLE_DRIVER")
	case VK_ERROR_TOO_MANY_OBJECTS:
		err = errors.New("VK_ERROR_TOO_MANY_OBJECTS")
	case VK_ERROR_FORMAT_NOT_SUPPORTED:
		err = errors.New("VK_ERROR_FORMAT_NOT_SUPPORTED")
	case VK_ERROR_FRAGMENTED_POOL:
		err = errors.New("VK_ERROR_FRAGMENTED_POOL")
	case VK_ERROR_UNKNOWN:
		err = errors.New("VK_ERROR_UNKNOWN")
	case VK_ERROR_OUT_OF_POOL_MEMORY:
		err = errors.New("VK_ERROR_OUT_OF_POOL_MEMORY")
	case VK_ERROR_INVALID_EXTERNAL_HANDLE:
		err = errors.New("VK_ERROR_INVALID_EXTERNAL_HANDLE")
	case VK_ERROR_FRAGMENTATION:
		err = errors.New("VK_ERROR_FRAGMENTATION")
	case VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS:
		err = errors.New("VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS")
	case VK_ERROR_SURFACE_LOST_KHR:
		err = errors.New("VK_ERROR_SURFACE_LOST_KHR")
	case VK_ERROR_NATIVE_WINDOW_IN_USE_KHR:
		err = errors.New("VK_ERROR_NATIVE_WINDOW_IN_USE_KHR")
	case VK_SUBOPTIMAL_KHR:
		err = errors.New("VK_SUBOPTIMAL_KHR")
	case VK_ERROR_OUT_OF_DATE_KHR:
		err = errors.New("VK_ERROR_OUT_OF_DATE_KHR")
	case VK_ERROR_INCOMPATIBLE_DISPLAY_KHR:
		err = errors.New("VK_ERROR_INCOMPATIBLE_DISPLAY_KHR")
	case VK_ERROR_VALIDATION_FAILED_EXT:
		err = errors.New("VK_ERROR_VALIDATION_FAILED_EXT")
	case VK_ERROR_INVALID_SHADER_NV:
		err = errors.New("VK_ERROR_INVALID_SHADER_NV")
	case VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT:
		err = errors.New("VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT")
	case VK_ERROR_NOT_PERMITTED_EXT:
		err = errors.New("VK_ERROR_NOT_PERMITTED_EXT")
	case VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT:
		err = errors.New("VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT")
	case VK_THREAD_IDLE_KHR:
		err = errors.New("VK_THREAD_IDLE_KHR")
	case VK_THREAD_DONE_KHR:
		err = errors.New("VK_THREAD_DONE_KHR")
	case VK_OPERATION_DEFERRED_KHR:
		err = errors.New("VK_OPERATION_DEFERRED_KHR")
	case VK_OPERATION_NOT_DEFERRED_KHR:
		err = errors.New("VK_OPERATION_NOT_DEFERRED_KHR")
	case VK_PIPELINE_COMPILE_REQUIRED_EXT:
		err = errors.New("VK_PIPELINE_COMPILE_REQUIRED_EXT")
	case VK_RESULT_MAX_ENUM:
		err = errors.New("VK_RESULT_MAX_ENUM")
	}
	return
}
