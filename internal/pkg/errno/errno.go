package errno

import (
	"backend/internal/pkg/response"
	"net/http"
)

var (
	AccessDenied = response.NewCodeError(http.StatusForbidden, "access denied")
)
