package config

import "time"

type (
	Config struct {
		App        App            `yaml:"APP"`
		Timezone   *time.Location `yaml:"TIMEZONE"`
		Location   string         `yaml:"LOCATION"`
		Logger     Logger         `yaml:"LOGGER"`
		HTTPServer HTTPServer     `yaml:"HTTP_SERVER"`
		Database   Database       `yaml:"DATABASE" json:"database"`
		Redis      Redis          `yaml:"REDIS"`
		//Nats       Nats           `yaml:"NATS"`
	}

	App struct {
		Environment string `yaml:"ENVIRONMENT" validate:"required,oneof=production development"`
	}

	Logger struct {
		Level         string `yaml:"LEVEL"`
		StoutLogLevel int8   `yaml:"STDOUT_LOG_LEVEL"`
		ELKLogLevel   int8   `yaml:"ELK_LOG_LEVEL"`
		Path          string `yaml:"PATH"`
	}

	HTTPServer struct {
		Listen            string        `yaml:"LISTEN"`
		ReadTimeout       time.Duration `yaml:"READ_TIMEOUT"`
		WriteTimeout      time.Duration `yaml:"WRITE_TIMEOUT"`
		ReadHeaderTimeout time.Duration `yaml:"READ_HEADER_TIMEOUT"`
		IdleTimeout       time.Duration `yaml:"IDLE_TIMEOUT"`
	}

	Database struct {
		Driver      string        `yaml:"DRIVER"`
		Host        string        `yaml:"HOST" json:"host"`
		Port        int           `yaml:"PORT"`
		Name        string        `yaml:"NAME"`
		User        string        `yaml:"USER"`
		Password    string        `yaml:"PASSWORD"`
		DialRetry   int           `yaml:"DIAL_RETRY"`
		MaxIdle     int           `yaml:"MAX_IDLE"`
		MaxConn     int           `yaml:"MAX_CONN"`
		IdleConn    int           `yaml:"IDLE_CONN"`
		DialTimeout time.Duration `yaml:"DIAL_TIMEOUT"`
	}

	Redis struct {
		Host        string        `yaml:"HOST"`
		Port        string        `yaml:"PORT"`
		DB          int           `yaml:"DB"`
		Username    string        `yaml:"USERNAME"`
		Password    string        `yaml:"PASSWORD"`
		DialTimeout time.Duration `yaml:"DIAL_TIMEOUT"`
	}

	//Nats struct {
	//	AuthEnabled    bool          `yaml:"AUTH_ENABLED" validate:"boolean"`
	//	Host           string        `yaml:"HOST"`
	//	Port           string        `yaml:"PORT"`
	//	ConnectWait    time.Duration `yaml:"CONNECT_WAIT"`
	//	DialTimeout    time.Duration `yaml:"DIAL_TIMEOUT"`
	//	FlushTimeout   time.Duration `yaml:"FLUSH_TIMEOUTt"`
	//	FlusherTimeout time.Duration `yaml:"FLUSHER_TIMEOUTt"`
	//	PingInterval   time.Duration `yaml:"PING_INTERVAL"`
	//	ConnectBufSize int           `yaml:"CONNECT_BUF_SIZR"`
	//	MaxChanLen     int           `yaml:"MAX_CHAN_LEN"`
	//	MaxPingOut     int           `yaml:"MAX_PING_OUT"`
	//	MaxReconnect   int           `yaml:"MAX_RECONNECT"`
	//	ClusterID      string        `yaml:"CLUSTER_ID"`
	//	ClientName     string        `yaml:"CLIENT_NAME"`
	//	Username       string        `yaml:"USERNAME"`
	//	Password       string        `yaml:"PASSWORD"`
	//}
)
