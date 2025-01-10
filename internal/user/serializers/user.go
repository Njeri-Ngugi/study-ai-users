package serializers

type CreateUserRequest struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	DateOfBirth    string `json:"date_of_birth" validate:"required"`
	Email          string `json:"email" validate:"required,custom_email"`
	Username       string `json:"username" validate:"required,min=3,max=10"`
	Password       string `json:"password" validate:"required,min=8,max=40"`
	CourseName     string `json:"course_name" validate:"required"`
	Institution    string `json:"institution" validate:"required"`
	YearOfStudy    uint16 `json:"year_of_study" validate:"required,number"`
	CompletionDate string `json:"completion_date" validate:"required"`
}
