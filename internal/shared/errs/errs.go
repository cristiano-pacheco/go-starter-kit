package errs

import (
	"net/http"

	"github.com/cristiano-pacheco/go-starter-kit/pkg/errs"
)

var (
	ErrRecordNotFound = errs.New("RECORD_NOT_FOUND", "Record not found", http.StatusNotFound, nil)
)
