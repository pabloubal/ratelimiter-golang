package mapper

import (
	"bytes"
	"io"
	"io/github/pabloubal/ratelimiter/dto"
	"io/github/pabloubal/ratelimiter/internal/domain"
	pb "io/github/pabloubal/ratelimiter/proto"
	"net/http"
	"net/url"
)

func RQToEntity(rq *http.Request) domain.Request {
	return domain.NewRequest(rq.Method, *rq.URL, rq.Body, rq.Header)
}

func GrpcRQToEntity(rq *pb.RequestEndpointRQ) (domain.Request, domain.RestApiError) {
	method := rq.Method.String()
	urlObj, err := url.Parse(rq.Url)

	if err != nil {
		return nil, domain.ErrorBadURLFormat
	}

	return domain.NewRequest(method, *urlObj, io.NopCloser(bytes.NewReader(rq.Body)), toHeadersMap(rq.Headers)), nil
}

func HttpResponseToGrpc(rs *http.Response) *pb.RequestEndpointRS {
	body, err := io.ReadAll(rs.Body)
	defer rs.Body.Close()

	if err != nil {
		return &pb.RequestEndpointRS{
			StatusCode: http.StatusInternalServerError,
			Headers:    nil,
			Body:       []byte(err.Error()),
		}
	}

	return &pb.RequestEndpointRS{
		StatusCode: int32(rs.StatusCode),
		Headers:    toHeaderEntries(rs.Header),
		Body:       body,
	}
}

func CreateToEntity(c dto.CreateEndpointDto) (domain.Endpoint, domain.RestApiError) {
	return domain.NewEndpoint(c.Path, c.Url, c.Limit, c.Method)
}

func CreateToDto(domain.Endpoint) dto.CreateEndpointDto {
	return dto.CreateEndpointDto{}
}

func toHeadersMap(h []*pb.HeaderEntry) map[string][]string {
	headers := make(map[string][]string)
	for _, header := range h {
		headers[header.Key] = header.Val
	}
	return headers
}

func toHeaderEntries(headers map[string][]string) []*pb.HeaderEntry {
	entries := make([]*pb.HeaderEntry, len(headers))
	i := 0
	for k, vs := range headers {
		entries[i] = &pb.HeaderEntry{
			Key: k,
			Val: vs,
		}
		i++
	}
	return entries
}
