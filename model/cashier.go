package model

import "time"

type Cashier struct {
	CashierId int        `db:"cashierId" json:"cashierId,omitempty"`
	Name      string     `db:"name" json:"name,omitempty"`
	Passcode  string     `db:"passcode" json:"passcode,omitempty"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	IsDelete  int        `db:"isDelete" json:"isDelete,omitempty"`
}
