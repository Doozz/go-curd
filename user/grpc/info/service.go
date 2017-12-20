package info

import (
	context "golang.org/x/net/context"
)

type UinfoServer struct{}

func (s *UinfoServer) GetInfo(ctx context.Context, req *InfoReq) (*InfoRep, error) {
	resp := &InfoRep{}
	if req.ID == 1 {
		resp.RetCode = "1"
		resp.RetMsg = "this 1"
	} else {
		resp.RetCode = "2"
		resp.RetMsg = "this 2"
	}
	return resp, nil
}
