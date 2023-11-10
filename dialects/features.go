package dialects

type Feature bool

const (
	Supported   Feature = true
	Unsupported Feature = false
)

type Features struct {
	// TruncateTable indicates if the dialect supports TRUNCATE TABLE statement.
	// If not, DELETE FROM will be used instead.
	TruncateTable Feature
}
