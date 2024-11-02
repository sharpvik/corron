package corron_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/sharpvik/corron"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (user *User) Validate() error {
	if len(user.Name) > 256 {
		return errors.New("name_too_long")
	}

	if user.Age < 18 {
		return errors.New("age_too_young")
	}

	return nil
}

func TestUnmarshal(t *testing.T) {
	var to User
	validUser := []byte(`{ "name": "Viktor", "age": 24 }`)
	invalidUser := []byte(`{ "name": "Viktor", "age": 15 }`)
	invalidType := []byte(`42`)
	assert.NoError(t, corron.Unmarshal(validUser, &to))
	err := corron.Unmarshal(invalidUser, &to)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "age_too_young")
	assert.Error(t, corron.Unmarshal(invalidType, &to))
}

func TestDecode(t *testing.T) {
	var to User
	validUser := strings.NewReader(`{ "name": "Viktor", "age": 24 }`)
	invalidUser := strings.NewReader(`{ "name": "Viktor", "age": 15 }`)
	invalidType := strings.NewReader(`42`)
	assert.NoError(t, corron.NewDecoder(validUser).Decode(&to))
	err := corron.NewDecoder(invalidUser).Decode(&to)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "age_too_young")
	assert.Error(t, corron.NewDecoder(invalidType).Decode(&to))
}
