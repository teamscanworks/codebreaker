package types

// todo: auto-generate types from json

// Chain Cosmos Chain.json is a metadata file that contains information about a cosmos sdk based
// chain.
type Chain struct {
	Bech32Prefix string              `json:"bech32_prefix"`
	ChainID      string              `json:"chain_id"`
	ChainName    string              `json:"chain_name"`
	Codebase     *Codebase           `json:"codebase,omitempty"`
	Genesis      *Genesis            `json:"genesis,omitempty"`
	NetworkType  *NetworkType        `json:"network_type,omitempty"`
	PrettyName   *string             `json:"pretty_name,omitempty"`
	Status       *Status             `json:"status,omitempty"`
	Codes        map[string]Contract `json:"codes,omitempty"`
}

// Contract json is a metadata file that contains information about CosmWasm contract code
// in Cosmos chains.
type Contract struct {
	CodeID          int    `json:"code_id"`
	Checksum        string `json:"checksum"`
	BuildInfo       string `json:"build_info"`
	BuildEnv        string `json:"build_env"`
	ModuleName      string `json:"module_name"`
	Repository      string `json:"repository"`  // todo: refactor like Genesis type
	ReleaseTag      string `json:"release_tag"` // todo: refactor like Genesis type
	Org             string `json:"org,omitempty"`
	SecurityContact string `json:"security_contact,omitempty"`
	Website         string `json:"website,omitempty"`
}

type Codebase struct {
	GitRepo            string `json:"git_repo"`
	RecommendedVersion string `json:"recommended_version"`
}

type Genesis struct {
	GenesisURL *string `json:"genesis_url,omitempty"`
}

type NetworkType string

const (
	Mainnet NetworkType = "mainnet"
	Testnet NetworkType = "testnet"
)

type Status string

const (
	Killed   Status = "killed"
	Live     Status = "live"
	Upcoming Status = "upcoming"
)
