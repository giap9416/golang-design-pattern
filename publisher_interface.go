package observer

type publisher interface {
	Register(Subscriber subscriber)
	Deregister(Subscriber subscriber)
	UpdateStatus()
	NotifyAll()
}
