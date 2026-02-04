package contacts

import "fmt"

// Email
const Email = "support@example.com" // global exportable variable

var support string // global non-exportable variable

// SetSupport sets the value of support variable
func SetSupport(s string) { //exportable function
	support = s
}

// GetContact returns supports name and email
func GetContact() string { // exportable function
	return fmt.Sprintf("%s <%s>", support, Email)
}
