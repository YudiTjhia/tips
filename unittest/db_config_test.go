package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DbConfig_Negative(t *testing.T) {

	_, err := NewDbConfig("", "", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "host is empty")

	_, err = NewDbConfig("localhost", "", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user is empty")

	_, err = NewDbConfig("localhost", "user", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "password is empty")

	_, err = NewDbConfig("localhost", "user", "pass", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "port is empty")

	_, err = NewDbConfig("localhost", "user", "pass", "abc", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "port is not a number")

	_, err = NewDbConfig("localhost", "user", "pass", "9999", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "db name is empty")

}

func Test_DbConfig_Normal(t *testing.T) { // e.g for CRUD test, or other normal test
	_, err := NewDbConfig("localhost", "user", "pass", "9999", "account_db")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test_DbConfig_InsertOnly(t *testing.T) { // for example to insert into db

}
