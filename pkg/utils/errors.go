package utils

import (
	"net/http"

	"github.com/labstack/echo"
)

// UnauthorizedError returns an echo instance of a http 401 Unauthorized
func UnauthorizedError(msg string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, msg)
}

// ForbiddenError returns an echo instance of a http 403 Forbidden
func ForbiddenError(msg string) error {
	return echo.NewHTTPError(http.StatusForbidden, msg)
}

// NotFoundError returns an echo instance of a http 404 NotFound
func NotFoundError(msg string) error {
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

// ConflictError returns an echo instance of a http 409 Conflict
func ConflictError(msg string) error {
	return echo.NewHTTPError(http.StatusConflict, msg)
}

// ValidationError returns an echo instance of a http 422 UnprocessableEntitty
func ValidationError(msg string) error {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
}

// InternalServerError returns an echo instance of a http 500 InternalServerError
func InternalServerError(msg string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, msg)
}
