package tcp_rpc

type Request struct {
	CallBackData any
	Id           int64
	ExpiredTime  int64 // 过期时间
}

func (req *Request) SetId(id int64) {
	req.Id = id
}

func (req *Request) IsExpired(curTime int64) bool {
	return req.ExpiredTime >= curTime
}

func (req *Request) SetExpiredTime(expiredTime int64) {
	req.ExpiredTime = expiredTime
}

func (req *Request) GetExpiredTime() int64 {
	return req.ExpiredTime
}
