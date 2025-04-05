package utils

type Response struct {
	State  map[string]any `json:"state"`
	Result map[string]any `json:"result"`
	Next   *NextTask      `json:"next"`
	Done   bool           `json:"done"`
	Error  any            `json:"error"`
}

type NextTask struct {
	Key     string         `json:"key"`
	Payload map[string]any `json:"payload"`
}

func ResponseSuccess(state, result map[string]any, nextTask *NextTask) Response {
	return Response{
		State:  state,
		Result: result,
		Next:   nextTask,
		Done:   nextTask == nil,
		Error:  nil,
	}
}

func ResponseFailed(state map[string]any, err any) Response {
	return Response{
		State:  state,
		Result: nil,
		Next:   nil,
		Done:   true,
		Error:  err,
	}
}
