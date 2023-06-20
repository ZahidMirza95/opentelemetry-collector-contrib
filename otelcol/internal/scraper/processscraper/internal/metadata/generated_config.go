// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for hostmetricsreceiver/process metrics.
type MetricsConfig struct {
	ProcessContextSwitches     MetricConfig `mapstructure:"process.context_switches"`
	ProcessCPUTime             MetricConfig `mapstructure:"process.cpu.time"`
	ProcessCPUUtilization      MetricConfig `mapstructure:"process.cpu.utilization"`
	ProcessDiskIo              MetricConfig `mapstructure:"process.disk.io"`
	ProcessDiskOperations      MetricConfig `mapstructure:"process.disk.operations"`
	ProcessMemoryUsage         MetricConfig `mapstructure:"process.memory.usage"`
	ProcessMemoryUtilization   MetricConfig `mapstructure:"process.memory.utilization"`
	ProcessMemoryVirtual       MetricConfig `mapstructure:"process.memory.virtual"`
	ProcessOpenFileDescriptors MetricConfig `mapstructure:"process.open_file_descriptors"`
	ProcessPagingFaults        MetricConfig `mapstructure:"process.paging.faults"`
	ProcessSignalsPending      MetricConfig `mapstructure:"process.signals_pending"`
	ProcessThreads             MetricConfig `mapstructure:"process.threads"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		ProcessContextSwitches: MetricConfig{
			Enabled: false,
		},
		ProcessCPUTime: MetricConfig{
			Enabled: true,
		},
		ProcessCPUUtilization: MetricConfig{
			Enabled: false,
		},
		ProcessDiskIo: MetricConfig{
			Enabled: true,
		},
		ProcessDiskOperations: MetricConfig{
			Enabled: false,
		},
		ProcessMemoryUsage: MetricConfig{
			Enabled: true,
		},
		ProcessMemoryUtilization: MetricConfig{
			Enabled: false,
		},
		ProcessMemoryVirtual: MetricConfig{
			Enabled: true,
		},
		ProcessOpenFileDescriptors: MetricConfig{
			Enabled: false,
		},
		ProcessPagingFaults: MetricConfig{
			Enabled: false,
		},
		ProcessSignalsPending: MetricConfig{
			Enabled: false,
		},
		ProcessThreads: MetricConfig{
			Enabled: false,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for hostmetricsreceiver/process resource attributes.
type ResourceAttributesConfig struct {
	ProcessCommand        ResourceAttributeConfig `mapstructure:"process.command"`
	ProcessCommandLine    ResourceAttributeConfig `mapstructure:"process.command_line"`
	ProcessExecutableName ResourceAttributeConfig `mapstructure:"process.executable.name"`
	ProcessExecutablePath ResourceAttributeConfig `mapstructure:"process.executable.path"`
	ProcessOwner          ResourceAttributeConfig `mapstructure:"process.owner"`
	ProcessParentPid      ResourceAttributeConfig `mapstructure:"process.parent_pid"`
	ProcessPid            ResourceAttributeConfig `mapstructure:"process.pid"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		ProcessCommand: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessCommandLine: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessExecutableName: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessExecutablePath: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessOwner: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessParentPid: ResourceAttributeConfig{
			Enabled: true,
		},
		ProcessPid: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/process metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
