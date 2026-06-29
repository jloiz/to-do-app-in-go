package utils

import (
	"encoding/json"
	"fmt"
	tps "to-do-app-in-go/types"
)

func ParseRequest(rawRequest []byte) (tps.TaskRequest, error) {
	var parsedRequest tps.TaskRequest
	taskReqStr := fmt.Sprintf("%s", rawRequest)
	fmt.Printf("Parsing request: \n %s\n", taskReqStr)
	err := json.Unmarshal([]byte(rawRequest), &parsedRequest)

	if err != nil {
		return tps.TaskRequest{}, fmt.Errorf("Invalid task request format: ", err)
	}

	return parsedRequest, nil
}
