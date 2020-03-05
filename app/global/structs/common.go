package structs

import "github.com/gorilla/websocket"

// EnvConfig dev.yaml格式
type EnvConfig struct {
	DBMaster    DbMaster    `yaml:"master"`
	DbSlave     DbSlave     `yaml:"slave"`
	API         API         `yaml:"api"`
	Log         Log         `yaml:"log"`
	DB          DB          `yaml:"db"`
	Redis       Redis       `yaml:"redis"`
	GrpcSetting GrpcSetting `yaml:"grpc_setting"`
}

// DbMaster 載入db的master環境設定
type DbMaster struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// DbSlave 載入db的slave環境設定
type DbSlave struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// API 載入各單位api環境設定
type API struct {
	PitayaGrpcServer string `yaml:"pitaya_grpc_server"`
	LemonGrpcServer  string `yaml:"lemon_grpc_server"`
}

// Log 載入Log設定檔規則
type Log struct {
	LogDir    string `yaml:"log_dir"`
	AccessLog string `yaml:"access_log"`
	ErrorLog  string `yaml:"error_log"`
}

// DB 對DB其他操作的設定
type DB struct {
	Debug bool `yaml:"debug"`
}

// Redis 載入redis設定
type Redis struct {
	RedisHost string `yaml:"redis_host"`
	RedisPort string `yaml:"redis_port"`
	RedisPwd  string `yaml:"redis_password"`
}

// APIResult 回傳API格式
type APIResult struct {
	ErrorCode   int         `json:"error_code"`
	ErrorMsg    string      `json:"error_msg"`
	LogIDentity string      `json:"log_id"`
	Result      interface{} `json:"result"`
}

// GrpcSetting grpc 自動註冊服務設定
type GrpcSetting struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// WsClient Websocket Client Info
type WsClient struct {
	Hub    *Hub           // ws 行為
	UID    string         // 用戶
	PID    string         // 代理
	Token  string         // 憑證
	GameID string         // 遊戲 ID
	Conn   websocket.Conn // ws 連線
	Send   chan []byte    // 訊息
}

// Hub maintains the set of active clients and broadcasts messages to the client
type Hub struct {
	Singal     chan SingalMsg       // Registered clients.
	Broadcast  chan []byte          // Inbound messages from the clients.
	Register   chan *WsClient       // Register requests from the clients.
	Unregister chan *WsClient       // Unregister requests from clients.
	List       map[string]*WsClient // Unregister requests from clients.
}

// SingalMsg 個人訊息
type SingalMsg struct {
	uuid string
	msg  []byte
}
