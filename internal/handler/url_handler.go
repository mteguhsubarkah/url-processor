package handler



import (
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
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
        generateErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
        return
    }

    var req domain.URLRequest
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()


    if err := decoder.Decode(&req); err != nil {
        generateErrorResponse(w, http.StatusBadRequest, "Invalid body request: "+err.Error())
        return
    }

    if strings.TrimSpace(req.URL) == "" {
        generateErrorResponse(w, http.StatusBadRequest, "URL is required")
        return
    }
    if strings.TrimSpace(string(req.Operation)) == "" {
        generateErrorResponse(w, http.StatusBadRequest, "Operation is required")
        return
    }

    // Validate operation type
    op, err := domain.ParseOperation(strings.ToLower(string(req.Operation)))
    if err != nil {
        generateErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    var inputURL string
    switch op {
    case domain.OperationAll, domain.OperationRedirection:
        // Normalized form
        inputURL, err = normalizeURL(req.URL)
        if err != nil {
            generateErrorResponse(w, http.StatusBadRequest, "Invalid URL: "+err.Error())
            return
        }
    case domain.OperationCanonical:
        // Keep original casing and path, just strip query & trailing slash
        inputURL, err = stripQueryAndSlash(req.URL)
        if err != nil {
            generateErrorResponse(w, http.StatusBadRequest, "Invalid URL: "+err.Error())
            return
        }
    }

    // Process the URL
    processedURL, err := h.urlService.ProcessURL(inputURL, op)
    if err != nil {
        generateErrorResponse(w, http.StatusInternalServerError, "Error Processing URL: "+err.Error())
        return
    }

    resp := domain.URLResponse{ProcessedURL: processedURL}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// normalizeURL: used for "all" and "redirection"
func normalizeURL(raw string) (string, error) {
    parsed, err := url.Parse(raw)
    if err != nil {
        return "", err
    }

    parsed.Host = strings.ToLower(parsed.Host)
    if !strings.HasPrefix(parsed.Host, "www.") {
        parsed.Host = "www." + parsed.Host
    }
    parsed.Path = strings.ToLower(parsed.Path)
    parsed.Path = strings.TrimSuffix(parsed.Path, "/")
    parsed.RawQuery = ""

    return parsed.String(), nil
}

// stripQueryAndSlash: used for "canonical"
func stripQueryAndSlash(raw string) (string, error) {
    parsed, err := url.Parse(raw)
    if err != nil {
        return "", err
    }
    parsed.Path = strings.TrimSuffix(parsed.Path, "/")
    parsed.RawQuery = ""
    return parsed.String(), nil
}




func generateErrorResponse(w http.ResponseWriter, status int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(domain.ErrorResponse{Message: msg})
}