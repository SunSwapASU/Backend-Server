package models

type RegisterCreds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileSetup struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PreferredContact string `json:"preferred_contact"`
	CampusName       string `json:"campus_name"`
	Major            string `json:"major"`
	GradYear         string `json:"grad_year"`
	Bio              string `json:"bio"`
}

type LoginCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
