package ws

type singalMsg struct {
	uuid string
	msg  []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Singal chan singalMsg

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// Unregister requests from clients.
	List map[string]*Client
}

// NewHub 回傳 struct
func NewHub() *Hub {
	return &Hub{
		List:       make(map[string]*Client),
		Singal:     make(chan singalMsg),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run 監聽 ws 事件
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.List[client.ID] = client
		case client := <-h.Unregister:
			if _, ok := h.List[client.ID]; ok {
				delete(h.List, client.ID)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for _, client := range h.List {
				select {
				case client.Send <- message:
				default:
					delete(h.List, client.ID)
					close(client.Send)
				}
			}
		case singalMsg := <-h.Singal:
			client := h.List[singalMsg.uuid]
			select {
			case client.Send <- singalMsg.msg:
			default:
				delete(h.List, client.ID)
				close(client.Send)
			}
		}
	}
}
