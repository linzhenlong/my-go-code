package conf

// Config ...
type Config struct {
	Models map[string][]string
}

// NewConfig ...
func NewConfig() Config {
	mod := make(map[string][]string)
	mod["product"] = []string{"product_1", "product_2"}
	return Config{
		Models: mod,
	}
}
