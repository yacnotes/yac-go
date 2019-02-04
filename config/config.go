package config

import (
	"os"
	"strconv"
	"yac-go/log"
)

type Config struct {
	Env         string `json:"env"`
	Port        int    `json:"port"`
	DatabaseDir string `json:"dbDir""`
	Log         struct {
		Stdout bool `json:"stdout"`
		Level  int  `json:"level"`
	} `json:"log"`
	AppInfo struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
}

func (c *Config) LoadFromEnv(version string) {
	c.AppInfo.Version = version
	c.AppInfo.Name = "yac-go"

	c.Env = getString("YAC_ENV", "debug")
	c.Port = getInt("YAC_PORT", 3000)
	c.DatabaseDir = getString("YAC_DB_DIR", "./.data/db")

	c.Log.Stdout = getBool("YAC_LOG_TO_STDOUT", true)
	c.Log.Level = getInt("YAC_LOG_LEVEL", log.LevelDebug)
}

func getInt(key string, def int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		val = def
	}
	return val
}

func getBool(key string, def bool) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		val = def
	}
	return val
}

func getString(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}
