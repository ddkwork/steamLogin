package conf

// config struct
type Config struct {
	Server struct {
		Name       string `yaml:"name"` //App 名称
		ServerName string `yaml:"server_name"`
		Address    string `yaml:"address"`  // IP地址
		Port       string `yaml:"port"`     //端口
		SiteUrl    string `yaml:"site_url"` //websocket url
	} `yaml:"server"`
}
