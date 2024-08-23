package entity

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}
type AppConfig struct {
	Port string `yaml:"port"`
}
type DBConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"dbname"`
}

type Product struct {
	Id       int `gorm:"primaryKey;autoIncrement:true"`
	Name     string	`json:"name"`
	Category string `json:"category"`
	Price    int `json:"price"`
	Stock    int `json:"stock"`
}
