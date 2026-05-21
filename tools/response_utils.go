package tools

import (
	"fmt"
	"io"
)

const defaultResponseLimitBytes int64 = 10 * 1024 * 1024 // 10MB

// readResponseBody reads up to limit bytes from r. If the response exceeds
// limit, it returns the truncated data along with an error.
func readResponseBody(r io.Reader, limit int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > limit {
		return data[:limit], fmt.Errorf("response body exceeds maximum size of %d bytes; try narrowing your query", limit)
	}
	return data, nil
}
