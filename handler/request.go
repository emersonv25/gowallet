package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

type CreateWalletRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (r *CreateWalletRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Document == "" {
		return errParamIsRequired("document", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	return nil
}
