package domain

type URLRequest struct {
    URL       string        `json:"url"`
    Operation OperationType `json:"operation"`
}

type URLResponse struct {
    ProcessedURL string `json:"processed_url"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}