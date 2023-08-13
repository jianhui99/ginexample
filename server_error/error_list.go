package server_error

var (
	// server error
	OrderAlreadyExists           = ServerError{code: 2001, msg: "Oops! Order already exists"}
	FailedToValidateStruct       = ServerError{code: 2002, msg: "Failed to validate struct"}
	CityAndCountryNotFoundFromIP = ServerError{code: 2003, msg: "City and country not found from IP"}

	// bad request error
	WebRec2Category1NotFound = BadRequestError{code: 4001, msg: "Web Recommend 2 Category 1 not found"}
	WebRec2Category2NotFound = BadRequestError{code: 4002, msg: "Web Recommend 2 Category 2 not found"}
)
