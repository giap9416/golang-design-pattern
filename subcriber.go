package observer

import "fmt"

type Subscriber struct {
	ID string
}

func (c *Subscriber) update(title string) {
	fmt.Printf("Notify to Subscriber %s for post: %s\n", c.ID, title)
}

func (c *Subscriber) getID() string {
	return c.ID
}
