package notification

type Notifier interface {
	Notify(msg string) error
}
