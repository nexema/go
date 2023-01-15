package identity

import (
	"github.com/example/models"
	"github.com/nexema/go/runtime"
	"io"
)

type User struct {
	Id string

	Name string

	SelectedColor models.Colors

	Account runtime.Nullable[CustomerAccount]

	Tags []string

	Claims map[string]runtime.Nullable[string]
}

func (u *User) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error

	encoder.EncodeString(u.Id)

	encoder.EncodeString(u.Name)

	encoder.EncodeUint8(u.SelectedColor.Index())

	accountBytes, err := u.Account.Encode()
	if err != nil {
		return nil, err
	}
	encoder.EncodeBinary(accountBytes)

	return encoder.TakeBytes(), nil
}

func (u *User) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *User) Decode(buffer []byte) error {
	return nil
}

func (u *User) MustDecode(buffer []byte) {
	err := u.Decode(buffer)
	if err != nil {
		panic(err)
	}
}

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

type AdminAccount struct {
	Active bool
}

func (u *AdminAccount) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error

	encoder.EncodeBool(u.Active)

	return encoder.TakeBytes(), nil
}

func (u *AdminAccount) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *AdminAccount) Decode(buffer []byte) error {
	return nil
}

func (u *AdminAccount) MustDecode(buffer []byte) {
	err := u.Decode(buffer)
	if err != nil {
		panic(err)
	}
}

type CustomerAccount struct {
	Age runtime.Nullable[uint8]
}

func (u *CustomerAccount) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error

	if u.Age.IsNull() {
		encoder.EncodeNull()
	} else {
		age := *u.Age.Value
		encoder.EncodeUint8(age)
	}

	return encoder.TakeBytes(), nil
}

func (u *CustomerAccount) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *CustomerAccount) Decode(buffer []byte) error {
	return nil
}

func (u *CustomerAccount) MustDecode(buffer []byte) {
	err := u.Decode(buffer)
	if err != nil {
		panic(err)
	}
}
