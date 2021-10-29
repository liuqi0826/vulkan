package vulkan

import (
	"syscall"
	"unsafe"
)

type (
	HWND      = uintptr
	HANDLE    = uintptr
	HINSTANCE = uintptr
	HMONITOR  = uintptr
	BOOL      = int32
	WORD      = uint16
	DWORD     = uint32
	WCHAR     = uint16
	LPWSTR    = *WCHAR
	LPCWSTR   = *WCHAR
	TCHAR     = uint16
	LPTSTR    = *TCHAR
	LPCTSTR   = *TCHAR
	LPVOID    = unsafe.Pointer
)
type SECURITY_ATTRIBUTES struct {
	NLength              DWORD
	LPSecurityDescriptor LPVOID
	BInheritHandle       BOOL
}

var (
	vulkan                                                                = syscall.NewLazyDLL("vulkan-1.dll")
	PFN_vkCreateInstance                                                  = vulkan.NewProc("vkCreateInstance")
	PFN_vkDestroyInstance                                                 = vulkan.NewProc("vkDestroyInstance")
	PFN_vkEnumeratePhysicalDevices                                        = vulkan.NewProc("vkEnumeratePhysicalDevices")
	PFN_vkGetPhysicalDeviceFeatures                                       = vulkan.NewProc("vkGetPhysicalDeviceFeatures")
	PFN_vkGetPhysicalDeviceFormatProperties                               = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties")
	PFN_vkGetPhysicalDeviceImageFormatProperties                          = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties")
	PFN_vkGetPhysicalDeviceProperties                                     = vulkan.NewProc("vkGetPhysicalDeviceProperties")
	PFN_vkGetPhysicalDeviceQueueFamilyProperties                          = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties")
	PFN_vkGetPhysicalDeviceMemoryProperties                               = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties")
	PFN_vkGetInstanceProcAddr                                             = vulkan.NewProc("vkGetInstanceProcAddr")
	PFN_vkGetDeviceProcAddr                                               = vulkan.NewProc("vkGetDeviceProcAddr")
	PFN_vkCreateDevice                                                    = vulkan.NewProc("vkCreateDevice")
	PFN_vkDestroyDevice                                                   = vulkan.NewProc("vkDestroyDevice")
	PFN_vkEnumerateInstanceExtensionProperties                            = vulkan.NewProc("vkEnumerateInstanceExtensionProperties")
	PFN_vkEnumerateDeviceExtensionProperties                              = vulkan.NewProc("vkEnumerateDeviceExtensionProperties")
	PFN_vkEnumerateInstanceLayerProperties                                = vulkan.NewProc("vkEnumerateInstanceLayerProperties")
	PFN_vkEnumerateDeviceLayerProperties                                  = vulkan.NewProc("vkEnumerateDeviceLayerProperties")
	PFN_vkGetDeviceQueue                                                  = vulkan.NewProc("vkGetDeviceQueue")
	PFN_vkQueueSubmit                                                     = vulkan.NewProc("vkQueueSubmit")
	PFN_vkQueueWaitIdle                                                   = vulkan.NewProc("vkQueueWaitIdle")
	PFN_vkDeviceWaitIdle                                                  = vulkan.NewProc("vkDeviceWaitIdle")
	PFN_vkAllocateMemory                                                  = vulkan.NewProc("vkAllocateMemory")
	PFN_vkFreeMemory                                                      = vulkan.NewProc("vkFreeMemory")
	PFN_vkMapMemory                                                       = vulkan.NewProc("vkMapMemory")
	PFN_vkUnmapMemory                                                     = vulkan.NewProc("vkUnmapMemory")
	PFN_vkFlushMappedMemoryRanges                                         = vulkan.NewProc("vkFlushMappedMemoryRanges")
	PFN_vkInvalidateMappedMemoryRanges                                    = vulkan.NewProc("vkInvalidateMappedMemoryRanges")
	PFN_vkGetDeviceMemoryCommitment                                       = vulkan.NewProc("vkGetDeviceMemoryCommitment")
	PFN_vkBindBufferMemory                                                = vulkan.NewProc("vkBindBufferMemory")
	PFN_vkBindImageMemory                                                 = vulkan.NewProc("vkBindImageMemory")
	PFN_vkGetBufferMemoryRequirements                                     = vulkan.NewProc("vkGetBufferMemoryRequirements")
	PFN_vkGetImageMemoryRequirements                                      = vulkan.NewProc("vkGetImageMemoryRequirements")
	PFN_vkGetImageSparseMemoryRequirements                                = vulkan.NewProc("vkGetImageSparseMemoryRequirements")
	PFN_vkGetPhysicalDeviceSparseImageFormatProperties                    = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties")
	PFN_vkQueueBindSparse                                                 = vulkan.NewProc("vkQueueBindSparse")
	PFN_vkCreateFence                                                     = vulkan.NewProc("vkCreateFence")
	PFN_vkDestroyFence                                                    = vulkan.NewProc("vkDestroyFence")
	PFN_vkResetFences                                                     = vulkan.NewProc("vkResetFences")
	PFN_vkGetFenceStatus                                                  = vulkan.NewProc("vkGetFenceStatus")
	PFN_vkWaitForFences                                                   = vulkan.NewProc("vkWaitForFences")
	PFN_vkCreateSemaphore                                                 = vulkan.NewProc("vkCreateSemaphore")
	PFN_vkDestroySemaphore                                                = vulkan.NewProc("vkDestroySemaphore")
	PFN_vkCreateEvent                                                     = vulkan.NewProc("vkCreateEvent")
	PFN_vkDestroyEvent                                                    = vulkan.NewProc("vkDestroyEvent")
	PFN_vkGetEventStatus                                                  = vulkan.NewProc("vkGetEventStatus")
	PFN_vkSetEvent                                                        = vulkan.NewProc("vkSetEvent")
	PFN_vkResetEvent                                                      = vulkan.NewProc("vkResetEvent")
	PFN_vkCreateQueryPool                                                 = vulkan.NewProc("vkCreateQueryPool")
	PFN_vkDestroyQueryPool                                                = vulkan.NewProc("vkDestroyQueryPool")
	PFN_vkGetQueryPoolResults                                             = vulkan.NewProc("vkGetQueryPoolResults")
	PFN_vkCreateBuffer                                                    = vulkan.NewProc("vkCreateBuffer")
	PFN_vkDestroyBuffer                                                   = vulkan.NewProc("vkDestroyBuffer")
	PFN_vkCreateBufferView                                                = vulkan.NewProc("vkCreateBufferView")
	PFN_vkDestroyBufferView                                               = vulkan.NewProc("vkDestroyBufferView")
	PFN_vkCreateImage                                                     = vulkan.NewProc("vkCreateImage")
	PFN_vkDestroyImage                                                    = vulkan.NewProc("vkDestroyImage")
	PFN_vkGetImageSubresourceLayout                                       = vulkan.NewProc("vkGetImageSubresourceLayout")
	PFN_vkCreateImageView                                                 = vulkan.NewProc("vkCreateImageView")
	PFN_vkDestroyImageView                                                = vulkan.NewProc("vkDestroyImageView")
	PFN_vkCreateShaderModule                                              = vulkan.NewProc("vkCreateShaderModule")
	PFN_vkDestroyShaderModule                                             = vulkan.NewProc("vkDestroyShaderModule")
	PFN_vkCreatePipelineCache                                             = vulkan.NewProc("vkCreatePipelineCache")
	PFN_vkDestroyPipelineCache                                            = vulkan.NewProc("vkDestroyPipelineCache")
	PFN_vkGetPipelineCacheData                                            = vulkan.NewProc("vkGetPipelineCacheData")
	PFN_vkMergePipelineCaches                                             = vulkan.NewProc("vkMergePipelineCaches")
	PFN_vkCreateGraphicsPipelines                                         = vulkan.NewProc("vkCreateGraphicsPipelines")
	PFN_vkCreateComputePipelines                                          = vulkan.NewProc("vkCreateComputePipelines")
	PFN_vkDestroyPipeline                                                 = vulkan.NewProc("vkDestroyPipeline")
	PFN_vkCreatePipelineLayout                                            = vulkan.NewProc("vkCreatePipelineLayout")
	PFN_vkDestroyPipelineLayout                                           = vulkan.NewProc("vkDestroyPipelineLayout")
	PFN_vkCreateSampler                                                   = vulkan.NewProc("vkCreateSampler")
	PFN_vkDestroySampler                                                  = vulkan.NewProc("vkDestroySampler")
	PFN_vkCreateDescriptorSetLayout                                       = vulkan.NewProc("vkCreateDescriptorSetLayout")
	PFN_vkDestroyDescriptorSetLayout                                      = vulkan.NewProc("vkDestroyDescriptorSetLayout")
	PFN_vkCreateDescriptorPool                                            = vulkan.NewProc("vkCreateDescriptorPool")
	PFN_vkDestroyDescriptorPool                                           = vulkan.NewProc("vkDestroyDescriptorPool")
	PFN_vkResetDescriptorPool                                             = vulkan.NewProc("vkResetDescriptorPool")
	PFN_vkAllocateDescriptorSets                                          = vulkan.NewProc("vkAllocateDescriptorSets")
	PFN_vkFreeDescriptorSets                                              = vulkan.NewProc("vkFreeDescriptorSets")
	PFN_vkUpdateDescriptorSets                                            = vulkan.NewProc("vkUpdateDescriptorSets")
	PFN_vkCreateFramebuffer                                               = vulkan.NewProc("vkCreateFramebuffer")
	PFN_vkDestroyFramebuffer                                              = vulkan.NewProc("vkDestroyFramebuffer")
	PFN_vkCreateRenderPass                                                = vulkan.NewProc("vkCreateRenderPass")
	PFN_vkDestroyRenderPass                                               = vulkan.NewProc("vkDestroyRenderPass")
	PFN_vkGetRenderAreaGranularity                                        = vulkan.NewProc("vkGetRenderAreaGranularity")
	PFN_vkCreateCommandPool                                               = vulkan.NewProc("vkCreateCommandPool")
	PFN_vkDestroyCommandPool                                              = vulkan.NewProc("vkDestroyCommandPool")
	PFN_vkResetCommandPool                                                = vulkan.NewProc("vkResetCommandPool")
	PFN_vkAllocateCommandBuffers                                          = vulkan.NewProc("vkAllocateCommandBuffers")
	PFN_vkFreeCommandBuffers                                              = vulkan.NewProc("vkFreeCommandBuffers")
	PFN_vkBeginCommandBuffer                                              = vulkan.NewProc("vkBeginCommandBuffer")
	PFN_vkEndCommandBuffer                                                = vulkan.NewProc("vkEndCommandBuffer")
	PFN_vkResetCommandBuffer                                              = vulkan.NewProc("vkResetCommandBuffer")
	PFN_vkCmdBindPipeline                                                 = vulkan.NewProc("vkCmdBindPipeline")
	PFN_vkCmdSetViewport                                                  = vulkan.NewProc("vkCmdSetViewport")
	PFN_vkCmdSetScissor                                                   = vulkan.NewProc("vkCmdSetScissor")
	PFN_vkCmdSetLineWidth                                                 = vulkan.NewProc("vkCmdSetLineWidth")
	PFN_vkCmdSetDepthBias                                                 = vulkan.NewProc("vkCmdSetDepthBias")
	PFN_vkCmdSetBlendConstants                                            = vulkan.NewProc("vkCmdSetBlendConstants")
	PFN_vkCmdSetDepthBounds                                               = vulkan.NewProc("vkCmdSetDepthBounds")
	PFN_vkCmdSetStencilCompareMask                                        = vulkan.NewProc("vkCmdSetStencilCompareMask")
	PFN_vkCmdSetStencilWriteMask                                          = vulkan.NewProc("vkCmdSetStencilWriteMask")
	PFN_vkCmdSetStencilReference                                          = vulkan.NewProc("vkCmdSetStencilReference")
	PFN_vkCmdBindDescriptorSets                                           = vulkan.NewProc("vkCmdBindDescriptorSets")
	PFN_vkCmdBindIndexBuffer                                              = vulkan.NewProc("vkCmdBindIndexBuffer")
	PFN_vkCmdBindVertexBuffers                                            = vulkan.NewProc("vkCmdBindVertexBuffers")
	PFN_vkCmdDraw                                                         = vulkan.NewProc("vkCmdDraw")
	PFN_vkCmdDrawIndexed                                                  = vulkan.NewProc("vkCmdDrawIndexed")
	PFN_vkCmdDrawIndirect                                                 = vulkan.NewProc("vkCmdDrawIndirect")
	PFN_vkCmdDrawIndexedIndirect                                          = vulkan.NewProc("vkCmdDrawIndexedIndirect")
	PFN_vkCmdDispatch                                                     = vulkan.NewProc("vkCmdDispatch")
	PFN_vkCmdDispatchIndirect                                             = vulkan.NewProc("vkCmdDispatchIndirect")
	PFN_vkCmdCopyBuffer                                                   = vulkan.NewProc("vkCmdCopyBuffer")
	PFN_vkCmdCopyImage                                                    = vulkan.NewProc("vkCmdCopyImage")
	PFN_vkCmdBlitImage                                                    = vulkan.NewProc("vkCmdBlitImage")
	PFN_vkCmdCopyBufferToImage                                            = vulkan.NewProc("vkCmdCopyBufferToImage")
	PFN_vkCmdCopyImageToBuffer                                            = vulkan.NewProc("vkCmdCopyImageToBuffer")
	PFN_vkCmdUpdateBuffer                                                 = vulkan.NewProc("vkCmdUpdateBuffer")
	PFN_vkCmdFillBuffer                                                   = vulkan.NewProc("vkCmdFillBuffer")
	PFN_vkCmdClearColorImage                                              = vulkan.NewProc("vkCmdClearColorImage")
	PFN_vkCmdClearDepthStencilImage                                       = vulkan.NewProc("vkCmdClearDepthStencilImage")
	PFN_vkCmdClearAttachments                                             = vulkan.NewProc("vkCmdClearAttachments")
	PFN_vkCmdResolveImage                                                 = vulkan.NewProc("vkCmdResolveImage")
	PFN_vkCmdSetEvent                                                     = vulkan.NewProc("vkCmdSetEvent")
	PFN_vkCmdResetEvent                                                   = vulkan.NewProc("vkCmdResetEvent")
	PFN_vkCmdWaitEvents                                                   = vulkan.NewProc("vkCmdWaitEvents")
	PFN_vkCmdPipelineBarrier                                              = vulkan.NewProc("vkCmdPipelineBarrier")
	PFN_vkCmdBeginQuery                                                   = vulkan.NewProc("vkCmdBeginQuery")
	PFN_vkCmdEndQuery                                                     = vulkan.NewProc("vkCmdEndQuery")
	PFN_vkCmdResetQueryPool                                               = vulkan.NewProc("vkCmdResetQueryPool")
	PFN_vkCmdWriteTimestamp                                               = vulkan.NewProc("vkCmdWriteTimestamp")
	PFN_vkCmdCopyQueryPoolResults                                         = vulkan.NewProc("vkCmdCopyQueryPoolResults")
	PFN_vkCmdPushConstants                                                = vulkan.NewProc("vkCmdPushConstants")
	PFN_vkCmdBeginRenderPass                                              = vulkan.NewProc("vkCmdBeginRenderPass")
	PFN_vkCmdNextSubpass                                                  = vulkan.NewProc("vkCmdNextSubpass")
	PFN_vkCmdEndRenderPass                                                = vulkan.NewProc("vkCmdEndRenderPass")
	PFN_vkCmdExecuteCommands                                              = vulkan.NewProc("vkCmdExecuteCommands")
	PFN_vkEnumerateInstanceVersion                                        = vulkan.NewProc("vkEnumerateInstanceVersion")
	PFN_vkBindBufferMemory2                                               = vulkan.NewProc("vkBindBufferMemory2")
	PFN_vkBindImageMemory2                                                = vulkan.NewProc("vkBindImageMemory2")
	PFN_vkGetDeviceGroupPeerMemoryFeatures                                = vulkan.NewProc("vkGetDeviceGroupPeerMemoryFeatures")
	PFN_vkCmdSetDeviceMask                                                = vulkan.NewProc("vkCmdSetDeviceMask")
	PFN_vkCmdDispatchBase                                                 = vulkan.NewProc("vkCmdDispatchBase")
	PFN_vkEnumeratePhysicalDeviceGroups                                   = vulkan.NewProc("vkEnumeratePhysicalDeviceGroups")
	PFN_vkGetImageMemoryRequirements2                                     = vulkan.NewProc("vkGetImageMemoryRequirements2")
	PFN_vkGetBufferMemoryRequirements2                                    = vulkan.NewProc("vkGetBufferMemoryRequirements2")
	PFN_vkGetImageSparseMemoryRequirements2                               = vulkan.NewProc("vkGetImageSparseMemoryRequirements2")
	PFN_vkGetPhysicalDeviceFeatures2                                      = vulkan.NewProc("vkGetPhysicalDeviceFeatures2")
	PFN_vkGetPhysicalDeviceProperties2                                    = vulkan.NewProc("vkGetPhysicalDeviceProperties2")
	PFN_vkGetPhysicalDeviceFormatProperties2                              = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties2")
	PFN_vkGetPhysicalDeviceImageFormatProperties2                         = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties2")
	PFN_vkGetPhysicalDeviceQueueFamilyProperties2                         = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties2")
	PFN_vkGetPhysicalDeviceMemoryProperties2                              = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties2")
	PFN_vkGetPhysicalDeviceSparseImageFormatProperties2                   = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties2")
	PFN_vkTrimCommandPool                                                 = vulkan.NewProc("vkTrimCommandPool")
	PFN_vkGetDeviceQueue2                                                 = vulkan.NewProc("vkGetDeviceQueue2")
	PFN_vkCreateSamplerYcbcrConversion                                    = vulkan.NewProc("vkCreateSamplerYcbcrConversion")
	PFN_vkDestroySamplerYcbcrConversion                                   = vulkan.NewProc("vkDestroySamplerYcbcrConversion")
	PFN_vkCreateDescriptorUpdateTemplate                                  = vulkan.NewProc("vkCreateDescriptorUpdateTemplate")
	PFN_vkDestroyDescriptorUpdateTemplate                                 = vulkan.NewProc("vkDestroyDescriptorUpdateTemplate")
	PFN_vkUpdateDescriptorSetWithTemplate                                 = vulkan.NewProc("vkUpdateDescriptorSetWithTemplate")
	PFN_vkGetPhysicalDeviceExternalBufferProperties                       = vulkan.NewProc("vkGetPhysicalDeviceExternalBufferProperties")
	PFN_vkGetPhysicalDeviceExternalFenceProperties                        = vulkan.NewProc("vkGetPhysicalDeviceExternalFenceProperties")
	PFN_vkGetPhysicalDeviceExternalSemaphoreProperties                    = vulkan.NewProc("vkGetPhysicalDeviceExternalSemaphoreProperties")
	PFN_vkGetDescriptorSetLayoutSupport                                   = vulkan.NewProc("vkGetDescriptorSetLayoutSupport")
	PFN_vkCmdDrawIndirectCount                                            = vulkan.NewProc("vkCmdDrawIndirectCount")
	PFN_vkCmdDrawIndexedIndirectCount                                     = vulkan.NewProc("vkCmdDrawIndexedIndirectCount")
	PFN_vkCreateRenderPass2                                               = vulkan.NewProc("vkCreateRenderPass2")
	PFN_vkCmdBeginRenderPass2                                             = vulkan.NewProc("vkCmdBeginRenderPass2")
	PFN_vkCmdNextSubpass2                                                 = vulkan.NewProc("vkCmdNextSubpass2")
	PFN_vkCmdEndRenderPass2                                               = vulkan.NewProc("vkCmdEndRenderPass2")
	PFN_vkResetQueryPool                                                  = vulkan.NewProc("vkResetQueryPool")
	PFN_vkGetSemaphoreCounterValue                                        = vulkan.NewProc("vkGetSemaphoreCounterValue")
	PFN_vkWaitSemaphores                                                  = vulkan.NewProc("vkWaitSemaphores")
	PFN_vkSignalSemaphore                                                 = vulkan.NewProc("vkSignalSemaphore")
	PFN_vkGetBufferDeviceAddress                                          = vulkan.NewProc("vkGetBufferDeviceAddress")
	PFN_vkGetBufferOpaqueCaptureAddress                                   = vulkan.NewProc("vkGetBufferOpaqueCaptureAddress")
	PFN_vkGetDeviceMemoryOpaqueCaptureAddress                             = vulkan.NewProc("vkGetDeviceMemoryOpaqueCaptureAddress")
	PFN_vkDestroySurfaceKHR                                               = vulkan.NewProc("vkDestroySurfaceKHR")
	PFN_vkGetPhysicalDeviceSurfaceSupportKHR                              = vulkan.NewProc("vkGetPhysicalDeviceSurfaceSupportKHR")
	PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR                         = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilitiesKHR")
	PFN_vkGetPhysicalDeviceSurfaceFormatsKHR                              = vulkan.NewProc("vkGetPhysicalDeviceSurfaceFormatsKHR")
	PFN_vkGetPhysicalDeviceSurfacePresentModesKHR                         = vulkan.NewProc("vkGetPhysicalDeviceSurfacePresentModesKHR")
	PFN_vkCreateSwapchainKHR                                              = vulkan.NewProc("vkCreateSwapchainKHR")
	PFN_vkDestroySwapchainKHR                                             = vulkan.NewProc("vkDestroySwapchainKHR")
	PFN_vkGetSwapchainImagesKHR                                           = vulkan.NewProc("vkGetSwapchainImagesKHR")
	PFN_vkAcquireNextImageKHR                                             = vulkan.NewProc("vkAcquireNextImageKHR")
	PFN_vkQueuePresentKHR                                                 = vulkan.NewProc("vkQueuePresentKHR")
	PFN_vkGetDeviceGroupPresentCapabilitiesKHR                            = vulkan.NewProc("vkGetDeviceGroupPresentCapabilitiesKHR")
	PFN_vkGetDeviceGroupSurfacePresentModesKHR                            = vulkan.NewProc("vkGetDeviceGroupSurfacePresentModesKHR")
	PFN_vkGetPhysicalDevicePresentRectanglesKHR                           = vulkan.NewProc("vkGetPhysicalDevicePresentRectanglesKHR")
	PFN_vkAcquireNextImage2KHR                                            = vulkan.NewProc("vkAcquireNextImage2KHR")
	PFN_vkGetPhysicalDeviceDisplayPropertiesKHR                           = vulkan.NewProc("vkGetPhysicalDeviceDisplayPropertiesKHR")
	PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR                      = vulkan.NewProc("vkGetPhysicalDeviceDisplayPlanePropertiesKHR")
	PFN_vkGetDisplayPlaneSupportedDisplaysKHR                             = vulkan.NewProc("vkGetDisplayPlaneSupportedDisplaysKHR")
	PFN_vkGetDisplayModePropertiesKHR                                     = vulkan.NewProc("vkGetDisplayModePropertiesKHR")
	PFN_vkCreateDisplayModeKHR                                            = vulkan.NewProc("vkCreateDisplayModeKHR")
	PFN_vkGetDisplayPlaneCapabilitiesKHR                                  = vulkan.NewProc("vkGetDisplayPlaneCapabilitiesKHR")
	PFN_vkCreateDisplayPlaneSurfaceKHR                                    = vulkan.NewProc("vkCreateDisplayPlaneSurfaceKHR")
	PFN_vkCreateSharedSwapchainsKHR                                       = vulkan.NewProc("vkCreateSharedSwapchainsKHR")
	PFN_vkGetPhysicalDeviceFeatures2KHR                                   = vulkan.NewProc("vkGetPhysicalDeviceFeatures2KHR")
	PFN_vkGetPhysicalDeviceProperties2KHR                                 = vulkan.NewProc("vkGetPhysicalDeviceProperties2KHR")
	PFN_vkGetPhysicalDeviceFormatProperties2KHR                           = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties2KHR")
	PFN_vkGetPhysicalDeviceImageFormatProperties2KHR                      = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties2KHR")
	PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR                      = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties2KHR")
	PFN_vkGetPhysicalDeviceMemoryProperties2KHR                           = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties2KHR")
	PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR                = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties2KHR")
	PFN_vkGetDeviceGroupPeerMemoryFeaturesKHR                             = vulkan.NewProc("vkGetDeviceGroupPeerMemoryFeaturesKHR")
	PFN_vkCmdSetDeviceMaskKHR                                             = vulkan.NewProc("vkCmdSetDeviceMaskKHR")
	PFN_vkCmdDispatchBaseKHR                                              = vulkan.NewProc("vkCmdDispatchBaseKHR")
	PFN_vkTrimCommandPoolKHR                                              = vulkan.NewProc("vkTrimCommandPoolKHR")
	PFN_vkEnumeratePhysicalDeviceGroupsKHR                                = vulkan.NewProc("vkEnumeratePhysicalDeviceGroupsKHR")
	PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR                    = vulkan.NewProc("vkGetPhysicalDeviceExternalBufferPropertiesKHR")
	PFN_vkGetMemoryFdKHR                                                  = vulkan.NewProc("vkGetMemoryFdKHR")
	PFN_vkGetMemoryFdPropertiesKHR                                        = vulkan.NewProc("vkGetMemoryFdPropertiesKHR")
	PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR                 = vulkan.NewProc("vkGetPhysicalDeviceExternalSemaphorePropertiesKHR")
	PFN_vkImportSemaphoreFdKHR                                            = vulkan.NewProc("vkImportSemaphoreFdKHR")
	PFN_vkGetSemaphoreFdKHR                                               = vulkan.NewProc("vkGetSemaphoreFdKHR")
	PFN_vkCmdPushDescriptorSetKHR                                         = vulkan.NewProc("vkCmdPushDescriptorSetKHR")
	PFN_vkCmdPushDescriptorSetWithTemplateKHR                             = vulkan.NewProc("vkCmdPushDescriptorSetWithTemplateKHR")
	PFN_vkCreateDescriptorUpdateTemplateKHR                               = vulkan.NewProc("vkCreateDescriptorUpdateTemplateKHR")
	PFN_vkDestroyDescriptorUpdateTemplateKHR                              = vulkan.NewProc("vkDestroyDescriptorUpdateTemplateKHR")
	PFN_vkUpdateDescriptorSetWithTemplateKHR                              = vulkan.NewProc("vkUpdateDescriptorSetWithTemplateKHR")
	PFN_vkCreateRenderPass2KHR                                            = vulkan.NewProc("vkCreateRenderPass2KHR")
	PFN_vkCmdBeginRenderPass2KHR                                          = vulkan.NewProc("vkCmdBeginRenderPass2KHR")
	PFN_vkCmdNextSubpass2KHR                                              = vulkan.NewProc("vkCmdNextSubpass2KHR")
	PFN_vkCmdEndRenderPass2KHR                                            = vulkan.NewProc("vkCmdEndRenderPass2KHR")
	PFN_vkGetSwapchainStatusKHR                                           = vulkan.NewProc("vkGetSwapchainStatusKHR")
	PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR                     = vulkan.NewProc("vkGetPhysicalDeviceExternalFencePropertiesKHR")
	PFN_vkImportFenceFdKHR                                                = vulkan.NewProc("vkImportFenceFdKHR")
	PFN_vkGetFenceFdKHR                                                   = vulkan.NewProc("vkGetFenceFdKHR")
	PFN_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR   = vulkan.NewProc("vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR")
	PFN_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR           = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR")
	PFN_vkAcquireProfilingLockKHR                                         = vulkan.NewProc("vkAcquireProfilingLockKHR")
	PFN_vkReleaseProfilingLockKHR                                         = vulkan.NewProc("vkReleaseProfilingLockKHR")
	PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR                        = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilities2KHR")
	PFN_vkGetPhysicalDeviceSurfaceFormats2KHR                             = vulkan.NewProc("vkGetPhysicalDeviceSurfaceFormats2KHR")
	PFN_vkGetPhysicalDeviceDisplayProperties2KHR                          = vulkan.NewProc("vkGetPhysicalDeviceDisplayProperties2KHR")
	PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR                     = vulkan.NewProc("vkGetPhysicalDeviceDisplayPlaneProperties2KHR")
	PFN_vkGetDisplayModeProperties2KHR                                    = vulkan.NewProc("vkGetDisplayModeProperties2KHR")
	PFN_vkGetDisplayPlaneCapabilities2KHR                                 = vulkan.NewProc("vkGetDisplayPlaneCapabilities2KHR")
	PFN_vkGetImageMemoryRequirements2KHR                                  = vulkan.NewProc("vkGetImageMemoryRequirements2KHR")
	PFN_vkGetBufferMemoryRequirements2KHR                                 = vulkan.NewProc("vkGetBufferMemoryRequirements2KHR")
	PFN_vkGetImageSparseMemoryRequirements2KHR                            = vulkan.NewProc("vkGetImageSparseMemoryRequirements2KHR")
	PFN_vkCreateSamplerYcbcrConversionKHR                                 = vulkan.NewProc("vkCreateSamplerYcbcrConversionKHR")
	PFN_vkDestroySamplerYcbcrConversionKHR                                = vulkan.NewProc("vkDestroySamplerYcbcrConversionKHR")
	PFN_vkBindBufferMemory2KHR                                            = vulkan.NewProc("vkBindBufferMemory2KHR")
	PFN_vkBindImageMemory2KHR                                             = vulkan.NewProc("vkBindImageMemory2KHR")
	PFN_vkGetDescriptorSetLayoutSupportKHR                                = vulkan.NewProc("vkGetDescriptorSetLayoutSupportKHR")
	PFN_vkCmdDrawIndirectCountKHR                                         = vulkan.NewProc("vkCmdDrawIndirectCountKHR")
	PFN_vkCmdDrawIndexedIndirectCountKHR                                  = vulkan.NewProc("vkCmdDrawIndexedIndirectCountKHR")
	PFN_vkGetSemaphoreCounterValueKHR                                     = vulkan.NewProc("vkGetSemaphoreCounterValueKHR")
	PFN_vkWaitSemaphoresKHR                                               = vulkan.NewProc("vkWaitSemaphoresKHR")
	PFN_vkSignalSemaphoreKHR                                              = vulkan.NewProc("vkSignalSemaphoreKHR")
	PFN_vkGetPhysicalDeviceFragmentShadingRatesKHR                        = vulkan.NewProc("vkGetPhysicalDeviceFragmentShadingRatesKHR")
	PFN_vkCmdSetFragmentShadingRateKHR                                    = vulkan.NewProc("vkCmdSetFragmentShadingRateKHR")
	PFN_vkWaitForPresentKHR                                               = vulkan.NewProc("vkWaitForPresentKHR")
	PFN_vkGetBufferDeviceAddressKHR                                       = vulkan.NewProc("vkGetBufferDeviceAddressKHR")
	PFN_vkGetBufferOpaqueCaptureAddressKHR                                = vulkan.NewProc("vkGetBufferOpaqueCaptureAddressKHR")
	PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR                          = vulkan.NewProc("vkGetDeviceMemoryOpaqueCaptureAddressKHR")
	PFN_vkCreateDeferredOperationKHR                                      = vulkan.NewProc("vkCreateDeferredOperationKHR")
	PFN_vkDestroyDeferredOperationKHR                                     = vulkan.NewProc("vkDestroyDeferredOperationKHR")
	PFN_vkGetDeferredOperationMaxConcurrencyKHR                           = vulkan.NewProc("vkGetDeferredOperationMaxConcurrencyKHR")
	PFN_vkGetDeferredOperationResultKHR                                   = vulkan.NewProc("vkGetDeferredOperationResultKHR")
	PFN_vkDeferredOperationJoinKHR                                        = vulkan.NewProc("vkDeferredOperationJoinKHR")
	PFN_vkGetPipelineExecutablePropertiesKHR                              = vulkan.NewProc("vkGetPipelineExecutablePropertiesKHR")
	PFN_vkGetPipelineExecutableStatisticsKHR                              = vulkan.NewProc("vkGetPipelineExecutableStatisticsKHR")
	PFN_vkGetPipelineExecutableInternalRepresentationsKHR                 = vulkan.NewProc("vkGetPipelineExecutableInternalRepresentationsKHR")
	PFN_vkCmdSetEvent2KHR                                                 = vulkan.NewProc("vkCmdSetEvent2KHR")
	PFN_vkCmdResetEvent2KHR                                               = vulkan.NewProc("vkCmdResetEvent2KHR")
	PFN_vkCmdWaitEvents2KHR                                               = vulkan.NewProc("vkCmdWaitEvents2KHR")
	PFN_vkCmdPipelineBarrier2KHR                                          = vulkan.NewProc("vkCmdPipelineBarrier2KHR")
	PFN_vkCmdWriteTimestamp2KHR                                           = vulkan.NewProc("vkCmdWriteTimestamp2KHR")
	PFN_vkQueueSubmit2KHR                                                 = vulkan.NewProc("vkQueueSubmit2KHR")
	PFN_vkCmdWriteBufferMarker2AMD                                        = vulkan.NewProc("vkCmdWriteBufferMarker2AMD")
	PFN_vkGetQueueCheckpointData2NV                                       = vulkan.NewProc("vkGetQueueCheckpointData2NV")
	PFN_vkCmdCopyBuffer2KHR                                               = vulkan.NewProc("vkCmdCopyBuffer2KHR")
	PFN_vkCmdCopyImage2KHR                                                = vulkan.NewProc("vkCmdCopyImage2KHR")
	PFN_vkCmdCopyBufferToImage2KHR                                        = vulkan.NewProc("vkCmdCopyBufferToImage2KHR")
	PFN_vkCmdCopyImageToBuffer2KHR                                        = vulkan.NewProc("vkCmdCopyImageToBuffer2KHR")
	PFN_vkCmdBlitImage2KHR                                                = vulkan.NewProc("vkCmdBlitImage2KHR")
	PFN_vkCmdResolveImage2KHR                                             = vulkan.NewProc("vkCmdResolveImage2KHR")
	PFN_vkCreateDebugReportCallbackEXT                                    = vulkan.NewProc("vkCreateDebugReportCallbackEXT")
	PFN_vkDestroyDebugReportCallbackEXT                                   = vulkan.NewProc("vkDestroyDebugReportCallbackEXT")
	PFN_vkDebugReportMessageEXT                                           = vulkan.NewProc("vkDebugReportMessageEXT")
	PFN_vkDebugMarkerSetObjectTagEXT                                      = vulkan.NewProc("vkDebugMarkerSetObjectTagEXT")
	PFN_vkDebugMarkerSetObjectNameEXT                                     = vulkan.NewProc("vkDebugMarkerSetObjectNameEXT")
	PFN_vkCmdDebugMarkerBeginEXT                                          = vulkan.NewProc("vkCmdDebugMarkerBeginEXT")
	PFN_vkCmdDebugMarkerEndEXT                                            = vulkan.NewProc("vkCmdDebugMarkerEndEXT")
	PFN_vkCmdDebugMarkerInsertEXT                                         = vulkan.NewProc("vkCmdDebugMarkerInsertEXT")
	PFN_vkCmdBindTransformFeedbackBuffersEXT                              = vulkan.NewProc("vkCmdBindTransformFeedbackBuffersEXT")
	PFN_vkCmdBeginTransformFeedbackEXT                                    = vulkan.NewProc("vkCmdBeginTransformFeedbackEXT")
	PFN_vkCmdEndTransformFeedbackEXT                                      = vulkan.NewProc("vkCmdEndTransformFeedbackEXT")
	PFN_vkCmdBeginQueryIndexedEXT                                         = vulkan.NewProc("vkCmdBeginQueryIndexedEXT")
	PFN_vkCmdEndQueryIndexedEXT                                           = vulkan.NewProc("vkCmdEndQueryIndexedEXT")
	PFN_vkCmdDrawIndirectByteCountEXT                                     = vulkan.NewProc("vkCmdDrawIndirectByteCountEXT")
	PFN_vkCreateCuModuleNVX                                               = vulkan.NewProc("vkCreateCuModuleNVX")
	PFN_vkCreateCuFunctionNVX                                             = vulkan.NewProc("vkCreateCuFunctionNVX")
	PFN_vkDestroyCuModuleNVX                                              = vulkan.NewProc("vkDestroyCuModuleNVX")
	PFN_vkDestroyCuFunctionNVX                                            = vulkan.NewProc("vkDestroyCuFunctionNVX")
	PFN_vkCmdCuLaunchKernelNVX                                            = vulkan.NewProc("vkCmdCuLaunchKernelNVX")
	PFN_vkGetImageViewHandleNVX                                           = vulkan.NewProc("vkGetImageViewHandleNVX")
	PFN_vkGetImageViewAddressNVX                                          = vulkan.NewProc("vkGetImageViewAddressNVX")
	PFN_vkCmdDrawIndirectCountAMD                                         = vulkan.NewProc("vkCmdDrawIndirectCountAMD")
	PFN_vkCmdDrawIndexedIndirectCountAMD                                  = vulkan.NewProc("vkCmdDrawIndexedIndirectCountAMD")
	PFN_vkGetShaderInfoAMD                                                = vulkan.NewProc("vkGetShaderInfoAMD")
	PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV                = vulkan.NewProc("vkGetPhysicalDeviceExternalImageFormatPropertiesNV")
	PFN_vkCmdBeginConditionalRenderingEXT                                 = vulkan.NewProc("vkCmdBeginConditionalRenderingEXT")
	PFN_vkCmdEndConditionalRenderingEXT                                   = vulkan.NewProc("vkCmdEndConditionalRenderingEXT")
	PFN_vkCmdSetViewportWScalingNV                                        = vulkan.NewProc("vkCmdSetViewportWScalingNV")
	PFN_vkReleaseDisplayEXT                                               = vulkan.NewProc("vkReleaseDisplayEXT")
	PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT                        = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilities2EXT")
	PFN_vkDisplayPowerControlEXT                                          = vulkan.NewProc("vkDisplayPowerControlEXT")
	PFN_vkRegisterDeviceEventEXT                                          = vulkan.NewProc("vkRegisterDeviceEventEXT")
	PFN_vkRegisterDisplayEventEXT                                         = vulkan.NewProc("vkRegisterDisplayEventEXT")
	PFN_vkGetSwapchainCounterEXT                                          = vulkan.NewProc("vkGetSwapchainCounterEXT")
	PFN_vkGetRefreshCycleDurationGOOGLE                                   = vulkan.NewProc("vkGetRefreshCycleDurationGOOGLE")
	PFN_vkGetPastPresentationTimingGOOGLE                                 = vulkan.NewProc("vkGetPastPresentationTimingGOOGLE")
	PFN_vkCmdSetDiscardRectangleEXT                                       = vulkan.NewProc("vkCmdSetDiscardRectangleEXT")
	PFN_vkSetHdrMetadataEXT                                               = vulkan.NewProc("vkSetHdrMetadataEXT")
	PFN_vkSetDebugUtilsObjectNameEXT                                      = vulkan.NewProc("vkSetDebugUtilsObjectNameEXT")
	PFN_vkSetDebugUtilsObjectTagEXT                                       = vulkan.NewProc("vkSetDebugUtilsObjectTagEXT")
	PFN_vkQueueBeginDebugUtilsLabelEXT                                    = vulkan.NewProc("vkQueueBeginDebugUtilsLabelEXT")
	PFN_vkQueueEndDebugUtilsLabelEXT                                      = vulkan.NewProc("vkQueueEndDebugUtilsLabelEXT")
	PFN_vkQueueInsertDebugUtilsLabelEXT                                   = vulkan.NewProc("vkQueueInsertDebugUtilsLabelEXT")
	PFN_vkCmdBeginDebugUtilsLabelEXT                                      = vulkan.NewProc("vkCmdBeginDebugUtilsLabelEXT")
	PFN_vkCmdEndDebugUtilsLabelEXT                                        = vulkan.NewProc("vkCmdEndDebugUtilsLabelEXT")
	PFN_vkCmdInsertDebugUtilsLabelEXT                                     = vulkan.NewProc("vkCmdInsertDebugUtilsLabelEXT")
	PFN_vkCreateDebugUtilsMessengerEXT                                    = vulkan.NewProc("vkCreateDebugUtilsMessengerEXT")
	PFN_vkDestroyDebugUtilsMessengerEXT                                   = vulkan.NewProc("vkDestroyDebugUtilsMessengerEXT")
	PFN_vkSubmitDebugUtilsMessageEXT                                      = vulkan.NewProc("vkSubmitDebugUtilsMessageEXT")
	PFN_vkCmdSetSampleLocationsEXT                                        = vulkan.NewProc("vkCmdSetSampleLocationsEXT")
	PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT                       = vulkan.NewProc("vkGetPhysicalDeviceMultisamplePropertiesEXT")
	PFN_vkGetImageDrmFormatModifierPropertiesEXT                          = vulkan.NewProc("vkGetImageDrmFormatModifierPropertiesEXT")
	PFN_vkCreateValidationCacheEXT                                        = vulkan.NewProc("vkCreateValidationCacheEXT")
	PFN_vkDestroyValidationCacheEXT                                       = vulkan.NewProc("vkDestroyValidationCacheEXT")
	PFN_vkMergeValidationCachesEXT                                        = vulkan.NewProc("vkMergeValidationCachesEXT")
	PFN_vkGetValidationCacheDataEXT                                       = vulkan.NewProc("vkGetValidationCacheDataEXT")
	PFN_vkCmdBindShadingRateImageNV                                       = vulkan.NewProc("vkCmdBindShadingRateImageNV")
	PFN_vkCmdSetViewportShadingRatePaletteNV                              = vulkan.NewProc("vkCmdSetViewportShadingRatePaletteNV")
	PFN_vkCmdSetCoarseSampleOrderNV                                       = vulkan.NewProc("vkCmdSetCoarseSampleOrderNV")
	PFN_vkCreateAccelerationStructureNV                                   = vulkan.NewProc("vkCreateAccelerationStructureNV")
	PFN_vkDestroyAccelerationStructureNV                                  = vulkan.NewProc("vkDestroyAccelerationStructureNV")
	PFN_vkGetAccelerationStructureMemoryRequirementsNV                    = vulkan.NewProc("vkGetAccelerationStructureMemoryRequirementsNV")
	PFN_vkBindAccelerationStructureMemoryNV                               = vulkan.NewProc("vkBindAccelerationStructureMemoryNV")
	PFN_vkCmdBuildAccelerationStructureNV                                 = vulkan.NewProc("vkCmdBuildAccelerationStructureNV")
	PFN_vkCmdCopyAccelerationStructureNV                                  = vulkan.NewProc("vkCmdCopyAccelerationStructureNV")
	PFN_vkCmdTraceRaysNV                                                  = vulkan.NewProc("vkCmdTraceRaysNV")
	PFN_vkCreateRayTracingPipelinesNV                                     = vulkan.NewProc("vkCreateRayTracingPipelinesNV")
	PFN_vkGetRayTracingShaderGroupHandlesKHR                              = vulkan.NewProc("vkGetRayTracingShaderGroupHandlesKHR")
	PFN_vkGetRayTracingShaderGroupHandlesNV                               = vulkan.NewProc("vkGetRayTracingShaderGroupHandlesNV")
	PFN_vkGetAccelerationStructureHandleNV                                = vulkan.NewProc("vkGetAccelerationStructureHandleNV")
	PFN_vkCmdWriteAccelerationStructuresPropertiesNV                      = vulkan.NewProc("vkCmdWriteAccelerationStructuresPropertiesNV")
	PFN_vkCompileDeferredNV                                               = vulkan.NewProc("vkCompileDeferredNV")
	PFN_vkGetMemoryHostPointerPropertiesEXT                               = vulkan.NewProc("vkGetMemoryHostPointerPropertiesEXT")
	PFN_vkCmdWriteBufferMarkerAMD                                         = vulkan.NewProc("vkCmdWriteBufferMarkerAMD")
	PFN_vkGetPhysicalDeviceCalibrateableTimeDomainsEXT                    = vulkan.NewProc("vkGetPhysicalDeviceCalibrateableTimeDomainsEXT")
	PFN_vkGetCalibratedTimestampsEXT                                      = vulkan.NewProc("vkGetCalibratedTimestampsEXT")
	PFN_vkCmdDrawMeshTasksNV                                              = vulkan.NewProc("vkCmdDrawMeshTasksNV")
	PFN_vkCmdDrawMeshTasksIndirectNV                                      = vulkan.NewProc("vkCmdDrawMeshTasksIndirectNV")
	PFN_vkCmdDrawMeshTasksIndirectCountNV                                 = vulkan.NewProc("vkCmdDrawMeshTasksIndirectCountNV")
	PFN_vkCmdSetExclusiveScissorNV                                        = vulkan.NewProc("vkCmdSetExclusiveScissorNV")
	PFN_vkCmdSetCheckpointNV                                              = vulkan.NewProc("vkCmdSetCheckpointNV")
	PFN_vkGetQueueCheckpointDataNV                                        = vulkan.NewProc("vkGetQueueCheckpointDataNV")
	PFN_vkInitializePerformanceApiINTEL                                   = vulkan.NewProc("vkInitializePerformanceApiINTEL")
	PFN_vkUninitializePerformanceApiINTEL                                 = vulkan.NewProc("vkUninitializePerformanceApiINTEL")
	PFN_vkCmdSetPerformanceMarkerINTEL                                    = vulkan.NewProc("vkCmdSetPerformanceMarkerINTEL")
	PFN_vkCmdSetPerformanceStreamMarkerINTEL                              = vulkan.NewProc("vkCmdSetPerformanceStreamMarkerINTEL")
	PFN_vkCmdSetPerformanceOverrideINTEL                                  = vulkan.NewProc("vkCmdSetPerformanceOverrideINTEL")
	PFN_vkAcquirePerformanceConfigurationINTEL                            = vulkan.NewProc("vkAcquirePerformanceConfigurationINTEL")
	PFN_vkReleasePerformanceConfigurationINTEL                            = vulkan.NewProc("vkReleasePerformanceConfigurationINTEL")
	PFN_vkQueueSetPerformanceConfigurationINTEL                           = vulkan.NewProc("vkQueueSetPerformanceConfigurationINTEL")
	PFN_vkGetPerformanceParameterINTEL                                    = vulkan.NewProc("vkGetPerformanceParameterINTEL")
	PFN_vkSetLocalDimmingAMD                                              = vulkan.NewProc("vkSetLocalDimmingAMD")
	PFN_vkGetBufferDeviceAddressEXT                                       = vulkan.NewProc("vkGetBufferDeviceAddressEXT")
	PFN_vkGetPhysicalDeviceToolPropertiesEXT                              = vulkan.NewProc("vkGetPhysicalDeviceToolPropertiesEXT")
	PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV                  = vulkan.NewProc("vkGetPhysicalDeviceCooperativeMatrixPropertiesNV")
	PFN_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV = vulkan.NewProc("vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV")
	PFN_vkCreateHeadlessSurfaceEXT                                        = vulkan.NewProc("vkCreateHeadlessSurfaceEXT")
	PFN_vkCmdSetLineStippleEXT                                            = vulkan.NewProc("vkCmdSetLineStippleEXT")
	PFN_vkResetQueryPoolEXT                                               = vulkan.NewProc("vkResetQueryPoolEXT")
	PFN_vkCmdSetCullModeEXT                                               = vulkan.NewProc("vkCmdSetCullModeEXT")
	PFN_vkCmdSetFrontFaceEXT                                              = vulkan.NewProc("vkCmdSetFrontFaceEXT")
	PFN_vkCmdSetPrimitiveTopologyEXT                                      = vulkan.NewProc("vkCmdSetPrimitiveTopologyEXT")
	PFN_vkCmdSetViewportWithCountEXT                                      = vulkan.NewProc("vkCmdSetViewportWithCountEXT")
	PFN_vkCmdSetScissorWithCountEXT                                       = vulkan.NewProc("vkCmdSetScissorWithCountEXT")
	PFN_vkCmdBindVertexBuffers2EXT                                        = vulkan.NewProc("vkCmdBindVertexBuffers2EXT")
	PFN_vkCmdSetDepthTestEnableEXT                                        = vulkan.NewProc("vkCmdSetDepthTestEnableEXT")
	PFN_vkCmdSetDepthWriteEnableEXT                                       = vulkan.NewProc("vkCmdSetDepthWriteEnableEXT")
	PFN_vkCmdSetDepthCompareOpEXT                                         = vulkan.NewProc("vkCmdSetDepthCompareOpEXT")
	PFN_vkCmdSetDepthBoundsTestEnableEXT                                  = vulkan.NewProc("vkCmdSetDepthBoundsTestEnableEXT")
	PFN_vkCmdSetStencilTestEnableEXT                                      = vulkan.NewProc("vkCmdSetStencilTestEnableEXT")
	PFN_vkCmdSetStencilOpEXT                                              = vulkan.NewProc("vkCmdSetStencilOpEXT")
	PFN_vkGetGeneratedCommandsMemoryRequirementsNV                        = vulkan.NewProc("vkGetGeneratedCommandsMemoryRequirementsNV")
	PFN_vkCmdPreprocessGeneratedCommandsNV                                = vulkan.NewProc("vkCmdPreprocessGeneratedCommandsNV")
	PFN_vkCmdExecuteGeneratedCommandsNV                                   = vulkan.NewProc("vkCmdExecuteGeneratedCommandsNV")
	PFN_vkCmdBindPipelineShaderGroupNV                                    = vulkan.NewProc("vkCmdBindPipelineShaderGroupNV")
	PFN_vkCreateIndirectCommandsLayoutNV                                  = vulkan.NewProc("vkCreateIndirectCommandsLayoutNV")
	PFN_vkDestroyIndirectCommandsLayoutNV                                 = vulkan.NewProc("vkDestroyIndirectCommandsLayoutNV")
	PFN_vkAcquireDrmDisplayEXT                                            = vulkan.NewProc("vkAcquireDrmDisplayEXT")
	PFN_vkGetDrmDisplayEXT                                                = vulkan.NewProc("vkGetDrmDisplayEXT")
	PFN_vkCreatePrivateDataSlotEXT                                        = vulkan.NewProc("vkCreatePrivateDataSlotEXT")
	PFN_vkDestroyPrivateDataSlotEXT                                       = vulkan.NewProc("vkDestroyPrivateDataSlotEXT")
	PFN_vkSetPrivateDataEXT                                               = vulkan.NewProc("vkSetPrivateDataEXT")
	PFN_vkGetPrivateDataEXT                                               = vulkan.NewProc("vkGetPrivateDataEXT")
	PFN_vkCmdSetFragmentShadingRateEnumNV                                 = vulkan.NewProc("vkCmdSetFragmentShadingRateEnumNV")
	PFN_vkAcquireWinrtDisplayNV                                           = vulkan.NewProc("vkAcquireWinrtDisplayNV")
	PFN_vkGetWinrtDisplayNV                                               = vulkan.NewProc("vkGetWinrtDisplayNV")
	PFN_vkCmdSetVertexInputEXT                                            = vulkan.NewProc("vkCmdSetVertexInputEXT")
	PFN_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI                   = vulkan.NewProc("vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI")
	PFN_vkCmdSubpassShadingHUAWEI                                         = vulkan.NewProc("vkCmdSubpassShadingHUAWEI")
	PFN_vkCmdBindInvocationMaskHUAWEI                                     = vulkan.NewProc("vkCmdBindInvocationMaskHUAWEI")
	PFN_vkGetMemoryRemoteAddressNV                                        = vulkan.NewProc("vkGetMemoryRemoteAddressNV")
	PFN_vkCmdSetPatchControlPointsEXT                                     = vulkan.NewProc("vkCmdSetPatchControlPointsEXT")
	PFN_vkCmdSetRasterizerDiscardEnableEXT                                = vulkan.NewProc("vkCmdSetRasterizerDiscardEnableEXT")
	PFN_vkCmdSetDepthBiasEnableEXT                                        = vulkan.NewProc("vkCmdSetDepthBiasEnableEXT")
	PFN_vkCmdSetLogicOpEXT                                                = vulkan.NewProc("vkCmdSetLogicOpEXT")
	PFN_vkCmdSetPrimitiveRestartEnableEXT                                 = vulkan.NewProc("vkCmdSetPrimitiveRestartEnableEXT")
	PFN_vkCmdSetColorWriteEnableEXT                                       = vulkan.NewProc("vkCmdSetColorWriteEnableEXT")
	PFN_vkCmdDrawMultiEXT                                                 = vulkan.NewProc("vkCmdDrawMultiEXT")
	PFN_vkCmdDrawMultiIndexedEXT                                          = vulkan.NewProc("vkCmdDrawMultiIndexedEXT")
	PFN_vkCreateAccelerationStructureKHR                                  = vulkan.NewProc("vkCreateAccelerationStructureKHR")
	PFN_vkDestroyAccelerationStructureKHR                                 = vulkan.NewProc("vkDestroyAccelerationStructureKHR")
	PFN_vkCmdBuildAccelerationStructuresKHR                               = vulkan.NewProc("vkCmdBuildAccelerationStructuresKHR")
	PFN_vkCmdBuildAccelerationStructuresIndirectKHR                       = vulkan.NewProc("vkCmdBuildAccelerationStructuresIndirectKHR")
	PFN_vkBuildAccelerationStructuresKHR                                  = vulkan.NewProc("vkBuildAccelerationStructuresKHR")
	PFN_vkCopyAccelerationStructureKHR                                    = vulkan.NewProc("vkCopyAccelerationStructureKHR")
	PFN_vkCopyAccelerationStructureToMemoryKHR                            = vulkan.NewProc("vkCopyAccelerationStructureToMemoryKHR")
	PFN_vkCopyMemoryToAccelerationStructureKHR                            = vulkan.NewProc("vkCopyMemoryToAccelerationStructureKHR")
	PFN_vkWriteAccelerationStructuresPropertiesKHR                        = vulkan.NewProc("vkWriteAccelerationStructuresPropertiesKHR")
	PFN_vkCmdCopyAccelerationStructureKHR                                 = vulkan.NewProc("vkCmdCopyAccelerationStructureKHR")
	PFN_vkCmdCopyAccelerationStructureToMemoryKHR                         = vulkan.NewProc("vkCmdCopyAccelerationStructureToMemoryKHR")
	PFN_vkCmdCopyMemoryToAccelerationStructureKHR                         = vulkan.NewProc("vkCmdCopyMemoryToAccelerationStructureKHR")
	PFN_vkGetAccelerationStructureDeviceAddressKHR                        = vulkan.NewProc("vkGetAccelerationStructureDeviceAddressKHR")
	PFN_vkCmdWriteAccelerationStructuresPropertiesKHR                     = vulkan.NewProc("vkCmdWriteAccelerationStructuresPropertiesKHR")
	PFN_vkGetDeviceAccelerationStructureCompatibilityKHR                  = vulkan.NewProc("vkGetDeviceAccelerationStructureCompatibilityKHR")
	PFN_vkGetAccelerationStructureBuildSizesKHR                           = vulkan.NewProc("vkGetAccelerationStructureBuildSizesKHR")
	PFN_vkCmdTraceRaysKHR                                                 = vulkan.NewProc("vkCmdTraceRaysKHR")
	PFN_vkCreateRayTracingPipelinesKHR                                    = vulkan.NewProc("vkCreateRayTracingPipelinesKHR")
	PFN_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR                 = vulkan.NewProc("vkGetRayTracingCaptureReplayShaderGroupHandlesKHR")
	PFN_vkCmdTraceRaysIndirectKHR                                         = vulkan.NewProc("vkCmdTraceRaysIndirectKHR")
	PFN_vkGetRayTracingShaderGroupStackSizeKHR                            = vulkan.NewProc("vkGetRayTracingShaderGroupStackSizeKHR")
	PFN_vkCmdSetRayTracingPipelineStackSizeKHR                            = vulkan.NewProc("vkCmdSetRayTracingPipelineStackSizeKHR")
	PFN_vkCreateWin32SurfaceKHR                                           = vulkan.NewProc("vkCreateWin32SurfaceKHR")
	PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR                    = vulkan.NewProc("vkGetPhysicalDeviceWin32PresentationSupportKHR")
	PFN_vkGetMemoryWin32HandleKHR                                         = vulkan.NewProc("vkGetMemoryWin32HandleKHR")
	PFN_vkGetMemoryWin32HandlePropertiesKHR                               = vulkan.NewProc("vkGetMemoryWin32HandlePropertiesKHR")
	PFN_vkImportSemaphoreWin32HandleKHR                                   = vulkan.NewProc("vkImportSemaphoreWin32HandleKHR")
	PFN_vkGetSemaphoreWin32HandleKHR                                      = vulkan.NewProc("vkGetSemaphoreWin32HandleKHR")
	PFN_vkImportFenceWin32HandleKHR                                       = vulkan.NewProc("vkImportFenceWin32HandleKHR")
	PFN_vkGetFenceWin32HandleKHR                                          = vulkan.NewProc("vkGetFenceWin32HandleKHR")
	PFN_vkGetMemoryWin32HandleNV                                          = vulkan.NewProc("vkGetMemoryWin32HandleNV")
	PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT                        = vulkan.NewProc("vkGetPhysicalDeviceSurfacePresentModes2EXT")
	PFN_vkAcquireFullScreenExclusiveModeEXT                               = vulkan.NewProc("vkAcquireFullScreenExclusiveModeEXT")
	PFN_vkReleaseFullScreenExclusiveModeEXT                               = vulkan.NewProc("vkReleaseFullScreenExclusiveModeEXT")
	PFN_vkGetDeviceGroupSurfacePresentModes2EXT                           = vulkan.NewProc("vkGetDeviceGroupSurfacePresentModes2EXT")
)

type (
	VkWin32SurfaceCreateFlagsKHR = VkFlags
)

const (
	VK_KHR_win32_surface                           uint32 = 1
	VK_KHR_WIN32_SURFACE_SPEC_VERSION              uint32 = 6
	VK_KHR_WIN32_SURFACE_EXTENSION_NAME            string = "VK_KHR_win32_surface"
	VK_KHR_external_memory_win32                   uint32 = 1
	VK_KHR_EXTERNAL_MEMORY_WIN32_SPEC_VERSION      uint32 = 1
	VK_KHR_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME    string = "VK_KHR_external_memory_win32"
	VK_KHR_win32_keyed_mutex                       uint32 = 1
	VK_KHR_WIN32_KEYED_MUTEX_SPEC_VERSION          uint32 = 1
	VK_KHR_WIN32_KEYED_MUTEX_EXTENSION_NAME        string = "VK_KHR_win32_keyed_mutex"
	VK_KHR_external_semaphore_win32                uint32 = 1
	VK_KHR_EXTERNAL_SEMAPHORE_WIN32_SPEC_VERSION   uint32 = 1
	VK_KHR_EXTERNAL_SEMAPHORE_WIN32_EXTENSION_NAME string = "VK_KHR_external_semaphore_win32"
	VK_KHR_external_fence_win32                    uint32 = 1
	VK_KHR_EXTERNAL_FENCE_WIN32_SPEC_VERSION       uint32 = 1
	VK_KHR_EXTERNAL_FENCE_WIN32_EXTENSION_NAME     string = "VK_KHR_external_fence_win32"
	VK_NV_external_memory_win32                    uint32 = 1
	VK_NV_EXTERNAL_MEMORY_WIN32_SPEC_VERSION       uint32 = 1
	VK_NV_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME     string = "VK_NV_external_memory_win32"
	VK_NV_win32_keyed_mutex                        uint32 = 1
	VK_NV_WIN32_KEYED_MUTEX_SPEC_VERSION           uint32 = 2
	VK_NV_WIN32_KEYED_MUTEX_EXTENSION_NAME         string = "VK_NV_win32_keyed_mutex"
	VK_EXT_full_screen_exclusive                   uint32 = 1
	VK_EXT_FULL_SCREEN_EXCLUSIVE_SPEC_VERSION      uint32 = 4
	VK_EXT_FULL_SCREEN_EXCLUSIVE_EXTENSION_NAME    string = "VK_EXT_full_screen_exclusive"
)

type VkFullScreenExclusiveEXT = uint32

const (
	VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT                VkFullScreenExclusiveEXT = 0
	VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT                VkFullScreenExclusiveEXT = 1
	VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT             VkFullScreenExclusiveEXT = 2
	VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT VkFullScreenExclusiveEXT = 3
	VK_FULL_SCREEN_EXCLUSIVE_MAX_ENUM_EXT               VkFullScreenExclusiveEXT = 0x7FFFFFFF
)

type VkWin32SurfaceCreateInfoKHR struct {
	SType     VkStructureType
	PNext     uintptr
	Flags     VkWin32SurfaceCreateFlagsKHR
	Hinstance HINSTANCE
	Hwnd      HWND
}
type VkImportMemoryWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalMemoryHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}
type VkExportMemoryWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       uintptr
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}
type VkMemoryWin32HandlePropertiesKHR struct {
	SType          VkStructureType
	PNext          uintptr
	MemoryTypeBits uint32
}
type VkMemoryGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}
type VkWin32KeyedMutexAcquireReleaseInfoKHR struct {
	SType            VkStructureType
	PNext            uintptr
	AcquireCount     uint32
	PAcquireSyncs    *VkDeviceMemory
	PAcquireKeys     *uint64
	PAcquireTimeouts *uint32
	ReleaseCount     uint32
	PReleaseSyncs    *VkDeviceMemory
	PReleaseKeys     *uint64
}
type VkImportSemaphoreWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Semaphore  VkSemaphore
	Flags      VkSemaphoreImportFlags
	HandleType VkExternalSemaphoreHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}
type VkExportSemaphoreWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       uintptr
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}
type VkD3D12FenceSubmitInfoKHR struct {
	SType                      VkStructureType
	PNext                      uintptr
	WaitSemaphoreValuesCount   uint32
	PWaitSemaphoreValues       *uint64
	SignalSemaphoreValuesCount uint32
	PSignalSemaphoreValues     *uint64
}
type VkSemaphoreGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkImportFenceWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Fence      VkFence
	Flags      VkFenceImportFlags
	HandleType VkExternalFenceHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}
type VkExportFenceWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       uintptr
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}
type VkFenceGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
}
type VkImportMemoryWin32HandleInfoNV struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalMemoryHandleTypeFlagsNV
	Handle     HANDLE
}
type VkExportMemoryWin32HandleInfoNV struct {
	SType       VkStructureType
	PNext       uintptr
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
}
type VkWin32KeyedMutexAcquireReleaseInfoNV struct {
	SType                       VkStructureType
	PNext                       uintptr
	AcquireCount                uint32
	PAcquireSyncs               *VkDeviceMemory
	PAcquireKeys                *uint64
	PAcquireTimeoutMilliseconds *uint32
	ReleaseCount                uint32
	PReleaseSyncs               *VkDeviceMemory
	PReleaseKeys                *uint64
}
type VkSurfaceFullScreenExclusiveInfoEXT struct {
	SType               VkStructureType
	PNext               uintptr
	FullScreenExclusive VkFullScreenExclusiveEXT
}
type VkSurfaceCapabilitiesFullScreenExclusiveEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	FullScreenExclusiveSupported VkBool32
}
type VkSurfaceFullScreenExclusiveWin32InfoEXT struct {
	SType    VkStructureType
	PNext    uintptr
	Hmonitor HMONITOR
}

func CreateWin32SurfaceKHR(
	instance VkInstance,
	pCreateInfo *VkWin32SurfaceCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pSurface VkSurfaceKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateWin32SurfaceKHR.Addr(),
		4,
		uintptr(instance),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSurface)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceWin32PresentationSupportKHR(
	physicalDevice VkPhysicalDevice,
	queueFamilyIndex uint32,
) (result VkBool32) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(queueFamilyIndex),
		0,
	)
	result = (VkBool32)(ret)
	return
}
func GetMemoryWin32HandleKHR(
	device VkDevice,
	pGetWin32HandleInfo *VkMemoryGetWin32HandleInfoKHR,
	pHandle *HANDLE,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetMemoryWin32HandleKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetWin32HandleInfo)),
		uintptr(unsafe.Pointer(pHandle)),
	)
	result = GetError((int32)(ret))
	return
}
func GetMemoryWin32HandlePropertiesKHR(
	device VkDevice,
	handleType VkExternalMemoryHandleTypeFlagBits,
	handle HANDLE,
	pMemoryWin32HandleProperties *VkMemoryWin32HandlePropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetMemoryWin32HandlePropertiesKHR.Addr(),
		4,
		uintptr(device),
		uintptr(handleType),
		uintptr(handle),
		uintptr(unsafe.Pointer(pMemoryWin32HandleProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func ImportSemaphoreWin32HandleKHR(
	device VkDevice,
	pImportSemaphoreWin32HandleInfo *VkImportSemaphoreWin32HandleInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkImportSemaphoreWin32HandleKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pImportSemaphoreWin32HandleInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetSemaphoreWin32HandleKHR(
	device VkDevice,
	pGetWin32HandleInfo *VkSemaphoreGetWin32HandleInfoKHR,
	pHandle *HANDLE,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetSemaphoreWin32HandleKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetWin32HandleInfo)),
		uintptr(unsafe.Pointer(pHandle)),
	)
	result = GetError((int32)(ret))
	return
}
func ImportFenceWin32HandleKHR(
	device VkDevice,
	pImportFenceWin32HandleInfo *VkImportFenceWin32HandleInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkImportFenceWin32HandleKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pImportFenceWin32HandleInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetFenceWin32HandleKHR(
	device VkDevice,
	pGetWin32HandleInfo *VkFenceGetWin32HandleInfoKHR,
	pHandle *HANDLE,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetFenceWin32HandleKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetWin32HandleInfo)),
		uintptr(unsafe.Pointer(pHandle)),
	)
	result = GetError((int32)(ret))
	return
}
func GetMemoryWin32HandleNV(
	device VkDevice,
	memory VkDeviceMemory,
	handleType VkExternalMemoryHandleTypeFlagsNV,
	pHandle *HANDLE,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetMemoryWin32HandleNV.Addr(),
		4,
		uintptr(device),
		uintptr(memory),
		uintptr(handleType),
		uintptr(unsafe.Pointer(pHandle)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfacePresentModes2EXT(
	physicalDevice VkPhysicalDevice,
	pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR,
	pPresentModeCount *uint32,
	pPresentModes *VkPresentModeKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pSurfaceInfo)),
		uintptr(unsafe.Pointer(pPresentModeCount)),
		uintptr(unsafe.Pointer(pPresentModes)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func AcquireFullScreenExclusiveModeEXT(
	device VkDevice,
	swapchain VkSwapchainKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquireFullScreenExclusiveModeEXT.Addr(),
		2,
		uintptr(device),
		uintptr(swapchain),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func ReleaseFullScreenExclusiveModeEXT(
	device VkDevice,
	swapchain VkSwapchainKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkReleaseFullScreenExclusiveModeEXT.Addr(),
		2,
		uintptr(device),
		uintptr(swapchain),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceGroupSurfacePresentModes2EXT(
	device VkDevice,
	pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR,
	pModes *VkDeviceGroupPresentModeFlagsKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceGroupSurfacePresentModes2EXT.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pSurfaceInfo)),
		uintptr(unsafe.Pointer(pModes)),
	)
	result = GetError((int32)(ret))
	return
}
