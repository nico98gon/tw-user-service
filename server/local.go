package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"user-service/pkg/aws"
	"user-service/services"

	"github.com/aws/aws-lambda-go/events"
)

func StartLocalServer() {
	aws.StartAWS()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		headers := make(map[string]string)
		for key, values := range r.Header {
			if len(values) > 0 {
				headers[key] = values[0]
			}
		}

		request := events.APIGatewayProxyRequest{
			HTTPMethod: r.Method,
			Path:       r.URL.Path,
			Headers:    headers,
			Body:       string(body),
		}

		response, err := services.LambdaExec(context.Background(), request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(response.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response.Body))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}