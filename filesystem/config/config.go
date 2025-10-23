package config

type MinioConfig struct {
	AccessKey    string
	AccessSecret string
	Domain       string
	Endpoint     string
	Bucket       string
}

type OssConfig struct {
	AccessKey    string
	AccessSecret string
	Domain       string
	Endpoint     string
	Region       string
	Bucket       string
}

type Config struct {
	Driver string
	Oss    *OssConfig
	Minio  *MinioConfig
}
