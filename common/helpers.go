package common

type Response struct {
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Status     int         `json:"status"`
	TotalCount int         `json:"totalcount"`
}

func ResponseHandler(errcode, language string, totalcount int, data interface{}) Response {
	ed := ErrorMessages[errcode]
	msg := ed.(map[string]interface{})[language].(string)
	status := ed.(map[string]interface{})["status"].(float64)
	r := Response{Msg: msg, Data: data, Status: int(status), TotalCount: totalcount}
	return r
}
