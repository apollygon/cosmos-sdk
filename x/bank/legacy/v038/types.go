package v038

// DONTCOVER
// nolint

const (
	ModuleName = "bank"
)

type (
	GenesisState struct {
		SendEnabled bool `json:"send_enabled" yaml:"send_enabled"`
	}
)
