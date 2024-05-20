package httputil

import (
	"net/http"
	"strings"
)

// MonitoringConfig structure of data for handling configuration for
// controlling what content is monitored.
type MonitoringConfig struct {
	Types                []string `export:"true"   json:"types,omitempty"   toml:"types,omitempty"      yaml:"types,omitempty"`
	Methods              []string `export:"true"   json:"methods,omitempty" toml:"methods,omitempty"    yaml:"methods,omitempty"`
	CheckMimeAccept      bool     `default:"false" export:"true"            json:"checkMimeAccept"      toml:"checkMimeAccept"      yaml:"checkMimeAccept"`
	CheckMimeContentType bool     `default:"true"  export:"true"            json:"checkMimeContentType" toml:"checkMimeContentType" yaml:"checkMimeContentType"`
	CheckAcceptEncoding  bool     `default:"true"  export:"true"            json:"checkAcceptEncoding"  toml:"checkAcceptEncoding"  yaml:"checkAcceptEncoding"`
	CheckContentEncoding bool     `default:"true"  export:"true"            json:"checkContentEncoding" toml:"checkContentEncoding" yaml:"checkContentEncoding"`
}

// CreateMonitoringConfig creates and initializes the monitoring configuration.
func CreateMonitoringConfig() *MonitoringConfig {
	config := MonitoringConfig{
		Types:                nil,
		Methods:              nil,
		CheckMimeAccept:      true,
		CheckMimeContentType: true,
		CheckAcceptEncoding:  true,
		CheckContentEncoding: true,
	}
	config.EnsureDefaults()

	return &config
}

// EnsureDefaults check Types and Methods for empty arrays and apply default values if found.
func (config *MonitoringConfig) EnsureDefaults() {
	if len(config.Methods) == 0 {
		config.Methods = []string{http.MethodGet}
	}

	if len(config.Types) == 0 {
		config.Types = []string{"text/html"}
	}
}

// EnsureProperFormat handle weird yaml parsing until the underlying issue can be resolved.
func (config *MonitoringConfig) EnsureProperFormat() {
	if len(config.Methods) == 1 && strings.HasPrefix(config.Methods[0], "║24║") {
		config.Methods = strings.Split(strings.ReplaceAll(config.Methods[0], "║24║", ""), "║")
	}

	if len(config.Types) == 1 && strings.HasPrefix(config.Types[0], "║24║") {
		config.Types = strings.Split(strings.ReplaceAll(config.Types[0], "║24║", ""), "║")
	}
}
