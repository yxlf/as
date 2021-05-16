package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (s *Setting, e error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("configs/")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{v}, nil
}
