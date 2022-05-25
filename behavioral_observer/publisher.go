package observer

import "fmt"

type post struct {
	subscriberList []subscriber
	tittle         string
	isNew          bool
}

func NewPost(tittle string) publisher {
	return &post{
		tittle: tittle,
	}
}
func (p *post) UpdateStatus() {
	fmt.Printf("New post published:  %s\n", p.tittle)
	p.isNew = true
	p.NotifyAll()
}
func (p *post) Register(o subscriber) {
	p.subscriberList = append(p.subscriberList, o)
}

func (p *post) Deregister(o subscriber) {
	p.subscriberList = removeFromslice(p.subscriberList, o)
}

func (p *post) NotifyAll() {
	for _, subscriber := range p.subscriberList {
		subscriber.update(p.tittle)
	}
}

func removeFromslice(subscriberList []subscriber, observerToRemove subscriber) []subscriber {
	subscriberListLength := len(subscriberList)
	for i, subscriber := range subscriberList {
		if observerToRemove.getID() == subscriber.getID() {
			subscriberList[subscriberListLength-1], subscriberList[i] = subscriberList[i], subscriberList[subscriberListLength-1]
			return subscriberList[:subscriberListLength-1]
		}
	}
	return subscriberList
}
