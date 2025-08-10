package handler

import (
    "encoding/json"
    "net/http"
    "url_processor/internal/domain"
    "url_processor/internal/service"
)

type URLHandler struct {
    urlService *service.URLService
}

func NewURLHandler(urlService *service.URLService) *URLHandler {
    return &URLHandler{urlService: urlService}
}

func (h *URLHandler) ProcessURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req domain.URLRequest

    // Ensuring correct body request
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    if err := decoder.Decode(&req); err != nil {
        http.Error(w, "invalid body request: "+err.Error(), http.StatusBadRequest)
        return
    }


    // Process the URL
    processedURL, err := h.urlService.ProcessURL(req.URL, req.Operation)
    if err != nil {
        http.Error(w, "processing error: "+err.Error(), http.StatusBadRequest)
        return
    }

    resp := domain.URLResponse{ProcessedURL: processedURL}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
