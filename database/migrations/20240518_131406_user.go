package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20240518_131406 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20240518_131406{}
	m.Created = "20240518_131406"

	migration.Register("User_20240518_131406", m)
}

// Run the migrations
func (m *User_20240518_131406) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY," +
		"name VARCHAR(200) NOT NULL," +
		"email VARCHAR(100) NOT NULL UNIQUE," +
		"phone VARCHAR(100) NOT NULL UNIQUE," +
		"password VARCHAR(200) NOT NULL)")

}

// Reverse the migrations
func (m *User_20240518_131406) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user")
}
