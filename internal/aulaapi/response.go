package aulaapi

import (
	"encoding/json"
	"fmt"
)

// WebResponseStatus is the status block returned by the API.
// Maps to WebResponseStatus from the decompiled assembly.
type WebResponseStatus struct {
	HTTPCode            int     `json:"httpCode"`
	BackendErrorCode    int     `json:"code"`
	Message             *string `json:"message,omitempty"`
	PresentedMessage    *string `json:"presentedMessage,omitempty"`
	SubCode             *int    `json:"subCode,omitempty"`
	HTMLContentIfError  *string `json:"htmlContentIfError,omitempty"`
}

// WebResponseStatusSubCode represents known sub-code constants.
type WebResponseStatusSubCode int

const (
	SubCodeAuthorizationDeniedAnyScope               WebResponseStatusSubCode = 1
	SubCodeAuthorizationDeniedInstitutionScope        WebResponseStatusSubCode = 2
	SubCodeAuthorizationDeniedGroupScope              WebResponseStatusSubCode = 3
	SubCodeAuthorizationDeniedProfileScope            WebResponseStatusSubCode = 4
	SubCodeAuthorizationDeniedBlockedCommunication    WebResponseStatusSubCode = 5
	SubCodeAuthorizationDeniedAccessNotGranted        WebResponseStatusSubCode = 6
	SubCodeAuthorizationDeniedUserDeactivated         WebResponseStatusSubCode = 7
	SubCodeAuthorizationStepUpRequired                WebResponseStatusSubCode = 8
	SubCodeInvalidToken                               WebResponseStatusSubCode = 9
	SubCodeOutOfSyncPresenceConfiguration             WebResponseStatusSubCode = 10
	SubCodeUnregisterDeviceFailed                     WebResponseStatusSubCode = 11
	SubCodeCrossMunicipalityTagging                   WebResponseStatusSubCode = 12
	SubCodeSessionExpired                             WebResponseStatusSubCode = 13
	SubCodeExceedingMaximumParticipants               WebResponseStatusSubCode = 14
	SubCodeDateAlreadyHasOccurrenceFromSameSeries     WebResponseStatusSubCode = 15
	SubCodeFirstRepeatingEventExceptionOutOfRange     WebResponseStatusSubCode = 16
	SubCodeLastRepeatingEventExceptionOutOfRange      WebResponseStatusSubCode = 17
	SubCodeDeactivatedInstitutionProfile              WebResponseStatusSubCode = 18
	SubCodeSecureDocsOnlyShareWithinOneMunicipality   WebResponseStatusSubCode = 19
)

// SubCodeFromCode converts a raw sub-code integer to a known variant.
// Returns 0 and false if the code is not recognized.
func SubCodeFromCode(code int) (WebResponseStatusSubCode, bool) {
	if code >= 1 && code <= 19 {
		return WebResponseStatusSubCode(code), true
	}
	return 0, false
}

// AulaServiceResponse is the generic API response envelope.
type AulaServiceResponse[T any] struct {
	Status WebResponseStatus `json:"status"`
	Data   T                 `json:"data"`
}

// DataArrayResponse is a pagination wrapper for array responses.
type DataArrayResponse[T any] struct {
	TotalHits int `json:"totalHits"`
	Results   []T `json:"results"`
}

// AulaErrorResponse is the error response wrapper.
type AulaErrorResponse[T any] struct {
	Status AulaErrorResponseStatus[T] `json:"status"`
}

// AulaErrorResponseStatus is the error status block.
type AulaErrorResponseStatus[T any] struct {
	ErrorInformation T `json:"errorInformation"`
}

// parseEnvelope parses the standard Aula response envelope and returns the data payload.
// It maps known sub-codes and error conditions to appropriate errors.
func parseEnvelope[T any](body []byte, httpStatusCode int) (T, error) {
	var zero T

	var envelope AulaServiceResponse[T]
	if err := json.Unmarshal(body, &envelope); err != nil {
		if httpStatusCode < 200 || httpStatusCode >= 300 {
			return zero, &APIError{
				Message: fmt.Sprintf("HTTP %d: %s", httpStatusCode, string(body)),
			}
		}
		return zero, fmt.Errorf("JSON error: %w", err)
	}

	// Check envelope-level errors via sub-code.
	if envelope.Status.SubCode != nil {
		sc := *envelope.Status.SubCode
		if code, ok := SubCodeFromCode(sc); ok {
			switch code {
			case SubCodeInvalidToken:
				return zero, ErrInvalidToken
			case SubCodeSessionExpired:
				return zero, ErrSessionExpired
			case SubCodeAuthorizationStepUpRequired:
				return zero, ErrStepUpRequired
			case SubCodeAuthorizationDeniedUserDeactivated:
				return zero, ErrUserDeactivated
			}
		}
	}

	// Non-zero backend error code with no mapped sub-code.
	if envelope.Status.BackendErrorCode != 0 {
		msg := fmt.Sprintf("backend error %d", envelope.Status.BackendErrorCode)
		if envelope.Status.Message != nil {
			msg = *envelope.Status.Message
		}
		return zero, &APIError{
			Message: msg,
			Status:  &envelope.Status,
		}
	}

	return envelope.Data, nil
}
