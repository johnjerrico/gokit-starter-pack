package vault

//IVault ...
type IVault interface {
	GetEnvOrDefaultConfig(path string, def interface{})
	WriteEncrypted(transitkey, path, value []byte)
	ReadEncrypted(transitkey, path string)
}
