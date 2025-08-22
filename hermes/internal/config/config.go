package config

import (
	"fmt"
	"hermes/log"
	"reflect"
	"strings"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var Cfg *Config

func Init(configPath string) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SMS_GATEWAY")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configPath == "" {
		configPath = "config.defaults.yaml"
	}
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		log.Fatalf("Error in reading config, error: %+v", err.Error())
	}

	// Initialize Cfg if it's nil
	if Cfg == nil {
		Cfg = &Config{}
	}

	hooks := []mapstructure.DecodeHookFunc{
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		TimeLocationDecodeHook(),
	}

	decoderConfig := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &Cfg,
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.ComposeDecodeHookFunc(hooks...),
		TagName:          "yaml",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		log.Fatalf("Error creating decoder, error: %+v", err.Error())
	}

	err = decoder.Decode(viper.AllSettings())
	if err != nil {
		log.Fatalf("Error in unmarshaling config, error: %+v", err.Error())
	}

	if Cfg.Logger.Level == "debug" {
		fmt.Printf("%#v \n", Cfg)
	}

	err = validator.New().Struct(Cfg)
	if err != nil {
		log.Fatalf("Error in validating config, error: %+v", err.Error())
	}

	log.Infof("using config: %s", viper.ConfigFileUsed())
}

func TimeLocationDecodeHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		var timeLocation *time.Location
		if t != reflect.TypeOf(timeLocation) {
			return data, nil
		}

		return time.LoadLocation(data.(string))
	}
}
