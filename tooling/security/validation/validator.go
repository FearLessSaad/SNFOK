package validation

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single validation error
type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// validatorInstance is a singleton validator to cache struct metadata
var (
	validatorInstance *validator.Validate
	validatorOnce     sync.Once
)

// getValidator lazily initializes and returns a singleton validator
func getValidator() *validator.Validate {
	validatorOnce.Do(func() {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())

		// Register custom "eq" validation
		validatorInstance.RegisterValidation("eq", func(fl validator.FieldLevel) bool {
			switch fl.Field().Kind() {
			case reflect.String:
				return fl.Field().String() == fl.Param()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return fmt.Sprintf("%d", fl.Field().Int()) == fl.Param()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return fmt.Sprintf("%d", fl.Field().Uint()) == fl.Param()
			case reflect.Float32, reflect.Float64:
				return fmt.Sprintf("%f", fl.Field().Float()) == fl.Param()
			default:
				return false
			}
		})
	})
	return validatorInstance
}

// ValidateStruct validates any struct and returns a slice of ValidationError
func ValidateStruct(data interface{}) []ValidationError {
	var errors []ValidationError

	// Get validator instance
	validate := getValidator()

	// Validate the struct
	err := validate.Struct(data)
	if err == nil {
		return errors
	}

	// Cache field tags to minimize reflection
	fieldTagCache := make(map[string]string)

	// Handle validation errors
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := getFieldJSONTag(data, err.Field(), fieldTagCache)
		errorMsg := buildErrorMessage(err)

		errors = append(errors, ValidationError{
			Field: fieldName,
			Error: errorMsg,
		})
	}

	return errors
}

// getFieldJSONTag retrieves the JSON tag for a struct field, using a cache
func getFieldJSONTag(data interface{}, fieldName string, cache map[string]string) string {
	// Check cache first
	if tag, ok := cache[fieldName]; ok {
		return tag
	}

	// Get struct value
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		cache[fieldName] = fieldName
		return fieldName
	}

	// Get field by name
	field, ok := value.Type().FieldByName(fieldName)
	if !ok {
		cache[fieldName] = fieldName
		return fieldName
	}

	// Extract JSON tag
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		cache[fieldName] = fieldName
		return fieldName
	}

	// Handle JSON tag with options (e.g., "name,omitempty")
	tag := strings.Split(jsonTag, ",")[0]
	cache[fieldName] = tag
	return tag
}

// buildErrorMessage constructs a user-friendly error message for a validation error
func buildErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Value must be at least %s", err.Param())
	case "max":
		return fmt.Sprintf("Value must not exceed %s", err.Param())
	case "gte":
		return fmt.Sprintf("Value must be greater than or equal to %s", err.Param())
	case "lte":
		return fmt.Sprintf("Value must be less than or equal to %s", err.Param())
	case "eq":
		return fmt.Sprintf("Value must be equal to %s", err.Param())
	case "uuid":
		return "Invalid UUID format"
	default:
		return fmt.Sprintf("Validation failed on %s", err.Tag())
	}
}
