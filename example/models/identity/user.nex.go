package identity

import (
	"github.com/example/models"
	"github.com/nexema/go/runtime"
)

type User struct {
	Id string

	Name string

	SelectedColor models.Colors

	Account CustomerAccount
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
	Age uint8
}

func (u *CustomerAccount) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error

	encoder.EncodeUint8(u.Age)

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
