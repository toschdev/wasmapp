package types

const (
	// ModuleName defines the module name
	ModuleName = "wasmapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wasmapp"
)

var (
	ParamsKey = []byte("p_wasmapp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
