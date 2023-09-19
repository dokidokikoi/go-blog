package config

type Oss struct {
	Host       string
	ID         string
	Secret     string
	UploadDir  string `mapstructure:"upload_dir"`
	ExpireTime int64  `mapstructure:"expire_time"`
}

const (
	ossKey string = "oss"
)

var OssConfig = &Oss{}

func init() {
	ossConfig := GetSpecConfig(ossKey)
	if ossConfig != nil {
		ossConfig.Unmarshal(OssConfig)
	}
}
