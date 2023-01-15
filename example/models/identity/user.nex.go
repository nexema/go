package identity

import (
	"github.com/nexema/go/runtime"
)

type Account struct {
	value      interface{}
	fieldIndex int64
}

type AccountWhichField int8

const (
	Account_NotSet AccountWhichField = -1

	Account_Admin AccountWhichField = 0

	Account_Customer AccountWhichField = 1
)

type accountBuilder struct{}

var AccountBuilder *accountBuilder = &accountBuilder{}

func (u *Account) IsSet() bool {
	return u.fieldIndex != -1
}

func (u *Account) Clear() {
	u.value = nil
	u.fieldIndex = -1
}

func (u *Account) WhichField() AccountWhichField {
	return AccountWhichField(u.fieldIndex)
}

func (u *Account) CurrentValue() interface{} {
	return u.value
}

func (u *Account) MergeFrom(buffer []byte) error {
	return nil
}

func (u *Account) MergeUsing(other *Account) error {
	u.fieldIndex = other.fieldIndex
	u.value = other.value
	return nil
}

func (u *Account) Clone() *Account {
	return &Account{
		fieldIndex: u.fieldIndex,
		value:      u.value,
	}
}

func (*accountBuilder) Admin(value CustomerAccount) *Account {
	return &Account{
		value:      value,
		fieldIndex: 0,
	}
}

func (u *Account) SetAdmin(value CustomerAccount) {
	u.value = value
	u.fieldIndex = 0
}

func (*accountBuilder) Customer(value CustomerAccount) *Account {
	return &Account{
		value:      value,
		fieldIndex: 1,
	}
}

func (u *Account) SetCustomer(value CustomerAccount) {
	u.value = value
	u.fieldIndex = 1
}
