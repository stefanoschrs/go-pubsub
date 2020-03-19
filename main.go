package pubsub

import "errors"

type Subscriber struct {
	Id      string
	channel chan []byte
}

var subscribers map[string]Subscriber

var (
	ErrExists = errors.New("exists")
	ErrNotFound   = errors.New("not found")
)

func (s Subscriber) Pub(data []byte) {
	s.channel <-data
}

func (s Subscriber) Sub() []byte {
	return <-s.channel
}

func (s Subscriber) GetChannel() chan []byte {
	return s.channel
}

func (s Subscriber) Close() {
	close(s.channel)
	delete(subscribers, s.Id)
}

func Create(id string) (s Subscriber, err error) {
	if subscribers == nil {
		subscribers = make(map[string]Subscriber)
	}

	if _, ok := subscribers[id]; ok {
		err = ErrExists
		return
	}

	s = Subscriber{
		Id:      id,
		channel: make(chan []byte),
	}

	subscribers[id] = s
	return
}

func Get(id string) (s Subscriber, err error) {
	if subscribers == nil {
		err = ErrNotFound
		return
	}

	if _, ok := subscribers[id]; !ok {
		err = ErrNotFound
		return
	}

	s = subscribers[id]
	return
}
