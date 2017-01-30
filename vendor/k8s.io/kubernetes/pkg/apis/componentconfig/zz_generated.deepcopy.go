// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package componentconfig

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	config "k8s.io/kubernetes/pkg/util/config"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_IPVar, InType: reflect.TypeOf(&IPVar{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeControllerManagerConfiguration, InType: reflect.TypeOf(&KubeControllerManagerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeProxyConfiguration, InType: reflect.TypeOf(&KubeProxyConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeSchedulerConfiguration, InType: reflect.TypeOf(&KubeSchedulerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAnonymousAuthentication, InType: reflect.TypeOf(&KubeletAnonymousAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAuthentication, InType: reflect.TypeOf(&KubeletAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAuthorization, InType: reflect.TypeOf(&KubeletAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletConfiguration, InType: reflect.TypeOf(&KubeletConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletWebhookAuthentication, InType: reflect.TypeOf(&KubeletWebhookAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletWebhookAuthorization, InType: reflect.TypeOf(&KubeletWebhookAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletX509Authentication, InType: reflect.TypeOf(&KubeletX509Authentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_LeaderElectionConfiguration, InType: reflect.TypeOf(&LeaderElectionConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration, InType: reflect.TypeOf(&PersistentVolumeRecyclerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PortRangeVar, InType: reflect.TypeOf(&PortRangeVar{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_VolumeConfiguration, InType: reflect.TypeOf(&VolumeConfiguration{})},
	)
}

func DeepCopy_componentconfig_IPVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IPVar)
		out := out.(*IPVar)
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		} else {
			out.Val = nil
		}
		return nil
	}
}

func DeepCopy_componentconfig_KubeControllerManagerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeControllerManagerConfiguration)
		out := out.(*KubeControllerManagerConfiguration)
		out.TypeMeta = in.TypeMeta
		out.Port = in.Port
		out.Address = in.Address
		out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
		out.CloudProvider = in.CloudProvider
		out.CloudConfigFile = in.CloudConfigFile
		out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
		out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
		out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
		out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
		out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
		out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
		out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
		out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
		out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
		out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
		out.LookupCacheSizeForRC = in.LookupCacheSizeForRC
		out.LookupCacheSizeForRS = in.LookupCacheSizeForRS
		out.LookupCacheSizeForDaemonSet = in.LookupCacheSizeForDaemonSet
		out.ServiceSyncPeriod = in.ServiceSyncPeriod
		out.NodeSyncPeriod = in.NodeSyncPeriod
		out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
		out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
		out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
		out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
		out.MinResyncPeriod = in.MinResyncPeriod
		out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
		out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
		out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
		out.PodEvictionTimeout = in.PodEvictionTimeout
		out.DeletingPodsQps = in.DeletingPodsQps
		out.DeletingPodsBurst = in.DeletingPodsBurst
		out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
		out.RegisterRetryCount = in.RegisterRetryCount
		out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
		out.NodeMonitorPeriod = in.NodeMonitorPeriod
		out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
		out.ClusterSigningCertFile = in.ClusterSigningCertFile
		out.ClusterSigningKeyFile = in.ClusterSigningKeyFile
		out.ApproveAllKubeletCSRsForGroup = in.ApproveAllKubeletCSRsForGroup
		out.EnableProfiling = in.EnableProfiling
		out.ClusterName = in.ClusterName
		out.ClusterCIDR = in.ClusterCIDR
		out.ServiceCIDR = in.ServiceCIDR
		out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
		out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
		out.ConfigureCloudRoutes = in.ConfigureCloudRoutes
		out.RootCAFile = in.RootCAFile
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.LeaderElection = in.LeaderElection
		out.VolumeConfiguration = in.VolumeConfiguration
		out.ControllerStartInterval = in.ControllerStartInterval
		out.EnableGarbageCollector = in.EnableGarbageCollector
		out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
		out.NodeEvictionRate = in.NodeEvictionRate
		out.SecondaryNodeEvictionRate = in.SecondaryNodeEvictionRate
		out.LargeClusterSizeThreshold = in.LargeClusterSizeThreshold
		out.UnhealthyZoneThreshold = in.UnhealthyZoneThreshold
		out.DisableAttachDetachReconcilerSync = in.DisableAttachDetachReconcilerSync
		out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
		return nil
	}
}

func DeepCopy_componentconfig_KubeProxyConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyConfiguration)
		out := out.(*KubeProxyConfiguration)
		out.TypeMeta = in.TypeMeta
		out.BindAddress = in.BindAddress
		out.ClusterCIDR = in.ClusterCIDR
		out.HealthzBindAddress = in.HealthzBindAddress
		out.HealthzPort = in.HealthzPort
		out.HostnameOverride = in.HostnameOverride
		if in.IPTablesMasqueradeBit != nil {
			in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
			*out = new(int32)
			**out = **in
		} else {
			out.IPTablesMasqueradeBit = nil
		}
		out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
		out.IPTablesMinSyncPeriod = in.IPTablesMinSyncPeriod
		out.KubeconfigPath = in.KubeconfigPath
		out.MasqueradeAll = in.MasqueradeAll
		out.Master = in.Master
		if in.OOMScoreAdj != nil {
			in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
			*out = new(int32)
			**out = **in
		} else {
			out.OOMScoreAdj = nil
		}
		out.Mode = in.Mode
		out.PortRange = in.PortRange
		out.ResourceContainer = in.ResourceContainer
		out.UDPIdleTimeout = in.UDPIdleTimeout
		out.ConntrackMax = in.ConntrackMax
		out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
		out.ConntrackMin = in.ConntrackMin
		out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
		out.ConntrackTCPCloseWaitTimeout = in.ConntrackTCPCloseWaitTimeout
		return nil
	}
}

func DeepCopy_componentconfig_KubeSchedulerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeSchedulerConfiguration)
		out := out.(*KubeSchedulerConfiguration)
		out.TypeMeta = in.TypeMeta
		out.Port = in.Port
		out.Address = in.Address
		out.AlgorithmProvider = in.AlgorithmProvider
		out.PolicyConfigFile = in.PolicyConfigFile
		out.EnableProfiling = in.EnableProfiling
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.SchedulerName = in.SchedulerName
		out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
		out.FailureDomains = in.FailureDomains
		out.LeaderElection = in.LeaderElection
		return nil
	}
}

func DeepCopy_componentconfig_KubeletAnonymousAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAnonymousAuthentication)
		out := out.(*KubeletAnonymousAuthentication)
		out.Enabled = in.Enabled
		return nil
	}
}

func DeepCopy_componentconfig_KubeletAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthentication)
		out := out.(*KubeletAuthentication)
		out.X509 = in.X509
		out.Webhook = in.Webhook
		out.Anonymous = in.Anonymous
		return nil
	}
}

func DeepCopy_componentconfig_KubeletAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthorization)
		out := out.(*KubeletAuthorization)
		out.Mode = in.Mode
		out.Webhook = in.Webhook
		return nil
	}
}

func DeepCopy_componentconfig_KubeletConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletConfiguration)
		out := out.(*KubeletConfiguration)
		out.TypeMeta = in.TypeMeta
		out.PodManifestPath = in.PodManifestPath
		out.SyncFrequency = in.SyncFrequency
		out.FileCheckFrequency = in.FileCheckFrequency
		out.HTTPCheckFrequency = in.HTTPCheckFrequency
		out.ManifestURL = in.ManifestURL
		out.ManifestURLHeader = in.ManifestURLHeader
		out.EnableServer = in.EnableServer
		out.Address = in.Address
		out.Port = in.Port
		out.ReadOnlyPort = in.ReadOnlyPort
		out.TLSCertFile = in.TLSCertFile
		out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
		out.CertDirectory = in.CertDirectory
		out.Authentication = in.Authentication
		out.Authorization = in.Authorization
		out.HostnameOverride = in.HostnameOverride
		out.PodInfraContainerImage = in.PodInfraContainerImage
		out.DockerEndpoint = in.DockerEndpoint
		out.RootDirectory = in.RootDirectory
		out.SeccompProfileRoot = in.SeccompProfileRoot
		out.AllowPrivileged = in.AllowPrivileged
		if in.HostNetworkSources != nil {
			in, out := &in.HostNetworkSources, &out.HostNetworkSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.HostNetworkSources = nil
		}
		if in.HostPIDSources != nil {
			in, out := &in.HostPIDSources, &out.HostPIDSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.HostPIDSources = nil
		}
		if in.HostIPCSources != nil {
			in, out := &in.HostIPCSources, &out.HostIPCSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.HostIPCSources = nil
		}
		out.RegistryPullQPS = in.RegistryPullQPS
		out.RegistryBurst = in.RegistryBurst
		out.EventRecordQPS = in.EventRecordQPS
		out.EventBurst = in.EventBurst
		out.EnableDebuggingHandlers = in.EnableDebuggingHandlers
		out.MinimumGCAge = in.MinimumGCAge
		out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
		out.MaxContainerCount = in.MaxContainerCount
		out.CAdvisorPort = in.CAdvisorPort
		out.HealthzPort = in.HealthzPort
		out.HealthzBindAddress = in.HealthzBindAddress
		out.OOMScoreAdj = in.OOMScoreAdj
		out.RegisterNode = in.RegisterNode
		out.ClusterDomain = in.ClusterDomain
		out.MasterServiceNamespace = in.MasterServiceNamespace
		out.ClusterDNS = in.ClusterDNS
		out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
		out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
		out.ImageMinimumGCAge = in.ImageMinimumGCAge
		out.ImageGCHighThresholdPercent = in.ImageGCHighThresholdPercent
		out.ImageGCLowThresholdPercent = in.ImageGCLowThresholdPercent
		out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
		out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
		out.NetworkPluginName = in.NetworkPluginName
		out.NetworkPluginMTU = in.NetworkPluginMTU
		out.NetworkPluginDir = in.NetworkPluginDir
		out.CNIConfDir = in.CNIConfDir
		out.CNIBinDir = in.CNIBinDir
		out.VolumePluginDir = in.VolumePluginDir
		out.CloudProvider = in.CloudProvider
		out.CloudConfigFile = in.CloudConfigFile
		out.KubeletCgroups = in.KubeletCgroups
		out.ExperimentalCgroupsPerQOS = in.ExperimentalCgroupsPerQOS
		out.CgroupDriver = in.CgroupDriver
		out.RuntimeCgroups = in.RuntimeCgroups
		out.SystemCgroups = in.SystemCgroups
		out.CgroupRoot = in.CgroupRoot
		out.ContainerRuntime = in.ContainerRuntime
		out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
		out.RemoteImageEndpoint = in.RemoteImageEndpoint
		out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
		out.RktPath = in.RktPath
		out.ExperimentalMounterPath = in.ExperimentalMounterPath
		out.RktAPIEndpoint = in.RktAPIEndpoint
		out.RktStage1Image = in.RktStage1Image
		out.LockFilePath = in.LockFilePath
		out.ExitOnLockContention = in.ExitOnLockContention
		out.HairpinMode = in.HairpinMode
		out.BabysitDaemons = in.BabysitDaemons
		out.MaxPods = in.MaxPods
		out.NvidiaGPUs = in.NvidiaGPUs
		out.DockerExecHandlerName = in.DockerExecHandlerName
		out.PodCIDR = in.PodCIDR
		out.ResolverConfig = in.ResolverConfig
		out.CPUCFSQuota = in.CPUCFSQuota
		out.Containerized = in.Containerized
		out.MaxOpenFiles = in.MaxOpenFiles
		out.ReconcileCIDR = in.ReconcileCIDR
		out.RegisterSchedulable = in.RegisterSchedulable
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.SerializeImagePulls = in.SerializeImagePulls
		out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
		out.NodeIP = in.NodeIP
		if in.NodeLabels != nil {
			in, out := &in.NodeLabels, &out.NodeLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.NodeLabels = nil
		}
		out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
		out.EnableCustomMetrics = in.EnableCustomMetrics
		out.EvictionHard = in.EvictionHard
		out.EvictionSoft = in.EvictionSoft
		out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
		out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
		out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
		out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
		out.ExperimentalKernelMemcgNotification = in.ExperimentalKernelMemcgNotification
		out.PodsPerCore = in.PodsPerCore
		out.EnableControllerAttachDetach = in.EnableControllerAttachDetach
		if in.SystemReserved != nil {
			in, out := &in.SystemReserved, &out.SystemReserved
			*out = make(config.ConfigurationMap)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.SystemReserved = nil
		}
		if in.KubeReserved != nil {
			in, out := &in.KubeReserved, &out.KubeReserved
			*out = make(config.ConfigurationMap)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.KubeReserved = nil
		}
		out.ProtectKernelDefaults = in.ProtectKernelDefaults
		out.MakeIPTablesUtilChains = in.MakeIPTablesUtilChains
		out.IPTablesMasqueradeBit = in.IPTablesMasqueradeBit
		out.IPTablesDropBit = in.IPTablesDropBit
		if in.AllowedUnsafeSysctls != nil {
			in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.AllowedUnsafeSysctls = nil
		}
		out.FeatureGates = in.FeatureGates
		out.EnableCRI = in.EnableCRI
		out.ExperimentalFailSwapOn = in.ExperimentalFailSwapOn
		out.ExperimentalCheckNodeCapabilitiesBeforeMount = in.ExperimentalCheckNodeCapabilitiesBeforeMount
		out.KeepTerminatedPodVolumes = in.KeepTerminatedPodVolumes
		return nil
	}
}

func DeepCopy_componentconfig_KubeletWebhookAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthentication)
		out := out.(*KubeletWebhookAuthentication)
		out.Enabled = in.Enabled
		out.CacheTTL = in.CacheTTL
		return nil
	}
}

func DeepCopy_componentconfig_KubeletWebhookAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthorization)
		out := out.(*KubeletWebhookAuthorization)
		out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
		out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
		return nil
	}
}

func DeepCopy_componentconfig_KubeletX509Authentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletX509Authentication)
		out := out.(*KubeletX509Authentication)
		out.ClientCAFile = in.ClientCAFile
		return nil
	}
}

func DeepCopy_componentconfig_LeaderElectionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LeaderElectionConfiguration)
		out := out.(*LeaderElectionConfiguration)
		out.LeaderElect = in.LeaderElect
		out.LeaseDuration = in.LeaseDuration
		out.RenewDeadline = in.RenewDeadline
		out.RetryPeriod = in.RetryPeriod
		return nil
	}
}

func DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PersistentVolumeRecyclerConfiguration)
		out := out.(*PersistentVolumeRecyclerConfiguration)
		out.MaximumRetry = in.MaximumRetry
		out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
		out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
		out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
		out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
		out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
		out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
		return nil
	}
}

func DeepCopy_componentconfig_PortRangeVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PortRangeVar)
		out := out.(*PortRangeVar)
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		} else {
			out.Val = nil
		}
		return nil
	}
}

func DeepCopy_componentconfig_VolumeConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*VolumeConfiguration)
		out := out.(*VolumeConfiguration)
		out.EnableHostPathProvisioning = in.EnableHostPathProvisioning
		out.EnableDynamicProvisioning = in.EnableDynamicProvisioning
		out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
		out.FlexVolumePluginDir = in.FlexVolumePluginDir
		return nil
	}
}
