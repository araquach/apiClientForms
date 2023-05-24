package models

import "time"

type SkinTest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	Salon     uint      `json:"salon"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	TestDone  bool      `json:"test_done"`
	Accept    bool      `json:"accept"`
	Email     string    `json:"email"`
	Signature string    `json:"signature"`
}

type ApiSkinTest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Salon     uint      `json:"salon"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type Extensions struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time ` json:"created_at"`
	Type        string    `json:"type"`
	Salon       uint      `json:"salon"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Final       bool      `json:"final"`
	Payment     bool      `json:"payment"`
	Maintenance bool      `json:"maintenance"`
	RemovalCost bool      `json:"removal_cost"`
	PreBook     bool      `json:"pre_book"`
	Email       string    `json:"email"`
	Mobile      string    `json:"mobile"`
	Signature   string    `json:"signature"`
}

type ApiExtensions struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Salon     uint      `json:"salon"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
