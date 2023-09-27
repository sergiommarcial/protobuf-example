# protobuf-example

**Step 1: Install Protocol Buffers Compiler (protoc)**

Before we start, you need to install the `protoc` compiler, which is used to compile `.proto` files into language-specific code.

**For Windows:**

1. Visit the Protocol Buffers GitHub releases page for Windows: https://github.com/protocolbuffers/protobuf/releases
2. Download the latest `protoc-{version}-win32.zip` file.
3. Extract the contents of the zip file to a directory.
4. Add the `bin` directory to your system's PATH.

**For macOS:**

You can use Homebrew to install `protobuf`:

1. Open Terminal.
2. Run the following command:

```shell
brew install protobuf
```

**Step 2: Define Your .proto File**

Create a `.proto` file that defines your service and message types. Here's an example:

```proto
syntax = "proto3";

package myservice;
option go_package = "/gen;gen";

service MyService {
    rpc GetData (DataRequest) returns (DataResponse);
}

message DataRequest {
    string message = 1;
}

message DataResponse {
    string reply = 1;
}
```

**Step 3: Generate Go Code**

To generate Go code from your `.proto` file, you need the `protoc-gen-go` and `protoc-gen-go-grpc` plugins. Install it using `go install`:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

and

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

**Step 4: Generate Code**

Now, you can generate Go code from your `.proto` file:

```shell
protoc --go_out=. --go-grpc_out=. path/to/your/protobuf.proto
```

Replace `path/to/your/protobuf.proto` with the actual path to your `.proto` file.

**Step 5: Implement the Service**

With the generated Go code, you can now implement your gRPC service. Here's an example:

```go
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sergiommarcial/gen" // Import your generated proto package

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	return &pb.DataResponse{Reply: "Hello, " + req.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterMyServiceServer(server, pb.UnimplementedMyServiceServer{})
	log.Println("Server started on :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

```

This code sets up a gRPC server that listens on port 50051 and implements the `GetData` RPC method.

**Step 6: Compile and Run**

Compile and run your Go gRPC server:

```shell
go run main.go
```

Now, you have a gRPC service running locally.

Please adjust the file paths and package names to match your project's structure. This example provides a basic overview of generating a Protocol Buffers service using Go, and it can be extended as needed for your specific use case.
