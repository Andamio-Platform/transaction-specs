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

type Seed struct {
	BlockSlot uint64
	BlockHash string
}

type V2ConfigWithSeed struct {
	V2ConfigStruct // embedded — brings all fields of V2ConfigStruct directly
	Seed           Seed
}

func (c *Config) CurrentV2() V2ConfigWithSeed {
	if c.Network == constants.MAINNET {
		return V2ConfigWithSeed{
			V2ConfigStruct: c.V2.Mainnet,
			Seed: Seed{
				BlockSlot: 0,
				BlockHash: "",
			},
		}
	}
	return V2ConfigWithSeed{
		V2ConfigStruct: c.V2.Preprod,
		Seed: Seed{
			BlockSlot: 107466420,
			BlockHash: "76bc905e0c90761540a5bf504c2a9463ad075eed6284cb06540f4e1092f959c6",
		},
	}
}
