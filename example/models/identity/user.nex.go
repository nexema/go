package identity

import (
	"github.com/nexema/go/runtime"
	"io"
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

func (u *Account) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	encoder.EncodeVarint(u.fieldIndex)
	switch u.fieldIndex {

	case 0:

		adminBytes, err := u.Admin.Encode()
		if err != nil {
			return nil, err
		}
		encoder.EncodeBinary(adminBytes)

	case 1:

		customerBytes, err := u.Customer.Encode()
		if err != nil {
			return nil, err
		}
		encoder.EncodeBinary(customerBytes)

	}

	return encoder.TakeBytes(), nil
}

func (u *Account) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *Account) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	u.fieldIndex, err = decoder.DecodeVarint()
	if err != nil {
		return err
	}

	switch fieldIndex {
	case -1:
		u.value = nil

	case 0:

		adminBytes, err := decoder.DecodeBinary()
		if err != nil {
			return nil, err
		}
		err = u.Admin.Decode(adminBytes)
		if err != nil {
			return err
		}

	case 1:

		customerBytes, err := decoder.DecodeBinary()
		if err != nil {
			return nil, err
		}
		err = u.Customer.Decode(customerBytes)
		if err != nil {
			return err
		}

	}

	return nil
}

func (u *Account) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
