package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "coursemanagement/proto"
)

var (
	serverAddr string
)

func init() {
	serverAddr = os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:50051"
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the service.
	client := pb.NewCourseServiceClient(conn)

	for {
		// Call the GetAllCourses RPC method.
		response, err := client.GetAllCourses(context.Background(), &pb.GetAllCoursesRequest{})
		if err != nil {
			log.Fatalf("could not get courses: %v", err)
		}

		// Print the response.
		fmt.Println("Courses:")
		for _, course := range response.Courses {
			fmt.Printf("ID: %d, Name: %s, Teacher ID: %d\n", course.Id, course.Name, course.TeacherId)
		}

		// Wait for a minute before making the next request.
		time.Sleep(1 * time.Minute)
	}
}
