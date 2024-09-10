package types

const (
	// ModuleName defines the module name
	ModuleName = "throughput"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_throughput"
)

var (
	ParamsKey = []byte("p_throughput")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
