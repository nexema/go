package identity

import (
	"github.com/example/models"
	"github.com/nexema/go/runtime"
	"io"
)

type User struct {
	Id            string
	Name          string
	SelectedColor models.Colors
	Account       runtime.Nullable[CustomerAccount]
	Tags          []string
	Claims        map[string]runtime.Nullable[string]
}

func (u *User) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	encoder.EncodeString(u.Id)

	encoder.EncodeString(u.Name)

	encoder.EncodeUint8(u.SelectedColor.Index())

	if u.Account.IsNull() {
		encoder.EncodeNull()
	} else {
		account := *u.Account.Value

		accountBytes, err := account.Encode()
		if err != nil {
			return nil, err
		}

		encoder.EncodeBinary(accountBytes)

	}

	encoder.BeginArray(int64(len(u.Tags)))
	for _, value := range u.Tags {

		encoder.EncodeString(value)

	}

	encoder.BeginMap(int64(len(u.Claims)))
	for key, value := range u.Claims {
		encoder.EncodeString(key)

		if value.IsNull() {
			encoder.EncodeNull()
		} else {
			encoder.EncodeString(*value.Value)
		}

	}

	return encoder.TakeBytes(), nil
}

func (u *User) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u User) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	u.Id, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.Name, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	selectedColorEnumIndex, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.SelectedColor = models.ColorsPicker.ByIndex(selectedColorEnumIndex)

	if decoder.IsNextNull() {
		u.Account.Clear()
	} else {

		accountBytes, err := decoder.DecodeBinary()
		if err != nil {
			return err
		}

		err = u.Account.Decode(accountBytes)
		if err != nil {
			return err
		}

	}

	tagsArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}

	u.Tags = make([]string, tagsArrayLen)
	for i := int64(0); i < tagsArrayLen; i++ {
		u.Tags[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}

	claimsMapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Claims = make(map[string]runtime.Nullable[string], claimsMapLen)
	for i := int64(0); i < claimsMapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		value, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.Claims[key] = value
	}

	return nil
}

func (u User) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
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
	Account_NotSet   AccountWhichField = -1
	Account_Admin    AccountWhichField = 0
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

		adminBytes, err := admin.Encode()
		if err != nil {
			return nil, err
		}

		encoder.EncodeBinary(adminBytes)

	case 1:

		customerBytes, err := customer.Encode()
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

func (u Account) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	u.fieldIndex, err = decoder.DecodeVarint()
	if err != nil {
		return err
	}

	switch u.fieldIndex {
	case -1:
		u.value = nil

	case 0:

		adminBytes, err := decoder.DecodeBinary()
		if err != nil {
			return err
		}

		err = u.Admin.Decode(adminBytes)
		if err != nil {
			return err
		}

	case 1:

		customerBytes, err := decoder.DecodeBinary()
		if err != nil {
			return err
		}

		err = u.Customer.Decode(customerBytes)
		if err != nil {
			return err
		}

	}

	return nil
}

func (u Account) MustDecode(reader io.Reader) {
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

func (u AdminAccount) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	u.Active, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	return nil
}

func (u AdminAccount) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

type CustomerAccount struct {
	Age runtime.Nullable[uint8]
}

func (u *CustomerAccount) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	if u.Age.IsNull() {
		encoder.EncodeNull()
	} else {
		age := *u.Age.Value

		encoder.EncodeUint8(u.Age)

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

func (u CustomerAccount) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	if decoder.IsNextNull() {
		u.Age.Clear()
	} else {

		u.Age, err = decoder.DecodeUint8()
		if err != nil {
			return err
		}

	}

	return nil
}

func (u CustomerAccount) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
