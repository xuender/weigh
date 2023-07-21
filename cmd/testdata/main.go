package main

import (
	"os"
	"path/filepath"

	"github.com/bytedance/sonic"
	"github.com/samber/lo"
	"github.com/xuender/kit/oss"
	"github.com/xuender/weigh/pb"
)

func main() {
	msg := &pb.Msg{
		Request: []*pb.Request{
			{
				URL:    "http://127.0.0.1:8080/",
				Method: "GET",
			},
			{
				URL:    "http://127.0.0.1:8080/",
				Method: "GET",
			},
			{
				URL:    "http://127.0.0.1:8080/",
				Method: "GET",
			},
		},
	}

	data := lo.Must1(sonic.Marshal(msg))
	lo.Must0(os.WriteFile(filepath.Join("proxy", "msg.json"), data, oss.DefaultFileMode))
}
