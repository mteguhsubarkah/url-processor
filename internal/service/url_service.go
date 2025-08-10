package service

import (
    "errors"
    "net/url"
    "strings"
    "url_processor/internal/domain"
)

type URLService struct{}

func NewURLService() *URLService {
    return &URLService{}
}

func (s *URLService) ProcessURL(rawURL string, operation domain.OperationType) (string, error) {
    switch operation {
    case domain.OperationCanonical:
        return s.canonical(rawURL)
    case domain.OperationRedirection:
        return s.redirection(rawURL)
    case domain.OperationAll:
        cleanedURL, err := s.canonical(rawURL)
        if err != nil {
            return "Can't parse the URL", err
        }
        return s.redirection(cleanedURL)
    default:
        return "", errors.New("Unsupported operation")
    }
}

func (s *URLService) canonical(rawURL string) (string, error) {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return "Can't parse the URL", err
    }

	// Clean up the url by removing the query parameter
    parsedURL.RawQuery = ""
	
	if parsedURL.Path != "/" {
		// Remove trailing slash
        parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")
    }

    return parsedURL.String(), nil
}

func (s *URLService) redirection(rawURL string) (string, error) {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return "", err
    }

	// redirect the host to byfood domain
	parsedURL.Host = "www.byfood.com"


    protocolScheme := strings.ToLower(parsedURL.Scheme) // take and lowercase the protocol scheme
    host := strings.ToLower(parsedURL.Host) // take and lowercase  the url host
    path := strings.ToLower(parsedURL.Path) // take and lowercase  the url path
   
    redirectedURL := protocolScheme + "://" + host + path

    if parsedURL.RawQuery != "" {
        redirectedURL += "?" + parsedURL.RawQuery
    }
    if parsedURL.Fragment != "" {
        redirectedURL += "#" + parsedURL.Fragment
    }

    return redirectedURL, nil
}
