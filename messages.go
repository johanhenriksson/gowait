package gowait

type MsgType string

const (
	InitMsg   MsgType = "task/init"
	ReturnMsg         = "task/return"
	ErrorMsg          = "task/error"
	LogMsg            = "task/log"
)

type BaseMessage struct {
	Type MsgType `json:"type"`
}

type InitMessage struct {
	Type    MsgType `json:"type"`
	ID      string  `json:"id"`
	Version string  `json:"version"`
	Task    Taskdef `json:"task"`
}

type ReturnMessage struct {
	Type       MsgType `json:"type"`
	ID         string  `json:"id"`
	Result     Result  `json:"result"`
	ResultType string  `json:"result_type"`
}

type ErrorMessage struct {
	Type  MsgType `json:"type"`
	ID    string  `json:"id"`
	Error string  `json:"error"`
}

type LogMessage struct {
	Type MsgType `json:"type"`
	ID   string  `json:"id"`
	File string  `json:"file"`
	Data string  `json:"data"`
}
