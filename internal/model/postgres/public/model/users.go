//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Users struct {
	ID       int32 `sql:"primary_key" db:"id" json:"idm,omitempty"`
	Login    string `db:"login" json:"login"`
	Password []byte `db:"password" json:"password,omitempty"`
}
