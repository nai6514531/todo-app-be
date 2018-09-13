package enity

type Result struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}
