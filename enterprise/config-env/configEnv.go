package configenv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jeremywohl/flatten"
	"github.com/tidwall/sjson"

	"github.com/rudderlabs/rudder-server/config"
	"github.com/rudderlabs/rudder-server/utils/logger"
)

type HandleT struct{}

var (
	configEnvReplacer string
	pkgLogger         logger.LoggerI
)

func loadConfig() {
	configEnvReplacer = config.GetString("BackendConfig.configEnvReplacer", "env.")
}

// ReplaceConfigWithEnvVariables : Replaces all env variables in the config
func (handle *HandleT) ReplaceConfigWithEnvVariables(workspaceConfig []byte) (updatedConfig []byte) {
	configMap := make(map[string]interface{}, 0)

	err := json.Unmarshal(workspaceConfig, &configMap)
	if err != nil {
		pkgLogger.Error("[ConfigEnv] Error while parsing request", err, string(workspaceConfig))
		return workspaceConfig
	}

	flattenedConfig, err := flatten.Flatten(configMap, "", flatten.DotStyle)
	if err != nil {
		pkgLogger.Errorf("[ConfigEnv] Failed to flatten workspace config: %v", err)
		return workspaceConfig
	}

	for configKey, v := range flattenedConfig {
		reflectType := reflect.TypeOf(v)
		if reflectType != nil && reflectType.String() == "string" {
			valString := v.(string)
			shouldReplace := strings.HasPrefix(strings.TrimSpace(valString), configEnvReplacer)
			if shouldReplace {
				envVariable := valString[len(configEnvReplacer):]
				envVarValue := config.GetEnv(envVariable, "")
				if envVarValue == "" {
					errorMessage := fmt.Sprintf("[ConfigEnv] Missing envVariable: %s. Either set it as envVariable or remove %s from the destination config.", envVariable, configEnvReplacer)
					pkgLogger.Errorf(errorMessage)
					continue
				}
				workspaceConfig, err = sjson.SetBytes(workspaceConfig, configKey, envVarValue)
				if err != nil {
					pkgLogger.Error("[ConfigEnv] Failed to set config for %s", configKey)
				}
			}
		}
	}

	return workspaceConfig
}
