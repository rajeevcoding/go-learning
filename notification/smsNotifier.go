package notification

import "fmt"

type SmsNotifier struct {
}

func (em *SmsNotifier) Notify(msg string) error {
	fmt.Printf("Notification from SMS: %s", msg)
	return nil
}
