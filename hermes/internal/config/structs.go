package config

import "time"

type (
	Config struct {
		App      App            `yaml:"APP"`
		Timezone *time.Location `yaml:"TIMEZONE"`
		Location string         `yaml:"LOCATION"`
		Logger   Logger         `yaml:"LOGGER"`
		Nats     Nats           `yaml:"NATS"`
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

	Nats struct {
		Address        string        `yaml:"ADDRESS"`
		ConnectWait    time.Duration `yaml:"CONNECT_WAIT"`
		DialTimeout    time.Duration `yaml:"DIAL_TIMEOUT"`
		FlushTimeout   time.Duration `yaml:"FLUSH_TIMEOUT"`
		FlusherTimeout time.Duration `yaml:"FLUSHER_TIMEOUT"`
		PingInterval   time.Duration `yaml:"PING_INTERVAL"`
		ConnectBufSize int           `yaml:"CONNECT_BUFFER_SIZE"`
		MaxChanLen     int           `yaml:"MAX_CHAN_LENGTH"`
		MaxPingOut     int           `yaml:"MAX_PING_OUT"`
	}
)
