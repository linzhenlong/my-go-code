package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}

var (

	// 请求体解析错误.
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSc:400,
		Error:Err{
			Error:"requestBody parse failed",
			ErrorCode:"001",
		},
	}

	// 用户验证失败
	ErrorNotAuthUser = ErrorResponse{
		HttpSc:401,
		Error:Err{
			Error:"user authentication failed",
			ErrorCode:"002",
		},
	}

	ErrorDBError = ErrorResponse {
		HttpSc: 500,
		Error: Err{
			Error: "DB ops failed",
			ErrorCode: "003",
		},
	}

	ErrorInternalFaults = ErrorResponse {
		HttpSc: 500,
		Error: Err{
			Error: "internal faults",
			ErrorCode: "004",
		},
	}
	
)
