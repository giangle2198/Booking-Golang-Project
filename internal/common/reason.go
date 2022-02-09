package common

type ReasonCode string

const (
	ReasonGeneralError    ReasonCode = "-1"
	ReasonDBError         ReasonCode = "-2"
	ReasonCacheError      ReasonCode = "-3"
	ReasonAdapterError    ReasonCode = "-4"
	ReasonNotFound        ReasonCode = "-5"
	ReasonInvalidArgument ReasonCode = "-6"
)

var reasonCodeValues = map[string]string{
	"-1": "General Error",
	"-2": "DB Error",
	"-3": "Cache Error",
	"-4": "Adapter Error",
	"-5": "Not Found",
	"-6": "Invalid Argument",
}

// Code represents reason code
func (rc ReasonCode) Code() string {
	return string(rc)
}

// Message represent reason message
func (rc ReasonCode) Message() string {
	if value, ok := reasonCodeValues[rc.Code()]; ok {
		return value
	}
	return ""
}

// ParseError parsing error code to message
func ParseError(err error) ReasonCode {
	if err == nil {
		return ReasonCode("")
	}
	return ReasonCode(err.Error())
}
