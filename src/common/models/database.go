package models

// Database ...
type Database struct {
	Type       string      `json:"type"`
	PostGreSQL *PostGreSQL `json:"postgresql,omitempty"`
}

// PostGreSQL ...
type PostGreSQL struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	Database     string `json:"database"`
	SSLMode      string `json:"sslmode"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}
