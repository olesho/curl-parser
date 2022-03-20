package parser

import (
	"errors"
	"fmt"
	"github.com/google/shlex"
	//"github.com/jessevdk/go-flags"
	"net/http"
	"strings"
)

func parseHeader(h string) (string, string, error) {
	parts := strings.SplitN(h, ":", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
	}
	return "", "", errors.New(fmt.Sprintf(`wrong header format: %v`, h))
}

func Parse(curl string) (*http.Request, error) {
	method := ""
	url := ""
	body := ""
	headers := make(map[string]string)

	parts, err := shlex.Split(curl)
	if err != nil {
		return nil, err
	}

	var trimmedParts []string
	for _, p := range parts {
		part := strings.TrimSpace(p)
		if part != "" {
			trimmedParts = append(trimmedParts, part)
		}
	}

	var currentPart, nextPart string
	for i := 1; i < len(trimmedParts); {
		currentPart = trimmedParts[i-1]
		nextPart = trimmedParts[i]
		if currentPart != "" {
			switch currentPart {
			case "-X":
				method = strings.ToUpper(nextPart)
				i++
			case "-H":
				k, v, err := parseHeader(nextPart)
				if err != nil {
					return nil, err
				}
				headers[strings.ToLower(k)] = v
				i++
			case "-d":
				body = nextPart
				i++
			case "--data-raw":
				body = nextPart
				i++
			case "--abstract-unix-socket":
				i++
			case "--alt-svc":
				i++
			case "--aws-sigv4":
				i++
			case "-a":
			case "--append":
			case "--anyauth":
			case "--basic":
			case "curl":
			case "-k":
			case "-v":
			case "-V":
			default:
				if !strings.HasPrefix(currentPart, "-") {
					url = currentPart
				}
			}
		}
		i++
	}

	if body != "" {
		if method == "" {
			method = "POST"
		}
	} else {
		if method == "" {
			method = "GET"
		}
	}

	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return r, nil
}
