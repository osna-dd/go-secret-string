package secret

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

var _ json.Marshaler = &String{}

func TestStringPtr(t *testing.T) {
	secret := "my-secret"
	s_ := New(secret)
	s := &s_

	t.Run("fmt.Printf", func(t *testing.T) {
		formats := []string{
			"%s",
			"%q",
			"%v",
			"%x",
			"%X",
			"%#v",
			"%+v",
		}
		for _, format := range formats {
			t.Run("'"+format+"'", func(t *testing.T) {
				assert := assert.New(t)
				assert.NotEqual(secret, fmt.Sprintf(format, s))
				t.Logf("'%s': %s", format, fmt.Sprintf(format, s))
			})
		}
	})

	t.Run("json.Marshal", func(t *testing.T) {
		assert := assert.New(t)
		bs, err := json.Marshal(&s)
		assert.Nil(bs)
		assert.Error(err)
	})

	t.Run("get secret value", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(secret, s.Get())
	})
}

func TestString(t *testing.T) {
	secret := "my-secret"
	s := New(secret)

	t.Run("fmt.Printf", func(t *testing.T) {
		formats := []string{
			"%s",
			"%q",
			"%v",
			"%x",
			"%X",
			"%#v",
			"%+v",
		}
		for _, format := range formats {
			t.Run("'"+format+"'", func(t *testing.T) {
				assert := assert.New(t)
				assert.NotEqual(secret, fmt.Sprintf(format, s))
				t.Logf("'%s': %s", format, fmt.Sprintf(format, s))
			})
		}
	})

	t.Run("json.Marshal", func(t *testing.T) {
		assert := assert.New(t)
		bs, err := json.Marshal(&s)
		assert.Nil(bs)
		assert.Error(err)
	})

	t.Run("get secret value", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(secret, s.Get())
	})
}

func TestString2Ptr(t *testing.T) {
	secret := "my-secret"
	s_ := String2(secret)
	s := &s_

	t.Run("fmt.Printf", func(t *testing.T) {
		formats := []string{
			"%s",
			"%q",
			"%v",
			"%x",
			"%X",
			"%#v",
			"%+v",
		}
		for _, format := range formats {
			t.Run("'"+format+"'", func(t *testing.T) {
				assert := assert.New(t)
				assert.NotEqual(secret, fmt.Sprintf(format, s))
				t.Logf("'%s': %s", format, fmt.Sprintf(format, s))
			})
		}
	})

	t.Run("json.Marshal", func(t *testing.T) {
		assert := assert.New(t)
		bs, err := json.Marshal(&s)
		assert.Nil(bs)
		assert.Error(err)
	})

	t.Run("get secret value", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(secret, string(*s))
	})
}
func TestString2(t *testing.T) {
	secret := "my-secret"
	s := String2(secret)

	t.Run("fmt.Printf", func(t *testing.T) {
		formats := []string{
			"%s",
			"%q",
			"%v",
			"%x",
			"%X",
			"%#v",
			"%+v",
		}
		for _, format := range formats {
			t.Run("'"+format+"'", func(t *testing.T) {
				assert := assert.New(t)
				assert.NotEqual(secret, fmt.Sprintf(format, s))
				t.Logf("'%s': %s", format, fmt.Sprintf(format, s))
			})
		}
	})

	t.Run("json.Marshal", func(t *testing.T) {
		assert := assert.New(t)
		bs, err := json.Marshal(&s)
		assert.Nil(bs)
		assert.Error(err)
	})

	t.Run("get secret value", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(secret, string(s))
	})
}
