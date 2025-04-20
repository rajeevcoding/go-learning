package notification

import "fmt"

type EmailNotifier struct {
}

func (em *EmailNotifier) Notify(msg string) error {
	fmt.Printf("Notification from Email: %s", msg)
	return nil
}
