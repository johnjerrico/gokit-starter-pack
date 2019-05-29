package vault

//IVault ...
type IVault interface {
	New(addr, token string)
	GetEnvOrDefaultConfig(def interface{}, path string)
	WriteEncrypted(transitkey, path, key, value string)
	ReadEncrypted(transitkey, path, key string)
}
