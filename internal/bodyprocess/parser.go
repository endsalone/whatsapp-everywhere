package bodyprocess

import (
	"io"
)

type Parser interface {
	Parse() (io.Reader, error)
	ContentType() string
}
