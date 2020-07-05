package dbcon

type DBConfig struct {
	Driver   string
	User     string
	Password string
	Server   string
	Port     uint
	DBName   string
	SSLMode  string
}
