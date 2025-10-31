package zpl

import (
	"errors"
	"fmt"
)

var ErrInvalidDensity = fmt.Errorf("invalid density: must be one of %v", allowedDensities())
var ErrInvalidOutputFormat = fmt.Errorf("invalid output format: must be one of %v", allowedOutputFormats())
var ErrNilInput = errors.New("nil input")
var ErrNilOutput = errors.New("nil output")
var ErrInvalidWidth = fmt.Errorf("invalid width: must be between 1 and %d inches", MaxLabelSizeInches)
var ErrInvalidHeight = fmt.Errorf("invalid height: must be between 1 and %d inches", MaxLabelSizeInches)
