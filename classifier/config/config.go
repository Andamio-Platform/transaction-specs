package config

import (
	"sync"

	"github.com/Salvionied/apollo/constants"
)

var (
	instance *Config
	once     sync.Once

	// courseStatePolicyIds is managed separately to allow updates
	courseStatePolicyIds      []string
	courseStatePolicyIdsMutex sync.RWMutex
)

type Config struct {
	// V1          V1Config
	V2      V2Config
	Network constants.Network
}

// Init initializes the singleton config instance
func Init(network constants.Network) {
	once.Do(func() {
		instance = load(network)
	})
}

// SetCourseStatePolicyIds updates the course state policy IDs
func SetCourseStatePolicyIds(ids []string) {
	courseStatePolicyIdsMutex.Lock()
	defer courseStatePolicyIdsMutex.Unlock()
	courseStatePolicyIds = ids
}

// GetCourseStatePolicyIds returns the current course state policy IDs
func GetCourseStatePolicyIds() []string {
	courseStatePolicyIdsMutex.RLock()
	defer courseStatePolicyIdsMutex.RUnlock()
	return courseStatePolicyIds
}

// Get returns the singleton config instance
func Get() *Config {
	if instance == nil {
		// Fallback or panic if not initialized?
		// For now, let's assume Init is always called.
		// Or we can lazy load with defaults if Init wasn't called, but the user requirement implies explicit Init.
		// Let's panic if not initialized to enforce the pattern.
		panic("Config not initialized. Call config.Init() first.")
	}
	return instance
}

func load(network constants.Network) *Config {

	v2Config := loadV2Config()

	return &Config{
		V2:      v2Config,
		Network: network,
	}
}

func (c *Config) CurrentV2() V2ConfigStruct {
	if c.Network == constants.MAINNET {
		return c.V2.Mainnet
	}
	return c.V2.Preprod
}
