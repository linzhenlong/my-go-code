package config

// Web...
type Web struct {
	Path string `toml:"path"`
	Pass string `toml:"pass"`
}

// WebList ...
type WebList struct {
	Name string `toml:"name"`
	Web  []Web  `toml:"web"`
}

// Proxy ...
type Proxy struct {
	Name string `toml:"name"`
}

// TomlConfig ...
type TomlConfig struct {
	Proxy   Proxy   `toml:"proxy"`
	WebList WebList `toml:"webList"`
}
