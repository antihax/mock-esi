package swaggerServer

import "time"
var _ time.Time

/* 
Internal server error */
type GetInsurancePricesInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
