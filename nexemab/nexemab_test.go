package nexemab

import (
	"bytes"
	"testing"
	"time"

	"github.com/nexema/go/buffer"
	"github.com/nexema/go/runtime"
	"github.com/stretchr/testify/require"
)

type user struct {
	ID          string
	Name        string
	BirthDate   runtime.Nullable[time.Time]
	Tags        []string
	Preferences runtime.Nullable[[]string]
	Active      bool
	Metadata    map[string]string
}

func (u *user) Encode(encoder *Encoder) {
	encoder.EncodeString(u.ID)
	encoder.EncodeString(u.Name)
	if u.BirthDate.IsNull() {
		encoder.EncodeNull()
	} else {
		encoder.EncodeInt64(u.BirthDate.Value.Unix())
	}

	encoder.BeginArray(int64(len(u.Tags)))
	for _, value := range u.Tags {
		encoder.EncodeString(value)
	}

	if u.Preferences.IsNull() {
		encoder.EncodeNull()
	} else {
		preferences := *u.Preferences.Value
		encoder.BeginArray(int64(len(preferences)))
		for _, value := range preferences {
			encoder.EncodeString(value)
		}
	}

	encoder.EncodeBool(u.Active)
	encoder.BeginMap(int64(len(u.Metadata)))
	for k, v := range u.Metadata {
		encoder.EncodeString(k)
		encoder.EncodeString(v)
	}
}

func (u *user) Decode(decoder *Decoder) error {
	var err error
	u.ID, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.Name, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	if !decoder.IsNextNull() {
		value, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.BirthDate.SetValue(time.Unix(value, 0))
	}

	arrLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}

	u.Tags = make([]string, arrLen)
	for i := int64(0); i < arrLen; i++ {
		value, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.Tags[i] = value
	}

	if !decoder.IsNextNull() {
		arrLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		u.Preferences = runtime.NewNullable(make([]string, arrLen))
		for i := int64(0); i < arrLen; i++ {
			value, err := decoder.DecodeString()
			if err != nil {
				return err
			}

			(*u.Preferences.Value)[i] = value
		}
	}

	u.Active, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	mapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Metadata = make(map[string]string, mapLen)
	for i := int64(0); i < mapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		value, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.Metadata[key] = value
	}

	return nil
}

func TestEncodeStruct(t *testing.T) {
	myUser := user{
		ID:          "12345",
		Name:        "Tomas Weigenast",
		BirthDate:   runtime.NewNull[time.Time](),
		Tags:        []string{"tomas", "weigenast", "12345"},
		Preferences: runtime.NewNullable([]string{"cats", "cars", "food"}),
		Active:      true,
		Metadata: map[string]string{
			"activationDate": "now",
		},
	}

	encoder := NewEncoder(buffer.NewBytesBuffer())
	myUser.Encode(encoder)
	buffer := encoder.Close()

	// restore it
	decoder := NewDecoder(bytes.NewBuffer(buffer))
	myUser = user{}
	err := myUser.Decode(decoder)
	require.Nil(t, err)
	require.Equal(t, myUser, user{
		ID:          "12345",
		Name:        "Tomas Weigenast",
		BirthDate:   runtime.NewNull[time.Time](),
		Tags:        []string{"tomas", "weigenast", "12345"},
		Preferences: runtime.NewNullable([]string{"cats", "cars", "food"}),
		Active:      true,
		Metadata:    map[string]string{"activationDate": "now"},
	})
}
