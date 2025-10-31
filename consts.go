package zpl

const (
	PDF = "pdf"
	PNG = "png"
	// MaxLabelSizeInches is the maximum label width/height accepted by the
	// Labelary API (15 inches).
	MaxLabelSizeInches = 15
)

func allowedDensities() []int {
	return []int{6, 8, 12, 24}
}

func allowedOutputFormats() []string {
	return []string{PDF, PNG}
}
