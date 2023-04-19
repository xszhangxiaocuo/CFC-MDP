package user

type LoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type LoginRsp struct {
	code string `json:"code"`
}
