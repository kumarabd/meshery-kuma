package config

import (
	"fmt"
	"io/ioutil"

	"github.com/layer5io/gokit/utils"

	"github.com/spf13/viper"
)

// Viper instance for configuration
type Viper struct {
	instance *viper.Viper
}

// NewViper intializes a viper instance and dependencies
func NewViper() (Handler, error) {

	err := ioutil.WriteFile(fmt.Sprintf("%s%s", filepath, filename), []byte(""), 0600)
	if err != nil {
		return nil, ErrViper(err)
	}

	v := viper.New()
	v.AddConfigPath(filepath)
	v.SetConfigType(filetype)
	v.SetConfigName(filename)
	v.AutomaticEnv()

	v.SetDefault("server", server)
	v.SetDefault("mesh", mesh)
	v.SetDefault("instance", instance)
	v.SetDefault("operations", operations)
	err = v.WriteConfig()
	if err != nil {
		return nil, ErrViper(err)
	}

	return &Viper{
		instance: v,
	}, nil
}

// SetKey sets a key value in viper
func (v *Viper) SetKey(key string, value string) error {
	v.instance.Set(key, value)
	err := v.instance.WriteConfig()
	if err != nil {
		return ErrViper(err)
	}
	return nil
}

// GetKey gets a key value from viper
func (v *Viper) GetKey(key string) (string, error) {
	err := v.instance.ReadInConfig()
	if err != nil {
		return " ", ErrViper(err)
	}

	s, err := utils.Marshal(v.instance.Get(key))
	if err != nil {
		return " ", ErrViper(err)
	}
	return s, nil
}

// Server provides server specific configuration
func (v *Viper) Server(result interface{}) error {
	s, err := v.GetKey("server")
	if err != nil {
		return ErrViper(err)
	}
	return utils.Unmarshal(s, &result)
}

// MeshSpec provides mesh specific configuration
func (v *Viper) MeshSpec(result interface{}) error {
	s, err := v.GetKey("mesh")
	if err != nil {
		return ErrViper(err)
	}
	return utils.Unmarshal(s, &result)
}

// MeshInstance provides mesh specific configuration
func (v *Viper) MeshInstance(result interface{}) error {
	s, err := v.GetKey("instance")
	if err != nil {
		return ErrViper(err)
	}
	return utils.Unmarshal(s, &result)
}

// Operations provides list of operations available
func (v *Viper) Operations(result interface{}) error {
	s, err := v.GetKey("operations")
	if err != nil {
		return ErrViper(err)
	}
	return utils.Unmarshal(s, &result)
}
