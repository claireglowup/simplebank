package mail

import (
	"techschool/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {

	config, err := util.LoadConfig("../.env")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test Email"
	content := `
		<h1> i love you mia </h1>
		<p>This is a test message from go</p>
	`
	to := []string{"rikymia01@gmail.com"}
	attachFiles := []string{"mia.txt"}
	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
