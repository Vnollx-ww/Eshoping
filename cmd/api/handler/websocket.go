package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/user"
	"Eshop/pkg/middlerware"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"go.uber.org/zap"
	"log"
	"sync"
)

// 定义 WebSocket 连接的结构体
type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	userid int64
}

// 定义 WebSocket 服务器
type WebSocketServer struct {
	clients    map[int64]*Client // 连接的客户端
	broadcast  chan []byte       // 广播消息的通道
	register   chan *Client      // 新客户端的通道
	unregister chan *Client      // 断开连接的客户端通道
	mu         sync.Mutex        // 锁，保护 clients 数据
}

var server *WebSocketServer

func init() {
	server = NewWebSocketServer()
	go HandleMessages()
}

// 创建新的 WebSocket 服务器
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:    make(map[int64]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}
func findClientByUserID(userID int64) *Client {
	server.mu.Lock()
	defer server.mu.Unlock()

	return server.clients[userID] // 直接根据 userID 查找目标客户端
}
func WebSocketConnections(ctx context.Context, c *app.RequestContext) {
	var upgrader = websocket.HertzUpgrader{}
	token := c.DefaultQuery("token", "")
	if token == "" {
		logger.Error("Token缺失")
		BadBaseResponse(c, "Token缺失")
		return
	}

	mc, err := middlerware.ParseToken(token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		BadBaseResponse(c, "token解析失败")
	}
	// 使用 HertzUpgrader将 HTTP 请求升级为 WebSocket 连接，并通过回调函数处理 WebSocket 连接
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		// 创建新的客户端结构体

		// 处理后续的连接
		client := &Client{
			conn:   conn,
			send:   make(chan []byte),
			userid: mc.UserId,
		}

		// 注册客户端
		server.register <- client

		// 启动 goroutine来处理读取和发送消息
		go client.ReadMessages()
		go client.SendMessages()
		select {}
		// 连接关闭时，取消注册并关闭连接
		/*defer func() {
			server.unregister <- client
			conn.Close()
		}()*/
	})

	if err != nil {
		log.Println("Error upgrading connection:", err)
	}
}

// 读取客户端的消息
func (client *Client) ReadMessages() {
	for {
		messageType, message, err := client.conn.ReadMessage()
		messageStr := string(message)
		//log.Println(messageStr)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("Connection closed normally")
			} else {
				log.Println("Error reading message:", err)
			}
			break
		}
		if messageType == websocket.TextMessage {
			// 假设前端传过来的消息是 JSON 格式
			// 示例：{"token": "user_token", "touserid": "123", "message": "Hello!"}
			var msgData struct {
				Type     string `json:"type"`
				Token    string `json:"token"`
				Touserid int64  `json:"touserid"`
				Message  string `json:"content"`
			}

			// 解析消息
			err = json.Unmarshal([]byte(messageStr), &msgData)
			if err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}

			// 调用 RPC 服务将消息存入数据库
			req := &user.SendMessageRequest{
				Token:    msgData.Token,
				TouserId: msgData.Touserid,
				Content:  msgData.Message,
			}
			//log.Println(req.Content)
			_, _ = rpc.SendMessage(context.Background(), req)
			targetClient := findClientByUserID(msgData.Touserid)
			if targetClient != nil {
				targetClient.send <- message
			}
		}
	}
}

// 向客户端发送消息
func (client *Client) SendMessages() {
	for message := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

func HandleMessages() {
	for {
		select {
		case client := <-server.register:
			server.mu.Lock()
			server.clients[client.userid] = client // 使用 userID 作为 key 注册
			server.mu.Unlock()
			log.Println("New client connected")
		case client := <-server.unregister:
			server.mu.Lock()
			delete(server.clients, client.userid) // 通过 userID 注销
			close(client.send)
			server.mu.Unlock()
			log.Println("Client disconnected")
		}
		// 不再处理 broadcast 消息，改为在 ReadMessages 中直接发送给目标客户端
	}
}
