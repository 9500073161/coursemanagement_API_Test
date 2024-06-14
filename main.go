package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"coursemanagement/handlers"
	"coursemanagement/mysqldbmodels"
	"coursemanagement/statemanager"

	"google.golang.org/grpc"

	pb "coursemanagement/proto"
)

var (
	port = flag.Int("port", 50051, "The gRPC server port")
)

// server is used to implement helloworld.GreeterServer.
type courseServer struct {
	pb.UnimplementedCourseServiceServer
	stateManager *statemanager.StateManager
	dbClient     *mysqldbmodels.DBClient
}

// GetAllCourses implements the GetAllCourses method of the CourseService.
func (s *courseServer) GetAllCourses(ctx context.Context, req *pb.GetAllCoursesRequest) (*pb.GetAllCoursesResponse, error) {
	dbClient := mysqldbmodels.DBClient{Conn: s.stateManager.GetDBConnection()}
	courses, err := dbClient.GetCourseRaw()
	if err != nil {
		return nil, err
	}

	var pbCourses []*pb.Course
	for _, course := range courses {
		pbCourses = append(pbCourses, &pb.Course{
			Id:        int32(course.ID),
			TeacherId: int32(course.TeacherID),
			Name:      course.Name,
		})
	}

	return &pb.GetAllCoursesResponse{Courses: pbCourses}, nil
}

func startGRPCServer(courseSrv *courseServer) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, courseSrv)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHTTPServer(sm *statemanager.StateManager) {
	// Create handlers
	courseHandler := handlers.NewCourseHandler(sm)
	studentHandler := handlers.NewStudentHandler(sm)
	teacherHandler := handlers.NewTeacherHandler(sm)
	enrollmentHandler := handlers.NewEntrollmentHandler(sm)

	// Define HTTP routes
	http.HandleFunc("/course", courseHandler.HandlersCourse)
	http.HandleFunc("/student", studentHandler.HandlersStudent)
	http.HandleFunc("/teacher", teacherHandler.HandlersTeacher)
	http.HandleFunc("/entrollment", enrollmentHandler.HandlersEntrollment)

	fmt.Println("HTTP server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Initialize the state manager
	sm, err := statemanager.InitStateManager()
	if err != nil {
		log.Fatal("Failed to initialize state manager:", err)
	}

	// Initialize DBClient
	dbClient := mysqldbmodels.DBClient{Conn: sm.GetDBConnection()}

	// Create course server with state manager and DBClient
	courseSrv := &courseServer{stateManager: sm, dbClient: &dbClient}

	// Start gRPC server in a separate goroutine
	go startGRPCServer(courseSrv)

	// Start HTTP server in the main goroutine
	startHTTPServer(sm)
}
