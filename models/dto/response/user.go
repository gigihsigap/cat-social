package response

type SignUpResponse struct {
    Name       string `json:"name"`
    Email      string `json:"email"`
    AccessToken string `json:"accessToken"`
}

type SignInResponse struct {
    Name       string `json:"name"`
    Email      string `json:"email"`
    AccessToken string `json:"accessToken"`
}