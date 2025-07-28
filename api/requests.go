package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

func unmarshal_json[T any](req *http.Request) (T, error) {
	var result T
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error getting json from request: %s", err)
		return result, fmt.Errorf("error gettin json from request: %s", err)
	}
	defer req.Body.Close()

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("error unmarshaling json to type %s: %s", reflect.TypeOf(result), err)
		return result, fmt.Errorf("error gettin json from request: %s", err)
	}
	return result, nil
}
