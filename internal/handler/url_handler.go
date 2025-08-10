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


// ProcessURL godoc
// @Summary Process a given URL
// @Description Cleans or redirects a URL based on the operation type
// @Tags URL
// @Accept  json
// @Produce  json
// @Param request body domain.URLRequest true "URL and operation"
// @Success 200 {object} domain.URLResponse
// @Failure 400 {object} domain.ErrorResponse "Invalid input or invalid body request"
// @Failure 405 {object} domain.ErrorResponse "Method Not Allowed"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /process-url [post]
func (h *URLHandler) ProcessURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var req domain.URLRequest

    // Ensuring correct body request
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    if err := decoder.Decode(&req); err != nil {
        generateErrorResponse(w, http.StatusBadRequest, "Invalid body request: "+err.Error())
        return
    }


    // Process the URL
    processedURL, err := h.urlService.ProcessURL(req.URL, req.Operation)
    if err != nil {
        generateErrorResponse(w, http.StatusInternalServerError, "Error Processing URL: "+err.Error())
        return
    }

    resp := domain.URLResponse{ProcessedURL: processedURL}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}


func generateErrorResponse(w http.ResponseWriter, status int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(domain.ErrorResponse{Message: msg})
}