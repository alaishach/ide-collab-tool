// Package schemas
package schemas

// every field needs to be pointer it's necessary for input validation, see "ParseBodyJSON"
// not sure how setting a default should be done but I guess it's something like setting up a tag like this and then in ParseBodyJSON we init accordingly with the correct type default:"value"
// For vield validation we are using the "github.com/go-playground/validator/v10" library

type SignupData struct {
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type LoginData struct {
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
