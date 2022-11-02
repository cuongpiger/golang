package main

import "fmt"

type (
	// server is subject
	server interface {
		handleRequest(string, string) (int, string)
	}

	// Application is real subject
	Application struct{}

	// Nginx is proxy
	Nginx struct {
		application       *Application
		maxAllowedRequest int
		rateLimiter       map[string]int
	}
)

// Nginx's collection of methods
func newNginxServer() *Nginx {
	return &Nginx{
		application:       new(Application),
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (s *Nginx) checkRateLimiting(url string) bool {
	if s.rateLimiter[url] == 0 {
		s.rateLimiter[url] = 1
	}

	if s.rateLimiter[url] > s.maxAllowedRequest {
		return false
	}

	s.rateLimiter[url] = s.rateLimiter[url] + 1
	return true
}

func (s *Nginx) handleRequest(url, method string) (int, string) {
	allowed := s.checkRateLimiting(url)
	if !allowed {
		return 403, "Not allowed"
	}

	return s.application.handleRequest(url, method)
}

// Application's collection of methods
func (s *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User created"
	}

	return 404, "Not OK"
}

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
