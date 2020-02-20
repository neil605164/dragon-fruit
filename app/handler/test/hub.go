package test

type singalMsg struct {
	uuid string
	msg  []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	singal chan singalMsg

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Unregister requests from clients.
	list map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		list:       make(map[string]*Client),
		singal:     make(chan singalMsg),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.list[client.id] = client
		case client := <-h.unregister:
			if _, ok := h.list[client.id]; ok {
				delete(h.list, client.id)
				close(client.send)
			}
		case message := <-h.broadcast:
			for _, client := range h.list {
				select {
				case client.send <- message:
				default:
					delete(h.list, client.id)
					close(client.send)
				}
			}
		case singalMsg := <-h.singal:
			client := h.list[singalMsg.uuid]
			select {
			case client.send <- singalMsg.msg:
			default:
				delete(h.list, client.id)
				close(client.send)
			}
		}
	}
}
