package tcp_rpc

import "time"

type RpcManager struct {
	requests  map[int64]*Request
	requestId int64 // request唯一Id
}

func CreateRpcManager() *RpcManager {
	return &RpcManager{
		requests: make(map[int64]*Request),
	}
}

func (mgr *RpcManager) Action() {
	curTime := time.Now().Unix()

	for _, req := range mgr.requests {
		if req.IsExpired(curTime) {
			delete(mgr.requests, req.Id)
		}
	}
}

func (mgr *RpcManager) AddRequest(request *Request) int64 {
	mgr.requestId++
	request.SetId(mgr.requestId)
	mgr.requests[mgr.requestId] = request
	return mgr.requestId
}

func (mgr *RpcManager) GetRequestById(requestId int64) *Request {
	return mgr.requests[requestId]
}
