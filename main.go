package main

import (
	pb "github.com/HouLPcode/swap"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterSwapServiceService(s.Service("swap.SwapService"), &swapServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
