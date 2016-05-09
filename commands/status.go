package commands

import (
	"errors"

	imap "github.com/emersion/imap/common"
	"github.com/emersion/imap/utf7"
)

// A STATUS command.
// See https://tools.ietf.org/html/rfc3501#section-6.3.10
type Status struct {
	Mailbox string
	Items []string
}

func (cmd *Status) Command() *imap.Command {
	mailbox, _ := utf7.Encoder.String(cmd.Mailbox)

	items := make([]interface{}, len(cmd.Items))
	for i, f := range cmd.Items {
		items[i] = f
	}

	return &imap.Command{
		Name: imap.Status,
		Arguments: []interface{}{mailbox, items},
	}
}

func (cmd *Status) Parse(fields []interface{}) error {
	if len(fields) < 2 {
		return errors.New("No enough arguments")
	}

	mailbox, ok := fields[0].(string)
	if !ok {
		return errors.New("Mailbox name is not a string")
	}

	var err error
	if cmd.Mailbox, err = utf7.Decoder.String(mailbox); err != nil {
		return err
	}

	items, ok := fields[1].([]interface{})
	if !ok {
		return errors.New("Items must be a list")
	}

	cmd.Items = make([]string, len(items))
	for i, v := range items {
		cmd.Items[i], _ = v.(string)
	}

	return nil
}
