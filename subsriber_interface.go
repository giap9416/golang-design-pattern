package observer

type subscriber interface {
	update(string)
	getID() string
}
