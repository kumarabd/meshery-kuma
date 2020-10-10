package config

import (
	"fmt"
	"os"

	"github.com/layer5io/gokit/utils"
)

var (

	// server holds server configuration
	server = map[string]string{
		"name":    "kuma-adapter",
		"port":    "10007",
		"version": "v1.0.0",
	}

	// mesh holds mesh configuration
	mesh = map[string]string{
		"name":     "kuma",
		"status":   "not installed",
		"traceurl": os.Getenv("TRACING_ENDPOINT"),
		"version":  "0.6.0",
	}

	instance = map[string]string{
		"installmode":     "flat",
		"installplatform": "kubernetes",
		"installzone":     " ",
		"mgmtaddr":        "0.0.0.0:8000",
		"kumaaddr":        "5681",
	}

	// operations holds the supported operations inside mesh
	operations = map[string]interface{}{
		InstallKumav071: map[string]interface{}{
			"type": "0",
			"properties": map[string]string{
				"description": "Kuma service mesh (0.7.1)",
				"version":     "0.7.1",
			},
		},
		InstallKumav070: map[string]interface{}{
			"type": "0",
			"properties": map[string]string{
				"description": "Kuma service mesh (0.7.0)",
				"version":     "0.7.0",
			},
		},
		InstallKumav060: map[string]interface{}{
			"type": "0",
			"properties": map[string]string{
				"description": "Kuma service mesh (0.6.0)",
				"version":     "0.6.0",
			},
		},
		InstallSampleBookInfo: map[string]interface{}{
			"type": "1",
			"properties": map[string]string{
				"description": "BookInfo Application",
				"version":     "latest",
			},
		},
		ValidateSmiConformance: map[string]interface{}{
			"type": "3",
			"properties": map[string]string{
				"description": "SMI conformance Test",
				"version":     "latest",
			},
		},
	}

	// Viper configuration
	filepath = fmt.Sprintf("%s/.meshery/kuma/", utils.GetHome())
	filename = "config"
	filetype = "yaml"
)
