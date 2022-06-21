package app

type hub_t struct {
	clients map[*client_t]bool
	regCh   chan *client_t
	unregCh chan *client_t
	bcastCh chan []byte
}

func newHub() *hub_t {
	return &hub_t{
		regCh:   make(chan *client_t),
		unregCh: make(chan *client_t),
		clients: make(map[*client_t]bool),
		bcastCh: make(chan []byte),
	}
}

func (h *hub_t) run() {
	for {
		select {
		case client := <-h.regCh:
			h.clients[client] = true
		case client := <-h.unregCh:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.sendCh)
			}
		case message := <-h.bcastCh:
			for client := range h.clients {
				select {
				case client.sendCh <- message:
				default:
					close(client.sendCh)
					delete(h.clients, client)
				}
			}
		}
	}
}
