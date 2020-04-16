package config

const (
	LANG = 2 // 1 CN; 2 EN;
)

func GetApolloString(key string, defValue string) string {
	return defValue
}

func GetApolloInt(key string, defValue int) int {
	return defValue
}
