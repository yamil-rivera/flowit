package configs

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// ValidateViperConfig takes a viper configuration and validates it section by section
func unmarshallConfig() (Flowit, error) {

	var flowit Flowit

	config := func(c *mapstructure.DecoderConfig) {
		c.ErrorUnused = true
		c.WeaklyTypedInput = false
	}

	err := viper.UnmarshalKey("flowit", &flowit, config)
	if err != nil {
		return flowit, errors.Wrap(err, "Validation error")
	}
	return flowit, nil
}
