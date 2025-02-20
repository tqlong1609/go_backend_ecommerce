package setting

type Config struct {
	Server     ServerSetting     `mapstructure:"server"`
	Mysql      MySQLSetting      `mapstructure:"mysql"`
	PostgreSQL PostgreSQLSetting `mapstructure:"postgresql"`
	Logger     LoggerSetting     `mapstructure:"logger"`
	Redis      RedisSetting      `mapstructure:"redis"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type MySQLSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"maxIdleConnection"`
	MaxOpenConns int    `mapstructure:"maxOpenConnection"`
	MaxLifeTime  int    `mapstructure:"maxLifeTime"`
}

type PostgreSQLSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"maxIdleConnection"`
	MaxOpenConns int    `mapstructure:"maxOpenConnection"`
	MaxLifeTime  int    `mapstructure:"maxLifeTime"`
}

type LoggerSetting struct {
	Log_level   string `mapstructure:"log_level"`
	File_path   string `mapstructure:"file_path"`
	Max_size    int    `mapstructure:"max_size"`
	Max_backups int    `mapstructure:"max_backups"`
	Max_age     int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
