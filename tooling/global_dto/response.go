package global_dto

// Meta contains additional metadata like pagination or timestamps.
type Meta struct {
	TotalCount  int64  `json:"total_count,omitempty"`  // For paginated responses
	CurrentPage int    `json:"current_page,omitempty"` // For paginated responses
	NextPage    *int   `json:"next_page,omitempty"`    // Nullable next page
	RequestID   string `json:"request_id,omitempty"`   // For tracing
	Code        int    `json:"code,omitempty"`
}

// Response is a generic API response structure.
type Response[T any] struct {
	Status  string `json:"status"`            // Response status ("success", "error", "partial")
	Message string `json:"message,omitempty"` // Human-readable message
	Data    *T     `json:"data,omitempty"`    // Response data (type-safe)
	Errors  []any  `json:"errors,omitempty"`  // List of error objects (can be strings, structs, etc.)
	Meta    *Meta  `json:"meta,omitempty"`    // Optional metadata
}
