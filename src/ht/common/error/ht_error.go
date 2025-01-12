package error

import (
	"context"
	"encoding/json"
	"ht-crm/src/ht/common/result"
)

type HtError struct {
	Code    int
	Msg     string
	context context.Context
}

func (e *HtError) Error() string {
	traceId := e.context.Value("traceId")
	if traceId == nil {
		traceId = ""
	}
	res, _ := json.Marshal(result.NewResult(e.Code, e.Msg, traceId.(string), nil))
	return string(res)
}

func NewHtError(ctx context.Context, code int, msg string) *HtError {
	return &HtError{Code: code, Msg: msg, context: ctx}
}

func NewHtErrorFromMsg(ctx context.Context, msg string) *HtError {
	return &HtError{Code: result.FAILURE.Code, Msg: msg, context: ctx}
}

func NewHtErrorFromTemplate(ctx context.Context, template result.Template) *HtError {
	return &HtError{Code: template.Code, Msg: template.Msg, context: ctx}
}