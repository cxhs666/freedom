package common

//统一的认证信息
type Auth struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatr"`
}
