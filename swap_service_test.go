package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/HouLPcode/swap"
	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/HouLPcode/swap/swap_mock.go -package=swap -self_package=github.com/HouLPcode/swap --source=stub/github.com/HouLPcode/swap/swap.trpc.go

func Test_swapServiceImpl_Swap(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	swapServiceService := pb.NewMockSwapServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := swapServiceService.EXPECT().Swap(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.SwapRequest) (*pb.SwapResponse, error) {
		s := &swapServiceImpl{}
		return s.Swap(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.SwapRequest
		rsp *pb.SwapResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.SwapResponse
			var err error
			if rsp, err = swapServiceService.Swap(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("swapServiceImpl.Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("swapServiceImpl.Swap() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
