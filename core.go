package vulkan

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

func VK_MAKE_VERSION(major, minor, patch uint32) (version uint32) {
	version = major<<22 | minor<<12 | patch
	return
}
func VK_MAKE_API_VERSION(variant, major, minor, patch uint32) (version uint32) {
	version = variant<<29 | major<<12 | minor<<12 | patch
	return
}

var VK_API_VERSION_1_0 uint32 = VK_MAKE_API_VERSION(0, 1, 0, 0)
var VK_API_VERSION_1_1 uint32 = VK_MAKE_API_VERSION(0, 1, 1, 0)
var VK_API_VERSION_1_2 uint32 = VK_MAKE_API_VERSION(0, 1, 2, 0)
var VK_HEADER_VERSION_COMPLETE uint32 = VK_MAKE_API_VERSION(0, 1, 2, VK_HEADER_VERSION)

type (
	VkBuffer                                               = uintptr
	VkImage                                                = uintptr
	VkInstance                                             = uintptr
	VkPhysicalDevice                                       = uintptr
	VkDevice                                               = uintptr
	VkQueue                                                = uintptr
	VkSemaphore                                            = uintptr
	VkCommandBuffer                                        = uintptr
	VkFence                                                = uintptr
	VkDeviceMemory                                         = uintptr
	VkEvent                                                = uintptr
	VkQueryPool                                            = uintptr
	VkBufferView                                           = uintptr
	VkImageView                                            = uintptr
	VkShaderModule                                         = uintptr
	VkPipelineCache                                        = uintptr
	VkPipelineLayout                                       = uintptr
	VkPipeline                                             = uintptr
	VkRenderPass                                           = uintptr
	VkDescriptorSetLayout                                  = uintptr
	VkSampler                                              = uintptr
	VkDescriptorSet                                        = uintptr
	VkDescriptorPool                                       = uintptr
	VkFramebuffer                                          = uintptr
	VkCommandPool                                          = uintptr
	VkSamplerYcbcrConversion                               = uintptr
	VkDescriptorUpdateTemplate                             = uintptr
	VkSurfaceKHR                                           = uintptr
	VkSwapchainKHR                                         = uintptr
	VkDisplayKHR                                           = uintptr
	VkDisplayModeKHR                                       = uintptr
	VkDeferredOperationKHR                                 = uintptr
	VkDebugReportCallbackEXT                               = uintptr
	VkCuModuleNVX                                          = uintptr
	VkCuFunctionNVX                                        = uintptr
	VkDebugUtilsMessengerEXT                               = uintptr
	VkValidationCacheEXT                                   = uintptr
	VkAccelerationStructureNV                              = uintptr
	VkPerformanceConfigurationINTEL                        = uintptr
	VkIndirectCommandsLayoutNV                             = uintptr
	VkPrivateDataSlotEXT                                   = uintptr
	VkAccelerationStructureKHR                             = uintptr
	VkBool32                                               = uint32
	VkDeviceAddress                                        = uint64
	VkDeviceSize                                           = uint64
	VkFlags                                                = uint32
	VkSampleMask                                           = uint32
	VkAccessFlags                                          = VkFlags
	VkImageAspectFlags                                     = VkFlags
	VkFormatFeatureFlags                                   = VkFlags
	VkImageCreateFlags                                     = VkFlags
	VkSampleCountFlags                                     = VkFlags
	VkImageUsageFlags                                      = VkFlags
	VkInstanceCreateFlags                                  = VkFlags
	VkMemoryHeapFlags                                      = VkFlags
	VkMemoryPropertyFlags                                  = VkFlags
	VkQueueFlags                                           = VkFlags
	VkDeviceCreateFlags                                    = VkFlags
	VkDeviceQueueCreateFlags                               = VkFlags
	VkPipelineStageFlags                                   = VkFlags
	VkMemoryMapFlags                                       = VkFlags
	VkSparseMemoryBindFlags                                = VkFlags
	VkSparseImageFormatFlags                               = VkFlags
	VkFenceCreateFlags                                     = VkFlags
	VkSemaphoreCreateFlags                                 = VkFlags
	VkEventCreateFlags                                     = VkFlags
	VkQueryPipelineStatisticFlags                          = VkFlags
	VkQueryPoolCreateFlags                                 = VkFlags
	VkQueryResultFlags                                     = VkFlags
	VkBufferCreateFlags                                    = VkFlags
	VkBufferUsageFlags                                     = VkFlags
	VkBufferViewCreateFlags                                = VkFlags
	VkImageViewCreateFlags                                 = VkFlags
	VkShaderModuleCreateFlags                              = VkFlags
	VkPipelineCacheCreateFlags                             = VkFlags
	VkColorComponentFlags                                  = VkFlags
	VkPipelineCreateFlags                                  = VkFlags
	VkPipelineShaderStageCreateFlags                       = VkFlags
	VkCullModeFlags                                        = VkFlags
	VkPipelineVertexInputStateCreateFlags                  = VkFlags
	VkPipelineInputAssemblyStateCreateFlags                = VkFlags
	VkPipelineTessellationStateCreateFlags                 = VkFlags
	VkPipelineViewportStateCreateFlags                     = VkFlags
	VkPipelineRasterizationStateCreateFlags                = VkFlags
	VkPipelineMultisampleStateCreateFlags                  = VkFlags
	VkPipelineDepthStencilStateCreateFlags                 = VkFlags
	VkPipelineColorBlendStateCreateFlags                   = VkFlags
	VkPipelineDynamicStateCreateFlags                      = VkFlags
	VkPipelineLayoutCreateFlags                            = VkFlags
	VkShaderStageFlags                                     = VkFlags
	VkSamplerCreateFlags                                   = VkFlags
	VkDescriptorPoolCreateFlags                            = VkFlags
	VkDescriptorPoolResetFlags                             = VkFlags
	VkDescriptorSetLayoutCreateFlags                       = VkFlags
	VkAttachmentDescriptionFlags                           = VkFlags
	VkDependencyFlags                                      = VkFlags
	VkFramebufferCreateFlags                               = VkFlags
	VkRenderPassCreateFlags                                = VkFlags
	VkSubpassDescriptionFlags                              = VkFlags
	VkCommandPoolCreateFlags                               = VkFlags
	VkCommandPoolResetFlags                                = VkFlags
	VkCommandBufferUsageFlags                              = VkFlags
	VkQueryControlFlags                                    = VkFlags
	VkCommandBufferResetFlags                              = VkFlags
	VkStencilFaceFlags                                     = VkFlags
	VkSubgroupFeatureFlags                                 = VkFlags
	VkPeerMemoryFeatureFlags                               = VkFlags
	VkMemoryAllocateFlags                                  = VkFlags
	VkCommandPoolTrimFlags                                 = VkFlags
	VkDescriptorUpdateTemplateCreateFlags                  = VkFlags
	VkExternalMemoryHandleTypeFlags                        = VkFlags
	VkExternalMemoryFeatureFlags                           = VkFlags
	VkExternalFenceHandleTypeFlags                         = VkFlags
	VkExternalFenceFeatureFlags                            = VkFlags
	VkFenceImportFlags                                     = VkFlags
	VkSemaphoreImportFlags                                 = VkFlags
	VkExternalSemaphoreHandleTypeFlags                     = VkFlags
	VkExternalSemaphoreFeatureFlags                        = VkFlags
	VkPhysicalDeviceVariablePointerFeatures                = VkPhysicalDeviceVariablePointersFeatures
	VkPhysicalDeviceShaderDrawParameterFeatures            = VkPhysicalDeviceShaderDrawParametersFeatures
	VkResolveModeFlags                                     = VkFlags
	VkDescriptorBindingFlags                               = VkFlags
	VkSemaphoreWaitFlags                                   = VkFlags
	VkCompositeAlphaFlagsKHR                               = VkFlags
	VkSurfaceTransformFlagsKHR                             = VkFlags
	VkSwapchainCreateFlagsKHR                              = VkFlags
	VkDeviceGroupPresentModeFlagsKHR                       = VkFlags
	VkDisplayModeCreateFlagsKHR                            = VkFlags
	VkDisplayPlaneAlphaFlagsKHR                            = VkFlags
	VkDisplaySurfaceCreateFlagsKHR                         = VkFlags
	VkRenderPassMultiviewCreateInfoKHR                     = VkRenderPassMultiviewCreateInfo
	VkPhysicalDeviceMultiviewFeaturesKHR                   = VkPhysicalDeviceMultiviewFeatures
	VkPhysicalDeviceMultiviewPropertiesKHR                 = VkPhysicalDeviceMultiviewProperties
	VkPhysicalDeviceFeatures2KHR                           = VkPhysicalDeviceFeatures2
	VkPhysicalDeviceProperties2KHR                         = VkPhysicalDeviceProperties2
	VkFormatProperties2KHR                                 = VkFormatProperties2
	VkImageFormatProperties2KHR                            = VkImageFormatProperties2
	VkPhysicalDeviceImageFormatInfo2KHR                    = VkPhysicalDeviceImageFormatInfo2
	VkQueueFamilyProperties2KHR                            = VkQueueFamilyProperties2
	VkPhysicalDeviceMemoryProperties2KHR                   = VkPhysicalDeviceMemoryProperties2
	VkSparseImageFormatProperties2KHR                      = VkSparseImageFormatProperties2
	VkPhysicalDeviceSparseImageFormatInfo2KHR              = VkPhysicalDeviceSparseImageFormatInfo2
	VkPeerMemoryFeatureFlagsKHR                            = VkPeerMemoryFeatureFlags
	VkPeerMemoryFeatureFlagBitsKHR                         = VkPeerMemoryFeatureFlagBits
	VkMemoryAllocateFlagsKHR                               = VkMemoryAllocateFlags
	VkMemoryAllocateFlagBitsKHR                            = VkMemoryAllocateFlagBits
	VkMemoryAllocateFlagsInfoKHR                           = VkMemoryAllocateFlagsInfo
	VkDeviceGroupRenderPassBeginInfoKHR                    = VkDeviceGroupRenderPassBeginInfo
	VkDeviceGroupCommandBufferBeginInfoKHR                 = VkDeviceGroupCommandBufferBeginInfo
	VkDeviceGroupSubmitInfoKHR                             = VkDeviceGroupSubmitInfo
	VkDeviceGroupBindSparseInfoKHR                         = VkDeviceGroupBindSparseInfo
	VkBindBufferMemoryDeviceGroupInfoKHR                   = VkBindBufferMemoryDeviceGroupInfo
	VkBindImageMemoryDeviceGroupInfoKHR                    = VkBindImageMemoryDeviceGroupInfo
	VkCommandPoolTrimFlagsKHR                              = VkCommandPoolTrimFlags
	VkPhysicalDeviceGroupPropertiesKHR                     = VkPhysicalDeviceGroupProperties
	VkDeviceGroupDeviceCreateInfoKHR                       = VkDeviceGroupDeviceCreateInfo
	VkExternalMemoryHandleTypeFlagsKHR                     = VkExternalMemoryHandleTypeFlags
	VkExternalMemoryHandleTypeFlagBitsKHR                  = VkExternalMemoryHandleTypeFlagBits
	VkExternalMemoryFeatureFlagsKHR                        = VkExternalMemoryFeatureFlags
	VkExternalMemoryFeatureFlagBitsKHR                     = VkExternalMemoryFeatureFlagBits
	VkExternalMemoryPropertiesKHR                          = VkExternalMemoryProperties
	VkPhysicalDeviceExternalImageFormatInfoKHR             = VkPhysicalDeviceExternalImageFormatInfo
	VkExternalImageFormatPropertiesKHR                     = VkExternalImageFormatProperties
	VkPhysicalDeviceExternalBufferInfoKHR                  = VkPhysicalDeviceExternalBufferInfo
	VkExternalBufferPropertiesKHR                          = VkExternalBufferProperties
	VkPhysicalDeviceIDPropertiesKHR                        = VkPhysicalDeviceIDProperties
	VkExternalMemoryImageCreateInfoKHR                     = VkExternalMemoryImageCreateInfo
	VkExternalMemoryBufferCreateInfoKHR                    = VkExternalMemoryBufferCreateInfo
	VkExportMemoryAllocateInfoKHR                          = VkExportMemoryAllocateInfo
	VkExternalSemaphoreHandleTypeFlagsKHR                  = VkExternalSemaphoreHandleTypeFlags
	VkExternalSemaphoreHandleTypeFlagBitsKHR               = VkExternalSemaphoreHandleTypeFlagBits
	VkExternalSemaphoreFeatureFlagsKHR                     = VkExternalSemaphoreFeatureFlags
	VkExternalSemaphoreFeatureFlagBitsKHR                  = VkExternalSemaphoreFeatureFlagBits
	VkPhysicalDeviceExternalSemaphoreInfoKHR               = VkPhysicalDeviceExternalSemaphoreInfo
	VkExternalSemaphorePropertiesKHR                       = VkExternalSemaphoreProperties
	VkSemaphoreImportFlagsKHR                              = VkSemaphoreImportFlags
	VkSemaphoreImportFlagBitsKHR                           = VkSemaphoreImportFlagBits
	VkExportSemaphoreCreateInfoKHR                         = VkExportSemaphoreCreateInfo
	VkPhysicalDeviceShaderFloat16Int8FeaturesKHR           = VkPhysicalDeviceShaderFloat16Int8Features
	VkPhysicalDeviceFloat16Int8FeaturesKHR                 = VkPhysicalDeviceShaderFloat16Int8Features
	VkPhysicalDevice16BitStorageFeaturesKHR                = VkPhysicalDevice16BitStorageFeatures
	VkDescriptorUpdateTemplateKHR                          = VkDescriptorUpdateTemplate
	VkDescriptorUpdateTemplateTypeKHR                      = VkDescriptorUpdateTemplateType
	VkDescriptorUpdateTemplateCreateFlagsKHR               = VkDescriptorUpdateTemplateCreateFlags
	VkDescriptorUpdateTemplateEntryKHR                     = VkDescriptorUpdateTemplateEntry
	VkDescriptorUpdateTemplateCreateInfoKHR                = VkDescriptorUpdateTemplateCreateInfo
	VkPhysicalDeviceImagelessFramebufferFeaturesKHR        = VkPhysicalDeviceImagelessFramebufferFeatures
	VkFramebufferAttachmentsCreateInfoKHR                  = VkFramebufferAttachmentsCreateInfo
	VkFramebufferAttachmentImageInfoKHR                    = VkFramebufferAttachmentImageInfo
	VkRenderPassAttachmentBeginInfoKHR                     = VkRenderPassAttachmentBeginInfo
	VkRenderPassCreateInfo2KHR                             = VkRenderPassCreateInfo2
	VkAttachmentDescription2KHR                            = VkAttachmentDescription2
	VkAttachmentReference2KHR                              = VkAttachmentReference2
	VkSubpassDescription2KHR                               = VkSubpassDescription2
	VkSubpassDependency2KHR                                = VkSubpassDependency2
	VkSubpassBeginInfoKHR                                  = VkSubpassBeginInfo
	VkSubpassEndInfoKHR                                    = VkSubpassEndInfo
	VkExternalFenceHandleTypeFlagsKHR                      = VkExternalFenceHandleTypeFlags
	VkExternalFenceHandleTypeFlagBitsKHR                   = VkExternalFenceHandleTypeFlagBits
	VkExternalFenceFeatureFlagsKHR                         = VkExternalFenceFeatureFlags
	VkExternalFenceFeatureFlagBitsKHR                      = VkExternalFenceFeatureFlagBits
	VkPhysicalDeviceExternalFenceInfoKHR                   = VkPhysicalDeviceExternalFenceInfo
	VkExternalFencePropertiesKHR                           = VkExternalFenceProperties
	VkFenceImportFlagsKHR                                  = VkFenceImportFlags
	VkFenceImportFlagBitsKHR                               = VkFenceImportFlagBits
	VkExportFenceCreateInfoKHR                             = VkExportFenceCreateInfo
	VkPerformanceCounterDescriptionFlagsKHR                = VkFlags
	VkAcquireProfilingLockFlagsKHR                         = VkFlags
	VkPointClippingBehaviorKHR                             = VkPointClippingBehavior
	VkTessellationDomainOriginKHR                          = VkTessellationDomainOrigin
	VkPhysicalDevicePointClippingPropertiesKHR             = VkPhysicalDevicePointClippingProperties
	VkRenderPassInputAttachmentAspectCreateInfoKHR         = VkRenderPassInputAttachmentAspectCreateInfo
	VkInputAttachmentAspectReferenceKHR                    = VkInputAttachmentAspectReference
	VkImageViewUsageCreateInfoKHR                          = VkImageViewUsageCreateInfo
	VkPipelineTessellationDomainOriginStateCreateInfoKHR   = VkPipelineTessellationDomainOriginStateCreateInfo
	VkPhysicalDeviceVariablePointerFeaturesKHR             = VkPhysicalDeviceVariablePointersFeatures
	VkPhysicalDeviceVariablePointersFeaturesKHR            = VkPhysicalDeviceVariablePointersFeatures
	VkMemoryDedicatedRequirementsKHR                       = VkMemoryDedicatedRequirements
	VkMemoryDedicatedAllocateInfoKHR                       = VkMemoryDedicatedAllocateInfo
	VkBufferMemoryRequirementsInfo2KHR                     = VkBufferMemoryRequirementsInfo2
	VkImageMemoryRequirementsInfo2KHR                      = VkImageMemoryRequirementsInfo2
	VkImageSparseMemoryRequirementsInfo2KHR                = VkImageSparseMemoryRequirementsInfo2
	VkMemoryRequirements2KHR                               = VkMemoryRequirements2
	VkSparseImageMemoryRequirements2KHR                    = VkSparseImageMemoryRequirements2
	VkImageFormatListCreateInfoKHR                         = VkImageFormatListCreateInfo
	VkSamplerYcbcrConversionKHR                            = VkSamplerYcbcrConversion
	VkSamplerYcbcrModelConversionKHR                       = VkSamplerYcbcrModelConversion
	VkSamplerYcbcrRangeKHR                                 = VkSamplerYcbcrRange
	VkChromaLocationKHR                                    = VkChromaLocation
	VkSamplerYcbcrConversionCreateInfoKHR                  = VkSamplerYcbcrConversionCreateInfo
	VkSamplerYcbcrConversionInfoKHR                        = VkSamplerYcbcrConversionInfo
	VkBindImagePlaneMemoryInfoKHR                          = VkBindImagePlaneMemoryInfo
	VkImagePlaneMemoryRequirementsInfoKHR                  = VkImagePlaneMemoryRequirementsInfo
	VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR      = VkPhysicalDeviceSamplerYcbcrConversionFeatures
	VkSamplerYcbcrConversionImageFormatPropertiesKHR       = VkSamplerYcbcrConversionImageFormatProperties
	VkBindBufferMemoryInfoKHR                              = VkBindBufferMemoryInfo
	VkBindImageMemoryInfoKHR                               = VkBindImageMemoryInfo
	VkPhysicalDeviceMaintenance3PropertiesKHR              = VkPhysicalDeviceMaintenance3Properties
	VkDescriptorSetLayoutSupportKHR                        = VkDescriptorSetLayoutSupport
	VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR = VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures
	VkPhysicalDevice8BitStorageFeaturesKHR                 = VkPhysicalDevice8BitStorageFeatures
	VkPhysicalDeviceShaderAtomicInt64FeaturesKHR           = VkPhysicalDeviceShaderAtomicInt64Features
	VkDriverIdKHR                                          = VkDriverId
	VkConformanceVersionKHR                                = VkConformanceVersion
	VkPhysicalDeviceDriverPropertiesKHR                    = VkPhysicalDeviceDriverProperties
	VkShaderFloatControlsIndependenceKHR                   = VkShaderFloatControlsIndependence
	VkPhysicalDeviceFloatControlsPropertiesKHR             = VkPhysicalDeviceFloatControlsProperties
	VkResolveModeFlagBitsKHR                               = VkResolveModeFlagBits
	VkResolveModeFlagsKHR                                  = VkResolveModeFlags
	VkSubpassDescriptionDepthStencilResolveKHR             = VkSubpassDescriptionDepthStencilResolve
	VkPhysicalDeviceDepthStencilResolvePropertiesKHR       = VkPhysicalDeviceDepthStencilResolveProperties
	VkSemaphoreTypeKHR                                     = VkSemaphoreType
	VkSemaphoreWaitFlagBitsKHR                             = VkSemaphoreWaitFlagBits
	VkSemaphoreWaitFlagsKHR                                = VkSemaphoreWaitFlags
	VkPhysicalDeviceTimelineSemaphoreFeaturesKHR           = VkPhysicalDeviceTimelineSemaphoreFeatures
	VkPhysicalDeviceTimelineSemaphorePropertiesKHR         = VkPhysicalDeviceTimelineSemaphoreProperties
	VkSemaphoreTypeCreateInfoKHR                           = VkSemaphoreTypeCreateInfo
	VkTimelineSemaphoreSubmitInfoKHR                       = VkTimelineSemaphoreSubmitInfo
	VkSemaphoreWaitInfoKHR                                 = VkSemaphoreWaitInfo
	VkSemaphoreSignalInfoKHR                               = VkSemaphoreSignalInfo
	VkPhysicalDeviceVulkanMemoryModelFeaturesKHR           = VkPhysicalDeviceVulkanMemoryModelFeatures
	VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR = VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures
	VkAttachmentReferenceStencilLayoutKHR                  = VkAttachmentReferenceStencilLayout
	VkAttachmentDescriptionStencilLayoutKHR                = VkAttachmentDescriptionStencilLayout
	VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR = VkPhysicalDeviceUniformBufferStandardLayoutFeatures
	VkPhysicalDeviceBufferDeviceAddressFeaturesKHR         = VkPhysicalDeviceBufferDeviceAddressFeatures
	VkBufferDeviceAddressInfoKHR                           = VkBufferDeviceAddressInfo
	VkBufferOpaqueCaptureAddressCreateInfoKHR              = VkBufferOpaqueCaptureAddressCreateInfo
	VkMemoryOpaqueCaptureAddressAllocateInfoKHR            = VkMemoryOpaqueCaptureAddressAllocateInfo
	VkDeviceMemoryOpaqueCaptureAddressInfoKHR              = VkDeviceMemoryOpaqueCaptureAddressInfo
	VkFlags64                                              = uint64
	VkPipelineStageFlags2KHR                               = VkFlags64
	VkPipelineStageFlagBits2KHR                            = VkFlags64
	VkAccessFlags2KHR                                      = VkFlags64
	VkAccessFlagBits2KHR                                   = VkFlags64
	VkSubmitFlagsKHR                                       = VkFlags
	VkDebugReportFlagsEXT                                  = VkFlags
	VkPipelineRasterizationStateStreamCreateFlagsEXT       = VkFlags
	VkExternalMemoryHandleTypeFlagsNV                      = VkFlags
	VkExternalMemoryFeatureFlagsNV                         = VkFlags
	VkConditionalRenderingFlagsEXT                         = VkFlags
	VkSurfaceCounterFlagsEXT                               = VkFlags
	VkPipelineViewportSwizzleStateCreateFlagsNV            = VkFlags
	VkPipelineDiscardRectangleStateCreateFlagsEXT          = VkFlags
	VkPipelineRasterizationConservativeStateCreateFlagsEXT = VkFlags
	VkPipelineRasterizationDepthClipStateCreateFlagsEXT    = VkFlags
	VkDebugUtilsMessengerCallbackDataFlagsEXT              = VkFlags
	VkDebugUtilsMessageTypeFlagsEXT                        = VkFlags
	VkDebugUtilsMessageSeverityFlagsEXT                    = VkFlags
	VkDebugUtilsMessengerCreateFlagsEXT                    = VkFlags
	VkSamplerReductionModeEXT                              = VkSamplerReductionMode
	VkSamplerReductionModeCreateInfoEXT                    = VkSamplerReductionModeCreateInfo
	VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT       = VkPhysicalDeviceSamplerFilterMinmaxProperties
	VkPipelineCoverageToColorStateCreateFlagsNV            = VkFlags
	VkPipelineCoverageModulationStateCreateFlagsNV         = VkFlags
	VkValidationCacheCreateFlagsEXT                        = VkFlags
	VkDescriptorBindingFlagBitsEXT                         = VkDescriptorBindingFlagBits
	VkDescriptorBindingFlagsEXT                            = VkDescriptorBindingFlags
	VkDescriptorSetLayoutBindingFlagsCreateInfoEXT         = VkDescriptorSetLayoutBindingFlagsCreateInfo
	VkPhysicalDeviceDescriptorIndexingFeaturesEXT          = VkPhysicalDeviceDescriptorIndexingFeatures
	VkPhysicalDeviceDescriptorIndexingPropertiesEXT        = VkPhysicalDeviceDescriptorIndexingProperties
	VkDescriptorSetVariableDescriptorCountAllocateInfoEXT  = VkDescriptorSetVariableDescriptorCountAllocateInfo
	VkDescriptorSetVariableDescriptorCountLayoutSupportEXT = VkDescriptorSetVariableDescriptorCountLayoutSupport
	VkRayTracingShaderGroupTypeNV                          = VkRayTracingShaderGroupTypeKHR
	VkGeometryTypeNV                                       = VkGeometryTypeKHR
	VkAccelerationStructureTypeNV                          = VkAccelerationStructureTypeKHR
	VkCopyAccelerationStructureModeNV                      = VkCopyAccelerationStructureModeKHR
	VkGeometryFlagsKHR                                     = VkFlags
	VkGeometryFlagsNV                                      = VkGeometryFlagsKHR
	VkGeometryFlagBitsNV                                   = VkGeometryFlagBitsKHR
	VkGeometryInstanceFlagsKHR                             = VkFlags
	VkGeometryInstanceFlagsNV                              = VkGeometryInstanceFlagsKHR
	VkGeometryInstanceFlagBitsNV                           = VkGeometryInstanceFlagBitsKHR
	VkBuildAccelerationStructureFlagsKHR                   = VkFlags
	VkBuildAccelerationStructureFlagsNV                    = VkBuildAccelerationStructureFlagsKHR
	VkBuildAccelerationStructureFlagBitsNV                 = VkBuildAccelerationStructureFlagBitsKHR
	VkTransformMatrixNV                                    = VkTransformMatrixKHR
	VkAabbPositionsNV                                      = VkAabbPositionsKHR
	VkAccelerationStructureInstanceNV                      = VkAccelerationStructureInstanceKHR
	VkPipelineCompilerControlFlagsAMD                      = VkFlags
	VkPipelineCreationFeedbackFlagsEXT                     = VkFlags
	VkQueryPoolCreateInfoINTEL                             = VkQueryPoolPerformanceQueryCreateInfoINTEL
	VkPhysicalDeviceScalarBlockLayoutFeaturesEXT           = VkPhysicalDeviceScalarBlockLayoutFeatures
	VkShaderCorePropertiesFlagsAMD                         = VkFlags
	VkPhysicalDeviceBufferAddressFeaturesEXT               = VkPhysicalDeviceBufferDeviceAddressFeaturesEXT
	VkBufferDeviceAddressInfoEXT                           = VkBufferDeviceAddressInfo
	VkToolPurposeFlagsEXT                                  = VkFlags
	VkImageStencilUsageCreateInfoEXT                       = VkImageStencilUsageCreateInfo
	VkPipelineCoverageReductionStateCreateFlagsNV          = VkFlags
	VkHeadlessSurfaceCreateFlagsEXT                        = VkFlags
	VkPhysicalDeviceHostQueryResetFeaturesEXT              = VkPhysicalDeviceHostQueryResetFeatures
	VkIndirectStateFlagsNV                                 = VkFlags
	VkIndirectCommandsLayoutUsageFlagsNV                   = VkFlags
	VkDeviceMemoryReportFlagsEXT                           = VkFlags
	VkPrivateDataSlotCreateFlagsEXT                        = VkFlags
	VkDeviceDiagnosticsConfigFlagsNV                       = VkFlags
	VkAccelerationStructureMotionInfoFlagsNV               = VkFlags
	VkAccelerationStructureMotionInstanceFlagsNV           = VkFlags
	VkRemoteAddressNV                                      = uintptr
	VkAccelerationStructureCreateFlagsKHR                  = VkFlags
)

const (
	VK_VERSION_1_0                                             uint32  = 1
	VK_NULL_HANDLE                                             uintptr = 0
	VK_HEADER_VERSION                                          uint32  = 189
	VK_UUID_SIZE                                               uint32  = 16
	VK_ATTACHMENT_UNUSED                                       uint32  = 0xFFFFFFFF
	VK_FALSE                                                   uint32  = 0
	VK_LOD_CLAMP_NONE                                          float32 = 1000.0
	VK_QUEUE_FAMILY_IGNORED                                    uint32  = 0xFFFFFFFF
	VK_REMAINING_ARRAY_LAYERS                                  uint32  = 0xFFFFFFFF
	VK_REMAINING_MIP_LEVELS                                    uint32  = 0xFFFFFFFF
	VK_SUBPASS_EXTERNAL                                        uint32  = 0xFFFFFFFF
	VK_TRUE                                                    uint32  = 1
	VK_WHOLE_SIZE                                              uint64  = 0xFFFFFFFFFFFFFFFF
	VK_MAX_MEMORY_TYPES                                        uint32  = 32
	VK_MAX_MEMORY_HEAPS                                        uint32  = 16
	VK_MAX_PHYSICAL_DEVICE_NAME_SIZE                           uint32  = 256
	VK_MAX_EXTENSION_NAME_SIZE                                 uint32  = 256
	VK_MAX_DESCRIPTION_SIZE                                    uint32  = 256
	VK_VERSION_1_1                                             uint32  = 1
	VK_MAX_DEVICE_GROUP_SIZE                                   uint32  = 32
	VK_LUID_SIZE                                               uint32  = 8
	VK_QUEUE_FAMILY_EXTERNAL                                   uint32  = 0x00000000
	VK_VERSION_1_2                                             uint32  = 1
	VK_MAX_DRIVER_NAME_SIZE                                    uint32  = 256
	VK_MAX_DRIVER_INFO_SIZE                                    uint32  = 256
	VK_KHR_surface                                             uint32  = 1
	VK_KHR_SURFACE_SPEC_VERSION                                uint32  = 25
	VK_KHR_SURFACE_EXTENSION_NAME                              string  = "VK_KHR_surface"
	VK_KHR_swapchain                                           uint32  = 1
	VK_KHR_SWAPCHAIN_SPEC_VERSION                              uint32  = 70
	VK_KHR_SWAPCHAIN_EXTENSION_NAME                            string  = "VK_KHR_swapchain"
	VK_KHR_display                                             uint32  = 1
	VK_KHR_DISPLAY_SPEC_VERSION                                uint32  = 23
	VK_KHR_DISPLAY_EXTENSION_NAME                              string  = "VK_KHR_display"
	VK_KHR_display_swapchain                                   uint32  = 1
	VK_KHR_DISPLAY_SWAPCHAIN_SPEC_VERSION                      uint32  = 10
	VK_KHR_DISPLAY_SWAPCHAIN_EXTENSION_NAME                    string  = "VK_KHR_display_swapchain"
	VK_KHR_sampler_mirror_clamp_to_edge                        uint32  = 1
	VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_SPEC_VERSION           uint32  = 3
	VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME         string  = "VK_KHR_sampler_mirror_clamp_to_edge"
	VK_KHR_multiview                                           uint32  = 1
	VK_KHR_MULTIVIEW_SPEC_VERSION                              uint32  = 1
	VK_KHR_MULTIVIEW_EXTENSION_NAME                            string  = "VK_KHR_multiview"
	VK_KHR_get_physical_device_properties2                     uint32  = 1
	VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_SPEC_VERSION       uint32  = 2
	VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_EXTENSION_NAME     string  = "VK_KHR_get_physical_device_properties2"
	VK_KHR_device_group                                        uint32  = 1
	VK_KHR_DEVICE_GROUP_SPEC_VERSION                           uint32  = 4
	VK_KHR_DEVICE_GROUP_EXTENSION_NAME                         string  = "VK_KHR_device_group"
	VK_KHR_shader_draw_parameters                              uint32  = 1
	VK_KHR_SHADER_DRAW_PARAMETERS_SPEC_VERSION                 uint32  = 1
	VK_KHR_SHADER_DRAW_PARAMETERS_EXTENSION_NAME               string  = "VK_KHR_shader_draw_parameters"
	VK_KHR_maintenance1                                        uint32  = 1
	VK_KHR_MAINTENANCE1_SPEC_VERSION                           uint32  = 2
	VK_KHR_MAINTENANCE1_EXTENSION_NAME                         string  = "VK_KHR_maintenance1"
	VK_KHR_device_group_creation                               uint32  = 1
	VK_KHR_DEVICE_GROUP_CREATION_SPEC_VERSION                  uint32  = 1
	VK_KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME                string  = "VK_KHR_device_group_creation"
	VK_MAX_DEVICE_GROUP_SIZE_KHR                                       = VK_MAX_DEVICE_GROUP_SIZE
	VK_KHR_external_memory_capabilities                        uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION           uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME         string  = "VK_KHR_external_memory_capabilities"
	VK_LUID_SIZE_KHR                                                   = VK_LUID_SIZE
	VK_KHR_external_memory                                     uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_SPEC_VERSION                        uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_EXTENSION_NAME                      string  = "VK_KHR_external_memory"
	VK_QUEUE_FAMILY_EXTERNAL_KHR                                       = VK_QUEUE_FAMILY_EXTERNAL
	VK_KHR_external_memory_fd                                  uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_FD_SPEC_VERSION                     uint32  = 1
	VK_KHR_EXTERNAL_MEMORY_FD_EXTENSION_NAME                   string  = "VK_KHR_external_memory_fd"
	VK_KHR_external_semaphore_capabilities                     uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_SPEC_VERSION        uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME      string  = "VK_KHR_external_semaphore_capabilities"
	VK_KHR_external_semaphore                                  uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_SPEC_VERSION                     uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME                   string  = "VK_KHR_external_semaphore"
	VK_KHR_external_semaphore_fd                               uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_FD_SPEC_VERSION                  uint32  = 1
	VK_KHR_EXTERNAL_SEMAPHORE_FD_EXTENSION_NAME                string  = "VK_KHR_external_semaphore_fd"
	VK_KHR_push_descriptor                                     uint32  = 1
	VK_KHR_PUSH_DESCRIPTOR_SPEC_VERSION                        uint32  = 2
	VK_KHR_PUSH_DESCRIPTOR_EXTENSION_NAME                      string  = "VK_KHR_push_descriptor"
	VK_KHR_shader_float16_int8                                 uint32  = 1
	VK_KHR_SHADER_FLOAT16_INT8_SPEC_VERSION                    uint32  = 1
	VK_KHR_SHADER_FLOAT16_INT8_EXTENSION_NAME                  string  = "VK_KHR_shader_float16_int8"
	VK_KHR_16bit_storage                                       uint32  = 1
	VK_KHR_16BIT_STORAGE_SPEC_VERSION                          uint32  = 1
	VK_KHR_16BIT_STORAGE_EXTENSION_NAME                        string  = "VK_KHR_16bit_storage"
	VK_KHR_incremental_present                                 uint32  = 1
	VK_KHR_INCREMENTAL_PRESENT_SPEC_VERSION                    uint32  = 2
	VK_KHR_INCREMENTAL_PRESENT_EXTENSION_NAME                  string  = "VK_KHR_incremental_present"
	VK_KHR_descriptor_update_template                          uint32  = 1
	VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_SPEC_VERSION             uint32  = 1
	VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME           string  = "VK_KHR_descriptor_update_template"
	VK_KHR_imageless_framebuffer                               uint32  = 1
	VK_KHR_IMAGELESS_FRAMEBUFFER_SPEC_VERSION                  uint32  = 1
	VK_KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME                string  = "VK_KHR_imageless_framebuffer"
	VK_KHR_create_renderpass2                                  uint32  = 1
	VK_KHR_CREATE_RENDERPASS_2_SPEC_VERSION                    uint32  = 1
	VK_KHR_CREATE_RENDERPASS_2_EXTENSION_NAME                  string  = "VK_KHR_create_renderpass2"
	VK_KHR_shared_presentable_image                            uint32  = 1
	VK_KHR_SHARED_PRESENTABLE_IMAGE_SPEC_VERSION               uint32  = 1
	VK_KHR_SHARED_PRESENTABLE_IMAGE_EXTENSION_NAME             string  = "VK_KHR_shared_presentable_image"
	VK_KHR_external_fence_capabilities                         uint32  = 1
	VK_KHR_EXTERNAL_FENCE_CAPABILITIES_SPEC_VERSION            uint32  = 1
	VK_KHR_EXTERNAL_FENCE_CAPABILITIES_EXTENSION_NAME          string  = "VK_KHR_external_fence_capabilities"
	VK_KHR_external_fence                                      uint32  = 1
	VK_KHR_EXTERNAL_FENCE_SPEC_VERSION                         uint32  = 1
	VK_KHR_EXTERNAL_FENCE_EXTENSION_NAME                       string  = "VK_KHR_external_fence"
	VK_KHR_external_fence_fd                                   uint32  = 1
	VK_KHR_EXTERNAL_FENCE_FD_SPEC_VERSION                      uint32  = 1
	VK_KHR_EXTERNAL_FENCE_FD_EXTENSION_NAME                    string  = "VK_KHR_external_fence_fd"
	VK_KHR_performance_query                                   uint32  = 1
	VK_KHR_PERFORMANCE_QUERY_SPEC_VERSION                      uint32  = 1
	VK_KHR_PERFORMANCE_QUERY_EXTENSION_NAME                    string  = "VK_KHR_performance_query"
	VK_KHR_maintenance2                                        uint32  = 1
	VK_KHR_MAINTENANCE2_SPEC_VERSION                           uint32  = 1
	VK_KHR_MAINTENANCE2_EXTENSION_NAME                         string  = "VK_KHR_maintenance2"
	VK_KHR_get_surface_capabilities2                           uint32  = 1
	VK_KHR_GET_SURFACE_CAPABILITIES_2_SPEC_VERSION             uint32  = 1
	VK_KHR_GET_SURFACE_CAPABILITIES_2_EXTENSION_NAME           string  = "VK_KHR_get_surface_capabilities2"
	VK_KHR_variable_pointers                                   uint32  = 1
	VK_KHR_VARIABLE_POINTERS_SPEC_VERSION                      uint32  = 1
	VK_KHR_VARIABLE_POINTERS_EXTENSION_NAME                    string  = "VK_KHR_variable_pointers"
	VK_KHR_get_display_properties2                             uint32  = 1
	VK_KHR_GET_DISPLAY_PROPERTIES_2_SPEC_VERSION               uint32  = 1
	VK_KHR_GET_DISPLAY_PROPERTIES_2_EXTENSION_NAME             string  = "VK_KHR_get_display_properties2"
	VK_KHR_dedicated_allocation                                uint32  = 1
	VK_KHR_DEDICATED_ALLOCATION_SPEC_VERSION                   uint32  = 3
	VK_KHR_DEDICATED_ALLOCATION_EXTENSION_NAME                 string  = "VK_KHR_dedicated_allocation"
	VK_KHR_storage_buffer_storage_class                        uint32  = 1
	VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_SPEC_VERSION           uint32  = 1
	VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME         string  = "VK_KHR_storage_buffer_storage_class"
	VK_KHR_relaxed_block_layout                                uint32  = 1
	VK_KHR_RELAXED_BLOCK_LAYOUT_SPEC_VERSION                   uint32  = 1
	VK_KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME                 string  = "VK_KHR_relaxed_block_layout"
	VK_KHR_get_memory_requirements2                            uint32  = 1
	VK_KHR_GET_MEMORY_REQUIREMENTS_2_SPEC_VERSION              uint32  = 1
	VK_KHR_GET_MEMORY_REQUIREMENTS_2_EXTENSION_NAME            string  = "VK_KHR_get_memory_requirements2"
	VK_KHR_image_format_list                                   uint32  = 1
	VK_KHR_IMAGE_FORMAT_LIST_SPEC_VERSION                      uint32  = 1
	VK_KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME                    string  = "VK_KHR_image_format_list"
	VK_KHR_sampler_ycbcr_conversion                            uint32  = 1
	VK_KHR_SAMPLER_YCBCR_CONVERSION_SPEC_VERSION               uint32  = 14
	VK_KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME             string  = "VK_KHR_sampler_ycbcr_conversion"
	VK_KHR_bind_memory2                                        uint32  = 1
	VK_KHR_BIND_MEMORY_2_SPEC_VERSION                          uint32  = 1
	VK_KHR_BIND_MEMORY_2_EXTENSION_NAME                        string  = "VK_KHR_bind_memory2"
	VK_KHR_maintenance3                                        uint32  = 1
	VK_KHR_MAINTENANCE3_SPEC_VERSION                           uint32  = 1
	VK_KHR_MAINTENANCE3_EXTENSION_NAME                         string  = "VK_KHR_maintenance3"
	VK_KHR_draw_indirect_count                                 uint32  = 1
	VK_KHR_DRAW_INDIRECT_COUNT_SPEC_VERSION                    uint32  = 1
	VK_KHR_DRAW_INDIRECT_COUNT_EXTENSION_NAME                  string  = "VK_KHR_draw_indirect_count"
	VK_KHR_shader_subgroup_extended_types                      uint32  = 1
	VK_KHR_SHADER_SUBGROUP_EXTENDED_TYPES_SPEC_VERSION         uint32  = 1
	VK_KHR_SHADER_SUBGROUP_EXTENDED_TYPES_EXTENSION_NAME       string  = "VK_KHR_shader_subgroup_extended_types"
	VK_KHR_8bit_storage                                        uint32  = 1
	VK_KHR_8BIT_STORAGE_SPEC_VERSION                           uint32  = 1
	VK_KHR_8BIT_STORAGE_EXTENSION_NAME                         string  = "VK_KHR_8bit_storage"
	VK_KHR_shader_atomic_int64                                 uint32  = 1
	VK_KHR_SHADER_ATOMIC_INT64_SPEC_VERSION                    uint32  = 1
	VK_KHR_SHADER_ATOMIC_INT64_EXTENSION_NAME                  string  = "VK_KHR_shader_atomic_int64"
	VK_KHR_shader_clock                                        uint32  = 1
	VK_KHR_SHADER_CLOCK_SPEC_VERSION                           uint32  = 1
	VK_KHR_SHADER_CLOCK_EXTENSION_NAME                         string  = "VK_KHR_shader_clock"
	VK_KHR_driver_properties                                   uint32  = 1
	VK_KHR_DRIVER_PROPERTIES_SPEC_VERSION                      uint32  = 1
	VK_KHR_DRIVER_PROPERTIES_EXTENSION_NAME                    string  = "VK_KHR_driver_properties"
	VK_MAX_DRIVER_NAME_SIZE_KHR                                        = VK_MAX_DRIVER_NAME_SIZE
	VK_MAX_DRIVER_INFO_SIZE_KHR                                        = VK_MAX_DRIVER_INFO_SIZE
	VK_KHR_shader_float_controls                               uint32  = 1
	VK_KHR_SHADER_FLOAT_CONTROLS_SPEC_VERSION                  uint32  = 4
	VK_KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME                string  = "VK_KHR_shader_float_controls"
	VK_KHR_depth_stencil_resolve                               uint32  = 1
	VK_KHR_DEPTH_STENCIL_RESOLVE_SPEC_VERSION                  uint32  = 1
	VK_KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME                string  = "VK_KHR_depth_stencil_resolve"
	VK_KHR_swapchain_mutable_format                            uint32  = 1
	VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_SPEC_VERSION               uint32  = 1
	VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_EXTENSION_NAME             string  = "VK_KHR_swapchain_mutable_format"
	VK_KHR_timeline_semaphore                                  uint32  = 1
	VK_KHR_TIMELINE_SEMAPHORE_SPEC_VERSION                     uint32  = 2
	VK_KHR_TIMELINE_SEMAPHORE_EXTENSION_NAME                   string  = "VK_KHR_timeline_semaphore"
	VK_KHR_vulkan_memory_model                                 uint32  = 1
	VK_KHR_VULKAN_MEMORY_MODEL_SPEC_VERSION                    uint32  = 3
	VK_KHR_VULKAN_MEMORY_MODEL_EXTENSION_NAME                  string  = "VK_KHR_vulkan_memory_model"
	VK_KHR_shader_terminate_invocation                         uint32  = 1
	VK_KHR_SHADER_TERMINATE_INVOCATION_SPEC_VERSION            uint32  = 1
	VK_KHR_SHADER_TERMINATE_INVOCATION_EXTENSION_NAME          string  = "VK_KHR_shader_terminate_invocation"
	VK_KHR_fragment_shading_rate                               uint32  = 1
	VK_KHR_FRAGMENT_SHADING_RATE_SPEC_VERSION                  uint32  = 1
	VK_KHR_FRAGMENT_SHADING_RATE_EXTENSION_NAME                string  = "VK_KHR_fragment_shading_rate"
	VK_KHR_spirv_1_4                                           uint32  = 1
	VK_KHR_SPIRV_1_4_SPEC_VERSION                              uint32  = 1
	VK_KHR_SPIRV_1_4_EXTENSION_NAME                            string  = "VK_KHR_spirv_1_4"
	VK_KHR_surface_protected_capabilities                      uint32  = 1
	VK_KHR_SURFACE_PROTECTED_CAPABILITIES_SPEC_VERSION         uint32  = 1
	VK_KHR_SURFACE_PROTECTED_CAPABILITIES_EXTENSION_NAME       string  = "VK_KHR_surface_protected_capabilities"
	VK_KHR_separate_depth_stencil_layouts                      uint32  = 1
	VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_SPEC_VERSION         uint32  = 1
	VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME       string  = "VK_KHR_separate_depth_stencil_layouts"
	VK_KHR_present_wait                                        uint32  = 1
	VK_KHR_PRESENT_WAIT_SPEC_VERSION                           uint32  = 1
	VK_KHR_PRESENT_WAIT_EXTENSION_NAME                         string  = "VK_KHR_present_wait"
	VK_KHR_uniform_buffer_standard_layout                      uint32  = 1
	VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_SPEC_VERSION         uint32  = 1
	VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME       string  = "VK_KHR_uniform_buffer_standard_layout"
	VK_KHR_buffer_device_address                               uint32  = 1
	VK_KHR_BUFFER_DEVICE_ADDRESS_SPEC_VERSION                  uint32  = 1
	VK_KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME                string  = "VK_KHR_buffer_device_address"
	VK_KHR_deferred_host_operations                            uint32  = 1
	VK_KHR_DEFERRED_HOST_OPERATIONS_SPEC_VERSION               uint32  = 4
	VK_KHR_DEFERRED_HOST_OPERATIONS_EXTENSION_NAME             string  = "VK_KHR_deferred_host_operations"
	VK_KHR_pipeline_executable_properties                      uint32  = 1
	VK_KHR_PIPELINE_EXECUTABLE_PROPERTIES_SPEC_VERSION         uint32  = 1
	VK_KHR_PIPELINE_EXECUTABLE_PROPERTIES_EXTENSION_NAME       string  = "VK_KHR_pipeline_executable_properties"
	VK_KHR_pipeline_library                                    uint32  = 1
	VK_KHR_PIPELINE_LIBRARY_SPEC_VERSION                       uint32  = 1
	VK_KHR_PIPELINE_LIBRARY_EXTENSION_NAME                     string  = "VK_KHR_pipeline_library"
	VK_KHR_shader_non_semantic_info                            uint32  = 1
	VK_KHR_SHADER_NON_SEMANTIC_INFO_SPEC_VERSION               uint32  = 1
	VK_KHR_SHADER_NON_SEMANTIC_INFO_EXTENSION_NAME             string  = "VK_KHR_shader_non_semantic_info"
	VK_KHR_present_id                                          uint32  = 1
	VK_KHR_PRESENT_ID_SPEC_VERSION                             uint32  = 1
	VK_KHR_PRESENT_ID_EXTENSION_NAME                           string  = "VK_KHR_present_id"
	VK_KHR_synchronization2                                    uint32  = 1
	VK_KHR_SYNCHRONIZATION_2_SPEC_VERSION                      uint32  = 1
	VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME                    string  = "VK_KHR_synchronization2"
	VK_KHR_shader_subgroup_uniform_control_flow                uint32  = 1
	VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_SPEC_VERSION   uint32  = 1
	VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_EXTENSION_NAME string  = "VK_KHR_shader_subgroup_uniform_control_flow"
	VK_KHR_zero_initialize_workgroup_memory                    uint32  = 1
	VK_KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_SPEC_VERSION       uint32  = 1
	VK_KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_EXTENSION_NAME     string  = "VK_KHR_zero_initialize_workgroup_memory"
	VK_KHR_workgroup_memory_explicit_layout                    uint32  = 1
	VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_SPEC_VERSION       uint32  = 1
	VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_EXTENSION_NAME     string  = "VK_KHR_workgroup_memory_explicit_layout"
	VK_KHR_copy_commands2                                      uint32  = 1
	VK_KHR_COPY_COMMANDS_2_SPEC_VERSION                        uint32  = 1
	VK_KHR_COPY_COMMANDS_2_EXTENSION_NAME                      string  = "VK_KHR_copy_commands2"
	VK_EXT_debug_report                                        uint32  = 1
	VK_EXT_DEBUG_REPORT_SPEC_VERSION                           uint32  = 10
	VK_EXT_DEBUG_REPORT_EXTENSION_NAME                         string  = "VK_EXT_debug_report"
	VK_NV_glsl_shader                                          uint32  = 1
	VK_NV_GLSL_SHADER_SPEC_VERSION                             uint32  = 1
	VK_NV_GLSL_SHADER_EXTENSION_NAME                           string  = "VK_NV_glsl_shader"
	VK_EXT_depth_range_unrestricted                            uint32  = 1
	VK_EXT_DEPTH_RANGE_UNRESTRICTED_SPEC_VERSION               uint32  = 1
	VK_EXT_DEPTH_RANGE_UNRESTRICTED_EXTENSION_NAME             string  = "VK_EXT_depth_range_unrestricted"
	VK_IMG_filter_cubic                                        uint32  = 1
	VK_IMG_FILTER_CUBIC_SPEC_VERSION                           uint32  = 1
	VK_IMG_FILTER_CUBIC_EXTENSION_NAME                         string  = "VK_IMG_filter_cubic"
	VK_AMD_rasterization_order                                 uint32  = 1
	VK_AMD_RASTERIZATION_ORDER_SPEC_VERSION                    uint32  = 1
	VK_AMD_RASTERIZATION_ORDER_EXTENSION_NAME                  string  = "VK_AMD_rasterization_order"
	VK_AMD_shader_trinary_minmax                               uint32  = 1
	VK_AMD_SHADER_TRINARY_MINMAX_SPEC_VERSION                  uint32  = 1
	VK_AMD_SHADER_TRINARY_MINMAX_EXTENSION_NAME                string  = "VK_AMD_shader_trinary_minmax"
	VK_AMD_shader_explicit_vertex_parameter                    uint32  = 1
	VK_AMD_SHADER_EXPLICIT_VERTEX_PARAMETER_SPEC_VERSION       uint32  = 1
	VK_AMD_SHADER_EXPLICIT_VERTEX_PARAMETER_EXTENSION_NAME     string  = "VK_AMD_shader_explicit_vertex_parameter"
	VK_EXT_debug_marker                                        uint32  = 1
	VK_EXT_DEBUG_MARKER_SPEC_VERSION                           uint32  = 4
	VK_EXT_DEBUG_MARKER_EXTENSION_NAME                         string  = "VK_EXT_debug_marker"
	VK_AMD_gcn_shader                                          uint32  = 1
	VK_AMD_GCN_SHADER_SPEC_VERSION                             uint32  = 1
	VK_AMD_GCN_SHADER_EXTENSION_NAME                           string  = "VK_AMD_gcn_shader"
	VK_NV_dedicated_allocation                                 uint32  = 1
	VK_NV_DEDICATED_ALLOCATION_SPEC_VERSION                    uint32  = 1
	VK_NV_DEDICATED_ALLOCATION_EXTENSION_NAME                  string  = "VK_NV_dedicated_allocation"
	VK_EXT_transform_feedback                                  uint32  = 1
	VK_EXT_TRANSFORM_FEEDBACK_SPEC_VERSION                     uint32  = 1
	VK_EXT_TRANSFORM_FEEDBACK_EXTENSION_NAME                   string  = "VK_EXT_transform_feedback"
	VK_NVX_binary_import                                       uint32  = 1
	VK_NVX_BINARY_IMPORT_SPEC_VERSION                          uint32  = 1
	VK_NVX_BINARY_IMPORT_EXTENSION_NAME                        string  = "VK_NVX_binary_import"
	VK_NVX_image_view_handle                                   uint32  = 1
	VK_NVX_IMAGE_VIEW_HANDLE_SPEC_VERSION                      uint32  = 2
	VK_NVX_IMAGE_VIEW_HANDLE_EXTENSION_NAME                    string  = "VK_NVX_image_view_handle"
	VK_AMD_draw_indirect_count                                 uint32  = 1
	VK_AMD_DRAW_INDIRECT_COUNT_SPEC_VERSION                    uint32  = 2
	VK_AMD_DRAW_INDIRECT_COUNT_EXTENSION_NAME                  string  = "VK_AMD_draw_indirect_count"
	VK_AMD_negative_viewport_height                            uint32  = 1
	VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_SPEC_VERSION               uint32  = 1
	VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_EXTENSION_NAME             string  = "VK_AMD_negative_viewport_height"
	VK_AMD_gpu_shader_half_float                               uint32  = 1
	VK_AMD_GPU_SHADER_HALF_FLOAT_SPEC_VERSION                  uint32  = 2
	VK_AMD_GPU_SHADER_HALF_FLOAT_EXTENSION_NAME                string  = "VK_AMD_gpu_shader_half_float"
	VK_AMD_shader_ballot                                       uint32  = 1
	VK_AMD_SHADER_BALLOT_SPEC_VERSION                          uint32  = 1
	VK_AMD_SHADER_BALLOT_EXTENSION_NAME                        string  = "VK_AMD_shader_ballot"
	VK_AMD_texture_gather_bias_lod                             uint32  = 1
	VK_AMD_TEXTURE_GATHER_BIAS_LOD_SPEC_VERSION                uint32  = 1
	VK_AMD_TEXTURE_GATHER_BIAS_LOD_EXTENSION_NAME              string  = "VK_AMD_texture_gather_bias_lod"
	VK_AMD_shader_info                                         uint32  = 1
	VK_AMD_SHADER_INFO_SPEC_VERSION                            uint32  = 1
	VK_AMD_SHADER_INFO_EXTENSION_NAME                          string  = "VK_AMD_shader_info"
	VK_AMD_shader_image_load_store_lod                         uint32  = 1
	VK_AMD_SHADER_IMAGE_LOAD_STORE_LOD_SPEC_VERSION            uint32  = 1
	VK_AMD_SHADER_IMAGE_LOAD_STORE_LOD_EXTENSION_NAME          string  = "VK_AMD_shader_image_load_store_lod"
	VK_NV_corner_sampled_image                                 uint32  = 1
	VK_NV_CORNER_SAMPLED_IMAGE_SPEC_VERSION                    uint32  = 2
	VK_NV_CORNER_SAMPLED_IMAGE_EXTENSION_NAME                  string  = "VK_NV_corner_sampled_image"
	VK_IMG_format_pvrtc                                        uint32  = 1
	VK_IMG_FORMAT_PVRTC_SPEC_VERSION                           uint32  = 1
	VK_IMG_FORMAT_PVRTC_EXTENSION_NAME                         string  = "VK_IMG_format_pvrtc"
	VK_NV_external_memory_capabilities                         uint32  = 1
	VK_NV_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION            uint32  = 1
	VK_NV_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME          string  = "VK_NV_external_memory_capabilities"
	VK_NV_external_memory                                      uint32  = 1
	VK_NV_EXTERNAL_MEMORY_SPEC_VERSION                         uint32  = 1
	VK_NV_EXTERNAL_MEMORY_EXTENSION_NAME                       string  = "VK_NV_external_memory"
	VK_EXT_validation_flags                                    uint32  = 1
	VK_EXT_VALIDATION_FLAGS_SPEC_VERSION                       uint32  = 2
	VK_EXT_VALIDATION_FLAGS_EXTENSION_NAME                     string  = "VK_EXT_validation_flags"
	VK_EXT_shader_subgroup_ballot                              uint32  = 1
	VK_EXT_SHADER_SUBGROUP_BALLOT_SPEC_VERSION                 uint32  = 1
	VK_EXT_SHADER_SUBGROUP_BALLOT_EXTENSION_NAME               string  = "VK_EXT_shader_subgroup_ballot"
	VK_EXT_shader_subgroup_vote                                uint32  = 1
	VK_EXT_SHADER_SUBGROUP_VOTE_SPEC_VERSION                   uint32  = 1
	VK_EXT_SHADER_SUBGROUP_VOTE_EXTENSION_NAME                 string  = "VK_EXT_shader_subgroup_vote"
	VK_EXT_texture_compression_astc_hdr                        uint32  = 1
	VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_SPEC_VERSION           uint32  = 1
	VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_EXTENSION_NAME         string  = "VK_EXT_texture_compression_astc_hdr"
	VK_EXT_astc_decode_mode                                    uint32  = 1
	VK_EXT_ASTC_DECODE_MODE_SPEC_VERSION                       uint32  = 1
	VK_EXT_ASTC_DECODE_MODE_EXTENSION_NAME                     string  = "VK_EXT_astc_decode_mode"
	VK_EXT_conditional_rendering                               uint32  = 1
	VK_EXT_CONDITIONAL_RENDERING_SPEC_VERSION                  uint32  = 2
	VK_EXT_CONDITIONAL_RENDERING_EXTENSION_NAME                string  = "VK_EXT_conditional_rendering"
	VK_NV_clip_space_w_scaling                                 uint32  = 1
	VK_NV_CLIP_SPACE_W_SCALING_SPEC_VERSION                    uint32  = 1
	VK_NV_CLIP_SPACE_W_SCALING_EXTENSION_NAME                  string  = "VK_NV_clip_space_w_scaling"
	VK_EXT_direct_mode_display                                 uint32  = 1
	VK_EXT_DIRECT_MODE_DISPLAY_SPEC_VERSION                    uint32  = 1
	VK_EXT_DIRECT_MODE_DISPLAY_EXTENSION_NAME                  string  = "VK_EXT_direct_mode_display"
	VK_EXT_display_surface_counter                             uint32  = 1
	VK_EXT_DISPLAY_SURFACE_COUNTER_SPEC_VERSION                uint32  = 1
	VK_EXT_DISPLAY_SURFACE_COUNTER_EXTENSION_NAME              string  = "VK_EXT_display_surface_counter"
	VK_EXT_display_control                                     uint32  = 1
	VK_EXT_DISPLAY_CONTROL_SPEC_VERSION                        uint32  = 1
	VK_EXT_DISPLAY_CONTROL_EXTENSION_NAME                      string  = "VK_EXT_display_control"
	VK_GOOGLE_display_timing                                   uint32  = 1
	VK_GOOGLE_DISPLAY_TIMING_SPEC_VERSION                      uint32  = 1
	VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME                    string  = "VK_GOOGLE_display_timing"
	VK_NV_sample_mask_override_coverage                        uint32  = 1
	VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_SPEC_VERSION           uint32  = 1
	VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_EXTENSION_NAME         string  = "VK_NV_sample_mask_override_coverage"
	VK_NV_geometry_shader_passthrough                          uint32  = 1
	VK_NV_GEOMETRY_SHADER_PASSTHROUGH_SPEC_VERSION             uint32  = 1
	VK_NV_GEOMETRY_SHADER_PASSTHROUGH_EXTENSION_NAME           string  = "VK_NV_geometry_shader_passthrough"
	VK_NV_viewport_array2                                      uint32  = 1
	VK_NV_VIEWPORT_ARRAY2_SPEC_VERSION                         uint32  = 1
	VK_NV_VIEWPORT_ARRAY2_EXTENSION_NAME                       string  = "VK_NV_viewport_array2"
	VK_NVX_multiview_per_view_attributes                       uint32  = 1
	VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_SPEC_VERSION          uint32  = 1
	VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_EXTENSION_NAME        string  = "VK_NVX_multiview_per_view_attributes"
	VK_NV_viewport_swizzle                                     uint32  = 1
	VK_NV_VIEWPORT_SWIZZLE_SPEC_VERSION                        uint32  = 1
	VK_NV_VIEWPORT_SWIZZLE_EXTENSION_NAME                      string  = "VK_NV_viewport_swizzle"
	VK_EXT_discard_rectangles                                  uint32  = 1
	VK_EXT_DISCARD_RECTANGLES_SPEC_VERSION                     uint32  = 1
	VK_EXT_DISCARD_RECTANGLES_EXTENSION_NAME                   string  = "VK_EXT_discard_rectangles"
	VK_EXT_conservative_rasterization                          uint32  = 1
	VK_EXT_CONSERVATIVE_RASTERIZATION_SPEC_VERSION             uint32  = 1
	VK_EXT_CONSERVATIVE_RASTERIZATION_EXTENSION_NAME           string  = "VK_EXT_conservative_rasterization"
	VK_EXT_depth_clip_enable                                   uint32  = 1
	VK_EXT_DEPTH_CLIP_ENABLE_SPEC_VERSION                      uint32  = 1
	VK_EXT_DEPTH_CLIP_ENABLE_EXTENSION_NAME                    string  = "VK_EXT_depth_clip_enable"
	VK_EXT_swapchain_colorspace                                uint32  = 1
	VK_EXT_SWAPCHAIN_COLOR_SPACE_SPEC_VERSION                  uint32  = 4
	VK_EXT_SWAPCHAIN_COLOR_SPACE_EXTENSION_NAME                string  = "VK_EXT_swapchain_colorspace"
	VK_EXT_hdr_metadata                                        uint32  = 1
	VK_EXT_HDR_METADATA_SPEC_VERSION                           uint32  = 2
	VK_EXT_HDR_METADATA_EXTENSION_NAME                         string  = "VK_EXT_hdr_metadata"
	VK_EXT_external_memory_dma_buf                             uint32  = 1
	VK_EXT_EXTERNAL_MEMORY_DMA_BUF_SPEC_VERSION                uint32  = 1
	VK_EXT_EXTERNAL_MEMORY_DMA_BUF_EXTENSION_NAME              string  = "VK_EXT_external_memory_dma_buf"
	VK_EXT_queue_family_foreign                                uint32  = 1
	VK_EXT_QUEUE_FAMILY_FOREIGN_SPEC_VERSION                   uint32  = 1
	VK_EXT_QUEUE_FAMILY_FOREIGN_EXTENSION_NAME                 string  = "VK_EXT_queue_family_foreign"
	VK_QUEUE_FAMILY_FOREIGN_EXT                                uint32  = 4294967293
	VK_EXT_debug_utils                                         uint32  = 1
	VK_EXT_DEBUG_UTILS_SPEC_VERSION                            uint32  = 2
	VK_EXT_DEBUG_UTILS_EXTENSION_NAME                          string  = "VK_EXT_debug_utils"
	VK_EXT_sampler_filter_minmax                               uint32  = 1
	VK_EXT_SAMPLER_FILTER_MINMAX_SPEC_VERSION                  uint32  = 2
	VK_EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME                string  = "VK_EXT_sampler_filter_minmax"
	VK_AMD_gpu_shader_int16                                    uint32  = 1
	VK_AMD_GPU_SHADER_INT16_SPEC_VERSION                       uint32  = 2
	VK_AMD_GPU_SHADER_INT16_EXTENSION_NAME                     string  = "VK_AMD_gpu_shader_int16"
	VK_AMD_mixed_attachment_samples                            uint32  = 1
	VK_AMD_MIXED_ATTACHMENT_SAMPLES_SPEC_VERSION               uint32  = 1
	VK_AMD_MIXED_ATTACHMENT_SAMPLES_EXTENSION_NAME             string  = "VK_AMD_mixed_attachment_samples"
	VK_AMD_shader_fragment_mask                                uint32  = 1
	VK_AMD_SHADER_FRAGMENT_MASK_SPEC_VERSION                   uint32  = 1
	VK_AMD_SHADER_FRAGMENT_MASK_EXTENSION_NAME                 string  = "VK_AMD_shader_fragment_mask"
	VK_EXT_inline_uniform_block                                uint32  = 1
	VK_EXT_INLINE_UNIFORM_BLOCK_SPEC_VERSION                   uint32  = 1
	VK_EXT_INLINE_UNIFORM_BLOCK_EXTENSION_NAME                 string  = "VK_EXT_inline_uniform_block"
	VK_EXT_shader_stencil_export                               uint32  = 1
	VK_EXT_SHADER_STENCIL_EXPORT_SPEC_VERSION                  uint32  = 1
	VK_EXT_SHADER_STENCIL_EXPORT_EXTENSION_NAME                string  = "VK_EXT_shader_stencil_export"
	VK_EXT_sample_locations                                    uint32  = 1
	VK_EXT_SAMPLE_LOCATIONS_SPEC_VERSION                       uint32  = 1
	VK_EXT_SAMPLE_LOCATIONS_EXTENSION_NAME                     string  = "VK_EXT_sample_locations"
	VK_EXT_blend_operation_advanced                            uint32  = 1
	VK_EXT_BLEND_OPERATION_ADVANCED_SPEC_VERSION               uint32  = 2
	VK_EXT_BLEND_OPERATION_ADVANCED_EXTENSION_NAME             string  = "VK_EXT_blend_operation_advanced"
	VK_NV_fragment_coverage_to_color                           uint32  = 1
	VK_NV_FRAGMENT_COVERAGE_TO_COLOR_SPEC_VERSION              uint32  = 1
	VK_NV_FRAGMENT_COVERAGE_TO_COLOR_EXTENSION_NAME            string  = "VK_NV_fragment_coverage_to_color"
	VK_NV_framebuffer_mixed_samples                            uint32  = 1
	VK_NV_FRAMEBUFFER_MIXED_SAMPLES_SPEC_VERSION               uint32  = 1
	VK_NV_FRAMEBUFFER_MIXED_SAMPLES_EXTENSION_NAME             string  = "VK_NV_framebuffer_mixed_samples"
	VK_NV_fill_rectangle                                       uint32  = 1
	VK_NV_FILL_RECTANGLE_SPEC_VERSION                          uint32  = 1
	VK_NV_FILL_RECTANGLE_EXTENSION_NAME                        string  = "VK_NV_fill_rectangle"
	VK_NV_shader_sm_builtins                                   uint32  = 1
	VK_NV_SHADER_SM_BUILTINS_SPEC_VERSION                      uint32  = 1
	VK_NV_SHADER_SM_BUILTINS_EXTENSION_NAME                    string  = "VK_NV_shader_sm_builtins"
	VK_EXT_post_depth_coverage                                 uint32  = 1
	VK_EXT_POST_DEPTH_COVERAGE_SPEC_VERSION                    uint32  = 1
	VK_EXT_POST_DEPTH_COVERAGE_EXTENSION_NAME                  string  = "VK_EXT_post_depth_coverage"
	VK_EXT_image_drm_format_modifier                           uint32  = 1
	VK_EXT_IMAGE_DRM_FORMAT_MODIFIER_SPEC_VERSION              uint32  = 1
	VK_EXT_IMAGE_DRM_FORMAT_MODIFIER_EXTENSION_NAME            string  = "VK_EXT_image_drm_format_modifier"
	VK_EXT_validation_cache                                    uint32  = 1
	VK_EXT_VALIDATION_CACHE_SPEC_VERSION                       uint32  = 1
	VK_EXT_VALIDATION_CACHE_EXTENSION_NAME                     string  = "VK_EXT_validation_cache"
	VK_EXT_descriptor_indexing                                 uint32  = 1
	VK_EXT_DESCRIPTOR_INDEXING_SPEC_VERSION                    uint32  = 2
	VK_EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME                  string  = "VK_EXT_descriptor_indexing"
	VK_EXT_shader_viewport_index_layer                         uint32  = 1
	VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_SPEC_VERSION            uint32  = 1
	VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME          string  = "VK_EXT_shader_viewport_index_layer"
	VK_NV_shading_rate_image                                   uint32  = 1
	VK_NV_SHADING_RATE_IMAGE_SPEC_VERSION                      uint32  = 3
	VK_NV_SHADING_RATE_IMAGE_EXTENSION_NAME                    string  = "VK_NV_shading_rate_image"
	VK_NV_ray_tracing                                          uint32  = 1
	VK_NV_RAY_TRACING_SPEC_VERSION                             uint32  = 3
	VK_NV_RAY_TRACING_EXTENSION_NAME                           string  = "VK_NV_ray_tracing"
	VK_SHADER_UNUSED_KHR                                       uint32  = 0xFFFFFFFF
	VK_SHADER_UNUSED_NV                                                = VK_SHADER_UNUSED_KHR
	VK_NV_representative_fragment_test                         uint32  = 1
	VK_NV_REPRESENTATIVE_FRAGMENT_TEST_SPEC_VERSION            uint32  = 2
	VK_NV_REPRESENTATIVE_FRAGMENT_TEST_EXTENSION_NAME          string  = "VK_NV_representative_fragment_test"
	VK_EXT_filter_cubic                                        uint32  = 1
	VK_EXT_FILTER_CUBIC_SPEC_VERSION                           uint32  = 3
	VK_EXT_FILTER_CUBIC_EXTENSION_NAME                         string  = "VK_EXT_filter_cubic"
	VK_QCOM_render_pass_shader_resolve                         uint32  = 1
	VK_QCOM_RENDER_PASS_SHADER_RESOLVE_SPEC_VERSION            uint32  = 4
	VK_QCOM_RENDER_PASS_SHADER_RESOLVE_EXTENSION_NAME          string  = "VK_QCOM_render_pass_shader_resolve"
	VK_EXT_global_priority                                     uint32  = 1
	VK_EXT_GLOBAL_PRIORITY_SPEC_VERSION                        uint32  = 2
	VK_EXT_GLOBAL_PRIORITY_EXTENSION_NAME                      string  = "VK_EXT_global_priority"
	VK_EXT_external_memory_host                                uint32  = 1
	VK_EXT_EXTERNAL_MEMORY_HOST_SPEC_VERSION                   uint32  = 1
	VK_EXT_EXTERNAL_MEMORY_HOST_EXTENSION_NAME                 string  = "VK_EXT_external_memory_host"
	VK_AMD_buffer_marker                                       uint32  = 1
	VK_AMD_BUFFER_MARKER_SPEC_VERSION                          uint32  = 1
	VK_AMD_BUFFER_MARKER_EXTENSION_NAME                        string  = "VK_AMD_buffer_marker"
	VK_AMD_pipeline_compiler_control                           uint32  = 1
	VK_AMD_PIPELINE_COMPILER_CONTROL_SPEC_VERSION              uint32  = 1
	VK_AMD_PIPELINE_COMPILER_CONTROL_EXTENSION_NAME            string  = "VK_AMD_pipeline_compiler_control"
	VK_EXT_calibrated_timestamps                               uint32  = 1
	VK_EXT_CALIBRATED_TIMESTAMPS_SPEC_VERSION                  uint32  = 2
	VK_EXT_CALIBRATED_TIMESTAMPS_EXTENSION_NAME                string  = "VK_EXT_calibrated_timestamps"
	VK_AMD_shader_core_properties                              uint32  = 1
	VK_AMD_SHADER_CORE_PROPERTIES_SPEC_VERSION                 uint32  = 2
	VK_AMD_SHADER_CORE_PROPERTIES_EXTENSION_NAME               string  = "VK_AMD_shader_core_properties"
	VK_AMD_memory_overallocation_behavior                      uint32  = 1
	VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_SPEC_VERSION         uint32  = 1
	VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_EXTENSION_NAME       string  = "VK_AMD_memory_overallocation_behavior"
	VK_EXT_vertex_attribute_divisor                            uint32  = 1
	VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_SPEC_VERSION               uint32  = 3
	VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_EXTENSION_NAME             string  = "VK_EXT_vertex_attribute_divisor"
	VK_EXT_pipeline_creation_feedback                          uint32  = 1
	VK_EXT_PIPELINE_CREATION_FEEDBACK_SPEC_VERSION             uint32  = 1
	VK_EXT_PIPELINE_CREATION_FEEDBACK_EXTENSION_NAME           string  = "VK_EXT_pipeline_creation_feedback"
	VK_NV_shader_subgroup_partitioned                          uint32  = 1
	VK_NV_SHADER_SUBGROUP_PARTITIONED_SPEC_VERSION             uint32  = 1
	VK_NV_SHADER_SUBGROUP_PARTITIONED_EXTENSION_NAME           string  = "VK_NV_shader_subgroup_partitioned"
	VK_NV_compute_shader_derivatives                           uint32  = 1
	VK_NV_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION              uint32  = 1
	VK_NV_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME            string  = "VK_NV_compute_shader_derivatives"
	VK_NV_mesh_shader                                          uint32  = 1
	VK_NV_MESH_SHADER_SPEC_VERSION                             uint32  = 1
	VK_NV_MESH_SHADER_EXTENSION_NAME                           string  = "VK_NV_mesh_shader"
	VK_NV_fragment_shader_barycentric                          uint32  = 1
	VK_NV_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION             uint32  = 1
	VK_NV_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME           string  = "VK_NV_fragment_shader_barycentric"
	VK_NV_shader_image_footprint                               uint32  = 1
	VK_NV_SHADER_IMAGE_FOOTPRINT_SPEC_VERSION                  uint32  = 2
	VK_NV_SHADER_IMAGE_FOOTPRINT_EXTENSION_NAME                string  = "VK_NV_shader_image_footprint"
	VK_NV_scissor_exclusive                                    uint32  = 1
	VK_NV_SCISSOR_EXCLUSIVE_SPEC_VERSION                       uint32  = 1
	VK_NV_SCISSOR_EXCLUSIVE_EXTENSION_NAME                     string  = "VK_NV_scissor_exclusive"
	VK_NV_device_diagnostic_checkpoints                        uint32  = 1
	VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_SPEC_VERSION           uint32  = 2
	VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_EXTENSION_NAME         string  = "VK_NV_device_diagnostic_checkpoints"
	VK_INTEL_shader_integer_functions2                         uint32  = 1
	VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_SPEC_VERSION           uint32  = 1
	VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_EXTENSION_NAME         string  = "VK_INTEL_shader_integer_functions2"
	VK_INTEL_performance_query                                 uint32  = 1
	VK_INTEL_PERFORMANCE_QUERY_SPEC_VERSION                    uint32  = 2
	VK_INTEL_PERFORMANCE_QUERY_EXTENSION_NAME                  string  = "VK_INTEL_performance_query"
	VK_EXT_pci_bus_info                                        uint32  = 1
	VK_EXT_PCI_BUS_INFO_SPEC_VERSION                           uint32  = 2
	VK_EXT_PCI_BUS_INFO_EXTENSION_NAME                         string  = "VK_EXT_pci_bus_info"
	VK_AMD_display_native_hdr                                  uint32  = 1
	VK_AMD_DISPLAY_NATIVE_HDR_SPEC_VERSION                     uint32  = 1
	VK_AMD_DISPLAY_NATIVE_HDR_EXTENSION_NAME                   string  = "VK_AMD_display_native_hdr"
	VK_EXT_fragment_density_map                                uint32  = 1
	VK_EXT_FRAGMENT_DENSITY_MAP_SPEC_VERSION                   uint32  = 1
	VK_EXT_FRAGMENT_DENSITY_MAP_EXTENSION_NAME                 string  = "VK_EXT_fragment_density_map"
	VK_EXT_scalar_block_layout                                 uint32  = 1
	VK_EXT_SCALAR_BLOCK_LAYOUT_SPEC_VERSION                    uint32  = 1
	VK_EXT_SCALAR_BLOCK_LAYOUT_EXTENSION_NAME                  string  = "VK_EXT_scalar_block_layout"
	VK_GOOGLE_hlsl_functionality1                              uint32  = 1
	VK_GOOGLE_HLSL_FUNCTIONALITY1_SPEC_VERSION                 uint32  = 1
	VK_GOOGLE_HLSL_FUNCTIONALITY1_EXTENSION_NAME               string  = "VK_GOOGLE_hlsl_functionality1"
	VK_GOOGLE_decorate_string                                  uint32  = 1
	VK_GOOGLE_DECORATE_STRING_SPEC_VERSION                     uint32  = 1
	VK_GOOGLE_DECORATE_STRING_EXTENSION_NAME                   string  = "VK_GOOGLE_decorate_string"
	VK_EXT_subgroup_size_control                               uint32  = 1
	VK_EXT_SUBGROUP_SIZE_CONTROL_SPEC_VERSION                  uint32  = 2
	VK_EXT_SUBGROUP_SIZE_CONTROL_EXTENSION_NAME                string  = "VK_EXT_subgroup_size_control"
	VK_AMD_shader_core_properties2                             uint32  = 1
	VK_AMD_SHADER_CORE_PROPERTIES_2_SPEC_VERSION               uint32  = 1
	VK_AMD_SHADER_CORE_PROPERTIES_2_EXTENSION_NAME             string  = "VK_AMD_shader_core_properties2"
	VK_AMD_device_coherent_memory                              uint32  = 1
	VK_AMD_DEVICE_COHERENT_MEMORY_SPEC_VERSION                 uint32  = 1
	VK_AMD_DEVICE_COHERENT_MEMORY_EXTENSION_NAME               string  = "VK_AMD_device_coherent_memory"
	VK_EXT_shader_image_atomic_int64                           uint32  = 1
	VK_EXT_SHADER_IMAGE_ATOMIC_INT64_SPEC_VERSION              uint32  = 1
	VK_EXT_SHADER_IMAGE_ATOMIC_INT64_EXTENSION_NAME            string  = "VK_EXT_shader_image_atomic_int64"
	VK_EXT_memory_budget                                       uint32  = 1
	VK_EXT_MEMORY_BUDGET_SPEC_VERSION                          uint32  = 1
	VK_EXT_MEMORY_BUDGET_EXTENSION_NAME                        string  = "VK_EXT_memory_budget"
	VK_EXT_memory_priority                                     uint32  = 1
	VK_EXT_MEMORY_PRIORITY_SPEC_VERSION                        uint32  = 1
	VK_EXT_MEMORY_PRIORITY_EXTENSION_NAME                      string  = "VK_EXT_memory_priority"
	VK_NV_dedicated_allocation_image_aliasing                  uint32  = 1
	VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_SPEC_VERSION     uint32  = 1
	VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_EXTENSION_NAME   string  = "VK_NV_dedicated_allocation_image_aliasing"
	VK_EXT_buffer_device_address                               uint32  = 1
	VK_EXT_BUFFER_DEVICE_ADDRESS_SPEC_VERSION                  uint32  = 2
	VK_EXT_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME                string  = "VK_EXT_buffer_device_address"
	VK_EXT_tooling_info                                        uint32  = 1
	VK_EXT_TOOLING_INFO_SPEC_VERSION                           uint32  = 1
	VK_EXT_TOOLING_INFO_EXTENSION_NAME                         string  = "VK_EXT_tooling_info"
	VK_EXT_separate_stencil_usage                              uint32  = 1
	VK_EXT_SEPARATE_STENCIL_USAGE_SPEC_VERSION                 uint32  = 1
	VK_EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME               string  = "VK_EXT_separate_stencil_usage"
	VK_EXT_validation_features                                 uint32  = 1
	VK_EXT_VALIDATION_FEATURES_SPEC_VERSION                    uint32  = 5
	VK_EXT_VALIDATION_FEATURES_EXTENSION_NAME                  string  = "VK_EXT_validation_features"
	VK_NV_cooperative_matrix                                   uint32  = 1
	VK_NV_COOPERATIVE_MATRIX_SPEC_VERSION                      uint32  = 1
	VK_NV_COOPERATIVE_MATRIX_EXTENSION_NAME                    string  = "VK_NV_cooperative_matrix"
	VK_NV_coverage_reduction_mode                              uint32  = 1
	VK_NV_COVERAGE_REDUCTION_MODE_SPEC_VERSION                 uint32  = 1
	VK_NV_COVERAGE_REDUCTION_MODE_EXTENSION_NAME               string  = "VK_NV_coverage_reduction_mode"
	VK_EXT_fragment_shader_interlock                           uint32  = 1
	VK_EXT_FRAGMENT_SHADER_INTERLOCK_SPEC_VERSION              uint32  = 1
	VK_EXT_FRAGMENT_SHADER_INTERLOCK_EXTENSION_NAME            string  = "VK_EXT_fragment_shader_interlock"
	VK_EXT_ycbcr_image_arrays                                  uint32  = 1
	VK_EXT_YCBCR_IMAGE_ARRAYS_SPEC_VERSION                     uint32  = 1
	VK_EXT_YCBCR_IMAGE_ARRAYS_EXTENSION_NAME                   string  = "VK_EXT_ycbcr_image_arrays"
	VK_EXT_provoking_vertex                                    uint32  = 1
	VK_EXT_PROVOKING_VERTEX_SPEC_VERSION                       uint32  = 1
	VK_EXT_PROVOKING_VERTEX_EXTENSION_NAME                     string  = "VK_EXT_provoking_vertex"
	VK_EXT_headless_surface                                    uint32  = 1
	VK_EXT_HEADLESS_SURFACE_SPEC_VERSION                       uint32  = 1
	VK_EXT_HEADLESS_SURFACE_EXTENSION_NAME                     string  = "VK_EXT_headless_surface"
	VK_EXT_line_rasterization                                  uint32  = 1
	VK_EXT_LINE_RASTERIZATION_SPEC_VERSION                     uint32  = 1
	VK_EXT_LINE_RASTERIZATION_EXTENSION_NAME                   string  = "VK_EXT_line_rasterization"
	VK_EXT_shader_atomic_float                                 uint32  = 1
	VK_EXT_SHADER_ATOMIC_FLOAT_SPEC_VERSION                    uint32  = 1
	VK_EXT_SHADER_ATOMIC_FLOAT_EXTENSION_NAME                  string  = "VK_EXT_shader_atomic_float"
	VK_EXT_host_query_reset                                    uint32  = 1
	VK_EXT_HOST_QUERY_RESET_SPEC_VERSION                       uint32  = 1
	VK_EXT_HOST_QUERY_RESET_EXTENSION_NAME                     string  = "VK_EXT_host_query_reset"
	VK_EXT_index_type_uint8                                    uint32  = 1
	VK_EXT_INDEX_TYPE_UINT8_SPEC_VERSION                       uint32  = 1
	VK_EXT_INDEX_TYPE_UINT8_EXTENSION_NAME                     string  = "VK_EXT_index_type_uint8"
	VK_EXT_extended_dynamic_state                              uint32  = 1
	VK_EXT_EXTENDED_DYNAMIC_STATE_SPEC_VERSION                 uint32  = 1
	VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME               string  = "VK_EXT_extended_dynamic_state"
	VK_EXT_shader_atomic_float2                                uint32  = 1
	VK_EXT_SHADER_ATOMIC_FLOAT_2_SPEC_VERSION                  uint32  = 1
	VK_EXT_SHADER_ATOMIC_FLOAT_2_EXTENSION_NAME                string  = "VK_EXT_shader_atomic_float2"
	VK_EXT_shader_demote_to_helper_invocation                  uint32  = 1
	VK_EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_SPEC_VERSION     uint32  = 1
	VK_EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_EXTENSION_NAME   string  = "VK_EXT_shader_demote_to_helper_invocation"
	VK_NV_device_generated_commands                            uint32  = 1
	VK_NV_DEVICE_GENERATED_COMMANDS_SPEC_VERSION               uint32  = 3
	VK_NV_DEVICE_GENERATED_COMMANDS_EXTENSION_NAME             string  = "VK_NV_device_generated_commands"
	VK_NV_inherited_viewport_scissor                           uint32  = 1
	VK_NV_INHERITED_VIEWPORT_SCISSOR_SPEC_VERSION              uint32  = 1
	VK_NV_INHERITED_VIEWPORT_SCISSOR_EXTENSION_NAME            string  = "VK_NV_inherited_viewport_scissor"
	VK_EXT_texel_buffer_alignment                              uint32  = 1
	VK_EXT_TEXEL_BUFFER_ALIGNMENT_SPEC_VERSION                 uint32  = 1
	VK_EXT_TEXEL_BUFFER_ALIGNMENT_EXTENSION_NAME               string  = "VK_EXT_texel_buffer_alignment"
	VK_QCOM_render_pass_transform                              uint32  = 1
	VK_QCOM_RENDER_PASS_TRANSFORM_SPEC_VERSION                 uint32  = 2
	VK_QCOM_RENDER_PASS_TRANSFORM_EXTENSION_NAME               string  = "VK_QCOM_render_pass_transform"
	VK_EXT_device_memory_report                                uint32  = 1
	VK_EXT_DEVICE_MEMORY_REPORT_SPEC_VERSION                   uint32  = 2
	VK_EXT_DEVICE_MEMORY_REPORT_EXTENSION_NAME                 string  = "VK_EXT_device_memory_report"
	VK_EXT_acquire_drm_display                                 uint32  = 1
	VK_EXT_ACQUIRE_DRM_DISPLAY_SPEC_VERSION                    uint32  = 1
	VK_EXT_ACQUIRE_DRM_DISPLAY_EXTENSION_NAME                  string  = "VK_EXT_acquire_drm_display"
	VK_EXT_robustness2                                         uint32  = 1
	VK_EXT_ROBUSTNESS_2_SPEC_VERSION                           uint32  = 1
	VK_EXT_ROBUSTNESS_2_EXTENSION_NAME                         string  = "VK_EXT_robustness2"
	VK_EXT_custom_border_color                                 uint32  = 1
	VK_EXT_CUSTOM_BORDER_COLOR_SPEC_VERSION                    uint32  = 12
	VK_EXT_CUSTOM_BORDER_COLOR_EXTENSION_NAME                  string  = "VK_EXT_custom_border_color"
	VK_GOOGLE_user_type                                        uint32  = 1
	VK_GOOGLE_USER_TYPE_SPEC_VERSION                           uint32  = 1
	VK_GOOGLE_USER_TYPE_EXTENSION_NAME                         string  = "VK_GOOGLE_user_type"
	VK_EXT_private_data                                        uint32  = 1
	VK_EXT_PRIVATE_DATA_SPEC_VERSION                           uint32  = 1
	VK_EXT_PRIVATE_DATA_EXTENSION_NAME                         string  = "VK_EXT_private_data"
	VK_EXT_pipeline_creation_cache_control                     uint32  = 1
	VK_EXT_PIPELINE_CREATION_CACHE_CONTROL_SPEC_VERSION        uint32  = 3
	VK_EXT_PIPELINE_CREATION_CACHE_CONTROL_EXTENSION_NAME      string  = "VK_EXT_pipeline_creation_cache_control"
	VK_NV_device_diagnostics_config                            uint32  = 1
	VK_NV_DEVICE_DIAGNOSTICS_CONFIG_SPEC_VERSION               uint32  = 1
	VK_NV_DEVICE_DIAGNOSTICS_CONFIG_EXTENSION_NAME             string  = "VK_NV_device_diagnostics_config"
	VK_QCOM_render_pass_store_ops                              uint32  = 1
	VK_QCOM_RENDER_PASS_STORE_OPS_SPEC_VERSION                 uint32  = 2
	VK_QCOM_RENDER_PASS_STORE_OPS_EXTENSION_NAME               string  = "VK_QCOM_render_pass_store_ops"
	VK_NV_fragment_shading_rate_enums                          uint32  = 1
	VK_NV_FRAGMENT_SHADING_RATE_ENUMS_SPEC_VERSION             uint32  = 1
	VK_NV_FRAGMENT_SHADING_RATE_ENUMS_EXTENSION_NAME           string  = "VK_NV_fragment_shading_rate_enums"
	VK_NV_ray_tracing_motion_blur                              uint32  = 1
	VK_NV_RAY_TRACING_MOTION_BLUR_SPEC_VERSION                 uint32  = 1
	VK_NV_RAY_TRACING_MOTION_BLUR_EXTENSION_NAME               string  = "VK_NV_ray_tracing_motion_blur"
	VK_EXT_ycbcr_2plane_444_formats                            uint32  = 1
	VK_EXT_YCBCR_2PLANE_444_FORMATS_SPEC_VERSION               uint32  = 1
	VK_EXT_YCBCR_2PLANE_444_FORMATS_EXTENSION_NAME             string  = "VK_EXT_ycbcr_2plane_444_formats"
	VK_EXT_fragment_density_map2                               uint32  = 1
	VK_EXT_FRAGMENT_DENSITY_MAP_2_SPEC_VERSION                 uint32  = 1
	VK_EXT_FRAGMENT_DENSITY_MAP_2_EXTENSION_NAME               string  = "VK_EXT_fragment_density_map2"
	VK_QCOM_rotated_copy_commands                              uint32  = 1
	VK_QCOM_ROTATED_COPY_COMMANDS_SPEC_VERSION                 uint32  = 1
	VK_QCOM_ROTATED_COPY_COMMANDS_EXTENSION_NAME               string  = "VK_QCOM_rotated_copy_commands"
	VK_EXT_image_robustness                                    uint32  = 1
	VK_EXT_IMAGE_ROBUSTNESS_SPEC_VERSION                       uint32  = 1
	VK_EXT_IMAGE_ROBUSTNESS_EXTENSION_NAME                     string  = "VK_EXT_image_robustness"
	VK_EXT_4444_formats                                        uint32  = 1
	VK_EXT_4444_FORMATS_SPEC_VERSION                           uint32  = 1
	VK_EXT_4444_FORMATS_EXTENSION_NAME                         string  = "VK_EXT_4444_formats"
	VK_NV_acquire_winrt_display                                uint32  = 1
	VK_NV_ACQUIRE_WINRT_DISPLAY_SPEC_VERSION                   uint32  = 1
	VK_NV_ACQUIRE_WINRT_DISPLAY_EXTENSION_NAME                 string  = "VK_NV_acquire_winrt_display"
	VK_VALVE_mutable_descriptor_type                           uint32  = 1
	VK_VALVE_MUTABLE_DESCRIPTOR_TYPE_SPEC_VERSION              uint32  = 1
	VK_VALVE_MUTABLE_DESCRIPTOR_TYPE_EXTENSION_NAME            string  = "VK_VALVE_mutable_descriptor_type"
	VK_EXT_vertex_input_dynamic_state                          uint32  = 1
	VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_SPEC_VERSION             uint32  = 2
	VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_EXTENSION_NAME           string  = "VK_EXT_vertex_input_dynamic_state"
	VK_EXT_physical_device_drm                                 uint32  = 1
	VK_EXT_PHYSICAL_DEVICE_DRM_SPEC_VERSION                    uint32  = 1
	VK_EXT_PHYSICAL_DEVICE_DRM_EXTENSION_NAME                  string  = "VK_EXT_physical_device_drm"
	VK_HUAWEI_subpass_shading                                  uint32  = 1
	VK_HUAWEI_SUBPASS_SHADING_SPEC_VERSION                     uint32  = 2
	VK_HUAWEI_SUBPASS_SHADING_EXTENSION_NAME                   string  = "VK_HUAWEI_subpass_shading"
	VK_HUAWEI_invocation_mask                                  uint32  = 1
	VK_HUAWEI_INVOCATION_MASK_SPEC_VERSION                     uint32  = 1
	VK_HUAWEI_INVOCATION_MASK_EXTENSION_NAME                   string  = "VK_HUAWEI_invocation_mask"
	VK_NV_external_memory_rdma                                 uint32  = 1
	VK_NV_EXTERNAL_MEMORY_RDMA_SPEC_VERSION                    uint32  = 1
	VK_NV_EXTERNAL_MEMORY_RDMA_EXTENSION_NAME                  string  = "VK_NV_external_memory_rdma"
	VK_EXT_extended_dynamic_state2                             uint32  = 1
	VK_EXT_EXTENDED_DYNAMIC_STATE_2_SPEC_VERSION               uint32  = 1
	VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME             string  = "VK_EXT_extended_dynamic_state2"
	VK_EXT_color_write_enable                                  uint32  = 1
	VK_EXT_COLOR_WRITE_ENABLE_SPEC_VERSION                     uint32  = 1
	VK_EXT_COLOR_WRITE_ENABLE_EXTENSION_NAME                   string  = "VK_EXT_color_write_enable"
	VK_EXT_global_priority_query                               uint32  = 1
	VK_MAX_GLOBAL_PRIORITY_SIZE_EXT                            uint32  = 16
	VK_EXT_GLOBAL_PRIORITY_QUERY_SPEC_VERSION                  uint32  = 1
	VK_EXT_GLOBAL_PRIORITY_QUERY_EXTENSION_NAME                string  = "VK_EXT_global_priority_query"
	VK_EXT_multi_draw                                          uint32  = 1
	VK_EXT_MULTI_DRAW_SPEC_VERSION                             uint32  = 1
	VK_EXT_MULTI_DRAW_EXTENSION_NAME                           string  = "VK_EXT_multi_draw"
	VK_EXT_load_store_op_none                                  uint32  = 1
	VK_EXT_LOAD_STORE_OP_NONE_SPEC_VERSION                     uint32  = 1
	VK_EXT_LOAD_STORE_OP_NONE_EXTENSION_NAME                   string  = "VK_EXT_load_store_op_none"
	VK_KHR_acceleration_structure                              uint32  = 1
	VK_KHR_ACCELERATION_STRUCTURE_SPEC_VERSION                 uint32  = 12
	VK_KHR_ACCELERATION_STRUCTURE_EXTENSION_NAME               string  = "VK_KHR_acceleration_structure"
	VK_KHR_ray_tracing_pipeline                                uint32  = 1
	VK_KHR_RAY_TRACING_PIPELINE_SPEC_VERSION                   uint32  = 1
	VK_KHR_RAY_TRACING_PIPELINE_EXTENSION_NAME                 string  = "VK_KHR_ray_tracing_pipeline"
	VK_KHR_ray_query                                           uint32  = 1
	VK_KHR_RAY_QUERY_SPEC_VERSION                              uint32  = 1
	VK_KHR_RAY_QUERY_EXTENSION_NAME                            string  = "VK_KHR_ray_query"
)

type VkResult = int32

const (
	VK_SUCCESS                                            VkResult = 0
	VK_NOT_READY                                          VkResult = 1
	VK_TIMEOUT                                            VkResult = 2
	VK_EVENT_SET                                          VkResult = 3
	VK_EVENT_RESET                                        VkResult = 4
	VK_INCOMPLETE                                         VkResult = 5
	VK_ERROR_OUT_OF_HOST_MEMORY                           VkResult = -1
	VK_ERROR_OUT_OF_DEVICE_MEMORY                         VkResult = -2
	VK_ERROR_INITIALIZATION_FAILED                        VkResult = -3
	VK_ERROR_DEVICE_LOST                                  VkResult = -4
	VK_ERROR_MEMORY_MAP_FAILED                            VkResult = -5
	VK_ERROR_LAYER_NOT_PRESENT                            VkResult = -6
	VK_ERROR_EXTENSION_NOT_PRESENT                        VkResult = -7
	VK_ERROR_FEATURE_NOT_PRESENT                          VkResult = -8
	VK_ERROR_INCOMPATIBLE_DRIVER                          VkResult = -9
	VK_ERROR_TOO_MANY_OBJECTS                             VkResult = -10
	VK_ERROR_FORMAT_NOT_SUPPORTED                         VkResult = -11
	VK_ERROR_FRAGMENTED_POOL                              VkResult = -12
	VK_ERROR_UNKNOWN                                      VkResult = -13
	VK_ERROR_OUT_OF_POOL_MEMORY                           VkResult = -1000069000
	VK_ERROR_INVALID_EXTERNAL_HANDLE                      VkResult = -1000072003
	VK_ERROR_FRAGMENTATION                                VkResult = -1000161000
	VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS               VkResult = -1000257000
	VK_ERROR_SURFACE_LOST_KHR                             VkResult = -1000000000
	VK_ERROR_NATIVE_WINDOW_IN_USE_KHR                     VkResult = -1000000001
	VK_SUBOPTIMAL_KHR                                     VkResult = 1000001003
	VK_ERROR_OUT_OF_DATE_KHR                              VkResult = -1000001004
	VK_ERROR_INCOMPATIBLE_DISPLAY_KHR                     VkResult = -1000003001
	VK_ERROR_VALIDATION_FAILED_EXT                        VkResult = -1000011001
	VK_ERROR_INVALID_SHADER_NV                            VkResult = -1000012000
	VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT VkResult = -1000158000
	VK_ERROR_NOT_PERMITTED_EXT                            VkResult = -1000174001
	VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT          VkResult = -1000255000
	VK_THREAD_IDLE_KHR                                    VkResult = 1000268000
	VK_THREAD_DONE_KHR                                    VkResult = 1000268001
	VK_OPERATION_DEFERRED_KHR                             VkResult = 1000268002
	VK_OPERATION_NOT_DEFERRED_KHR                         VkResult = 1000268003
	VK_PIPELINE_COMPILE_REQUIRED_EXT                      VkResult = 1000297000
	VK_ERROR_OUT_OF_POOL_MEMORY_KHR                       VkResult = VK_ERROR_OUT_OF_POOL_MEMORY
	VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR                  VkResult = VK_ERROR_INVALID_EXTERNAL_HANDLE
	VK_ERROR_FRAGMENTATION_EXT                            VkResult = VK_ERROR_FRAGMENTATION
	VK_ERROR_INVALID_DEVICE_ADDRESS_EXT                   VkResult = VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS
	VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR           VkResult = VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS
	VK_ERROR_PIPELINE_COMPILE_REQUIRED_EXT                VkResult = VK_PIPELINE_COMPILE_REQUIRED_EXT
	VK_RESULT_MAX_ENUM                                    VkResult = 0x7FFFFFFF
)

type VkStructureType = uint32

const (
	VK_STRUCTURE_TYPE_APPLICATION_INFO                                                  VkStructureType = 0
	VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO                                              VkStructureType = 1
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO                                          VkStructureType = 2
	VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO                                                VkStructureType = 3
	VK_STRUCTURE_TYPE_SUBMIT_INFO                                                       VkStructureType = 4
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO                                              VkStructureType = 5
	VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE                                               VkStructureType = 6
	VK_STRUCTURE_TYPE_BIND_SPARSE_INFO                                                  VkStructureType = 7
	VK_STRUCTURE_TYPE_FENCE_CREATE_INFO                                                 VkStructureType = 8
	VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO                                             VkStructureType = 9
	VK_STRUCTURE_TYPE_EVENT_CREATE_INFO                                                 VkStructureType = 10
	VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO                                            VkStructureType = 11
	VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO                                                VkStructureType = 12
	VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO                                           VkStructureType = 13
	VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO                                                 VkStructureType = 14
	VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO                                            VkStructureType = 15
	VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO                                         VkStructureType = 16
	VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO                                        VkStructureType = 17
	VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO                                 VkStructureType = 18
	VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO                           VkStructureType = 19
	VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO                         VkStructureType = 20
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO                           VkStructureType = 21
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO                               VkStructureType = 22
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO                          VkStructureType = 23
	VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO                            VkStructureType = 24
	VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO                          VkStructureType = 25
	VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO                            VkStructureType = 26
	VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO                                VkStructureType = 27
	VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO                                     VkStructureType = 28
	VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO                                      VkStructureType = 29
	VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO                                       VkStructureType = 30
	VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO                                               VkStructureType = 31
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO                                 VkStructureType = 32
	VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO                                       VkStructureType = 33
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO                                      VkStructureType = 34
	VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET                                              VkStructureType = 35
	VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET                                               VkStructureType = 36
	VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO                                           VkStructureType = 37
	VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO                                           VkStructureType = 38
	VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO                                          VkStructureType = 39
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO                                      VkStructureType = 40
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO                                   VkStructureType = 41
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO                                         VkStructureType = 42
	VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO                                            VkStructureType = 43
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER                                             VkStructureType = 44
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER                                              VkStructureType = 45
	VK_STRUCTURE_TYPE_MEMORY_BARRIER                                                    VkStructureType = 46
	VK_STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO                                       VkStructureType = 47
	VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO                                         VkStructureType = 48
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES                               VkStructureType = 1000094000
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO                                           VkStructureType = 1000157000
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO                                            VkStructureType = 1000157001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES                            VkStructureType = 1000083000
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS                                     VkStructureType = 1000127000
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO                                    VkStructureType = 1000127001
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO                                        VkStructureType = 1000060000
	VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO                               VkStructureType = 1000060003
	VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO                            VkStructureType = 1000060004
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO                                          VkStructureType = 1000060005
	VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO                                     VkStructureType = 1000060006
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO                              VkStructureType = 1000060013
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO                               VkStructureType = 1000060014
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES                                  VkStructureType = 1000070000
	VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO                                   VkStructureType = 1000070001
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2                                 VkStructureType = 1000146000
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2                                  VkStructureType = 1000146001
	VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2                           VkStructureType = 1000146002
	VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2                                             VkStructureType = 1000146003
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2                                VkStructureType = 1000146004
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2                                        VkStructureType = 1000059000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2                                      VkStructureType = 1000059001
	VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2                                               VkStructureType = 1000059002
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2                                         VkStructureType = 1000059003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2                               VkStructureType = 1000059004
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2                                         VkStructureType = 1000059005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2                               VkStructureType = 1000059006
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2                                  VkStructureType = 1000059007
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2                        VkStructureType = 1000059008
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES                         VkStructureType = 1000117000
	VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO                   VkStructureType = 1000117001
	VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO                                      VkStructureType = 1000117002
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO             VkStructureType = 1000117003
	VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO                                 VkStructureType = 1000053000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES                                VkStructureType = 1000053001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES                              VkStructureType = 1000053002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES                        VkStructureType = 1000120000
	VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO                                             VkStructureType = 1000145000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES                         VkStructureType = 1000145001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES                       VkStructureType = 1000145002
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2                                               VkStructureType = 1000145003
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO                              VkStructureType = 1000156000
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO                                     VkStructureType = 1000156001
	VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO                                      VkStructureType = 1000156002
	VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO                              VkStructureType = 1000156003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES                 VkStructureType = 1000156004
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES                  VkStructureType = 1000156005
	VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO                            VkStructureType = 1000085000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO                        VkStructureType = 1000071000
	VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES                                  VkStructureType = 1000071001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO                              VkStructureType = 1000071002
	VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES                                        VkStructureType = 1000071003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES                                     VkStructureType = 1000071004
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO                                VkStructureType = 1000072000
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO                                 VkStructureType = 1000072001
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO                                       VkStructureType = 1000072002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO                               VkStructureType = 1000112000
	VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES                                         VkStructureType = 1000112001
	VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO                                          VkStructureType = 1000113000
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO                                      VkStructureType = 1000077000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO                           VkStructureType = 1000076000
	VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES                                     VkStructureType = 1000076001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES                          VkStructureType = 1000168000
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT                                     VkStructureType = 1000168001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES                   VkStructureType = 1000063000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES                               VkStructureType = 49
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES                             VkStructureType = 50
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES                               VkStructureType = 51
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES                             VkStructureType = 52
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO                                     VkStructureType = 1000147000
	VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2                                          VkStructureType = 1000109000
	VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2                                            VkStructureType = 1000109001
	VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2                                             VkStructureType = 1000109002
	VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2                                              VkStructureType = 1000109003
	VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2                                         VkStructureType = 1000109004
	VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO                                                VkStructureType = 1000109005
	VK_STRUCTURE_TYPE_SUBPASS_END_INFO                                                  VkStructureType = 1000109006
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES                             VkStructureType = 1000177000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES                                 VkStructureType = 1000196000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES                      VkStructureType = 1000180000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES                      VkStructureType = 1000082000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES                         VkStructureType = 1000197000
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO                   VkStructureType = 1000161000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES                      VkStructureType = 1000161001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES                    VkStructureType = 1000161002
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO            VkStructureType = 1000161003
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT           VkStructureType = 1000161004
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES                  VkStructureType = 1000199000
	VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE                         VkStructureType = 1000199001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES                      VkStructureType = 1000221000
	VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO                                   VkStructureType = 1000246000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES                  VkStructureType = 1000130000
	VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO                                VkStructureType = 1000130001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES                      VkStructureType = 1000211000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES                    VkStructureType = 1000108000
	VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO                               VkStructureType = 1000108001
	VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO                                 VkStructureType = 1000108002
	VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO                                 VkStructureType = 1000108003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES           VkStructureType = 1000253000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES           VkStructureType = 1000175000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES           VkStructureType = 1000241000
	VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT                               VkStructureType = 1000241001
	VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT                             VkStructureType = 1000241002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES                         VkStructureType = 1000261000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES                       VkStructureType = 1000207000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES                     VkStructureType = 1000207001
	VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO                                        VkStructureType = 1000207002
	VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO                                    VkStructureType = 1000207003
	VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO                                               VkStructureType = 1000207004
	VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO                                             VkStructureType = 1000207005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES                    VkStructureType = 1000257000
	VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO                                        VkStructureType = 1000244001
	VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO                         VkStructureType = 1000257002
	VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO                       VkStructureType = 1000257003
	VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO                         VkStructureType = 1000257004
	VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR                                         VkStructureType = 1000001000
	VK_STRUCTURE_TYPE_PRESENT_INFO_KHR                                                  VkStructureType = 1000001001
	VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR                             VkStructureType = 1000060007
	VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR                                   VkStructureType = 1000060008
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR                              VkStructureType = 1000060009
	VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR                                       VkStructureType = 1000060010
	VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR                                     VkStructureType = 1000060011
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR                            VkStructureType = 1000060012
	VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR                                      VkStructureType = 1000002000
	VK_STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR                                   VkStructureType = 1000002001
	VK_STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR                                          VkStructureType = 1000003000
	VK_STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR                                      VkStructureType = 1000004000
	VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR                                       VkStructureType = 1000005000
	VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR                                   VkStructureType = 1000006000
	VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR                                   VkStructureType = 1000008000
	VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR                                     VkStructureType = 1000009000
	VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT                             VkStructureType = 1000011000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD              VkStructureType = 1000018000
	VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT                                 VkStructureType = 1000022000
	VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT                                  VkStructureType = 1000022001
	VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT                                      VkStructureType = 1000022002
	VK_STRUCTURE_TYPE_VIDEO_PROFILE_KHR                                                 VkStructureType = 1000023000
	VK_STRUCTURE_TYPE_VIDEO_CAPABILITIES_KHR                                            VkStructureType = 1000023001
	VK_STRUCTURE_TYPE_VIDEO_PICTURE_RESOURCE_KHR                                        VkStructureType = 1000023002
	VK_STRUCTURE_TYPE_VIDEO_GET_MEMORY_PROPERTIES_KHR                                   VkStructureType = 1000023003
	VK_STRUCTURE_TYPE_VIDEO_BIND_MEMORY_KHR                                             VkStructureType = 1000023004
	VK_STRUCTURE_TYPE_VIDEO_SESSION_CREATE_INFO_KHR                                     VkStructureType = 1000023005
	VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_CREATE_INFO_KHR                          VkStructureType = 1000023006
	VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_UPDATE_INFO_KHR                          VkStructureType = 1000023007
	VK_STRUCTURE_TYPE_VIDEO_BEGIN_CODING_INFO_KHR                                       VkStructureType = 1000023008
	VK_STRUCTURE_TYPE_VIDEO_END_CODING_INFO_KHR                                         VkStructureType = 1000023009
	VK_STRUCTURE_TYPE_VIDEO_CODING_CONTROL_INFO_KHR                                     VkStructureType = 1000023010
	VK_STRUCTURE_TYPE_VIDEO_REFERENCE_SLOT_KHR                                          VkStructureType = 1000023011
	VK_STRUCTURE_TYPE_VIDEO_QUEUE_FAMILY_PROPERTIES_2_KHR                               VkStructureType = 1000023012
	VK_STRUCTURE_TYPE_VIDEO_PROFILES_KHR                                                VkStructureType = 1000023013
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_FORMAT_INFO_KHR                             VkStructureType = 1000023014
	VK_STRUCTURE_TYPE_VIDEO_FORMAT_PROPERTIES_KHR                                       VkStructureType = 1000023015
	VK_STRUCTURE_TYPE_VIDEO_DECODE_INFO_KHR                                             VkStructureType = 1000024000
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV                         VkStructureType = 1000026000
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV                        VkStructureType = 1000026001
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV                      VkStructureType = 1000026002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT                   VkStructureType = 1000028000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT                 VkStructureType = 1000028001
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT               VkStructureType = 1000028002
	VK_STRUCTURE_TYPE_CU_MODULE_CREATE_INFO_NVX                                         VkStructureType = 1000029000
	VK_STRUCTURE_TYPE_CU_FUNCTION_CREATE_INFO_NVX                                       VkStructureType = 1000029001
	VK_STRUCTURE_TYPE_CU_LAUNCH_INFO_NVX                                                VkStructureType = 1000029002
	VK_STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX                                        VkStructureType = 1000030000
	VK_STRUCTURE_TYPE_IMAGE_VIEW_ADDRESS_PROPERTIES_NVX                                 VkStructureType = 1000030001
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_CAPABILITIES_EXT                                VkStructureType = 1000038000
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_CREATE_INFO_EXT                         VkStructureType = 1000038001
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_CREATE_INFO_EXT              VkStructureType = 1000038002
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_ADD_INFO_EXT                 VkStructureType = 1000038003
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_VCL_FRAME_INFO_EXT                              VkStructureType = 1000038004
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_DPB_SLOT_INFO_EXT                               VkStructureType = 1000038005
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_NALU_SLICE_EXT                                  VkStructureType = 1000038006
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_EMIT_PICTURE_PARAMETERS_EXT                     VkStructureType = 1000038007
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_PROFILE_EXT                                     VkStructureType = 1000038008
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_CAPABILITIES_EXT                                VkStructureType = 1000040000
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_CREATE_INFO_EXT                         VkStructureType = 1000040001
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PICTURE_INFO_EXT                                VkStructureType = 1000040002
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_MVC_EXT                                         VkStructureType = 1000040003
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PROFILE_EXT                                     VkStructureType = 1000040004
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_CREATE_INFO_EXT              VkStructureType = 1000040005
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_ADD_INFO_EXT                 VkStructureType = 1000040006
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_DPB_SLOT_INFO_EXT                               VkStructureType = 1000040007
	VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD                          VkStructureType = 1000041000
	VK_STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP                         VkStructureType = 1000049000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV                  VkStructureType = 1000050000
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV                              VkStructureType = 1000056000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV                                    VkStructureType = 1000056001
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV                                VkStructureType = 1000057000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV                                VkStructureType = 1000057001
	VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV                         VkStructureType = 1000058000
	VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT                                              VkStructureType = 1000061000
	VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN                                         VkStructureType = 1000062000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT         VkStructureType = 1000066000
	VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT                                   VkStructureType = 1000067000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT                          VkStructureType = 1000067001
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR                               VkStructureType = 1000073000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR                               VkStructureType = 1000073001
	VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR                                VkStructureType = 1000073002
	VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR                                  VkStructureType = 1000073003
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR                                         VkStructureType = 1000074000
	VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR                                          VkStructureType = 1000074001
	VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR                                            VkStructureType = 1000074002
	VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR                        VkStructureType = 1000075000
	VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                            VkStructureType = 1000078000
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                            VkStructureType = 1000078001
	VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR                                       VkStructureType = 1000078002
	VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR                               VkStructureType = 1000078003
	VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR                                      VkStructureType = 1000079000
	VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR                                         VkStructureType = 1000079001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR                    VkStructureType = 1000080000
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT         VkStructureType = 1000081000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT                VkStructureType = 1000081001
	VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT                              VkStructureType = 1000081002
	VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR                                               VkStructureType = 1000084000
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV                  VkStructureType = 1000087000
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT                                        VkStructureType = 1000090000
	VK_STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT                                            VkStructureType = 1000091000
	VK_STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT                                             VkStructureType = 1000091001
	VK_STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT                                            VkStructureType = 1000091002
	VK_STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT                                 VkStructureType = 1000091003
	VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE                                         VkStructureType = 1000092000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX      VkStructureType = 1000097000
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV                    VkStructureType = 1000098000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT                  VkStructureType = 1000099000
	VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT                  VkStructureType = 1000099001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT         VkStructureType = 1000101000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT         VkStructureType = 1000101001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT                    VkStructureType = 1000102000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT           VkStructureType = 1000102001
	VK_STRUCTURE_TYPE_HDR_METADATA_EXT                                                  VkStructureType = 1000105000
	VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR                           VkStructureType = 1000111000
	VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR                                VkStructureType = 1000114000
	VK_STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR                                VkStructureType = 1000114001
	VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR                                   VkStructureType = 1000114002
	VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR                                          VkStructureType = 1000115000
	VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR                                             VkStructureType = 1000115001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_FEATURES_KHR                    VkStructureType = 1000116000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_PROPERTIES_KHR                  VkStructureType = 1000116001
	VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR                            VkStructureType = 1000116002
	VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR                                 VkStructureType = 1000116003
	VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR                                   VkStructureType = 1000116004
	VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_KHR                                           VkStructureType = 1000116005
	VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_DESCRIPTION_KHR                               VkStructureType = 1000116006
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR                                VkStructureType = 1000119000
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR                                        VkStructureType = 1000119001
	VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR                                              VkStructureType = 1000119002
	VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR                                          VkStructureType = 1000121000
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR                                    VkStructureType = 1000121001
	VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR                                     VkStructureType = 1000121002
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR                                          VkStructureType = 1000121003
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR                                  VkStructureType = 1000121004
	VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK                                       VkStructureType = 1000122000
	VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK                                     VkStructureType = 1000123000
	VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT                                  VkStructureType = 1000128000
	VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT                                   VkStructureType = 1000128001
	VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT                                             VkStructureType = 1000128002
	VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT                           VkStructureType = 1000128003
	VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT                             VkStructureType = 1000128004
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID                             VkStructureType = 1000129000
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID                        VkStructureType = 1000129001
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID                 VkStructureType = 1000129002
	VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                       VkStructureType = 1000129003
	VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                   VkStructureType = 1000129004
	VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID                                           VkStructureType = 1000129005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT                 VkStructureType = 1000138000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT               VkStructureType = 1000138001
	VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT                     VkStructureType = 1000138002
	VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT              VkStructureType = 1000138003
	VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT                                         VkStructureType = 1000143000
	VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT                       VkStructureType = 1000143001
	VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT                   VkStructureType = 1000143002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT                   VkStructureType = 1000143003
	VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT                                        VkStructureType = 1000143004
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT             VkStructureType = 1000148000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT           VkStructureType = 1000148001
	VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT               VkStructureType = 1000148002
	VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV                   VkStructureType = 1000149000
	VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_KHR                   VkStructureType = 1000150007
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR                    VkStructureType = 1000150000
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DEVICE_ADDRESS_INFO_KHR                    VkStructureType = 1000150002
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR                    VkStructureType = 1000150003
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR                VkStructureType = 1000150004
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR                VkStructureType = 1000150005
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR                               VkStructureType = 1000150006
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_VERSION_INFO_KHR                           VkStructureType = 1000150009
	VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR                              VkStructureType = 1000150010
	VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR                    VkStructureType = 1000150011
	VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR                    VkStructureType = 1000150012
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR               VkStructureType = 1000150013
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR             VkStructureType = 1000150014
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_KHR                            VkStructureType = 1000150017
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_SIZES_INFO_KHR                       VkStructureType = 1000150020
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_FEATURES_KHR                 VkStructureType = 1000347000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_PROPERTIES_KHR               VkStructureType = 1000347001
	VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_KHR                              VkStructureType = 1000150015
	VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_KHR                          VkStructureType = 1000150016
	VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_INTERFACE_CREATE_INFO_KHR                    VkStructureType = 1000150018
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_QUERY_FEATURES_KHR                            VkStructureType = 1000348013
	VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV                 VkStructureType = 1000152000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV                    VkStructureType = 1000154000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV                  VkStructureType = 1000154001
	VK_STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT                           VkStructureType = 1000158000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT                VkStructureType = 1000158002
	VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT                    VkStructureType = 1000158003
	VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT                VkStructureType = 1000158004
	VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT                          VkStructureType = 1000158005
	VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT                                  VkStructureType = 1000160000
	VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT                    VkStructureType = 1000160001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR                   VkStructureType = 1000163000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR                 VkStructureType = 1000163001
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV         VkStructureType = 1000164000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV                    VkStructureType = 1000164001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV                  VkStructureType = 1000164002
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV        VkStructureType = 1000164005
	VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV                               VkStructureType = 1000165000
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV                             VkStructureType = 1000165001
	VK_STRUCTURE_TYPE_GEOMETRY_NV                                                       VkStructureType = 1000165003
	VK_STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV                                             VkStructureType = 1000165004
	VK_STRUCTURE_TYPE_GEOMETRY_AABB_NV                                                  VkStructureType = 1000165005
	VK_STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV                        VkStructureType = 1000165006
	VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV                    VkStructureType = 1000165007
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV                VkStructureType = 1000165008
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV                         VkStructureType = 1000165009
	VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV                           VkStructureType = 1000165011
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV                                    VkStructureType = 1000165012
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV          VkStructureType = 1000166000
	VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV        VkStructureType = 1000166001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT                  VkStructureType = 1000170000
	VK_STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT               VkStructureType = 1000170001
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT                      VkStructureType = 1000174000
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT                               VkStructureType = 1000178000
	VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT                                VkStructureType = 1000178001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT               VkStructureType = 1000178002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CLOCK_FEATURES_KHR                         VkStructureType = 1000181000
	VK_STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD                         VkStructureType = 1000183000
	VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT                                     VkStructureType = 1000184000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD                        VkStructureType = 1000185000
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_CAPABILITIES_EXT                                VkStructureType = 1000187000
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_CREATE_INFO_EXT                         VkStructureType = 1000187001
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_CREATE_INFO_EXT              VkStructureType = 1000187002
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_ADD_INFO_EXT                 VkStructureType = 1000187003
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PROFILE_EXT                                     VkStructureType = 1000187004
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PICTURE_INFO_EXT                                VkStructureType = 1000187005
	VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_DPB_SLOT_INFO_EXT                               VkStructureType = 1000187006
	VK_STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD                      VkStructureType = 1000189000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT           VkStructureType = 1000190000
	VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT               VkStructureType = 1000190001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT             VkStructureType = 1000190002
	VK_STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP                                           VkStructureType = 1000191000
	VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT                        VkStructureType = 1000192000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV            VkStructureType = 1000201000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV                           VkStructureType = 1000202000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV                         VkStructureType = 1000202001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV           VkStructureType = 1000203000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV                VkStructureType = 1000204000
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV          VkStructureType = 1000205000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV                     VkStructureType = 1000205002
	VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV                                                VkStructureType = 1000206000
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV                             VkStructureType = 1000206001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL         VkStructureType = 1000209000
	VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL                    VkStructureType = 1000210000
	VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL                             VkStructureType = 1000210001
	VK_STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL                                     VkStructureType = 1000210002
	VK_STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL                              VkStructureType = 1000210003
	VK_STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL                                   VkStructureType = 1000210004
	VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL                      VkStructureType = 1000210005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT                       VkStructureType = 1000212000
	VK_STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD                       VkStructureType = 1000213000
	VK_STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD                      VkStructureType = 1000213001
	VK_STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA                             VkStructureType = 1000214000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TERMINATE_INVOCATION_FEATURES_KHR          VkStructureType = 1000215000
	VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT                                     VkStructureType = 1000217000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT                 VkStructureType = 1000218000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT               VkStructureType = 1000218001
	VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT                  VkStructureType = 1000218002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT              VkStructureType = 1000225000
	VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT      VkStructureType = 1000225001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT                VkStructureType = 1000225002
	VK_STRUCTURE_TYPE_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR                         VkStructureType = 1000226000
	VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_STATE_CREATE_INFO_KHR              VkStructureType = 1000226001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_PROPERTIES_KHR              VkStructureType = 1000226002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_FEATURES_KHR                VkStructureType = 1000226003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_KHR                         VkStructureType = 1000226004
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD                      VkStructureType = 1000227000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD                      VkStructureType = 1000229000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_ATOMIC_INT64_FEATURES_EXT            VkStructureType = 1000234000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT                      VkStructureType = 1000237000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT                      VkStructureType = 1000238000
	VK_STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT                                 VkStructureType = 1000238001
	VK_STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR                                VkStructureType = 1000239000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV   VkStructureType = 1000240000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT                VkStructureType = 1000244000
	VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT                             VkStructureType = 1000244002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES_EXT                               VkStructureType = 1000245000
	VK_STRUCTURE_TYPE_VALIDATION_FEATURES_EXT                                           VkStructureType = 1000247000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_WAIT_FEATURES_KHR                         VkStructureType = 1000248000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV                    VkStructureType = 1000249000
	VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV                                  VkStructureType = 1000249001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV                  VkStructureType = 1000249002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV               VkStructureType = 1000250000
	VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV                  VkStructureType = 1000250001
	VK_STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV                          VkStructureType = 1000250002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT            VkStructureType = 1000251000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT                   VkStructureType = 1000252000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_FEATURES_EXT                     VkStructureType = 1000254000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_PROVOKING_VERTEX_STATE_CREATE_INFO_EXT     VkStructureType = 1000254001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_PROPERTIES_EXT                   VkStructureType = 1000254002
	VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT                            VkStructureType = 1000255000
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT                    VkStructureType = 1000255002
	VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT                      VkStructureType = 1000255001
	VK_STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT                                  VkStructureType = 1000256000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT                   VkStructureType = 1000259000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT                 VkStructureType = 1000259001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT                 VkStructureType = 1000259002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_FEATURES_EXT                  VkStructureType = 1000260000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_EXT                     VkStructureType = 1000265000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT               VkStructureType = 1000267000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR       VkStructureType = 1000269000
	VK_STRUCTURE_TYPE_PIPELINE_INFO_KHR                                                 VkStructureType = 1000269001
	VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR                                VkStructureType = 1000269002
	VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR                                      VkStructureType = 1000269003
	VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR                                 VkStructureType = 1000269004
	VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR                   VkStructureType = 1000269005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_2_FEATURES_EXT                VkStructureType = 1000273000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES_EXT   VkStructureType = 1000276000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_PROPERTIES_NV           VkStructureType = 1000277000
	VK_STRUCTURE_TYPE_GRAPHICS_SHADER_GROUP_CREATE_INFO_NV                              VkStructureType = 1000277001
	VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_SHADER_GROUPS_CREATE_INFO_NV                    VkStructureType = 1000277002
	VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_TOKEN_NV                                 VkStructureType = 1000277003
	VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NV                           VkStructureType = 1000277004
	VK_STRUCTURE_TYPE_GENERATED_COMMANDS_INFO_NV                                        VkStructureType = 1000277005
	VK_STRUCTURE_TYPE_GENERATED_COMMANDS_MEMORY_REQUIREMENTS_INFO_NV                    VkStructureType = 1000277006
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_FEATURES_NV             VkStructureType = 1000277007
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INHERITED_VIEWPORT_SCISSOR_FEATURES_NV            VkStructureType = 1000278000
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_VIEWPORT_SCISSOR_INFO_NV               VkStructureType = 1000278001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_FEATURES_EXT               VkStructureType = 1000281000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES_EXT             VkStructureType = 1000281001
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDER_PASS_TRANSFORM_INFO_QCOM        VkStructureType = 1000282000
	VK_STRUCTURE_TYPE_RENDER_PASS_TRANSFORM_BEGIN_INFO_QCOM                             VkStructureType = 1000282001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_MEMORY_REPORT_FEATURES_EXT                 VkStructureType = 1000284000
	VK_STRUCTURE_TYPE_DEVICE_DEVICE_MEMORY_REPORT_CREATE_INFO_EXT                       VkStructureType = 1000284001
	VK_STRUCTURE_TYPE_DEVICE_MEMORY_REPORT_CALLBACK_DATA_EXT                            VkStructureType = 1000284002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_FEATURES_EXT                         VkStructureType = 1000286000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_PROPERTIES_EXT                       VkStructureType = 1000286001
	VK_STRUCTURE_TYPE_SAMPLER_CUSTOM_BORDER_COLOR_CREATE_INFO_EXT                       VkStructureType = 1000287000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_PROPERTIES_EXT                VkStructureType = 1000287001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_FEATURES_EXT                  VkStructureType = 1000287002
	VK_STRUCTURE_TYPE_PIPELINE_LIBRARY_CREATE_INFO_KHR                                  VkStructureType = 1000290000
	VK_STRUCTURE_TYPE_PRESENT_ID_KHR                                                    VkStructureType = 1000294000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_ID_FEATURES_KHR                           VkStructureType = 1000294001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES_EXT                         VkStructureType = 1000295000
	VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO_EXT                               VkStructureType = 1000295001
	VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO_EXT                                 VkStructureType = 1000295002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_CREATION_CACHE_CONTROL_FEATURES_EXT      VkStructureType = 1000297000
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_INFO_KHR                                             VkStructureType = 1000299000
	VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_INFO_KHR                                VkStructureType = 1000299001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DIAGNOSTICS_CONFIG_FEATURES_NV                    VkStructureType = 1000300000
	VK_STRUCTURE_TYPE_DEVICE_DIAGNOSTICS_CONFIG_CREATE_INFO_NV                          VkStructureType = 1000300001
	VK_STRUCTURE_TYPE_MEMORY_BARRIER_2_KHR                                              VkStructureType = 1000314000
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2_KHR                                       VkStructureType = 1000314001
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2_KHR                                        VkStructureType = 1000314002
	VK_STRUCTURE_TYPE_DEPENDENCY_INFO_KHR                                               VkStructureType = 1000314003
	VK_STRUCTURE_TYPE_SUBMIT_INFO_2_KHR                                                 VkStructureType = 1000314004
	VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO_KHR                                         VkStructureType = 1000314005
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO_KHR                                    VkStructureType = 1000314006
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES_KHR                    VkStructureType = 1000314007
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_2_NV                           VkStructureType = 1000314008
	VK_STRUCTURE_TYPE_CHECKPOINT_DATA_2_NV                                              VkStructureType = 1000314009
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_FEATURES_KHR VkStructureType = 1000323000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES_KHR     VkStructureType = 1000325000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_PROPERTIES_NV         VkStructureType = 1000326000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_FEATURES_NV           VkStructureType = 1000326001
	VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_ENUM_STATE_CREATE_INFO_NV          VkStructureType = 1000326002
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_MOTION_TRIANGLES_DATA_NV          VkStructureType = 1000327000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MOTION_BLUR_FEATURES_NV               VkStructureType = 1000327001
	VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MOTION_INFO_NV                             VkStructureType = 1000327002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_2_PLANE_444_FORMATS_FEATURES_EXT            VkStructureType = 1000330000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_FEATURES_EXT               VkStructureType = 1000332000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_PROPERTIES_EXT             VkStructureType = 1000332001
	VK_STRUCTURE_TYPE_COPY_COMMAND_TRANSFORM_INFO_QCOM                                  VkStructureType = 1000333000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES_EXT                     VkStructureType = 1000335000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_FEATURES_KHR     VkStructureType = 1000336000
	VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2_KHR                                            VkStructureType = 1000337000
	VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2_KHR                                             VkStructureType = 1000337001
	VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2_KHR                                   VkStructureType = 1000337002
	VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2_KHR                                   VkStructureType = 1000337003
	VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2_KHR                                             VkStructureType = 1000337004
	VK_STRUCTURE_TYPE_RESOLVE_IMAGE_INFO_2_KHR                                          VkStructureType = 1000337005
	VK_STRUCTURE_TYPE_BUFFER_COPY_2_KHR                                                 VkStructureType = 1000337006
	VK_STRUCTURE_TYPE_IMAGE_COPY_2_KHR                                                  VkStructureType = 1000337007
	VK_STRUCTURE_TYPE_IMAGE_BLIT_2_KHR                                                  VkStructureType = 1000337008
	VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2_KHR                                           VkStructureType = 1000337009
	VK_STRUCTURE_TYPE_IMAGE_RESOLVE_2_KHR                                               VkStructureType = 1000337010
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT                         VkStructureType = 1000340000
	VK_STRUCTURE_TYPE_DIRECTFB_SURFACE_CREATE_INFO_EXT                                  VkStructureType = 1000346000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MUTABLE_DESCRIPTOR_TYPE_FEATURES_VALVE            VkStructureType = 1000351000
	VK_STRUCTURE_TYPE_MUTABLE_DESCRIPTOR_TYPE_CREATE_INFO_VALVE                         VkStructureType = 1000351002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_INPUT_DYNAMIC_STATE_FEATURES_EXT           VkStructureType = 1000352000
	VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT                            VkStructureType = 1000352001
	VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT                          VkStructureType = 1000352002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRM_PROPERTIES_EXT                                VkStructureType = 1000353000
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_ZIRCON_HANDLE_INFO_FUCHSIA                          VkStructureType = 1000364000
	VK_STRUCTURE_TYPE_MEMORY_ZIRCON_HANDLE_PROPERTIES_FUCHSIA                           VkStructureType = 1000364001
	VK_STRUCTURE_TYPE_MEMORY_GET_ZIRCON_HANDLE_INFO_FUCHSIA                             VkStructureType = 1000364002
	VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_ZIRCON_HANDLE_INFO_FUCHSIA                       VkStructureType = 1000365000
	VK_STRUCTURE_TYPE_SEMAPHORE_GET_ZIRCON_HANDLE_INFO_FUCHSIA                          VkStructureType = 1000365001
	VK_STRUCTURE_TYPE_SUBPASS_SHADING_PIPELINE_CREATE_INFO_HUAWEI                       VkStructureType = 1000369000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_FEATURES_HUAWEI                   VkStructureType = 1000369001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_PROPERTIES_HUAWEI                 VkStructureType = 1000369002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INVOCATION_MASK_FEATURES_HUAWEI                   VkStructureType = 1000370000
	VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV                                 VkStructureType = 1000371000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_RDMA_FEATURES_NV                  VkStructureType = 1000371001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT             VkStructureType = 1000377000
	VK_STRUCTURE_TYPE_SCREEN_SURFACE_CREATE_INFO_QNX                                    VkStructureType = 1000378000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COLOR_WRITE_ENABLE_FEATURES_EXT                   VkStructureType = 1000381000
	VK_STRUCTURE_TYPE_PIPELINE_COLOR_WRITE_CREATE_INFO_EXT                              VkStructureType = 1000381001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_EXT                VkStructureType = 1000388000
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_EXT                       VkStructureType = 1000388001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_FEATURES_EXT                           VkStructureType = 1000392000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_PROPERTIES_EXT                         VkStructureType = 1000392001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES                         VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES                    VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT                                      VkStructureType = VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT
	VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR                            VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR                          VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR                                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR                                           VkStructureType = VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR                                     VkStructureType = VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR                           VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR                                     VkStructureType = VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR                           VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR                              VkStructureType = VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR                    VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR                           VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR                        VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR                                      VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR                          VkStructureType = VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR                           VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR                              VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR                               VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR                    VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR                              VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR                          VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR                            VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR                                   VkStructureType = VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR                       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR                                  VkStructureType = VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT16_INT8_FEATURES_KHR                         VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR                        VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR                        VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES2_EXT                                         VkStructureType = VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR                VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES
	VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR                           VkStructureType = VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO
	VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO
	VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO
	VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR                                      VkStructureType = VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2
	VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR                                        VkStructureType = VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
	VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR                                         VkStructureType = VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2
	VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR                                          VkStructureType = VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2
	VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR                                     VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
	VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR                                            VkStructureType = VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
	VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR                                              VkStructureType = VK_STRUCTURE_TYPE_SUBPASS_END_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR                           VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR                                     VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR                                      VkStructureType = VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR               VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR                                  VkStructureType = VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR         VkStructureType = VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR                    VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR                                VkStructureType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT              VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES
	VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT                            VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR                             VkStructureType = VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR                              VkStructureType = VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR                       VkStructureType = VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR                                         VkStructureType = VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR                            VkStructureType = VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR                          VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR                                  VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR                          VkStructureType = VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR             VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR              VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR                                       VkStructureType = VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR                                        VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT               VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT                VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT        VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT       VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR                      VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR                         VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR                             VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR              VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES
	VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR                     VkStructureType = VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR                   VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES_KHR                 VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES
	VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO
	VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO_KHR                                VkStructureType = VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO
	VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR                                           VkStructureType = VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO
	VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR                                         VkStructureType = VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO
	VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL                                      VkStructureType = VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES
	VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR                           VkStructureType = VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT
	VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR                         VkStructureType = VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_ADDRESS_FEATURES_EXT                       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT
	VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT                                    VkStructureType = VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO
	VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT                               VkStructureType = VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR                VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
	VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO
	VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO
	VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO_KHR                   VkStructureType = VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES
	VK_STRUCTURE_TYPE_MAX_ENUM                                                          VkStructureType = 0x7FFFFFFF
)

type VkImageLayout = uint32

const (
	VK_IMAGE_LAYOUT_UNDEFINED                                      VkImageLayout = 0
	VK_IMAGE_LAYOUT_GENERAL                                        VkImageLayout = 1
	VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL                       VkImageLayout = 2
	VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL               VkImageLayout = 3
	VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL                VkImageLayout = 4
	VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL                       VkImageLayout = 5
	VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL                           VkImageLayout = 6
	VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL                           VkImageLayout = 7
	VK_IMAGE_LAYOUT_PREINITIALIZED                                 VkImageLayout = 8
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL     VkImageLayout = 1000117000
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL     VkImageLayout = 1000117001
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL                       VkImageLayout = 1000241000
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL                        VkImageLayout = 1000241001
	VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL                     VkImageLayout = 1000241002
	VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL                      VkImageLayout = 1000241003
	VK_IMAGE_LAYOUT_PRESENT_SRC_KHR                                VkImageLayout = 1000001002
	VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR                           VkImageLayout = 1000024000
	VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR                           VkImageLayout = 1000024001
	VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR                           VkImageLayout = 1000024002
	VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR                             VkImageLayout = 1000111000
	VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT               VkImageLayout = 1000218000
	VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR   VkImageLayout = 1000164003
	VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR                           VkImageLayout = 1000299000
	VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR                           VkImageLayout = 1000299001
	VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR                           VkImageLayout = 1000299002
	VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR                          VkImageLayout = 1000314000
	VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR                         VkImageLayout = 1000314001
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV                        VkImageLayout = VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR                   VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR                    VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL
	VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR                 VkImageLayout = VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL
	VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR                  VkImageLayout = VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL
	VK_IMAGE_LAYOUT_MAX_ENUM                                       VkImageLayout = 0x7FFFFFFF
)

type VkObjectType = uint32

const (
	VK_OBJECT_TYPE_UNKNOWN                         VkObjectType = 0
	VK_OBJECT_TYPE_INSTANCE                        VkObjectType = 1
	VK_OBJECT_TYPE_PHYSICAL_DEVICE                 VkObjectType = 2
	VK_OBJECT_TYPE_DEVICE                          VkObjectType = 3
	VK_OBJECT_TYPE_QUEUE                           VkObjectType = 4
	VK_OBJECT_TYPE_SEMAPHORE                       VkObjectType = 5
	VK_OBJECT_TYPE_COMMAND_BUFFER                  VkObjectType = 6
	VK_OBJECT_TYPE_FENCE                           VkObjectType = 7
	VK_OBJECT_TYPE_DEVICE_MEMORY                   VkObjectType = 8
	VK_OBJECT_TYPE_BUFFER                          VkObjectType = 9
	VK_OBJECT_TYPE_IMAGE                           VkObjectType = 10
	VK_OBJECT_TYPE_EVENT                           VkObjectType = 11
	VK_OBJECT_TYPE_QUERY_POOL                      VkObjectType = 12
	VK_OBJECT_TYPE_BUFFER_VIEW                     VkObjectType = 13
	VK_OBJECT_TYPE_IMAGE_VIEW                      VkObjectType = 14
	VK_OBJECT_TYPE_SHADER_MODULE                   VkObjectType = 15
	VK_OBJECT_TYPE_PIPELINE_CACHE                  VkObjectType = 16
	VK_OBJECT_TYPE_PIPELINE_LAYOUT                 VkObjectType = 17
	VK_OBJECT_TYPE_RENDER_PASS                     VkObjectType = 18
	VK_OBJECT_TYPE_PIPELINE                        VkObjectType = 19
	VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT           VkObjectType = 20
	VK_OBJECT_TYPE_SAMPLER                         VkObjectType = 21
	VK_OBJECT_TYPE_DESCRIPTOR_POOL                 VkObjectType = 22
	VK_OBJECT_TYPE_DESCRIPTOR_SET                  VkObjectType = 23
	VK_OBJECT_TYPE_FRAMEBUFFER                     VkObjectType = 24
	VK_OBJECT_TYPE_COMMAND_POOL                    VkObjectType = 25
	VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION        VkObjectType = 1000156000
	VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE      VkObjectType = 1000085000
	VK_OBJECT_TYPE_SURFACE_KHR                     VkObjectType = 1000000000
	VK_OBJECT_TYPE_SWAPCHAIN_KHR                   VkObjectType = 1000001000
	VK_OBJECT_TYPE_DISPLAY_KHR                     VkObjectType = 1000002000
	VK_OBJECT_TYPE_DISPLAY_MODE_KHR                VkObjectType = 1000002001
	VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT       VkObjectType = 1000011000
	VK_OBJECT_TYPE_VIDEO_SESSION_KHR               VkObjectType = 1000023000
	VK_OBJECT_TYPE_VIDEO_SESSION_PARAMETERS_KHR    VkObjectType = 1000023001
	VK_OBJECT_TYPE_CU_MODULE_NVX                   VkObjectType = 1000029000
	VK_OBJECT_TYPE_CU_FUNCTION_NVX                 VkObjectType = 1000029001
	VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT       VkObjectType = 1000128000
	VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR      VkObjectType = 1000150000
	VK_OBJECT_TYPE_VALIDATION_CACHE_EXT            VkObjectType = 1000160000
	VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV       VkObjectType = 1000165000
	VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL VkObjectType = 1000210000
	VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR          VkObjectType = 1000268000
	VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NV     VkObjectType = 1000277000
	VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT           VkObjectType = 1000295000
	VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR  VkObjectType = VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR    VkObjectType = VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
	VK_OBJECT_TYPE_MAX_ENUM                        VkObjectType = 0x7FFFFFFF
)

type VkPipelineCacheHeaderVersion = uint32

const (
	VK_PIPELINE_CACHE_HEADER_VERSION_ONE      VkPipelineCacheHeaderVersion = 1
	VK_PIPELINE_CACHE_HEADER_VERSION_MAX_ENUM VkPipelineCacheHeaderVersion = 0x7FFFFFFF
)

type VkVendorId = uint32

const (
	VK_VENDOR_ID_VIV      VkVendorId = 0x10001
	VK_VENDOR_ID_VSI      VkVendorId = 0x10002
	VK_VENDOR_ID_KAZAN    VkVendorId = 0x10003
	VK_VENDOR_ID_CODEPLAY VkVendorId = 0x10004
	VK_VENDOR_ID_MESA     VkVendorId = 0x10005
	VK_VENDOR_ID_POCL     VkVendorId = 0x10006
	VK_VENDOR_ID_MAX_ENUM VkVendorId = 0x7FFFFFFF
)

type VkSystemAllocationScope = uint32

const (
	VK_SYSTEM_ALLOCATION_SCOPE_COMMAND  VkSystemAllocationScope = 0
	VK_SYSTEM_ALLOCATION_SCOPE_OBJECT   VkSystemAllocationScope = 1
	VK_SYSTEM_ALLOCATION_SCOPE_CACHE    VkSystemAllocationScope = 2
	VK_SYSTEM_ALLOCATION_SCOPE_DEVICE   VkSystemAllocationScope = 3
	VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE VkSystemAllocationScope = 4
	VK_SYSTEM_ALLOCATION_SCOPE_MAX_ENUM VkSystemAllocationScope = 0x7FFFFFFF
)

type VkInternalAllocationType = uint32

const (
	VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE VkInternalAllocationType = 0
	VK_INTERNAL_ALLOCATION_TYPE_MAX_ENUM   VkInternalAllocationType = 0x7FFFFFFF
)

type VkFormat = uint32

const (
	VK_FORMAT_UNDEFINED                                      VkFormat = 0
	VK_FORMAT_R4G4_UNORM_PACK8                               VkFormat = 1
	VK_FORMAT_R4G4B4A4_UNORM_PACK16                          VkFormat = 2
	VK_FORMAT_B4G4R4A4_UNORM_PACK16                          VkFormat = 3
	VK_FORMAT_R5G6B5_UNORM_PACK16                            VkFormat = 4
	VK_FORMAT_B5G6R5_UNORM_PACK16                            VkFormat = 5
	VK_FORMAT_R5G5B5A1_UNORM_PACK16                          VkFormat = 6
	VK_FORMAT_B5G5R5A1_UNORM_PACK16                          VkFormat = 7
	VK_FORMAT_A1R5G5B5_UNORM_PACK16                          VkFormat = 8
	VK_FORMAT_R8_UNORM                                       VkFormat = 9
	VK_FORMAT_R8_SNORM                                       VkFormat = 10
	VK_FORMAT_R8_USCALED                                     VkFormat = 11
	VK_FORMAT_R8_SSCALED                                     VkFormat = 12
	VK_FORMAT_R8_UINT                                        VkFormat = 13
	VK_FORMAT_R8_SINT                                        VkFormat = 14
	VK_FORMAT_R8_SRGB                                        VkFormat = 15
	VK_FORMAT_R8G8_UNORM                                     VkFormat = 16
	VK_FORMAT_R8G8_SNORM                                     VkFormat = 17
	VK_FORMAT_R8G8_USCALED                                   VkFormat = 18
	VK_FORMAT_R8G8_SSCALED                                   VkFormat = 19
	VK_FORMAT_R8G8_UINT                                      VkFormat = 20
	VK_FORMAT_R8G8_SINT                                      VkFormat = 21
	VK_FORMAT_R8G8_SRGB                                      VkFormat = 22
	VK_FORMAT_R8G8B8_UNORM                                   VkFormat = 23
	VK_FORMAT_R8G8B8_SNORM                                   VkFormat = 24
	VK_FORMAT_R8G8B8_USCALED                                 VkFormat = 25
	VK_FORMAT_R8G8B8_SSCALED                                 VkFormat = 26
	VK_FORMAT_R8G8B8_UINT                                    VkFormat = 27
	VK_FORMAT_R8G8B8_SINT                                    VkFormat = 28
	VK_FORMAT_R8G8B8_SRGB                                    VkFormat = 29
	VK_FORMAT_B8G8R8_UNORM                                   VkFormat = 30
	VK_FORMAT_B8G8R8_SNORM                                   VkFormat = 31
	VK_FORMAT_B8G8R8_USCALED                                 VkFormat = 32
	VK_FORMAT_B8G8R8_SSCALED                                 VkFormat = 33
	VK_FORMAT_B8G8R8_UINT                                    VkFormat = 34
	VK_FORMAT_B8G8R8_SINT                                    VkFormat = 35
	VK_FORMAT_B8G8R8_SRGB                                    VkFormat = 36
	VK_FORMAT_R8G8B8A8_UNORM                                 VkFormat = 37
	VK_FORMAT_R8G8B8A8_SNORM                                 VkFormat = 38
	VK_FORMAT_R8G8B8A8_USCALED                               VkFormat = 39
	VK_FORMAT_R8G8B8A8_SSCALED                               VkFormat = 40
	VK_FORMAT_R8G8B8A8_UINT                                  VkFormat = 41
	VK_FORMAT_R8G8B8A8_SINT                                  VkFormat = 42
	VK_FORMAT_R8G8B8A8_SRGB                                  VkFormat = 43
	VK_FORMAT_B8G8R8A8_UNORM                                 VkFormat = 44
	VK_FORMAT_B8G8R8A8_SNORM                                 VkFormat = 45
	VK_FORMAT_B8G8R8A8_USCALED                               VkFormat = 46
	VK_FORMAT_B8G8R8A8_SSCALED                               VkFormat = 47
	VK_FORMAT_B8G8R8A8_UINT                                  VkFormat = 48
	VK_FORMAT_B8G8R8A8_SINT                                  VkFormat = 49
	VK_FORMAT_B8G8R8A8_SRGB                                  VkFormat = 50
	VK_FORMAT_A8B8G8R8_UNORM_PACK32                          VkFormat = 51
	VK_FORMAT_A8B8G8R8_SNORM_PACK32                          VkFormat = 52
	VK_FORMAT_A8B8G8R8_USCALED_PACK32                        VkFormat = 53
	VK_FORMAT_A8B8G8R8_SSCALED_PACK32                        VkFormat = 54
	VK_FORMAT_A8B8G8R8_UINT_PACK32                           VkFormat = 55
	VK_FORMAT_A8B8G8R8_SINT_PACK32                           VkFormat = 56
	VK_FORMAT_A8B8G8R8_SRGB_PACK32                           VkFormat = 57
	VK_FORMAT_A2R10G10B10_UNORM_PACK32                       VkFormat = 58
	VK_FORMAT_A2R10G10B10_SNORM_PACK32                       VkFormat = 59
	VK_FORMAT_A2R10G10B10_USCALED_PACK32                     VkFormat = 60
	VK_FORMAT_A2R10G10B10_SSCALED_PACK32                     VkFormat = 61
	VK_FORMAT_A2R10G10B10_UINT_PACK32                        VkFormat = 62
	VK_FORMAT_A2R10G10B10_SINT_PACK32                        VkFormat = 63
	VK_FORMAT_A2B10G10R10_UNORM_PACK32                       VkFormat = 64
	VK_FORMAT_A2B10G10R10_SNORM_PACK32                       VkFormat = 65
	VK_FORMAT_A2B10G10R10_USCALED_PACK32                     VkFormat = 66
	VK_FORMAT_A2B10G10R10_SSCALED_PACK32                     VkFormat = 67
	VK_FORMAT_A2B10G10R10_UINT_PACK32                        VkFormat = 68
	VK_FORMAT_A2B10G10R10_SINT_PACK32                        VkFormat = 69
	VK_FORMAT_R16_UNORM                                      VkFormat = 70
	VK_FORMAT_R16_SNORM                                      VkFormat = 71
	VK_FORMAT_R16_USCALED                                    VkFormat = 72
	VK_FORMAT_R16_SSCALED                                    VkFormat = 73
	VK_FORMAT_R16_UINT                                       VkFormat = 74
	VK_FORMAT_R16_SINT                                       VkFormat = 75
	VK_FORMAT_R16_SFLOAT                                     VkFormat = 76
	VK_FORMAT_R16G16_UNORM                                   VkFormat = 77
	VK_FORMAT_R16G16_SNORM                                   VkFormat = 78
	VK_FORMAT_R16G16_USCALED                                 VkFormat = 79
	VK_FORMAT_R16G16_SSCALED                                 VkFormat = 80
	VK_FORMAT_R16G16_UINT                                    VkFormat = 81
	VK_FORMAT_R16G16_SINT                                    VkFormat = 82
	VK_FORMAT_R16G16_SFLOAT                                  VkFormat = 83
	VK_FORMAT_R16G16B16_UNORM                                VkFormat = 84
	VK_FORMAT_R16G16B16_SNORM                                VkFormat = 85
	VK_FORMAT_R16G16B16_USCALED                              VkFormat = 86
	VK_FORMAT_R16G16B16_SSCALED                              VkFormat = 87
	VK_FORMAT_R16G16B16_UINT                                 VkFormat = 88
	VK_FORMAT_R16G16B16_SINT                                 VkFormat = 89
	VK_FORMAT_R16G16B16_SFLOAT                               VkFormat = 90
	VK_FORMAT_R16G16B16A16_UNORM                             VkFormat = 91
	VK_FORMAT_R16G16B16A16_SNORM                             VkFormat = 92
	VK_FORMAT_R16G16B16A16_USCALED                           VkFormat = 93
	VK_FORMAT_R16G16B16A16_SSCALED                           VkFormat = 94
	VK_FORMAT_R16G16B16A16_UINT                              VkFormat = 95
	VK_FORMAT_R16G16B16A16_SINT                              VkFormat = 96
	VK_FORMAT_R16G16B16A16_SFLOAT                            VkFormat = 97
	VK_FORMAT_R32_UINT                                       VkFormat = 98
	VK_FORMAT_R32_SINT                                       VkFormat = 99
	VK_FORMAT_R32_SFLOAT                                     VkFormat = 100
	VK_FORMAT_R32G32_UINT                                    VkFormat = 101
	VK_FORMAT_R32G32_SINT                                    VkFormat = 102
	VK_FORMAT_R32G32_SFLOAT                                  VkFormat = 103
	VK_FORMAT_R32G32B32_UINT                                 VkFormat = 104
	VK_FORMAT_R32G32B32_SINT                                 VkFormat = 105
	VK_FORMAT_R32G32B32_SFLOAT                               VkFormat = 106
	VK_FORMAT_R32G32B32A32_UINT                              VkFormat = 107
	VK_FORMAT_R32G32B32A32_SINT                              VkFormat = 108
	VK_FORMAT_R32G32B32A32_SFLOAT                            VkFormat = 109
	VK_FORMAT_R64_UINT                                       VkFormat = 110
	VK_FORMAT_R64_SINT                                       VkFormat = 111
	VK_FORMAT_R64_SFLOAT                                     VkFormat = 112
	VK_FORMAT_R64G64_UINT                                    VkFormat = 113
	VK_FORMAT_R64G64_SINT                                    VkFormat = 114
	VK_FORMAT_R64G64_SFLOAT                                  VkFormat = 115
	VK_FORMAT_R64G64B64_UINT                                 VkFormat = 116
	VK_FORMAT_R64G64B64_SINT                                 VkFormat = 117
	VK_FORMAT_R64G64B64_SFLOAT                               VkFormat = 118
	VK_FORMAT_R64G64B64A64_UINT                              VkFormat = 119
	VK_FORMAT_R64G64B64A64_SINT                              VkFormat = 120
	VK_FORMAT_R64G64B64A64_SFLOAT                            VkFormat = 121
	VK_FORMAT_B10G11R11_UFLOAT_PACK32                        VkFormat = 122
	VK_FORMAT_E5B9G9R9_UFLOAT_PACK32                         VkFormat = 123
	VK_FORMAT_D16_UNORM                                      VkFormat = 124
	VK_FORMAT_X8_D24_UNORM_PACK32                            VkFormat = 125
	VK_FORMAT_D32_SFLOAT                                     VkFormat = 126
	VK_FORMAT_S8_UINT                                        VkFormat = 127
	VK_FORMAT_D16_UNORM_S8_UINT                              VkFormat = 128
	VK_FORMAT_D24_UNORM_S8_UINT                              VkFormat = 129
	VK_FORMAT_D32_SFLOAT_S8_UINT                             VkFormat = 130
	VK_FORMAT_BC1_RGB_UNORM_BLOCK                            VkFormat = 131
	VK_FORMAT_BC1_RGB_SRGB_BLOCK                             VkFormat = 132
	VK_FORMAT_BC1_RGBA_UNORM_BLOCK                           VkFormat = 133
	VK_FORMAT_BC1_RGBA_SRGB_BLOCK                            VkFormat = 134
	VK_FORMAT_BC2_UNORM_BLOCK                                VkFormat = 135
	VK_FORMAT_BC2_SRGB_BLOCK                                 VkFormat = 136
	VK_FORMAT_BC3_UNORM_BLOCK                                VkFormat = 137
	VK_FORMAT_BC3_SRGB_BLOCK                                 VkFormat = 138
	VK_FORMAT_BC4_UNORM_BLOCK                                VkFormat = 139
	VK_FORMAT_BC4_SNORM_BLOCK                                VkFormat = 140
	VK_FORMAT_BC5_UNORM_BLOCK                                VkFormat = 141
	VK_FORMAT_BC5_SNORM_BLOCK                                VkFormat = 142
	VK_FORMAT_BC6H_UFLOAT_BLOCK                              VkFormat = 143
	VK_FORMAT_BC6H_SFLOAT_BLOCK                              VkFormat = 144
	VK_FORMAT_BC7_UNORM_BLOCK                                VkFormat = 145
	VK_FORMAT_BC7_SRGB_BLOCK                                 VkFormat = 146
	VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK                        VkFormat = 147
	VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK                         VkFormat = 148
	VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK                      VkFormat = 149
	VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK                       VkFormat = 150
	VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK                      VkFormat = 151
	VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK                       VkFormat = 152
	VK_FORMAT_EAC_R11_UNORM_BLOCK                            VkFormat = 153
	VK_FORMAT_EAC_R11_SNORM_BLOCK                            VkFormat = 154
	VK_FORMAT_EAC_R11G11_UNORM_BLOCK                         VkFormat = 155
	VK_FORMAT_EAC_R11G11_SNORM_BLOCK                         VkFormat = 156
	VK_FORMAT_ASTC_4x4_UNORM_BLOCK                           VkFormat = 157
	VK_FORMAT_ASTC_4x4_SRGB_BLOCK                            VkFormat = 158
	VK_FORMAT_ASTC_5x4_UNORM_BLOCK                           VkFormat = 159
	VK_FORMAT_ASTC_5x4_SRGB_BLOCK                            VkFormat = 160
	VK_FORMAT_ASTC_5x5_UNORM_BLOCK                           VkFormat = 161
	VK_FORMAT_ASTC_5x5_SRGB_BLOCK                            VkFormat = 162
	VK_FORMAT_ASTC_6x5_UNORM_BLOCK                           VkFormat = 163
	VK_FORMAT_ASTC_6x5_SRGB_BLOCK                            VkFormat = 164
	VK_FORMAT_ASTC_6x6_UNORM_BLOCK                           VkFormat = 165
	VK_FORMAT_ASTC_6x6_SRGB_BLOCK                            VkFormat = 166
	VK_FORMAT_ASTC_8x5_UNORM_BLOCK                           VkFormat = 167
	VK_FORMAT_ASTC_8x5_SRGB_BLOCK                            VkFormat = 168
	VK_FORMAT_ASTC_8x6_UNORM_BLOCK                           VkFormat = 169
	VK_FORMAT_ASTC_8x6_SRGB_BLOCK                            VkFormat = 170
	VK_FORMAT_ASTC_8x8_UNORM_BLOCK                           VkFormat = 171
	VK_FORMAT_ASTC_8x8_SRGB_BLOCK                            VkFormat = 172
	VK_FORMAT_ASTC_10x5_UNORM_BLOCK                          VkFormat = 173
	VK_FORMAT_ASTC_10x5_SRGB_BLOCK                           VkFormat = 174
	VK_FORMAT_ASTC_10x6_UNORM_BLOCK                          VkFormat = 175
	VK_FORMAT_ASTC_10x6_SRGB_BLOCK                           VkFormat = 176
	VK_FORMAT_ASTC_10x8_UNORM_BLOCK                          VkFormat = 177
	VK_FORMAT_ASTC_10x8_SRGB_BLOCK                           VkFormat = 178
	VK_FORMAT_ASTC_10x10_UNORM_BLOCK                         VkFormat = 179
	VK_FORMAT_ASTC_10x10_SRGB_BLOCK                          VkFormat = 180
	VK_FORMAT_ASTC_12x10_UNORM_BLOCK                         VkFormat = 181
	VK_FORMAT_ASTC_12x10_SRGB_BLOCK                          VkFormat = 182
	VK_FORMAT_ASTC_12x12_UNORM_BLOCK                         VkFormat = 183
	VK_FORMAT_ASTC_12x12_SRGB_BLOCK                          VkFormat = 184
	VK_FORMAT_G8B8G8R8_422_UNORM                             VkFormat = 1000156000
	VK_FORMAT_B8G8R8G8_422_UNORM                             VkFormat = 1000156001
	VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM                      VkFormat = 1000156002
	VK_FORMAT_G8_B8R8_2PLANE_420_UNORM                       VkFormat = 1000156003
	VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM                      VkFormat = 1000156004
	VK_FORMAT_G8_B8R8_2PLANE_422_UNORM                       VkFormat = 1000156005
	VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM                      VkFormat = 1000156006
	VK_FORMAT_R10X6_UNORM_PACK16                             VkFormat = 1000156007
	VK_FORMAT_R10X6G10X6_UNORM_2PACK16                       VkFormat = 1000156008
	VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16             VkFormat = 1000156009
	VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16         VkFormat = 1000156010
	VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16         VkFormat = 1000156011
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16     VkFormat = 1000156012
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16      VkFormat = 1000156013
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16     VkFormat = 1000156014
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16      VkFormat = 1000156015
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16     VkFormat = 1000156016
	VK_FORMAT_R12X4_UNORM_PACK16                             VkFormat = 1000156017
	VK_FORMAT_R12X4G12X4_UNORM_2PACK16                       VkFormat = 1000156018
	VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16             VkFormat = 1000156019
	VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16         VkFormat = 1000156020
	VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16         VkFormat = 1000156021
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16     VkFormat = 1000156022
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16      VkFormat = 1000156023
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16     VkFormat = 1000156024
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16      VkFormat = 1000156025
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16     VkFormat = 1000156026
	VK_FORMAT_G16B16G16R16_422_UNORM                         VkFormat = 1000156027
	VK_FORMAT_B16G16R16G16_422_UNORM                         VkFormat = 1000156028
	VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM                   VkFormat = 1000156029
	VK_FORMAT_G16_B16R16_2PLANE_420_UNORM                    VkFormat = 1000156030
	VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM                   VkFormat = 1000156031
	VK_FORMAT_G16_B16R16_2PLANE_422_UNORM                    VkFormat = 1000156032
	VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM                   VkFormat = 1000156033
	VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054000
	VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054001
	VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054002
	VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054003
	VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054004
	VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054005
	VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054006
	VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054007
	VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT                      VkFormat = 1000066000
	VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT                      VkFormat = 1000066001
	VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT                      VkFormat = 1000066002
	VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT                      VkFormat = 1000066003
	VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT                      VkFormat = 1000066004
	VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT                      VkFormat = 1000066005
	VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT                      VkFormat = 1000066006
	VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT                      VkFormat = 1000066007
	VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT                     VkFormat = 1000066008
	VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT                     VkFormat = 1000066009
	VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT                     VkFormat = 1000066010
	VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT                    VkFormat = 1000066011
	VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT                    VkFormat = 1000066012
	VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT                    VkFormat = 1000066013
	VK_FORMAT_G8_B8R8_2PLANE_444_UNORM_EXT                   VkFormat = 1000330000
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16_EXT  VkFormat = 1000330001
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16_EXT  VkFormat = 1000330002
	VK_FORMAT_G16_B16R16_2PLANE_444_UNORM_EXT                VkFormat = 1000330003
	VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT                      VkFormat = 1000340000
	VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT                      VkFormat = 1000340001
	VK_FORMAT_G8B8G8R8_422_UNORM_KHR                         VkFormat = VK_FORMAT_G8B8G8R8_422_UNORM
	VK_FORMAT_B8G8R8G8_422_UNORM_KHR                         VkFormat = VK_FORMAT_B8G8R8G8_422_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	VK_FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR                   VkFormat = VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	VK_FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR                   VkFormat = VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM
	VK_FORMAT_R10X6_UNORM_PACK16_KHR                         VkFormat = VK_FORMAT_R10X6_UNORM_PACK16
	VK_FORMAT_R10X6G10X6_UNORM_2PACK16_KHR                   VkFormat = VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR         VkFormat = VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	VK_FORMAT_R12X4_UNORM_PACK16_KHR                         VkFormat = VK_FORMAT_R12X4_UNORM_PACK16
	VK_FORMAT_R12X4G12X4_UNORM_2PACK16_KHR                   VkFormat = VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR         VkFormat = VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	VK_FORMAT_G16B16G16R16_422_UNORM_KHR                     VkFormat = VK_FORMAT_G16B16G16R16_422_UNORM
	VK_FORMAT_B16G16R16G16_422_UNORM_KHR                     VkFormat = VK_FORMAT_B16G16R16G16_422_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	VK_FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR                VkFormat = VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	VK_FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR                VkFormat = VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM
	VK_FORMAT_MAX_ENUM                                       VkFormat = 0x7FFFFFFF
)

type VkImageTiling = uint32

const (
	VK_IMAGE_TILING_OPTIMAL                 VkImageTiling = 0
	VK_IMAGE_TILING_LINEAR                  VkImageTiling = 1
	VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT VkImageTiling = 1000158000
	VK_IMAGE_TILING_MAX_ENUM                VkImageTiling = 0x7FFFFFFF
)

type VkImageType = uint32

const (
	VK_IMAGE_TYPE_1D       VkImageType = 0
	VK_IMAGE_TYPE_2D       VkImageType = 1
	VK_IMAGE_TYPE_3D       VkImageType = 2
	VK_IMAGE_TYPE_MAX_ENUM VkImageType = 0x7FFFFFFF
)

type VkPhysicalDeviceType = uint32

const (
	VK_PHYSICAL_DEVICE_TYPE_OTHER          VkPhysicalDeviceType = 0
	VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU VkPhysicalDeviceType = 1
	VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU   VkPhysicalDeviceType = 2
	VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU    VkPhysicalDeviceType = 3
	VK_PHYSICAL_DEVICE_TYPE_CPU            VkPhysicalDeviceType = 4
	VK_PHYSICAL_DEVICE_TYPE_MAX_ENUM       VkPhysicalDeviceType = 0x7FFFFFFF
)

type VkQueryType = uint32

const (
	VK_QUERY_TYPE_OCCLUSION                                     VkQueryType = 0
	VK_QUERY_TYPE_PIPELINE_STATISTICS                           VkQueryType = 1
	VK_QUERY_TYPE_TIMESTAMP                                     VkQueryType = 2
	VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR                        VkQueryType = 1000023000
	VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT                 VkQueryType = 1000028004
	VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR                         VkQueryType = 1000116000
	VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR     VkQueryType = 1000150000
	VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR VkQueryType = 1000150001
	VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV      VkQueryType = 1000165000
	VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL                       VkQueryType = 1000210000
	VK_QUERY_TYPE_VIDEO_ENCODE_BITSTREAM_BUFFER_RANGE_KHR       VkQueryType = 1000299000
	VK_QUERY_TYPE_MAX_ENUM                                      VkQueryType = 0x7FFFFFFF
)

type VkSharingMode = uint32

const (
	VK_SHARING_MODE_EXCLUSIVE  VkSharingMode = 0
	VK_SHARING_MODE_CONCURRENT VkSharingMode = 1
	VK_SHARING_MODE_MAX_ENUM   VkSharingMode = 0x7FFFFFFF
)

type VkComponentSwizzle = uint32

const (
	VK_COMPONENT_SWIZZLE_IDENTITY VkComponentSwizzle = 0
	VK_COMPONENT_SWIZZLE_ZERO     VkComponentSwizzle = 1
	VK_COMPONENT_SWIZZLE_ONE      VkComponentSwizzle = 2
	VK_COMPONENT_SWIZZLE_R        VkComponentSwizzle = 3
	VK_COMPONENT_SWIZZLE_G        VkComponentSwizzle = 4
	VK_COMPONENT_SWIZZLE_B        VkComponentSwizzle = 5
	VK_COMPONENT_SWIZZLE_A        VkComponentSwizzle = 6
	VK_COMPONENT_SWIZZLE_MAX_ENUM VkComponentSwizzle = 0x7FFFFFFF
)

type VkImageViewType = uint32

const (
	VK_IMAGE_VIEW_TYPE_1D         VkImageViewType = 0
	VK_IMAGE_VIEW_TYPE_2D         VkImageViewType = 1
	VK_IMAGE_VIEW_TYPE_3D         VkImageViewType = 2
	VK_IMAGE_VIEW_TYPE_CUBE       VkImageViewType = 3
	VK_IMAGE_VIEW_TYPE_1D_ARRAY   VkImageViewType = 4
	VK_IMAGE_VIEW_TYPE_2D_ARRAY   VkImageViewType = 5
	VK_IMAGE_VIEW_TYPE_CUBE_ARRAY VkImageViewType = 6
	VK_IMAGE_VIEW_TYPE_MAX_ENUM   VkImageViewType = 0x7FFFFFFF
)

type VkBlendFactor = uint32

const (
	VK_BLEND_FACTOR_ZERO                     VkBlendFactor = 0
	VK_BLEND_FACTOR_ONE                      VkBlendFactor = 1
	VK_BLEND_FACTOR_SRC_COLOR                VkBlendFactor = 2
	VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR      VkBlendFactor = 3
	VK_BLEND_FACTOR_DST_COLOR                VkBlendFactor = 4
	VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR      VkBlendFactor = 5
	VK_BLEND_FACTOR_SRC_ALPHA                VkBlendFactor = 6
	VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA      VkBlendFactor = 7
	VK_BLEND_FACTOR_DST_ALPHA                VkBlendFactor = 8
	VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA      VkBlendFactor = 9
	VK_BLEND_FACTOR_CONSTANT_COLOR           VkBlendFactor = 10
	VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR VkBlendFactor = 11
	VK_BLEND_FACTOR_CONSTANT_ALPHA           VkBlendFactor = 12
	VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA VkBlendFactor = 13
	VK_BLEND_FACTOR_SRC_ALPHA_SATURATE       VkBlendFactor = 14
	VK_BLEND_FACTOR_SRC1_COLOR               VkBlendFactor = 15
	VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR     VkBlendFactor = 16
	VK_BLEND_FACTOR_SRC1_ALPHA               VkBlendFactor = 17
	VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA     VkBlendFactor = 18
	VK_BLEND_FACTOR_MAX_ENUM                 VkBlendFactor = 0x7FFFFFFF
)

type VkBlendOp = uint32

const (
	VK_BLEND_OP_ADD                    VkBlendOp = 0
	VK_BLEND_OP_SUBTRACT               VkBlendOp = 1
	VK_BLEND_OP_REVERSE_SUBTRACT       VkBlendOp = 2
	VK_BLEND_OP_MIN                    VkBlendOp = 3
	VK_BLEND_OP_MAX                    VkBlendOp = 4
	VK_BLEND_OP_ZERO_EXT               VkBlendOp = 1000148000
	VK_BLEND_OP_SRC_EXT                VkBlendOp = 1000148001
	VK_BLEND_OP_DST_EXT                VkBlendOp = 1000148002
	VK_BLEND_OP_SRC_OVER_EXT           VkBlendOp = 1000148003
	VK_BLEND_OP_DST_OVER_EXT           VkBlendOp = 1000148004
	VK_BLEND_OP_SRC_IN_EXT             VkBlendOp = 1000148005
	VK_BLEND_OP_DST_IN_EXT             VkBlendOp = 1000148006
	VK_BLEND_OP_SRC_OUT_EXT            VkBlendOp = 1000148007
	VK_BLEND_OP_DST_OUT_EXT            VkBlendOp = 1000148008
	VK_BLEND_OP_SRC_ATOP_EXT           VkBlendOp = 1000148009
	VK_BLEND_OP_DST_ATOP_EXT           VkBlendOp = 1000148010
	VK_BLEND_OP_XOR_EXT                VkBlendOp = 1000148011
	VK_BLEND_OP_MULTIPLY_EXT           VkBlendOp = 1000148012
	VK_BLEND_OP_SCREEN_EXT             VkBlendOp = 1000148013
	VK_BLEND_OP_OVERLAY_EXT            VkBlendOp = 1000148014
	VK_BLEND_OP_DARKEN_EXT             VkBlendOp = 1000148015
	VK_BLEND_OP_LIGHTEN_EXT            VkBlendOp = 1000148016
	VK_BLEND_OP_COLORDODGE_EXT         VkBlendOp = 1000148017
	VK_BLEND_OP_COLORBURN_EXT          VkBlendOp = 1000148018
	VK_BLEND_OP_HARDLIGHT_EXT          VkBlendOp = 1000148019
	VK_BLEND_OP_SOFTLIGHT_EXT          VkBlendOp = 1000148020
	VK_BLEND_OP_DIFFERENCE_EXT         VkBlendOp = 1000148021
	VK_BLEND_OP_EXCLUSION_EXT          VkBlendOp = 1000148022
	VK_BLEND_OP_INVERT_EXT             VkBlendOp = 1000148023
	VK_BLEND_OP_INVERT_RGB_EXT         VkBlendOp = 1000148024
	VK_BLEND_OP_LINEARDODGE_EXT        VkBlendOp = 1000148025
	VK_BLEND_OP_LINEARBURN_EXT         VkBlendOp = 1000148026
	VK_BLEND_OP_VIVIDLIGHT_EXT         VkBlendOp = 1000148027
	VK_BLEND_OP_LINEARLIGHT_EXT        VkBlendOp = 1000148028
	VK_BLEND_OP_PINLIGHT_EXT           VkBlendOp = 1000148029
	VK_BLEND_OP_HARDMIX_EXT            VkBlendOp = 1000148030
	VK_BLEND_OP_HSL_HUE_EXT            VkBlendOp = 1000148031
	VK_BLEND_OP_HSL_SATURATION_EXT     VkBlendOp = 1000148032
	VK_BLEND_OP_HSL_COLOR_EXT          VkBlendOp = 1000148033
	VK_BLEND_OP_HSL_LUMINOSITY_EXT     VkBlendOp = 1000148034
	VK_BLEND_OP_PLUS_EXT               VkBlendOp = 1000148035
	VK_BLEND_OP_PLUS_CLAMPED_EXT       VkBlendOp = 1000148036
	VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT VkBlendOp = 1000148037
	VK_BLEND_OP_PLUS_DARKER_EXT        VkBlendOp = 1000148038
	VK_BLEND_OP_MINUS_EXT              VkBlendOp = 1000148039
	VK_BLEND_OP_MINUS_CLAMPED_EXT      VkBlendOp = 1000148040
	VK_BLEND_OP_CONTRAST_EXT           VkBlendOp = 1000148041
	VK_BLEND_OP_INVERT_OVG_EXT         VkBlendOp = 1000148042
	VK_BLEND_OP_RED_EXT                VkBlendOp = 1000148043
	VK_BLEND_OP_GREEN_EXT              VkBlendOp = 1000148044
	VK_BLEND_OP_BLUE_EXT               VkBlendOp = 1000148045
	VK_BLEND_OP_MAX_ENUM               VkBlendOp = 0x7FFFFFFF
)

type VkCompareOp = uint32

const (
	VK_COMPARE_OP_NEVER            VkCompareOp = 0
	VK_COMPARE_OP_LESS             VkCompareOp = 1
	VK_COMPARE_OP_EQUAL            VkCompareOp = 2
	VK_COMPARE_OP_LESS_OR_EQUAL    VkCompareOp = 3
	VK_COMPARE_OP_GREATER          VkCompareOp = 4
	VK_COMPARE_OP_NOT_EQUAL        VkCompareOp = 5
	VK_COMPARE_OP_GREATER_OR_EQUAL VkCompareOp = 6
	VK_COMPARE_OP_ALWAYS           VkCompareOp = 7
	VK_COMPARE_OP_MAX_ENUM         VkCompareOp = 0x7FFFFFFF
)

type VkDynamicState = uint32

const (
	VK_DYNAMIC_STATE_VIEWPORT                            VkDynamicState = 0
	VK_DYNAMIC_STATE_SCISSOR                             VkDynamicState = 1
	VK_DYNAMIC_STATE_LINE_WIDTH                          VkDynamicState = 2
	VK_DYNAMIC_STATE_DEPTH_BIAS                          VkDynamicState = 3
	VK_DYNAMIC_STATE_BLEND_CONSTANTS                     VkDynamicState = 4
	VK_DYNAMIC_STATE_DEPTH_BOUNDS                        VkDynamicState = 5
	VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK                VkDynamicState = 6
	VK_DYNAMIC_STATE_STENCIL_WRITE_MASK                  VkDynamicState = 7
	VK_DYNAMIC_STATE_STENCIL_REFERENCE                   VkDynamicState = 8
	VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV               VkDynamicState = 1000087000
	VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT               VkDynamicState = 1000099000
	VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT                VkDynamicState = 1000143000
	VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR VkDynamicState = 1000347000
	VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV    VkDynamicState = 1000164004
	VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV     VkDynamicState = 1000164006
	VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV                VkDynamicState = 1000205001
	VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR           VkDynamicState = 1000226000
	VK_DYNAMIC_STATE_LINE_STIPPLE_EXT                    VkDynamicState = 1000259000
	VK_DYNAMIC_STATE_CULL_MODE_EXT                       VkDynamicState = 1000267000
	VK_DYNAMIC_STATE_FRONT_FACE_EXT                      VkDynamicState = 1000267001
	VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT              VkDynamicState = 1000267002
	VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT             VkDynamicState = 1000267003
	VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT              VkDynamicState = 1000267004
	VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT     VkDynamicState = 1000267005
	VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT               VkDynamicState = 1000267006
	VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT              VkDynamicState = 1000267007
	VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT                VkDynamicState = 1000267008
	VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT        VkDynamicState = 1000267009
	VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT             VkDynamicState = 1000267010
	VK_DYNAMIC_STATE_STENCIL_OP_EXT                      VkDynamicState = 1000267011
	VK_DYNAMIC_STATE_VERTEX_INPUT_EXT                    VkDynamicState = 1000352000
	VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT            VkDynamicState = 1000377000
	VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT       VkDynamicState = 1000377001
	VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT               VkDynamicState = 1000377002
	VK_DYNAMIC_STATE_LOGIC_OP_EXT                        VkDynamicState = 1000377003
	VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT        VkDynamicState = 1000377004
	VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT              VkDynamicState = 1000381000
	VK_DYNAMIC_STATE_MAX_ENUM                            VkDynamicState = 0x7FFFFFFF
)

type VkFrontFace = uint32

const (
	VK_FRONT_FACE_COUNTER_CLOCKWISE VkFrontFace = 0
	VK_FRONT_FACE_CLOCKWISE         VkFrontFace = 1
	VK_FRONT_FACE_MAX_ENUM          VkFrontFace = 0x7FFFFFFF
)

type VkVertexInputRate = uint32

const (
	VK_VERTEX_INPUT_RATE_VERTEX   VkVertexInputRate = 0
	VK_VERTEX_INPUT_RATE_INSTANCE VkVertexInputRate = 1
	VK_VERTEX_INPUT_RATE_MAX_ENUM VkVertexInputRate = 0x7FFFFFFF
)

type VkPrimitiveTopology = uint32

const (
	VK_PRIMITIVE_TOPOLOGY_POINT_LIST                    VkPrimitiveTopology = 0
	VK_PRIMITIVE_TOPOLOGY_LINE_LIST                     VkPrimitiveTopology = 1
	VK_PRIMITIVE_TOPOLOGY_LINE_STRIP                    VkPrimitiveTopology = 2
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST                 VkPrimitiveTopology = 3
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP                VkPrimitiveTopology = 4
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN                  VkPrimitiveTopology = 5
	VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY      VkPrimitiveTopology = 6
	VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY     VkPrimitiveTopology = 7
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY  VkPrimitiveTopology = 8
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY VkPrimitiveTopology = 9
	VK_PRIMITIVE_TOPOLOGY_PATCH_LIST                    VkPrimitiveTopology = 10
	VK_PRIMITIVE_TOPOLOGY_MAX_ENUM                      VkPrimitiveTopology = 0x7FFFFFFF
)

type VkPolygonMode = uint32

const (
	VK_POLYGON_MODE_FILL              VkPolygonMode = 0
	VK_POLYGON_MODE_LINE              VkPolygonMode = 1
	VK_POLYGON_MODE_POINT             VkPolygonMode = 2
	VK_POLYGON_MODE_FILL_RECTANGLE_NV VkPolygonMode = 1000153000
	VK_POLYGON_MODE_MAX_ENUM          VkPolygonMode = 0x7FFFFFFF
)

type VkStencilOp = uint32

const (
	VK_STENCIL_OP_KEEP                VkStencilOp = 0
	VK_STENCIL_OP_ZERO                VkStencilOp = 1
	VK_STENCIL_OP_REPLACE             VkStencilOp = 2
	VK_STENCIL_OP_INCREMENT_AND_CLAMP VkStencilOp = 3
	VK_STENCIL_OP_DECREMENT_AND_CLAMP VkStencilOp = 4
	VK_STENCIL_OP_INVERT              VkStencilOp = 5
	VK_STENCIL_OP_INCREMENT_AND_WRAP  VkStencilOp = 6
	VK_STENCIL_OP_DECREMENT_AND_WRAP  VkStencilOp = 7
	VK_STENCIL_OP_MAX_ENUM            VkStencilOp = 0x7FFFFFFF
)

type VkLogicOp = uint32

const (
	VK_LOGIC_OP_CLEAR         VkLogicOp = 0
	VK_LOGIC_OP_AND           VkLogicOp = 1
	VK_LOGIC_OP_AND_REVERSE   VkLogicOp = 2
	VK_LOGIC_OP_COPY          VkLogicOp = 3
	VK_LOGIC_OP_AND_INVERTED  VkLogicOp = 4
	VK_LOGIC_OP_NO_OP         VkLogicOp = 5
	VK_LOGIC_OP_XOR           VkLogicOp = 6
	VK_LOGIC_OP_OR            VkLogicOp = 7
	VK_LOGIC_OP_NOR           VkLogicOp = 8
	VK_LOGIC_OP_EQUIVALENT    VkLogicOp = 9
	VK_LOGIC_OP_INVERT        VkLogicOp = 10
	VK_LOGIC_OP_OR_REVERSE    VkLogicOp = 11
	VK_LOGIC_OP_COPY_INVERTED VkLogicOp = 12
	VK_LOGIC_OP_OR_INVERTED   VkLogicOp = 13
	VK_LOGIC_OP_NAND          VkLogicOp = 14
	VK_LOGIC_OP_SET           VkLogicOp = 15
	VK_LOGIC_OP_MAX_ENUM      VkLogicOp = 0x7FFFFFFF
)

type VkBorderColor = uint32

const (
	VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK VkBorderColor = 0
	VK_BORDER_COLOR_INT_TRANSPARENT_BLACK   VkBorderColor = 1
	VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK      VkBorderColor = 2
	VK_BORDER_COLOR_INT_OPAQUE_BLACK        VkBorderColor = 3
	VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE      VkBorderColor = 4
	VK_BORDER_COLOR_INT_OPAQUE_WHITE        VkBorderColor = 5
	VK_BORDER_COLOR_FLOAT_CUSTOM_EXT        VkBorderColor = 1000287003
	VK_BORDER_COLOR_INT_CUSTOM_EXT          VkBorderColor = 1000287004
	VK_BORDER_COLOR_MAX_ENUM                VkBorderColor = 0x7FFFFFFF
)

type VkFilter = uint32

const (
	VK_FILTER_NEAREST   VkFilter = 0
	VK_FILTER_LINEAR    VkFilter = 1
	VK_FILTER_CUBIC_IMG VkFilter = 1000015000
	VK_FILTER_CUBIC_EXT VkFilter = VK_FILTER_CUBIC_IMG
	VK_FILTER_MAX_ENUM  VkFilter = 0x7FFFFFFF
)

type VkSamplerAddressMode = uint32

const (
	VK_SAMPLER_ADDRESS_MODE_REPEAT                   VkSamplerAddressMode = 0
	VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT          VkSamplerAddressMode = 1
	VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE            VkSamplerAddressMode = 2
	VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER          VkSamplerAddressMode = 3
	VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE     VkSamplerAddressMode = 4
	VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR VkSamplerAddressMode = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
	VK_SAMPLER_ADDRESS_MODE_MAX_ENUM                 VkSamplerAddressMode = 0x7FFFFFFF
)

type VkSamplerMipmapMode = uint32

const (
	VK_SAMPLER_MIPMAP_MODE_NEAREST  VkSamplerMipmapMode = 0
	VK_SAMPLER_MIPMAP_MODE_LINEAR   VkSamplerMipmapMode = 1
	VK_SAMPLER_MIPMAP_MODE_MAX_ENUM VkSamplerMipmapMode = 0x7FFFFFFF
)

type VkDescriptorType = uint32

const (
	VK_DESCRIPTOR_TYPE_SAMPLER                    VkDescriptorType = 0
	VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER     VkDescriptorType = 1
	VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE              VkDescriptorType = 2
	VK_DESCRIPTOR_TYPE_STORAGE_IMAGE              VkDescriptorType = 3
	VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER       VkDescriptorType = 4
	VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER       VkDescriptorType = 5
	VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER             VkDescriptorType = 6
	VK_DESCRIPTOR_TYPE_STORAGE_BUFFER             VkDescriptorType = 7
	VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC     VkDescriptorType = 8
	VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC     VkDescriptorType = 9
	VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT           VkDescriptorType = 10
	VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT   VkDescriptorType = 1000138000
	VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR VkDescriptorType = 1000150000
	VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV  VkDescriptorType = 1000165000
	VK_DESCRIPTOR_TYPE_MUTABLE_VALVE              VkDescriptorType = 1000351000
	VK_DESCRIPTOR_TYPE_MAX_ENUM                   VkDescriptorType = 0x7FFFFFFF
)

type VkAttachmentLoadOp = uint32

const (
	VK_ATTACHMENT_LOAD_OP_LOAD      VkAttachmentLoadOp = 0
	VK_ATTACHMENT_LOAD_OP_CLEAR     VkAttachmentLoadOp = 1
	VK_ATTACHMENT_LOAD_OP_DONT_CARE VkAttachmentLoadOp = 2
	VK_ATTACHMENT_LOAD_OP_NONE_EXT  VkAttachmentLoadOp = 1000400000
	VK_ATTACHMENT_LOAD_OP_MAX_ENUM  VkAttachmentLoadOp = 0x7FFFFFFF
)

type VkAttachmentStoreOp = uint32

const (
	VK_ATTACHMENT_STORE_OP_STORE     VkAttachmentStoreOp = 0
	VK_ATTACHMENT_STORE_OP_DONT_CARE VkAttachmentStoreOp = 1
	VK_ATTACHMENT_STORE_OP_NONE_EXT  VkAttachmentStoreOp = 1000301000
	VK_ATTACHMENT_STORE_OP_NONE_QCOM VkAttachmentStoreOp = VK_ATTACHMENT_STORE_OP_NONE_EXT
	VK_ATTACHMENT_STORE_OP_MAX_ENUM  VkAttachmentStoreOp = 0x7FFFFFFF
)

type VkPipelineBindPoint = uint32

const (
	VK_PIPELINE_BIND_POINT_GRAPHICS               VkPipelineBindPoint = 0
	VK_PIPELINE_BIND_POINT_COMPUTE                VkPipelineBindPoint = 1
	VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR        VkPipelineBindPoint = 1000165000
	VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI VkPipelineBindPoint = 1000369003
	VK_PIPELINE_BIND_POINT_RAY_TRACING_NV         VkPipelineBindPoint = VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR
	VK_PIPELINE_BIND_POINT_MAX_ENUM               VkPipelineBindPoint = 0x7FFFFFFF
)

type VkCommandBufferLevel = uint32

const (
	VK_COMMAND_BUFFER_LEVEL_PRIMARY   VkCommandBufferLevel = 0
	VK_COMMAND_BUFFER_LEVEL_SECONDARY VkCommandBufferLevel = 1
	VK_COMMAND_BUFFER_LEVEL_MAX_ENUM  VkCommandBufferLevel = 0x7FFFFFFF
)

type VkIndexType = uint32

const (
	VK_INDEX_TYPE_UINT16    VkIndexType = 0
	VK_INDEX_TYPE_UINT32    VkIndexType = 1
	VK_INDEX_TYPE_NONE_KHR  VkIndexType = 1000165000
	VK_INDEX_TYPE_UINT8_EXT VkIndexType = 1000265000
	VK_INDEX_TYPE_NONE_NV   VkIndexType = VK_INDEX_TYPE_NONE_KHR
	VK_INDEX_TYPE_MAX_ENUM  VkIndexType = 0x7FFFFFFF
)

type VkSubpassContents = uint32

const (
	VK_SUBPASS_CONTENTS_INLINE                    VkSubpassContents = 0
	VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS VkSubpassContents = 1
	VK_SUBPASS_CONTENTS_MAX_ENUM                  VkSubpassContents = 0x7FFFFFFF
)

type VkAccessFlagBits = uint32

const (
	VK_ACCESS_INDIRECT_COMMAND_READ_BIT                     VkAccessFlagBits = 0x00000001
	VK_ACCESS_INDEX_READ_BIT                                VkAccessFlagBits = 0x00000002
	VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT                     VkAccessFlagBits = 0x00000004
	VK_ACCESS_UNIFORM_READ_BIT                              VkAccessFlagBits = 0x00000008
	VK_ACCESS_INPUT_ATTACHMENT_READ_BIT                     VkAccessFlagBits = 0x00000010
	VK_ACCESS_SHADER_READ_BIT                               VkAccessFlagBits = 0x00000020
	VK_ACCESS_SHADER_WRITE_BIT                              VkAccessFlagBits = 0x00000040
	VK_ACCESS_COLOR_ATTACHMENT_READ_BIT                     VkAccessFlagBits = 0x00000080
	VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT                    VkAccessFlagBits = 0x00000100
	VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT             VkAccessFlagBits = 0x00000200
	VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT            VkAccessFlagBits = 0x00000400
	VK_ACCESS_TRANSFER_READ_BIT                             VkAccessFlagBits = 0x00000800
	VK_ACCESS_TRANSFER_WRITE_BIT                            VkAccessFlagBits = 0x00001000
	VK_ACCESS_HOST_READ_BIT                                 VkAccessFlagBits = 0x00002000
	VK_ACCESS_HOST_WRITE_BIT                                VkAccessFlagBits = 0x00004000
	VK_ACCESS_MEMORY_READ_BIT                               VkAccessFlagBits = 0x00008000
	VK_ACCESS_MEMORY_WRITE_BIT                              VkAccessFlagBits = 0x00010000
	VK_ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT              VkAccessFlagBits = 0x02000000
	VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT       VkAccessFlagBits = 0x04000000
	VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT      VkAccessFlagBits = 0x08000000
	VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT            VkAccessFlagBits = 0x00100000
	VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT     VkAccessFlagBits = 0x00080000
	VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR           VkAccessFlagBits = 0x00200000
	VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR          VkAccessFlagBits = 0x00400000
	VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT             VkAccessFlagBits = 0x01000000
	VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR VkAccessFlagBits = 0x00800000
	VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV                VkAccessFlagBits = 0x00020000
	VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_NV               VkAccessFlagBits = 0x00040000
	VK_ACCESS_NONE_KHR                                      VkAccessFlagBits = 0
	VK_ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV                VkAccessFlagBits = VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR
	VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV            VkAccessFlagBits = VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR
	VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV           VkAccessFlagBits = VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR
	VK_ACCESS_FLAG_BITS_MAX_ENUM                            VkAccessFlagBits = 0x7FFFFFFF
)

type VkImageAspectFlagBits = uint32

const (
	VK_IMAGE_ASPECT_COLOR_BIT              VkImageAspectFlagBits = 0x00000001
	VK_IMAGE_ASPECT_DEPTH_BIT              VkImageAspectFlagBits = 0x00000002
	VK_IMAGE_ASPECT_STENCIL_BIT            VkImageAspectFlagBits = 0x00000004
	VK_IMAGE_ASPECT_METADATA_BIT           VkImageAspectFlagBits = 0x00000008
	VK_IMAGE_ASPECT_PLANE_0_BIT            VkImageAspectFlagBits = 0x00000010
	VK_IMAGE_ASPECT_PLANE_1_BIT            VkImageAspectFlagBits = 0x00000020
	VK_IMAGE_ASPECT_PLANE_2_BIT            VkImageAspectFlagBits = 0x00000040
	VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT VkImageAspectFlagBits = 0x00000080
	VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT VkImageAspectFlagBits = 0x00000100
	VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT VkImageAspectFlagBits = 0x00000200
	VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT VkImageAspectFlagBits = 0x00000400
	VK_IMAGE_ASPECT_PLANE_0_BIT_KHR        VkImageAspectFlagBits = VK_IMAGE_ASPECT_PLANE_0_BIT
	VK_IMAGE_ASPECT_PLANE_1_BIT_KHR        VkImageAspectFlagBits = VK_IMAGE_ASPECT_PLANE_1_BIT
	VK_IMAGE_ASPECT_PLANE_2_BIT_KHR        VkImageAspectFlagBits = VK_IMAGE_ASPECT_PLANE_2_BIT
	VK_IMAGE_ASPECT_FLAG_BITS_MAX_ENUM     VkImageAspectFlagBits = 0x7FFFFFFF
)

type VkFormatFeatureFlagBits = uint32

const (
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT                                                               VkFormatFeatureFlagBits = 0x00000001
	VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT                                                               VkFormatFeatureFlagBits = 0x00000002
	VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT                                                        VkFormatFeatureFlagBits = 0x00000004
	VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT                                                        VkFormatFeatureFlagBits = 0x00000008
	VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT                                                        VkFormatFeatureFlagBits = 0x00000010
	VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT                                                 VkFormatFeatureFlagBits = 0x00000020
	VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT                                                               VkFormatFeatureFlagBits = 0x00000040
	VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT                                                            VkFormatFeatureFlagBits = 0x00000080
	VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT                                                      VkFormatFeatureFlagBits = 0x00000100
	VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT                                                    VkFormatFeatureFlagBits = 0x00000200
	VK_FORMAT_FEATURE_BLIT_SRC_BIT                                                                    VkFormatFeatureFlagBits = 0x00000400
	VK_FORMAT_FEATURE_BLIT_DST_BIT                                                                    VkFormatFeatureFlagBits = 0x00000800
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT                                                 VkFormatFeatureFlagBits = 0x00001000
	VK_FORMAT_FEATURE_TRANSFER_SRC_BIT                                                                VkFormatFeatureFlagBits = 0x00004000
	VK_FORMAT_FEATURE_TRANSFER_DST_BIT                                                                VkFormatFeatureFlagBits = 0x00008000
	VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT                                                     VkFormatFeatureFlagBits = 0x00020000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT                                VkFormatFeatureFlagBits = 0x00040000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT               VkFormatFeatureFlagBits = 0x00080000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT               VkFormatFeatureFlagBits = 0x00100000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT     VkFormatFeatureFlagBits = 0x00200000
	VK_FORMAT_FEATURE_DISJOINT_BIT                                                                    VkFormatFeatureFlagBits = 0x00400000
	VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT                                                      VkFormatFeatureFlagBits = 0x00800000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT                                                 VkFormatFeatureFlagBits = 0x00010000
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG                                              VkFormatFeatureFlagBits = 0x00002000
	VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR                                                     VkFormatFeatureFlagBits = 0x02000000
	VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR                                                        VkFormatFeatureFlagBits = 0x04000000
	VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR                                    VkFormatFeatureFlagBits = 0x20000000
	VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT                                                    VkFormatFeatureFlagBits = 0x01000000
	VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR                                        VkFormatFeatureFlagBits = 0x40000000
	VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR                                                      VkFormatFeatureFlagBits = 0x08000000
	VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR                                                        VkFormatFeatureFlagBits = 0x10000000
	VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR                                                            VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_TRANSFER_SRC_BIT
	VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR                                                            VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_TRANSFER_DST_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT                                             VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT
	VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR                                                 VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR                            VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR           VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR           VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	VK_FORMAT_FEATURE_DISJOINT_BIT_KHR                                                                VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_DISJOINT_BIT
	VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR                                                  VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT                                              VkFormatFeatureFlagBits = VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG
	VK_FORMAT_FEATURE_FLAG_BITS_MAX_ENUM                                                              VkFormatFeatureFlagBits = 0x7FFFFFFF
)

type VkImageCreateFlagBits = uint32

const (
	VK_IMAGE_CREATE_SPARSE_BINDING_BIT                        VkImageCreateFlagBits = 0x00000001
	VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT                      VkImageCreateFlagBits = 0x00000002
	VK_IMAGE_CREATE_SPARSE_ALIASED_BIT                        VkImageCreateFlagBits = 0x00000004
	VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT                        VkImageCreateFlagBits = 0x00000008
	VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT                       VkImageCreateFlagBits = 0x00000010
	VK_IMAGE_CREATE_ALIAS_BIT                                 VkImageCreateFlagBits = 0x00000400
	VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT           VkImageCreateFlagBits = 0x00000040
	VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT                   VkImageCreateFlagBits = 0x00000020
	VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT           VkImageCreateFlagBits = 0x00000080
	VK_IMAGE_CREATE_EXTENDED_USAGE_BIT                        VkImageCreateFlagBits = 0x00000100
	VK_IMAGE_CREATE_PROTECTED_BIT                             VkImageCreateFlagBits = 0x00000800
	VK_IMAGE_CREATE_DISJOINT_BIT                              VkImageCreateFlagBits = 0x00000200
	VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV                     VkImageCreateFlagBits = 0x00002000
	VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT VkImageCreateFlagBits = 0x00001000
	VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT                        VkImageCreateFlagBits = 0x00004000
	VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR       VkImageCreateFlagBits = VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT
	VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR               VkImageCreateFlagBits = VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR       VkImageCreateFlagBits = VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR                    VkImageCreateFlagBits = VK_IMAGE_CREATE_EXTENDED_USAGE_BIT
	VK_IMAGE_CREATE_DISJOINT_BIT_KHR                          VkImageCreateFlagBits = VK_IMAGE_CREATE_DISJOINT_BIT
	VK_IMAGE_CREATE_ALIAS_BIT_KHR                             VkImageCreateFlagBits = VK_IMAGE_CREATE_ALIAS_BIT
	VK_IMAGE_CREATE_FLAG_BITS_MAX_ENUM                        VkImageCreateFlagBits = 0x7FFFFFFF
)

type VkSampleCountFlagBits = uint32

const (
	VK_SAMPLE_COUNT_1_BIT              VkSampleCountFlagBits = 0x00000001
	VK_SAMPLE_COUNT_2_BIT              VkSampleCountFlagBits = 0x00000002
	VK_SAMPLE_COUNT_4_BIT              VkSampleCountFlagBits = 0x00000004
	VK_SAMPLE_COUNT_8_BIT              VkSampleCountFlagBits = 0x00000008
	VK_SAMPLE_COUNT_16_BIT             VkSampleCountFlagBits = 0x00000010
	VK_SAMPLE_COUNT_32_BIT             VkSampleCountFlagBits = 0x00000020
	VK_SAMPLE_COUNT_64_BIT             VkSampleCountFlagBits = 0x00000040
	VK_SAMPLE_COUNT_FLAG_BITS_MAX_ENUM VkSampleCountFlagBits = 0x7FFFFFFF
)

type VkImageUsageFlagBits = uint32

const (
	VK_IMAGE_USAGE_TRANSFER_SRC_BIT                         VkImageUsageFlagBits = 0x00000001
	VK_IMAGE_USAGE_TRANSFER_DST_BIT                         VkImageUsageFlagBits = 0x00000002
	VK_IMAGE_USAGE_SAMPLED_BIT                              VkImageUsageFlagBits = 0x00000004
	VK_IMAGE_USAGE_STORAGE_BIT                              VkImageUsageFlagBits = 0x00000008
	VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT                     VkImageUsageFlagBits = 0x00000010
	VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT             VkImageUsageFlagBits = 0x00000020
	VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT                 VkImageUsageFlagBits = 0x00000040
	VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT                     VkImageUsageFlagBits = 0x00000080
	VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR                 VkImageUsageFlagBits = 0x00000400
	VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR                 VkImageUsageFlagBits = 0x00000800
	VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR                 VkImageUsageFlagBits = 0x00001000
	VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT             VkImageUsageFlagBits = 0x00000200
	VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR VkImageUsageFlagBits = 0x00000100
	VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR                 VkImageUsageFlagBits = 0x00002000
	VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR                 VkImageUsageFlagBits = 0x00004000
	VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR                 VkImageUsageFlagBits = 0x00008000
	VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI               VkImageUsageFlagBits = 0x00040000
	VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV                VkImageUsageFlagBits = VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	VK_IMAGE_USAGE_FLAG_BITS_MAX_ENUM                       VkImageUsageFlagBits = 0x7FFFFFFF
)

type VkMemoryHeapFlagBits = uint32

const (
	VK_MEMORY_HEAP_DEVICE_LOCAL_BIT       VkMemoryHeapFlagBits = 0x00000001
	VK_MEMORY_HEAP_MULTI_INSTANCE_BIT     VkMemoryHeapFlagBits = 0x00000002
	VK_MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR VkMemoryHeapFlagBits = VK_MEMORY_HEAP_MULTI_INSTANCE_BIT
	VK_MEMORY_HEAP_FLAG_BITS_MAX_ENUM     VkMemoryHeapFlagBits = 0x7FFFFFFF
)

type VkMemoryPropertyFlagBits = uint32

const (
	VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT        VkMemoryPropertyFlagBits = 0x00000001
	VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT        VkMemoryPropertyFlagBits = 0x00000002
	VK_MEMORY_PROPERTY_HOST_COHERENT_BIT       VkMemoryPropertyFlagBits = 0x00000004
	VK_MEMORY_PROPERTY_HOST_CACHED_BIT         VkMemoryPropertyFlagBits = 0x00000008
	VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT    VkMemoryPropertyFlagBits = 0x00000010
	VK_MEMORY_PROPERTY_PROTECTED_BIT           VkMemoryPropertyFlagBits = 0x00000020
	VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD VkMemoryPropertyFlagBits = 0x00000040
	VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD VkMemoryPropertyFlagBits = 0x00000080
	VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV     VkMemoryPropertyFlagBits = 0x00000100
	VK_MEMORY_PROPERTY_FLAG_BITS_MAX_ENUM      VkMemoryPropertyFlagBits = 0x7FFFFFFF
)

type VkQueueFlagBits = uint32

const (
	VK_QUEUE_GRAPHICS_BIT         VkQueueFlagBits = 0x00000001
	VK_QUEUE_COMPUTE_BIT          VkQueueFlagBits = 0x00000002
	VK_QUEUE_TRANSFER_BIT         VkQueueFlagBits = 0x00000004
	VK_QUEUE_SPARSE_BINDING_BIT   VkQueueFlagBits = 0x00000008
	VK_QUEUE_PROTECTED_BIT        VkQueueFlagBits = 0x00000010
	VK_QUEUE_VIDEO_DECODE_BIT_KHR VkQueueFlagBits = 0x00000020
	VK_QUEUE_VIDEO_ENCODE_BIT_KHR VkQueueFlagBits = 0x00000040
	VK_QUEUE_FLAG_BITS_MAX_ENUM   VkQueueFlagBits = 0x7FFFFFFF
)

type VkDeviceQueueCreateFlagBits = uint32

const (
	VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT      VkDeviceQueueCreateFlagBits = 0x00000001
	VK_DEVICE_QUEUE_CREATE_FLAG_BITS_MAX_ENUM VkDeviceQueueCreateFlagBits = 0x7FFFFFFF
)

type VkPipelineStageFlagBits = uint32

const (
	VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT                          VkPipelineStageFlagBits = 0x00000001
	VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT                        VkPipelineStageFlagBits = 0x00000002
	VK_PIPELINE_STAGE_VERTEX_INPUT_BIT                         VkPipelineStageFlagBits = 0x00000004
	VK_PIPELINE_STAGE_VERTEX_SHADER_BIT                        VkPipelineStageFlagBits = 0x00000008
	VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT          VkPipelineStageFlagBits = 0x00000010
	VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT       VkPipelineStageFlagBits = 0x00000020
	VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT                      VkPipelineStageFlagBits = 0x00000040
	VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT                      VkPipelineStageFlagBits = 0x00000080
	VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT                 VkPipelineStageFlagBits = 0x00000100
	VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT                  VkPipelineStageFlagBits = 0x00000200
	VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT              VkPipelineStageFlagBits = 0x00000400
	VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT                       VkPipelineStageFlagBits = 0x00000800
	VK_PIPELINE_STAGE_TRANSFER_BIT                             VkPipelineStageFlagBits = 0x00001000
	VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT                       VkPipelineStageFlagBits = 0x00002000
	VK_PIPELINE_STAGE_HOST_BIT                                 VkPipelineStageFlagBits = 0x00004000
	VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT                         VkPipelineStageFlagBits = 0x00008000
	VK_PIPELINE_STAGE_ALL_COMMANDS_BIT                         VkPipelineStageFlagBits = 0x00010000
	VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT               VkPipelineStageFlagBits = 0x01000000
	VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT            VkPipelineStageFlagBits = 0x00040000
	VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR     VkPipelineStageFlagBits = 0x02000000
	VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR               VkPipelineStageFlagBits = 0x00200000
	VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV                       VkPipelineStageFlagBits = 0x00080000
	VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV                       VkPipelineStageFlagBits = 0x00100000
	VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT         VkPipelineStageFlagBits = 0x00800000
	VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR VkPipelineStageFlagBits = 0x00400000
	VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV                VkPipelineStageFlagBits = 0x00020000
	VK_PIPELINE_STAGE_NONE_KHR                                 VkPipelineStageFlagBits = 0
	VK_PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV                VkPipelineStageFlagBits = VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV                VkPipelineStageFlagBits = VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR
	VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV      VkPipelineStageFlagBits = VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR
	VK_PIPELINE_STAGE_FLAG_BITS_MAX_ENUM                       VkPipelineStageFlagBits = 0x7FFFFFFF
)

type VkSparseMemoryBindFlagBits = uint32

const (
	VK_SPARSE_MEMORY_BIND_METADATA_BIT       VkSparseMemoryBindFlagBits = 0x00000001
	VK_SPARSE_MEMORY_BIND_FLAG_BITS_MAX_ENUM VkSparseMemoryBindFlagBits = 0x7FFFFFFF
)

type VkSparseImageFormatFlagBits = uint32

const (
	VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT         VkSparseImageFormatFlagBits = 0x00000001
	VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT       VkSparseImageFormatFlagBits = 0x00000002
	VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT VkSparseImageFormatFlagBits = 0x00000004
	VK_SPARSE_IMAGE_FORMAT_FLAG_BITS_MAX_ENUM         VkSparseImageFormatFlagBits = 0x7FFFFFFF
)

type VkFenceCreateFlagBits = uint32

const (
	VK_FENCE_CREATE_SIGNALED_BIT       VkFenceCreateFlagBits = 0x00000001
	VK_FENCE_CREATE_FLAG_BITS_MAX_ENUM VkFenceCreateFlagBits = 0x7FFFFFFF
)

type VkEventCreateFlagBits = uint32

const (
	VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR VkEventCreateFlagBits = 0x00000001
	VK_EVENT_CREATE_FLAG_BITS_MAX_ENUM  VkEventCreateFlagBits = 0x7FFFFFFF
)

type VkQueryPipelineStatisticFlagBits = uint32

const (
	VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT                    VkQueryPipelineStatisticFlagBits = 0x00000001
	VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT                  VkQueryPipelineStatisticFlagBits = 0x00000002
	VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT                  VkQueryPipelineStatisticFlagBits = 0x00000004
	VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT                VkQueryPipelineStatisticFlagBits = 0x00000008
	VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT                 VkQueryPipelineStatisticFlagBits = 0x00000010
	VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT                       VkQueryPipelineStatisticFlagBits = 0x00000020
	VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT                        VkQueryPipelineStatisticFlagBits = 0x00000040
	VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT                VkQueryPipelineStatisticFlagBits = 0x00000080
	VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT        VkQueryPipelineStatisticFlagBits = 0x00000100
	VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x00000200
	VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT                 VkQueryPipelineStatisticFlagBits = 0x00000400
	VK_QUERY_PIPELINE_STATISTIC_FLAG_BITS_MAX_ENUM                             VkQueryPipelineStatisticFlagBits = 0x7FFFFFFF
)

type VkQueryResultFlagBits = uint32

const (
	VK_QUERY_RESULT_64_BIT                VkQueryResultFlagBits = 0x00000001
	VK_QUERY_RESULT_WAIT_BIT              VkQueryResultFlagBits = 0x00000002
	VK_QUERY_RESULT_WITH_AVAILABILITY_BIT VkQueryResultFlagBits = 0x00000004
	VK_QUERY_RESULT_PARTIAL_BIT           VkQueryResultFlagBits = 0x00000008
	VK_QUERY_RESULT_WITH_STATUS_BIT_KHR   VkQueryResultFlagBits = 0x00000010
	VK_QUERY_RESULT_FLAG_BITS_MAX_ENUM    VkQueryResultFlagBits = 0x7FFFFFFF
)

type VkBufferCreateFlagBits = uint32

const (
	VK_BUFFER_CREATE_SPARSE_BINDING_BIT                    VkBufferCreateFlagBits = 0x00000001
	VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT                  VkBufferCreateFlagBits = 0x00000002
	VK_BUFFER_CREATE_SPARSE_ALIASED_BIT                    VkBufferCreateFlagBits = 0x00000004
	VK_BUFFER_CREATE_PROTECTED_BIT                         VkBufferCreateFlagBits = 0x00000008
	VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT     VkBufferCreateFlagBits = 0x00000010
	VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT VkBufferCreateFlagBits = VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT
	VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR VkBufferCreateFlagBits = VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT
	VK_BUFFER_CREATE_FLAG_BITS_MAX_ENUM                    VkBufferCreateFlagBits = 0x7FFFFFFF
)

type VkBufferUsageFlagBits = uint32

const (
	VK_BUFFER_USAGE_TRANSFER_SRC_BIT                                     VkBufferUsageFlagBits = 0x00000001
	VK_BUFFER_USAGE_TRANSFER_DST_BIT                                     VkBufferUsageFlagBits = 0x00000002
	VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT                             VkBufferUsageFlagBits = 0x00000004
	VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT                             VkBufferUsageFlagBits = 0x00000008
	VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT                                   VkBufferUsageFlagBits = 0x00000010
	VK_BUFFER_USAGE_STORAGE_BUFFER_BIT                                   VkBufferUsageFlagBits = 0x00000020
	VK_BUFFER_USAGE_INDEX_BUFFER_BIT                                     VkBufferUsageFlagBits = 0x00000040
	VK_BUFFER_USAGE_VERTEX_BUFFER_BIT                                    VkBufferUsageFlagBits = 0x00000080
	VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT                                  VkBufferUsageFlagBits = 0x00000100
	VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT                            VkBufferUsageFlagBits = 0x00020000
	VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR                             VkBufferUsageFlagBits = 0x00002000
	VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR                             VkBufferUsageFlagBits = 0x00004000
	VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT                    VkBufferUsageFlagBits = 0x00000800
	VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT            VkBufferUsageFlagBits = 0x00001000
	VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT                        VkBufferUsageFlagBits = 0x00000200
	VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR VkBufferUsageFlagBits = 0x00080000
	VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR               VkBufferUsageFlagBits = 0x00100000
	VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR                         VkBufferUsageFlagBits = 0x00000400
	VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR                             VkBufferUsageFlagBits = 0x00008000
	VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR                             VkBufferUsageFlagBits = 0x00010000
	VK_BUFFER_USAGE_RAY_TRACING_BIT_NV                                   VkBufferUsageFlagBits = VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR
	VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT                        VkBufferUsageFlagBits = VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT
	VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR                        VkBufferUsageFlagBits = VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT
	VK_BUFFER_USAGE_FLAG_BITS_MAX_ENUM                                   VkBufferUsageFlagBits = 0x7FFFFFFF
)

type VkImageViewCreateFlagBits = uint32

const (
	VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT  VkImageViewCreateFlagBits = 0x00000001
	VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT VkImageViewCreateFlagBits = 0x00000002
	VK_IMAGE_VIEW_CREATE_FLAG_BITS_MAX_ENUM                    VkImageViewCreateFlagBits = 0x7FFFFFFF
)

type VkShaderModuleCreateFlagBits = uint32

const (
	VK_SHADER_MODULE_CREATE_FLAG_BITS_MAX_ENUM VkShaderModuleCreateFlagBits = 0x7FFFFFFF
)

type VkPipelineCacheCreateFlagBits = uint32

const (
	VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT VkPipelineCacheCreateFlagBits = 0x00000001
	VK_PIPELINE_CACHE_CREATE_FLAG_BITS_MAX_ENUM              VkPipelineCacheCreateFlagBits = 0x7FFFFFFF
)

type VkColorComponentFlagBits = uint32

const (
	VK_COLOR_COMPONENT_R_BIT              VkColorComponentFlagBits = 0x00000001
	VK_COLOR_COMPONENT_G_BIT              VkColorComponentFlagBits = 0x00000002
	VK_COLOR_COMPONENT_B_BIT              VkColorComponentFlagBits = 0x00000004
	VK_COLOR_COMPONENT_A_BIT              VkColorComponentFlagBits = 0x00000008
	VK_COLOR_COMPONENT_FLAG_BITS_MAX_ENUM VkColorComponentFlagBits = 0x7FFFFFFF
)

type VkPipelineCreateFlagBits = uint32

const (
	VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT                               VkPipelineCreateFlagBits = 0x00000001
	VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT                                  VkPipelineCreateFlagBits = 0x00000002
	VK_PIPELINE_CREATE_DERIVATIVE_BIT                                         VkPipelineCreateFlagBits = 0x00000004
	VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT                       VkPipelineCreateFlagBits = 0x00000008
	VK_PIPELINE_CREATE_DISPATCH_BASE_BIT                                      VkPipelineCreateFlagBits = 0x00000010
	VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR            VkPipelineCreateFlagBits = 0x00004000
	VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR        VkPipelineCreateFlagBits = 0x00008000
	VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR               VkPipelineCreateFlagBits = 0x00010000
	VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR       VkPipelineCreateFlagBits = 0x00020000
	VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR                     VkPipelineCreateFlagBits = 0x00001000
	VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR                         VkPipelineCreateFlagBits = 0x00002000
	VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR VkPipelineCreateFlagBits = 0x00080000
	VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV                                   VkPipelineCreateFlagBits = 0x00000020
	VK_PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR                             VkPipelineCreateFlagBits = 0x00000040
	VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR               VkPipelineCreateFlagBits = 0x00000080
	VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV                               VkPipelineCreateFlagBits = 0x00040000
	VK_PIPELINE_CREATE_LIBRARY_BIT_KHR                                        VkPipelineCreateFlagBits = 0x00000800
	VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT              VkPipelineCreateFlagBits = 0x00000100
	VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT_EXT                        VkPipelineCreateFlagBits = 0x00000200
	VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV                        VkPipelineCreateFlagBits = 0x00100000
	VK_PIPELINE_CREATE_DISPATCH_BASE                                          VkPipelineCreateFlagBits = VK_PIPELINE_CREATE_DISPATCH_BASE_BIT
	VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR                   VkPipelineCreateFlagBits = VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
	VK_PIPELINE_CREATE_DISPATCH_BASE_KHR                                      VkPipelineCreateFlagBits = VK_PIPELINE_CREATE_DISPATCH_BASE
	VK_PIPELINE_CREATE_FLAG_BITS_MAX_ENUM                                     VkPipelineCreateFlagBits = 0x7FFFFFFF
)

type VkPipelineShaderStageCreateFlagBits = uint32

const (
	VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT VkPipelineShaderStageCreateFlagBits = 0x00000001
	VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT      VkPipelineShaderStageCreateFlagBits = 0x00000002
	VK_PIPELINE_SHADER_STAGE_CREATE_FLAG_BITS_MAX_ENUM                  VkPipelineShaderStageCreateFlagBits = 0x7FFFFFFF
)

type VkShaderStageFlagBits = uint32

const (
	VK_SHADER_STAGE_VERTEX_BIT                  VkShaderStageFlagBits = 0x00000001
	VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT    VkShaderStageFlagBits = 0x00000002
	VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT VkShaderStageFlagBits = 0x00000004
	VK_SHADER_STAGE_GEOMETRY_BIT                VkShaderStageFlagBits = 0x00000008
	VK_SHADER_STAGE_FRAGMENT_BIT                VkShaderStageFlagBits = 0x00000010
	VK_SHADER_STAGE_COMPUTE_BIT                 VkShaderStageFlagBits = 0x00000020
	VK_SHADER_STAGE_ALL_GRAPHICS                VkShaderStageFlagBits = 0x0000001F
	VK_SHADER_STAGE_ALL                         VkShaderStageFlagBits = 0x7FFFFFFF
	VK_SHADER_STAGE_RAYGEN_BIT_KHR              VkShaderStageFlagBits = 0x00000100
	VK_SHADER_STAGE_ANY_HIT_BIT_KHR             VkShaderStageFlagBits = 0x00000200
	VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR         VkShaderStageFlagBits = 0x00000400
	VK_SHADER_STAGE_MISS_BIT_KHR                VkShaderStageFlagBits = 0x00000800
	VK_SHADER_STAGE_INTERSECTION_BIT_KHR        VkShaderStageFlagBits = 0x00001000
	VK_SHADER_STAGE_CALLABLE_BIT_KHR            VkShaderStageFlagBits = 0x00002000
	VK_SHADER_STAGE_TASK_BIT_NV                 VkShaderStageFlagBits = 0x00000040
	VK_SHADER_STAGE_MESH_BIT_NV                 VkShaderStageFlagBits = 0x00000080
	VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI  VkShaderStageFlagBits = 0x00004000
	VK_SHADER_STAGE_RAYGEN_BIT_NV               VkShaderStageFlagBits = VK_SHADER_STAGE_RAYGEN_BIT_KHR
	VK_SHADER_STAGE_ANY_HIT_BIT_NV              VkShaderStageFlagBits = VK_SHADER_STAGE_ANY_HIT_BIT_KHR
	VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV          VkShaderStageFlagBits = VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR
	VK_SHADER_STAGE_MISS_BIT_NV                 VkShaderStageFlagBits = VK_SHADER_STAGE_MISS_BIT_KHR
	VK_SHADER_STAGE_INTERSECTION_BIT_NV         VkShaderStageFlagBits = VK_SHADER_STAGE_INTERSECTION_BIT_KHR
	VK_SHADER_STAGE_CALLABLE_BIT_NV             VkShaderStageFlagBits = VK_SHADER_STAGE_CALLABLE_BIT_KHR
	VK_SHADER_STAGE_FLAG_BITS_MAX_ENUM          VkShaderStageFlagBits = 0x7FFFFFFF
)

type VkCullModeFlagBits = uint32

const (
	VK_CULL_MODE_NONE               VkCullModeFlagBits = 0
	VK_CULL_MODE_FRONT_BIT          VkCullModeFlagBits = 0x00000001
	VK_CULL_MODE_BACK_BIT           VkCullModeFlagBits = 0x00000002
	VK_CULL_MODE_FRONT_AND_BACK     VkCullModeFlagBits = 0x00000003
	VK_CULL_MODE_FLAG_BITS_MAX_ENUM VkCullModeFlagBits = 0x7FFFFFFF
)

type VkSamplerCreateFlagBits = uint32

const (
	VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT                       VkSamplerCreateFlagBits = 0x00000001
	VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT VkSamplerCreateFlagBits = 0x00000002
	VK_SAMPLER_CREATE_FLAG_BITS_MAX_ENUM                       VkSamplerCreateFlagBits = 0x7FFFFFFF
)

type VkDescriptorPoolCreateFlagBits = uint32

const (
	VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT   VkDescriptorPoolCreateFlagBits = 0x00000001
	VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT     VkDescriptorPoolCreateFlagBits = 0x00000002
	VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE       VkDescriptorPoolCreateFlagBits = 0x00000004
	VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT VkDescriptorPoolCreateFlagBits = VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT
	VK_DESCRIPTOR_POOL_CREATE_FLAG_BITS_MAX_ENUM        VkDescriptorPoolCreateFlagBits = 0x7FFFFFFF
)

type VkDescriptorSetLayoutCreateFlagBits = uint32

const (
	VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT     VkDescriptorSetLayoutCreateFlagBits = 0x00000002
	VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR        VkDescriptorSetLayoutCreateFlagBits = 0x00000001
	VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE       VkDescriptorSetLayoutCreateFlagBits = 0x00000004
	VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT VkDescriptorSetLayoutCreateFlagBits = VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT
	VK_DESCRIPTOR_SET_LAYOUT_CREATE_FLAG_BITS_MAX_ENUM             VkDescriptorSetLayoutCreateFlagBits = 0x7FFFFFFF
)

type VkAttachmentDescriptionFlagBits = uint32

const (
	VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT      VkAttachmentDescriptionFlagBits = 0x00000001
	VK_ATTACHMENT_DESCRIPTION_FLAG_BITS_MAX_ENUM VkAttachmentDescriptionFlagBits = 0x7FFFFFFF
)

type VkDependencyFlagBits = uint32

const (
	VK_DEPENDENCY_BY_REGION_BIT        VkDependencyFlagBits = 0x00000001
	VK_DEPENDENCY_DEVICE_GROUP_BIT     VkDependencyFlagBits = 0x00000004
	VK_DEPENDENCY_VIEW_LOCAL_BIT       VkDependencyFlagBits = 0x00000002
	VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR   VkDependencyFlagBits = VK_DEPENDENCY_VIEW_LOCAL_BIT
	VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR VkDependencyFlagBits = VK_DEPENDENCY_DEVICE_GROUP_BIT
	VK_DEPENDENCY_FLAG_BITS_MAX_ENUM   VkDependencyFlagBits = 0x7FFFFFFF
)

type VkFramebufferCreateFlagBits = uint32

const (
	VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT      VkFramebufferCreateFlagBits = 0x00000001
	VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR  VkFramebufferCreateFlagBits = VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT
	VK_FRAMEBUFFER_CREATE_FLAG_BITS_MAX_ENUM VkFramebufferCreateFlagBits = 0x7FFFFFFF
)

type VkRenderPassCreateFlagBits = uint32

const (
	VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM VkRenderPassCreateFlagBits = 0x00000002
	VK_RENDER_PASS_CREATE_FLAG_BITS_MAX_ENUM VkRenderPassCreateFlagBits = 0x7FFFFFFF
)

type VkSubpassDescriptionFlagBits = uint32

const (
	VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX      VkSubpassDescriptionFlagBits = 0x00000001
	VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX VkSubpassDescriptionFlagBits = 0x00000002
	VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM         VkSubpassDescriptionFlagBits = 0x00000004
	VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM          VkSubpassDescriptionFlagBits = 0x00000008
	VK_SUBPASS_DESCRIPTION_FLAG_BITS_MAX_ENUM               VkSubpassDescriptionFlagBits = 0x7FFFFFFF
)

type VkCommandPoolCreateFlagBits = uint32

const (
	VK_COMMAND_POOL_CREATE_TRANSIENT_BIT            VkCommandPoolCreateFlagBits = 0x00000001
	VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT VkCommandPoolCreateFlagBits = 0x00000002
	VK_COMMAND_POOL_CREATE_PROTECTED_BIT            VkCommandPoolCreateFlagBits = 0x00000004
	VK_COMMAND_POOL_CREATE_FLAG_BITS_MAX_ENUM       VkCommandPoolCreateFlagBits = 0x7FFFFFFF
)

type VkCommandPoolResetFlagBits = uint32

const (
	VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT VkCommandPoolResetFlagBits = 0x00000001
	VK_COMMAND_POOL_RESET_FLAG_BITS_MAX_ENUM    VkCommandPoolResetFlagBits = 0x7FFFFFFF
)

type VkCommandBufferUsageFlagBits = uint32

const (
	VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT      VkCommandBufferUsageFlagBits = 0x00000001
	VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT VkCommandBufferUsageFlagBits = 0x00000002
	VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT     VkCommandBufferUsageFlagBits = 0x00000004
	VK_COMMAND_BUFFER_USAGE_FLAG_BITS_MAX_ENUM       VkCommandBufferUsageFlagBits = 0x7FFFFFFF
)

type VkQueryControlFlagBits = uint32

const (
	VK_QUERY_CONTROL_PRECISE_BIT        VkQueryControlFlagBits = 0x00000001
	VK_QUERY_CONTROL_FLAG_BITS_MAX_ENUM VkQueryControlFlagBits = 0x7FFFFFFF
)

type VkCommandBufferResetFlagBits = uint32

const (
	VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT VkCommandBufferResetFlagBits = 0x00000001
	VK_COMMAND_BUFFER_RESET_FLAG_BITS_MAX_ENUM    VkCommandBufferResetFlagBits = 0x7FFFFFFF
)

type VkStencilFaceFlagBits = uint32

const (
	VK_STENCIL_FACE_FRONT_BIT          VkStencilFaceFlagBits = 0x00000001
	VK_STENCIL_FACE_BACK_BIT           VkStencilFaceFlagBits = 0x00000002
	VK_STENCIL_FACE_FRONT_AND_BACK     VkStencilFaceFlagBits = 0x00000003
	VK_STENCIL_FRONT_AND_BACK          VkStencilFaceFlagBits = VK_STENCIL_FACE_FRONT_AND_BACK
	VK_STENCIL_FACE_FLAG_BITS_MAX_ENUM VkStencilFaceFlagBits = 0x7FFFFFFF
)

type VkPointClippingBehavior = uint32

const (
	VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES           VkPointClippingBehavior = 0
	VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY     VkPointClippingBehavior = 1
	VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES_KHR       VkPointClippingBehavior = VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES
	VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR VkPointClippingBehavior = VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY
	VK_POINT_CLIPPING_BEHAVIOR_MAX_ENUM                  VkPointClippingBehavior = 0x7FFFFFFF
)

type VkTessellationDomainOrigin = uint32

const (
	VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT     VkTessellationDomainOrigin = 0
	VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT     VkTessellationDomainOrigin = 1
	VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR VkTessellationDomainOrigin = VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT
	VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR VkTessellationDomainOrigin = VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT
	VK_TESSELLATION_DOMAIN_ORIGIN_MAX_ENUM       VkTessellationDomainOrigin = 0x7FFFFFFF
)

type VkSamplerYcbcrModelConversion = uint32

const (
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY       VkSamplerYcbcrModelConversion = 0
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY     VkSamplerYcbcrModelConversion = 1
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709          VkSamplerYcbcrModelConversion = 2
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601          VkSamplerYcbcrModelConversion = 3
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020         VkSamplerYcbcrModelConversion = 4
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR   VkSamplerYcbcrModelConversion = VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR VkSamplerYcbcrModelConversion = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR      VkSamplerYcbcrModelConversion = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR      VkSamplerYcbcrModelConversion = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR     VkSamplerYcbcrModelConversion = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	VK_SAMPLER_YCBCR_MODEL_CONVERSION_MAX_ENUM           VkSamplerYcbcrModelConversion = 0x7FFFFFFF
)

type VkSamplerYcbcrRange = uint32

const (
	VK_SAMPLER_YCBCR_RANGE_ITU_FULL       VkSamplerYcbcrRange = 0
	VK_SAMPLER_YCBCR_RANGE_ITU_NARROW     VkSamplerYcbcrRange = 1
	VK_SAMPLER_YCBCR_RANGE_ITU_FULL_KHR   VkSamplerYcbcrRange = VK_SAMPLER_YCBCR_RANGE_ITU_FULL
	VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR VkSamplerYcbcrRange = VK_SAMPLER_YCBCR_RANGE_ITU_NARROW
	VK_SAMPLER_YCBCR_RANGE_MAX_ENUM       VkSamplerYcbcrRange = 0x7FFFFFFF
)

type VkChromaLocation = uint32

const (
	VK_CHROMA_LOCATION_COSITED_EVEN     VkChromaLocation = 0
	VK_CHROMA_LOCATION_MIDPOINT         VkChromaLocation = 1
	VK_CHROMA_LOCATION_COSITED_EVEN_KHR VkChromaLocation = VK_CHROMA_LOCATION_COSITED_EVEN
	VK_CHROMA_LOCATION_MIDPOINT_KHR     VkChromaLocation = VK_CHROMA_LOCATION_MIDPOINT
	VK_CHROMA_LOCATION_MAX_ENUM         VkChromaLocation = 0x7FFFFFFF
)

type VkDescriptorUpdateTemplateType = uint32

const (
	VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET       VkDescriptorUpdateTemplateType = 0
	VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR VkDescriptorUpdateTemplateType = 1
	VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR   VkDescriptorUpdateTemplateType = VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_MAX_ENUM             VkDescriptorUpdateTemplateType = 0x7FFFFFFF
)

type VkSubgroupFeatureFlagBits = uint32

const (
	VK_SUBGROUP_FEATURE_BASIC_BIT            VkSubgroupFeatureFlagBits = 0x00000001
	VK_SUBGROUP_FEATURE_VOTE_BIT             VkSubgroupFeatureFlagBits = 0x00000002
	VK_SUBGROUP_FEATURE_ARITHMETIC_BIT       VkSubgroupFeatureFlagBits = 0x00000004
	VK_SUBGROUP_FEATURE_BALLOT_BIT           VkSubgroupFeatureFlagBits = 0x00000008
	VK_SUBGROUP_FEATURE_SHUFFLE_BIT          VkSubgroupFeatureFlagBits = 0x00000010
	VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT VkSubgroupFeatureFlagBits = 0x00000020
	VK_SUBGROUP_FEATURE_CLUSTERED_BIT        VkSubgroupFeatureFlagBits = 0x00000040
	VK_SUBGROUP_FEATURE_QUAD_BIT             VkSubgroupFeatureFlagBits = 0x00000080
	VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV   VkSubgroupFeatureFlagBits = 0x00000100
	VK_SUBGROUP_FEATURE_FLAG_BITS_MAX_ENUM   VkSubgroupFeatureFlagBits = 0x7FFFFFFF
)

type VkPeerMemoryFeatureFlagBits = uint32

const (
	VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT        VkPeerMemoryFeatureFlagBits = 0x00000001
	VK_PEER_MEMORY_FEATURE_COPY_DST_BIT        VkPeerMemoryFeatureFlagBits = 0x00000002
	VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT     VkPeerMemoryFeatureFlagBits = 0x00000004
	VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT     VkPeerMemoryFeatureFlagBits = 0x00000008
	VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR    VkPeerMemoryFeatureFlagBits = VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT
	VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR    VkPeerMemoryFeatureFlagBits = VK_PEER_MEMORY_FEATURE_COPY_DST_BIT
	VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR VkPeerMemoryFeatureFlagBits = VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT
	VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR VkPeerMemoryFeatureFlagBits = VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT
	VK_PEER_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM  VkPeerMemoryFeatureFlagBits = 0x7FFFFFFF
)

type VkMemoryAllocateFlagBits = uint32

const (
	VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT                       VkMemoryAllocateFlagBits = 0x00000001
	VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT                    VkMemoryAllocateFlagBits = 0x00000002
	VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT     VkMemoryAllocateFlagBits = 0x00000004
	VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR                   VkMemoryAllocateFlagBits = VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT
	VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR                VkMemoryAllocateFlagBits = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT
	VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR VkMemoryAllocateFlagBits = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT
	VK_MEMORY_ALLOCATE_FLAG_BITS_MAX_ENUM                    VkMemoryAllocateFlagBits = 0x7FFFFFFF
)

type VkExternalMemoryHandleTypeFlagBits = uint32

const (
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT                       VkExternalMemoryHandleTypeFlagBits = 0x00000001
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT                    VkExternalMemoryHandleTypeFlagBits = 0x00000002
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT                VkExternalMemoryHandleTypeFlagBits = 0x00000004
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT                   VkExternalMemoryHandleTypeFlagBits = 0x00000008
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT               VkExternalMemoryHandleTypeFlagBits = 0x00000010
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT                      VkExternalMemoryHandleTypeFlagBits = 0x00000020
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT                  VkExternalMemoryHandleTypeFlagBits = 0x00000040
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT                     VkExternalMemoryHandleTypeFlagBits = 0x00000200
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID VkExternalMemoryHandleTypeFlagBits = 0x00000400
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT             VkExternalMemoryHandleTypeFlagBits = 0x00000080
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT  VkExternalMemoryHandleTypeFlagBits = 0x00000100
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA              VkExternalMemoryHandleTypeFlagBits = 0x00000800
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV                 VkExternalMemoryHandleTypeFlagBits = 0x00001000
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR                   VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR                VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR            VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR               VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR           VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR                  VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR              VkExternalMemoryHandleTypeFlagBits = VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_FLAG_BITS_MAX_ENUM                  VkExternalMemoryHandleTypeFlagBits = 0x7FFFFFFF
)

type VkExternalMemoryFeatureFlagBits = uint32

const (
	VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT     VkExternalMemoryFeatureFlagBits = 0x00000001
	VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT         VkExternalMemoryFeatureFlagBits = 0x00000002
	VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT         VkExternalMemoryFeatureFlagBits = 0x00000004
	VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR VkExternalMemoryFeatureFlagBits = VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT
	VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR     VkExternalMemoryFeatureFlagBits = VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT
	VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR     VkExternalMemoryFeatureFlagBits = VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT
	VK_EXTERNAL_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM     VkExternalMemoryFeatureFlagBits = 0x7FFFFFFF
)

type VkExternalFenceHandleTypeFlagBits = uint32

const (
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT            VkExternalFenceHandleTypeFlagBits = 0x00000001
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT         VkExternalFenceHandleTypeFlagBits = 0x00000002
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     VkExternalFenceHandleTypeFlagBits = 0x00000004
	VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT              VkExternalFenceHandleTypeFlagBits = 0x00000008
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        VkExternalFenceHandleTypeFlagBits = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     VkExternalFenceHandleTypeFlagBits = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR VkExternalFenceHandleTypeFlagBits = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR          VkExternalFenceHandleTypeFlagBits = VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT
	VK_EXTERNAL_FENCE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       VkExternalFenceHandleTypeFlagBits = 0x7FFFFFFF
)

type VkExternalFenceFeatureFlagBits = uint32

const (
	VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT     VkExternalFenceFeatureFlagBits = 0x00000001
	VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT     VkExternalFenceFeatureFlagBits = 0x00000002
	VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR VkExternalFenceFeatureFlagBits = VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT_KHR VkExternalFenceFeatureFlagBits = VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT
	VK_EXTERNAL_FENCE_FEATURE_FLAG_BITS_MAX_ENUM VkExternalFenceFeatureFlagBits = 0x7FFFFFFF
)

type VkFenceImportFlagBits = uint32

const (
	VK_FENCE_IMPORT_TEMPORARY_BIT      VkFenceImportFlagBits = 0x00000001
	VK_FENCE_IMPORT_TEMPORARY_BIT_KHR  VkFenceImportFlagBits = VK_FENCE_IMPORT_TEMPORARY_BIT
	VK_FENCE_IMPORT_FLAG_BITS_MAX_ENUM VkFenceImportFlagBits = 0x7FFFFFFF
)

type VkSemaphoreImportFlagBits = uint32

const (
	VK_SEMAPHORE_IMPORT_TEMPORARY_BIT      VkSemaphoreImportFlagBits = 0x00000001
	VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR  VkSemaphoreImportFlagBits = VK_SEMAPHORE_IMPORT_TEMPORARY_BIT
	VK_SEMAPHORE_IMPORT_FLAG_BITS_MAX_ENUM VkSemaphoreImportFlagBits = 0x7FFFFFFF
)

type VkExternalSemaphoreHandleTypeFlagBits = uint32

const (
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT            VkExternalSemaphoreHandleTypeFlagBits = 0x00000001
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT         VkExternalSemaphoreHandleTypeFlagBits = 0x00000002
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     VkExternalSemaphoreHandleTypeFlagBits = 0x00000004
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT          VkExternalSemaphoreHandleTypeFlagBits = 0x00000008
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT              VkExternalSemaphoreHandleTypeFlagBits = 0x00000010
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA VkExternalSemaphoreHandleTypeFlagBits = 0x00000080
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D11_FENCE_BIT          VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR      VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR          VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT
	VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       VkExternalSemaphoreHandleTypeFlagBits = 0x7FFFFFFF
)

type VkExternalSemaphoreFeatureFlagBits = uint32

const (
	VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT     VkExternalSemaphoreFeatureFlagBits = 0x00000001
	VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT     VkExternalSemaphoreFeatureFlagBits = 0x00000002
	VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR VkExternalSemaphoreFeatureFlagBits = VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT
	VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR VkExternalSemaphoreFeatureFlagBits = VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT
	VK_EXTERNAL_SEMAPHORE_FEATURE_FLAG_BITS_MAX_ENUM VkExternalSemaphoreFeatureFlagBits = 0x7FFFFFFF
)

type VkDriverId = uint32

const (
	VK_DRIVER_ID_AMD_PROPRIETARY               VkDriverId = 1
	VK_DRIVER_ID_AMD_OPEN_SOURCE               VkDriverId = 2
	VK_DRIVER_ID_MESA_RADV                     VkDriverId = 3
	VK_DRIVER_ID_NVIDIA_PROPRIETARY            VkDriverId = 4
	VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS     VkDriverId = 5
	VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA        VkDriverId = 6
	VK_DRIVER_ID_IMAGINATION_PROPRIETARY       VkDriverId = 7
	VK_DRIVER_ID_QUALCOMM_PROPRIETARY          VkDriverId = 8
	VK_DRIVER_ID_ARM_PROPRIETARY               VkDriverId = 9
	VK_DRIVER_ID_GOOGLE_SWIFTSHADER            VkDriverId = 10
	VK_DRIVER_ID_GGP_PROPRIETARY               VkDriverId = 11
	VK_DRIVER_ID_BROADCOM_PROPRIETARY          VkDriverId = 12
	VK_DRIVER_ID_MESA_LLVMPIPE                 VkDriverId = 13
	VK_DRIVER_ID_MOLTENVK                      VkDriverId = 14
	VK_DRIVER_ID_COREAVI_PROPRIETARY           VkDriverId = 15
	VK_DRIVER_ID_JUICE_PROPRIETARY             VkDriverId = 16
	VK_DRIVER_ID_VERISILICON_PROPRIETARY       VkDriverId = 17
	VK_DRIVER_ID_AMD_PROPRIETARY_KHR           VkDriverId = VK_DRIVER_ID_AMD_PROPRIETARY
	VK_DRIVER_ID_AMD_OPEN_SOURCE_KHR           VkDriverId = VK_DRIVER_ID_AMD_OPEN_SOURCE
	VK_DRIVER_ID_MESA_RADV_KHR                 VkDriverId = VK_DRIVER_ID_MESA_RADV
	VK_DRIVER_ID_NVIDIA_PROPRIETARY_KHR        VkDriverId = VK_DRIVER_ID_NVIDIA_PROPRIETARY
	VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR VkDriverId = VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS
	VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR    VkDriverId = VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA
	VK_DRIVER_ID_IMAGINATION_PROPRIETARY_KHR   VkDriverId = VK_DRIVER_ID_IMAGINATION_PROPRIETARY
	VK_DRIVER_ID_QUALCOMM_PROPRIETARY_KHR      VkDriverId = VK_DRIVER_ID_QUALCOMM_PROPRIETARY
	VK_DRIVER_ID_ARM_PROPRIETARY_KHR           VkDriverId = VK_DRIVER_ID_ARM_PROPRIETARY
	VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR        VkDriverId = VK_DRIVER_ID_GOOGLE_SWIFTSHADER
	VK_DRIVER_ID_GGP_PROPRIETARY_KHR           VkDriverId = VK_DRIVER_ID_GGP_PROPRIETARY
	VK_DRIVER_ID_BROADCOM_PROPRIETARY_KHR      VkDriverId = VK_DRIVER_ID_BROADCOM_PROPRIETARY
	VK_DRIVER_ID_MAX_ENUM                      VkDriverId = 0x7FFFFFFF
)

type VkShaderFloatControlsIndependence = uint32

const (
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY     VkShaderFloatControlsIndependence = 0
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL             VkShaderFloatControlsIndependence = 1
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE            VkShaderFloatControlsIndependence = 2
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR VkShaderFloatControlsIndependence = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR         VkShaderFloatControlsIndependence = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR        VkShaderFloatControlsIndependence = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE
	VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_MAX_ENUM        VkShaderFloatControlsIndependence = 0x7FFFFFFF
)

type VkSamplerReductionMode = uint32

const (
	VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE     VkSamplerReductionMode = 0
	VK_SAMPLER_REDUCTION_MODE_MIN                  VkSamplerReductionMode = 1
	VK_SAMPLER_REDUCTION_MODE_MAX                  VkSamplerReductionMode = 2
	VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT VkSamplerReductionMode = VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE
	VK_SAMPLER_REDUCTION_MODE_MIN_EXT              VkSamplerReductionMode = VK_SAMPLER_REDUCTION_MODE_MIN
	VK_SAMPLER_REDUCTION_MODE_MAX_EXT              VkSamplerReductionMode = VK_SAMPLER_REDUCTION_MODE_MAX
	VK_SAMPLER_REDUCTION_MODE_MAX_ENUM             VkSamplerReductionMode = 0x7FFFFFFF
)

type VkSemaphoreType = uint32

const (
	VK_SEMAPHORE_TYPE_BINARY       VkSemaphoreType = 0
	VK_SEMAPHORE_TYPE_TIMELINE     VkSemaphoreType = 1
	VK_SEMAPHORE_TYPE_BINARY_KHR   VkSemaphoreType = VK_SEMAPHORE_TYPE_BINARY
	VK_SEMAPHORE_TYPE_TIMELINE_KHR VkSemaphoreType = VK_SEMAPHORE_TYPE_TIMELINE
	VK_SEMAPHORE_TYPE_MAX_ENUM     VkSemaphoreType = 0x7FFFFFFF
)

type VkResolveModeFlagBits = uint32

const (
	VK_RESOLVE_MODE_NONE                VkResolveModeFlagBits = 0
	VK_RESOLVE_MODE_SAMPLE_ZERO_BIT     VkResolveModeFlagBits = 0x00000001
	VK_RESOLVE_MODE_AVERAGE_BIT         VkResolveModeFlagBits = 0x00000002
	VK_RESOLVE_MODE_MIN_BIT             VkResolveModeFlagBits = 0x00000004
	VK_RESOLVE_MODE_MAX_BIT             VkResolveModeFlagBits = 0x00000008
	VK_RESOLVE_MODE_NONE_KHR            VkResolveModeFlagBits = VK_RESOLVE_MODE_NONE
	VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR VkResolveModeFlagBits = VK_RESOLVE_MODE_SAMPLE_ZERO_BIT
	VK_RESOLVE_MODE_AVERAGE_BIT_KHR     VkResolveModeFlagBits = VK_RESOLVE_MODE_AVERAGE_BIT
	VK_RESOLVE_MODE_MIN_BIT_KHR         VkResolveModeFlagBits = VK_RESOLVE_MODE_MIN_BIT
	VK_RESOLVE_MODE_MAX_BIT_KHR         VkResolveModeFlagBits = VK_RESOLVE_MODE_MAX_BIT
	VK_RESOLVE_MODE_FLAG_BITS_MAX_ENUM  VkResolveModeFlagBits = 0x7FFFFFFF
)

type VkDescriptorBindingFlagBits = uint32

const (
	VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT               VkDescriptorBindingFlagBits = 0x00000001
	VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT     VkDescriptorBindingFlagBits = 0x00000002
	VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT                 VkDescriptorBindingFlagBits = 0x00000004
	VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT       VkDescriptorBindingFlagBits = 0x00000008
	VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT           VkDescriptorBindingFlagBits = VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT
	VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT VkDescriptorBindingFlagBits = VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT
	VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT             VkDescriptorBindingFlagBits = VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT
	VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT   VkDescriptorBindingFlagBits = VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT
	VK_DESCRIPTOR_BINDING_FLAG_BITS_MAX_ENUM                  VkDescriptorBindingFlagBits = 0x7FFFFFFF
)

type VkSemaphoreWaitFlagBits = uint32

const (
	VK_SEMAPHORE_WAIT_ANY_BIT            VkSemaphoreWaitFlagBits = 0x00000001
	VK_SEMAPHORE_WAIT_ANY_BIT_KHR        VkSemaphoreWaitFlagBits = VK_SEMAPHORE_WAIT_ANY_BIT
	VK_SEMAPHORE_WAIT_FLAG_BITS_MAX_ENUM VkSemaphoreWaitFlagBits = 0x7FFFFFFF
)

type VkPresentModeKHR = uint32

const (
	VK_PRESENT_MODE_IMMEDIATE_KHR                 VkPresentModeKHR = 0
	VK_PRESENT_MODE_MAILBOX_KHR                   VkPresentModeKHR = 1
	VK_PRESENT_MODE_FIFO_KHR                      VkPresentModeKHR = 2
	VK_PRESENT_MODE_FIFO_RELAXED_KHR              VkPresentModeKHR = 3
	VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR     VkPresentModeKHR = 1000111000
	VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR VkPresentModeKHR = 1000111001
	VK_PRESENT_MODE_MAX_ENUM_KHR                  VkPresentModeKHR = 0x7FFFFFFF
)

type VkColorSpaceKHR = uint32

const (
	VK_COLOR_SPACE_SRGB_NONLINEAR_KHR          VkColorSpaceKHR = 0
	VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT    VkColorSpaceKHR = 1000104001
	VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT    VkColorSpaceKHR = 1000104002
	VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT       VkColorSpaceKHR = 1000104003
	VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT        VkColorSpaceKHR = 1000104004
	VK_COLOR_SPACE_BT709_LINEAR_EXT            VkColorSpaceKHR = 1000104005
	VK_COLOR_SPACE_BT709_NONLINEAR_EXT         VkColorSpaceKHR = 1000104006
	VK_COLOR_SPACE_BT2020_LINEAR_EXT           VkColorSpaceKHR = 1000104007
	VK_COLOR_SPACE_HDR10_ST2084_EXT            VkColorSpaceKHR = 1000104008
	VK_COLOR_SPACE_DOLBYVISION_EXT             VkColorSpaceKHR = 1000104009
	VK_COLOR_SPACE_HDR10_HLG_EXT               VkColorSpaceKHR = 1000104010
	VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT         VkColorSpaceKHR = 1000104011
	VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT      VkColorSpaceKHR = 1000104012
	VK_COLOR_SPACE_PASS_THROUGH_EXT            VkColorSpaceKHR = 1000104013
	VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT VkColorSpaceKHR = 1000104014
	VK_COLOR_SPACE_DISPLAY_NATIVE_AMD          VkColorSpaceKHR = 1000213000
	VK_COLORSPACE_SRGB_NONLINEAR_KHR           VkColorSpaceKHR = VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
	VK_COLOR_SPACE_DCI_P3_LINEAR_EXT           VkColorSpaceKHR = VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT
	VK_COLOR_SPACE_MAX_ENUM_KHR                VkColorSpaceKHR = 0x7FFFFFFF
)

type VkSurfaceTransformFlagBitsKHR = uint32

const (
	VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR                     VkSurfaceTransformFlagBitsKHR = 0x00000001
	VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR                    VkSurfaceTransformFlagBitsKHR = 0x00000002
	VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR                   VkSurfaceTransformFlagBitsKHR = 0x00000004
	VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR                   VkSurfaceTransformFlagBitsKHR = 0x00000008
	VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR            VkSurfaceTransformFlagBitsKHR = 0x00000010
	VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR  VkSurfaceTransformFlagBitsKHR = 0x00000020
	VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x00000040
	VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x00000080
	VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR                      VkSurfaceTransformFlagBitsKHR = 0x00000100
	VK_SURFACE_TRANSFORM_FLAG_BITS_MAX_ENUM_KHR               VkSurfaceTransformFlagBitsKHR = 0x7FFFFFFF
)

type VkCompositeAlphaFlagBitsKHR = uint32

const (
	VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR          VkCompositeAlphaFlagBitsKHR = 0x00000001
	VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR  VkCompositeAlphaFlagBitsKHR = 0x00000002
	VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR VkCompositeAlphaFlagBitsKHR = 0x00000004
	VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR         VkCompositeAlphaFlagBitsKHR = 0x00000008
	VK_COMPOSITE_ALPHA_FLAG_BITS_MAX_ENUM_KHR  VkCompositeAlphaFlagBitsKHR = 0x7FFFFFFF
)

type VkSwapchainCreateFlagBitsKHR = uint32

const (
	VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR VkSwapchainCreateFlagBitsKHR = 0x00000001
	VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR                   VkSwapchainCreateFlagBitsKHR = 0x00000002
	VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR              VkSwapchainCreateFlagBitsKHR = 0x00000004
	VK_SWAPCHAIN_CREATE_FLAG_BITS_MAX_ENUM_KHR              VkSwapchainCreateFlagBitsKHR = 0x7FFFFFFF
)

type VkDeviceGroupPresentModeFlagBitsKHR = uint32

const (
	VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR              VkDeviceGroupPresentModeFlagBitsKHR = 0x00000001
	VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR             VkDeviceGroupPresentModeFlagBitsKHR = 0x00000002
	VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR                VkDeviceGroupPresentModeFlagBitsKHR = 0x00000004
	VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR VkDeviceGroupPresentModeFlagBitsKHR = 0x00000008
	VK_DEVICE_GROUP_PRESENT_MODE_FLAG_BITS_MAX_ENUM_KHR     VkDeviceGroupPresentModeFlagBitsKHR = 0x7FFFFFFF
)

type VkDisplayPlaneAlphaFlagBitsKHR = uint32

const (
	VK_DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR                  VkDisplayPlaneAlphaFlagBitsKHR = 0x00000001
	VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR                  VkDisplayPlaneAlphaFlagBitsKHR = 0x00000002
	VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR               VkDisplayPlaneAlphaFlagBitsKHR = 0x00000004
	VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR VkDisplayPlaneAlphaFlagBitsKHR = 0x00000008
	VK_DISPLAY_PLANE_ALPHA_FLAG_BITS_MAX_ENUM_KHR          VkDisplayPlaneAlphaFlagBitsKHR = 0x7FFFFFFF
)

type VkPerformanceCounterUnitKHR = uint32

const (
	VK_PERFORMANCE_COUNTER_UNIT_GENERIC_KHR          VkPerformanceCounterUnitKHR = 0
	VK_PERFORMANCE_COUNTER_UNIT_PERCENTAGE_KHR       VkPerformanceCounterUnitKHR = 1
	VK_PERFORMANCE_COUNTER_UNIT_NANOSECONDS_KHR      VkPerformanceCounterUnitKHR = 2
	VK_PERFORMANCE_COUNTER_UNIT_BYTES_KHR            VkPerformanceCounterUnitKHR = 3
	VK_PERFORMANCE_COUNTER_UNIT_BYTES_PER_SECOND_KHR VkPerformanceCounterUnitKHR = 4
	VK_PERFORMANCE_COUNTER_UNIT_KELVIN_KHR           VkPerformanceCounterUnitKHR = 5
	VK_PERFORMANCE_COUNTER_UNIT_WATTS_KHR            VkPerformanceCounterUnitKHR = 6
	VK_PERFORMANCE_COUNTER_UNIT_VOLTS_KHR            VkPerformanceCounterUnitKHR = 7
	VK_PERFORMANCE_COUNTER_UNIT_AMPS_KHR             VkPerformanceCounterUnitKHR = 8
	VK_PERFORMANCE_COUNTER_UNIT_HERTZ_KHR            VkPerformanceCounterUnitKHR = 9
	VK_PERFORMANCE_COUNTER_UNIT_CYCLES_KHR           VkPerformanceCounterUnitKHR = 10
	VK_PERFORMANCE_COUNTER_UNIT_MAX_ENUM_KHR         VkPerformanceCounterUnitKHR = 0x7FFFFFFF
)

type VkPerformanceCounterScopeKHR = uint32

const (
	VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR VkPerformanceCounterScopeKHR = 0
	VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR    VkPerformanceCounterScopeKHR = 1
	VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR        VkPerformanceCounterScopeKHR = 2
	VK_QUERY_SCOPE_COMMAND_BUFFER_KHR               VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR
	VK_QUERY_SCOPE_RENDER_PASS_KHR                  VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR
	VK_QUERY_SCOPE_COMMAND_KHR                      VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR
	VK_PERFORMANCE_COUNTER_SCOPE_MAX_ENUM_KHR       VkPerformanceCounterScopeKHR = 0x7FFFFFFF
)

type VkPerformanceCounterStorageKHR = uint32

const (
	VK_PERFORMANCE_COUNTER_STORAGE_INT32_KHR    VkPerformanceCounterStorageKHR = 0
	VK_PERFORMANCE_COUNTER_STORAGE_INT64_KHR    VkPerformanceCounterStorageKHR = 1
	VK_PERFORMANCE_COUNTER_STORAGE_UINT32_KHR   VkPerformanceCounterStorageKHR = 2
	VK_PERFORMANCE_COUNTER_STORAGE_UINT64_KHR   VkPerformanceCounterStorageKHR = 3
	VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32_KHR  VkPerformanceCounterStorageKHR = 4
	VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64_KHR  VkPerformanceCounterStorageKHR = 5
	VK_PERFORMANCE_COUNTER_STORAGE_MAX_ENUM_KHR VkPerformanceCounterStorageKHR = 0x7FFFFFFF
)

type VkPerformanceCounterDescriptionFlagBitsKHR = uint32

const (
	VK_PERFORMANCE_COUNTER_DESCRIPTION_PERFORMANCE_IMPACTING_BIT_KHR VkPerformanceCounterDescriptionFlagBitsKHR = 0x00000001
	VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_BIT_KHR VkPerformanceCounterDescriptionFlagBitsKHR = 0x00000002
	VK_PERFORMANCE_COUNTER_DESCRIPTION_PERFORMANCE_IMPACTING_KHR     VkPerformanceCounterDescriptionFlagBitsKHR = VK_PERFORMANCE_COUNTER_DESCRIPTION_PERFORMANCE_IMPACTING_BIT_KHR
	VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_KHR     VkPerformanceCounterDescriptionFlagBitsKHR = VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_BIT_KHR
	VK_PERFORMANCE_COUNTER_DESCRIPTION_FLAG_BITS_MAX_ENUM_KHR        VkPerformanceCounterDescriptionFlagBitsKHR = 0x7FFFFFFF
)

type VkAcquireProfilingLockFlagBitsKHR = uint32

const (
	VK_ACQUIRE_PROFILING_LOCK_FLAG_BITS_MAX_ENUM_KHR VkAcquireProfilingLockFlagBitsKHR = 0x7FFFFFFF
)

type VkFragmentShadingRateCombinerOpKHR = uint32

const (
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR     VkFragmentShadingRateCombinerOpKHR = 0
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR  VkFragmentShadingRateCombinerOpKHR = 1
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR      VkFragmentShadingRateCombinerOpKHR = 2
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR      VkFragmentShadingRateCombinerOpKHR = 3
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR      VkFragmentShadingRateCombinerOpKHR = 4
	VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_ENUM_KHR VkFragmentShadingRateCombinerOpKHR = 0x7FFFFFFF
)

type VkPipelineExecutableStatisticFormatKHR = uint32

const (
	VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR   VkPipelineExecutableStatisticFormatKHR = 0
	VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR    VkPipelineExecutableStatisticFormatKHR = 1
	VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR   VkPipelineExecutableStatisticFormatKHR = 2
	VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR  VkPipelineExecutableStatisticFormatKHR = 3
	VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_MAX_ENUM_KHR VkPipelineExecutableStatisticFormatKHR = 0x7FFFFFFF
)

type VkSubmitFlagBitsKHR = uint32

const (
	VK_SUBMIT_PROTECTED_BIT_KHR      VkSubmitFlagBitsKHR = 0x00000001
	VK_SUBMIT_FLAG_BITS_MAX_ENUM_KHR VkSubmitFlagBitsKHR = 0x7FFFFFFF
)

type VkDebugReportObjectTypeEXT = uint32

const (
	VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT                        VkDebugReportObjectTypeEXT = 0
	VK_DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT                       VkDebugReportObjectTypeEXT = 1
	VK_DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT                VkDebugReportObjectTypeEXT = 2
	VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT                         VkDebugReportObjectTypeEXT = 3
	VK_DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT                          VkDebugReportObjectTypeEXT = 4
	VK_DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT                      VkDebugReportObjectTypeEXT = 5
	VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT                 VkDebugReportObjectTypeEXT = 6
	VK_DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT                          VkDebugReportObjectTypeEXT = 7
	VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT                  VkDebugReportObjectTypeEXT = 8
	VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT                         VkDebugReportObjectTypeEXT = 9
	VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT                          VkDebugReportObjectTypeEXT = 10
	VK_DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT                          VkDebugReportObjectTypeEXT = 11
	VK_DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT                     VkDebugReportObjectTypeEXT = 12
	VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT                    VkDebugReportObjectTypeEXT = 13
	VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT                     VkDebugReportObjectTypeEXT = 14
	VK_DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT                  VkDebugReportObjectTypeEXT = 15
	VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT                 VkDebugReportObjectTypeEXT = 16
	VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT                VkDebugReportObjectTypeEXT = 17
	VK_DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT                    VkDebugReportObjectTypeEXT = 18
	VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT                       VkDebugReportObjectTypeEXT = 19
	VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT          VkDebugReportObjectTypeEXT = 20
	VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT                        VkDebugReportObjectTypeEXT = 21
	VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT                VkDebugReportObjectTypeEXT = 22
	VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT                 VkDebugReportObjectTypeEXT = 23
	VK_DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT                    VkDebugReportObjectTypeEXT = 24
	VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT                   VkDebugReportObjectTypeEXT = 25
	VK_DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT                    VkDebugReportObjectTypeEXT = 26
	VK_DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT                  VkDebugReportObjectTypeEXT = 27
	VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT      VkDebugReportObjectTypeEXT = 28
	VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT                    VkDebugReportObjectTypeEXT = 29
	VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT               VkDebugReportObjectTypeEXT = 30
	VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT           VkDebugReportObjectTypeEXT = 33
	VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT       VkDebugReportObjectTypeEXT = 1000156000
	VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT     VkDebugReportObjectTypeEXT = 1000085000
	VK_DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXT                  VkDebugReportObjectTypeEXT = 1000029000
	VK_DEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT                VkDebugReportObjectTypeEXT = 1000029001
	VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR_EXT     VkDebugReportObjectTypeEXT = 1000150000
	VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT      VkDebugReportObjectTypeEXT = 1000165000
	VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT                   VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT
	VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT               VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT
	VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT
	VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT   VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT
	VK_DEBUG_REPORT_OBJECT_TYPE_MAX_ENUM_EXT                       VkDebugReportObjectTypeEXT = 0x7FFFFFFF
)

type VkDebugReportFlagBitsEXT = uint32

const (
	VK_DEBUG_REPORT_INFORMATION_BIT_EXT         VkDebugReportFlagBitsEXT = 0x00000001
	VK_DEBUG_REPORT_WARNING_BIT_EXT             VkDebugReportFlagBitsEXT = 0x00000002
	VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT VkDebugReportFlagBitsEXT = 0x00000004
	VK_DEBUG_REPORT_ERROR_BIT_EXT               VkDebugReportFlagBitsEXT = 0x00000008
	VK_DEBUG_REPORT_DEBUG_BIT_EXT               VkDebugReportFlagBitsEXT = 0x00000010
	VK_DEBUG_REPORT_FLAG_BITS_MAX_ENUM_EXT      VkDebugReportFlagBitsEXT = 0x7FFFFFFF
)

type VkRasterizationOrderAMD = uint32

const (
	VK_RASTERIZATION_ORDER_STRICT_AMD   VkRasterizationOrderAMD = 0
	VK_RASTERIZATION_ORDER_RELAXED_AMD  VkRasterizationOrderAMD = 1
	VK_RASTERIZATION_ORDER_MAX_ENUM_AMD VkRasterizationOrderAMD = 0x7FFFFFFF
)

type VkShaderInfoTypeAMD = uint32

const (
	VK_SHADER_INFO_TYPE_STATISTICS_AMD  VkShaderInfoTypeAMD = 0
	VK_SHADER_INFO_TYPE_BINARY_AMD      VkShaderInfoTypeAMD = 1
	VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD VkShaderInfoTypeAMD = 2
	VK_SHADER_INFO_TYPE_MAX_ENUM_AMD    VkShaderInfoTypeAMD = 0x7FFFFFFF
)

type VkExternalMemoryHandleTypeFlagBitsNV = uint32

const (
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV     VkExternalMemoryHandleTypeFlagBitsNV = 0x00000001
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV VkExternalMemoryHandleTypeFlagBitsNV = 0x00000002
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV      VkExternalMemoryHandleTypeFlagBitsNV = 0x00000004
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_KMT_BIT_NV  VkExternalMemoryHandleTypeFlagBitsNV = 0x00000008
	VK_EXTERNAL_MEMORY_HANDLE_TYPE_FLAG_BITS_MAX_ENUM_NV   VkExternalMemoryHandleTypeFlagBitsNV = 0x7FFFFFFF
)

type VkExternalMemoryFeatureFlagBitsNV = uint32

const (
	VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV VkExternalMemoryFeatureFlagBitsNV = 0x00000001
	VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV     VkExternalMemoryFeatureFlagBitsNV = 0x00000002
	VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV     VkExternalMemoryFeatureFlagBitsNV = 0x00000004
	VK_EXTERNAL_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM_NV VkExternalMemoryFeatureFlagBitsNV = 0x7FFFFFFF
)

type VkValidationCheckEXT = uint32

const (
	VK_VALIDATION_CHECK_ALL_EXT      VkValidationCheckEXT = 0
	VK_VALIDATION_CHECK_SHADERS_EXT  VkValidationCheckEXT = 1
	VK_VALIDATION_CHECK_MAX_ENUM_EXT VkValidationCheckEXT = 0x7FFFFFFF
)

type VkConditionalRenderingFlagBitsEXT = uint32

const (
	VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT       VkConditionalRenderingFlagBitsEXT = 0x00000001
	VK_CONDITIONAL_RENDERING_FLAG_BITS_MAX_ENUM_EXT VkConditionalRenderingFlagBitsEXT = 0x7FFFFFFF
)

type VkSurfaceCounterFlagBitsEXT = uint32

const (
	VK_SURFACE_COUNTER_VBLANK_BIT_EXT         VkSurfaceCounterFlagBitsEXT = 0x00000001
	VK_SURFACE_COUNTER_VBLANK_EXT             VkSurfaceCounterFlagBitsEXT = VK_SURFACE_COUNTER_VBLANK_BIT_EXT
	VK_SURFACE_COUNTER_FLAG_BITS_MAX_ENUM_EXT VkSurfaceCounterFlagBitsEXT = 0x7FFFFFFF
)

type VkDisplayPowerStateEXT = uint32

const (
	VK_DISPLAY_POWER_STATE_OFF_EXT      VkDisplayPowerStateEXT = 0
	VK_DISPLAY_POWER_STATE_SUSPEND_EXT  VkDisplayPowerStateEXT = 1
	VK_DISPLAY_POWER_STATE_ON_EXT       VkDisplayPowerStateEXT = 2
	VK_DISPLAY_POWER_STATE_MAX_ENUM_EXT VkDisplayPowerStateEXT = 0x7FFFFFFF
)

type VkDeviceEventTypeEXT = uint32

const (
	VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT VkDeviceEventTypeEXT = 0
	VK_DEVICE_EVENT_TYPE_MAX_ENUM_EXT        VkDeviceEventTypeEXT = 0x7FFFFFFF
)

type VkDisplayEventTypeEXT = uint32

const (
	VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT VkDisplayEventTypeEXT = 0
	VK_DISPLAY_EVENT_TYPE_MAX_ENUM_EXT        VkDisplayEventTypeEXT = 0x7FFFFFFF
)

type VkViewportCoordinateSwizzleNV = uint32

const (
	VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV VkViewportCoordinateSwizzleNV = 0
	VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV VkViewportCoordinateSwizzleNV = 1
	VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV VkViewportCoordinateSwizzleNV = 2
	VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV VkViewportCoordinateSwizzleNV = 3
	VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV VkViewportCoordinateSwizzleNV = 4
	VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV VkViewportCoordinateSwizzleNV = 5
	VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV VkViewportCoordinateSwizzleNV = 6
	VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV VkViewportCoordinateSwizzleNV = 7
	VK_VIEWPORT_COORDINATE_SWIZZLE_MAX_ENUM_NV   VkViewportCoordinateSwizzleNV = 0x7FFFFFFF
)

type VkDiscardRectangleModeEXT = uint32

const (
	VK_DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT VkDiscardRectangleModeEXT = 0
	VK_DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT VkDiscardRectangleModeEXT = 1
	VK_DISCARD_RECTANGLE_MODE_MAX_ENUM_EXT  VkDiscardRectangleModeEXT = 0x7FFFFFFF
)

type VkConservativeRasterizationModeEXT = uint32

const (
	VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT      VkConservativeRasterizationModeEXT = 0
	VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT  VkConservativeRasterizationModeEXT = 1
	VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT VkConservativeRasterizationModeEXT = 2
	VK_CONSERVATIVE_RASTERIZATION_MODE_MAX_ENUM_EXT      VkConservativeRasterizationModeEXT = 0x7FFFFFFF
)

type VkDebugUtilsMessageSeverityFlagBitsEXT = uint32

const (
	VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT        VkDebugUtilsMessageSeverityFlagBitsEXT = 0x00000001
	VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT           VkDebugUtilsMessageSeverityFlagBitsEXT = 0x00000010
	VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT        VkDebugUtilsMessageSeverityFlagBitsEXT = 0x00000100
	VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT          VkDebugUtilsMessageSeverityFlagBitsEXT = 0x00001000
	VK_DEBUG_UTILS_MESSAGE_SEVERITY_FLAG_BITS_MAX_ENUM_EXT VkDebugUtilsMessageSeverityFlagBitsEXT = 0x7FFFFFFF
)

type VkDebugUtilsMessageTypeFlagBitsEXT = uint32

const (
	VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT        VkDebugUtilsMessageTypeFlagBitsEXT = 0x00000001
	VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT     VkDebugUtilsMessageTypeFlagBitsEXT = 0x00000002
	VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT    VkDebugUtilsMessageTypeFlagBitsEXT = 0x00000004
	VK_DEBUG_UTILS_MESSAGE_TYPE_FLAG_BITS_MAX_ENUM_EXT VkDebugUtilsMessageTypeFlagBitsEXT = 0x7FFFFFFF
)

type VkBlendOverlapEXT = uint32

const (
	VK_BLEND_OVERLAP_UNCORRELATED_EXT VkBlendOverlapEXT = 0
	VK_BLEND_OVERLAP_DISJOINT_EXT     VkBlendOverlapEXT = 1
	VK_BLEND_OVERLAP_CONJOINT_EXT     VkBlendOverlapEXT = 2
	VK_BLEND_OVERLAP_MAX_ENUM_EXT     VkBlendOverlapEXT = 0x7FFFFFFF
)

type VkCoverageModulationModeNV = uint32

const (
	VK_COVERAGE_MODULATION_MODE_NONE_NV     VkCoverageModulationModeNV = 0
	VK_COVERAGE_MODULATION_MODE_RGB_NV      VkCoverageModulationModeNV = 1
	VK_COVERAGE_MODULATION_MODE_ALPHA_NV    VkCoverageModulationModeNV = 2
	VK_COVERAGE_MODULATION_MODE_RGBA_NV     VkCoverageModulationModeNV = 3
	VK_COVERAGE_MODULATION_MODE_MAX_ENUM_NV VkCoverageModulationModeNV = 0x7FFFFFFF
)

type VkValidationCacheHeaderVersionEXT = uint32

const (
	VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT      VkValidationCacheHeaderVersionEXT = 1
	VK_VALIDATION_CACHE_HEADER_VERSION_MAX_ENUM_EXT VkValidationCacheHeaderVersionEXT = 0x7FFFFFFF
)

type VkShadingRatePaletteEntryNV = uint32

const (
	VK_SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV              VkShadingRatePaletteEntryNV = 0
	VK_SHADING_RATE_PALETTE_ENTRY_16_INVOCATIONS_PER_PIXEL_NV    VkShadingRatePaletteEntryNV = 1
	VK_SHADING_RATE_PALETTE_ENTRY_8_INVOCATIONS_PER_PIXEL_NV     VkShadingRatePaletteEntryNV = 2
	VK_SHADING_RATE_PALETTE_ENTRY_4_INVOCATIONS_PER_PIXEL_NV     VkShadingRatePaletteEntryNV = 3
	VK_SHADING_RATE_PALETTE_ENTRY_2_INVOCATIONS_PER_PIXEL_NV     VkShadingRatePaletteEntryNV = 4
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_PIXEL_NV      VkShadingRatePaletteEntryNV = 5
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X1_PIXELS_NV VkShadingRatePaletteEntryNV = 6
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_1X2_PIXELS_NV VkShadingRatePaletteEntryNV = 7
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X2_PIXELS_NV VkShadingRatePaletteEntryNV = 8
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X2_PIXELS_NV VkShadingRatePaletteEntryNV = 9
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X4_PIXELS_NV VkShadingRatePaletteEntryNV = 10
	VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV VkShadingRatePaletteEntryNV = 11
	VK_SHADING_RATE_PALETTE_ENTRY_MAX_ENUM_NV                    VkShadingRatePaletteEntryNV = 0x7FFFFFFF
)

type VkCoarseSampleOrderTypeNV = uint32

const (
	VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV      VkCoarseSampleOrderTypeNV = 0
	VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV       VkCoarseSampleOrderTypeNV = 1
	VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV  VkCoarseSampleOrderTypeNV = 2
	VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV VkCoarseSampleOrderTypeNV = 3
	VK_COARSE_SAMPLE_ORDER_TYPE_MAX_ENUM_NV     VkCoarseSampleOrderTypeNV = 0x7FFFFFFF
)

type VkRayTracingShaderGroupTypeKHR = uint32

const (
	VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR              VkRayTracingShaderGroupTypeKHR = 0
	VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR  VkRayTracingShaderGroupTypeKHR = 1
	VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR VkRayTracingShaderGroupTypeKHR = 2
	VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV               VkRayTracingShaderGroupTypeKHR = VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR
	VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV   VkRayTracingShaderGroupTypeKHR = VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR
	VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV  VkRayTracingShaderGroupTypeKHR = VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR
	VK_RAY_TRACING_SHADER_GROUP_TYPE_MAX_ENUM_KHR             VkRayTracingShaderGroupTypeKHR = 0x7FFFFFFF
)

type VkGeometryTypeKHR = uint32

const (
	VK_GEOMETRY_TYPE_TRIANGLES_KHR VkGeometryTypeKHR = 0
	VK_GEOMETRY_TYPE_AABBS_KHR     VkGeometryTypeKHR = 1
	VK_GEOMETRY_TYPE_INSTANCES_KHR VkGeometryTypeKHR = 2
	VK_GEOMETRY_TYPE_TRIANGLES_NV  VkGeometryTypeKHR = VK_GEOMETRY_TYPE_TRIANGLES_KHR
	VK_GEOMETRY_TYPE_AABBS_NV      VkGeometryTypeKHR = VK_GEOMETRY_TYPE_AABBS_KHR
	VK_GEOMETRY_TYPE_MAX_ENUM_KHR  VkGeometryTypeKHR = 0x7FFFFFFF
)

type VkAccelerationStructureTypeKHR = uint32

const (
	VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR    VkAccelerationStructureTypeKHR = 0
	VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR VkAccelerationStructureTypeKHR = 1
	VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR      VkAccelerationStructureTypeKHR = 2
	VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV     VkAccelerationStructureTypeKHR = VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR
	VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV  VkAccelerationStructureTypeKHR = VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR
	VK_ACCELERATION_STRUCTURE_TYPE_MAX_ENUM_KHR     VkAccelerationStructureTypeKHR = 0x7FFFFFFF
)

type VkCopyAccelerationStructureModeKHR = uint32

const (
	VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR       VkCopyAccelerationStructureModeKHR = 0
	VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR     VkCopyAccelerationStructureModeKHR = 1
	VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR   VkCopyAccelerationStructureModeKHR = 2
	VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR VkCopyAccelerationStructureModeKHR = 3
	VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV        VkCopyAccelerationStructureModeKHR = VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR
	VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV      VkCopyAccelerationStructureModeKHR = VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR
	VK_COPY_ACCELERATION_STRUCTURE_MODE_MAX_ENUM_KHR    VkCopyAccelerationStructureModeKHR = 0x7FFFFFFF
)

type VkAccelerationStructureMemoryRequirementsTypeNV = uint32

const (
	VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV         VkAccelerationStructureMemoryRequirementsTypeNV = 0
	VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV  VkAccelerationStructureMemoryRequirementsTypeNV = 1
	VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV VkAccelerationStructureMemoryRequirementsTypeNV = 2
	VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_MAX_ENUM_NV       VkAccelerationStructureMemoryRequirementsTypeNV = 0x7FFFFFFF
)

type VkGeometryFlagBitsKHR = uint32

const (
	VK_GEOMETRY_OPAQUE_BIT_KHR                          VkGeometryFlagBitsKHR = 0x00000001
	VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR VkGeometryFlagBitsKHR = 0x00000002
	VK_GEOMETRY_OPAQUE_BIT_NV                           VkGeometryFlagBitsKHR = VK_GEOMETRY_OPAQUE_BIT_KHR
	VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV  VkGeometryFlagBitsKHR = VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR
	VK_GEOMETRY_FLAG_BITS_MAX_ENUM_KHR                  VkGeometryFlagBitsKHR = 0x7FFFFFFF
)

type VkGeometryInstanceFlagBitsKHR = uint32

const (
	VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR    VkGeometryInstanceFlagBitsKHR = 0x00000001
	VK_GEOMETRY_INSTANCE_TRIANGLE_FLIP_FACING_BIT_KHR            VkGeometryInstanceFlagBitsKHR = 0x00000002
	VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR                    VkGeometryInstanceFlagBitsKHR = 0x00000004
	VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR                 VkGeometryInstanceFlagBitsKHR = 0x00000008
	VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_KHR VkGeometryInstanceFlagBitsKHR = VK_GEOMETRY_INSTANCE_TRIANGLE_FLIP_FACING_BIT_KHR
	VK_GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV            VkGeometryInstanceFlagBitsKHR = VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR
	VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV  VkGeometryInstanceFlagBitsKHR = VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_KHR
	VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV                     VkGeometryInstanceFlagBitsKHR = VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR
	VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV                  VkGeometryInstanceFlagBitsKHR = VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR
	VK_GEOMETRY_INSTANCE_FLAG_BITS_MAX_ENUM_KHR                  VkGeometryInstanceFlagBitsKHR = 0x7FFFFFFF
)

type VkBuildAccelerationStructureFlagBitsKHR = uint32

const (
	VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR      VkBuildAccelerationStructureFlagBitsKHR = 0x00000001
	VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR  VkBuildAccelerationStructureFlagBitsKHR = 0x00000002
	VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x00000004
	VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x00000008
	VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR        VkBuildAccelerationStructureFlagBitsKHR = 0x00000010
	VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV             VkBuildAccelerationStructureFlagBitsKHR = 0x00000020
	VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV       VkBuildAccelerationStructureFlagBitsKHR = VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR
	VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV   VkBuildAccelerationStructureFlagBitsKHR = VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR
	VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV  VkBuildAccelerationStructureFlagBitsKHR = VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR
	VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV  VkBuildAccelerationStructureFlagBitsKHR = VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR
	VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV         VkBuildAccelerationStructureFlagBitsKHR = VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR
	VK_BUILD_ACCELERATION_STRUCTURE_FLAG_BITS_MAX_ENUM_KHR    VkBuildAccelerationStructureFlagBitsKHR = 0x7FFFFFFF
)

type VkQueueGlobalPriorityEXT = uint32

const (
	VK_QUEUE_GLOBAL_PRIORITY_LOW_EXT      VkQueueGlobalPriorityEXT = 128
	VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT   VkQueueGlobalPriorityEXT = 256
	VK_QUEUE_GLOBAL_PRIORITY_HIGH_EXT     VkQueueGlobalPriorityEXT = 512
	VK_QUEUE_GLOBAL_PRIORITY_REALTIME_EXT VkQueueGlobalPriorityEXT = 1024
	VK_QUEUE_GLOBAL_PRIORITY_MAX_ENUM_EXT VkQueueGlobalPriorityEXT = 0x7FFFFFFF
)

type VkPipelineCompilerControlFlagBitsAMD = uint32

const (
	VK_PIPELINE_COMPILER_CONTROL_FLAG_BITS_MAX_ENUM_AMD VkPipelineCompilerControlFlagBitsAMD = 0x7FFFFFFF
)

type VkTimeDomainEXT = uint32

const (
	VK_TIME_DOMAIN_DEVICE_EXT                    VkTimeDomainEXT = 0
	VK_TIME_DOMAIN_CLOCK_MONOTONIC_EXT           VkTimeDomainEXT = 1
	VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT       VkTimeDomainEXT = 2
	VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT VkTimeDomainEXT = 3
	VK_TIME_DOMAIN_MAX_ENUM_EXT                  VkTimeDomainEXT = 0x7FFFFFFF
)

type VkMemoryOverallocationBehaviorAMD = uint32

const (
	VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD    VkMemoryOverallocationBehaviorAMD = 0
	VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD    VkMemoryOverallocationBehaviorAMD = 1
	VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD VkMemoryOverallocationBehaviorAMD = 2
	VK_MEMORY_OVERALLOCATION_BEHAVIOR_MAX_ENUM_AMD   VkMemoryOverallocationBehaviorAMD = 0x7FFFFFFF
)

type VkPipelineCreationFeedbackFlagBitsEXT = uint32

const (
	VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT                          VkPipelineCreationFeedbackFlagBitsEXT = 0x00000001
	VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT VkPipelineCreationFeedbackFlagBitsEXT = 0x00000002
	VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT     VkPipelineCreationFeedbackFlagBitsEXT = 0x00000004
	VK_PIPELINE_CREATION_FEEDBACK_FLAG_BITS_MAX_ENUM_EXT                 VkPipelineCreationFeedbackFlagBitsEXT = 0x7FFFFFFF
)

type VkPerformanceConfigurationTypeINTEL = uint32

const (
	VK_PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL VkPerformanceConfigurationTypeINTEL = 0
	VK_PERFORMANCE_CONFIGURATION_TYPE_MAX_ENUM_INTEL                                  VkPerformanceConfigurationTypeINTEL = 0x7FFFFFFF
)

type VkQueryPoolSamplingModeINTEL = uint32

const (
	VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL   VkQueryPoolSamplingModeINTEL = 0
	VK_QUERY_POOL_SAMPLING_MODE_MAX_ENUM_INTEL VkQueryPoolSamplingModeINTEL = 0x7FFFFFFF
)

type VkPerformanceOverrideTypeINTEL = uint32

const (
	VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL    VkPerformanceOverrideTypeINTEL = 0
	VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL VkPerformanceOverrideTypeINTEL = 1
	VK_PERFORMANCE_OVERRIDE_TYPE_MAX_ENUM_INTEL         VkPerformanceOverrideTypeINTEL = 0x7FFFFFFF
)

type VkPerformanceParameterTypeINTEL = uint32

const (
	VK_PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL    VkPerformanceParameterTypeINTEL = 0
	VK_PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL VkPerformanceParameterTypeINTEL = 1
	VK_PERFORMANCE_PARAMETER_TYPE_MAX_ENUM_INTEL                 VkPerformanceParameterTypeINTEL = 0x7FFFFFFF
)

type VkPerformanceValueTypeINTEL = uint32

const (
	VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL   VkPerformanceValueTypeINTEL = 0
	VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL   VkPerformanceValueTypeINTEL = 1
	VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL    VkPerformanceValueTypeINTEL = 2
	VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL     VkPerformanceValueTypeINTEL = 3
	VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL   VkPerformanceValueTypeINTEL = 4
	VK_PERFORMANCE_VALUE_TYPE_MAX_ENUM_INTEL VkPerformanceValueTypeINTEL = 0x7FFFFFFF
)

type VkShaderCorePropertiesFlagBitsAMD = uint32

const (
	VK_SHADER_CORE_PROPERTIES_FLAG_BITS_MAX_ENUM_AMD VkShaderCorePropertiesFlagBitsAMD = 0x7FFFFFFF
)

type VkToolPurposeFlagBitsEXT = uint32

const (
	VK_TOOL_PURPOSE_VALIDATION_BIT_EXT          VkToolPurposeFlagBitsEXT = 0x00000001
	VK_TOOL_PURPOSE_PROFILING_BIT_EXT           VkToolPurposeFlagBitsEXT = 0x00000002
	VK_TOOL_PURPOSE_TRACING_BIT_EXT             VkToolPurposeFlagBitsEXT = 0x00000004
	VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT_EXT VkToolPurposeFlagBitsEXT = 0x00000008
	VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT_EXT  VkToolPurposeFlagBitsEXT = 0x00000010
	VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT     VkToolPurposeFlagBitsEXT = 0x00000020
	VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT       VkToolPurposeFlagBitsEXT = 0x00000040
	VK_TOOL_PURPOSE_FLAG_BITS_MAX_ENUM_EXT      VkToolPurposeFlagBitsEXT = 0x7FFFFFFF
)

type VkValidationFeatureEnableEXT = uint32

const (
	VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT                      VkValidationFeatureEnableEXT = 0
	VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT VkValidationFeatureEnableEXT = 1
	VK_VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT                    VkValidationFeatureEnableEXT = 2
	VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT                      VkValidationFeatureEnableEXT = 3
	VK_VALIDATION_FEATURE_ENABLE_SYNCHRONIZATION_VALIDATION_EXT        VkValidationFeatureEnableEXT = 4
	VK_VALIDATION_FEATURE_ENABLE_MAX_ENUM_EXT                          VkValidationFeatureEnableEXT = 0x7FFFFFFF
)

type VkValidationFeatureDisableEXT = uint32

const (
	VK_VALIDATION_FEATURE_DISABLE_ALL_EXT                     VkValidationFeatureDisableEXT = 0
	VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT                 VkValidationFeatureDisableEXT = 1
	VK_VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT           VkValidationFeatureDisableEXT = 2
	VK_VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT          VkValidationFeatureDisableEXT = 3
	VK_VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT        VkValidationFeatureDisableEXT = 4
	VK_VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT             VkValidationFeatureDisableEXT = 5
	VK_VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT          VkValidationFeatureDisableEXT = 6
	VK_VALIDATION_FEATURE_DISABLE_SHADER_VALIDATION_CACHE_EXT VkValidationFeatureDisableEXT = 7
	VK_VALIDATION_FEATURE_DISABLE_MAX_ENUM_EXT                VkValidationFeatureDisableEXT = 0x7FFFFFFF
)

type VkComponentTypeNV = uint32

const (
	VK_COMPONENT_TYPE_FLOAT16_NV  VkComponentTypeNV = 0
	VK_COMPONENT_TYPE_FLOAT32_NV  VkComponentTypeNV = 1
	VK_COMPONENT_TYPE_FLOAT64_NV  VkComponentTypeNV = 2
	VK_COMPONENT_TYPE_SINT8_NV    VkComponentTypeNV = 3
	VK_COMPONENT_TYPE_SINT16_NV   VkComponentTypeNV = 4
	VK_COMPONENT_TYPE_SINT32_NV   VkComponentTypeNV = 5
	VK_COMPONENT_TYPE_SINT64_NV   VkComponentTypeNV = 6
	VK_COMPONENT_TYPE_UINT8_NV    VkComponentTypeNV = 7
	VK_COMPONENT_TYPE_UINT16_NV   VkComponentTypeNV = 8
	VK_COMPONENT_TYPE_UINT32_NV   VkComponentTypeNV = 9
	VK_COMPONENT_TYPE_UINT64_NV   VkComponentTypeNV = 10
	VK_COMPONENT_TYPE_MAX_ENUM_NV VkComponentTypeNV = 0x7FFFFFFF
)

type VkScopeNV = uint32

const (
	VK_SCOPE_DEVICE_NV       VkScopeNV = 1
	VK_SCOPE_WORKGROUP_NV    VkScopeNV = 2
	VK_SCOPE_SUBGROUP_NV     VkScopeNV = 3
	VK_SCOPE_QUEUE_FAMILY_NV VkScopeNV = 5
	VK_SCOPE_MAX_ENUM_NV     VkScopeNV = 0x7FFFFFFF
)

type VkCoverageReductionModeNV = uint32

const (
	VK_COVERAGE_REDUCTION_MODE_MERGE_NV    VkCoverageReductionModeNV = 0
	VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV VkCoverageReductionModeNV = 1
	VK_COVERAGE_REDUCTION_MODE_MAX_ENUM_NV VkCoverageReductionModeNV = 0x7FFFFFFF
)

type VkProvokingVertexModeEXT = uint32

const (
	VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT VkProvokingVertexModeEXT = 0
	VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT  VkProvokingVertexModeEXT = 1
	VK_PROVOKING_VERTEX_MODE_MAX_ENUM_EXT     VkProvokingVertexModeEXT = 0x7FFFFFFF
)

type VkLineRasterizationModeEXT = uint32

const (
	VK_LINE_RASTERIZATION_MODE_DEFAULT_EXT            VkLineRasterizationModeEXT = 0
	VK_LINE_RASTERIZATION_MODE_RECTANGULAR_EXT        VkLineRasterizationModeEXT = 1
	VK_LINE_RASTERIZATION_MODE_BRESENHAM_EXT          VkLineRasterizationModeEXT = 2
	VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT VkLineRasterizationModeEXT = 3
	VK_LINE_RASTERIZATION_MODE_MAX_ENUM_EXT           VkLineRasterizationModeEXT = 0x7FFFFFFF
)

type VkIndirectCommandsTokenTypeNV = uint32

const (
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV  VkIndirectCommandsTokenTypeNV = 0
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV   VkIndirectCommandsTokenTypeNV = 1
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV  VkIndirectCommandsTokenTypeNV = 2
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV VkIndirectCommandsTokenTypeNV = 3
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV VkIndirectCommandsTokenTypeNV = 4
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV  VkIndirectCommandsTokenTypeNV = 5
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV          VkIndirectCommandsTokenTypeNV = 6
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV    VkIndirectCommandsTokenTypeNV = 7
	VK_INDIRECT_COMMANDS_TOKEN_TYPE_MAX_ENUM_NV      VkIndirectCommandsTokenTypeNV = 0x7FFFFFFF
)

type VkIndirectStateFlagBitsNV = uint32

const (
	VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV VkIndirectStateFlagBitsNV = 0x00000001
	VK_INDIRECT_STATE_FLAG_BITS_MAX_ENUM_NV VkIndirectStateFlagBitsNV = 0x7FFFFFFF
)

type VkIndirectCommandsLayoutUsageFlagBitsNV = uint32

const (
	VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV VkIndirectCommandsLayoutUsageFlagBitsNV = 0x00000001
	VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV   VkIndirectCommandsLayoutUsageFlagBitsNV = 0x00000002
	VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV VkIndirectCommandsLayoutUsageFlagBitsNV = 0x00000004
	VK_INDIRECT_COMMANDS_LAYOUT_USAGE_FLAG_BITS_MAX_ENUM_NV      VkIndirectCommandsLayoutUsageFlagBitsNV = 0x7FFFFFFF
)

type VkDeviceMemoryReportEventTypeEXT = uint32

const (
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT          VkDeviceMemoryReportEventTypeEXT = 0
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT              VkDeviceMemoryReportEventTypeEXT = 1
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT            VkDeviceMemoryReportEventTypeEXT = 2
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT          VkDeviceMemoryReportEventTypeEXT = 3
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT VkDeviceMemoryReportEventTypeEXT = 4
	VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_MAX_ENUM_EXT          VkDeviceMemoryReportEventTypeEXT = 0x7FFFFFFF
)

type VkPrivateDataSlotCreateFlagBitsEXT = uint32

const (
	VK_PRIVATE_DATA_SLOT_CREATE_FLAG_BITS_MAX_ENUM_EXT VkPrivateDataSlotCreateFlagBitsEXT = 0x7FFFFFFF
)

type VkDeviceDiagnosticsConfigFlagBitsNV = uint32

const (
	VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV     VkDeviceDiagnosticsConfigFlagBitsNV = 0x00000001
	VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV     VkDeviceDiagnosticsConfigFlagBitsNV = 0x00000002
	VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV VkDeviceDiagnosticsConfigFlagBitsNV = 0x00000004
	VK_DEVICE_DIAGNOSTICS_CONFIG_FLAG_BITS_MAX_ENUM_NV               VkDeviceDiagnosticsConfigFlagBitsNV = 0x7FFFFFFF
)

type VkFragmentShadingRateTypeNV = uint32

const (
	VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV VkFragmentShadingRateTypeNV = 0
	VK_FRAGMENT_SHADING_RATE_TYPE_ENUMS_NV         VkFragmentShadingRateTypeNV = 1
	VK_FRAGMENT_SHADING_RATE_TYPE_MAX_ENUM_NV      VkFragmentShadingRateTypeNV = 0x7FFFFFFF
)

type VkFragmentShadingRateNV = uint32

const (
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV      VkFragmentShadingRateNV = 0
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_1X2_PIXELS_NV VkFragmentShadingRateNV = 1
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X1_PIXELS_NV VkFragmentShadingRateNV = 4
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X2_PIXELS_NV VkFragmentShadingRateNV = 5
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X4_PIXELS_NV VkFragmentShadingRateNV = 6
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X2_PIXELS_NV VkFragmentShadingRateNV = 9
	VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X4_PIXELS_NV VkFragmentShadingRateNV = 10
	VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV     VkFragmentShadingRateNV = 11
	VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV     VkFragmentShadingRateNV = 12
	VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV     VkFragmentShadingRateNV = 13
	VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV    VkFragmentShadingRateNV = 14
	VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV              VkFragmentShadingRateNV = 15
	VK_FRAGMENT_SHADING_RATE_MAX_ENUM_NV                    VkFragmentShadingRateNV = 0x7FFFFFFF
)

type VkAccelerationStructureMotionInstanceTypeNV = uint32

const (
	VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV        VkAccelerationStructureMotionInstanceTypeNV = 0
	VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV VkAccelerationStructureMotionInstanceTypeNV = 1
	VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV    VkAccelerationStructureMotionInstanceTypeNV = 2
	VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MAX_ENUM_NV      VkAccelerationStructureMotionInstanceTypeNV = 0x7FFFFFFF
)

type VkBuildAccelerationStructureModeKHR = uint32

const (
	VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR    VkBuildAccelerationStructureModeKHR = 0
	VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR   VkBuildAccelerationStructureModeKHR = 1
	VK_BUILD_ACCELERATION_STRUCTURE_MODE_MAX_ENUM_KHR VkBuildAccelerationStructureModeKHR = 0x7FFFFFFF
)

type VkAccelerationStructureBuildTypeKHR = uint32

const (
	VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR           VkAccelerationStructureBuildTypeKHR = 0
	VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR         VkAccelerationStructureBuildTypeKHR = 1
	VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR VkAccelerationStructureBuildTypeKHR = 2
	VK_ACCELERATION_STRUCTURE_BUILD_TYPE_MAX_ENUM_KHR       VkAccelerationStructureBuildTypeKHR = 0x7FFFFFFF
)

type VkAccelerationStructureCompatibilityKHR = uint32

const (
	VK_ACCELERATION_STRUCTURE_COMPATIBILITY_COMPATIBLE_KHR   VkAccelerationStructureCompatibilityKHR = 0
	VK_ACCELERATION_STRUCTURE_COMPATIBILITY_INCOMPATIBLE_KHR VkAccelerationStructureCompatibilityKHR = 1
	VK_ACCELERATION_STRUCTURE_COMPATIBILITY_MAX_ENUM_KHR     VkAccelerationStructureCompatibilityKHR = 0x7FFFFFFF
)

type VkAccelerationStructureCreateFlagBitsKHR = uint32

const (
	VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR VkAccelerationStructureCreateFlagBitsKHR = 0x00000001
	VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV                         VkAccelerationStructureCreateFlagBitsKHR = 0x00000004
	VK_ACCELERATION_STRUCTURE_CREATE_FLAG_BITS_MAX_ENUM_KHR                VkAccelerationStructureCreateFlagBitsKHR = 0x7FFFFFFF
)

type VkShaderGroupShaderKHR = uint32

const (
	VK_SHADER_GROUP_SHADER_GENERAL_KHR      VkShaderGroupShaderKHR = 0
	VK_SHADER_GROUP_SHADER_CLOSEST_HIT_KHR  VkShaderGroupShaderKHR = 1
	VK_SHADER_GROUP_SHADER_ANY_HIT_KHR      VkShaderGroupShaderKHR = 2
	VK_SHADER_GROUP_SHADER_INTERSECTION_KHR VkShaderGroupShaderKHR = 3
	VK_SHADER_GROUP_SHADER_MAX_ENUM_KHR     VkShaderGroupShaderKHR = 0x7FFFFFFF
)

type VkClearColorValue struct {
	value [16]byte
}

func (this *VkClearColorValue) GetFloat32() (float [4]float32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, float[0])
	binary.Read(buffer, binary.BigEndian, float[1])
	binary.Read(buffer, binary.BigEndian, float[2])
	binary.Read(buffer, binary.BigEndian, float[3])
	return
}
func (this *VkClearColorValue) SetFloat32(float [4]float32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, float[0])
	binary.Write(buffer, binary.LittleEndian, float[1])
	binary.Write(buffer, binary.LittleEndian, float[2])
	binary.Write(buffer, binary.LittleEndian, float[3])
}
func (this *VkClearColorValue) GetInt32() (int32_t [4]int32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, int32_t[0])
	binary.Read(buffer, binary.BigEndian, int32_t[1])
	binary.Read(buffer, binary.BigEndian, int32_t[2])
	binary.Read(buffer, binary.BigEndian, int32_t[3])
	return
}
func (this *VkClearColorValue) SetInt32(int32_t [4]int32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, int32_t[0])
	binary.Write(buffer, binary.LittleEndian, int32_t[1])
	binary.Write(buffer, binary.LittleEndian, int32_t[2])
	binary.Write(buffer, binary.LittleEndian, int32_t[3])
}
func (this *VkClearColorValue) GetUint32() (uint32_t [4]uint32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, uint32_t[0])
	binary.Read(buffer, binary.BigEndian, uint32_t[1])
	binary.Read(buffer, binary.BigEndian, uint32_t[2])
	binary.Read(buffer, binary.BigEndian, uint32_t[3])
	return
}
func (this *VkClearColorValue) SetUint32(uint32_t [4]uint32) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, uint32_t[0])
	binary.Write(buffer, binary.LittleEndian, uint32_t[1])
	binary.Write(buffer, binary.LittleEndian, uint32_t[2])
	binary.Write(buffer, binary.LittleEndian, uint32_t[3])
}

type VkClearValue struct {
	value [16]byte
}

func (this *VkClearValue) GetColor() (color VkClearColorValue) {
	copy(color.value[:], this.value[:])
	return
}
func (this *VkClearValue) SetColor(color VkClearColorValue) {
	copy(this.value[:], color.value[:])
}
func (this *VkClearValue) GetDepthStencil() (depthStencil VkClearDepthStencilValue) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, depthStencil.Depth)
	binary.Write(buffer, binary.LittleEndian, depthStencil.Stencil)
	return
}
func (this *VkClearValue) SetDepthStencil(depthStencil VkClearDepthStencilValue) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, depthStencil.Depth)
	binary.Write(buffer, binary.LittleEndian, depthStencil.Stencil)
}

type VkPerformanceCounterResultKHR struct {
	value [8]byte
}

func (this *VkPerformanceCounterResultKHR) GetInt32() (int32_t int32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, int32_t)
	return
}
func (this *VkPerformanceCounterResultKHR) SetInt32(int32_t int32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, int32_t)
}
func (this *VkPerformanceCounterResultKHR) GetInt64() (int64_t int64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, int64_t)
	return
}
func (this *VkPerformanceCounterResultKHR) SetInt64(int64_t int64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, int64_t)
}
func (this *VkPerformanceCounterResultKHR) GetUint32() (uint32_t uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, uint32_t)
	return
}
func (this *VkPerformanceCounterResultKHR) SetUint32(uint32_t uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, uint32_t)
}
func (this *VkPerformanceCounterResultKHR) GetUint64() (uint64_t uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, uint64_t)
	return
}
func (this *VkPerformanceCounterResultKHR) SetUint64(uint64_t uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, uint64_t)
}
func (this *VkPerformanceCounterResultKHR) GetFloat32() (float float32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, float)
	return
}
func (this *VkPerformanceCounterResultKHR) SetFloat32(float float32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, float)
}
func (this *VkPerformanceCounterResultKHR) GetFloat64() (double float64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, double)
	return
}
func (this *VkPerformanceCounterResultKHR) SetFloat64(double float64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, double)
}

type VkPipelineExecutableStatisticValueKHR struct {
	value [8]byte
}

func (this *VkPipelineExecutableStatisticValueKHR) GetVkBool32() (VkBool32 uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, VkBool32)
	return
}
func (this *VkPipelineExecutableStatisticValueKHR) SetVkBool32(VkBool32 uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, VkBool32)
}
func (this *VkPipelineExecutableStatisticValueKHR) GetInt64() (int64_t int64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, int64_t)
	return
}
func (this *VkPipelineExecutableStatisticValueKHR) SetInt64(int64_t int64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, int64_t)
}
func (this *VkPipelineExecutableStatisticValueKHR) GetUint64() (uint64_t uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, uint64_t)
	return
}
func (this *VkPipelineExecutableStatisticValueKHR) SetUint64(uint64_t uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, uint64_t)
}
func (this *VkPipelineExecutableStatisticValueKHR) GetFloat64() (double float64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, double)
	return
}
func (this *VkPipelineExecutableStatisticValueKHR) SetFloat64(double float64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, double)
}

type VkPerformanceValueDataINTEL struct {
	value [8]byte
}

func (this *VkPerformanceValueDataINTEL) GetUint32() (value32 uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, value32)
	return
}
func (this *VkPerformanceValueDataINTEL) SetUint32(value32 uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, value32)
}
func (this *VkPerformanceValueDataINTEL) GetUint64() (value64 uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, value64)
	return
}
func (this *VkPerformanceValueDataINTEL) SetUint64(value64 uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, value64)
}
func (this *VkPerformanceValueDataINTEL) GetFloat32() (valueFloat float32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, valueFloat)
	return
}
func (this *VkPerformanceValueDataINTEL) SetFloat32(valueFloat float32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, valueFloat)
}
func (this *VkPerformanceValueDataINTEL) GetVkBool32() (valueBool uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Read(buffer, binary.BigEndian, valueBool)
	return
}
func (this *VkPerformanceValueDataINTEL) SetVkBool32(valueBool uint32) {
	var buffer = bytes.NewBuffer(this.value[:4])
	binary.Write(buffer, binary.LittleEndian, valueBool)
}
func (this *VkPerformanceValueDataINTEL) GetVkString() (valueString unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, valueString)
	return
}
func (this *VkPerformanceValueDataINTEL) SetVkString(valueString unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, valueString)
}

type VkDeviceOrHostAddressConstKHR struct {
	value [8]byte
}

func (this *VkDeviceOrHostAddressConstKHR) GetVkDeviceAddress() (deviceAddress uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, deviceAddress)
	return
}
func (this *VkDeviceOrHostAddressConstKHR) SetVkDeviceAddress(deviceAddress uint64) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, deviceAddress)
}
func (this *VkDeviceOrHostAddressConstKHR) GetPointor() (hostAddress unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, hostAddress)
	return
}
func (this *VkDeviceOrHostAddressConstKHR) SetPointor(hostAddress unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, hostAddress)
}

type VkAccelerationStructureMotionInstanceDataNV struct {
	value [144]byte
}

func (this *VkAccelerationStructureMotionInstanceDataNV) GetVkAccelerationStructureInstanceKHR() (staticInstance VkAccelerationStructureInstanceKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, staticInstance.Transform)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Read(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
	return
}
func (this *VkAccelerationStructureMotionInstanceDataNV) SetVkAccelerationStructureInstanceKHR(staticInstance VkAccelerationStructureInstanceKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, staticInstance.Transform)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Write(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
}
func (this *VkAccelerationStructureMotionInstanceDataNV) GetVkAccelerationStructureMatrixMotionInstanceNV() (staticInstance VkAccelerationStructureMatrixMotionInstanceNV) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, staticInstance.TransformT0)
	binary.Read(buffer, binary.LittleEndian, staticInstance.TransformT1)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Read(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
	return
}
func (this *VkAccelerationStructureMotionInstanceDataNV) SetVkAccelerationStructureMatrixMotionInstanceNV(staticInstance VkAccelerationStructureMatrixMotionInstanceNV) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, staticInstance.TransformT0)
	binary.Write(buffer, binary.LittleEndian, staticInstance.TransformT1)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Write(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
}
func (this *VkAccelerationStructureMotionInstanceDataNV) GetVkAccelerationStructureSRTMotionInstanceNV() (staticInstance VkAccelerationStructureSRTMotionInstanceNV) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, staticInstance.TransformT0)
	binary.Read(buffer, binary.LittleEndian, staticInstance.TransformT1)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Read(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Read(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Read(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
	return
}
func (this *VkAccelerationStructureMotionInstanceDataNV) SetVkAccelerationStructureSRTMotionInstanceNV(staticInstance VkAccelerationStructureSRTMotionInstanceNV) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, staticInstance.TransformT0)
	binary.Write(buffer, binary.LittleEndian, staticInstance.TransformT1)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceCustomIndex)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Mask)
	binary.Write(buffer, binary.LittleEndian, staticInstance.InstanceShaderBindingTableRecordOffset)
	binary.Write(buffer, binary.LittleEndian, staticInstance.Flags)
	binary.Write(buffer, binary.LittleEndian, staticInstance.AccelerationStructureReference)
}

type VkDeviceOrHostAddressKHR struct {
	value [8]byte
}

func (this *VkDeviceOrHostAddressKHR) GetVkDeviceAddress() (deviceAddress VkDeviceAddress) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, deviceAddress)
	return
}
func (this *VkDeviceOrHostAddressKHR) SetVkDeviceAddress(deviceAddress VkDeviceAddress) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, deviceAddress)
}
func (this *VkDeviceOrHostAddressKHR) GetPointor() (hostAddress unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.BigEndian, hostAddress)
	return
}
func (this *VkDeviceOrHostAddressKHR) SetPointor(hostAddress unsafe.Pointer) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, hostAddress)
}

type VkAccelerationStructureGeometryDataKHR struct {
	value [64]byte
}

func (this *VkAccelerationStructureGeometryDataKHR) GetVkAccelerationStructureGeometryTrianglesDataKHR() (triangles VkAccelerationStructureGeometryTrianglesDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, triangles.SType)
	binary.Read(buffer, binary.LittleEndian, triangles.PNext)
	binary.Read(buffer, binary.LittleEndian, triangles.VertexFormat)
	binary.Read(buffer, binary.LittleEndian, triangles.VertexData)
	binary.Read(buffer, binary.LittleEndian, triangles.VertexStride)
	binary.Read(buffer, binary.LittleEndian, triangles.MaxVertex)
	binary.Read(buffer, binary.LittleEndian, triangles.IndexType)
	binary.Read(buffer, binary.LittleEndian, triangles.IndexData)
	binary.Read(buffer, binary.LittleEndian, triangles.TransformData)
	return
}
func (this *VkAccelerationStructureGeometryDataKHR) SetVkAccelerationStructureGeometryTrianglesDataKHR(triangles VkAccelerationStructureGeometryTrianglesDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, triangles.SType)
	binary.Write(buffer, binary.LittleEndian, triangles.PNext)
	binary.Write(buffer, binary.LittleEndian, triangles.VertexFormat)
	binary.Write(buffer, binary.LittleEndian, triangles.VertexData)
	binary.Write(buffer, binary.LittleEndian, triangles.VertexStride)
	binary.Write(buffer, binary.LittleEndian, triangles.MaxVertex)
	binary.Write(buffer, binary.LittleEndian, triangles.IndexType)
	binary.Write(buffer, binary.LittleEndian, triangles.IndexData)
	binary.Write(buffer, binary.LittleEndian, triangles.TransformData)
}
func (this *VkAccelerationStructureGeometryDataKHR) GetVkAccelerationStructureGeometryAabbsDataKHR() (aabbs VkAccelerationStructureGeometryAabbsDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, aabbs.SType)
	binary.Read(buffer, binary.LittleEndian, aabbs.PNext)
	binary.Read(buffer, binary.LittleEndian, aabbs.Data)
	binary.Read(buffer, binary.LittleEndian, aabbs.Stride)
	return
}
func (this *VkAccelerationStructureGeometryDataKHR) SetVkAccelerationStructureGeometryAabbsDataKHR(aabbs VkAccelerationStructureGeometryAabbsDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, aabbs.SType)
	binary.Write(buffer, binary.LittleEndian, aabbs.PNext)
	binary.Write(buffer, binary.LittleEndian, aabbs.Data)
	binary.Write(buffer, binary.LittleEndian, aabbs.Stride)
}
func (this *VkAccelerationStructureGeometryDataKHR) GetVkAccelerationStructureGeometryInstancesDataKHR() (instances VkAccelerationStructureGeometryInstancesDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Read(buffer, binary.LittleEndian, instances.SType)
	binary.Read(buffer, binary.LittleEndian, instances.PNext)
	binary.Read(buffer, binary.LittleEndian, instances.ArrayOfPointers)
	binary.Read(buffer, binary.LittleEndian, instances.Data)
	return
}
func (this *VkAccelerationStructureGeometryDataKHR) SetVkAccelerationStructureGeometryInstancesDataKHR(instances VkAccelerationStructureGeometryInstancesDataKHR) {
	var buffer = bytes.NewBuffer(this.value[:])
	binary.Write(buffer, binary.LittleEndian, instances.SType)
	binary.Write(buffer, binary.LittleEndian, instances.PNext)
	binary.Write(buffer, binary.LittleEndian, instances.ArrayOfPointers)
	binary.Write(buffer, binary.LittleEndian, instances.Data)
}

type VkExtent2D struct {
	Width  uint32
	Height uint32
}
type VkExtent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}
type VkOffset2D struct {
	X int32
	Y int32
}
type VkOffset3D struct {
	X int32
	Y int32
	Z int32
}
type VkRect2D struct {
	Offset VkOffset2D
	Extent VkExtent2D
}
type VkBaseInStructure struct {
	SType VkStructureType
}
type VkBaseOutStructure struct {
	SType VkStructureType
}
type VkBufferMemoryBarrier struct {
	SType               VkStructureType
	PNext               uintptr
	SrcAccessMask       VkAccessFlags
	DstAccessMask       VkAccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              VkBuffer
	Offset              VkDeviceSize
	Size                VkDeviceSize
}
type VkDispatchIndirectCommand struct {
	X uint32
	Y uint32
	Z uint32
}
type VkDrawIndexedIndirectCommand struct {
	IndexCount    uint32
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}
type VkDrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}
type VkImageSubresourceRange struct {
	AspectMask     VkImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}
type VkImageMemoryBarrier struct {
	SType               VkStructureType
	PNext               uintptr
	SrcAccessMask       VkAccessFlags
	DstAccessMask       VkAccessFlags
	OldLayout           VkImageLayout
	NewLayout           VkImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               VkImage
	SubresourceRange    VkImageSubresourceRange
}
type VkMemoryBarrier struct {
	SType         VkStructureType
	PNext         uintptr
	SrcAccessMask VkAccessFlags
	DstAccessMask VkAccessFlags
}
type VkPipelineCacheHeaderVersionOne struct {
	HeaderSize        uint32
	HeaderVersion     VkPipelineCacheHeaderVersion
	VendorID          uint32
	DeviceID          uint32
	PipelineCacheUUID [16]uint8
}
type VkAllocationCallbacks struct {
	PUserData             uintptr
	PfnAllocation         uintptr
	PfnReallocation       uintptr
	PfnFree               uintptr
	PfnInternalAllocation uintptr
	PfnInternalFree       uintptr
}
type VkApplicationInfo struct {
	SType              VkStructureType
	PNext              uintptr
	PApplicationName   *byte
	ApplicationVersion uint32
	PEngineName        *byte
	EngineVersion      uint32
	ApiVersion         uint32
}
type VkFormatProperties struct {
	LinearTilingFeatures  VkFormatFeatureFlags
	OptimalTilingFeatures VkFormatFeatureFlags
	BufferFeatures        VkFormatFeatureFlags
}
type VkImageFormatProperties struct {
	MaxExtent       VkExtent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    VkSampleCountFlags
	MaxResourceSize VkDeviceSize
}
type VkInstanceCreateInfo struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkInstanceCreateFlags
	PApplicationInfo        *VkApplicationInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     **byte
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames **byte
}
type VkMemoryHeap struct {
	Size  VkDeviceSize
	Flags VkMemoryHeapFlags
}
type VkMemoryType struct {
	PropertyFlags VkMemoryPropertyFlags
	HeapIndex     uint32
}
type VkPhysicalDeviceFeatures struct {
	RobustBufferAccess                      VkBool32
	FullDrawIndexUint32                     VkBool32
	ImageCubeArray                          VkBool32
	IndependentBlend                        VkBool32
	GeometryShader                          VkBool32
	TessellationShader                      VkBool32
	SampleRateShading                       VkBool32
	DualSrcBlend                            VkBool32
	LogicOp                                 VkBool32
	MultiDrawIndirect                       VkBool32
	DrawIndirectFirstInstance               VkBool32
	DepthClamp                              VkBool32
	DepthBiasClamp                          VkBool32
	FillModeNonSolid                        VkBool32
	DepthBounds                             VkBool32
	WideLines                               VkBool32
	LargePoints                             VkBool32
	AlphaToOne                              VkBool32
	MultiViewport                           VkBool32
	SamplerAnisotropy                       VkBool32
	TextureCompressionETC2                  VkBool32
	TextureCompressionASTC_LDR              VkBool32
	TextureCompressionBC                    VkBool32
	OcclusionQueryPrecise                   VkBool32
	PipelineStatisticsQuery                 VkBool32
	VertexPipelineStoresAndAtomics          VkBool32
	FragmentStoresAndAtomics                VkBool32
	ShaderTessellationAndGeometryPointSize  VkBool32
	ShaderImageGatherExtended               VkBool32
	ShaderStorageImageExtendedFormats       VkBool32
	ShaderStorageImageMultisample           VkBool32
	ShaderStorageImageReadWithoutFormat     VkBool32
	ShaderStorageImageWriteWithoutFormat    VkBool32
	ShaderUniformBufferArrayDynamicIndexing VkBool32
	ShaderSampledImageArrayDynamicIndexing  VkBool32
	ShaderStorageBufferArrayDynamicIndexing VkBool32
	ShaderStorageImageArrayDynamicIndexing  VkBool32
	ShaderClipDistance                      VkBool32
	ShaderCullDistance                      VkBool32
	ShaderFloat64                           VkBool32
	ShaderInt64                             VkBool32
	ShaderInt16                             VkBool32
	ShaderResourceResidency                 VkBool32
	ShaderResourceMinLod                    VkBool32
	SparseBinding                           VkBool32
	SparseResidencyBuffer                   VkBool32
	SparseResidencyImage2D                  VkBool32
	SparseResidencyImage3D                  VkBool32
	SparseResidency2Samples                 VkBool32
	SparseResidency4Samples                 VkBool32
	SparseResidency8Samples                 VkBool32
	SparseResidency16Samples                VkBool32
	SparseResidencyAliased                  VkBool32
	VariableMultisampleRate                 VkBool32
	InheritedQueries                        VkBool32
}
type VkPhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxMemoryAllocationCount                        uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          VkDeviceSize
	SparseAddressSpaceSize                          VkDeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uint64
	MinTexelBufferOffsetAlignment                   VkDeviceSize
	MinUniformBufferOffsetAlignment                 VkDeviceSize
	MinStorageBufferOffsetAlignment                 VkDeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    VkSampleCountFlags
	FramebufferDepthSampleCounts                    VkSampleCountFlags
	FramebufferStencilSampleCounts                  VkSampleCountFlags
	FramebufferNoAttachmentsSampleCounts            VkSampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   VkSampleCountFlags
	SampledImageIntegerSampleCounts                 VkSampleCountFlags
	SampledImageDepthSampleCounts                   VkSampleCountFlags
	SampledImageStencilSampleCounts                 VkSampleCountFlags
	StorageImageSampleCounts                        VkSampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     VkBool32
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     VkBool32
	StandardSampleLocations                         VkBool32
	OptimalBufferCopyOffsetAlignment                VkDeviceSize
	OptimalBufferCopyRowPitchAlignment              VkDeviceSize
	NonCoherentAtomSize                             VkDeviceSize
}
type VkPhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]VkMemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]VkMemoryHeap
}
type VkPhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            VkBool32
	ResidencyStandard2DMultisampleBlockShape VkBool32
	ResidencyStandard3DBlockShape            VkBool32
	ResidencyAlignedMipSize                  VkBool32
	ResidencyNonResidentStrict               VkBool32
}
type VkPhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        VkPhysicalDeviceType
	DeviceName        [256]byte
	PipelineCacheUUID [16]uint8
	Limits            VkPhysicalDeviceLimits
	SparseProperties  VkPhysicalDeviceSparseProperties
}
type VkQueueFamilyProperties struct {
	QueueFlags                  VkQueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity VkExtent3D
}
type VkDeviceQueueCreateInfo struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	PQueuePriorities *float32
}
type VkDeviceCreateInfo struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkDeviceCreateFlags
	QueueCreateInfoCount    uint32
	PQueueCreateInfos       *VkDeviceQueueCreateInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     **byte
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames **byte
	PEnabledFeatures        *VkPhysicalDeviceFeatures
}
type VkExtensionProperties struct {
	ExtensionName [256]byte
	SpecVersion   uint32
}
type VkLayerProperties struct {
	LayerName             [256]byte
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [256]byte
}
type VkSubmitInfo struct {
	SType                VkStructureType
	PNext                uintptr
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *VkSemaphore
	PWaitDstStageMask    *VkPipelineStageFlags
	CommandBufferCount   uint32
	PCommandBuffers      *VkCommandBuffer
	SignalSemaphoreCount uint32
	PSignalSemaphores    *VkSemaphore
}
type VkMappedMemoryRange struct {
	SType  VkStructureType
	PNext  uintptr
	Memory VkDeviceMemory
	Offset VkDeviceSize
	Size   VkDeviceSize
}
type VkMemoryAllocateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	AllocationSize  VkDeviceSize
	MemoryTypeIndex uint32
}
type VkMemoryRequirements struct {
	Size           VkDeviceSize
	Alignment      VkDeviceSize
	MemoryTypeBits uint32
}
type VkSparseMemoryBind struct {
	ResourceOffset VkDeviceSize
	Size           VkDeviceSize
	Memory         VkDeviceMemory
	MemoryOffset   VkDeviceSize
	Flags          VkSparseMemoryBindFlags
}
type VkSparseBufferMemoryBindInfo struct {
	Buffer    VkBuffer
	BindCount uint32
	PBinds    *VkSparseMemoryBind
}
type VkSparseImageOpaqueMemoryBindInfo struct {
	Image     VkImage
	BindCount uint32
	PBinds    *VkSparseMemoryBind
}
type VkImageSubresource struct {
	AspectMask VkImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}
type VkSparseImageMemoryBind struct {
	Subresource  VkImageSubresource
	Offset       VkOffset3D
	Extent       VkExtent3D
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
	Flags        VkSparseMemoryBindFlags
}
type VkSparseImageMemoryBindInfo struct {
	Image     VkImage
	BindCount uint32
	PBinds    *VkSparseImageMemoryBind
}
type VkBindSparseInfo struct {
	SType                VkStructureType
	PNext                uintptr
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *VkSemaphore
	BufferBindCount      uint32
	PBufferBinds         *VkSparseBufferMemoryBindInfo
	ImageOpaqueBindCount uint32
	PImageOpaqueBinds    *VkSparseImageOpaqueMemoryBindInfo
	ImageBindCount       uint32
	PImageBinds          *VkSparseImageMemoryBindInfo
	SignalSemaphoreCount uint32
	PSignalSemaphores    *VkSemaphore
}
type VkSparseImageFormatProperties struct {
	AspectMask       VkImageAspectFlags
	ImageGranularity VkExtent3D
	Flags            VkSparseImageFormatFlags
}
type VkSparseImageMemoryRequirements struct {
	FormatProperties     VkSparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     VkDeviceSize
	ImageMipTailOffset   VkDeviceSize
	ImageMipTailStride   VkDeviceSize
}
type VkFenceCreateInfo struct {
	SType VkStructureType
	PNext uintptr
	Flags VkFenceCreateFlags
}
type VkSemaphoreCreateInfo struct {
	SType VkStructureType
	PNext uintptr
	Flags VkSemaphoreCreateFlags
}
type VkEventCreateInfo struct {
	SType VkStructureType
	PNext uintptr
	Flags VkEventCreateFlags
}
type VkQueryPoolCreateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	Flags              VkQueryPoolCreateFlags
	QueryType          VkQueryType
	QueryCount         uint32
	PipelineStatistics VkQueryPipelineStatisticFlags
}
type VkBufferCreateInfo struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkBufferCreateFlags
	Size                  VkDeviceSize
	Usage                 VkBufferUsageFlags
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}
type VkBufferViewCreateInfo struct {
	SType  VkStructureType
	PNext  uintptr
	Flags  VkBufferViewCreateFlags
	Buffer VkBuffer
	Format VkFormat
	Offset VkDeviceSize
	Range  VkDeviceSize
}
type VkImageCreateInfo struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkImageCreateFlags
	ImageType             VkImageType
	Format                VkFormat
	Extent                VkExtent3D
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               VkSampleCountFlagBits
	Tiling                VkImageTiling
	Usage                 VkImageUsageFlags
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	InitialLayout         VkImageLayout
}
type VkSubresourceLayout struct {
	Offset     VkDeviceSize
	Size       VkDeviceSize
	RowPitch   VkDeviceSize
	ArrayPitch VkDeviceSize
	DepthPitch VkDeviceSize
}
type VkComponentMapping struct {
	R VkComponentSwizzle
	G VkComponentSwizzle
	B VkComponentSwizzle
	A VkComponentSwizzle
}
type VkImageViewCreateInfo struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkImageViewCreateFlags
	Image            VkImage
	ViewType         VkImageViewType
	Format           VkFormat
	Components       VkComponentMapping
	SubresourceRange VkImageSubresourceRange
}
type VkShaderModuleCreateInfo struct {
	SType    VkStructureType
	PNext    uintptr
	Flags    VkShaderModuleCreateFlags
	CodeSize uint64
	PCode    *uint32
}
type VkPipelineCacheCreateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkPipelineCacheCreateFlags
	InitialDataSize uint64
	PInitialData    uintptr
}
type VkSpecializationMapEntry struct {
	AntID  uint32
	Offset uint32
	Size   uint64
}
type VkSpecializationInfo struct {
	MapEntryCount uint32
	PMapEntries   *VkSpecializationMapEntry
	DataSize      uint64
	PData         uintptr
}
type VkPipelineShaderStageCreateInfo struct {
	SType               VkStructureType
	PNext               uintptr
	Flags               VkPipelineShaderStageCreateFlags
	Stage               VkShaderStageFlagBits
	Module              VkShaderModule
	PName               *byte
	PSpecializationInfo *VkSpecializationInfo
}
type VkComputePipelineCreateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	Flags              VkPipelineCreateFlags
	Stage              VkPipelineShaderStageCreateInfo
	Layout             VkPipelineLayout
	BasePipelineHandle VkPipeline
	BasePipelineIndex  int32
}
type VkVertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VkVertexInputRate
}
type VkVertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   VkFormat
	Offset   uint32
}
type VkPipelineVertexInputStateCreateInfo struct {
	SType                           VkStructureType
	PNext                           uintptr
	Flags                           VkPipelineVertexInputStateCreateFlags
	VertexBindingDescriptionCount   uint32
	PVertexBindingDescriptions      *VkVertexInputBindingDescription
	VertexAttributeDescriptionCount uint32
	PVertexAttributeDescriptions    *VkVertexInputAttributeDescription
}
type VkPipelineInputAssemblyStateCreateInfo struct {
	SType                  VkStructureType
	PNext                  uintptr
	Flags                  VkPipelineInputAssemblyStateCreateFlags
	Topology               VkPrimitiveTopology
	PrimitiveRestartEnable VkBool32
}
type VkPipelineTessellationStateCreateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	Flags              VkPipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}
type VkViewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}
type VkPipelineViewportStateCreateInfo struct {
	SType         VkStructureType
	PNext         uintptr
	Flags         VkPipelineViewportStateCreateFlags
	ViewportCount uint32
	PViewports    *VkViewport
	ScissorCount  uint32
	PScissors     *VkRect2D
}
type VkPipelineRasterizationStateCreateInfo struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkPipelineRasterizationStateCreateFlags
	DepthClampEnable        VkBool32
	RasterizerDiscardEnable VkBool32
	PolygonMode             VkPolygonMode
	CullMode                VkCullModeFlags
	FrontFace               VkFrontFace
	DepthBiasEnable         VkBool32
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}
type VkPipelineMultisampleStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkPipelineMultisampleStateCreateFlags
	RasterizationSamples  VkSampleCountFlagBits
	SampleShadingEnable   VkBool32
	MinSampleShading      float32
	PSampleMask           *VkSampleMask
	AlphaToCoverageEnable VkBool32
	AlphaToOneEnable      VkBool32
}
type VkStencilOpState struct {
	FailOp      VkStencilOp
	PassOp      VkStencilOp
	DepthFailOp VkStencilOp
	CompareOp   VkCompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}
type VkPipelineDepthStencilStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkPipelineDepthStencilStateCreateFlags
	DepthTestEnable       VkBool32
	DepthWriteEnable      VkBool32
	DepthCompareOp        VkCompareOp
	DepthBoundsTestEnable VkBool32
	StencilTestEnable     VkBool32
	Front                 VkStencilOpState
	Back                  VkStencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}
type VkPipelineColorBlendAttachmentState struct {
	BlendEnable         VkBool32
	SrcColorBlendFactor VkBlendFactor
	DstColorBlendFactor VkBlendFactor
	ColorBlendOp        VkBlendOp
	SrcAlphaBlendFactor VkBlendFactor
	DstAlphaBlendFactor VkBlendFactor
	AlphaBlendOp        VkBlendOp
	ColorWriteMask      VkColorComponentFlags
}
type VkPipelineColorBlendStateCreateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkPipelineColorBlendStateCreateFlags
	LogicOpEnable   VkBool32
	LogicOp         VkLogicOp
	AttachmentCount uint32
	PAttachments    *VkPipelineColorBlendAttachmentState
	BlendConstants  [4]float32
}
type VkPipelineDynamicStateCreateInfo struct {
	SType             VkStructureType
	PNext             uintptr
	Flags             VkPipelineDynamicStateCreateFlags
	DynamicStateCount uint32
	PDynamicStates    *VkDynamicState
}
type VkGraphicsPipelineCreateInfo struct {
	SType               VkStructureType
	PNext               uintptr
	Flags               VkPipelineCreateFlags
	StageCount          uint32
	PStages             *VkPipelineShaderStageCreateInfo
	PVertexInputState   *VkPipelineVertexInputStateCreateInfo
	PInputAssemblyState *VkPipelineInputAssemblyStateCreateInfo
	PTessellationState  *VkPipelineTessellationStateCreateInfo
	PViewportState      *VkPipelineViewportStateCreateInfo
	PRasterizationState *VkPipelineRasterizationStateCreateInfo
	PMultisampleState   *VkPipelineMultisampleStateCreateInfo
	PDepthStencilState  *VkPipelineDepthStencilStateCreateInfo
	PColorBlendState    *VkPipelineColorBlendStateCreateInfo
	PDynamicState       *VkPipelineDynamicStateCreateInfo
	Layout              VkPipelineLayout
	RenderPass          VkRenderPass
	Subpass             uint32
	BasePipelineHandle  VkPipeline
	BasePipelineIndex   int32
}
type VkPushConstantRange struct {
	StageFlags VkShaderStageFlags
	Offset     uint32
	Size       uint32
}
type VkPipelineLayoutCreateInfo struct {
	SType                  VkStructureType
	PNext                  uintptr
	Flags                  VkPipelineLayoutCreateFlags
	SetLayoutCount         uint32
	PSetLayouts            *VkDescriptorSetLayout
	PushConstantRangeCount uint32
	PPushConstantRanges    *VkPushConstantRange
}
type VkSamplerCreateInfo struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkSamplerCreateFlags
	MagFilter               VkFilter
	MinFilter               VkFilter
	MipmapMode              VkSamplerMipmapMode
	AddressModeU            VkSamplerAddressMode
	AddressModeV            VkSamplerAddressMode
	AddressModeW            VkSamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        VkBool32
	MaxAnisotropy           float32
	CompareEnable           VkBool32
	CompareOp               VkCompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             VkBorderColor
	UnnormalizedCoordinates VkBool32
}
type VkCopyDescriptorSet struct {
	SType           VkStructureType
	PNext           uintptr
	SrcSet          VkDescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          VkDescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}
type VkDescriptorBufferInfo struct {
	Buffer VkBuffer
	Offset VkDeviceSize
	Range  VkDeviceSize
}
type VkDescriptorImageInfo struct {
	Sampler     VkSampler
	ImageView   VkImageView
	ImageLayout VkImageLayout
}
type VkDescriptorPoolSize struct {
	Type            VkDescriptorType
	DescriptorCount uint32
}
type VkDescriptorPoolCreateInfo struct {
	SType         VkStructureType
	PNext         uintptr
	Flags         VkDescriptorPoolCreateFlags
	MaxSets       uint32
	PoolSizeCount uint32
	PPoolSizes    *VkDescriptorPoolSize
}
type VkDescriptorSetAllocateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	DescriptorPool     VkDescriptorPool
	DescriptorSetCount uint32
	PSetLayouts        *VkDescriptorSetLayout
}
type VkDescriptorSetLayoutBinding struct {
	Binding            uint32
	DescriptorType     VkDescriptorType
	DescriptorCount    uint32
	StageFlags         VkShaderStageFlags
	PImmutableSamplers *VkSampler
}
type VkDescriptorSetLayoutCreateInfo struct {
	SType        VkStructureType
	PNext        uintptr
	Flags        VkDescriptorSetLayoutCreateFlags
	BindingCount uint32
	PBindings    *VkDescriptorSetLayoutBinding
}
type VkWriteDescriptorSet struct {
	SType            VkStructureType
	PNext            uintptr
	DstSet           VkDescriptorSet
	DstBinding       uint32
	DstArrayElement  uint32
	DescriptorCount  uint32
	DescriptorType   VkDescriptorType
	PImageInfo       *VkDescriptorImageInfo
	PBufferInfo      *VkDescriptorBufferInfo
	PTexelBufferView *VkBufferView
}
type VkAttachmentDescription struct {
	Flags          VkAttachmentDescriptionFlags
	Format         VkFormat
	Samples        VkSampleCountFlagBits
	LoadOp         VkAttachmentLoadOp
	StoreOp        VkAttachmentStoreOp
	StencilLoadOp  VkAttachmentLoadOp
	StencilStoreOp VkAttachmentStoreOp
	InitialLayout  VkImageLayout
	FinalLayout    VkImageLayout
}
type VkAttachmentReference struct {
	Attachment uint32
	Layout     VkImageLayout
}
type VkFramebufferCreateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkFramebufferCreateFlags
	RenderPass      VkRenderPass
	AttachmentCount uint32
	PAttachments    *VkImageView
	Width           uint32
	Height          uint32
	Layers          uint32
}
type VkSubpassDescription struct {
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	InputAttachmentCount    uint32
	PInputAttachments       *VkAttachmentReference
	ColorAttachmentCount    uint32
	PColorAttachments       *VkAttachmentReference
	PResolveAttachments     *VkAttachmentReference
	PDepthStencilAttachment *VkAttachmentReference
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}
type VkSubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    VkPipelineStageFlags
	DstStageMask    VkPipelineStageFlags
	SrcAccessMask   VkAccessFlags
	DstAccessMask   VkAccessFlags
	DependencyFlags VkDependencyFlags
}
type VkRenderPassCreateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkRenderPassCreateFlags
	AttachmentCount uint32
	PAttachments    *VkAttachmentDescription
	SubpassCount    uint32
	PSubpasses      *VkSubpassDescription
	DependencyCount uint32
	PDependencies   *VkSubpassDependency
}
type VkCommandPoolCreateInfo struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkCommandPoolCreateFlags
	QueueFamilyIndex uint32
}
type VkCommandBufferAllocateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	CommandPool        VkCommandPool
	Level              VkCommandBufferLevel
	CommandBufferCount uint32
}
type VkCommandBufferInheritanceInfo struct {
	SType                VkStructureType
	PNext                uintptr
	RenderPass           VkRenderPass
	Subpass              uint32
	Framebuffer          VkFramebuffer
	OcclusionQueryEnable VkBool32
	QueryFlags           VkQueryControlFlags
	PipelineStatistics   VkQueryPipelineStatisticFlags
}
type VkCommandBufferBeginInfo struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkCommandBufferUsageFlags
	PInheritanceInfo *VkCommandBufferInheritanceInfo
}
type VkBufferCopy struct {
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}
type VkImageSubresourceLayers struct {
	AspectMask     VkImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}
type VkBufferImageCopy struct {
	BufferOffset      VkDeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}
type VkClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}
type VkClearAttachment struct {
	AspectMask      VkImageAspectFlags
	ColorAttachment uint32
	ClearValue      VkClearValue
}
type VkClearRect struct {
	Rect           VkRect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}
type VkImageBlit struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     [2]VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     [2]VkOffset3D
}
type VkImageCopy struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}
type VkImageResolve struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}
type VkRenderPassBeginInfo struct {
	SType           VkStructureType
	PNext           uintptr
	RenderPass      VkRenderPass
	Framebuffer     VkFramebuffer
	RenderArea      VkRect2D
	ClearValueCount uint32
	PClearValues    *VkClearValue
}
type VkPhysicalDeviceSubgroupProperties struct {
	SType                     VkStructureType
	PNext                     uintptr
	SubgroupSize              uint32
	SupportedStages           VkShaderStageFlags
	SupportedOperations       VkSubgroupFeatureFlags
	QuadOperationsInAllStages VkBool32
}
type VkBindBufferMemoryInfo struct {
	SType        VkStructureType
	PNext        uintptr
	Buffer       VkBuffer
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}
type VkBindImageMemoryInfo struct {
	SType        VkStructureType
	PNext        uintptr
	Image        VkImage
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}
type VkPhysicalDevice16BitStorageFeatures struct {
	SType                              VkStructureType
	PNext                              uintptr
	StorageBuffer16BitAccess           VkBool32
	UniformAndStorageBuffer16BitAccess VkBool32
	StoragePushConstant16              VkBool32
	StorageInputOutput16               VkBool32
}
type VkMemoryDedicatedRequirements struct {
	SType                       VkStructureType
	PNext                       uintptr
	PrefersDedicatedAllocation  VkBool32
	RequiresDedicatedAllocation VkBool32
}
type VkMemoryDedicatedAllocateInfo struct {
	SType  VkStructureType
	PNext  uintptr
	Image  VkImage
	Buffer VkBuffer
}
type VkMemoryAllocateFlagsInfo struct {
	SType      VkStructureType
	PNext      uintptr
	Flags      VkMemoryAllocateFlags
	DeviceMask uint32
}
type VkDeviceGroupRenderPassBeginInfo struct {
	SType                 VkStructureType
	PNext                 uintptr
	DeviceMask            uint32
	DeviceRenderAreaCount uint32
	PDeviceRenderAreas    *VkRect2D
}
type VkDeviceGroupCommandBufferBeginInfo struct {
	SType      VkStructureType
	PNext      uintptr
	DeviceMask uint32
}
type VkDeviceGroupSubmitInfo struct {
	SType                         VkStructureType
	PNext                         uintptr
	WaitSemaphoreCount            uint32
	PWaitSemaphoreDeviceIndices   *uint32
	CommandBufferCount            uint32
	PCommandBufferDeviceMasks     *uint32
	SignalSemaphoreCount          uint32
	PSignalSemaphoreDeviceIndices *uint32
}
type VkDeviceGroupBindSparseInfo struct {
	SType               VkStructureType
	PNext               uintptr
	ResourceDeviceIndex uint32
	MemoryDeviceIndex   uint32
}
type VkBindBufferMemoryDeviceGroupInfo struct {
	SType            VkStructureType
	PNext            uintptr
	DeviceIndexCount uint32
	PDeviceIndices   *uint32
}
type VkBindImageMemoryDeviceGroupInfo struct {
	SType                        VkStructureType
	PNext                        uintptr
	DeviceIndexCount             uint32
	PDeviceIndices               *uint32
	SplitInstanceBindRegionCount uint32
	PSplitInstanceBindRegions    *VkRect2D
}
type VkPhysicalDeviceGroupProperties struct {
	SType               VkStructureType
	PNext               uintptr
	PhysicalDeviceCount uint32
	PhysicalDevices     [32]VkPhysicalDevice
	SubsetAllocation    VkBool32
}
type VkDeviceGroupDeviceCreateInfo struct {
	SType               VkStructureType
	PNext               uintptr
	PhysicalDeviceCount uint32
	PPhysicalDevices    *VkPhysicalDevice
}
type VkBufferMemoryRequirementsInfo2 struct {
	SType  VkStructureType
	PNext  uintptr
	Buffer VkBuffer
}
type VkImageMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext uintptr
	Image VkImage
}
type VkImageSparseMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext uintptr
	Image VkImage
}
type VkMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              uintptr
	MemoryRequirements VkMemoryRequirements
}
type VkSparseImageMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              uintptr
	MemoryRequirements VkSparseImageMemoryRequirements
}
type VkPhysicalDeviceFeatures2 struct {
	SType    VkStructureType
	PNext    uintptr
	Features VkPhysicalDeviceFeatures
}
type VkPhysicalDeviceProperties2 struct {
	SType      VkStructureType
	PNext      uintptr
	Properties VkPhysicalDeviceProperties
}
type VkFormatProperties2 struct {
	SType            VkStructureType
	PNext            uintptr
	FormatProperties VkFormatProperties
}
type VkImageFormatProperties2 struct {
	SType                 VkStructureType
	PNext                 uintptr
	ImageFormatProperties VkImageFormatProperties
}
type VkPhysicalDeviceImageFormatInfo2 struct {
	SType  VkStructureType
	PNext  uintptr
	Format VkFormat
	Type   VkImageType
	Tiling VkImageTiling
	Usage  VkImageUsageFlags
	Flags  VkImageCreateFlags
}
type VkQueueFamilyProperties2 struct {
	SType                 VkStructureType
	PNext                 uintptr
	QueueFamilyProperties VkQueueFamilyProperties
}
type VkPhysicalDeviceMemoryProperties2 struct {
	SType            VkStructureType
	PNext            uintptr
	MemoryProperties VkPhysicalDeviceMemoryProperties
}
type VkSparseImageFormatProperties2 struct {
	SType      VkStructureType
	PNext      uintptr
	Properties VkSparseImageFormatProperties
}
type VkPhysicalDeviceSparseImageFormatInfo2 struct {
	SType   VkStructureType
	PNext   uintptr
	Format  VkFormat
	Type    VkImageType
	Samples VkSampleCountFlagBits
	Usage   VkImageUsageFlags
	Tiling  VkImageTiling
}
type VkPhysicalDevicePointClippingProperties struct {
	SType                 VkStructureType
	PNext                 uintptr
	PointClippingBehavior VkPointClippingBehavior
}
type VkInputAttachmentAspectReference struct {
	Subpass              uint32
	InputAttachmentIndex uint32
	AspectMask           VkImageAspectFlags
}
type VkRenderPassInputAttachmentAspectCreateInfo struct {
	SType                VkStructureType
	PNext                uintptr
	AspectReferenceCount uint32
	PAspectReferences    *VkInputAttachmentAspectReference
}
type VkImageViewUsageCreateInfo struct {
	SType VkStructureType
	PNext uintptr
	Usage VkImageUsageFlags
}
type VkPipelineTessellationDomainOriginStateCreateInfo struct {
	SType        VkStructureType
	PNext        uintptr
	DomainOrigin VkTessellationDomainOrigin
}
type VkRenderPassMultiviewCreateInfo struct {
	SType                VkStructureType
	PNext                uintptr
	SubpassCount         uint32
	PViewMasks           *uint32
	DependencyCount      uint32
	PViewOffsets         *int32
	CorrelationMaskCount uint32
	PCorrelationMasks    *uint32
}
type VkPhysicalDeviceMultiviewFeatures struct {
	SType                       VkStructureType
	PNext                       uintptr
	Multiview                   VkBool32
	MultiviewGeometryShader     VkBool32
	MultiviewTessellationShader VkBool32
}
type VkPhysicalDeviceMultiviewProperties struct {
	SType                     VkStructureType
	PNext                     uintptr
	MaxMultiviewViewCount     uint32
	MaxMultiviewInstanceIndex uint32
}
type VkPhysicalDeviceVariablePointersFeatures struct {
	SType                         VkStructureType
	PNext                         uintptr
	VariablePointersStorageBuffer VkBool32
	VariablePointers              VkBool32
}
type VkPhysicalDeviceProtectedMemoryFeatures struct {
	SType           VkStructureType
	PNext           uintptr
	ProtectedMemory VkBool32
}
type VkPhysicalDeviceProtectedMemoryProperties struct {
	SType            VkStructureType
	PNext            uintptr
	ProtectedNoFault VkBool32
}
type VkDeviceQueueInfo2 struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueIndex       uint32
}
type VkProtectedSubmitInfo struct {
	SType           VkStructureType
	PNext           uintptr
	ProtectedSubmit VkBool32
}
type VkSamplerYcbcrConversionCreateInfo struct {
	SType                  VkStructureType
	PNext                  uintptr
	Format                 VkFormat
	YcbcrModel             VkSamplerYcbcrModelConversion
	YcbcrRange             VkSamplerYcbcrRange
	Components             VkComponentMapping
	XChromaOffset          VkChromaLocation
	YChromaOffset          VkChromaLocation
	ChromaFilter           VkFilter
	ForceExplicitReruction VkBool32
}
type VkSamplerYcbcrConversionInfo struct {
	SType      VkStructureType
	PNext      uintptr
	Conversion VkSamplerYcbcrConversion
}
type VkBindImagePlaneMemoryInfo struct {
	SType       VkStructureType
	PNext       uintptr
	PlaneAspect VkImageAspectFlagBits
}
type VkImagePlaneMemoryRequirementsInfo struct {
	SType       VkStructureType
	PNext       uintptr
	PlaneAspect VkImageAspectFlagBits
}
type VkPhysicalDeviceSamplerYcbcrConversionFeatures struct {
	SType                  VkStructureType
	PNext                  uintptr
	SamplerYcbcrConversion VkBool32
}
type VkSamplerYcbcrConversionImageFormatProperties struct {
	SType                               VkStructureType
	PNext                               uintptr
	CombinedImageSamplerDescriptorCount uint32
}
type VkDescriptorUpdateTemplateEntry struct {
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  VkDescriptorType
	Offset          uint64
	Stride          uint64
}
type VkDescriptorUpdateTemplateCreateInfo struct {
	SType                      VkStructureType
	PNext                      uintptr
	Flags                      VkDescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntryCount uint32
	PDescriptorUpdateEntries   *VkDescriptorUpdateTemplateEntry
	TemplateType               VkDescriptorUpdateTemplateType
	DescriptorSetLayout        VkDescriptorSetLayout
	PipelineBindPoint          VkPipelineBindPoint
	PipelineLayout             VkPipelineLayout
	SET                        uint32
}
type VkExternalMemoryProperties struct {
	ExternalMemoryFeatures        VkExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes VkExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         VkExternalMemoryHandleTypeFlags
}
type VkPhysicalDeviceExternalImageFormatInfo struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalMemoryHandleTypeFlagBits
}
type VkExternalImageFormatProperties struct {
	SType                    VkStructureType
	PNext                    uintptr
	ExternalMemoryProperties VkExternalMemoryProperties
}
type VkPhysicalDeviceExternalBufferInfo struct {
	SType      VkStructureType
	PNext      uintptr
	Flags      VkBufferCreateFlags
	Usage      VkBufferUsageFlags
	HandleType VkExternalMemoryHandleTypeFlagBits
}
type VkExternalBufferProperties struct {
	SType                    VkStructureType
	PNext                    uintptr
	ExternalMemoryProperties VkExternalMemoryProperties
}
type VkPhysicalDeviceIDProperties struct {
	SType           VkStructureType
	PNext           uintptr
	DeviceUUID      [16]uint8
	DriverUUID      [16]uint8
	DeviceLUID      [8]uint8
	DeviceNodeMask  uint32
	DeviceLUIDValid VkBool32
}
type VkExternalMemoryImageCreateInfo struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalMemoryHandleTypeFlags
}
type VkExternalMemoryBufferCreateInfo struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalMemoryHandleTypeFlags
}
type VkExportMemoryAllocateInfo struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalMemoryHandleTypeFlags
}
type VkPhysicalDeviceExternalFenceInfo struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalFenceHandleTypeFlagBits
}
type VkExternalFenceProperties struct {
	SType                         VkStructureType
	PNext                         uintptr
	ExportFromImportedHandleTypes VkExternalFenceHandleTypeFlags
	CompatibleHandleTypes         VkExternalFenceHandleTypeFlags
	ExternalFenceFeatures         VkExternalFenceFeatureFlags
}
type VkExportFenceCreateInfo struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalFenceHandleTypeFlags
}
type VkExportSemaphoreCreateInfo struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalSemaphoreHandleTypeFlags
}
type VkPhysicalDeviceExternalSemaphoreInfo struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkExternalSemaphoreProperties struct {
	SType                         VkStructureType
	PNext                         uintptr
	ExportFromImportedHandleTypes VkExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         VkExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     VkExternalSemaphoreFeatureFlags
}
type VkPhysicalDeviceMaintenance3Properties struct {
	SType                   VkStructureType
	PNext                   uintptr
	MaxPerSetDescriptors    uint32
	MaxMemoryAllocationSize VkDeviceSize
}
type VkDescriptorSetLayoutSupport struct {
	SType     VkStructureType
	PNext     uintptr
	Supported VkBool32
}
type VkPhysicalDeviceShaderDrawParametersFeatures struct {
	SType                VkStructureType
	PNext                uintptr
	ShaderDrawParameters VkBool32
}
type VkPhysicalDeviceVulkan11Features struct {
	SType                              VkStructureType
	PNext                              uintptr
	StorageBuffer16BitAccess           VkBool32
	UniformAndStorageBuffer16BitAccess VkBool32
	StoragePushConstant16              VkBool32
	StorageInputOutput16               VkBool32
	Multiview                          VkBool32
	MultiviewGeometryShader            VkBool32
	MultiviewTessellationShader        VkBool32
	VariablePointersStorageBuffer      VkBool32
	VariablePointers                   VkBool32
	ProtectedMemory                    VkBool32
	SamplerYcbcrConversion             VkBool32
	ShaderDrawParameters               VkBool32
}
type VkPhysicalDeviceVulkan11Properties struct {
	SType                             VkStructureType
	PNext                             uintptr
	DeviceUUID                        [16]uint8
	DriverUUID                        [16]uint8
	DeviceLUID                        [8]uint8
	DeviceNodeMask                    uint32
	DeviceLUIDValid                   VkBool32
	SubgroupSize                      uint32
	SubgroupSupportedStages           VkShaderStageFlags
	SubgroupSupportedOperations       VkSubgroupFeatureFlags
	SubgroupQuadOperationsInAllStages VkBool32
	PointClippingBehavior             VkPointClippingBehavior
	MaxMultiviewViewCount             uint32
	MaxMultiviewInstanceIndex         uint32
	ProtectedNoFault                  VkBool32
	MaxPerSetDescriptors              uint32
	MaxMemoryAllocationSize           VkDeviceSize
}
type VkPhysicalDeviceVulkan12Features struct {
	SType                                              VkStructureType
	PNext                                              uintptr
	SamplerMirrorClampToEdge                           VkBool32
	DrawIndirectCount                                  VkBool32
	StorageBuffer8BitAccess                            VkBool32
	UniformAndStorageBuffer8BitAccess                  VkBool32
	StoragePushConstant8                               VkBool32
	ShaderBufferInt64Atomics                           VkBool32
	ShaderSharedInt64Atomics                           VkBool32
	ShaderFloat16                                      VkBool32
	ShaderInt8                                         VkBool32
	DescriptorIndexing                                 VkBool32
	ShaderInputAttachmentArrayDynamicIndexing          VkBool32
	ShaderUniformTexelBufferArrayDynamicIndexing       VkBool32
	ShaderStorageTexelBufferArrayDynamicIndexing       VkBool32
	ShaderUniformBufferArrayNonUniformIndexing         VkBool32
	ShaderSampledImageArrayNonUniformIndexing          VkBool32
	ShaderStorageBufferArrayNonUniformIndexing         VkBool32
	ShaderStorageImageArrayNonUniformIndexing          VkBool32
	ShaderInputAttachmentArrayNonUniformIndexing       VkBool32
	ShaderUniformTexelBufferArrayNonUniformIndexing    VkBool32
	ShaderStorageTexelBufferArrayNonUniformIndexing    VkBool32
	DescriptorBindingUniformBufferUpdateAfterBind      VkBool32
	DescriptorBindingSampledImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageBufferUpdateAfterBind      VkBool32
	DescriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingUpdateUnusedWhilePending          VkBool32
	DescriptorBindingPartiallyBound                    VkBool32
	DescriptorBindingVariableDescriptorCount           VkBool32
	RuntimeDescriptorArray                             VkBool32
	SamplerFilterMinmax                                VkBool32
	ScalarBlockLayout                                  VkBool32
	ImagelessFramebuffer                               VkBool32
	UniformBufferStandardLayout                        VkBool32
	ShaderSubgroupExtendedTypes                        VkBool32
	SeparateDepthStencilLayouts                        VkBool32
	HostQueryReset                                     VkBool32
	TimelineSemaphore                                  VkBool32
	BufferDeviceAddress                                VkBool32
	BufferDeviceAddressCaptureReplay                   VkBool32
	BufferDeviceAddressMultiDevice                     VkBool32
	VulkanMemoryModel                                  VkBool32
	VulkanMemoryModelDeviceScope                       VkBool32
	VulkanMemoryModelAvailabilityVisibilityChains      VkBool32
	ShaderOutputViewportIndex                          VkBool32
	ShaderOutputLayer                                  VkBool32
	SubgroupBroadcastDynamicId                         VkBool32
}
type VkConformanceVersion struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}
type VkPhysicalDeviceVulkan12Properties struct {
	SType                                                VkStructureType
	PNext                                                uintptr
	DriverID                                             VkDriverId
	DriverName                                           [256]byte
	DriverInfo                                           [256]byte
	ConformanceVersion                                   VkConformanceVersion
	DenormBehaviorIndependence                           VkShaderFloatControlsIndependence
	RoundingModeIndependence                             VkShaderFloatControlsIndependence
	ShaderSignedZeroInfNanPreserveFloat16                VkBool32
	ShaderSignedZeroInfNanPreserveFloat32                VkBool32
	ShaderSignedZeroInfNanPreserveFloat64                VkBool32
	ShaderDenormPreserveFloat16                          VkBool32
	ShaderDenormPreserveFloat32                          VkBool32
	ShaderDenormPreserveFloat64                          VkBool32
	ShaderDenormFlushToZeroFloat16                       VkBool32
	ShaderDenormFlushToZeroFloat32                       VkBool32
	ShaderDenormFlushToZeroFloat64                       VkBool32
	ShaderRoundingModeRTEFloat16                         VkBool32
	ShaderRoundingModeRTEFloat32                         VkBool32
	ShaderRoundingModeRTEFloat64                         VkBool32
	ShaderRoundingModeRTZFloat16                         VkBool32
	ShaderRoundingModeRTZFloat32                         VkBool32
	ShaderRoundingModeRTZFloat64                         VkBool32
	MaxUpdateAfterBindDescriptorsInAllPools              uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     VkBool32
	ShaderSampledImageArrayNonUniformIndexingNative      VkBool32
	ShaderStorageBufferArrayNonUniformIndexingNative     VkBool32
	ShaderStorageImageArrayNonUniformIndexingNative      VkBool32
	ShaderInputAttachmentArrayNonUniformIndexingNative   VkBool32
	RobustBufferAccessUpdateAfterBind                    VkBool32
	QuadDivergentImplicitLod                             VkBool32
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageUpdateAfterBindResources                  uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
	SupportedDepthResolveModes                           VkResolveModeFlags
	SupportedStencilResolveModes                         VkResolveModeFlags
	IndependentResolveNone                               VkBool32
	IndependentResolve                                   VkBool32
	FilterMinmaxSingleComponentFormats                   VkBool32
	FilterMinmaxImageComponentMapping                    VkBool32
	MaxTimelineSemaphoreValueDifference                  uint64
	FramebufferIntegerColorSampleCounts                  VkSampleCountFlags
}
type VkImageFormatListCreateInfo struct {
	SType           VkStructureType
	PNext           uintptr
	ViewFormatCount uint32
	PViewFormats    *VkFormat
}
type VkAttachmentDescription2 struct {
	SType          VkStructureType
	PNext          uintptr
	Flags          VkAttachmentDescriptionFlags
	Format         VkFormat
	Samples        VkSampleCountFlagBits
	LoadOp         VkAttachmentLoadOp
	StoreOp        VkAttachmentStoreOp
	StencilLoadOp  VkAttachmentLoadOp
	StencilStoreOp VkAttachmentStoreOp
	InitialLayout  VkImageLayout
	FinalLayout    VkImageLayout
}
type VkAttachmentReference2 struct {
	SType      VkStructureType
	PNext      uintptr
	Attachment uint32
	Layout     VkImageLayout
	AspectMask VkImageAspectFlags
}
type VkSubpassDescription2 struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	ViewMask                uint32
	InputAttachmentCount    uint32
	PInputAttachments       *VkAttachmentReference2
	ColorAttachmentCount    uint32
	PColorAttachments       *VkAttachmentReference2
	PResolveAttachments     *VkAttachmentReference2
	PDepthStencilAttachment *VkAttachmentReference2
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}
type VkSubpassDependency2 struct {
	SType           VkStructureType
	PNext           uintptr
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    VkPipelineStageFlags
	DstStageMask    VkPipelineStageFlags
	SrcAccessMask   VkAccessFlags
	DstAccessMask   VkAccessFlags
	DependencyFlags VkDependencyFlags
	ViewOffset      int32
}
type VkRenderPassCreateInfo2 struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkRenderPassCreateFlags
	AttachmentCount         uint32
	PAttachments            *VkAttachmentDescription2
	SubpassCount            uint32
	PSubpasses              *VkSubpassDescription2
	DependencyCount         uint32
	PDependencies           *VkSubpassDependency2
	CorrelatedViewMaskCount uint32
	PCorrelatedViewMasks    *uint32
}
type VkSubpassBeginInfo struct {
	SType    VkStructureType
	PNext    uintptr
	Contents VkSubpassContents
}
type VkSubpassEndInfo struct {
	SType VkStructureType
	PNext uintptr
}
type VkPhysicalDevice8BitStorageFeatures struct {
	SType                             VkStructureType
	PNext                             uintptr
	StorageBuffer8BitAccess           VkBool32
	UniformAndStorageBuffer8BitAccess VkBool32
	StoragePushConstant8              VkBool32
}
type VkPhysicalDeviceDriverProperties struct {
	SType              VkStructureType
	PNext              uintptr
	DriverID           VkDriverId
	DriverName         [256]byte
	DriverInfo         [256]byte
	ConformanceVersion VkConformanceVersion
}
type VkPhysicalDeviceShaderAtomicInt64Features struct {
	SType                    VkStructureType
	PNext                    uintptr
	ShaderBufferInt64Atomics VkBool32
	ShaderSharedInt64Atomics VkBool32
}
type VkPhysicalDeviceShaderFloat16Int8Features struct {
	SType         VkStructureType
	PNext         uintptr
	ShaderFloat16 VkBool32
	ShaderInt8    VkBool32
}
type VkPhysicalDeviceFloatControlsProperties struct {
	SType                                 VkStructureType
	PNext                                 uintptr
	DenormBehaviorIndependence            VkShaderFloatControlsIndependence
	RoundingModeIndependence              VkShaderFloatControlsIndependence
	ShaderSignedZeroInfNanPreserveFloat16 VkBool32
	ShaderSignedZeroInfNanPreserveFloat32 VkBool32
	ShaderSignedZeroInfNanPreserveFloat64 VkBool32
	ShaderDenormPreserveFloat16           VkBool32
	ShaderDenormPreserveFloat32           VkBool32
	ShaderDenormPreserveFloat64           VkBool32
	ShaderDenormFlushToZeroFloat16        VkBool32
	ShaderDenormFlushToZeroFloat32        VkBool32
	ShaderDenormFlushToZeroFloat64        VkBool32
	ShaderRoundingModeRTEFloat16          VkBool32
	ShaderRoundingModeRTEFloat32          VkBool32
	ShaderRoundingModeRTEFloat64          VkBool32
	ShaderRoundingModeRTZFloat16          VkBool32
	ShaderRoundingModeRTZFloat32          VkBool32
	ShaderRoundingModeRTZFloat64          VkBool32
}
type VkDescriptorSetLayoutBindingFlagsCreateInfo struct {
	SType         VkStructureType
	PNext         uintptr
	BindingCount  uint32
	PBindingFlags *VkDescriptorBindingFlags
}
type VkPhysicalDeviceDescriptorIndexingFeatures struct {
	SType                                              VkStructureType
	PNext                                              uintptr
	ShaderInputAttachmentArrayDynamicIndexing          VkBool32
	ShaderUniformTexelBufferArrayDynamicIndexing       VkBool32
	ShaderStorageTexelBufferArrayDynamicIndexing       VkBool32
	ShaderUniformBufferArrayNonUniformIndexing         VkBool32
	ShaderSampledImageArrayNonUniformIndexing          VkBool32
	ShaderStorageBufferArrayNonUniformIndexing         VkBool32
	ShaderStorageImageArrayNonUniformIndexing          VkBool32
	ShaderInputAttachmentArrayNonUniformIndexing       VkBool32
	ShaderUniformTexelBufferArrayNonUniformIndexing    VkBool32
	ShaderStorageTexelBufferArrayNonUniformIndexing    VkBool32
	DescriptorBindingUniformBufferUpdateAfterBind      VkBool32
	DescriptorBindingSampledImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageBufferUpdateAfterBind      VkBool32
	DescriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingUpdateUnusedWhilePending          VkBool32
	DescriptorBindingPartiallyBound                    VkBool32
	DescriptorBindingVariableDescriptorCount           VkBool32
	RuntimeDescriptorArray                             VkBool32
}
type VkPhysicalDeviceDescriptorIndexingProperties struct {
	SType                                                VkStructureType
	PNext                                                uintptr
	MaxUpdateAfterBindDescriptorsInAllPools              uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     VkBool32
	ShaderSampledImageArrayNonUniformIndexingNative      VkBool32
	ShaderStorageBufferArrayNonUniformIndexingNative     VkBool32
	ShaderStorageImageArrayNonUniformIndexingNative      VkBool32
	ShaderInputAttachmentArrayNonUniformIndexingNative   VkBool32
	RobustBufferAccessUpdateAfterBind                    VkBool32
	QuadDivergentImplicitLod                             VkBool32
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageUpdateAfterBindResources                  uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
}
type VkDescriptorSetVariableDescriptorCountAllocateInfo struct {
	SType              VkStructureType
	PNext              uintptr
	DescriptorSetCount uint32
	PDescriptorCounts  *uint32
}
type VkDescriptorSetVariableDescriptorCountLayoutSupport struct {
	SType                      VkStructureType
	PNext                      uintptr
	MaxVariableDescriptorCount uint32
}
type VkSubpassDescriptionDepthStencilResolve struct {
	SType                          VkStructureType
	PNext                          uintptr
	DepthResolveMode               VkResolveModeFlagBits
	StencilResolveMode             VkResolveModeFlagBits
	PDepthStencilResolveAttachment *VkAttachmentReference2
}
type VkPhysicalDeviceDepthStencilResolveProperties struct {
	SType                        VkStructureType
	PNext                        uintptr
	SupportedDepthResolveModes   VkResolveModeFlags
	SupportedStencilResolveModes VkResolveModeFlags
	IndependentResolveNone       VkBool32
	IndependentResolve           VkBool32
}
type VkPhysicalDeviceScalarBlockLayoutFeatures struct {
	SType             VkStructureType
	PNext             uintptr
	ScalarBlockLayout VkBool32
}
type VkImageStencilUsageCreateInfo struct {
	SType        VkStructureType
	PNext        uintptr
	StencilUsage VkImageUsageFlags
}
type VkSamplerReductionModeCreateInfo struct {
	SType         VkStructureType
	PNext         uintptr
	ReductionMode VkSamplerReductionMode
}
type VkPhysicalDeviceSamplerFilterMinmaxProperties struct {
	SType                              VkStructureType
	PNext                              uintptr
	FilterMinmaxSingleComponentFormats VkBool32
	FilterMinmaxImageComponentMapping  VkBool32
}
type VkPhysicalDeviceVulkanMemoryModelFeatures struct {
	SType                                         VkStructureType
	PNext                                         uintptr
	VulkanMemoryModel                             VkBool32
	VulkanMemoryModelDeviceScope                  VkBool32
	VulkanMemoryModelAvailabilityVisibilityChains VkBool32
}
type VkPhysicalDeviceImagelessFramebufferFeatures struct {
	SType                VkStructureType
	PNext                uintptr
	ImagelessFramebuffer VkBool32
}
type VkFramebufferAttachmentImageInfo struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkImageCreateFlags
	Usage           VkImageUsageFlags
	Width           uint32
	Height          uint32
	LayerCount      uint32
	ViewFormatCount uint32
	PViewFormats    *VkFormat
}
type VkFramebufferAttachmentsCreateInfo struct {
	SType                    VkStructureType
	PNext                    uintptr
	AttachmentImageInfoCount uint32
	PAttachmentImageInfos    *VkFramebufferAttachmentImageInfo
}
type VkRenderPassAttachmentBeginInfo struct {
	SType           VkStructureType
	PNext           uintptr
	AttachmentCount uint32
	PAttachments    *VkImageView
}
type VkPhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	SType                       VkStructureType
	PNext                       uintptr
	UniformBufferStandardLayout VkBool32
}
type VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	SType                       VkStructureType
	PNext                       uintptr
	ShaderSubgroupExtendedTypes VkBool32
}
type VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	SType                       VkStructureType
	PNext                       uintptr
	SeparateDepthStencilLayouts VkBool32
}
type VkAttachmentReferenceStencilLayout struct {
	SType         VkStructureType
	PNext         uintptr
	StencilLayout VkImageLayout
}
type VkAttachmentDescriptionStencilLayout struct {
	SType                VkStructureType
	PNext                uintptr
	StencilInitialLayout VkImageLayout
	StencilFinalLayout   VkImageLayout
}
type VkPhysicalDeviceHostQueryResetFeatures struct {
	SType          VkStructureType
	PNext          uintptr
	HostQueryReset VkBool32
}
type VkPhysicalDeviceTimelineSemaphoreFeatures struct {
	SType             VkStructureType
	PNext             uintptr
	TimelineSemaphore VkBool32
}
type VkPhysicalDeviceTimelineSemaphoreProperties struct {
	SType                               VkStructureType
	PNext                               uintptr
	MaxTimelineSemaphoreValueDifference uint64
}
type VkSemaphoreTypeCreateInfo struct {
	SType         VkStructureType
	PNext         uintptr
	SemaphoreType VkSemaphoreType
	InitialValue  uint64
}
type VkTimelineSemaphoreSubmitInfo struct {
	SType                     VkStructureType
	PNext                     uintptr
	WaitSemaphoreValueCount   uint32
	PWaitSemaphoreValues      *uint64
	SignalSemaphoreValueCount uint32
	PSignalSemaphoreValues    *uint64
}
type VkSemaphoreWaitInfo struct {
	SType          VkStructureType
	PNext          uintptr
	Flags          VkSemaphoreWaitFlags
	SemaphoreCount uint32
	PSemaphores    *VkSemaphore
	PValues        *uint64
}
type VkSemaphoreSignalInfo struct {
	SType     VkStructureType
	PNext     uintptr
	Semaphore VkSemaphore
	Value     uint64
}
type VkPhysicalDeviceBufferDeviceAddressFeatures struct {
	SType                            VkStructureType
	PNext                            uintptr
	BufferDeviceAddress              VkBool32
	BufferDeviceAddressCaptureReplay VkBool32
	BufferDeviceAddressMultiDevice   VkBool32
}
type VkBufferDeviceAddressInfo struct {
	SType  VkStructureType
	PNext  uintptr
	Buffer VkBuffer
}
type VkBufferOpaqueCaptureAddressCreateInfo struct {
	SType                VkStructureType
	PNext                uintptr
	OpaqueCaptureAddress uint64
}
type VkMemoryOpaqueCaptureAddressAllocateInfo struct {
	SType                VkStructureType
	PNext                uintptr
	OpaqueCaptureAddress uint64
}
type VkDeviceMemoryOpaqueCaptureAddressInfo struct {
	SType  VkStructureType
	PNext  uintptr
	Memory VkDeviceMemory
}
type VkSurfaceCapabilitiesKHR struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           VkExtent2D
	MinImageExtent          VkExtent2D
	MaxImageExtent          VkExtent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     VkSurfaceTransformFlagsKHR
	CurrentTransform        VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha VkCompositeAlphaFlagsKHR
	SupportedUsageFlags     VkImageUsageFlags
}
type VkSurfaceFormatKHR struct {
	Format     VkFormat
	ColorSpace VkColorSpaceKHR
}
type VkSwapchainCreateInfoKHR struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkSwapchainCreateFlagsKHR
	Surface               VkSurfaceKHR
	MinImageCount         uint32
	ImageFormat           VkFormat
	ImageColorSpace       VkColorSpaceKHR
	ImageExtent           VkExtent2D
	ImageArrayLayers      uint32
	ImageUsage            VkImageUsageFlags
	ImageSharingMode      VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	PreTransform          VkSurfaceTransformFlagBitsKHR
	CompositeAlpha        VkCompositeAlphaFlagBitsKHR
	PresentMode           VkPresentModeKHR
	Clipped               VkBool32
	OldSwapchain          VkSwapchainKHR
}
type VkPresentInfoKHR struct {
	SType              VkStructureType
	PNext              uintptr
	WaitSemaphoreCount uint32
	PWaitSemaphores    *VkSemaphore
	SwapchainCount     uint32
	PSwapchains        *VkSwapchainKHR
	PImageIndices      *uint32
	PResults           *VkResult
}
type VkImageSwapchainCreateInfoKHR struct {
	SType     VkStructureType
	PNext     uintptr
	Swapchain VkSwapchainKHR
}
type VkBindImageMemorySwapchainInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Swapchain  VkSwapchainKHR
	ImageIndex uint32
}
type VkAcquireNextImageInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Swapchain  VkSwapchainKHR
	Timeout    uint64
	Semaphore  VkSemaphore
	Fence      VkFence
	DeviceMask uint32
}
type VkDeviceGroupPresentCapabilitiesKHR struct {
	SType       VkStructureType
	PNext       uintptr
	PresentMask [32]uint32
	Modes       VkDeviceGroupPresentModeFlagsKHR
}
type VkDeviceGroupPresentInfoKHR struct {
	SType          VkStructureType
	PNext          uintptr
	SwapchainCount uint32
	PDeviceMasks   *uint32
	Mode           VkDeviceGroupPresentModeFlagBitsKHR
}
type VkDeviceGroupSwapchainCreateInfoKHR struct {
	SType VkStructureType
	PNext uintptr
	Modes VkDeviceGroupPresentModeFlagsKHR
}
type VkDisplayModeParametersKHR struct {
	VisibleRegion VkExtent2D
	RefreshRate   uint32
}
type VkDisplayModeCreateInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Flags      VkDisplayModeCreateFlagsKHR
	Parameters VkDisplayModeParametersKHR
}
type VkDisplayModePropertiesKHR struct {
	DisplayMode VkDisplayModeKHR
	Parameters  VkDisplayModeParametersKHR
}
type VkDisplayPlaneCapabilitiesKHR struct {
	SupportedAlpha VkDisplayPlaneAlphaFlagsKHR
	MinSrcPosition VkOffset2D
	MaxSrcPosition VkOffset2D
	MinSrcExtent   VkExtent2D
	MaxSrcExtent   VkExtent2D
	MinDstPosition VkOffset2D
	MaxDstPosition VkOffset2D
	MinDstExtent   VkExtent2D
	MaxDstExtent   VkExtent2D
}
type VkDisplayPlanePropertiesKHR struct {
	CurrentDisplay    VkDisplayKHR
	CurrentStackIndex uint32
}
type VkDisplayPropertiesKHR struct {
	Display              VkDisplayKHR
	DisplayName          *byte
	PhysicalDimensions   VkExtent2D
	PhysicalResolution   VkExtent2D
	SupportedTransforms  VkSurfaceTransformFlagsKHR
	PlaneReorderPossible VkBool32
	PersistentContent    VkBool32
}
type VkDisplaySurfaceCreateInfoKHR struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkDisplaySurfaceCreateFlagsKHR
	DisplayMode     VkDisplayModeKHR
	PlaneIndex      uint32
	PlaneStackIndex uint32
	Transform       VkSurfaceTransformFlagBitsKHR
	GlobalAlpha     float32
	AlphaMode       VkDisplayPlaneAlphaFlagBitsKHR
	ImageExtent     VkExtent2D
}
type VkDisplayPresentInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	SrcRect    VkRect2D
	DstRect    VkRect2D
	Persistent VkBool32
}
type VkImportMemoryFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	HandleType VkExternalMemoryHandleTypeFlagBits
	FD         int
}
type VkMemoryFdPropertiesKHR struct {
	SType          VkStructureType
	PNext          uintptr
	MemoryTypeBits uint32
}
type VkMemoryGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}
type VkImportSemaphoreFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Semaphore  VkSemaphore
	Flags      VkSemaphoreImportFlags
	HandleType VkExternalSemaphoreHandleTypeFlagBits
	FD         int
}
type VkSemaphoreGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkPhysicalDevicePushDescriptorPropertiesKHR struct {
	SType              VkStructureType
	PNext              uintptr
	MaxPushDescriptors uint32
}
type VkRectLayerKHR struct {
	Offset VkOffset2D
	Extent VkExtent2D
	Layer  uint32
}
type VkPresentRegionKHR struct {
	RectangleCount uint32
	PRectangles    *VkRectLayerKHR
}
type VkPresentRegionsKHR struct {
	SType          VkStructureType
	PNext          uintptr
	SwapchainCount uint32
	PRegions       *VkPresentRegionKHR
}
type VkSharedPresentSurfaceCapabilitiesKHR struct {
	SType                            VkStructureType
	PNext                            uintptr
	SharedPresentSupportedUsageFlags VkImageUsageFlags
}
type VkImportFenceFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Fence      VkFence
	Flags      VkFenceImportFlags
	HandleType VkExternalFenceHandleTypeFlagBits
	FD         int
}
type VkFenceGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      uintptr
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
}
type VkPhysicalDevicePerformanceQueryFeaturesKHR struct {
	SType                                VkStructureType
	PNext                                uintptr
	PerformanceCounterQueryPools         VkBool32
	PerformanceCounterMultipleQueryPools VkBool32
}
type VkPhysicalDevicePerformanceQueryPropertiesKHR struct {
	SType                         VkStructureType
	PNext                         uintptr
	AllowCommandBufferQueryCopies VkBool32
}
type VkPerformanceCounterKHR struct {
	SType   VkStructureType
	PNext   uintptr
	Unit    VkPerformanceCounterUnitKHR
	Scope   VkPerformanceCounterScopeKHR
	Storage VkPerformanceCounterStorageKHR
	Uuid    [16]uint8
}
type VkPerformanceCounterDescriptionKHR struct {
	SType       VkStructureType
	PNext       uintptr
	Flags       VkPerformanceCounterDescriptionFlagsKHR
	Name        [256]byte
	Category    [256]byte
	Description [256]byte
}
type VkQueryPoolPerformanceCreateInfoKHR struct {
	SType             VkStructureType
	PNext             uintptr
	QueueFamilyIndex  uint32
	CounterIndexCount uint32
	PCounterIndices   *uint32
}
type VkAcquireProfilingLockInfoKHR struct {
	SType   VkStructureType
	PNext   uintptr
	Flags   VkAcquireProfilingLockFlagsKHR
	Timeout uint64
}
type VkPerformanceQuerySubmitInfoKHR struct {
	SType            VkStructureType
	PNext            uintptr
	CounterPassIndex uint32
}
type VkPhysicalDeviceSurfaceInfo2KHR struct {
	SType   VkStructureType
	PNext   uintptr
	Surface VkSurfaceKHR
}
type VkSurfaceCapabilities2KHR struct {
	SType               VkStructureType
	PNext               uintptr
	SurfaceCapabilities VkSurfaceCapabilitiesKHR
}
type VkSurfaceFormat2KHR struct {
	SType         VkStructureType
	PNext         uintptr
	SurfaceFormat VkSurfaceFormatKHR
}
type VkDisplayProperties2KHR struct {
	SType             VkStructureType
	PNext             uintptr
	DisplayProperties VkDisplayPropertiesKHR
}
type VkDisplayPlaneProperties2KHR struct {
	SType                  VkStructureType
	PNext                  uintptr
	DisplayPlaneProperties VkDisplayPlanePropertiesKHR
}
type VkDisplayModeProperties2KHR struct {
	SType                 VkStructureType
	PNext                 uintptr
	DisplayModeProperties VkDisplayModePropertiesKHR
}
type VkDisplayPlaneInfo2KHR struct {
	SType      VkStructureType
	PNext      uintptr
	Mode       VkDisplayModeKHR
	PlaneIndex uint32
}
type VkDisplayPlaneCapabilities2KHR struct {
	SType        VkStructureType
	PNext        uintptr
	Capabilities VkDisplayPlaneCapabilitiesKHR
}
type VkPhysicalDeviceShaderClockFeaturesKHR struct {
	SType               VkStructureType
	PNext               uintptr
	ShaderSubgroupClock VkBool32
	ShaderDeviceClock   VkBool32
}
type VkPhysicalDeviceShaderTerminateInvocationFeaturesKHR struct {
	SType                     VkStructureType
	PNext                     uintptr
	ShaderTerminateInvocation VkBool32
}
type VkFragmentShadingRateAttachmentInfoKHR struct {
	SType                          VkStructureType
	PNext                          uintptr
	PFragmentShadingRateAttachment *VkAttachmentReference2
	ShadingRateAttachmentTexelSize VkExtent2D
}
type VkPipelineFragmentShadingRateStateCreateInfoKHR struct {
	SType        VkStructureType
	PNext        uintptr
	FragmentSize VkExtent2D
	CombinerOps  [2]VkFragmentShadingRateCombinerOpKHR
}
type VkPhysicalDeviceFragmentShadingRateFeaturesKHR struct {
	SType                         VkStructureType
	PNext                         uintptr
	PipelineFragmentShadingRate   VkBool32
	PrimitiveFragmentShadingRate  VkBool32
	AttachmentFragmentShadingRate VkBool32
}
type VkPhysicalDeviceFragmentShadingRatePropertiesKHR struct {
	SType                                                VkStructureType
	PNext                                                uintptr
	MinFragmentShadingRateAttachmentTexelSize            VkExtent2D
	MaxFragmentShadingRateAttachmentTexelSize            VkExtent2D
	MaxFragmentShadingRateAttachmentTexelSizeAspectRatio uint32
	PrimitiveFragmentShadingRateWithMultipleViewports    VkBool32
	LayeredShadingRateAttachments                        VkBool32
	FragmentShadingRateNonTrivialCombinerOps             VkBool32
	MaxFragmentSize                                      VkExtent2D
	MaxFragmentSizeAspectRatio                           uint32
	MaxFragmentShadingRateCoverageSamples                uint32
	MaxFragmentShadingRateRasterizationSamples           VkSampleCountFlagBits
	FragmentShadingRateWithShaderDepthStencilWrites      VkBool32
	FragmentShadingRateWithSampleMask                    VkBool32
	FragmentShadingRateWithShaderSampleMask              VkBool32
	FragmentShadingRateWithConservativeRasterization     VkBool32
	FragmentShadingRateWithFragmentShaderInterlock       VkBool32
	FragmentShadingRateWithCustomSampleLocations         VkBool32
	FragmentShadingRateStrictMultiplyCombiner            VkBool32
}
type VkPhysicalDeviceFragmentShadingRateKHR struct {
	SType        VkStructureType
	PNext        uintptr
	SampleCounts VkSampleCountFlags
	FragmentSize VkExtent2D
}
type VkSurfaceProtectedCapabilitiesKHR struct {
	SType             VkStructureType
	PNext             uintptr
	SupportsProtected VkBool32
}
type VkPhysicalDevicePresentWaitFeaturesKHR struct {
	SType       VkStructureType
	PNext       uintptr
	PresentWait VkBool32
}
type VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR struct {
	SType                  VkStructureType
	PNext                  uintptr
	PipelineExecutableInfo VkBool32
}
type VkPipelineInfoKHR struct {
	SType    VkStructureType
	PNext    uintptr
	Pipeline VkPipeline
}
type VkPipelineExecutablePropertiesKHR struct {
	SType        VkStructureType
	PNext        uintptr
	Stages       VkShaderStageFlags
	Name         [256]byte
	Description  [256]byte
	SubgroupSize uint32
}
type VkPipelineExecutableInfoKHR struct {
	SType           VkStructureType
	PNext           uintptr
	Pipeline        VkPipeline
	ExecutableIndex uint32
}
type VkPipelineExecutableStatisticKHR struct {
	SType       VkStructureType
	PNext       uintptr
	Name        [256]byte
	Description [256]byte
	Format      VkPipelineExecutableStatisticFormatKHR
	Value       VkPipelineExecutableStatisticValueKHR
}
type VkPipelineExecutableInternalRepresentationKHR struct {
	SType       VkStructureType
	PNext       uintptr
	Name        [256]byte
	Description [256]byte
	IsText      VkBool32
	DataSize    uint64
	PData       uintptr
}
type VkPipelineLibraryCreateInfoKHR struct {
	SType        VkStructureType
	PNext        uintptr
	LibraryCount uint32
	PLibraries   *VkPipeline
}
type VkPresentIdKHR struct {
	SType          VkStructureType
	PNext          uintptr
	SwapchainCount uint32
	PPresentIds    *uint64
}
type VkPhysicalDevicePresentIdFeaturesKHR struct {
	SType     VkStructureType
	PNext     uintptr
	PresentId VkBool32
}
type VkMemoryBarrier2KHR struct {
	SType         VkStructureType
	PNext         uintptr
	SrcStageMask  VkPipelineStageFlags2KHR
	SrcAccessMask VkAccessFlags2KHR
	DstStageMask  VkPipelineStageFlags2KHR
	DstAccessMask VkAccessFlags2KHR
}
type VkBufferMemoryBarrier2KHR struct {
	SType               VkStructureType
	PNext               uintptr
	SrcStageMask        VkPipelineStageFlags2KHR
	SrcAccessMask       VkAccessFlags2KHR
	DstStageMask        VkPipelineStageFlags2KHR
	DstAccessMask       VkAccessFlags2KHR
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              VkBuffer
	Offset              VkDeviceSize
	Size                VkDeviceSize
}
type VkImageMemoryBarrier2KHR struct {
	SType               VkStructureType
	PNext               uintptr
	SrcStageMask        VkPipelineStageFlags2KHR
	SrcAccessMask       VkAccessFlags2KHR
	DstStageMask        VkPipelineStageFlags2KHR
	DstAccessMask       VkAccessFlags2KHR
	OldLayout           VkImageLayout
	NewLayout           VkImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               VkImage
	SubresourceRange    VkImageSubresourceRange
}
type VkDependencyInfoKHR struct {
	SType                    VkStructureType
	PNext                    uintptr
	DependencyFlags          VkDependencyFlags
	MemoryBarrierCount       uint32
	PMemoryBarriers          *VkMemoryBarrier2KHR
	BufferMemoryBarrierCount uint32
	PBufferMemoryBarriers    *VkBufferMemoryBarrier2KHR
	ImageMemoryBarrierCount  uint32
	PImageMemoryBarriers     *VkImageMemoryBarrier2KHR
}
type VkSemaphoreSubmitInfoKHR struct {
	SType       VkStructureType
	PNext       uintptr
	Semaphore   VkSemaphore
	Value       uint64
	StageMask   VkPipelineStageFlags2KHR
	DeviceIndex uint32
}
type VkCommandBufferSubmitInfoKHR struct {
	SType         VkStructureType
	PNext         uintptr
	CommandBuffer VkCommandBuffer
	DeviceMask    uint32
}
type VkSubmitInfo2KHR struct {
	SType                    VkStructureType
	PNext                    uintptr
	Flags                    VkSubmitFlagsKHR
	WaitSemaphoreInfoCount   uint32
	PWaitSemaphoreInfos      *VkSemaphoreSubmitInfoKHR
	CommandBufferInfoCount   uint32
	PCommandBufferInfos      *VkCommandBufferSubmitInfoKHR
	SignalSemaphoreInfoCount uint32
	PSignalSemaphoreInfos    *VkSemaphoreSubmitInfoKHR
}
type VkPhysicalDeviceSynchronization2FeaturesKHR struct {
	SType            VkStructureType
	PNext            uintptr
	Synchronization2 VkBool32
}
type VkQueueFamilyCheckpointProperties2NV struct {
	SType                        VkStructureType
	PNext                        uintptr
	CheckpointExecutionStageMask VkPipelineStageFlags2KHR
}
type VkCheckpointData2NV struct {
	SType             VkStructureType
	PNext             uintptr
	Stage             VkPipelineStageFlags2KHR
	PCheckpointMarker uintptr
}
type VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR struct {
	SType                            VkStructureType
	PNext                            uintptr
	ShaderSubgroupUniformControlFlow VkBool32
}
type VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR struct {
	SType                               VkStructureType
	PNext                               uintptr
	ShaderZeroInitializeWorkgroupMemory VkBool32
}
type VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR struct {
	SType                                          VkStructureType
	PNext                                          uintptr
	WorkgroupMemoryExplicitLayout                  VkBool32
	WorkgroupMemoryExplicitLayoutScalarBlockLayout VkBool32
	WorkgroupMemoryExplicitLayout8BitAccess        VkBool32
	WorkgroupMemoryExplicitLayout16BitAccess       VkBool32
}
type VkBufferCopy2KHR struct {
	SType     VkStructureType
	PNext     uintptr
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}
type VkCopyBufferInfo2KHR struct {
	SType       VkStructureType
	PNext       uintptr
	SrcBuffer   VkBuffer
	DstBuffer   VkBuffer
	RegionCount uint32
	PRegions    *VkBufferCopy2KHR
}
type VkImageCopy2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}
type VkCopyImageInfo2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageCopy2KHR
}
type VkBufferImageCopy2KHR struct {
	SType             VkStructureType
	PNext             uintptr
	BufferOffset      VkDeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}
type VkCopyBufferToImageInfo2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcBuffer      VkBuffer
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkBufferImageCopy2KHR
}
type VkCopyImageToBufferInfo2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstBuffer      VkBuffer
	RegionCount    uint32
	PRegions       *VkBufferImageCopy2KHR
}
type VkImageBlit2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     [2]VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     [2]VkOffset3D
}
type VkBlitImageInfo2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageBlit2KHR
	Filter         VkFilter
}
type VkImageResolve2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}
type VkResolveImageInfo2KHR struct {
	SType          VkStructureType
	PNext          uintptr
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageResolve2KHR
}
type VkDebugReportCallbackCreateInfoEXT struct {
	SType       VkStructureType
	PNext       uintptr
	Flags       VkDebugReportFlagsEXT
	PfnCallback uintptr
	PUserData   uintptr
}
type VkPipelineRasterizationStateRasterizationOrderAMD struct {
	SType              VkStructureType
	PNext              uintptr
	RasterizationOrder VkRasterizationOrderAMD
}
type VkDebugMarkerObjectNameInfoEXT struct {
	SType       VkStructureType
	PNext       uintptr
	ObjectType  VkDebugReportObjectTypeEXT
	Object      uint64
	PObjectName *byte
}
type VkDebugMarkerObjectTagInfoEXT struct {
	SType      VkStructureType
	PNext      uintptr
	ObjectType VkDebugReportObjectTypeEXT
	Object     uint64
	TagName    uint64
	TagSize    uint64
	PTag       uintptr
}
type VkDebugMarkerMarkerInfoEXT struct {
	SType       VkStructureType
	PNext       uintptr
	PMarkerName *byte
	Color       [4]float32
}
type VkDedicatedAllocationImageCreateInfoNV struct {
	SType               VkStructureType
	PNext               uintptr
	DedicatedAllocation VkBool32
}
type VkDedicatedAllocationBufferCreateInfoNV struct {
	SType               VkStructureType
	PNext               uintptr
	DedicatedAllocation VkBool32
}
type VkDedicatedAllocationMemoryAllocateInfoNV struct {
	SType  VkStructureType
	PNext  uintptr
	Image  VkImage
	Buffer VkBuffer
}
type VkPhysicalDeviceTransformFeedbackFeaturesEXT struct {
	SType             VkStructureType
	PNext             uintptr
	TransformFeedback VkBool32
	GeometryStreams   VkBool32
}
type VkPhysicalDeviceTransformFeedbackPropertiesEXT struct {
	SType                                      VkStructureType
	PNext                                      uintptr
	MaxTransformFeedbackStreams                uint32
	MaxTransformFeedbackBuffers                uint32
	MaxTransformFeedbackBufferSize             VkDeviceSize
	MaxTransformFeedbackStreamDataSize         uint32
	MaxTransformFeedbackBufferDataSize         uint32
	MaxTransformFeedbackBufferDataStride       uint32
	TransformFeedbackQueries                   VkBool32
	TransformFeedbackStreamsLinesTriangles     VkBool32
	TransformFeedbackRasterizationStreamSelect VkBool32
	TransformFeedbackDraw                      VkBool32
}
type VkPipelineRasterizationStateStreamCreateInfoEXT struct {
	SType               VkStructureType
	PNext               uintptr
	Flags               VkPipelineRasterizationStateStreamCreateFlagsEXT
	RasterizationStream uint32
}
type VkCuModuleCreateInfoNVX struct {
	SType    VkStructureType
	PNext    uintptr
	DataSize uint64
	PData    uintptr
}
type VkCuFunctionCreateInfoNVX struct {
	SType  VkStructureType
	PNext  uintptr
	Module VkCuModuleNVX
	PName  *byte
}
type VkCuLaunchInfoNVX struct {
	SType          VkStructureType
	PNext          uintptr
	Function       VkCuFunctionNVX
	GridDimX       uint32
	GridDimY       uint32
	GridDimZ       uint32
	BlockDimX      uint32
	BlockDimY      uint32
	BlockDimZ      uint32
	SharedMemBytes uint32
	ParamCount     uint64
	ExtraCount     uint64
}
type VkImageViewHandleInfoNVX struct {
	SType          VkStructureType
	PNext          uintptr
	ImageView      VkImageView
	DescriptorType VkDescriptorType
	Sampler        VkSampler
}
type VkImageViewAddressPropertiesNVX struct {
	SType         VkStructureType
	PNext         uintptr
	DeviceAddress VkDeviceAddress
	Size          VkDeviceSize
}
type VkTextureLODGatherFormatPropertiesAMD struct {
	SType                           VkStructureType
	PNext                           uintptr
	SupportsTextureGatherLODBiasAMD VkBool32
}
type VkShaderResourceUsageAMD struct {
	NumUsedVgprs             uint32
	NumUsedSgprs             uint32
	LdsSizePerLocalWorkGroup uint32
	LdsUsageSizeInBytes      uint64
	ScratchMemUsageInBytes   uint64
}
type VkShaderStatisticsInfoAMD struct {
	ShaderStageMask      VkShaderStageFlags
	ResourceUsage        VkShaderResourceUsageAMD
	NumPhysicalVgprs     uint32
	NumPhysicalSgprs     uint32
	NumAvailableVgprs    uint32
	NumAvailableSgprs    uint32
	ComputeWorkGroupSize [3]uint32
}
type VkPhysicalDeviceCornerSampledImageFeaturesNV struct {
	SType              VkStructureType
	PNext              uintptr
	CornerSampledImage VkBool32
}
type VkExternalImageFormatPropertiesNV struct {
	ImageFormatProperties         VkImageFormatProperties
	ExternalMemoryFeatures        VkExternalMemoryFeatureFlagsNV
	ExportFromImportedHandleTypes VkExternalMemoryHandleTypeFlagsNV
	CompatibleHandleTypes         VkExternalMemoryHandleTypeFlagsNV
}
type VkExternalMemoryImageCreateInfoNV struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalMemoryHandleTypeFlagsNV
}
type VkExportMemoryAllocateInfoNV struct {
	SType       VkStructureType
	PNext       uintptr
	HandleTypes VkExternalMemoryHandleTypeFlagsNV
}
type VkValidationFlagsEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	DisabledValidationCheckCount uint32
	PDisabledValidationChecks    *VkValidationCheckEXT
}
type VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT struct {
	SType                      VkStructureType
	PNext                      uintptr
	TextureCompressionASTC_HDR VkBool32
}
type VkImageViewASTCDecodeModeEXT struct {
	SType      VkStructureType
	PNext      uintptr
	DecodeMode VkFormat
}
type VkPhysicalDeviceASTCDecodeFeaturesEXT struct {
	SType                    VkStructureType
	PNext                    uintptr
	DecodeModeSharedExponent VkBool32
}
type VkConditionalRenderingBeginInfoEXT struct {
	SType  VkStructureType
	PNext  uintptr
	Buffer VkBuffer
	Offset VkDeviceSize
	Flags  VkConditionalRenderingFlagsEXT
}
type VkPhysicalDeviceConditionalRenderingFeaturesEXT struct {
	SType                         VkStructureType
	PNext                         uintptr
	ConditionalRendering          VkBool32
	InheritedConditionalRendering VkBool32
}
type VkCommandBufferInheritanceConditionalRenderingInfoEXT struct {
	SType                      VkStructureType
	PNext                      uintptr
	ConditionalRenderingEnable VkBool32
}
type VkViewportWScalingNV struct {
	Xcoeff float32
	Ycoeff float32
}
type VkPipelineViewportWScalingStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  uintptr
	ViewportWScalingEnable VkBool32
	ViewportCount          uint32
	PViewportWScalings     *VkViewportWScalingNV
}
type VkSurfaceCapabilities2EXT struct {
	SType                    VkStructureType
	PNext                    uintptr
	MinImageCount            uint32
	MaxImageCount            uint32
	CurrentExtent            VkExtent2D
	MinImageExtent           VkExtent2D
	MaxImageExtent           VkExtent2D
	MaxImageArrayLayers      uint32
	SupportedTransforms      VkSurfaceTransformFlagsKHR
	CurrentTransform         VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha  VkCompositeAlphaFlagsKHR
	SupportedUsageFlags      VkImageUsageFlags
	SupportedSurfaceCounters VkSurfaceCounterFlagsEXT
}
type VkDisplayPowerInfoEXT struct {
	SType      VkStructureType
	PNext      uintptr
	PowerState VkDisplayPowerStateEXT
}
type VkDeviceEventInfoEXT struct {
	SType       VkStructureType
	PNext       uintptr
	DeviceEvent VkDeviceEventTypeEXT
}
type VkDisplayEventInfoEXT struct {
	SType        VkStructureType
	PNext        uintptr
	DisplayEvent VkDisplayEventTypeEXT
}
type VkSwapchainCounterCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	SurfaceCounters VkSurfaceCounterFlagsEXT
}
type VkRefreshCycleDurationGOOGLE struct {
	RefreshDuration uint64
}
type VkPastPresentationTimingGOOGLE struct {
	PresentID           uint32
	DesiredPresentTime  uint64
	ActualPresentTime   uint64
	EarliestPresentTime uint64
	PresentMargin       uint64
}
type VkPresentTimeGOOGLE struct {
	PresentID          uint32
	DesiredPresentTime uint64
}
type VkPresentTimesInfoGOOGLE struct {
	SType          VkStructureType
	PNext          uintptr
	SwapchainCount uint32
	PTimes         *VkPresentTimeGOOGLE
}
type VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX struct {
	SType                        VkStructureType
	PNext                        uintptr
	PerViewPositionAllComponents VkBool32
}
type VkViewportSwizzleNV struct {
	X VkViewportCoordinateSwizzleNV
	Y VkViewportCoordinateSwizzleNV
	Z VkViewportCoordinateSwizzleNV
	W VkViewportCoordinateSwizzleNV
}
type VkPipelineViewportSwizzleStateCreateInfoNV struct {
	SType             VkStructureType
	PNext             uintptr
	Flags             VkPipelineViewportSwizzleStateCreateFlagsNV
	ViewportCount     uint32
	PViewportSwizzles *VkViewportSwizzleNV
}
type VkPhysicalDeviceDiscardRectanglePropertiesEXT struct {
	SType                VkStructureType
	PNext                uintptr
	MaxDiscardRectangles uint32
}
type VkPipelineDiscardRectangleStateCreateInfoEXT struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkPipelineDiscardRectangleStateCreateFlagsEXT
	DiscardRectangleMode  VkDiscardRectangleModeEXT
	DiscardRectangleCount uint32
	PDiscardRectangles    *VkRect2D
}
type VkPhysicalDeviceConservativeRasterizationPropertiesEXT struct {
	SType                                       VkStructureType
	PNext                                       uintptr
	PrimitiveOverestimationSize                 float32
	MaxExtraPrimitiveOverestimationSize         float32
	ExtraPrimitiveOverestimationSizeGranularity float32
	PrimitiveUnderestimation                    VkBool32
	ConservativePointAndLineRasterization       VkBool32
	DegenerateTrianglesRasterized               VkBool32
	DegenerateLinesRasterized                   VkBool32
	FullyCoveredFragmentShaderInputVariable     VkBool32
	ConservativeRasterizationPostDepthCoverage  VkBool32
}
type VkPipelineRasterizationConservativeStateCreateInfoEXT struct {
	SType                            VkStructureType
	PNext                            uintptr
	Flags                            VkPipelineRasterizationConservativeStateCreateFlagsEXT
	ConservativeRasterizationMode    VkConservativeRasterizationModeEXT
	ExtraPrimitiveOverestimationSize float32
}
type VkPhysicalDeviceDepthClipEnableFeaturesEXT struct {
	SType           VkStructureType
	PNext           uintptr
	DepthClipEnable VkBool32
}
type VkPipelineRasterizationDepthClipStateCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkPipelineRasterizationDepthClipStateCreateFlagsEXT
	DepthClipEnable VkBool32
}
type VkXYColorEXT struct {
	X float32
	Y float32
}
type VkHdrMetadataEXT struct {
	SType                     VkStructureType
	PNext                     uintptr
	DisplayPrimaryRed         VkXYColorEXT
	DisplayPrimaryGreen       VkXYColorEXT
	DisplayPrimaryBlue        VkXYColorEXT
	WhitePoint                VkXYColorEXT
	MaxLuminance              float32
	MinLuminance              float32
	MaxContentLightLevel      float32
	MaxFrameAverageLightLevel float32
}
type VkDebugUtilsLabelEXT struct {
	SType      VkStructureType
	PNext      uintptr
	PLabelName *byte
	Color      [4]float32
}
type VkDebugUtilsObjectNameInfoEXT struct {
	SType        VkStructureType
	PNext        uintptr
	ObjectType   VkObjectType
	ObjectHandle uint64
	PObjectName  *byte
}
type VkDebugUtilsMessengerCallbackDataEXT struct {
	SType            VkStructureType
	PNext            uintptr
	Flags            VkDebugUtilsMessengerCallbackDataFlagsEXT
	PMessageIdName   *byte
	MessageIdNumber  int32
	PMessage         *byte
	QueueLabelCount  uint32
	PQueueLabels     *VkDebugUtilsLabelEXT
	CmdBufLabelCount uint32
	PCmdBufLabels    *VkDebugUtilsLabelEXT
	ObjectCount      uint32
	PObjects         *VkDebugUtilsObjectNameInfoEXT
}
type VkDebugUtilsMessengerCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkDebugUtilsMessengerCreateFlagsEXT
	MessageSeverity VkDebugUtilsMessageSeverityFlagsEXT
	MessageType     VkDebugUtilsMessageTypeFlagsEXT
	PfnUserCallback uintptr
	PUserData       uintptr
}
type VkDebugUtilsObjectTagInfoEXT struct {
	SType        VkStructureType
	PNext        uintptr
	ObjectType   VkObjectType
	ObjectHandle uint64
	TagName      uint64
	TagSize      uint64
	PTag         uintptr
}
type VkPhysicalDeviceInlineUniformBlockFeaturesEXT struct {
	SType                                              VkStructureType
	PNext                                              uintptr
	InlineUniformBlock                                 VkBool32
	DescriptorBindingInlineUniformBlockUpdateAfterBind VkBool32
}
type VkPhysicalDeviceInlineUniformBlockPropertiesEXT struct {
	SType                                                   VkStructureType
	PNext                                                   uintptr
	MaxInlineUniformBlockSize                               uint32
	MaxPerStageDescriptorInlineUniformBlocks                uint32
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks uint32
	MaxDescriptorSetInlineUniformBlocks                     uint32
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks      uint32
}
type VkWriteDescriptorSetInlineUniformBlockEXT struct {
	SType    VkStructureType
	PNext    uintptr
	DataSize uint32
	PData    uintptr
}
type VkDescriptorPoolInlineUniformBlockCreateInfoEXT struct {
	SType                         VkStructureType
	PNext                         uintptr
	MaxInlineUniformBlockBindings uint32
}
type VkSampleLocationEXT struct {
	X float32
	Y float32
}
type VkSampleLocationsInfoEXT struct {
	SType                   VkStructureType
	PNext                   uintptr
	SampleLocationsPerPixel VkSampleCountFlagBits
	SampleLocationGridSize  VkExtent2D
	SampleLocationsCount    uint32
	PSampleLocations        *VkSampleLocationEXT
}
type VkAttachmentSampleLocationsEXT struct {
	AttachmentIndex     uint32
	SampleLocationsInfo VkSampleLocationsInfoEXT
}
type VkSubpassSampleLocationsEXT struct {
	SubpassIndex        uint32
	SampleLocationsInfo VkSampleLocationsInfoEXT
}
type VkRenderPassSampleLocationsBeginInfoEXT struct {
	SType                                 VkStructureType
	PNext                                 uintptr
	AttachmentInitialSampleLocationsCount uint32
	PAttachmentInitialSampleLocations     *VkAttachmentSampleLocationsEXT
	PostSubpassSampleLocationsCount       uint32
	PPostSubpassSampleLocations           *VkSubpassSampleLocationsEXT
}
type VkPipelineSampleLocationsStateCreateInfoEXT struct {
	SType                 VkStructureType
	PNext                 uintptr
	SampleLocationsEnable VkBool32
	SampleLocationsInfo   VkSampleLocationsInfoEXT
}
type VkPhysicalDeviceSampleLocationsPropertiesEXT struct {
	SType                         VkStructureType
	PNext                         uintptr
	SampleLocationSampleCounts    VkSampleCountFlags
	MaxSampleLocationGridSize     VkExtent2D
	SampleLocationCoordinateRange [2]float32
	SampleLocationSubPixelBits    uint32
	VariableSampleLocations       VkBool32
}
type VkMultisamplePropertiesEXT struct {
	SType                     VkStructureType
	PNext                     uintptr
	MaxSampleLocationGridSize VkExtent2D
}
type VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT struct {
	SType                           VkStructureType
	PNext                           uintptr
	AdvancedBlendCoherentOperations VkBool32
}
type VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT struct {
	SType                                 VkStructureType
	PNext                                 uintptr
	AdvancedBlendMaxColorAttachments      uint32
	AdvancedBlendIndependentBlend         VkBool32
	AdvancedBlendNonPremultipliedSrcColor VkBool32
	AdvancedBlendNonPremultipliedDstColor VkBool32
	AdvancedBlendCorrelatedOverlap        VkBool32
	AdvancedBlendAllOperations            VkBool32
}
type VkPipelineColorBlendAdvancedStateCreateInfoEXT struct {
	SType            VkStructureType
	PNext            uintptr
	SrcPremultiplied VkBool32
	DstPremultiplied VkBool32
	BlendOverlap     VkBlendOverlapEXT
}
type VkPipelineCoverageToColorStateCreateInfoNV struct {
	SType                   VkStructureType
	PNext                   uintptr
	Flags                   VkPipelineCoverageToColorStateCreateFlagsNV
	CoverageToColorEnable   VkBool32
	CoverageToColorLocation uint32
}
type VkPipelineCoverageModulationStateCreateInfoNV struct {
	SType                         VkStructureType
	PNext                         uintptr
	Flags                         VkPipelineCoverageModulationStateCreateFlagsNV
	CoverageModulationMode        VkCoverageModulationModeNV
	CoverageModulationTableEnable VkBool32
	CoverageModulationTableCount  uint32
	PCoverageModulationTable      *float32
}
type VkPhysicalDeviceShaderSMBuiltinsPropertiesNV struct {
	SType            VkStructureType
	PNext            uintptr
	ShaderSMCount    uint32
	ShaderWarpsPerSM uint32
}
type VkPhysicalDeviceShaderSMBuiltinsFeaturesNV struct {
	SType            VkStructureType
	PNext            uintptr
	ShaderSMBuiltins VkBool32
}
type VkDrmFormatModifierPropertiesEXT struct {
	DrmFormatModifier               uint64
	DrmFormatModifierPlaneCount     uint32
	DrmFormatModifierTilingFeatures VkFormatFeatureFlags
}
type VkDrmFormatModifierPropertiesListEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	DrmFormatModifierCount       uint32
	PDrmFormatModifierProperties *VkDrmFormatModifierPropertiesEXT
}
type VkPhysicalDeviceImageDrmFormatModifierInfoEXT struct {
	SType                 VkStructureType
	PNext                 uintptr
	DrmFormatModifier     uint64
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}
type VkImageDrmFormatModifierListCreateInfoEXT struct {
	SType                  VkStructureType
	PNext                  uintptr
	DrmFormatModifierCount uint32
	PDrmFormatModifiers    *uint64
}
type VkImageDrmFormatModifierExplicitCreateInfoEXT struct {
	SType                       VkStructureType
	PNext                       uintptr
	DrmFormatModifier           uint64
	DrmFormatModifierPlaneCount uint32
	PPlaneLayouts               *VkSubresourceLayout
}
type VkImageDrmFormatModifierPropertiesEXT struct {
	SType             VkStructureType
	PNext             uintptr
	DrmFormatModifier uint64
}
type VkValidationCacheCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkValidationCacheCreateFlagsEXT
	InitialDataSize uint64
	PInitialData    uintptr
}
type VkShaderModuleValidationCacheCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	ValidationCache VkValidationCacheEXT
}
type VkShadingRatePaletteNV struct {
	ShadingRatePaletteEntryCount uint32
	PShadingRatePaletteEntries   *VkShadingRatePaletteEntryNV
}
type VkPipelineViewportShadingRateImageStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  uintptr
	ShadingRateImageEnable VkBool32
	ViewportCount          uint32
	PShadingRatePalettes   *VkShadingRatePaletteNV
}
type VkPhysicalDeviceShadingRateImageFeaturesNV struct {
	SType                        VkStructureType
	PNext                        uintptr
	ShadingRateImage             VkBool32
	ShadingRateCoarseSampleOrder VkBool32
}
type VkPhysicalDeviceShadingRateImagePropertiesNV struct {
	SType                       VkStructureType
	PNext                       uintptr
	ShadingRateTexelSize        VkExtent2D
	ShadingRatePaletteSize      uint32
	ShadingRateMaxCoarseSamples uint32
}
type VkCoarseSampleLocationNV struct {
	PixelX uint32
	PixelY uint32
	Sample uint32
}
type VkCoarseSampleOrderCustomNV struct {
	ShadingRate         VkShadingRatePaletteEntryNV
	SampleCount         uint32
	SampleLocationCount uint32
	PSampleLocations    *VkCoarseSampleLocationNV
}
type VkPipelineViewportCoarseSampleOrderStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  uintptr
	SampleOrderType        VkCoarseSampleOrderTypeNV
	CustomSampleOrderCount uint32
	PCustomSampleOrders    *VkCoarseSampleOrderCustomNV
}
type VkRayTracingShaderGroupCreateInfoNV struct {
	SType              VkStructureType
	PNext              uintptr
	Type               VkRayTracingShaderGroupTypeKHR
	GeneralShader      uint32
	ClosestHitShader   uint32
	AnyHitShader       uint32
	IntersectionShader uint32
}
type VkRayTracingPipelineCreateInfoNV struct {
	SType              VkStructureType
	PNext              uintptr
	Flags              VkPipelineCreateFlags
	StageCount         uint32
	PStages            *VkPipelineShaderStageCreateInfo
	GroupCount         uint32
	PGroups            *VkRayTracingShaderGroupCreateInfoNV
	MaxRecursionDepth  uint32
	Layout             VkPipelineLayout
	BasePipelineHandle VkPipeline
	BasePipelineIndex  int32
}
type VkGeometryTrianglesNV struct {
	SType           VkStructureType
	PNext           uintptr
	VertexData      VkBuffer
	VertexOffset    VkDeviceSize
	VertexCount     uint32
	VertexStride    VkDeviceSize
	VertexFormat    VkFormat
	IndexData       VkBuffer
	IndexOffset     VkDeviceSize
	IndexCount      uint32
	IndexType       VkIndexType
	TransformData   VkBuffer
	TransformOffset VkDeviceSize
}
type VkGeometryAABBNV struct {
	SType    VkStructureType
	PNext    uintptr
	AabbData VkBuffer
	NumAABBs uint32
	Stride   uint32
	Offset   VkDeviceSize
}
type VkGeometryDataNV struct {
	Triangles VkGeometryTrianglesNV
	Aabbs     VkGeometryAABBNV
}
type VkGeometryNV struct {
	SType        VkStructureType
	PNext        uintptr
	GeometryType VkGeometryTypeKHR
	Geometry     VkGeometryDataNV
	Flags        VkGeometryFlagsKHR
}
type VkAccelerationStructureInfoNV struct {
	SType         VkStructureType
	PNext         uintptr
	Type          VkAccelerationStructureTypeNV
	Flags         VkBuildAccelerationStructureFlagsNV
	InstanceCount uint32
	GeometryCount uint32
	PGeometries   *VkGeometryNV
}
type VkAccelerationStructureCreateInfoNV struct {
	SType         VkStructureType
	PNext         uintptr
	CompactedSize VkDeviceSize
	Info          VkAccelerationStructureInfoNV
}
type VkBindAccelerationStructureMemoryInfoNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	AccelerationStructure VkAccelerationStructureNV
	Memory                VkDeviceMemory
	MemoryOffset          VkDeviceSize
	DeviceIndexCount      uint32
	PDeviceIndices        *uint32
}
type VkWriteDescriptorSetAccelerationStructureNV struct {
	SType                      VkStructureType
	PNext                      uintptr
	AccelerationStructureCount uint32
	PAccelerationStructures    *VkAccelerationStructureNV
}
type VkAccelerationStructureMemoryRequirementsInfoNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	Type                  VkAccelerationStructureMemoryRequirementsTypeNV
	AccelerationStructure VkAccelerationStructureNV
}
type VkPhysicalDeviceRayTracingPropertiesNV struct {
	SType                                  VkStructureType
	PNext                                  uintptr
	ShaderGroupHandleSize                  uint32
	MaxRecursionDepth                      uint32
	MaxShaderGroupStride                   uint32
	ShaderGroupBaseAlignment               uint32
	MaxGeometryCount                       uint64
	MaxInstanceCount                       uint64
	MaxTriangleCount                       uint64
	MaxDescriptorSetAccelerationStructures uint32
}
type VkTransformMatrixKHR struct {
	Matrix [3][4]float32
}
type VkAabbPositionsKHR struct {
	MinX float32
	MinY float32
	MinZ float32
	MaxX float32
	MaxY float32
	MaxZ float32
}
type VkAccelerationStructureInstanceKHR struct {
	Transform                              VkTransformMatrixKHR
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}
type VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV struct {
	SType                      VkStructureType
	PNext                      uintptr
	RepresentativeFragmentTest VkBool32
}
type VkPipelineRepresentativeFragmentTestStateCreateInfoNV struct {
	SType                            VkStructureType
	PNext                            uintptr
	RepresentativeFragmentTestEnable VkBool32
}
type VkPhysicalDeviceImageViewImageFormatInfoEXT struct {
	SType         VkStructureType
	PNext         uintptr
	ImageViewType VkImageViewType
}
type VkFilterCubicImageViewImageFormatPropertiesEXT struct {
	SType             VkStructureType
	PNext             uintptr
	FilterCubic       VkBool32
	FilterCubicMinmax VkBool32
}
type VkDeviceQueueGlobalPriorityCreateInfoEXT struct {
	SType          VkStructureType
	PNext          uintptr
	GlobalPriority VkQueueGlobalPriorityEXT
}
type VkImportMemoryHostPointerInfoEXT struct {
	SType        VkStructureType
	PNext        uintptr
	HandleType   VkExternalMemoryHandleTypeFlagBits
	PHostPointer uintptr
}
type VkMemoryHostPointerPropertiesEXT struct {
	SType          VkStructureType
	PNext          uintptr
	MemoryTypeBits uint32
}
type VkPhysicalDeviceExternalMemoryHostPropertiesEXT struct {
	SType                           VkStructureType
	PNext                           uintptr
	MinImportedHostPointerAlignment VkDeviceSize
}
type VkPipelineCompilerControlCreateInfoAMD struct {
	SType                VkStructureType
	PNext                uintptr
	CompilerControlFlags VkPipelineCompilerControlFlagsAMD
}
type VkCalibratedTimestampInfoEXT struct {
	SType      VkStructureType
	PNext      uintptr
	TimeDomain VkTimeDomainEXT
}
type VkPhysicalDeviceShaderCorePropertiesAMD struct {
	SType                      VkStructureType
	PNext                      uintptr
	ShaderEngineCount          uint32
	ShaderArraysPerEngineCount uint32
	ComputeUnitsPerShaderArray uint32
	SimdPerComputeUnit         uint32
	WavefrontsPerSimd          uint32
	WavefrontSize              uint32
	SgprsPerSimd               uint32
	MinSgprAllocation          uint32
	MaxSgprAllocation          uint32
	SgprAllocationGranularity  uint32
	VgprsPerSimd               uint32
	MinVgprAllocation          uint32
	MaxVgprAllocation          uint32
	VgprAllocationGranularity  uint32
}
type VkDeviceMemoryOverallocationCreateInfoAMD struct {
	SType                  VkStructureType
	PNext                  uintptr
	OverallocationBehavior VkMemoryOverallocationBehaviorAMD
}
type VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT struct {
	SType                  VkStructureType
	PNext                  uintptr
	MaxVertexAttribDivisor uint32
}
type VkVertexInputBindingDivisorDescriptionEXT struct {
	Binding uint32
	Divisor uint32
}
type VkPipelineVertexInputDivisorStateCreateInfoEXT struct {
	SType                     VkStructureType
	PNext                     uintptr
	VertexBindingDivisorCount uint32
	PVertexBindingDivisors    *VkVertexInputBindingDivisorDescriptionEXT
}
type VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT struct {
	SType                                  VkStructureType
	PNext                                  uintptr
	VertexAttributeInstanceRateDivisor     VkBool32
	VertexAttributeInstanceRateZeroDivisor VkBool32
}
type VkPipelineCreationFeedbackEXT struct {
	Flags    VkPipelineCreationFeedbackFlagsEXT
	Duration uint64
}
type VkPipelineCreationFeedbackCreateInfoEXT struct {
	SType                              VkStructureType
	PNext                              uintptr
	PPipelineCreationFeedback          *VkPipelineCreationFeedbackEXT
	PipelineStageCreationFeedbackCount uint32
	PPipelineStageCreationFeedbacks    *VkPipelineCreationFeedbackEXT
}
type VkPhysicalDeviceComputeShaderDerivativesFeaturesNV struct {
	SType                        VkStructureType
	PNext                        uintptr
	ComputeDerivativeGroupQuads  VkBool32
	ComputeDerivativeGroupLinear VkBool32
}
type VkPhysicalDeviceMeshShaderFeaturesNV struct {
	SType      VkStructureType
	PNext      uintptr
	TaskShader VkBool32
	MeshShader VkBool32
}
type VkPhysicalDeviceMeshShaderPropertiesNV struct {
	SType                             VkStructureType
	PNext                             uintptr
	MaxDrawMeshTasksCount             uint32
	MaxTaskWorkGroupInvocations       uint32
	MaxTaskWorkGroupSize              [3]uint32
	MaxTaskTotalMemorySize            uint32
	MaxTaskOutputCount                uint32
	MaxMeshWorkGroupInvocations       uint32
	MaxMeshWorkGroupSize              [3]uint32
	MaxMeshTotalMemorySize            uint32
	MaxMeshOutputVertices             uint32
	MaxMeshOutputPrimitives           uint32
	MaxMeshMultiviewViewCount         uint32
	MeshOutputPerVertexGranularity    uint32
	MeshOutputPerPrimitiveGranularity uint32
}
type VkDrawMeshTasksIndirectCommandNV struct {
	TaskCount uint32
	FirstTask uint32
}
type VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV struct {
	SType                     VkStructureType
	PNext                     uintptr
	FragmentShaderBarycentric VkBool32
}
type VkPhysicalDeviceShaderImageFootprintFeaturesNV struct {
	SType          VkStructureType
	PNext          uintptr
	ImageFootprint VkBool32
}
type VkPipelineViewportExclusiveScissorStateCreateInfoNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	ExclusiveScissorCount uint32
	PExclusiveScissors    *VkRect2D
}
type VkPhysicalDeviceExclusiveScissorFeaturesNV struct {
	SType            VkStructureType
	PNext            uintptr
	ExclusiveScissor VkBool32
}
type VkQueueFamilyCheckpointPropertiesNV struct {
	SType                        VkStructureType
	PNext                        uintptr
	CheckpointExecutionStageMask VkPipelineStageFlags
}
type VkCheckpointDataNV struct {
	SType             VkStructureType
	PNext             uintptr
	Stage             VkPipelineStageFlagBits
	PCheckpointMarker uintptr
}
type VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL struct {
	SType                   VkStructureType
	PNext                   uintptr
	ShaderIntegerFunctions2 VkBool32
}
type VkPerformanceValueINTEL struct {
	Type VkPerformanceValueTypeINTEL
	Data VkPerformanceValueDataINTEL
}
type VkInitializePerformanceApiInfoINTEL struct {
	SType     VkStructureType
	PNext     uintptr
	PUserData uintptr
}
type VkQueryPoolPerformanceQueryCreateInfoINTEL struct {
	SType                       VkStructureType
	PNext                       uintptr
	PerformanceCountersSampling VkQueryPoolSamplingModeINTEL
}
type VkPerformanceMarkerInfoINTEL struct {
	SType  VkStructureType
	PNext  uintptr
	Marker uint64
}
type VkPerformanceStreamMarkerInfoINTEL struct {
	SType  VkStructureType
	PNext  uintptr
	Marker uint32
}
type VkPerformanceOverrideInfoINTEL struct {
	SType     VkStructureType
	PNext     uintptr
	Type      VkPerformanceOverrideTypeINTEL
	Enable    VkBool32
	Parameter uint64
}
type VkPerformanceConfigurationAcquireInfoINTEL struct {
	SType VkStructureType
	PNext uintptr
	Type  VkPerformanceConfigurationTypeINTEL
}
type VkPhysicalDevicePCIBusInfoPropertiesEXT struct {
	SType       VkStructureType
	PNext       uintptr
	PciDomain   uint32
	PciBus      uint32
	PciDevice   uint32
	PciFunction uint32
}
type VkDisplayNativeHdrSurfaceCapabilitiesAMD struct {
	SType               VkStructureType
	PNext               uintptr
	LocalDimmingSupport VkBool32
}
type VkSwapchainDisplayNativeHdrCreateInfoAMD struct {
	SType              VkStructureType
	PNext              uintptr
	LocalDimmingEnable VkBool32
}
type VkPhysicalDeviceFragmentDensityMapFeaturesEXT struct {
	SType                                 VkStructureType
	PNext                                 uintptr
	FragmentDensityMap                    VkBool32
	FragmentDensityMapDynamic             VkBool32
	FragmentDensityMapNonSubsampledImages VkBool32
}
type VkPhysicalDeviceFragmentDensityMapPropertiesEXT struct {
	SType                       VkStructureType
	PNext                       uintptr
	MinFragmentDensityTexelSize VkExtent2D
	MaxFragmentDensityTexelSize VkExtent2D
	FragmentDensityInvocations  VkBool32
}
type VkRenderPassFragmentDensityMapCreateInfoEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	FragmentDensityMapAttachment VkAttachmentReference
}
type VkPhysicalDeviceSubgroupSizeControlFeaturesEXT struct {
	SType                VkStructureType
	PNext                uintptr
	SubgroupSizeControl  VkBool32
	ComputeFullSubgroups VkBool32
}
type VkPhysicalDeviceSubgroupSizeControlPropertiesEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	MinSubgroupSize              uint32
	MaxSubgroupSize              uint32
	MaxComputeWorkgroupSubgroups uint32
	RequiredSubgroupSizeStages   VkShaderStageFlags
}
type VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT struct {
	SType                VkStructureType
	PNext                uintptr
	RequiredSubgroupSize uint32
}
type VkPhysicalDeviceShaderCoreProperties2AMD struct {
	SType                  VkStructureType
	PNext                  uintptr
	ShaderCoreFeatures     VkShaderCorePropertiesFlagsAMD
	ActiveComputeUnitCount uint32
}
type VkPhysicalDeviceCoherentMemoryFeaturesAMD struct {
	SType                VkStructureType
	PNext                uintptr
	DeviceCoherentMemory VkBool32
}
type VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT struct {
	SType                   VkStructureType
	PNext                   uintptr
	ShaderImageInt64Atomics VkBool32
	SparseImageInt64Atomics VkBool32
}
type VkPhysicalDeviceMemoryBudgetPropertiesEXT struct {
	SType      VkStructureType
	PNext      uintptr
	HeapBudget [16]VkDeviceSize
	HeapUsage  [16]VkDeviceSize
}
type VkPhysicalDeviceMemoryPriorityFeaturesEXT struct {
	SType          VkStructureType
	PNext          uintptr
	MemoryPriority VkBool32
}
type VkMemoryPriorityAllocateInfoEXT struct {
	SType    VkStructureType
	PNext    uintptr
	Priority float32
}
type VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV struct {
	SType                            VkStructureType
	PNext                            uintptr
	DedicatedAllocationImageAliasing VkBool32
}
type VkPhysicalDeviceBufferDeviceAddressFeaturesEXT struct {
	SType                            VkStructureType
	PNext                            uintptr
	BufferDeviceAddress              VkBool32
	BufferDeviceAddressCaptureReplay VkBool32
	BufferDeviceAddressMultiDevice   VkBool32
}
type VkBufferDeviceAddressCreateInfoEXT struct {
	SType         VkStructureType
	PNext         uintptr
	DeviceAddress VkDeviceAddress
}
type VkPhysicalDeviceToolPropertiesEXT struct {
	SType       VkStructureType
	PNext       uintptr
	Name        [256]byte
	Version     [256]byte
	Purposes    VkToolPurposeFlagsEXT
	Description [256]byte
	Layer       [256]byte
}
type VkValidationFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          uintptr
	EnabledValidationFeatureCount  uint32
	PEnabledValidationFeatures     *VkValidationFeatureEnableEXT
	DisabledValidationFeatureCount uint32
	PDisabledValidationFeatures    *VkValidationFeatureDisableEXT
}
type VkCooperativeMatrixPropertiesNV struct {
	SType VkStructureType
	PNext uintptr
	MSize uint32
	NSize uint32
	KSize uint32
	AType VkComponentTypeNV
	BType VkComponentTypeNV
	CType VkComponentTypeNV
	DType VkComponentTypeNV
	Scope VkScopeNV
}
type VkPhysicalDeviceCooperativeMatrixFeaturesNV struct {
	SType                               VkStructureType
	PNext                               uintptr
	CooperativeMatrix                   VkBool32
	CooperativeMatrixRobustBufferAccess VkBool32
}
type VkPhysicalDeviceCooperativeMatrixPropertiesNV struct {
	SType                            VkStructureType
	PNext                            uintptr
	CooperativeMatrixSupportedStages VkShaderStageFlags
}
type VkPhysicalDeviceCoverageReductionModeFeaturesNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	CoverageReductionMode VkBool32
}
type VkPipelineCoverageReductionStateCreateInfoNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	Flags                 VkPipelineCoverageReductionStateCreateFlagsNV
	CoverageReductionMode VkCoverageReductionModeNV
}
type VkFramebufferMixedSamplesCombinationNV struct {
	SType                 VkStructureType
	PNext                 uintptr
	CoverageReductionMode VkCoverageReductionModeNV
	RasterizationSamples  VkSampleCountFlagBits
	DepthStencilSamples   VkSampleCountFlags
	ColorSamples          VkSampleCountFlags
}
type VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT struct {
	SType                              VkStructureType
	PNext                              uintptr
	FragmentShaderSampleInterlock      VkBool32
	FragmentShaderPixelInterlock       VkBool32
	FragmentShaderShadingRateInterlock VkBool32
}
type VkPhysicalDeviceYcbcrImageArraysFeaturesEXT struct {
	SType            VkStructureType
	PNext            uintptr
	YcbcrImageArrays VkBool32
}
type VkPhysicalDeviceProvokingVertexFeaturesEXT struct {
	SType                                     VkStructureType
	PNext                                     uintptr
	ProvokingVertexLast                       VkBool32
	TransformFeedbackPreservesProvokingVertex VkBool32
}
type VkPhysicalDeviceProvokingVertexPropertiesEXT struct {
	SType                                                VkStructureType
	PNext                                                uintptr
	ProvokingVertexModePerPipeline                       VkBool32
	TransformFeedbackPreservesTriangleFanProvokingVertex VkBool32
}
type VkPipelineRasterizationProvokingVertexStateCreateInfoEXT struct {
	SType               VkStructureType
	PNext               uintptr
	ProvokingVertexMode VkProvokingVertexModeEXT
}
type VkHeadlessSurfaceCreateInfoEXT struct {
	SType VkStructureType
	PNext uintptr
	Flags VkHeadlessSurfaceCreateFlagsEXT
}
type VkPhysicalDeviceLineRasterizationFeaturesEXT struct {
	SType                    VkStructureType
	PNext                    uintptr
	RectangularLines         VkBool32
	BresenhamLines           VkBool32
	SmoothLines              VkBool32
	StippledRectangularLines VkBool32
	StippledBresenhamLines   VkBool32
	StippledSmoothLines      VkBool32
}
type VkPhysicalDeviceLineRasterizationPropertiesEXT struct {
	SType                     VkStructureType
	PNext                     uintptr
	LineSubPixelPrecisionBits uint32
}
type VkPipelineRasterizationLineStateCreateInfoEXT struct {
	SType                 VkStructureType
	PNext                 uintptr
	LineRasterizationMode VkLineRasterizationModeEXT
	StippledLineEnable    VkBool32
	LineStippleFactor     uint32
	LineStipplePattern    uint16
}
type VkPhysicalDeviceShaderAtomicFloatFeaturesEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	ShaderBufferFloat32Atomics   VkBool32
	ShaderBufferFloat32AtomicAdd VkBool32
	ShaderBufferFloat64Atomics   VkBool32
	ShaderBufferFloat64AtomicAdd VkBool32
	ShaderSharedFloat32Atomics   VkBool32
	ShaderSharedFloat32AtomicAdd VkBool32
	ShaderSharedFloat64Atomics   VkBool32
	ShaderSharedFloat64AtomicAdd VkBool32
	ShaderImageFloat32Atomics    VkBool32
	ShaderImageFloat32AtomicAdd  VkBool32
	SparseImageFloat32Atomics    VkBool32
	SparseImageFloat32AtomicAdd  VkBool32
}
type VkPhysicalDeviceIndexTypeUint8FeaturesEXT struct {
	SType          VkStructureType
	PNext          uintptr
	IndexTypeUint8 VkBool32
}
type VkPhysicalDeviceExtendedDynamicStateFeaturesEXT struct {
	SType                VkStructureType
	PNext                uintptr
	ExtendedDynamicState VkBool32
}
type VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT struct {
	SType                           VkStructureType
	PNext                           uintptr
	ShaderBufferFloat16Atomics      VkBool32
	ShaderBufferFloat16AtomicAdd    VkBool32
	ShaderBufferFloat16AtomicMinMax VkBool32
	ShaderBufferFloat32AtomicMinMax VkBool32
	ShaderBufferFloat64AtomicMinMax VkBool32
	ShaderSharedFloat16Atomics      VkBool32
	ShaderSharedFloat16AtomicAdd    VkBool32
	ShaderSharedFloat16AtomicMinMax VkBool32
	ShaderSharedFloat32AtomicMinMax VkBool32
	ShaderSharedFloat64AtomicMinMax VkBool32
	ShaderImageFloat32AtomicMinMax  VkBool32
	SparseImageFloat32AtomicMinMax  VkBool32
}
type VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          uintptr
	ShaderDemoteToHelperInvocation VkBool32
}
type VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV struct {
	SType                                    VkStructureType
	PNext                                    uintptr
	MaxGraphicsShaderGroupCount              uint32
	MaxIndirectSequenceCount                 uint32
	MaxIndirectCommandsTokenCount            uint32
	MaxIndirectCommandsStreamCount           uint32
	MaxIndirectCommandsTokenOffset           uint32
	MaxIndirectCommandsStreamStride          uint32
	MinSequencesCountBufferOffsetAlignment   uint32
	MinSequencesIndexBufferOffsetAlignment   uint32
	MinIndirectCommandsBufferOffsetAlignment uint32
}
type VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV struct {
	SType                   VkStructureType
	PNext                   uintptr
	DeviceGeneratedCommands VkBool32
}
type VkGraphicsShaderGroupCreateInfoNV struct {
	SType              VkStructureType
	PNext              uintptr
	StageCount         uint32
	PStages            *VkPipelineShaderStageCreateInfo
	PVertexInputState  *VkPipelineVertexInputStateCreateInfo
	PTessellationState *VkPipelineTessellationStateCreateInfo
}
type VkGraphicsPipelineShaderGroupsCreateInfoNV struct {
	SType         VkStructureType
	PNext         uintptr
	GroupCount    uint32
	PGroups       *VkGraphicsShaderGroupCreateInfoNV
	PipelineCount uint32
	PPipelines    *VkPipeline
}
type VkBindShaderGroupIndirectCommandNV struct {
	GroupIndex uint32
}
type VkBindIndexBufferIndirectCommandNV struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	IndexType     VkIndexType
}
type VkBindVertexBufferIndirectCommandNV struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	Stride        uint32
}
type VkSetStateFlagsIndirectCommandNV struct {
	Data uint32
}
type VkIndirectCommandsStreamNV struct {
	Buffer VkBuffer
	Offset VkDeviceSize
}
type VkIndirectCommandsLayoutTokenNV struct {
	SType                   VkStructureType
	PNext                   uintptr
	TokenType               VkIndirectCommandsTokenTypeNV
	Stream                  uint32
	Offset                  uint32
	VertexBindingUnit       uint32
	VertexDynamicStride     VkBool32
	PushantPipelineLayout   VkPipelineLayout
	PushantShaderStageFlags VkShaderStageFlags
	PushantOffset           uint32
	PushantSize             uint32
	IndirectStateFlags      VkIndirectStateFlagsNV
	IndexTypeCount          uint32
	PIndexTypes             *VkIndexType
	PIndexTypeValues        *uint32
}
type VkIndirectCommandsLayoutCreateInfoNV struct {
	SType             VkStructureType
	PNext             uintptr
	Flags             VkIndirectCommandsLayoutUsageFlagsNV
	PipelineBindPoint VkPipelineBindPoint
	TokenCount        uint32
	PTokens           *VkIndirectCommandsLayoutTokenNV
	StreamCount       uint32
	PStreamStrides    *uint32
}
type VkGeneratedCommandsInfoNV struct {
	SType                  VkStructureType
	PNext                  uintptr
	PipelineBindPoint      VkPipelineBindPoint
	Pipeline               VkPipeline
	IndirectCommandsLayout VkIndirectCommandsLayoutNV
	StreamCount            uint32
	PStreams               *VkIndirectCommandsStreamNV
	SequencesCount         uint32
	PreprocessBuffer       VkBuffer
	PreprocessOffset       VkDeviceSize
	PreprocessSize         VkDeviceSize
	SequencesCountBuffer   VkBuffer
	SequencesCountOffset   VkDeviceSize
	SequencesIndexBuffer   VkBuffer
	SequencesIndexOffset   VkDeviceSize
}
type VkGeneratedCommandsMemoryRequirementsInfoNV struct {
	SType                  VkStructureType
	PNext                  uintptr
	PipelineBindPoint      VkPipelineBindPoint
	Pipeline               VkPipeline
	IndirectCommandsLayout VkIndirectCommandsLayoutNV
	MaxSequencesCount      uint32
}
type VkPhysicalDeviceInheritedViewportScissorFeaturesNV struct {
	SType                      VkStructureType
	PNext                      uintptr
	InheritedViewportScissor2D VkBool32
}
type VkCommandBufferInheritanceViewportScissorInfoNV struct {
	SType              VkStructureType
	PNext              uintptr
	ViewportScissor2D  VkBool32
	ViewportDepthCount uint32
	PViewportDepths    *VkViewport
}
type VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT struct {
	SType                VkStructureType
	PNext                uintptr
	TexelBufferAlignment VkBool32
}
type VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT struct {
	SType                                        VkStructureType
	PNext                                        uintptr
	StorageTexelBufferOffsetAlignmentBytes       VkDeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment VkBool32
	UniformTexelBufferOffsetAlignmentBytes       VkDeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment VkBool32
}
type VkRenderPassTransformBeginInfoQCOM struct {
	SType     VkStructureType
	PNext     uintptr
	Transform VkSurfaceTransformFlagBitsKHR
}
type VkCommandBufferInheritanceRenderPassTransformInfoQCOM struct {
	SType      VkStructureType
	PNext      uintptr
	Transform  VkSurfaceTransformFlagBitsKHR
	RenderArea VkRect2D
}
type VkPhysicalDeviceDeviceMemoryReportFeaturesEXT struct {
	SType              VkStructureType
	PNext              uintptr
	DeviceMemoryReport VkBool32
}
type VkDeviceMemoryReportCallbackDataEXT struct {
	SType          VkStructureType
	PNext          uintptr
	Flags          VkDeviceMemoryReportFlagsEXT
	Type           VkDeviceMemoryReportEventTypeEXT
	MemoryObjectId uint64
	Size           VkDeviceSize
	ObjectType     VkObjectType
	ObjectHandle   uint64
	HeapIndex      uint32
}
type VkDeviceDeviceMemoryReportCreateInfoEXT struct {
	SType           VkStructureType
	PNext           uintptr
	Flags           VkDeviceMemoryReportFlagsEXT
	PfnUserCallback uintptr
	PUserData       uintptr
}
type VkPhysicalDeviceRobustness2FeaturesEXT struct {
	SType               VkStructureType
	PNext               uintptr
	RobustBufferAccess2 VkBool32
	RobustImageAccess2  VkBool32
	NullDescriptor      VkBool32
}
type VkPhysicalDeviceRobustness2PropertiesEXT struct {
	SType                                  VkStructureType
	PNext                                  uintptr
	RobustStorageBufferAccessSizeAlignment VkDeviceSize
	RobustUniformBufferAccessSizeAlignment VkDeviceSize
}
type VkSamplerCustomBorderColorCreateInfoEXT struct {
	SType             VkStructureType
	PNext             uintptr
	CustomBorderColor VkClearColorValue
	Format            VkFormat
}
type VkPhysicalDeviceCustomBorderColorPropertiesEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	MaxCustomBorderColorSamplers uint32
}
type VkPhysicalDeviceCustomBorderColorFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          uintptr
	CustomBorderColors             VkBool32
	CustomBorderColorWithoutFormat VkBool32
}
type VkPhysicalDevicePrivateDataFeaturesEXT struct {
	SType       VkStructureType
	PNext       uintptr
	PrivateData VkBool32
}
type VkDevicePrivateDataCreateInfoEXT struct {
	SType                       VkStructureType
	PNext                       uintptr
	PrivateDataSlotRequestCount uint32
}
type VkPrivateDataSlotCreateInfoEXT struct {
	SType VkStructureType
	PNext uintptr
	Flags VkPrivateDataSlotCreateFlagsEXT
}
type VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT struct {
	SType                        VkStructureType
	PNext                        uintptr
	PipelineCreationCacheControl VkBool32
}
type VkPhysicalDeviceDiagnosticsConfigFeaturesNV struct {
	SType             VkStructureType
	PNext             uintptr
	DiagnosticsConfig VkBool32
}
type VkDeviceDiagnosticsConfigCreateInfoNV struct {
	SType VkStructureType
	PNext uintptr
	Flags VkDeviceDiagnosticsConfigFlagsNV
}
type VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV struct {
	SType                            VkStructureType
	PNext                            uintptr
	FragmentShadingRateEnums         VkBool32
	SupersampleFragmentShadingRates  VkBool32
	NoInvocationFragmentShadingRates VkBool32
}
type VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV struct {
	SType                                 VkStructureType
	PNext                                 uintptr
	MaxFragmentShadingRateInvocationCount VkSampleCountFlagBits
}
type VkPipelineFragmentShadingRateEnumStateCreateInfoNV struct {
	SType           VkStructureType
	PNext           uintptr
	ShadingRateType VkFragmentShadingRateTypeNV
	ShadingRate     VkFragmentShadingRateNV
	CombinerOps     [2]VkFragmentShadingRateCombinerOpKHR
}
type VkAccelerationStructureGeometryMotionTrianglesDataNV struct {
	SType      VkStructureType
	PNext      uintptr
	VertexData VkDeviceOrHostAddressConstKHR
}
type VkAccelerationStructureMotionInfoNV struct {
	SType        VkStructureType
	PNext        uintptr
	MaxInstances uint32
	Flags        VkAccelerationStructureMotionInfoFlagsNV
}
type VkAccelerationStructureMatrixMotionInstanceNV struct {
	TransformT0                            VkTransformMatrixKHR
	TransformT1                            VkTransformMatrixKHR
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}
type VkSRTDataNV struct {
	SX  float32
	A   float32
	B   float32
	PVX float32
	SY  float32
	C   float32
	PVY float32
	SZ  float32
	PVZ float32
	QX  float32
	QY  float32
	QZ  float32
	QW  float32
	TX  float32
	TY  float32
	TZ  float32
}
type VkAccelerationStructureSRTMotionInstanceNV struct {
	TransformT0                            VkSRTDataNV
	TransformT1                            VkSRTDataNV
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}
type VkAccelerationStructureMotionInstanceNV struct {
	Type  VkAccelerationStructureMotionInstanceTypeNV
	Flags VkAccelerationStructureMotionInstanceFlagsNV
	Data  VkAccelerationStructureMotionInstanceDataNV
}
type VkPhysicalDeviceRayTracingMotionBlurFeaturesNV struct {
	SType                                         VkStructureType
	PNext                                         uintptr
	RayTracingMotionBlur                          VkBool32
	RayTracingMotionBlurPipelineTraceRaysIndirect VkBool32
}
type VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT struct {
	SType                 VkStructureType
	PNext                 uintptr
	Ycbcr2plane444Formats VkBool32
}
type VkPhysicalDeviceFragmentDensityMap2FeaturesEXT struct {
	SType                      VkStructureType
	PNext                      uintptr
	FragmentDensityMapDeferred VkBool32
}
type VkPhysicalDeviceFragmentDensityMap2PropertiesEXT struct {
	SType                                VkStructureType
	PNext                                uintptr
	SubsampledLoads                      VkBool32
	SubsampledCoarseReructionEarlyAccess VkBool32
	MaxSubsampledArrayLayers             uint32
	MaxDescriptorSetSubsampledSamplers   uint32
}
type VkCopyCommandTransformInfoQCOM struct {
	SType     VkStructureType
	PNext     uintptr
	Transform VkSurfaceTransformFlagBitsKHR
}
type VkPhysicalDeviceImageRobustnessFeaturesEXT struct {
	SType             VkStructureType
	PNext             uintptr
	RobustImageAccess VkBool32
}
type VkPhysicalDevice4444FormatsFeaturesEXT struct {
	SType          VkStructureType
	PNext          uintptr
	FormatA4R4G4B4 VkBool32
	FormatA4B4G4R4 VkBool32
}
type VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE struct {
	SType                 VkStructureType
	PNext                 uintptr
	MutableDescriptorType VkBool32
}
type VkMutableDescriptorTypeListVALVE struct {
	DescriptorTypeCount uint32
	PDescriptorTypes    *VkDescriptorType
}
type VkMutableDescriptorTypeCreateInfoVALVE struct {
	SType                          VkStructureType
	PNext                          uintptr
	MutableDescriptorTypeListCount uint32
	PMutableDescriptorTypeLists    *VkMutableDescriptorTypeListVALVE
}
type VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT struct {
	SType                   VkStructureType
	PNext                   uintptr
	VertexInputDynamicState VkBool32
}
type VkVertexInputBindingDescription2EXT struct {
	SType     VkStructureType
	PNext     uintptr
	Binding   uint32
	Stride    uint32
	InputRate VkVertexInputRate
	Divisor   uint32
}
type VkVertexInputAttributeDescription2EXT struct {
	SType    VkStructureType
	PNext    uintptr
	Location uint32
	Binding  uint32
	Format   VkFormat
	Offset   uint32
}
type VkPhysicalDeviceDrmPropertiesEXT struct {
	SType        VkStructureType
	PNext        uintptr
	HasPrimary   VkBool32
	HasRender    VkBool32
	PrimaryMajor int64
	PrimaryMinor int64
	RenderMajor  int64
	RenderMinor  int64
}
type VkSubpassShadingPipelineCreateInfoHUAWEI struct {
	SType      VkStructureType
	PNext      uintptr
	RenderPass VkRenderPass
	Subpass    uint32
}
type VkPhysicalDeviceSubpassShadingFeaturesHUAWEI struct {
	SType          VkStructureType
	PNext          uintptr
	SubpassShading VkBool32
}
type VkPhysicalDeviceSubpassShadingPropertiesHUAWEI struct {
	SType                                     VkStructureType
	PNext                                     uintptr
	MaxSubpassShadingWorkgroupSizeAspectRatio uint32
}
type VkPhysicalDeviceInvocationMaskFeaturesHUAWEI struct {
	SType          VkStructureType
	PNext          uintptr
	InvocationMask VkBool32
}
type VkMemoryGetRemoteAddressInfoNV struct {
	SType      VkStructureType
	PNext      uintptr
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}
type VkPhysicalDeviceExternalMemoryRDMAFeaturesNV struct {
	SType              VkStructureType
	PNext              uintptr
	ExternalMemoryRDMA VkBool32
}
type VkPhysicalDeviceExtendedDynamicState2FeaturesEXT struct {
	SType                                   VkStructureType
	PNext                                   uintptr
	ExtendedDynamicState2                   VkBool32
	ExtendedDynamicState2LogicOp            VkBool32
	ExtendedDynamicState2PatchControlPoints VkBool32
}
type VkPhysicalDeviceColorWriteEnableFeaturesEXT struct {
	SType            VkStructureType
	PNext            uintptr
	ColorWriteEnable VkBool32
}
type VkPipelineColorWriteCreateInfoEXT struct {
	SType              VkStructureType
	PNext              uintptr
	AttachmentCount    uint32
	PColorWriteEnables *VkBool32
}
type VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT struct {
	SType               VkStructureType
	PNext               uintptr
	GlobalPriorityQuery VkBool32
}
type VkQueueFamilyGlobalPriorityPropertiesEXT struct {
	SType         VkStructureType
	PNext         uintptr
	PriorityCount uint32
	Priorities    [16]VkQueueGlobalPriorityEXT
}
type VkPhysicalDeviceMultiDrawFeaturesEXT struct {
	SType     VkStructureType
	PNext     uintptr
	MultiDraw VkBool32
}
type VkPhysicalDeviceMultiDrawPropertiesEXT struct {
	SType             VkStructureType
	PNext             uintptr
	MaxMultiDrawCount uint32
}
type VkMultiDrawInfoEXT struct {
	FirstVertex uint32
	VertexCount uint32
}
type VkMultiDrawIndexedInfoEXT struct {
	FirstIndex   uint32
	IndexCount   uint32
	VertexOffset int32
}
type VkAccelerationStructureBuildRangeInfoKHR struct {
	PrimitiveCount  uint32
	PrimitiveOffset uint32
	FirstVertex     uint32
	TransformOffset uint32
}
type VkAccelerationStructureGeometryTrianglesDataKHR struct {
	SType         VkStructureType
	PNext         uintptr
	VertexFormat  VkFormat
	VertexData    VkDeviceOrHostAddressConstKHR
	VertexStride  VkDeviceSize
	MaxVertex     uint32
	IndexType     VkIndexType
	IndexData     VkDeviceOrHostAddressConstKHR
	TransformData VkDeviceOrHostAddressConstKHR
}
type VkAccelerationStructureGeometryAabbsDataKHR struct {
	SType  VkStructureType
	PNext  uintptr
	Data   VkDeviceOrHostAddressConstKHR
	Stride VkDeviceSize
}
type VkAccelerationStructureGeometryInstancesDataKHR struct {
	SType           VkStructureType
	PNext           uintptr
	ArrayOfPointers VkBool32
	Data            VkDeviceOrHostAddressConstKHR
}
type VkAccelerationStructureGeometryKHR struct {
	SType        VkStructureType
	PNext        uintptr
	GeometryType VkGeometryTypeKHR
	Geometry     VkAccelerationStructureGeometryDataKHR
	Flags        VkGeometryFlagsKHR
}
type VkAccelerationStructureBuildGeometryInfoKHR struct {
	SType                    VkStructureType
	PNext                    uintptr
	Type                     VkAccelerationStructureTypeKHR
	Flags                    VkBuildAccelerationStructureFlagsKHR
	Mode                     VkBuildAccelerationStructureModeKHR
	SrcAccelerationStructure VkAccelerationStructureKHR
	DstAccelerationStructure VkAccelerationStructureKHR
	GeometryCount            uint32
	PGeometries              *VkAccelerationStructureGeometryKHR
	PpGeometries             **VkAccelerationStructureGeometryKHR
	ScratchData              VkDeviceOrHostAddressKHR
}
type VkAccelerationStructureCreateInfoKHR struct {
	SType         VkStructureType
	PNext         uintptr
	CreateFlags   VkAccelerationStructureCreateFlagsKHR
	Buffer        VkBuffer
	Offset        VkDeviceSize
	Size          VkDeviceSize
	Type          VkAccelerationStructureTypeKHR
	DeviceAddress VkDeviceAddress
}
type VkWriteDescriptorSetAccelerationStructureKHR struct {
	SType                      VkStructureType
	PNext                      uintptr
	AccelerationStructureCount uint32
	PAccelerationStructures    *VkAccelerationStructureKHR
}
type VkPhysicalDeviceAccelerationStructureFeaturesKHR struct {
	SType                                                 VkStructureType
	PNext                                                 uintptr
	AccelerationStructure                                 VkBool32
	AccelerationStructureCaptureReplay                    VkBool32
	AccelerationStructureIndirectBuild                    VkBool32
	AccelerationStructureHostCommands                     VkBool32
	DescriptorBindingAccelerationStructureUpdateAfterBind VkBool32
}
type VkPhysicalDeviceAccelerationStructurePropertiesKHR struct {
	SType                                                      VkStructureType
	PNext                                                      uintptr
	MaxGeometryCount                                           uint64
	MaxInstanceCount                                           uint64
	MaxPrimitiveCount                                          uint64
	MaxPerStageDescriptorAccelerationStructures                uint32
	MaxPerStageDescriptorUpdateAfterBindAccelerationStructures uint32
	MaxDescriptorSetAccelerationStructures                     uint32
	MaxDescriptorSetUpdateAfterBindAccelerationStructures      uint32
	MinAccelerationStructureScratchOffsetAlignment             uint32
}
type VkAccelerationStructureDeviceAddressInfoKHR struct {
	SType                 VkStructureType
	PNext                 uintptr
	AccelerationStructure VkAccelerationStructureKHR
}
type VkAccelerationStructureVersionInfoKHR struct {
	SType        VkStructureType
	PNext        uintptr
	PVersionData *uint8
}
type VkCopyAccelerationStructureToMemoryInfoKHR struct {
	SType VkStructureType
	PNext uintptr
	SRC   VkAccelerationStructureKHR
	DST   VkDeviceOrHostAddressKHR
	Mode  VkCopyAccelerationStructureModeKHR
}
type VkCopyMemoryToAccelerationStructureInfoKHR struct {
	SType VkStructureType
	PNext uintptr
	SRC   VkDeviceOrHostAddressConstKHR
	DST   VkAccelerationStructureKHR
	Mode  VkCopyAccelerationStructureModeKHR
}
type VkCopyAccelerationStructureInfoKHR struct {
	SType VkStructureType
	PNext uintptr
	SRC   VkAccelerationStructureKHR
	DST   VkAccelerationStructureKHR
	Mode  VkCopyAccelerationStructureModeKHR
}
type VkAccelerationStructureBuildSizesInfoKHR struct {
	SType                     VkStructureType
	PNext                     uintptr
	AccelerationStructureSize VkDeviceSize
	UpdateScratchSize         VkDeviceSize
	BuildScratchSize          VkDeviceSize
}
type VkRayTracingShaderGroupCreateInfoKHR struct {
	SType                           VkStructureType
	PNext                           uintptr
	Type                            VkRayTracingShaderGroupTypeKHR
	GeneralShader                   uint32
	ClosestHitShader                uint32
	AnyHitShader                    uint32
	IntersectionShader              uint32
	PShaderGroupCaptureReplayHandle uintptr
}
type VkRayTracingPipelineInterfaceCreateInfoKHR struct {
	SType                          VkStructureType
	PNext                          uintptr
	MaxPipelineRayPayloadSize      uint32
	MaxPipelineRayHitAttributeSize uint32
}
type VkRayTracingPipelineCreateInfoKHR struct {
	SType                        VkStructureType
	PNext                        uintptr
	Flags                        VkPipelineCreateFlags
	StageCount                   uint32
	PStages                      *VkPipelineShaderStageCreateInfo
	GroupCount                   uint32
	PGroups                      *VkRayTracingShaderGroupCreateInfoKHR
	MaxPipelineRayRecursionDepth uint32
	PLibraryInfo                 *VkPipelineLibraryCreateInfoKHR
	PLibraryInterface            *VkRayTracingPipelineInterfaceCreateInfoKHR
	PDynamicState                *VkPipelineDynamicStateCreateInfo
	Layout                       VkPipelineLayout
	BasePipelineHandle           VkPipeline
	BasePipelineIndex            int32
}
type VkPhysicalDeviceRayTracingPipelineFeaturesKHR struct {
	SType                                                 VkStructureType
	PNext                                                 uintptr
	RayTracingPipeline                                    VkBool32
	RayTracingPipelineShaderGroupHandleCaptureReplay      VkBool32
	RayTracingPipelineShaderGroupHandleCaptureReplayMixed VkBool32
	RayTracingPipelineTraceRaysIndirect                   VkBool32
	RayTraversalPrimitiveCulling                          VkBool32
}
type VkPhysicalDeviceRayTracingPipelinePropertiesKHR struct {
	SType                              VkStructureType
	PNext                              uintptr
	ShaderGroupHandleSize              uint32
	MaxRayRecursionDepth               uint32
	MaxShaderGroupStride               uint32
	ShaderGroupBaseAlignment           uint32
	ShaderGroupHandleCaptureReplaySize uint32
	MaxRayDispatchInvocationCount      uint32
	ShaderGroupHandleAlignment         uint32
	MaxRayHitAttributeSize             uint32
}
type VkStridedDeviceAddressRegionKHR struct {
	DeviceAddress VkDeviceAddress
	Stride        VkDeviceSize
	Size          VkDeviceSize
}
type VkTraceRaysIndirectCommandKHR struct {
	Width  uint32
	Height uint32
	Depth  uint32
}
type VkPhysicalDeviceRayQueryFeaturesKHR struct {
	SType    VkStructureType
	PNext    uintptr
	RayQuery VkBool32
}

func CreateInstance(
	pCreateInfo *VkInstanceCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pInstance VkInstance, result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCreateInstance.Addr(),
		3,
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pInstance)),
	)
	result = GetError((int32)(ret))
	return
}
func DestroyInstance(
	instance VkInstance,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyInstance.Addr(),
		2,
		uintptr(instance),
		uintptr(unsafe.Pointer(pAllocator)),
		0,
	)
	return
}
func EnumeratePhysicalDevices(
	instance VkInstance,
	pPhysicalDeviceCount *uint32,
	pPhysicalDevices *VkPhysicalDevice,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumeratePhysicalDevices.Addr(),
		3,
		uintptr(instance),
		uintptr(unsafe.Pointer(pPhysicalDeviceCount)),
		uintptr(unsafe.Pointer(pPhysicalDevices)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceFeatures(
	physicalDevice VkPhysicalDevice,
	pFeatures *VkPhysicalDeviceFeatures,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFeatures.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFeatures)),
		0,
	)
	return
}
func GetPhysicalDeviceFormatProperties(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	pFormatProperties *VkFormatProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFormatProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(unsafe.Pointer(pFormatProperties)),
	)
	return
}
func GetPhysicalDeviceImageFormatProperties(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	vtype VkImageType,
	tiling VkImageTiling,
	usage VkImageUsageFlags,
	flags VkImageCreateFlags,
	pImageFormatProperties *VkImageFormatProperties,
) (result error) {
	ret, _, _ := syscall.Syscall9(
		PFN_vkGetPhysicalDeviceImageFormatProperties.Addr(),
		7,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(vtype),
		uintptr(tiling),
		uintptr(usage),
		uintptr(flags),
		uintptr(unsafe.Pointer(pImageFormatProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceProperties(
	physicalDevice VkPhysicalDevice,
	pProperties *VkPhysicalDeviceProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceProperties.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pProperties)),
		0,
	)
	return
}
func GetPhysicalDeviceQueueFamilyProperties(
	physicalDevice VkPhysicalDevice,
	pQueueFamilyPropertyCount *uint32,
	pQueueFamilyProperties *VkQueueFamilyProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceQueueFamilyProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)),
		uintptr(unsafe.Pointer(pQueueFamilyProperties)),
	)
	return
}
func GetPhysicalDeviceMemoryProperties(
	physicalDevice VkPhysicalDevice,
	pMemoryProperties *VkPhysicalDeviceMemoryProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceMemoryProperties.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pMemoryProperties)),
		0,
	)
	return
}
func GetInstanceProcAddr(
	instance VkInstance,
	pName string,
) (result uintptr) {
	charptrpName := TEXT(pName)
	ret, _, _ := syscall.Syscall(
		PFN_vkGetInstanceProcAddr.Addr(),
		2,
		uintptr(instance),
		uintptr(unsafe.Pointer(charptrpName)),
		0,
	)
	result = (uintptr)(ret)
	return
}
func GetDeviceProcAddr(
	device VkDevice,
	pName string,
) (result uintptr) {
	charptrpName := TEXT(pName)
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceProcAddr.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(charptrpName)),
		0,
	)
	result = (uintptr)(ret)
	return
}
func CreateDevice(
	physicalDevice VkPhysicalDevice,
	pCreateInfo *VkDeviceCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pDevice VkDevice, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDevice.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pDevice)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDevice(
	device VkDevice,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDevice.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pAllocator)),
		0,
	)
	return
}
func EnumerateInstanceExtensionProperties(
	pLayerName string,
	pPropertyCount *uint32,
	pProperties *VkExtensionProperties,
) (result error) {
	charptrpLayerName := TEXT(pLayerName)
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumerateInstanceExtensionProperties.Addr(),
		3,
		uintptr(unsafe.Pointer(charptrpLayerName)),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func EnumerateDeviceExtensionProperties(
	physicalDevice VkPhysicalDevice,
	pLayerName string,
	pPropertyCount *uint32,
	pProperties *VkExtensionProperties,
) (result error) {
	charptrpLayerName := TEXT(pLayerName)
	ret, _, _ := syscall.Syscall6(
		PFN_vkEnumerateDeviceExtensionProperties.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(charptrpLayerName)),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func EnumerateInstanceLayerProperties(
	pPropertyCount *uint32,
	pProperties *VkLayerProperties,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumerateInstanceLayerProperties.Addr(),
		2,
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func EnumerateDeviceLayerProperties(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkLayerProperties,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumerateDeviceLayerProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceQueue(
	device VkDevice,
	queueFamilyIndex uint32,
	queueIndex uint32,
	pQueue *VkQueue,
) {
	syscall.Syscall6(
		PFN_vkGetDeviceQueue.Addr(),
		4,
		uintptr(device),
		uintptr(queueFamilyIndex),
		uintptr(queueIndex),
		uintptr(unsafe.Pointer(pQueue)),
		0,
		0,
	)
	return
}
func QueueSubmit(
	queue VkQueue,
	submitCount uint32,
	pSubmits *VkSubmitInfo,
	fence VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkQueueSubmit.Addr(),
		4,
		uintptr(queue),
		uintptr(submitCount),
		uintptr(unsafe.Pointer(pSubmits)),
		uintptr(fence),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func QueueWaitIdle(
	queue VkQueue,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkQueueWaitIdle.Addr(),
		1,
		uintptr(queue),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DeviceWaitIdle(
	device VkDevice,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkDeviceWaitIdle.Addr(),
		1,
		uintptr(device),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func AllocateMemory(
	device VkDevice,
	pAllocateInfo *VkMemoryAllocateInfo,
	pAllocator *VkAllocationCallbacks,
	pMemory *VkDeviceMemory,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkAllocateMemory.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pAllocateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(pMemory)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func FreeMemory(
	device VkDevice,
	memory VkDeviceMemory,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkFreeMemory.Addr(),
		3,
		uintptr(device),
		uintptr(memory),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func MapMemory(
	device VkDevice,
	memory VkDeviceMemory,
	offset VkDeviceSize,
	size VkDeviceSize,
	flags VkMemoryMapFlags,
	ppData **interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkMapMemory.Addr(),
		6,
		uintptr(device),
		uintptr(memory),
		uintptr(offset),
		uintptr(size),
		uintptr(flags),
		uintptr(unsafe.Pointer(ppData)),
	)
	result = GetError((int32)(ret))
	return
}
func UnmapMemory(
	device VkDevice,
	memory VkDeviceMemory,
) {
	syscall.Syscall(
		PFN_vkUnmapMemory.Addr(),
		2,
		uintptr(device),
		uintptr(memory),
		0,
	)
	return
}
func FlushMappedMemoryRanges(
	device VkDevice,
	memoryRangeCount uint32,
	pMemoryRanges *VkMappedMemoryRange,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkFlushMappedMemoryRanges.Addr(),
		3,
		uintptr(device),
		uintptr(memoryRangeCount),
		uintptr(unsafe.Pointer(pMemoryRanges)),
	)
	result = GetError((int32)(ret))
	return
}
func InvalidateMappedMemoryRanges(
	device VkDevice,
	memoryRangeCount uint32,
	pMemoryRanges *VkMappedMemoryRange,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkInvalidateMappedMemoryRanges.Addr(),
		3,
		uintptr(device),
		uintptr(memoryRangeCount),
		uintptr(unsafe.Pointer(pMemoryRanges)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceMemoryCommitment(
	device VkDevice,
	memory VkDeviceMemory,
	pCommittedMemoryInBytes *VkDeviceSize,
) {
	syscall.Syscall(
		PFN_vkGetDeviceMemoryCommitment.Addr(),
		3,
		uintptr(device),
		uintptr(memory),
		uintptr(unsafe.Pointer(pCommittedMemoryInBytes)),
	)
	return
}
func BindBufferMemory(
	device VkDevice,
	buffer VkBuffer,
	memory VkDeviceMemory,
	memoryOffset VkDeviceSize,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkBindBufferMemory.Addr(),
		4,
		uintptr(device),
		uintptr(buffer),
		uintptr(memory),
		uintptr(memoryOffset),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func BindImageMemory(
	device VkDevice,
	image VkImage,
	memory VkDeviceMemory,
	memoryOffset VkDeviceSize,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkBindImageMemory.Addr(),
		4,
		uintptr(device),
		uintptr(image),
		uintptr(memory),
		uintptr(memoryOffset),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetBufferMemoryRequirements(
	device VkDevice,
	buffer VkBuffer,
	pMemoryRequirements *VkMemoryRequirements,
) {
	syscall.Syscall(
		PFN_vkGetBufferMemoryRequirements.Addr(),
		3,
		uintptr(device),
		uintptr(buffer),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetImageMemoryRequirements(
	device VkDevice,
	image VkImage,
	pMemoryRequirements *VkMemoryRequirements,
) {
	syscall.Syscall(
		PFN_vkGetImageMemoryRequirements.Addr(),
		3,
		uintptr(device),
		uintptr(image),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetImageSparseMemoryRequirements(
	device VkDevice,
	image VkImage,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements,
) {
	syscall.Syscall6(
		PFN_vkGetImageSparseMemoryRequirements.Addr(),
		4,
		uintptr(device),
		uintptr(image),
		uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)),
		uintptr(unsafe.Pointer(pSparseMemoryRequirements)),
		0,
		0,
	)
	return
}
func GetPhysicalDeviceSparseImageFormatProperties(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	vtype VkImageType,
	samples VkSampleCountFlagBits,
	usage VkImageUsageFlags,
	tiling VkImageTiling,
	pPropertyCount *uint32,
	pProperties *VkSparseImageFormatProperties,
) {
	syscall.Syscall9(
		PFN_vkGetPhysicalDeviceSparseImageFormatProperties.Addr(),
		8,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(vtype),
		uintptr(samples),
		uintptr(usage),
		uintptr(tiling),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
	)
	return
}
func QueueBindSparse(
	queue VkQueue,
	bindInfoCount uint32,
	pBindInfo *VkBindSparseInfo,
	fence VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkQueueBindSparse.Addr(),
		4,
		uintptr(queue),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfo)),
		uintptr(fence),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateFence(
	device VkDevice,
	pCreateInfo *VkFenceCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pFence VkFence, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateFence.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pFence)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyFence(
	device VkDevice,
	fence VkFence,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyFence.Addr(),
		3,
		uintptr(device),
		uintptr(fence),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func ResetFences(
	device VkDevice,
	fenceCount uint32,
	pFences *VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkResetFences.Addr(),
		3,
		uintptr(device),
		uintptr(fenceCount),
		uintptr(unsafe.Pointer(pFences)),
	)
	result = GetError((int32)(ret))
	return
}
func GetFenceStatus(
	device VkDevice,
	fence VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetFenceStatus.Addr(),
		2,
		uintptr(device),
		uintptr(fence),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func WaitForFences(
	device VkDevice,
	fenceCount uint32,
	pFences *VkFence,
	waitAll VkBool32,
	timeout uint64,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkWaitForFences.Addr(),
		5,
		uintptr(device),
		uintptr(fenceCount),
		uintptr(unsafe.Pointer(pFences)),
		uintptr(waitAll),
		uintptr(timeout),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateSemaphore(
	device VkDevice,
	pCreateInfo *VkSemaphoreCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pSemaphore VkSemaphore, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSemaphore.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSemaphore)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroySemaphore(
	device VkDevice,
	semaphore VkSemaphore,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySemaphore.Addr(),
		3,
		uintptr(device),
		uintptr(semaphore),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateEvent(
	device VkDevice,
	pCreateInfo *VkEventCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pEvent VkEvent, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateEvent.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pEvent)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyEvent(
	device VkDevice,
	event VkEvent,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyEvent.Addr(),
		3,
		uintptr(device),
		uintptr(event),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetEventStatus(
	device VkDevice,
	event VkEvent,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetEventStatus.Addr(),
		2,
		uintptr(device),
		uintptr(event),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func SetEvent(
	device VkDevice,
	event VkEvent,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkSetEvent.Addr(),
		2,
		uintptr(device),
		uintptr(event),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func ResetEvent(
	device VkDevice,
	event VkEvent,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkResetEvent.Addr(),
		2,
		uintptr(device),
		uintptr(event),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateQueryPool(
	device VkDevice,
	pCreateInfo *VkQueryPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pQueryPool VkQueryPool, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateQueryPool.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pQueryPool)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyQueryPool(
	device VkDevice,
	queryPool VkQueryPool,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyQueryPool.Addr(),
		3,
		uintptr(device),
		uintptr(queryPool),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetQueryPoolResults(
	device VkDevice,
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
	dataSize uint64,
	pData *interface{},
	stride VkDeviceSize,
	flags VkQueryResultFlags,
) (result error) {
	ret, _, _ := syscall.Syscall9(
		PFN_vkGetQueryPoolResults.Addr(),
		8,
		uintptr(device),
		uintptr(queryPool),
		uintptr(firstQuery),
		uintptr(queryCount),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
		uintptr(stride),
		uintptr(flags),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateBuffer(
	device VkDevice,
	pCreateInfo *VkBufferCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pBuffer VkBuffer, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateBuffer.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pBuffer)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyBuffer(
	device VkDevice,
	buffer VkBuffer,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyBuffer.Addr(),
		3,
		uintptr(device),
		uintptr(buffer),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateBufferView(
	device VkDevice,
	pCreateInfo *VkBufferViewCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pView VkBufferView, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateBufferView.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pView)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyBufferView(
	device VkDevice,
	bufferView VkBufferView,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyBufferView.Addr(),
		3,
		uintptr(device),
		uintptr(bufferView),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateImage(
	device VkDevice,
	pCreateInfo *VkImageCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pImage VkImage, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateImage.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pImage)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyImage(
	device VkDevice,
	image VkImage,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyImage.Addr(),
		3,
		uintptr(device),
		uintptr(image),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetImageSubresourceLayout(
	device VkDevice,
	image VkImage,
	pSubresource *VkImageSubresource,
	pLayout *VkSubresourceLayout,
) {
	syscall.Syscall6(
		PFN_vkGetImageSubresourceLayout.Addr(),
		4,
		uintptr(device),
		uintptr(image),
		uintptr(unsafe.Pointer(pSubresource)),
		uintptr(unsafe.Pointer(pLayout)),
		0,
		0,
	)
	return
}
func CreateImageView(
	device VkDevice,
	pCreateInfo *VkImageViewCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pView VkImageView, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateImageView.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pView)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyImageView(
	device VkDevice,
	imageView VkImageView,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyImageView.Addr(),
		3,
		uintptr(device),
		uintptr(imageView),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateShaderModule(
	device VkDevice,
	pCreateInfo *VkShaderModuleCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pShaderModule VkShaderModule, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateShaderModule.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pShaderModule)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyShaderModule(
	device VkDevice,
	shaderModule VkShaderModule,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyShaderModule.Addr(),
		3,
		uintptr(device),
		uintptr(shaderModule),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreatePipelineCache(
	device VkDevice,
	pCreateInfo *VkPipelineCacheCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pPipelineCache VkPipelineCache, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreatePipelineCache.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelineCache)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyPipelineCache(
	device VkDevice,
	pipelineCache VkPipelineCache,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyPipelineCache.Addr(),
		3,
		uintptr(device),
		uintptr(pipelineCache),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetPipelineCacheData(
	device VkDevice,
	pipelineCache VkPipelineCache,
	pDataSize *uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPipelineCacheData.Addr(),
		4,
		uintptr(device),
		uintptr(pipelineCache),
		uintptr(unsafe.Pointer(pDataSize)),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func MergePipelineCaches(
	device VkDevice,
	dstCache VkPipelineCache,
	srcCacheCount uint32,
	pSrcCaches *VkPipelineCache,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkMergePipelineCaches.Addr(),
		4,
		uintptr(device),
		uintptr(dstCache),
		uintptr(srcCacheCount),
		uintptr(unsafe.Pointer(pSrcCaches)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateGraphicsPipelines(
	device VkDevice,
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkGraphicsPipelineCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pPipelines VkPipeline, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateGraphicsPipelines.Addr(),
		6,
		uintptr(device),
		uintptr(pipelineCache),
		uintptr(createInfoCount),
		uintptr(unsafe.Pointer(pCreateInfos)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelines)),
	)
	result = GetError((int32)(ret))
	return
}
func CreateComputePipelines(
	device VkDevice,
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkComputePipelineCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pPipelines VkPipeline, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateComputePipelines.Addr(),
		6,
		uintptr(device),
		uintptr(pipelineCache),
		uintptr(createInfoCount),
		uintptr(unsafe.Pointer(pCreateInfos)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelines)),
	)
	result = GetError((int32)(ret))
	return
}
func DestroyPipeline(
	device VkDevice,
	pipeline VkPipeline,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyPipeline.Addr(),
		3,
		uintptr(device),
		uintptr(pipeline),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreatePipelineLayout(
	device VkDevice,
	pCreateInfo *VkPipelineLayoutCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pPipelineLayout VkPipelineLayout, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreatePipelineLayout.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelineLayout)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyPipelineLayout(
	device VkDevice,
	pipelineLayout VkPipelineLayout,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyPipelineLayout.Addr(),
		3,
		uintptr(device),
		uintptr(pipelineLayout),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateSampler(
	device VkDevice,
	pCreateInfo *VkSamplerCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pSampler VkSampler, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSampler.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSampler)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroySampler(
	device VkDevice,
	sampler VkSampler,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySampler.Addr(),
		3,
		uintptr(device),
		uintptr(sampler),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateDescriptorSetLayout(
	device VkDevice,
	pCreateInfo *VkDescriptorSetLayoutCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pSetLayout VkDescriptorSetLayout, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDescriptorSetLayout.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSetLayout)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDescriptorSetLayout(
	device VkDevice,
	descriptorSetLayout VkDescriptorSetLayout,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDescriptorSetLayout.Addr(),
		3,
		uintptr(device),
		uintptr(descriptorSetLayout),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateDescriptorPool(
	device VkDevice,
	pCreateInfo *VkDescriptorPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pDescriptorPool VkDescriptorPool, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDescriptorPool.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pDescriptorPool)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDescriptorPool(
	device VkDevice,
	descriptorPool VkDescriptorPool,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDescriptorPool.Addr(),
		3,
		uintptr(device),
		uintptr(descriptorPool),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func ResetDescriptorPool(
	device VkDevice,
	descriptorPool VkDescriptorPool,
	flags VkDescriptorPoolResetFlags,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkResetDescriptorPool.Addr(),
		3,
		uintptr(device),
		uintptr(descriptorPool),
		uintptr(flags),
	)
	result = GetError((int32)(ret))
	return
}
func AllocateDescriptorSets(
	device VkDevice,
	pAllocateInfo *VkDescriptorSetAllocateInfo,
	pDescriptorSets *VkDescriptorSet,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAllocateDescriptorSets.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pAllocateInfo)),
		uintptr(unsafe.Pointer(pDescriptorSets)),
	)
	result = GetError((int32)(ret))
	return
}
func FreeDescriptorSets(
	device VkDevice,
	descriptorPool VkDescriptorPool,
	descriptorSetCount uint32,
	pDescriptorSets *VkDescriptorSet,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkFreeDescriptorSets.Addr(),
		4,
		uintptr(device),
		uintptr(descriptorPool),
		uintptr(descriptorSetCount),
		uintptr(unsafe.Pointer(pDescriptorSets)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func UpdateDescriptorSets(
	device VkDevice,
	descriptorWriteCount uint32,
	pDescriptorWrites *VkWriteDescriptorSet,
	descriptorCopyCount uint32,
	pDescriptorCopies *VkCopyDescriptorSet,
) {
	syscall.Syscall6(
		PFN_vkUpdateDescriptorSets.Addr(),
		5,
		uintptr(device),
		uintptr(descriptorWriteCount),
		uintptr(unsafe.Pointer(pDescriptorWrites)),
		uintptr(descriptorCopyCount),
		uintptr(unsafe.Pointer(pDescriptorCopies)),
		0,
	)
	return
}
func CreateFramebuffer(
	device VkDevice,
	pCreateInfo *VkFramebufferCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pFramebuffer VkFramebuffer, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateFramebuffer.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pFramebuffer)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyFramebuffer(
	device VkDevice,
	framebuffer VkFramebuffer,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyFramebuffer.Addr(),
		3,
		uintptr(device),
		uintptr(framebuffer),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateRenderPass(
	device VkDevice,
	pCreateInfo *VkRenderPassCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pRenderPass VkRenderPass, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateRenderPass.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pRenderPass)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyRenderPass(
	device VkDevice,
	renderPass VkRenderPass,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyRenderPass.Addr(),
		3,
		uintptr(device),
		uintptr(renderPass),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetRenderAreaGranularity(
	device VkDevice,
	renderPass VkRenderPass,
	pGranularity *VkExtent2D,
) {
	syscall.Syscall(
		PFN_vkGetRenderAreaGranularity.Addr(),
		3,
		uintptr(device),
		uintptr(renderPass),
		uintptr(unsafe.Pointer(pGranularity)),
	)
	return
}
func CreateCommandPool(
	device VkDevice,
	pCreateInfo *VkCommandPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pCommandPool VkCommandPool, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateCommandPool.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pCommandPool)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyCommandPool(
	device VkDevice,
	commandPool VkCommandPool,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyCommandPool.Addr(),
		3,
		uintptr(device),
		uintptr(commandPool),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func ResetCommandPool(
	device VkDevice,
	commandPool VkCommandPool,
	flags VkCommandPoolResetFlags,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkResetCommandPool.Addr(),
		3,
		uintptr(device),
		uintptr(commandPool),
		uintptr(flags),
	)
	result = GetError((int32)(ret))
	return
}
func AllocateCommandBuffers(
	device VkDevice,
	pAllocateInfo *VkCommandBufferAllocateInfo,
	pCommandBuffers *VkCommandBuffer,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAllocateCommandBuffers.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pAllocateInfo)),
		uintptr(unsafe.Pointer(pCommandBuffers)),
	)
	result = GetError((int32)(ret))
	return
}
func FreeCommandBuffers(
	device VkDevice,
	commandPool VkCommandPool,
	commandBufferCount uint32,
	pCommandBuffers *VkCommandBuffer,
) {
	syscall.Syscall6(
		PFN_vkFreeCommandBuffers.Addr(),
		4,
		uintptr(device),
		uintptr(commandPool),
		uintptr(commandBufferCount),
		uintptr(unsafe.Pointer(pCommandBuffers)),
		0,
		0,
	)
	return
}
func BeginCommandBuffer(
	commandBuffer VkCommandBuffer,
	pBeginInfo *VkCommandBufferBeginInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBeginCommandBuffer.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pBeginInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func EndCommandBuffer(
	commandBuffer VkCommandBuffer,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEndCommandBuffer.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func ResetCommandBuffer(
	commandBuffer VkCommandBuffer,
	flags VkCommandBufferResetFlags,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkResetCommandBuffer.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(flags),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdBindPipeline(
	commandBuffer VkCommandBuffer,
	pipelineBindPoint VkPipelineBindPoint,
	pipeline VkPipeline,
) {
	syscall.Syscall(
		PFN_vkCmdBindPipeline.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(pipelineBindPoint),
		uintptr(pipeline),
	)
	return
}
func CmdSetViewport(
	commandBuffer VkCommandBuffer,
	firstViewport uint32,
	viewportCount uint32,
	pViewports *VkViewport,
) {
	syscall.Syscall6(
		PFN_vkCmdSetViewport.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstViewport),
		uintptr(viewportCount),
		uintptr(unsafe.Pointer(pViewports)),
		0,
		0,
	)
	return
}
func CmdSetScissor(
	commandBuffer VkCommandBuffer,
	firstScissor uint32,
	scissorCount uint32,
	pScissors *VkRect2D,
) {
	syscall.Syscall6(
		PFN_vkCmdSetScissor.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstScissor),
		uintptr(scissorCount),
		uintptr(unsafe.Pointer(pScissors)),
		0,
		0,
	)
	return
}
func CmdSetLineWidth(
	commandBuffer VkCommandBuffer,
	lineWidth float32,
) {
	syscall.Syscall(
		PFN_vkCmdSetLineWidth.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(lineWidth),
		0,
	)
	return
}
func CmdSetDepthBias(
	commandBuffer VkCommandBuffer,
	depthBiasConstantFactor float32,
	depthBiasClamp float32,
	depthBiasSlopeFactor float32,
) {
	syscall.Syscall6(
		PFN_vkCmdSetDepthBias.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(depthBiasConstantFactor),
		uintptr(depthBiasClamp),
		uintptr(depthBiasSlopeFactor),
		0,
		0,
	)
	return
}
func CmdSetBlendConstants(
	commandBuffer VkCommandBuffer,
	blendConstants [4]float32,
) {
	syscall.Syscall(
		PFN_vkCmdSetBlendConstants.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(&blendConstants[0])),
		0,
	)
	return
}
func CmdSetDepthBounds(
	commandBuffer VkCommandBuffer,
	minDepthBounds float32,
	maxDepthBounds float32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthBounds.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(minDepthBounds),
		uintptr(maxDepthBounds),
	)
	return
}
func CmdSetStencilCompareMask(
	commandBuffer VkCommandBuffer,
	faceMask VkStencilFaceFlags,
	compareMask uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetStencilCompareMask.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(faceMask),
		uintptr(compareMask),
	)
	return
}
func CmdSetStencilWriteMask(
	commandBuffer VkCommandBuffer,
	faceMask VkStencilFaceFlags,
	writeMask uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetStencilWriteMask.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(faceMask),
		uintptr(writeMask),
	)
	return
}
func CmdSetStencilReference(
	commandBuffer VkCommandBuffer,
	faceMask VkStencilFaceFlags,
	reference uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetStencilReference.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(faceMask),
		uintptr(reference),
	)
	return
}
func CmdBindDescriptorSets(
	commandBuffer VkCommandBuffer,
	pipelineBindPoint VkPipelineBindPoint,
	layout VkPipelineLayout,
	firstSet uint32,
	descriptorSetCount uint32,
	pDescriptorSets *VkDescriptorSet,
	dynamicOffsetCount uint32,
	pDynamicOffsets *uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdBindDescriptorSets.Addr(),
		8,
		uintptr(commandBuffer),
		uintptr(pipelineBindPoint),
		uintptr(layout),
		uintptr(firstSet),
		uintptr(descriptorSetCount),
		uintptr(unsafe.Pointer(pDescriptorSets)),
		uintptr(dynamicOffsetCount),
		uintptr(unsafe.Pointer(pDynamicOffsets)),
		0,
	)
	return
}
func CmdBindIndexBuffer(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	indexType VkIndexType,
) {
	syscall.Syscall6(
		PFN_vkCmdBindIndexBuffer.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(indexType),
		0,
		0,
	)
	return
}
func CmdBindVertexBuffers(
	commandBuffer VkCommandBuffer,
	firstBinding uint32,
	bindingCount uint32,
	pBuffers *VkBuffer,
	pOffsets *VkDeviceSize,
) {
	syscall.Syscall6(
		PFN_vkCmdBindVertexBuffers.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(firstBinding),
		uintptr(bindingCount),
		uintptr(unsafe.Pointer(pBuffers)),
		uintptr(unsafe.Pointer(pOffsets)),
		0,
	)
	return
}
func CmdDraw(
	commandBuffer VkCommandBuffer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDraw.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(vertexCount),
		uintptr(instanceCount),
		uintptr(firstVertex),
		uintptr(firstInstance),
		0,
	)
	return
}
func CmdDrawIndexed(
	commandBuffer VkCommandBuffer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	vertexOffset int32,
	firstInstance uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDrawIndexed.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(indexCount),
		uintptr(instanceCount),
		uintptr(firstIndex),
		uintptr(vertexOffset),
		uintptr(firstInstance),
	)
	return
}
func CmdDrawIndirect(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	drawCount uint32,
	stride uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDrawIndirect.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(drawCount),
		uintptr(stride),
		0,
	)
	return
}
func CmdDrawIndexedIndirect(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	drawCount uint32,
	stride uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDrawIndexedIndirect.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(drawCount),
		uintptr(stride),
		0,
	)
	return
}
func CmdDispatch(
	commandBuffer VkCommandBuffer,
	groupCountX uint32,
	groupCountY uint32,
	groupCountZ uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDispatch.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(groupCountX),
		uintptr(groupCountY),
		uintptr(groupCountZ),
		0,
		0,
	)
	return
}
func CmdDispatchIndirect(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
) {
	syscall.Syscall(
		PFN_vkCmdDispatchIndirect.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
	)
	return
}
func CmdCopyBuffer(
	commandBuffer VkCommandBuffer,
	srcBuffer VkBuffer,
	dstBuffer VkBuffer,
	regionCount uint32,
	pRegions *VkBufferCopy,
) {
	syscall.Syscall6(
		PFN_vkCmdCopyBuffer.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(srcBuffer),
		uintptr(dstBuffer),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
		0,
	)
	return
}
func CmdCopyImage(
	commandBuffer VkCommandBuffer,
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageCopy,
) {
	syscall.Syscall9(
		PFN_vkCmdCopyImage.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(srcImage),
		uintptr(srcImageLayout),
		uintptr(dstImage),
		uintptr(dstImageLayout),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
		0,
		0,
	)
	return
}
func CmdBlitImage(
	commandBuffer VkCommandBuffer,
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageBlit,
	filter VkFilter,
) {
	syscall.Syscall9(
		PFN_vkCmdBlitImage.Addr(),
		8,
		uintptr(commandBuffer),
		uintptr(srcImage),
		uintptr(srcImageLayout),
		uintptr(dstImage),
		uintptr(dstImageLayout),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
		uintptr(filter),
		0,
	)
	return
}
func CmdCopyBufferToImage(
	commandBuffer VkCommandBuffer,
	srcBuffer VkBuffer,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkBufferImageCopy,
) {
	syscall.Syscall6(
		PFN_vkCmdCopyBufferToImage.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(srcBuffer),
		uintptr(dstImage),
		uintptr(dstImageLayout),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
	)
	return
}
func CmdCopyImageToBuffer(
	commandBuffer VkCommandBuffer,
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstBuffer VkBuffer,
	regionCount uint32,
	pRegions *VkBufferImageCopy,
) {
	syscall.Syscall6(
		PFN_vkCmdCopyImageToBuffer.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(srcImage),
		uintptr(srcImageLayout),
		uintptr(dstBuffer),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
	)
	return
}
func CmdUpdateBuffer(
	commandBuffer VkCommandBuffer,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	dataSize VkDeviceSize,
	pData *interface{},
) {
	syscall.Syscall6(
		PFN_vkCmdUpdateBuffer.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(dstBuffer),
		uintptr(dstOffset),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
		0,
	)
	return
}
func CmdFillBuffer(
	commandBuffer VkCommandBuffer,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	size VkDeviceSize,
	data uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdFillBuffer.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(dstBuffer),
		uintptr(dstOffset),
		uintptr(size),
		uintptr(data),
		0,
	)
	return
}
func CmdClearColorImage(
	commandBuffer VkCommandBuffer,
	image VkImage,
	imageLayout VkImageLayout,
	pColor *VkClearColorValue,
	rangeCount uint32,
	pRanges *VkImageSubresourceRange,
) {
	syscall.Syscall6(
		PFN_vkCmdClearColorImage.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(image),
		uintptr(imageLayout),
		uintptr(unsafe.Pointer(pColor)),
		uintptr(rangeCount),
		uintptr(unsafe.Pointer(pRanges)),
	)
	return
}
func CmdClearDepthStencilImage(
	commandBuffer VkCommandBuffer,
	image VkImage,
	imageLayout VkImageLayout,
	pDepthStencil *VkClearDepthStencilValue,
	rangeCount uint32,
	pRanges *VkImageSubresourceRange,
) {
	syscall.Syscall6(
		PFN_vkCmdClearDepthStencilImage.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(image),
		uintptr(imageLayout),
		uintptr(unsafe.Pointer(pDepthStencil)),
		uintptr(rangeCount),
		uintptr(unsafe.Pointer(pRanges)),
	)
	return
}
func CmdClearAttachments(
	commandBuffer VkCommandBuffer,
	attachmentCount uint32,
	pAttachments *VkClearAttachment,
	rectCount uint32,
	pRects *VkClearRect,
) {
	syscall.Syscall6(
		PFN_vkCmdClearAttachments.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(attachmentCount),
		uintptr(unsafe.Pointer(pAttachments)),
		uintptr(rectCount),
		uintptr(unsafe.Pointer(pRects)),
		0,
	)
	return
}
func CmdResolveImage(
	commandBuffer VkCommandBuffer,
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageResolve,
) {
	syscall.Syscall9(
		PFN_vkCmdResolveImage.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(srcImage),
		uintptr(srcImageLayout),
		uintptr(dstImage),
		uintptr(dstImageLayout),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(pRegions)),
		0,
		0,
	)
	return
}
func CmdSetEvent(
	commandBuffer VkCommandBuffer,
	event VkEvent,
	stageMask VkPipelineStageFlags,
) {
	syscall.Syscall(
		PFN_vkCmdSetEvent.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(event),
		uintptr(stageMask),
	)
	return
}
func CmdResetEvent(
	commandBuffer VkCommandBuffer,
	event VkEvent,
	stageMask VkPipelineStageFlags,
) {
	syscall.Syscall(
		PFN_vkCmdResetEvent.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(event),
		uintptr(stageMask),
	)
	return
}
func CmdWaitEvents(
	commandBuffer VkCommandBuffer,
	eventCount uint32,
	pEvents *VkEvent,
	srcStageMask VkPipelineStageFlags,
	dstStageMask VkPipelineStageFlags,
	memoryBarrierCount uint32,
	pMemoryBarriers *VkMemoryBarrier,
	bufferMemoryBarrierCount uint32,
	pBufferMemoryBarriers *VkBufferMemoryBarrier,
	imageMemoryBarrierCount uint32,
	pImageMemoryBarriers *VkImageMemoryBarrier,
) {
	syscall.Syscall12(
		PFN_vkCmdWaitEvents.Addr(),
		11,
		uintptr(commandBuffer),
		uintptr(eventCount),
		uintptr(unsafe.Pointer(pEvents)),
		uintptr(srcStageMask),
		uintptr(dstStageMask),
		uintptr(memoryBarrierCount),
		uintptr(unsafe.Pointer(pMemoryBarriers)),
		uintptr(bufferMemoryBarrierCount),
		uintptr(unsafe.Pointer(pBufferMemoryBarriers)),
		uintptr(imageMemoryBarrierCount),
		uintptr(unsafe.Pointer(pImageMemoryBarriers)),
		0,
	)
	return
}
func CmdPipelineBarrier(
	commandBuffer VkCommandBuffer,
	srcStageMask VkPipelineStageFlags,
	dstStageMask VkPipelineStageFlags,
	dependencyFlags VkDependencyFlags,
	memoryBarrierCount uint32,
	pMemoryBarriers *VkMemoryBarrier,
	bufferMemoryBarrierCount uint32,
	pBufferMemoryBarriers *VkBufferMemoryBarrier,
	imageMemoryBarrierCount uint32,
	pImageMemoryBarriers *VkImageMemoryBarrier,
) {
	syscall.Syscall12(
		PFN_vkCmdPipelineBarrier.Addr(),
		10,
		uintptr(commandBuffer),
		uintptr(srcStageMask),
		uintptr(dstStageMask),
		uintptr(dependencyFlags),
		uintptr(memoryBarrierCount),
		uintptr(unsafe.Pointer(pMemoryBarriers)),
		uintptr(bufferMemoryBarrierCount),
		uintptr(unsafe.Pointer(pBufferMemoryBarriers)),
		uintptr(imageMemoryBarrierCount),
		uintptr(unsafe.Pointer(pImageMemoryBarriers)),
		0,
		0,
	)
	return
}
func CmdBeginQuery(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	query uint32,
	flags VkQueryControlFlags,
) {
	syscall.Syscall6(
		PFN_vkCmdBeginQuery.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(query),
		uintptr(flags),
		0,
		0,
	)
	return
}
func CmdEndQuery(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	query uint32,
) {
	syscall.Syscall(
		PFN_vkCmdEndQuery.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(query),
	)
	return
}
func CmdResetQueryPool(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdResetQueryPool.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(firstQuery),
		uintptr(queryCount),
		0,
		0,
	)
	return
}
func CmdWriteTimestamp(
	commandBuffer VkCommandBuffer,
	pipelineStage VkPipelineStageFlagBits,
	queryPool VkQueryPool,
	query uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteTimestamp.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(pipelineStage),
		uintptr(queryPool),
		uintptr(query),
		0,
		0,
	)
	return
}
func CmdCopyQueryPoolResults(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	stride VkDeviceSize,
	flags VkQueryResultFlags,
) {
	syscall.Syscall9(
		PFN_vkCmdCopyQueryPoolResults.Addr(),
		8,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(firstQuery),
		uintptr(queryCount),
		uintptr(dstBuffer),
		uintptr(dstOffset),
		uintptr(stride),
		uintptr(flags),
		0,
	)
	return
}
func CmdPushConstants(
	commandBuffer VkCommandBuffer,
	layout VkPipelineLayout,
	stageFlags VkShaderStageFlags,
	offset uint32,
	size uint32,
	pValues *interface{},
) {
	syscall.Syscall6(
		PFN_vkCmdPushConstants.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(layout),
		uintptr(stageFlags),
		uintptr(offset),
		uintptr(size),
		uintptr(unsafe.Pointer(pValues)),
	)
	return
}
func CmdBeginRenderPass(
	commandBuffer VkCommandBuffer,
	pRenderPassBegin *VkRenderPassBeginInfo,
	contents VkSubpassContents,
) {
	syscall.Syscall(
		PFN_vkCmdBeginRenderPass.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pRenderPassBegin)),
		uintptr(contents),
	)
	return
}
func CmdNextSubpass(
	commandBuffer VkCommandBuffer,
	contents VkSubpassContents,
) {
	syscall.Syscall(
		PFN_vkCmdNextSubpass.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(contents),
		0,
	)
	return
}
func CmdEndRenderPass(
	commandBuffer VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdEndRenderPass.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	return
}
func CmdExecuteCommands(
	commandBuffer VkCommandBuffer,
	commandBufferCount uint32,
	pCommandBuffers *VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdExecuteCommands.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(commandBufferCount),
		uintptr(unsafe.Pointer(pCommandBuffers)),
	)
	return
}
func EnumerateInstanceVersion(
	pApiVersion *uint32,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumerateInstanceVersion.Addr(),
		1,
		uintptr(unsafe.Pointer(pApiVersion)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func BindBufferMemory2(
	device VkDevice,
	bindInfoCount uint32,
	pBindInfos *VkBindBufferMemoryInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBindBufferMemory2.Addr(),
		3,
		uintptr(device),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfos)),
	)
	result = GetError((int32)(ret))
	return
}
func BindImageMemory2(
	device VkDevice,
	bindInfoCount uint32,
	pBindInfos *VkBindImageMemoryInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBindImageMemory2.Addr(),
		3,
		uintptr(device),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfos)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceGroupPeerMemoryFeatures(
	device VkDevice,
	heapIndex uint32,
	localDeviceIndex uint32,
	remoteDeviceIndex uint32,
	pPeerMemoryFeatures *VkPeerMemoryFeatureFlags,
) {
	syscall.Syscall6(
		PFN_vkGetDeviceGroupPeerMemoryFeatures.Addr(),
		5,
		uintptr(device),
		uintptr(heapIndex),
		uintptr(localDeviceIndex),
		uintptr(remoteDeviceIndex),
		uintptr(unsafe.Pointer(pPeerMemoryFeatures)),
		0,
	)
	return
}
func CmdSetDeviceMask(
	commandBuffer VkCommandBuffer,
	deviceMask uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDeviceMask.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(deviceMask),
		0,
	)
	return
}
func CmdDispatchBase(
	commandBuffer VkCommandBuffer,
	baseGroupX uint32,
	baseGroupY uint32,
	baseGroupZ uint32,
	groupCountX uint32,
	groupCountY uint32,
	groupCountZ uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDispatchBase.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(baseGroupX),
		uintptr(baseGroupY),
		uintptr(baseGroupZ),
		uintptr(groupCountX),
		uintptr(groupCountY),
		uintptr(groupCountZ),
		0,
		0,
	)
	return
}
func EnumeratePhysicalDeviceGroups(
	instance VkInstance,
	pPhysicalDeviceGroupCount *uint32,
	pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumeratePhysicalDeviceGroups.Addr(),
		3,
		uintptr(instance),
		uintptr(unsafe.Pointer(pPhysicalDeviceGroupCount)),
		uintptr(unsafe.Pointer(pPhysicalDeviceGroupProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetImageMemoryRequirements2(
	device VkDevice,
	pInfo *VkImageMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
	syscall.Syscall(
		PFN_vkGetImageMemoryRequirements2.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetBufferMemoryRequirements2(
	device VkDevice,
	pInfo *VkBufferMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
	syscall.Syscall(
		PFN_vkGetBufferMemoryRequirements2.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetImageSparseMemoryRequirements2(
	device VkDevice,
	pInfo *VkImageSparseMemoryRequirementsInfo2,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements2,
) {
	syscall.Syscall6(
		PFN_vkGetImageSparseMemoryRequirements2.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)),
		uintptr(unsafe.Pointer(pSparseMemoryRequirements)),
		0,
		0,
	)
	return
}
func GetPhysicalDeviceFeatures2(
	physicalDevice VkPhysicalDevice,
	pFeatures *VkPhysicalDeviceFeatures2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFeatures2.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFeatures)),
		0,
	)
	return
}
func GetPhysicalDeviceProperties2(
	physicalDevice VkPhysicalDevice,
	pProperties *VkPhysicalDeviceProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceProperties2.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pProperties)),
		0,
	)
	return
}
func GetPhysicalDeviceFormatProperties2(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	pFormatProperties *VkFormatProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFormatProperties2.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(unsafe.Pointer(pFormatProperties)),
	)
	return
}
func GetPhysicalDeviceImageFormatProperties2(
	physicalDevice VkPhysicalDevice,
	pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2,
	pImageFormatProperties *VkImageFormatProperties2,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceImageFormatProperties2.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pImageFormatInfo)),
		uintptr(unsafe.Pointer(pImageFormatProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceQueueFamilyProperties2(
	physicalDevice VkPhysicalDevice,
	pQueueFamilyPropertyCount *uint32,
	pQueueFamilyProperties *VkQueueFamilyProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceQueueFamilyProperties2.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)),
		uintptr(unsafe.Pointer(pQueueFamilyProperties)),
	)
	return
}
func GetPhysicalDeviceMemoryProperties2(
	physicalDevice VkPhysicalDevice,
	pMemoryProperties *VkPhysicalDeviceMemoryProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceMemoryProperties2.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pMemoryProperties)),
		0,
	)
	return
}
func GetPhysicalDeviceSparseImageFormatProperties2(
	physicalDevice VkPhysicalDevice,
	pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2,
	pPropertyCount *uint32,
	pProperties *VkSparseImageFormatProperties2,
) {
	syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSparseImageFormatProperties2.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFormatInfo)),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	return
}
func TrimCommandPool(
	device VkDevice,
	commandPool VkCommandPool,
	flags VkCommandPoolTrimFlags,
) {
	syscall.Syscall(
		PFN_vkTrimCommandPool.Addr(),
		3,
		uintptr(device),
		uintptr(commandPool),
		uintptr(flags),
	)
	return
}
func GetDeviceQueue2(
	device VkDevice,
	pQueueInfo *VkDeviceQueueInfo2,
	pQueue *VkQueue,
) {
	syscall.Syscall(
		PFN_vkGetDeviceQueue2.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pQueueInfo)),
		uintptr(unsafe.Pointer(pQueue)),
	)
	return
}
func CreateSamplerYcbcrConversion(
	device VkDevice,
	pCreateInfo *VkSamplerYcbcrConversionCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pYcbcrConversion VkSamplerYcbcrConversion, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSamplerYcbcrConversion.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pYcbcrConversion)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroySamplerYcbcrConversion(
	device VkDevice,
	ycbcrConversion VkSamplerYcbcrConversion,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySamplerYcbcrConversion.Addr(),
		3,
		uintptr(device),
		uintptr(ycbcrConversion),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CreateDescriptorUpdateTemplate(
	device VkDevice,
	pCreateInfo *VkDescriptorUpdateTemplateCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pDescriptorUpdateTemplate VkDescriptorUpdateTemplate, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDescriptorUpdateTemplate.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pDescriptorUpdateTemplate)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDescriptorUpdateTemplate(
	device VkDevice,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDescriptorUpdateTemplate.Addr(),
		3,
		uintptr(device),
		uintptr(descriptorUpdateTemplate),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func UpdateDescriptorSetWithTemplate(
	device VkDevice,
	descriptorSet VkDescriptorSet,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pData *interface{},
) {
	syscall.Syscall6(
		PFN_vkUpdateDescriptorSetWithTemplate.Addr(),
		4,
		uintptr(device),
		uintptr(descriptorSet),
		uintptr(descriptorUpdateTemplate),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	return
}
func GetPhysicalDeviceExternalBufferProperties(
	physicalDevice VkPhysicalDevice,
	pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo,
	pExternalBufferProperties *VkExternalBufferProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalBufferProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalBufferInfo)),
		uintptr(unsafe.Pointer(pExternalBufferProperties)),
	)
	return
}
func GetPhysicalDeviceExternalFenceProperties(
	physicalDevice VkPhysicalDevice,
	pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo,
	pExternalFenceProperties *VkExternalFenceProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalFenceProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalFenceInfo)),
		uintptr(unsafe.Pointer(pExternalFenceProperties)),
	)
	return
}
func GetPhysicalDeviceExternalSemaphoreProperties(
	physicalDevice VkPhysicalDevice,
	pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo,
	pExternalSemaphoreProperties *VkExternalSemaphoreProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalSemaphoreProperties.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalSemaphoreInfo)),
		uintptr(unsafe.Pointer(pExternalSemaphoreProperties)),
	)
	return
}
func GetDescriptorSetLayoutSupport(
	device VkDevice,
	pCreateInfo *VkDescriptorSetLayoutCreateInfo,
	pSupport *VkDescriptorSetLayoutSupport,
) {
	syscall.Syscall(
		PFN_vkGetDescriptorSetLayoutSupport.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pSupport)),
	)
	return
}
func CmdDrawIndirectCount(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndirectCount.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func CmdDrawIndexedIndirectCount(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndexedIndirectCount.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func CreateRenderPass2(
	device VkDevice,
	pCreateInfo *VkRenderPassCreateInfo2,
	pAllocator *VkAllocationCallbacks,
) (pRenderPass VkRenderPass, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateRenderPass2.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pRenderPass)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdBeginRenderPass2(
	commandBuffer VkCommandBuffer,
	pRenderPassBegin *VkRenderPassBeginInfo,
	pSubpassBeginInfo *VkSubpassBeginInfo,
) {
	syscall.Syscall(
		PFN_vkCmdBeginRenderPass2.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pRenderPassBegin)),
		uintptr(unsafe.Pointer(pSubpassBeginInfo)),
	)
	return
}
func CmdNextSubpass2(
	commandBuffer VkCommandBuffer,
	pSubpassBeginInfo *VkSubpassBeginInfo,
	pSubpassEndInfo *VkSubpassEndInfo,
) {
	syscall.Syscall(
		PFN_vkCmdNextSubpass2.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pSubpassBeginInfo)),
		uintptr(unsafe.Pointer(pSubpassEndInfo)),
	)
	return
}
func CmdEndRenderPass2(
	commandBuffer VkCommandBuffer,
	pSubpassEndInfo *VkSubpassEndInfo,
) {
	syscall.Syscall(
		PFN_vkCmdEndRenderPass2.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pSubpassEndInfo)),
		0,
	)
	return
}
func ResetQueryPool(
	device VkDevice,
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
) {
	syscall.Syscall6(
		PFN_vkResetQueryPool.Addr(),
		4,
		uintptr(device),
		uintptr(queryPool),
		uintptr(firstQuery),
		uintptr(queryCount),
		0,
		0,
	)
	return
}
func GetSemaphoreCounterValue(
	device VkDevice,
	semaphore VkSemaphore,
	pValue *uint64,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetSemaphoreCounterValue.Addr(),
		3,
		uintptr(device),
		uintptr(semaphore),
		uintptr(unsafe.Pointer(pValue)),
	)
	result = GetError((int32)(ret))
	return
}
func WaitSemaphores(
	device VkDevice,
	pWaitInfo *VkSemaphoreWaitInfo,
	timeout uint64,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkWaitSemaphores.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pWaitInfo)),
		uintptr(timeout),
	)
	result = GetError((int32)(ret))
	return
}
func SignalSemaphore(
	device VkDevice,
	pSignalInfo *VkSemaphoreSignalInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkSignalSemaphore.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pSignalInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetBufferDeviceAddress(
	device VkDevice,
	pInfo *VkBufferDeviceAddressInfo,
) (result VkDeviceAddress) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetBufferDeviceAddress.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (VkDeviceAddress)(ret)
	return
}
func GetBufferOpaqueCaptureAddress(
	device VkDevice,
	pInfo *VkBufferDeviceAddressInfo,
) (result uint64) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetBufferOpaqueCaptureAddress.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (uint64)(ret)
	return
}
func GetDeviceMemoryOpaqueCaptureAddress(
	device VkDevice,
	pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo,
) (result uint64) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceMemoryOpaqueCaptureAddress.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (uint64)(ret)
	return
}
func DestroySurfaceKHR(
	instance VkInstance,
	surface VkSurfaceKHR,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySurfaceKHR.Addr(),
		3,
		uintptr(instance),
		uintptr(surface),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetPhysicalDeviceSurfaceSupportKHR(
	physicalDevice VkPhysicalDevice,
	queueFamilyIndex uint32,
	surface VkSurfaceKHR,
	pSupported *VkBool32,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSurfaceSupportKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(queueFamilyIndex),
		uintptr(surface),
		uintptr(unsafe.Pointer(pSupported)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfaceCapabilitiesKHR(
	physicalDevice VkPhysicalDevice,
	surface VkSurfaceKHR,
	pSurfaceCapabilities *VkSurfaceCapabilitiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(surface),
		uintptr(unsafe.Pointer(pSurfaceCapabilities)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfaceFormatsKHR(
	physicalDevice VkPhysicalDevice,
	surface VkSurfaceKHR,
	pSurfaceFormatCount *uint32,
	pSurfaceFormats *VkSurfaceFormatKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSurfaceFormatsKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(surface),
		uintptr(unsafe.Pointer(pSurfaceFormatCount)),
		uintptr(unsafe.Pointer(pSurfaceFormats)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfacePresentModesKHR(
	physicalDevice VkPhysicalDevice,
	surface VkSurfaceKHR,
	pPresentModeCount *uint32,
	pPresentModes *VkPresentModeKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSurfacePresentModesKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(surface),
		uintptr(unsafe.Pointer(pPresentModeCount)),
		uintptr(unsafe.Pointer(pPresentModes)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateSwapchainKHR(
	device VkDevice,
	pCreateInfo *VkSwapchainCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pSwapchain VkSwapchainKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSwapchainKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSwapchain)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroySwapchainKHR(
	device VkDevice,
	swapchain VkSwapchainKHR,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySwapchainKHR.Addr(),
		3,
		uintptr(device),
		uintptr(swapchain),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetSwapchainImagesKHR(
	device VkDevice,
	swapchain VkSwapchainKHR,
	pSwapchainImageCount *uint32,
	pSwapchainImages *VkImage,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetSwapchainImagesKHR.Addr(),
		4,
		uintptr(device),
		uintptr(swapchain),
		uintptr(unsafe.Pointer(pSwapchainImageCount)),
		uintptr(unsafe.Pointer(pSwapchainImages)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func AcquireNextImageKHR(
	device VkDevice,
	swapchain VkSwapchainKHR,
	timeout uint64,
	semaphore VkSemaphore,
	fence VkFence,
	pImageIndex *uint32,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkAcquireNextImageKHR.Addr(),
		6,
		uintptr(device),
		uintptr(swapchain),
		uintptr(timeout),
		uintptr(semaphore),
		uintptr(fence),
		uintptr(unsafe.Pointer(pImageIndex)),
	)
	result = GetError((int32)(ret))
	return
}
func QueuePresentKHR(
	queue VkQueue,
	pPresentInfo *VkPresentInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkQueuePresentKHR.Addr(),
		2,
		uintptr(queue),
		uintptr(unsafe.Pointer(pPresentInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceGroupPresentCapabilitiesKHR(
	device VkDevice,
	pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceGroupPresentCapabilitiesKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pDeviceGroupPresentCapabilities)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDeviceGroupSurfacePresentModesKHR(
	device VkDevice,
	surface VkSurfaceKHR,
	pModes *VkDeviceGroupPresentModeFlagsKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceGroupSurfacePresentModesKHR.Addr(),
		3,
		uintptr(device),
		uintptr(surface),
		uintptr(unsafe.Pointer(pModes)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDevicePresentRectanglesKHR(
	physicalDevice VkPhysicalDevice,
	surface VkSurfaceKHR,
	pRectCount *uint32,
	pRects *VkRect2D,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDevicePresentRectanglesKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(surface),
		uintptr(unsafe.Pointer(pRectCount)),
		uintptr(unsafe.Pointer(pRects)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func AcquireNextImage2KHR(
	device VkDevice,
	pAcquireInfo *VkAcquireNextImageInfoKHR,
	pImageIndex *uint32,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquireNextImage2KHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pAcquireInfo)),
		uintptr(unsafe.Pointer(pImageIndex)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceDisplayPropertiesKHR(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkDisplayPropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceDisplayPropertiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceDisplayPlanePropertiesKHR(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkDisplayPlanePropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDisplayPlaneSupportedDisplaysKHR(
	physicalDevice VkPhysicalDevice,
	planeIndex uint32,
	pDisplayCount *uint32,
	pDisplays *VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetDisplayPlaneSupportedDisplaysKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(planeIndex),
		uintptr(unsafe.Pointer(pDisplayCount)),
		uintptr(unsafe.Pointer(pDisplays)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDisplayModePropertiesKHR(
	physicalDevice VkPhysicalDevice,
	display VkDisplayKHR,
	pPropertyCount *uint32,
	pProperties *VkDisplayModePropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetDisplayModePropertiesKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(display),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateDisplayModeKHR(
	physicalDevice VkPhysicalDevice,
	display VkDisplayKHR,
	pCreateInfo *VkDisplayModeCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pMode VkDisplayModeKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDisplayModeKHR.Addr(),
		5,
		uintptr(physicalDevice),
		uintptr(display),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pMode)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDisplayPlaneCapabilitiesKHR(
	physicalDevice VkPhysicalDevice,
	mode VkDisplayModeKHR,
	planeIndex uint32,
	pCapabilities *VkDisplayPlaneCapabilitiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetDisplayPlaneCapabilitiesKHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(mode),
		uintptr(planeIndex),
		uintptr(unsafe.Pointer(pCapabilities)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateDisplayPlaneSurfaceKHR(
	instance VkInstance,
	pCreateInfo *VkDisplaySurfaceCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pSurface VkSurfaceKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDisplayPlaneSurfaceKHR.Addr(),
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
func CreateSharedSwapchainsKHR(
	device VkDevice,
	swapchainCount uint32,
	pCreateInfos *VkSwapchainCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pSwapchains VkSwapchainKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSharedSwapchainsKHR.Addr(),
		5,
		uintptr(device),
		uintptr(swapchainCount),
		uintptr(unsafe.Pointer(pCreateInfos)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pSwapchains)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceFeatures2KHR(
	physicalDevice VkPhysicalDevice,
	pFeatures *VkPhysicalDeviceFeatures2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFeatures2KHR.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFeatures)),
		0,
	)
	return
}
func GetPhysicalDeviceProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pProperties *VkPhysicalDeviceProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceProperties2KHR.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pProperties)),
		0,
	)
	return
}
func GetPhysicalDeviceFormatProperties2KHR(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	pFormatProperties *VkFormatProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceFormatProperties2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(unsafe.Pointer(pFormatProperties)),
	)
	return
}
func GetPhysicalDeviceImageFormatProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2,
	pImageFormatProperties *VkImageFormatProperties2,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceImageFormatProperties2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pImageFormatInfo)),
		uintptr(unsafe.Pointer(pImageFormatProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceQueueFamilyProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pQueueFamilyPropertyCount *uint32,
	pQueueFamilyProperties *VkQueueFamilyProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)),
		uintptr(unsafe.Pointer(pQueueFamilyProperties)),
	)
	return
}
func GetPhysicalDeviceMemoryProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pMemoryProperties *VkPhysicalDeviceMemoryProperties2,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceMemoryProperties2KHR.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pMemoryProperties)),
		0,
	)
	return
}
func GetPhysicalDeviceSparseImageFormatProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2,
	pPropertyCount *uint32,
	pProperties *VkSparseImageFormatProperties2,
) {
	syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFormatInfo)),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	return
}
func GetDeviceGroupPeerMemoryFeaturesKHR(
	device VkDevice,
	heapIndex uint32,
	localDeviceIndex uint32,
	remoteDeviceIndex uint32,
	pPeerMemoryFeatures *VkPeerMemoryFeatureFlags,
) {
	syscall.Syscall6(
		PFN_vkGetDeviceGroupPeerMemoryFeaturesKHR.Addr(),
		5,
		uintptr(device),
		uintptr(heapIndex),
		uintptr(localDeviceIndex),
		uintptr(remoteDeviceIndex),
		uintptr(unsafe.Pointer(pPeerMemoryFeatures)),
		0,
	)
	return
}
func CmdSetDeviceMaskKHR(
	commandBuffer VkCommandBuffer,
	deviceMask uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDeviceMaskKHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(deviceMask),
		0,
	)
	return
}
func CmdDispatchBaseKHR(
	commandBuffer VkCommandBuffer,
	baseGroupX uint32,
	baseGroupY uint32,
	baseGroupZ uint32,
	groupCountX uint32,
	groupCountY uint32,
	groupCountZ uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDispatchBaseKHR.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(baseGroupX),
		uintptr(baseGroupY),
		uintptr(baseGroupZ),
		uintptr(groupCountX),
		uintptr(groupCountY),
		uintptr(groupCountZ),
		0,
		0,
	)
	return
}
func TrimCommandPoolKHR(
	device VkDevice,
	commandPool VkCommandPool,
	flags VkCommandPoolTrimFlags,
) {
	syscall.Syscall(
		PFN_vkTrimCommandPoolKHR.Addr(),
		3,
		uintptr(device),
		uintptr(commandPool),
		uintptr(flags),
	)
	return
}
func EnumeratePhysicalDeviceGroupsKHR(
	instance VkInstance,
	pPhysicalDeviceGroupCount *uint32,
	pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkEnumeratePhysicalDeviceGroupsKHR.Addr(),
		3,
		uintptr(instance),
		uintptr(unsafe.Pointer(pPhysicalDeviceGroupCount)),
		uintptr(unsafe.Pointer(pPhysicalDeviceGroupProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceExternalBufferPropertiesKHR(
	physicalDevice VkPhysicalDevice,
	pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo,
	pExternalBufferProperties *VkExternalBufferProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalBufferInfo)),
		uintptr(unsafe.Pointer(pExternalBufferProperties)),
	)
	return
}
func GetMemoryFdKHR(
	device VkDevice,
	pGetFdInfo *VkMemoryGetFdInfoKHR,
	pFd *int,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetMemoryFdKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetFdInfo)),
		uintptr(unsafe.Pointer(pFd)),
	)
	result = GetError((int32)(ret))
	return
}
func GetMemoryFdPropertiesKHR(
	device VkDevice,
	handleType VkExternalMemoryHandleTypeFlagBits,
	fd int,
	pMemoryFdProperties *VkMemoryFdPropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetMemoryFdPropertiesKHR.Addr(),
		4,
		uintptr(device),
		uintptr(handleType),
		uintptr(fd),
		uintptr(unsafe.Pointer(pMemoryFdProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceExternalSemaphorePropertiesKHR(
	physicalDevice VkPhysicalDevice,
	pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo,
	pExternalSemaphoreProperties *VkExternalSemaphoreProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalSemaphoreInfo)),
		uintptr(unsafe.Pointer(pExternalSemaphoreProperties)),
	)
	return
}
func ImportSemaphoreFdKHR(
	device VkDevice,
	pImportSemaphoreFdInfo *VkImportSemaphoreFdInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkImportSemaphoreFdKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pImportSemaphoreFdInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetSemaphoreFdKHR(
	device VkDevice,
	pGetFdInfo *VkSemaphoreGetFdInfoKHR,
	pFd *int,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetSemaphoreFdKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetFdInfo)),
		uintptr(unsafe.Pointer(pFd)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdPushDescriptorSetKHR(
	commandBuffer VkCommandBuffer,
	pipelineBindPoint VkPipelineBindPoint,
	layout VkPipelineLayout,
	set uint32,
	descriptorWriteCount uint32,
	pDescriptorWrites *VkWriteDescriptorSet,
) {
	syscall.Syscall6(
		PFN_vkCmdPushDescriptorSetKHR.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(pipelineBindPoint),
		uintptr(layout),
		uintptr(set),
		uintptr(descriptorWriteCount),
		uintptr(unsafe.Pointer(pDescriptorWrites)),
	)
	return
}
func CmdPushDescriptorSetWithTemplateKHR(
	commandBuffer VkCommandBuffer,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	layout VkPipelineLayout,
	set uint32,
	pData *interface{},
) {
	syscall.Syscall6(
		PFN_vkCmdPushDescriptorSetWithTemplateKHR.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(descriptorUpdateTemplate),
		uintptr(layout),
		uintptr(set),
		uintptr(unsafe.Pointer(pData)),
		0,
	)
	return
}
func CreateDescriptorUpdateTemplateKHR(
	device VkDevice,
	pCreateInfo *VkDescriptorUpdateTemplateCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pDescriptorUpdateTemplate VkDescriptorUpdateTemplate, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDescriptorUpdateTemplateKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pDescriptorUpdateTemplate)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDescriptorUpdateTemplateKHR(
	device VkDevice,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDescriptorUpdateTemplateKHR.Addr(),
		3,
		uintptr(device),
		uintptr(descriptorUpdateTemplate),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func UpdateDescriptorSetWithTemplateKHR(
	device VkDevice,
	descriptorSet VkDescriptorSet,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pData *interface{},
) {
	syscall.Syscall6(
		PFN_vkUpdateDescriptorSetWithTemplateKHR.Addr(),
		4,
		uintptr(device),
		uintptr(descriptorSet),
		uintptr(descriptorUpdateTemplate),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	return
}
func CreateRenderPass2KHR(
	device VkDevice,
	pCreateInfo *VkRenderPassCreateInfo2,
	pAllocator *VkAllocationCallbacks,
) (pRenderPass VkRenderPass, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateRenderPass2KHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pRenderPass)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdBeginRenderPass2KHR(
	commandBuffer VkCommandBuffer,
	pRenderPassBegin *VkRenderPassBeginInfo,
	pSubpassBeginInfo *VkSubpassBeginInfo,
) {
	syscall.Syscall(
		PFN_vkCmdBeginRenderPass2KHR.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pRenderPassBegin)),
		uintptr(unsafe.Pointer(pSubpassBeginInfo)),
	)
	return
}
func CmdNextSubpass2KHR(
	commandBuffer VkCommandBuffer,
	pSubpassBeginInfo *VkSubpassBeginInfo,
	pSubpassEndInfo *VkSubpassEndInfo,
) {
	syscall.Syscall(
		PFN_vkCmdNextSubpass2KHR.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pSubpassBeginInfo)),
		uintptr(unsafe.Pointer(pSubpassEndInfo)),
	)
	return
}
func CmdEndRenderPass2KHR(
	commandBuffer VkCommandBuffer,
	pSubpassEndInfo *VkSubpassEndInfo,
) {
	syscall.Syscall(
		PFN_vkCmdEndRenderPass2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pSubpassEndInfo)),
		0,
	)
	return
}
func GetSwapchainStatusKHR(
	device VkDevice,
	swapchain VkSwapchainKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetSwapchainStatusKHR.Addr(),
		2,
		uintptr(device),
		uintptr(swapchain),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceExternalFencePropertiesKHR(
	physicalDevice VkPhysicalDevice,
	pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo,
	pExternalFenceProperties *VkExternalFenceProperties,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pExternalFenceInfo)),
		uintptr(unsafe.Pointer(pExternalFenceProperties)),
	)
	return
}
func ImportFenceFdKHR(
	device VkDevice,
	pImportFenceFdInfo *VkImportFenceFdInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkImportFenceFdKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pImportFenceFdInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetFenceFdKHR(
	device VkDevice,
	pGetFdInfo *VkFenceGetFdInfoKHR,
	pFd *int,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetFenceFdKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pGetFdInfo)),
		uintptr(unsafe.Pointer(pFd)),
	)
	result = GetError((int32)(ret))
	return
}
func EnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(
	physicalDevice VkPhysicalDevice,
	queueFamilyIndex uint32,
	pCounterCount *uint32,
	pCounters *VkPerformanceCounterKHR,
	pCounterDescriptions *VkPerformanceCounterDescriptionKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.Addr(),
		5,
		uintptr(physicalDevice),
		uintptr(queueFamilyIndex),
		uintptr(unsafe.Pointer(pCounterCount)),
		uintptr(unsafe.Pointer(pCounters)),
		uintptr(unsafe.Pointer(pCounterDescriptions)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(
	physicalDevice VkPhysicalDevice,
	pPerformanceQueryCreateInfo *VkQueryPoolPerformanceCreateInfoKHR,
	pNumPasses *uint32,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPerformanceQueryCreateInfo)),
		uintptr(unsafe.Pointer(pNumPasses)),
	)
	return
}
func AcquireProfilingLockKHR(
	device VkDevice,
	pInfo *VkAcquireProfilingLockInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquireProfilingLockKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func ReleaseProfilingLockKHR(
	device VkDevice,
) {
	syscall.Syscall(
		PFN_vkReleaseProfilingLockKHR.Addr(),
		1,
		uintptr(device),
		0,
		0,
	)
	return
}
func GetPhysicalDeviceSurfaceCapabilities2KHR(
	physicalDevice VkPhysicalDevice,
	pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR,
	pSurfaceCapabilities *VkSurfaceCapabilities2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pSurfaceInfo)),
		uintptr(unsafe.Pointer(pSurfaceCapabilities)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfaceFormats2KHR(
	physicalDevice VkPhysicalDevice,
	pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR,
	pSurfaceFormatCount *uint32,
	pSurfaceFormats *VkSurfaceFormat2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPhysicalDeviceSurfaceFormats2KHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pSurfaceInfo)),
		uintptr(unsafe.Pointer(pSurfaceFormatCount)),
		uintptr(unsafe.Pointer(pSurfaceFormats)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceDisplayProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkDisplayProperties2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceDisplayProperties2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceDisplayPlaneProperties2KHR(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkDisplayPlaneProperties2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDisplayModeProperties2KHR(
	physicalDevice VkPhysicalDevice,
	display VkDisplayKHR,
	pPropertyCount *uint32,
	pProperties *VkDisplayModeProperties2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetDisplayModeProperties2KHR.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(display),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetDisplayPlaneCapabilities2KHR(
	physicalDevice VkPhysicalDevice,
	pDisplayPlaneInfo *VkDisplayPlaneInfo2KHR,
	pCapabilities *VkDisplayPlaneCapabilities2KHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDisplayPlaneCapabilities2KHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pDisplayPlaneInfo)),
		uintptr(unsafe.Pointer(pCapabilities)),
	)
	result = GetError((int32)(ret))
	return
}
func GetImageMemoryRequirements2KHR(
	device VkDevice,
	pInfo *VkImageMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
	syscall.Syscall(
		PFN_vkGetImageMemoryRequirements2KHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetBufferMemoryRequirements2KHR(
	device VkDevice,
	pInfo *VkBufferMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
	syscall.Syscall(
		PFN_vkGetBufferMemoryRequirements2KHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func GetImageSparseMemoryRequirements2KHR(
	device VkDevice,
	pInfo *VkImageSparseMemoryRequirementsInfo2,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements2,
) {
	syscall.Syscall6(
		PFN_vkGetImageSparseMemoryRequirements2KHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)),
		uintptr(unsafe.Pointer(pSparseMemoryRequirements)),
		0,
		0,
	)
	return
}
func CreateSamplerYcbcrConversionKHR(
	device VkDevice,
	pCreateInfo *VkSamplerYcbcrConversionCreateInfo,
	pAllocator *VkAllocationCallbacks,
) (pYcbcrConversion VkSamplerYcbcrConversion, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateSamplerYcbcrConversionKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pYcbcrConversion)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroySamplerYcbcrConversionKHR(
	device VkDevice,
	ycbcrConversion VkSamplerYcbcrConversion,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroySamplerYcbcrConversionKHR.Addr(),
		3,
		uintptr(device),
		uintptr(ycbcrConversion),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func BindBufferMemory2KHR(
	device VkDevice,
	bindInfoCount uint32,
	pBindInfos *VkBindBufferMemoryInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBindBufferMemory2KHR.Addr(),
		3,
		uintptr(device),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfos)),
	)
	result = GetError((int32)(ret))
	return
}
func BindImageMemory2KHR(
	device VkDevice,
	bindInfoCount uint32,
	pBindInfos *VkBindImageMemoryInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBindImageMemory2KHR.Addr(),
		3,
		uintptr(device),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfos)),
	)
	result = GetError((int32)(ret))
	return
}
func GetDescriptorSetLayoutSupportKHR(
	device VkDevice,
	pCreateInfo *VkDescriptorSetLayoutCreateInfo,
	pSupport *VkDescriptorSetLayoutSupport,
) {
	syscall.Syscall(
		PFN_vkGetDescriptorSetLayoutSupportKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pSupport)),
	)
	return
}
func CmdDrawIndirectCountKHR(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndirectCountKHR.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func CmdDrawIndexedIndirectCountKHR(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndexedIndirectCountKHR.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func GetSemaphoreCounterValueKHR(
	device VkDevice,
	semaphore VkSemaphore,
	pValue *uint64,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetSemaphoreCounterValueKHR.Addr(),
		3,
		uintptr(device),
		uintptr(semaphore),
		uintptr(unsafe.Pointer(pValue)),
	)
	result = GetError((int32)(ret))
	return
}
func WaitSemaphoresKHR(
	device VkDevice,
	pWaitInfo *VkSemaphoreWaitInfo,
	timeout uint64,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkWaitSemaphoresKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pWaitInfo)),
		uintptr(timeout),
	)
	result = GetError((int32)(ret))
	return
}
func SignalSemaphoreKHR(
	device VkDevice,
	pSignalInfo *VkSemaphoreSignalInfo,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkSignalSemaphoreKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pSignalInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceFragmentShadingRatesKHR(
	physicalDevice VkPhysicalDevice,
	pFragmentShadingRateCount *uint32,
	pFragmentShadingRates *VkPhysicalDeviceFragmentShadingRateKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceFragmentShadingRatesKHR.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pFragmentShadingRateCount)),
		uintptr(unsafe.Pointer(pFragmentShadingRates)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetFragmentShadingRateKHR(
	commandBuffer VkCommandBuffer,
	pFragmentSize *VkExtent2D,
	combinerOps [2]VkFragmentShadingRateCombinerOpKHR,
) {
	syscall.Syscall(
		PFN_vkCmdSetFragmentShadingRateKHR.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pFragmentSize)),
		uintptr(unsafe.Pointer(&combinerOps[0])),
	)
	return
}
func WaitForPresentKHR(
	device VkDevice,
	swapchain VkSwapchainKHR,
	presentId uint64,
	timeout uint64,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkWaitForPresentKHR.Addr(),
		4,
		uintptr(device),
		uintptr(swapchain),
		uintptr(presentId),
		uintptr(timeout),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetBufferDeviceAddressKHR(
	device VkDevice,
	pInfo *VkBufferDeviceAddressInfo,
) (result VkDeviceAddress) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetBufferDeviceAddressKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (VkDeviceAddress)(ret)
	return
}
func GetBufferOpaqueCaptureAddressKHR(
	device VkDevice,
	pInfo *VkBufferDeviceAddressInfo,
) (result uint64) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetBufferOpaqueCaptureAddressKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (uint64)(ret)
	return
}
func GetDeviceMemoryOpaqueCaptureAddressKHR(
	device VkDevice,
	pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo,
) (result uint64) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (uint64)(ret)
	return
}
func CreateDeferredOperationKHR(
	device VkDevice,
	pAllocator *VkAllocationCallbacks,
) (pDeferredOperation VkDeferredOperationKHR, result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCreateDeferredOperationKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pDeferredOperation)),
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDeferredOperationKHR(
	device VkDevice,
	operation VkDeferredOperationKHR,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDeferredOperationKHR.Addr(),
		3,
		uintptr(device),
		uintptr(operation),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetDeferredOperationMaxConcurrencyKHR(
	device VkDevice,
	operation VkDeferredOperationKHR,
) (result uint32) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeferredOperationMaxConcurrencyKHR.Addr(),
		2,
		uintptr(device),
		uintptr(operation),
		0,
	)
	result = (uint32)(ret)
	return
}
func GetDeferredOperationResultKHR(
	device VkDevice,
	operation VkDeferredOperationKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeferredOperationResultKHR.Addr(),
		2,
		uintptr(device),
		uintptr(operation),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DeferredOperationJoinKHR(
	device VkDevice,
	operation VkDeferredOperationKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkDeferredOperationJoinKHR.Addr(),
		2,
		uintptr(device),
		uintptr(operation),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPipelineExecutablePropertiesKHR(
	device VkDevice,
	pPipelineInfo *VkPipelineInfoKHR,
	pExecutableCount *uint32,
	pProperties *VkPipelineExecutablePropertiesKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPipelineExecutablePropertiesKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pPipelineInfo)),
		uintptr(unsafe.Pointer(pExecutableCount)),
		uintptr(unsafe.Pointer(pProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPipelineExecutableStatisticsKHR(
	device VkDevice,
	pExecutableInfo *VkPipelineExecutableInfoKHR,
	pStatisticCount *uint32,
	pStatistics *VkPipelineExecutableStatisticKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPipelineExecutableStatisticsKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pExecutableInfo)),
		uintptr(unsafe.Pointer(pStatisticCount)),
		uintptr(unsafe.Pointer(pStatistics)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPipelineExecutableInternalRepresentationsKHR(
	device VkDevice,
	pExecutableInfo *VkPipelineExecutableInfoKHR,
	pInternalRepresentationCount *uint32,
	pInternalRepresentations *VkPipelineExecutableInternalRepresentationKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPipelineExecutableInternalRepresentationsKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pExecutableInfo)),
		uintptr(unsafe.Pointer(pInternalRepresentationCount)),
		uintptr(unsafe.Pointer(pInternalRepresentations)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetEvent2KHR(
	commandBuffer VkCommandBuffer,
	event VkEvent,
	pDependencyInfo *VkDependencyInfoKHR,
) {
	syscall.Syscall(
		PFN_vkCmdSetEvent2KHR.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(event),
		uintptr(unsafe.Pointer(pDependencyInfo)),
	)
	return
}
func CmdResetEvent2KHR(
	commandBuffer VkCommandBuffer,
	event VkEvent,
	stageMask VkPipelineStageFlags2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdResetEvent2KHR.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(event),
		uintptr(stageMask),
	)
	return
}
func CmdWaitEvents2KHR(
	commandBuffer VkCommandBuffer,
	eventCount uint32,
	pEvents *VkEvent,
	pDependencyInfos *VkDependencyInfoKHR,
) {
	syscall.Syscall6(
		PFN_vkCmdWaitEvents2KHR.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(eventCount),
		uintptr(unsafe.Pointer(pEvents)),
		uintptr(unsafe.Pointer(pDependencyInfos)),
		0,
		0,
	)
	return
}
func CmdPipelineBarrier2KHR(
	commandBuffer VkCommandBuffer,
	pDependencyInfo *VkDependencyInfoKHR,
) {
	syscall.Syscall(
		PFN_vkCmdPipelineBarrier2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pDependencyInfo)),
		0,
	)
	return
}
func CmdWriteTimestamp2KHR(
	commandBuffer VkCommandBuffer,
	stage VkPipelineStageFlags2KHR,
	queryPool VkQueryPool,
	query uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteTimestamp2KHR.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(stage),
		uintptr(queryPool),
		uintptr(query),
		0,
		0,
	)
	return
}
func QueueSubmit2KHR(
	queue VkQueue,
	submitCount uint32,
	pSubmits *VkSubmitInfo2KHR,
	fence VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkQueueSubmit2KHR.Addr(),
		4,
		uintptr(queue),
		uintptr(submitCount),
		uintptr(unsafe.Pointer(pSubmits)),
		uintptr(fence),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdWriteBufferMarker2AMD(
	commandBuffer VkCommandBuffer,
	stage VkPipelineStageFlags2KHR,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	marker uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteBufferMarker2AMD.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(stage),
		uintptr(dstBuffer),
		uintptr(dstOffset),
		uintptr(marker),
		0,
	)
	return
}
func GetQueueCheckpointData2NV(
	queue VkQueue,
	pCheckpointDataCount *uint32,
	pCheckpointData *VkCheckpointData2NV,
) {
	syscall.Syscall(
		PFN_vkGetQueueCheckpointData2NV.Addr(),
		3,
		uintptr(queue),
		uintptr(unsafe.Pointer(pCheckpointDataCount)),
		uintptr(unsafe.Pointer(pCheckpointData)),
	)
	return
}
func CmdCopyBuffer2KHR(
	commandBuffer VkCommandBuffer,
	pCopyBufferInfo *VkCopyBufferInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyBuffer2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pCopyBufferInfo)),
		0,
	)
	return
}
func CmdCopyImage2KHR(
	commandBuffer VkCommandBuffer,
	pCopyImageInfo *VkCopyImageInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyImage2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pCopyImageInfo)),
		0,
	)
	return
}
func CmdCopyBufferToImage2KHR(
	commandBuffer VkCommandBuffer,
	pCopyBufferToImageInfo *VkCopyBufferToImageInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyBufferToImage2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pCopyBufferToImageInfo)),
		0,
	)
	return
}
func CmdCopyImageToBuffer2KHR(
	commandBuffer VkCommandBuffer,
	pCopyImageToBufferInfo *VkCopyImageToBufferInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyImageToBuffer2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pCopyImageToBufferInfo)),
		0,
	)
	return
}
func CmdBlitImage2KHR(
	commandBuffer VkCommandBuffer,
	pBlitImageInfo *VkBlitImageInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdBlitImage2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pBlitImageInfo)),
		0,
	)
	return
}
func CmdResolveImage2KHR(
	commandBuffer VkCommandBuffer,
	pResolveImageInfo *VkResolveImageInfo2KHR,
) {
	syscall.Syscall(
		PFN_vkCmdResolveImage2KHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pResolveImageInfo)),
		0,
	)
	return
}
func CreateDebugReportCallbackEXT(
	instance VkInstance,
	pCreateInfo *VkDebugReportCallbackCreateInfoEXT,
	pAllocator *VkAllocationCallbacks,
) (pCallback VkDebugReportCallbackEXT, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDebugReportCallbackEXT.Addr(),
		4,
		uintptr(instance),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pCallback)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDebugReportCallbackEXT(
	instance VkInstance,
	callback VkDebugReportCallbackEXT,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDebugReportCallbackEXT.Addr(),
		3,
		uintptr(instance),
		uintptr(callback),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func DebugReportMessageEXT(
	instance VkInstance,
	flags VkDebugReportFlagsEXT,
	objectType VkDebugReportObjectTypeEXT,
	object uint64,
	location uint64,
	messageCode int32,
	pLayerPrefix string,
	pMessage string,
) {
	charptrpLayerPrefix := TEXT(pLayerPrefix)
	charptrpMessage := TEXT(pMessage)
	syscall.Syscall9(
		PFN_vkDebugReportMessageEXT.Addr(),
		8,
		uintptr(instance),
		uintptr(flags),
		uintptr(objectType),
		uintptr(object),
		uintptr(location),
		uintptr(messageCode),
		uintptr(unsafe.Pointer(charptrpLayerPrefix)),
		uintptr(unsafe.Pointer(charptrpMessage)),
		0,
	)
	return
}
func DebugMarkerSetObjectTagEXT(
	device VkDevice,
	pTagInfo *VkDebugMarkerObjectTagInfoEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkDebugMarkerSetObjectTagEXT.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pTagInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DebugMarkerSetObjectNameEXT(
	device VkDevice,
	pNameInfo *VkDebugMarkerObjectNameInfoEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkDebugMarkerSetObjectNameEXT.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pNameInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdDebugMarkerBeginEXT(
	commandBuffer VkCommandBuffer,
	pMarkerInfo *VkDebugMarkerMarkerInfoEXT,
) {
	syscall.Syscall(
		PFN_vkCmdDebugMarkerBeginEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pMarkerInfo)),
		0,
	)
	return
}
func CmdDebugMarkerEndEXT(
	commandBuffer VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdDebugMarkerEndEXT.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	return
}
func CmdDebugMarkerInsertEXT(
	commandBuffer VkCommandBuffer,
	pMarkerInfo *VkDebugMarkerMarkerInfoEXT,
) {
	syscall.Syscall(
		PFN_vkCmdDebugMarkerInsertEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pMarkerInfo)),
		0,
	)
	return
}
func CmdBindTransformFeedbackBuffersEXT(
	commandBuffer VkCommandBuffer,
	firstBinding uint32,
	bindingCount uint32,
	pBuffers *VkBuffer,
	pOffsets *VkDeviceSize,
	pSizes *VkDeviceSize,
) {
	syscall.Syscall6(
		PFN_vkCmdBindTransformFeedbackBuffersEXT.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(firstBinding),
		uintptr(bindingCount),
		uintptr(unsafe.Pointer(pBuffers)),
		uintptr(unsafe.Pointer(pOffsets)),
		uintptr(unsafe.Pointer(pSizes)),
	)
	return
}
func CmdBeginTransformFeedbackEXT(
	commandBuffer VkCommandBuffer,
	firstCounterBuffer uint32,
	counterBufferCount uint32,
	pCounterBuffers *VkBuffer,
	pCounterBufferOffsets *VkDeviceSize,
) {
	syscall.Syscall6(
		PFN_vkCmdBeginTransformFeedbackEXT.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(firstCounterBuffer),
		uintptr(counterBufferCount),
		uintptr(unsafe.Pointer(pCounterBuffers)),
		uintptr(unsafe.Pointer(pCounterBufferOffsets)),
		0,
	)
	return
}
func CmdEndTransformFeedbackEXT(
	commandBuffer VkCommandBuffer,
	firstCounterBuffer uint32,
	counterBufferCount uint32,
	pCounterBuffers *VkBuffer,
	pCounterBufferOffsets *VkDeviceSize,
) {
	syscall.Syscall6(
		PFN_vkCmdEndTransformFeedbackEXT.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(firstCounterBuffer),
		uintptr(counterBufferCount),
		uintptr(unsafe.Pointer(pCounterBuffers)),
		uintptr(unsafe.Pointer(pCounterBufferOffsets)),
		0,
	)
	return
}
func CmdBeginQueryIndexedEXT(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	query uint32,
	flags VkQueryControlFlags,
	index uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdBeginQueryIndexedEXT.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(query),
		uintptr(flags),
		uintptr(index),
		0,
	)
	return
}
func CmdEndQueryIndexedEXT(
	commandBuffer VkCommandBuffer,
	queryPool VkQueryPool,
	query uint32,
	index uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdEndQueryIndexedEXT.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(queryPool),
		uintptr(query),
		uintptr(index),
		0,
		0,
	)
	return
}
func CmdDrawIndirectByteCountEXT(
	commandBuffer VkCommandBuffer,
	instanceCount uint32,
	firstInstance uint32,
	counterBuffer VkBuffer,
	counterBufferOffset VkDeviceSize,
	counterOffset uint32,
	vertexStride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndirectByteCountEXT.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(instanceCount),
		uintptr(firstInstance),
		uintptr(counterBuffer),
		uintptr(counterBufferOffset),
		uintptr(counterOffset),
		uintptr(vertexStride),
		0,
		0,
	)
	return
}
func CreateCuModuleNVX(
	device VkDevice,
	pCreateInfo *VkCuModuleCreateInfoNVX,
	pAllocator *VkAllocationCallbacks,
) (pModule VkCuModuleNVX, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateCuModuleNVX.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pModule)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreateCuFunctionNVX(
	device VkDevice,
	pCreateInfo *VkCuFunctionCreateInfoNVX,
	pAllocator *VkAllocationCallbacks,
) (pFunction VkCuFunctionNVX, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateCuFunctionNVX.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pFunction)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyCuModuleNVX(
	device VkDevice,
	module VkCuModuleNVX,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyCuModuleNVX.Addr(),
		3,
		uintptr(device),
		uintptr(module),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func DestroyCuFunctionNVX(
	device VkDevice,
	function VkCuFunctionNVX,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyCuFunctionNVX.Addr(),
		3,
		uintptr(device),
		uintptr(function),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CmdCuLaunchKernelNVX(
	commandBuffer VkCommandBuffer,
	pLaunchInfo *VkCuLaunchInfoNVX,
) {
	syscall.Syscall(
		PFN_vkCmdCuLaunchKernelNVX.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pLaunchInfo)),
		0,
	)
	return
}
func GetImageViewHandleNVX(
	device VkDevice,
	pInfo *VkImageViewHandleInfoNVX,
) (result uint32) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetImageViewHandleNVX.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (uint32)(ret)
	return
}
func GetImageViewAddressNVX(
	device VkDevice,
	imageView VkImageView,
	pProperties *VkImageViewAddressPropertiesNVX,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetImageViewAddressNVX.Addr(),
		3,
		uintptr(device),
		uintptr(imageView),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdDrawIndirectCountAMD(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndirectCountAMD.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func CmdDrawIndexedIndirectCountAMD(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawIndexedIndirectCountAMD.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func GetShaderInfoAMD(
	device VkDevice,
	pipeline VkPipeline,
	shaderStage VkShaderStageFlagBits,
	infoType VkShaderInfoTypeAMD,
	pInfoSize *uint64,
	pInfo *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetShaderInfoAMD.Addr(),
		6,
		uintptr(device),
		uintptr(pipeline),
		uintptr(shaderStage),
		uintptr(infoType),
		uintptr(unsafe.Pointer(pInfoSize)),
		uintptr(unsafe.Pointer(pInfo)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceExternalImageFormatPropertiesNV(
	physicalDevice VkPhysicalDevice,
	format VkFormat,
	vtype VkImageType,
	tiling VkImageTiling,
	usage VkImageUsageFlags,
	flags VkImageCreateFlags,
	externalHandleType VkExternalMemoryHandleTypeFlagsNV,
	pExternalImageFormatProperties *VkExternalImageFormatPropertiesNV,
) (result error) {
	ret, _, _ := syscall.Syscall9(
		PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV.Addr(),
		8,
		uintptr(physicalDevice),
		uintptr(format),
		uintptr(vtype),
		uintptr(tiling),
		uintptr(usage),
		uintptr(flags),
		uintptr(externalHandleType),
		uintptr(unsafe.Pointer(pExternalImageFormatProperties)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdBeginConditionalRenderingEXT(
	commandBuffer VkCommandBuffer,
	pConditionalRenderingBegin *VkConditionalRenderingBeginInfoEXT,
) {
	syscall.Syscall(
		PFN_vkCmdBeginConditionalRenderingEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pConditionalRenderingBegin)),
		0,
	)
	return
}
func CmdEndConditionalRenderingEXT(
	commandBuffer VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdEndConditionalRenderingEXT.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	return
}
func CmdSetViewportWScalingNV(
	commandBuffer VkCommandBuffer,
	firstViewport uint32,
	viewportCount uint32,
	pViewportWScalings *VkViewportWScalingNV,
) {
	syscall.Syscall6(
		PFN_vkCmdSetViewportWScalingNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstViewport),
		uintptr(viewportCount),
		uintptr(unsafe.Pointer(pViewportWScalings)),
		0,
		0,
	)
	return
}
func ReleaseDisplayEXT(
	physicalDevice VkPhysicalDevice,
	display VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkReleaseDisplayEXT.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(display),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSurfaceCapabilities2EXT(
	physicalDevice VkPhysicalDevice,
	surface VkSurfaceKHR,
	pSurfaceCapabilities *VkSurfaceCapabilities2EXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(surface),
		uintptr(unsafe.Pointer(pSurfaceCapabilities)),
	)
	result = GetError((int32)(ret))
	return
}
func DisplayPowerControlEXT(
	device VkDevice,
	display VkDisplayKHR,
	pDisplayPowerInfo *VkDisplayPowerInfoEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkDisplayPowerControlEXT.Addr(),
		3,
		uintptr(device),
		uintptr(display),
		uintptr(unsafe.Pointer(pDisplayPowerInfo)),
	)
	result = GetError((int32)(ret))
	return
}
func RegisterDeviceEventEXT(
	device VkDevice,
	pDeviceEventInfo *VkDeviceEventInfoEXT,
	pAllocator *VkAllocationCallbacks,
	pFence *VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkRegisterDeviceEventEXT.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pDeviceEventInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(pFence)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func RegisterDisplayEventEXT(
	device VkDevice,
	display VkDisplayKHR,
	pDisplayEventInfo *VkDisplayEventInfoEXT,
	pAllocator *VkAllocationCallbacks,
	pFence *VkFence,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkRegisterDisplayEventEXT.Addr(),
		5,
		uintptr(device),
		uintptr(display),
		uintptr(unsafe.Pointer(pDisplayEventInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(pFence)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetSwapchainCounterEXT(
	device VkDevice,
	swapchain VkSwapchainKHR,
	counter VkSurfaceCounterFlagBitsEXT,
	pCounterValue *uint64,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetSwapchainCounterEXT.Addr(),
		4,
		uintptr(device),
		uintptr(swapchain),
		uintptr(counter),
		uintptr(unsafe.Pointer(pCounterValue)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetRefreshCycleDurationGOOGLE(
	device VkDevice,
	swapchain VkSwapchainKHR,
	pDisplayTimingProperties *VkRefreshCycleDurationGOOGLE,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetRefreshCycleDurationGOOGLE.Addr(),
		3,
		uintptr(device),
		uintptr(swapchain),
		uintptr(unsafe.Pointer(pDisplayTimingProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPastPresentationTimingGOOGLE(
	device VkDevice,
	swapchain VkSwapchainKHR,
	pPresentationTimingCount *uint32,
	pPresentationTimings *VkPastPresentationTimingGOOGLE,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetPastPresentationTimingGOOGLE.Addr(),
		4,
		uintptr(device),
		uintptr(swapchain),
		uintptr(unsafe.Pointer(pPresentationTimingCount)),
		uintptr(unsafe.Pointer(pPresentationTimings)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetDiscardRectangleEXT(
	commandBuffer VkCommandBuffer,
	firstDiscardRectangle uint32,
	discardRectangleCount uint32,
	pDiscardRectangles *VkRect2D,
) {
	syscall.Syscall6(
		PFN_vkCmdSetDiscardRectangleEXT.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstDiscardRectangle),
		uintptr(discardRectangleCount),
		uintptr(unsafe.Pointer(pDiscardRectangles)),
		0,
		0,
	)
	return
}
func SetHdrMetadataEXT(
	device VkDevice,
	swapchainCount uint32,
	pSwapchains *VkSwapchainKHR,
	pMetadata *VkHdrMetadataEXT,
) {
	syscall.Syscall6(
		PFN_vkSetHdrMetadataEXT.Addr(),
		4,
		uintptr(device),
		uintptr(swapchainCount),
		uintptr(unsafe.Pointer(pSwapchains)),
		uintptr(unsafe.Pointer(pMetadata)),
		0,
		0,
	)
	return
}
func SetDebugUtilsObjectNameEXT(
	device VkDevice,
	pNameInfo *VkDebugUtilsObjectNameInfoEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkSetDebugUtilsObjectNameEXT.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pNameInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func SetDebugUtilsObjectTagEXT(
	device VkDevice,
	pTagInfo *VkDebugUtilsObjectTagInfoEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkSetDebugUtilsObjectTagEXT.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pTagInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func QueueBeginDebugUtilsLabelEXT(
	queue VkQueue,
	pLabelInfo *VkDebugUtilsLabelEXT,
) {
	syscall.Syscall(
		PFN_vkQueueBeginDebugUtilsLabelEXT.Addr(),
		2,
		uintptr(queue),
		uintptr(unsafe.Pointer(pLabelInfo)),
		0,
	)
	return
}
func QueueEndDebugUtilsLabelEXT(
	queue VkQueue,
) {
	syscall.Syscall(
		PFN_vkQueueEndDebugUtilsLabelEXT.Addr(),
		1,
		uintptr(queue),
		0,
		0,
	)
	return
}
func QueueInsertDebugUtilsLabelEXT(
	queue VkQueue,
	pLabelInfo *VkDebugUtilsLabelEXT,
) {
	syscall.Syscall(
		PFN_vkQueueInsertDebugUtilsLabelEXT.Addr(),
		2,
		uintptr(queue),
		uintptr(unsafe.Pointer(pLabelInfo)),
		0,
	)
	return
}
func CmdBeginDebugUtilsLabelEXT(
	commandBuffer VkCommandBuffer,
	pLabelInfo *VkDebugUtilsLabelEXT,
) {
	syscall.Syscall(
		PFN_vkCmdBeginDebugUtilsLabelEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pLabelInfo)),
		0,
	)
	return
}
func CmdEndDebugUtilsLabelEXT(
	commandBuffer VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdEndDebugUtilsLabelEXT.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	return
}
func CmdInsertDebugUtilsLabelEXT(
	commandBuffer VkCommandBuffer,
	pLabelInfo *VkDebugUtilsLabelEXT,
) {
	syscall.Syscall(
		PFN_vkCmdInsertDebugUtilsLabelEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pLabelInfo)),
		0,
	)
	return
}
func CreateDebugUtilsMessengerEXT(
	instance VkInstance,
	pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT,
	pAllocator *VkAllocationCallbacks,
) (pMessenger VkDebugUtilsMessengerEXT, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateDebugUtilsMessengerEXT.Addr(),
		4,
		uintptr(instance),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pMessenger)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyDebugUtilsMessengerEXT(
	instance VkInstance,
	messenger VkDebugUtilsMessengerEXT,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyDebugUtilsMessengerEXT.Addr(),
		3,
		uintptr(instance),
		uintptr(messenger),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func SubmitDebugUtilsMessageEXT(
	instance VkInstance,
	messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT,
	messageTypes VkDebugUtilsMessageTypeFlagsEXT,
	pCallbackData *VkDebugUtilsMessengerCallbackDataEXT,
) {
	syscall.Syscall6(
		PFN_vkSubmitDebugUtilsMessageEXT.Addr(),
		4,
		uintptr(instance),
		uintptr(messageSeverity),
		uintptr(messageTypes),
		uintptr(unsafe.Pointer(pCallbackData)),
		0,
		0,
	)
	return
}
func CmdSetSampleLocationsEXT(
	commandBuffer VkCommandBuffer,
	pSampleLocationsInfo *VkSampleLocationsInfoEXT,
) {
	syscall.Syscall(
		PFN_vkCmdSetSampleLocationsEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pSampleLocationsInfo)),
		0,
	)
	return
}
func GetPhysicalDeviceMultisamplePropertiesEXT(
	physicalDevice VkPhysicalDevice,
	samples VkSampleCountFlagBits,
	pMultisampleProperties *VkMultisamplePropertiesEXT,
) {
	syscall.Syscall(
		PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(samples),
		uintptr(unsafe.Pointer(pMultisampleProperties)),
	)
	return
}
func GetImageDrmFormatModifierPropertiesEXT(
	device VkDevice,
	image VkImage,
	pProperties *VkImageDrmFormatModifierPropertiesEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetImageDrmFormatModifierPropertiesEXT.Addr(),
		3,
		uintptr(device),
		uintptr(image),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func CreateValidationCacheEXT(
	device VkDevice,
	pCreateInfo *VkValidationCacheCreateInfoEXT,
	pAllocator *VkAllocationCallbacks,
) (pValidationCache VkValidationCacheEXT, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateValidationCacheEXT.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pValidationCache)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyValidationCacheEXT(
	device VkDevice,
	validationCache VkValidationCacheEXT,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyValidationCacheEXT.Addr(),
		3,
		uintptr(device),
		uintptr(validationCache),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func MergeValidationCachesEXT(
	device VkDevice,
	dstCache VkValidationCacheEXT,
	srcCacheCount uint32,
	pSrcCaches *VkValidationCacheEXT,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkMergeValidationCachesEXT.Addr(),
		4,
		uintptr(device),
		uintptr(dstCache),
		uintptr(srcCacheCount),
		uintptr(unsafe.Pointer(pSrcCaches)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetValidationCacheDataEXT(
	device VkDevice,
	validationCache VkValidationCacheEXT,
	pDataSize *uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetValidationCacheDataEXT.Addr(),
		4,
		uintptr(device),
		uintptr(validationCache),
		uintptr(unsafe.Pointer(pDataSize)),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdBindShadingRateImageNV(
	commandBuffer VkCommandBuffer,
	imageView VkImageView,
	imageLayout VkImageLayout,
) {
	syscall.Syscall(
		PFN_vkCmdBindShadingRateImageNV.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(imageView),
		uintptr(imageLayout),
	)
	return
}
func CmdSetViewportShadingRatePaletteNV(
	commandBuffer VkCommandBuffer,
	firstViewport uint32,
	viewportCount uint32,
	pShadingRatePalettes *VkShadingRatePaletteNV,
) {
	syscall.Syscall6(
		PFN_vkCmdSetViewportShadingRatePaletteNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstViewport),
		uintptr(viewportCount),
		uintptr(unsafe.Pointer(pShadingRatePalettes)),
		0,
		0,
	)
	return
}
func CmdSetCoarseSampleOrderNV(
	commandBuffer VkCommandBuffer,
	sampleOrderType VkCoarseSampleOrderTypeNV,
	customSampleOrderCount uint32,
	pCustomSampleOrders *VkCoarseSampleOrderCustomNV,
) {
	syscall.Syscall6(
		PFN_vkCmdSetCoarseSampleOrderNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(sampleOrderType),
		uintptr(customSampleOrderCount),
		uintptr(unsafe.Pointer(pCustomSampleOrders)),
		0,
		0,
	)
	return
}
func CreateAccelerationStructureNV(
	device VkDevice,
	pCreateInfo *VkAccelerationStructureCreateInfoNV,
	pAllocator *VkAllocationCallbacks,
) (pAccelerationStructure VkAccelerationStructureNV, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateAccelerationStructureNV.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pAccelerationStructure)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyAccelerationStructureNV(
	device VkDevice,
	accelerationStructure VkAccelerationStructureNV,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyAccelerationStructureNV.Addr(),
		3,
		uintptr(device),
		uintptr(accelerationStructure),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func GetAccelerationStructureMemoryRequirementsNV(
	device VkDevice,
	pInfo *VkAccelerationStructureMemoryRequirementsInfoNV,
	pMemoryRequirements *VkMemoryRequirements2KHR,
) {
	syscall.Syscall(
		PFN_vkGetAccelerationStructureMemoryRequirementsNV.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func BindAccelerationStructureMemoryNV(
	device VkDevice,
	bindInfoCount uint32,
	pBindInfos *VkBindAccelerationStructureMemoryInfoNV,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkBindAccelerationStructureMemoryNV.Addr(),
		3,
		uintptr(device),
		uintptr(bindInfoCount),
		uintptr(unsafe.Pointer(pBindInfos)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdBuildAccelerationStructureNV(
	commandBuffer VkCommandBuffer,
	pInfo *VkAccelerationStructureInfoNV,
	instanceData VkBuffer,
	instanceOffset VkDeviceSize,
	update VkBool32,
	dst VkAccelerationStructureNV,
	src VkAccelerationStructureNV,
	scratch VkBuffer,
	scratchOffset VkDeviceSize,
) {
	syscall.Syscall9(
		PFN_vkCmdBuildAccelerationStructureNV.Addr(),
		9,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(instanceData),
		uintptr(instanceOffset),
		uintptr(update),
		uintptr(dst),
		uintptr(src),
		uintptr(scratch),
		uintptr(scratchOffset),
	)
	return
}
func CmdCopyAccelerationStructureNV(
	commandBuffer VkCommandBuffer,
	dst VkAccelerationStructureNV,
	src VkAccelerationStructureNV,
	mode VkCopyAccelerationStructureModeKHR,
) {
	syscall.Syscall6(
		PFN_vkCmdCopyAccelerationStructureNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(dst),
		uintptr(src),
		uintptr(mode),
		0,
		0,
	)
	return
}
func CmdTraceRaysNV(
	commandBuffer VkCommandBuffer,
	raygenShaderBindingTableBuffer VkBuffer,
	raygenShaderBindingOffset VkDeviceSize,
	missShaderBindingTableBuffer VkBuffer,
	missShaderBindingOffset VkDeviceSize,
	missShaderBindingStride VkDeviceSize,
	hitShaderBindingTableBuffer VkBuffer,
	hitShaderBindingOffset VkDeviceSize,
	hitShaderBindingStride VkDeviceSize,
	callableShaderBindingTableBuffer VkBuffer,
	callableShaderBindingOffset VkDeviceSize,
	callableShaderBindingStride VkDeviceSize,
	width uint32,
	height uint32,
	depth uint32,
) {
	syscall.Syscall15(
		PFN_vkCmdTraceRaysNV.Addr(),
		15,
		uintptr(commandBuffer),
		uintptr(raygenShaderBindingTableBuffer),
		uintptr(raygenShaderBindingOffset),
		uintptr(missShaderBindingTableBuffer),
		uintptr(missShaderBindingOffset),
		uintptr(missShaderBindingStride),
		uintptr(hitShaderBindingTableBuffer),
		uintptr(hitShaderBindingOffset),
		uintptr(hitShaderBindingStride),
		uintptr(callableShaderBindingTableBuffer),
		uintptr(callableShaderBindingOffset),
		uintptr(callableShaderBindingStride),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
	)
	return
}
func CreateRayTracingPipelinesNV(
	device VkDevice,
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkRayTracingPipelineCreateInfoNV,
	pAllocator *VkAllocationCallbacks,
) (pPipelines VkPipeline, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateRayTracingPipelinesNV.Addr(),
		6,
		uintptr(device),
		uintptr(pipelineCache),
		uintptr(createInfoCount),
		uintptr(unsafe.Pointer(pCreateInfos)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelines)),
	)
	result = GetError((int32)(ret))
	return
}
func GetRayTracingShaderGroupHandlesKHR(
	device VkDevice,
	pipeline VkPipeline,
	firstGroup uint32,
	groupCount uint32,
	dataSize uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetRayTracingShaderGroupHandlesKHR.Addr(),
		6,
		uintptr(device),
		uintptr(pipeline),
		uintptr(firstGroup),
		uintptr(groupCount),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
	)
	result = GetError((int32)(ret))
	return
}
func GetRayTracingShaderGroupHandlesNV(
	device VkDevice,
	pipeline VkPipeline,
	firstGroup uint32,
	groupCount uint32,
	dataSize uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetRayTracingShaderGroupHandlesNV.Addr(),
		6,
		uintptr(device),
		uintptr(pipeline),
		uintptr(firstGroup),
		uintptr(groupCount),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
	)
	result = GetError((int32)(ret))
	return
}
func GetAccelerationStructureHandleNV(
	device VkDevice,
	accelerationStructure VkAccelerationStructureNV,
	dataSize uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetAccelerationStructureHandleNV.Addr(),
		4,
		uintptr(device),
		uintptr(accelerationStructure),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdWriteAccelerationStructuresPropertiesNV(
	commandBuffer VkCommandBuffer,
	accelerationStructureCount uint32,
	pAccelerationStructures *VkAccelerationStructureNV,
	queryType VkQueryType,
	queryPool VkQueryPool,
	firstQuery uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteAccelerationStructuresPropertiesNV.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(accelerationStructureCount),
		uintptr(unsafe.Pointer(pAccelerationStructures)),
		uintptr(queryType),
		uintptr(queryPool),
		uintptr(firstQuery),
	)
	return
}
func CompileDeferredNV(
	device VkDevice,
	pipeline VkPipeline,
	shader uint32,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCompileDeferredNV.Addr(),
		3,
		uintptr(device),
		uintptr(pipeline),
		uintptr(shader),
	)
	result = GetError((int32)(ret))
	return
}
func GetMemoryHostPointerPropertiesEXT(
	device VkDevice,
	handleType VkExternalMemoryHandleTypeFlagBits,
	pHostPointer *interface{},
	pMemoryHostPointerProperties *VkMemoryHostPointerPropertiesEXT,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetMemoryHostPointerPropertiesEXT.Addr(),
		4,
		uintptr(device),
		uintptr(handleType),
		uintptr(unsafe.Pointer(pHostPointer)),
		uintptr(unsafe.Pointer(pMemoryHostPointerProperties)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdWriteBufferMarkerAMD(
	commandBuffer VkCommandBuffer,
	pipelineStage VkPipelineStageFlagBits,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	marker uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteBufferMarkerAMD.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(pipelineStage),
		uintptr(dstBuffer),
		uintptr(dstOffset),
		uintptr(marker),
		0,
	)
	return
}
func GetPhysicalDeviceCalibrateableTimeDomainsEXT(
	physicalDevice VkPhysicalDevice,
	pTimeDomainCount *uint32,
	pTimeDomains *VkTimeDomainEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pTimeDomainCount)),
		uintptr(unsafe.Pointer(pTimeDomains)),
	)
	result = GetError((int32)(ret))
	return
}
func GetCalibratedTimestampsEXT(
	device VkDevice,
	timestampCount uint32,
	pTimestampInfos *VkCalibratedTimestampInfoEXT,
	pTimestamps *uint64,
	pMaxDeviation *uint64,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetCalibratedTimestampsEXT.Addr(),
		5,
		uintptr(device),
		uintptr(timestampCount),
		uintptr(unsafe.Pointer(pTimestampInfos)),
		uintptr(unsafe.Pointer(pTimestamps)),
		uintptr(unsafe.Pointer(pMaxDeviation)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdDrawMeshTasksNV(
	commandBuffer VkCommandBuffer,
	taskCount uint32,
	firstTask uint32,
) {
	syscall.Syscall(
		PFN_vkCmdDrawMeshTasksNV.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(taskCount),
		uintptr(firstTask),
	)
	return
}
func CmdDrawMeshTasksIndirectNV(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	drawCount uint32,
	stride uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDrawMeshTasksIndirectNV.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(drawCount),
		uintptr(stride),
		0,
	)
	return
}
func CmdDrawMeshTasksIndirectCountNV(
	commandBuffer VkCommandBuffer,
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawMeshTasksIndirectCountNV.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(buffer),
		uintptr(offset),
		uintptr(countBuffer),
		uintptr(countBufferOffset),
		uintptr(maxDrawCount),
		uintptr(stride),
		0,
		0,
	)
	return
}
func CmdSetExclusiveScissorNV(
	commandBuffer VkCommandBuffer,
	firstExclusiveScissor uint32,
	exclusiveScissorCount uint32,
	pExclusiveScissors *VkRect2D,
) {
	syscall.Syscall6(
		PFN_vkCmdSetExclusiveScissorNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(firstExclusiveScissor),
		uintptr(exclusiveScissorCount),
		uintptr(unsafe.Pointer(pExclusiveScissors)),
		0,
		0,
	)
	return
}
func CmdSetCheckpointNV(
	commandBuffer VkCommandBuffer,
	pCheckpointMarker *interface{},
) {
	syscall.Syscall(
		PFN_vkCmdSetCheckpointNV.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pCheckpointMarker)),
		0,
	)
	return
}
func GetQueueCheckpointDataNV(
	queue VkQueue,
	pCheckpointDataCount *uint32,
	pCheckpointData *VkCheckpointDataNV,
) {
	syscall.Syscall(
		PFN_vkGetQueueCheckpointDataNV.Addr(),
		3,
		uintptr(queue),
		uintptr(unsafe.Pointer(pCheckpointDataCount)),
		uintptr(unsafe.Pointer(pCheckpointData)),
	)
	return
}
func InitializePerformanceApiINTEL(
	device VkDevice,
	pInitializeInfo *VkInitializePerformanceApiInfoINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkInitializePerformanceApiINTEL.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInitializeInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func UninitializePerformanceApiINTEL(
	device VkDevice,
) {
	syscall.Syscall(
		PFN_vkUninitializePerformanceApiINTEL.Addr(),
		1,
		uintptr(device),
		0,
		0,
	)
	return
}
func CmdSetPerformanceMarkerINTEL(
	commandBuffer VkCommandBuffer,
	pMarkerInfo *VkPerformanceMarkerInfoINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCmdSetPerformanceMarkerINTEL.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pMarkerInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetPerformanceStreamMarkerINTEL(
	commandBuffer VkCommandBuffer,
	pMarkerInfo *VkPerformanceStreamMarkerInfoINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCmdSetPerformanceStreamMarkerINTEL.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pMarkerInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetPerformanceOverrideINTEL(
	commandBuffer VkCommandBuffer,
	pOverrideInfo *VkPerformanceOverrideInfoINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCmdSetPerformanceOverrideINTEL.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pOverrideInfo)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func AcquirePerformanceConfigurationINTEL(
	device VkDevice,
	pAcquireInfo *VkPerformanceConfigurationAcquireInfoINTEL,
	pConfiguration *VkPerformanceConfigurationINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquirePerformanceConfigurationINTEL.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pAcquireInfo)),
		uintptr(unsafe.Pointer(pConfiguration)),
	)
	result = GetError((int32)(ret))
	return
}
func ReleasePerformanceConfigurationINTEL(
	device VkDevice,
	configuration VkPerformanceConfigurationINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkReleasePerformanceConfigurationINTEL.Addr(),
		2,
		uintptr(device),
		uintptr(configuration),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func QueueSetPerformanceConfigurationINTEL(
	queue VkQueue,
	configuration VkPerformanceConfigurationINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkQueueSetPerformanceConfigurationINTEL.Addr(),
		2,
		uintptr(queue),
		uintptr(configuration),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPerformanceParameterINTEL(
	device VkDevice,
	parameter VkPerformanceParameterTypeINTEL,
	pValue *VkPerformanceValueINTEL,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPerformanceParameterINTEL.Addr(),
		3,
		uintptr(device),
		uintptr(parameter),
		uintptr(unsafe.Pointer(pValue)),
	)
	result = GetError((int32)(ret))
	return
}
func SetLocalDimmingAMD(
	device VkDevice,
	swapChain VkSwapchainKHR,
	localDimmingEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkSetLocalDimmingAMD.Addr(),
		3,
		uintptr(device),
		uintptr(swapChain),
		uintptr(localDimmingEnable),
	)
	return
}
func GetBufferDeviceAddressEXT(
	device VkDevice,
	pInfo *VkBufferDeviceAddressInfo,
) (result VkDeviceAddress) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetBufferDeviceAddressEXT.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (VkDeviceAddress)(ret)
	return
}
func GetPhysicalDeviceToolPropertiesEXT(
	physicalDevice VkPhysicalDevice,
	pToolCount *uint32,
	pToolProperties *VkPhysicalDeviceToolPropertiesEXT,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceToolPropertiesEXT.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pToolCount)),
		uintptr(unsafe.Pointer(pToolProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceCooperativeMatrixPropertiesNV(
	physicalDevice VkPhysicalDevice,
	pPropertyCount *uint32,
	pProperties *VkCooperativeMatrixPropertiesNV,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pPropertyCount)),
		uintptr(unsafe.Pointer(pProperties)),
	)
	result = GetError((int32)(ret))
	return
}
func GetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(
	physicalDevice VkPhysicalDevice,
	pCombinationCount *uint32,
	pCombinations *VkFramebufferMixedSamplesCombinationNV,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(unsafe.Pointer(pCombinationCount)),
		uintptr(unsafe.Pointer(pCombinations)),
	)
	result = GetError((int32)(ret))
	return
}
func CreateHeadlessSurfaceEXT(
	instance VkInstance,
	pCreateInfo *VkHeadlessSurfaceCreateInfoEXT,
	pAllocator *VkAllocationCallbacks,
) (pSurface VkSurfaceKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateHeadlessSurfaceEXT.Addr(),
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
func CmdSetLineStippleEXT(
	commandBuffer VkCommandBuffer,
	lineStippleFactor uint32,
	lineStipplePattern uint16,
) {
	syscall.Syscall(
		PFN_vkCmdSetLineStippleEXT.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(lineStippleFactor),
		uintptr(lineStipplePattern),
	)
	return
}
func ResetQueryPoolEXT(
	device VkDevice,
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
) {
	syscall.Syscall6(
		PFN_vkResetQueryPoolEXT.Addr(),
		4,
		uintptr(device),
		uintptr(queryPool),
		uintptr(firstQuery),
		uintptr(queryCount),
		0,
		0,
	)
	return
}
func CmdSetCullModeEXT(
	commandBuffer VkCommandBuffer,
	cullMode VkCullModeFlags,
) {
	syscall.Syscall(
		PFN_vkCmdSetCullModeEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(cullMode),
		0,
	)
	return
}
func CmdSetFrontFaceEXT(
	commandBuffer VkCommandBuffer,
	frontFace VkFrontFace,
) {
	syscall.Syscall(
		PFN_vkCmdSetFrontFaceEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(frontFace),
		0,
	)
	return
}
func CmdSetPrimitiveTopologyEXT(
	commandBuffer VkCommandBuffer,
	primitiveTopology VkPrimitiveTopology,
) {
	syscall.Syscall(
		PFN_vkCmdSetPrimitiveTopologyEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(primitiveTopology),
		0,
	)
	return
}
func CmdSetViewportWithCountEXT(
	commandBuffer VkCommandBuffer,
	viewportCount uint32,
	pViewports *VkViewport,
) {
	syscall.Syscall(
		PFN_vkCmdSetViewportWithCountEXT.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(viewportCount),
		uintptr(unsafe.Pointer(pViewports)),
	)
	return
}
func CmdSetScissorWithCountEXT(
	commandBuffer VkCommandBuffer,
	scissorCount uint32,
	pScissors *VkRect2D,
) {
	syscall.Syscall(
		PFN_vkCmdSetScissorWithCountEXT.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(scissorCount),
		uintptr(unsafe.Pointer(pScissors)),
	)
	return
}
func CmdBindVertexBuffers2EXT(
	commandBuffer VkCommandBuffer,
	firstBinding uint32,
	bindingCount uint32,
	pBuffers *VkBuffer,
	pOffsets *VkDeviceSize,
	pSizes *VkDeviceSize,
	pStrides *VkDeviceSize,
) {
	syscall.Syscall9(
		PFN_vkCmdBindVertexBuffers2EXT.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(firstBinding),
		uintptr(bindingCount),
		uintptr(unsafe.Pointer(pBuffers)),
		uintptr(unsafe.Pointer(pOffsets)),
		uintptr(unsafe.Pointer(pSizes)),
		uintptr(unsafe.Pointer(pStrides)),
		0,
		0,
	)
	return
}
func CmdSetDepthTestEnableEXT(
	commandBuffer VkCommandBuffer,
	depthTestEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthTestEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(depthTestEnable),
		0,
	)
	return
}
func CmdSetDepthWriteEnableEXT(
	commandBuffer VkCommandBuffer,
	depthWriteEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthWriteEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(depthWriteEnable),
		0,
	)
	return
}
func CmdSetDepthCompareOpEXT(
	commandBuffer VkCommandBuffer,
	depthCompareOp VkCompareOp,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthCompareOpEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(depthCompareOp),
		0,
	)
	return
}
func CmdSetDepthBoundsTestEnableEXT(
	commandBuffer VkCommandBuffer,
	depthBoundsTestEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthBoundsTestEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(depthBoundsTestEnable),
		0,
	)
	return
}
func CmdSetStencilTestEnableEXT(
	commandBuffer VkCommandBuffer,
	stencilTestEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetStencilTestEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(stencilTestEnable),
		0,
	)
	return
}
func CmdSetStencilOpEXT(
	commandBuffer VkCommandBuffer,
	faceMask VkStencilFaceFlags,
	failOp VkStencilOp,
	passOp VkStencilOp,
	depthFailOp VkStencilOp,
	compareOp VkCompareOp,
) {
	syscall.Syscall6(
		PFN_vkCmdSetStencilOpEXT.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(faceMask),
		uintptr(failOp),
		uintptr(passOp),
		uintptr(depthFailOp),
		uintptr(compareOp),
	)
	return
}
func GetGeneratedCommandsMemoryRequirementsNV(
	device VkDevice,
	pInfo *VkGeneratedCommandsMemoryRequirementsInfoNV,
	pMemoryRequirements *VkMemoryRequirements2,
) {
	syscall.Syscall(
		PFN_vkGetGeneratedCommandsMemoryRequirementsNV.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		uintptr(unsafe.Pointer(pMemoryRequirements)),
	)
	return
}
func CmdPreprocessGeneratedCommandsNV(
	commandBuffer VkCommandBuffer,
	pGeneratedCommandsInfo *VkGeneratedCommandsInfoNV,
) {
	syscall.Syscall(
		PFN_vkCmdPreprocessGeneratedCommandsNV.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pGeneratedCommandsInfo)),
		0,
	)
	return
}
func CmdExecuteGeneratedCommandsNV(
	commandBuffer VkCommandBuffer,
	isPreprocessed VkBool32,
	pGeneratedCommandsInfo *VkGeneratedCommandsInfoNV,
) {
	syscall.Syscall(
		PFN_vkCmdExecuteGeneratedCommandsNV.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(isPreprocessed),
		uintptr(unsafe.Pointer(pGeneratedCommandsInfo)),
	)
	return
}
func CmdBindPipelineShaderGroupNV(
	commandBuffer VkCommandBuffer,
	pipelineBindPoint VkPipelineBindPoint,
	pipeline VkPipeline,
	groupIndex uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdBindPipelineShaderGroupNV.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(pipelineBindPoint),
		uintptr(pipeline),
		uintptr(groupIndex),
		0,
		0,
	)
	return
}
func CreateIndirectCommandsLayoutNV(
	device VkDevice,
	pCreateInfo *VkIndirectCommandsLayoutCreateInfoNV,
	pAllocator *VkAllocationCallbacks,
) (pIndirectCommandsLayout VkIndirectCommandsLayoutNV, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateIndirectCommandsLayoutNV.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pIndirectCommandsLayout)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyIndirectCommandsLayoutNV(
	device VkDevice,
	indirectCommandsLayout VkIndirectCommandsLayoutNV,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyIndirectCommandsLayoutNV.Addr(),
		3,
		uintptr(device),
		uintptr(indirectCommandsLayout),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func AcquireDrmDisplayEXT(
	physicalDevice VkPhysicalDevice,
	drmFd int32,
	display VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquireDrmDisplayEXT.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(drmFd),
		uintptr(display),
	)
	result = GetError((int32)(ret))
	return
}
func GetDrmDisplayEXT(
	physicalDevice VkPhysicalDevice,
	drmFd int32,
	connectorId uint32,
	display *VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetDrmDisplayEXT.Addr(),
		4,
		uintptr(physicalDevice),
		uintptr(drmFd),
		uintptr(connectorId),
		uintptr(unsafe.Pointer(display)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CreatePrivateDataSlotEXT(
	device VkDevice,
	pCreateInfo *VkPrivateDataSlotCreateInfoEXT,
	pAllocator *VkAllocationCallbacks,
) (pPrivateDataSlot VkPrivateDataSlotEXT, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreatePrivateDataSlotEXT.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPrivateDataSlot)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyPrivateDataSlotEXT(
	device VkDevice,
	privateDataSlot VkPrivateDataSlotEXT,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyPrivateDataSlotEXT.Addr(),
		3,
		uintptr(device),
		uintptr(privateDataSlot),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func SetPrivateDataEXT(
	device VkDevice,
	objectType VkObjectType,
	objectHandle uint64,
	privateDataSlot VkPrivateDataSlotEXT,
	data uint64,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkSetPrivateDataEXT.Addr(),
		5,
		uintptr(device),
		uintptr(objectType),
		uintptr(objectHandle),
		uintptr(privateDataSlot),
		uintptr(data),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetPrivateDataEXT(
	device VkDevice,
	objectType VkObjectType,
	objectHandle uint64,
	privateDataSlot VkPrivateDataSlotEXT,
	pData *uint64,
) {
	syscall.Syscall6(
		PFN_vkGetPrivateDataEXT.Addr(),
		5,
		uintptr(device),
		uintptr(objectType),
		uintptr(objectHandle),
		uintptr(privateDataSlot),
		uintptr(unsafe.Pointer(pData)),
		0,
	)
	return
}
func CmdSetFragmentShadingRateEnumNV(
	commandBuffer VkCommandBuffer,
	shadingRate VkFragmentShadingRateNV,
	combinerOps [2]VkFragmentShadingRateCombinerOpKHR,
) {
	syscall.Syscall(
		PFN_vkCmdSetFragmentShadingRateEnumNV.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(shadingRate),
		uintptr(unsafe.Pointer(&combinerOps[0])),
	)
	return
}
func AcquireWinrtDisplayNV(
	physicalDevice VkPhysicalDevice,
	display VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkAcquireWinrtDisplayNV.Addr(),
		2,
		uintptr(physicalDevice),
		uintptr(display),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetWinrtDisplayNV(
	physicalDevice VkPhysicalDevice,
	deviceRelativeId uint32,
	pDisplay *VkDisplayKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetWinrtDisplayNV.Addr(),
		3,
		uintptr(physicalDevice),
		uintptr(deviceRelativeId),
		uintptr(unsafe.Pointer(pDisplay)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetVertexInputEXT(
	commandBuffer VkCommandBuffer,
	vertexBindingDescriptionCount uint32,
	pVertexBindingDescriptions *VkVertexInputBindingDescription2EXT,
	vertexAttributeDescriptionCount uint32,
	pVertexAttributeDescriptions *VkVertexInputAttributeDescription2EXT,
) {
	syscall.Syscall6(
		PFN_vkCmdSetVertexInputEXT.Addr(),
		5,
		uintptr(commandBuffer),
		uintptr(vertexBindingDescriptionCount),
		uintptr(unsafe.Pointer(pVertexBindingDescriptions)),
		uintptr(vertexAttributeDescriptionCount),
		uintptr(unsafe.Pointer(pVertexAttributeDescriptions)),
		0,
	)
	return
}
func GetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(
	device VkDevice,
	renderpass VkRenderPass,
	pMaxWorkgroupSize *VkExtent2D,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI.Addr(),
		3,
		uintptr(device),
		uintptr(renderpass),
		uintptr(unsafe.Pointer(pMaxWorkgroupSize)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdSubpassShadingHUAWEI(
	commandBuffer VkCommandBuffer,
) {
	syscall.Syscall(
		PFN_vkCmdSubpassShadingHUAWEI.Addr(),
		1,
		uintptr(commandBuffer),
		0,
		0,
	)
	return
}
func CmdBindInvocationMaskHUAWEI(
	commandBuffer VkCommandBuffer,
	imageView VkImageView,
	imageLayout VkImageLayout,
) {
	syscall.Syscall(
		PFN_vkCmdBindInvocationMaskHUAWEI.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(imageView),
		uintptr(imageLayout),
	)
	return
}
func GetMemoryRemoteAddressNV(
	device VkDevice,
	pMemoryGetRemoteAddressInfo *VkMemoryGetRemoteAddressInfoNV,
	pAddress *VkRemoteAddressNV,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetMemoryRemoteAddressNV.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pMemoryGetRemoteAddressInfo)),
		uintptr(unsafe.Pointer(pAddress)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdSetPatchControlPointsEXT(
	commandBuffer VkCommandBuffer,
	patchControlPoints uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetPatchControlPointsEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(patchControlPoints),
		0,
	)
	return
}
func CmdSetRasterizerDiscardEnableEXT(
	commandBuffer VkCommandBuffer,
	rasterizerDiscardEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetRasterizerDiscardEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(rasterizerDiscardEnable),
		0,
	)
	return
}
func CmdSetDepthBiasEnableEXT(
	commandBuffer VkCommandBuffer,
	depthBiasEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetDepthBiasEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(depthBiasEnable),
		0,
	)
	return
}
func CmdSetLogicOpEXT(
	commandBuffer VkCommandBuffer,
	logicOp VkLogicOp,
) {
	syscall.Syscall(
		PFN_vkCmdSetLogicOpEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(logicOp),
		0,
	)
	return
}
func CmdSetPrimitiveRestartEnableEXT(
	commandBuffer VkCommandBuffer,
	primitiveRestartEnable VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetPrimitiveRestartEnableEXT.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(primitiveRestartEnable),
		0,
	)
	return
}
func CmdSetColorWriteEnableEXT(
	commandBuffer VkCommandBuffer,
	attachmentCount uint32,
	pColorWriteEnables *VkBool32,
) {
	syscall.Syscall(
		PFN_vkCmdSetColorWriteEnableEXT.Addr(),
		3,
		uintptr(commandBuffer),
		uintptr(attachmentCount),
		uintptr(unsafe.Pointer(pColorWriteEnables)),
	)
	return
}
func CmdDrawMultiEXT(
	commandBuffer VkCommandBuffer,
	drawCount uint32,
	pVertexInfo *VkMultiDrawInfoEXT,
	instanceCount uint32,
	firstInstance uint32,
	stride uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdDrawMultiEXT.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(drawCount),
		uintptr(unsafe.Pointer(pVertexInfo)),
		uintptr(instanceCount),
		uintptr(firstInstance),
		uintptr(stride),
	)
	return
}
func CmdDrawMultiIndexedEXT(
	commandBuffer VkCommandBuffer,
	drawCount uint32,
	pIndexInfo *VkMultiDrawIndexedInfoEXT,
	instanceCount uint32,
	firstInstance uint32,
	stride uint32,
	pVertexOffset *int32,
) {
	syscall.Syscall9(
		PFN_vkCmdDrawMultiIndexedEXT.Addr(),
		7,
		uintptr(commandBuffer),
		uintptr(drawCount),
		uintptr(unsafe.Pointer(pIndexInfo)),
		uintptr(instanceCount),
		uintptr(firstInstance),
		uintptr(stride),
		uintptr(unsafe.Pointer(pVertexOffset)),
		0,
		0,
	)
	return
}
func CreateAccelerationStructureKHR(
	device VkDevice,
	pCreateInfo *VkAccelerationStructureCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pAccelerationStructure VkAccelerationStructureKHR, result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkCreateAccelerationStructureKHR.Addr(),
		4,
		uintptr(device),
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pAccelerationStructure)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func DestroyAccelerationStructureKHR(
	device VkDevice,
	accelerationStructure VkAccelerationStructureKHR,
	pAllocator *VkAllocationCallbacks,
) {
	syscall.Syscall(
		PFN_vkDestroyAccelerationStructureKHR.Addr(),
		3,
		uintptr(device),
		uintptr(accelerationStructure),
		uintptr(unsafe.Pointer(pAllocator)),
	)
	return
}
func CmdBuildAccelerationStructuresKHR(
	commandBuffer VkCommandBuffer,
	infoCount uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
) {
	syscall.Syscall6(
		PFN_vkCmdBuildAccelerationStructuresKHR.Addr(),
		4,
		uintptr(commandBuffer),
		uintptr(infoCount),
		uintptr(unsafe.Pointer(pInfos)),
		uintptr(unsafe.Pointer(ppBuildRangeInfos)),
		0,
		0,
	)
	return
}
func CmdBuildAccelerationStructuresIndirectKHR(
	commandBuffer VkCommandBuffer,
	infoCount uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	pIndirectDeviceAddresses *VkDeviceAddress,
	pIndirectStrides *uint32,
	ppMaxPrimitiveCounts **uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdBuildAccelerationStructuresIndirectKHR.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(infoCount),
		uintptr(unsafe.Pointer(pInfos)),
		uintptr(unsafe.Pointer(pIndirectDeviceAddresses)),
		uintptr(unsafe.Pointer(pIndirectStrides)),
		uintptr(unsafe.Pointer(ppMaxPrimitiveCounts)),
	)
	return
}
func BuildAccelerationStructuresKHR(
	device VkDevice,
	deferredOperation VkDeferredOperationKHR,
	infoCount uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkBuildAccelerationStructuresKHR.Addr(),
		5,
		uintptr(device),
		uintptr(deferredOperation),
		uintptr(infoCount),
		uintptr(unsafe.Pointer(pInfos)),
		uintptr(unsafe.Pointer(ppBuildRangeInfos)),
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CopyAccelerationStructureKHR(
	device VkDevice,
	deferredOperation VkDeferredOperationKHR,
	pInfo *VkCopyAccelerationStructureInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCopyAccelerationStructureKHR.Addr(),
		3,
		uintptr(device),
		uintptr(deferredOperation),
		uintptr(unsafe.Pointer(pInfo)),
	)
	result = GetError((int32)(ret))
	return
}
func CopyAccelerationStructureToMemoryKHR(
	device VkDevice,
	deferredOperation VkDeferredOperationKHR,
	pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCopyAccelerationStructureToMemoryKHR.Addr(),
		3,
		uintptr(device),
		uintptr(deferredOperation),
		uintptr(unsafe.Pointer(pInfo)),
	)
	result = GetError((int32)(ret))
	return
}
func CopyMemoryToAccelerationStructureKHR(
	device VkDevice,
	deferredOperation VkDeferredOperationKHR,
	pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
) (result error) {
	ret, _, _ := syscall.Syscall(
		PFN_vkCopyMemoryToAccelerationStructureKHR.Addr(),
		3,
		uintptr(device),
		uintptr(deferredOperation),
		uintptr(unsafe.Pointer(pInfo)),
	)
	result = GetError((int32)(ret))
	return
}
func WriteAccelerationStructuresPropertiesKHR(
	device VkDevice,
	accelerationStructureCount uint32,
	pAccelerationStructures *VkAccelerationStructureKHR,
	queryType VkQueryType,
	dataSize uint64,
	pData *interface{},
	stride uint64,
) (result error) {
	ret, _, _ := syscall.Syscall9(
		PFN_vkWriteAccelerationStructuresPropertiesKHR.Addr(),
		7,
		uintptr(device),
		uintptr(accelerationStructureCount),
		uintptr(unsafe.Pointer(pAccelerationStructures)),
		uintptr(queryType),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
		uintptr(stride),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func CmdCopyAccelerationStructureKHR(
	commandBuffer VkCommandBuffer,
	pInfo *VkCopyAccelerationStructureInfoKHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyAccelerationStructureKHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	return
}
func CmdCopyAccelerationStructureToMemoryKHR(
	commandBuffer VkCommandBuffer,
	pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyAccelerationStructureToMemoryKHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	return
}
func CmdCopyMemoryToAccelerationStructureKHR(
	commandBuffer VkCommandBuffer,
	pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
) {
	syscall.Syscall(
		PFN_vkCmdCopyMemoryToAccelerationStructureKHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	return
}
func GetAccelerationStructureDeviceAddressKHR(
	device VkDevice,
	pInfo *VkAccelerationStructureDeviceAddressInfoKHR,
) (result VkDeviceAddress) {
	ret, _, _ := syscall.Syscall(
		PFN_vkGetAccelerationStructureDeviceAddressKHR.Addr(),
		2,
		uintptr(device),
		uintptr(unsafe.Pointer(pInfo)),
		0,
	)
	result = (VkDeviceAddress)(ret)
	return
}
func CmdWriteAccelerationStructuresPropertiesKHR(
	commandBuffer VkCommandBuffer,
	accelerationStructureCount uint32,
	pAccelerationStructures *VkAccelerationStructureKHR,
	queryType VkQueryType,
	queryPool VkQueryPool,
	firstQuery uint32,
) {
	syscall.Syscall6(
		PFN_vkCmdWriteAccelerationStructuresPropertiesKHR.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(accelerationStructureCount),
		uintptr(unsafe.Pointer(pAccelerationStructures)),
		uintptr(queryType),
		uintptr(queryPool),
		uintptr(firstQuery),
	)
	return
}
func GetDeviceAccelerationStructureCompatibilityKHR(
	device VkDevice,
	pVersionInfo *VkAccelerationStructureVersionInfoKHR,
	pCompatibility *VkAccelerationStructureCompatibilityKHR,
) {
	syscall.Syscall(
		PFN_vkGetDeviceAccelerationStructureCompatibilityKHR.Addr(),
		3,
		uintptr(device),
		uintptr(unsafe.Pointer(pVersionInfo)),
		uintptr(unsafe.Pointer(pCompatibility)),
	)
	return
}
func GetAccelerationStructureBuildSizesKHR(
	device VkDevice,
	buildType VkAccelerationStructureBuildTypeKHR,
	pBuildInfo *VkAccelerationStructureBuildGeometryInfoKHR,
	pMaxPrimitiveCounts *uint32,
	pSizeInfo *VkAccelerationStructureBuildSizesInfoKHR,
) {
	syscall.Syscall6(
		PFN_vkGetAccelerationStructureBuildSizesKHR.Addr(),
		5,
		uintptr(device),
		uintptr(buildType),
		uintptr(unsafe.Pointer(pBuildInfo)),
		uintptr(unsafe.Pointer(pMaxPrimitiveCounts)),
		uintptr(unsafe.Pointer(pSizeInfo)),
		0,
	)
	return
}
func CmdTraceRaysKHR(
	commandBuffer VkCommandBuffer,
	pRaygenShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pMissShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pHitShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pCallableShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	width uint32,
	height uint32,
	depth uint32,
) {
	syscall.Syscall9(
		PFN_vkCmdTraceRaysKHR.Addr(),
		8,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pRaygenShaderBindingTable)),
		uintptr(unsafe.Pointer(pMissShaderBindingTable)),
		uintptr(unsafe.Pointer(pHitShaderBindingTable)),
		uintptr(unsafe.Pointer(pCallableShaderBindingTable)),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		0,
	)
	return
}
func CreateRayTracingPipelinesKHR(
	device VkDevice,
	deferredOperation VkDeferredOperationKHR,
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkRayTracingPipelineCreateInfoKHR,
	pAllocator *VkAllocationCallbacks,
) (pPipelines VkPipeline, result error) {
	ret, _, _ := syscall.Syscall9(
		PFN_vkCreateRayTracingPipelinesKHR.Addr(),
		7,
		uintptr(device),
		uintptr(deferredOperation),
		uintptr(pipelineCache),
		uintptr(createInfoCount),
		uintptr(unsafe.Pointer(pCreateInfos)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(&pPipelines)),
		0,
		0,
	)
	result = GetError((int32)(ret))
	return
}
func GetRayTracingCaptureReplayShaderGroupHandlesKHR(
	device VkDevice,
	pipeline VkPipeline,
	firstGroup uint32,
	groupCount uint32,
	dataSize uint64,
	pData *interface{},
) (result error) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.Addr(),
		6,
		uintptr(device),
		uintptr(pipeline),
		uintptr(firstGroup),
		uintptr(groupCount),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(pData)),
	)
	result = GetError((int32)(ret))
	return
}
func CmdTraceRaysIndirectKHR(
	commandBuffer VkCommandBuffer,
	pRaygenShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pMissShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pHitShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	pCallableShaderBindingTable *VkStridedDeviceAddressRegionKHR,
	indirectDeviceAddress VkDeviceAddress,
) {
	syscall.Syscall6(
		PFN_vkCmdTraceRaysIndirectKHR.Addr(),
		6,
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(pRaygenShaderBindingTable)),
		uintptr(unsafe.Pointer(pMissShaderBindingTable)),
		uintptr(unsafe.Pointer(pHitShaderBindingTable)),
		uintptr(unsafe.Pointer(pCallableShaderBindingTable)),
		uintptr(indirectDeviceAddress),
	)
	return
}
func GetRayTracingShaderGroupStackSizeKHR(
	device VkDevice,
	pipeline VkPipeline,
	group uint32,
	groupShader VkShaderGroupShaderKHR,
) (result VkDeviceSize) {
	ret, _, _ := syscall.Syscall6(
		PFN_vkGetRayTracingShaderGroupStackSizeKHR.Addr(),
		4,
		uintptr(device),
		uintptr(pipeline),
		uintptr(group),
		uintptr(groupShader),
		0,
		0,
	)
	result = (VkDeviceSize)(ret)
	return
}
func CmdSetRayTracingPipelineStackSizeKHR(
	commandBuffer VkCommandBuffer,
	pipelineStackSize uint32,
) {
	syscall.Syscall(
		PFN_vkCmdSetRayTracingPipelineStackSizeKHR.Addr(),
		2,
		uintptr(commandBuffer),
		uintptr(pipelineStackSize),
		0,
	)
	return
}
