package utils

import (
	"encoding/json"
	"fmt"
	tps "to-do-app-in-go/types"
	hlp "to-do-app-in-go/helpers"
	cnsts "to-do-app-in-go/constants"
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

func ValidateRequest(request tps.TaskRequest) error {
		if hlp.FindInArray(request.Status, cnsts.ValidStatuses) == -1 {
		fmt.Printf("Invalid status. Status must be one of: %v", cnsts.ValidStatuses)
		return fmt.Errorf("Invalid status. Status must be one of: %v", cnsts.ValidStatuses)
	}
	return nil
}
