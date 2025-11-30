package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type V2Config struct {
	Mainnet V2ConfigStruct `json:"mainnet"`
	Preprod V2ConfigStruct `json:"preprod"`
}

type V2ConfigStruct struct {
	CourseGovernanceV2 struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"courseGovernanceV2"`
	GlobalStateV2S struct {
		SCAddress string `json:"sCAddress"`
		SCTxRef   string `json:"sCTxRef"`
	} `json:"globalStateV2S"`
	IndexAdmin string `json:"indexAdmin"`
	IndexMS    struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"indexMS"`
	IndexRefMS struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"indexRefMS"`
	IndexStakingSh       string `json:"indexStakingSh"`
	InitIndexPolId       string `json:"initIndexPolId"`
	InstanceAdmin        string `json:"instanceAdmin"`
	InstanceGovernanceV2 struct {
		SCAddress string `json:"sCAddress"`
		SCTxRef   string `json:"sCTxRef"`
	} `json:"instanceGovernanceV2"`
	InstanceProvidedV2MS struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"instanceProvidedV2MS"`
	InstanceStakingScrSh string `json:"instanceStakingScrSh"`
	InstanceV2MS         struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"instanceV2MS"`
	LocalStateRef struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"localStateRef"`
	ModuleScriptsV2 struct {
		MSCAddress  string `json:"mSCAddress"`
		MSCPolicyID string `json:"mSCPolicyID"`
		MSCTxRef    string `json:"mSCTxRef"`
	} `json:"moduleScriptsV2"`
	ReferenceAddr         string `json:"referenceAddr"`
	V2GlobalStateObsTxRef string `json:"v2GlobalStateObsTxRef"`
}

func loadV2Config() V2Config {
	paths := []string{
		"./config/v2-preprod.json",
		"../../config/v2-preprod.json",
	}

	var data []byte
	var err error
	var loadedPath string

	for _, path := range paths {
		data, err = os.ReadFile(path)
		if err == nil {
			loadedPath = path
			break
		}
	}

	if err != nil {
		panic(fmt.Sprintf("failed to read V2 config file. Tried paths: %v. Last error: %v", paths, err))
	}

	var preprodConfig V2ConfigStruct
	if err := json.Unmarshal(data, &preprodConfig); err != nil {
		panic(fmt.Sprintf("failed to parse V2 config file at %s: %v", loadedPath, err))
	}

	return V2Config{
		Mainnet: V2ConfigStruct{},
		Preprod: preprodConfig,
	}
}
