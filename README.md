# HTTP/1.1 Server in Go

Learning Golang, decided to implement an http/1.1 server in Go

## Start

```bash
# Clone the repository
git clone https://github.com/danishjuneja/http-server-go.git
cd http-server-go

# Run the server
go run ./cmd/httpserver

# Test it out!
curl http://localhost:42069/
curl http://localhost:42069/video
curl http://localhost:42069/httpbin/stream/100
```

Visit `http://localhost:42069` in your browser to see the server in action!

## 🏗️ Architecture Overview

```mermaid
graph TB
    Client[HTTP Client] --> TCP[TCP Connection]
    TCP --> Server[HTTP Server]

    Server --> Parser[HTTP Request Parser]
    Parser --> Router[Request Router]

    Router --> HTML[HTML Responses]
    Router --> Proxy[HTTPBin Proxy]
    Router --> Video[Binary Video Handler]

    Proxy --> HTTPBin[httpbin.org]
    HTTPBin --> Chunked[Chunked Transfer Encoding]

    HTML --> Response[Response Writer]
    Video --> Response
    Chunked --> Response

    Response --> HTTP[HTTP Response Formatting]
    HTTP --> TCP

    subgraph "Core Components"
        Parser
        Router
        Response
    end
```
