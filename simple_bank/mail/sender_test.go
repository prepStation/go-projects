package mail

import (
	"testing"

	"github.com/prepStation/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "a test subject"
	content := `
	<h1>Hello World</h1>
	<p>hey click here to go to google <a href="https://www.google.com"></a></p>
	`
	to := []string{"sassim882@gmail.com"}
	attachFiles := []string{"../README.md"}
	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
