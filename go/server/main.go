package main

import (
	"fmt"
	"net"
	"os"

	// protoc で自動生成されたパッケージ
	"github.com/ymmt2005/grpc-tutorial/go/deepthought"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

const portNumber = 13333

func main() {
	serv := grpc.NewServer()

	// 実装した Server を登録
	deepthought.RegisterComputeServer(serv, &Server{})

	// Healthcheck サーバーの作成
	hs := health.NewServer()
	healthgrpc.RegisterHealthServer(serv, hs)
	hs.Resume()

	// 待ち受けソケットを作成
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}

	// ヘルスチェックの結果を表示
	fmt.Println("Healthcheck server is running.")

	// gRPC サーバーでリクエストの受付を開始
	// l は Close されてから戻るので、main 関数での Close は不要
	serv.Serve(l)
}
