package helper

import "strconv"

func TransformId(requestId string) int {
	id, err := strconv.Atoi(requestId)
	PanicIfError(err)

	return id
}

func RequestCheck(request interface{}, model interface{}) interface{} {
	if request != "" {
		return request
	}
	return model
}
