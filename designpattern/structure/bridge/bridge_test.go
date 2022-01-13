package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification(t *testing.T) {
	// email := &EmailMSGSender{}
	email := NewEmailMSGSender([]string{"1025@qq.com"})
	errorNotification := NewErrorNotification(email)
	err := errorNotification.Notify("hi")
	assert.Nil(t, err)
}
