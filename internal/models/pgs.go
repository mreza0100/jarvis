package models

type PostgresConnConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// runnerport.PostgresConfig{
// 	Host:     "30.0.0.10",
// 	Port:     5432,
// 	Username: "mamad",
// 	Password: "mamadspass",
// 	Database: "mamad_db",
// }
