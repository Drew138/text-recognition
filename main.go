package main

import (
	"fmt"
	"log"
	"net"

	"github.com/drew138/text-recognition/equations"
	"github.com/drew138/text-recognition/solver"
	"google.golang.org/grpc"
)

func main() {

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	s := equations
	gs := grpc.NewServer()

	equations.RegisterEquationServiceServer(gs, &s)

	if err := gs.Serve(l); err != nil {
		log.Fatal("Failed to serve gRPC server")
	}

	Ab := [][]float64{
		{6, -2, 2, 4, 12},
		{12, -8, 6, 10, 34},
		{3, -13, 9, 3, 27},
		{-6, 4, 1, -18, -38},
	}

	sol, err := solver.SolveSystem(Ab)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sol)
}
