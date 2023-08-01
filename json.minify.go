package lets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// return a minified JSON input string
// return an error encountered during minifiying or reading minified bytes
func JsonMinify(jsonB []byte) ([]byte, error) {
	var buff *bytes.Buffer = new(bytes.Buffer)
	errCompact := json.Compact(buff, jsonB)
	if errCompact != nil {
		newErr := fmt.Errorf("failure encountered compacting json := %v", errCompact)
		return []byte{}, newErr
	}

	b, err := io.ReadAll(buff)
	if err != nil {
		readErr := fmt.Errorf("read buffer error encountered := %v", err)
		return []byte{}, readErr
	}

	return b, nil
}
