package main

import (
	"context"

	pb "github.com/HouLPcode/swap"
)

type swapServiceImpl struct {
	pb.UnimplementedSwapService
}

// Swap 执行切换操作
func (s *swapServiceImpl) Swap(
	ctx context.Context,
	req *pb.SwapRequest,
) (*pb.SwapResponse, error) {
	rsp := &pb.SwapResponse{}
	return rsp, nil
}
