package config

import "time"

type (
	Config struct {
		Logger Logger `yaml:"LOGGER"`
		Nats   Nats   `yaml:"NATS"`
	}

	Logger struct {
		Level string `yaml:"LEVEL"`
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
