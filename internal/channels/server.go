package channels

func NewServerChannel() ServerChannel {
	return make(chan error)
}

type ServerChannel chan error

func (ch ServerChannel) Stop() {
	ch <- nil
	close(ch)
}
