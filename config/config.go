package config

// Values is an instance of config.
var Values Config

// Config interface.
type Config interface {
	MySQL() string  // MySQLのURIをかえすよ
	DBName() string // Database名
}

// Init initalize Values.
func Init(env string) {
	if Values != nil {
		panic("already initialized")
	}
	switch env {
	case "test":
		Values = TestConfig{}
	default:
		Values = TestConfig{}
	}
}

// TestConfig is config for test.
type TestConfig struct{}

// MySQL to implement Config interface.
func (c TestConfig) MySQL() string {
	return "travis@"
}

// DBName to implement Config interface.
func (c TestConfig) DBName() string {
	return "animapi_test"
}
