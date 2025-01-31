package serializers

type CreateUserRequest struct {
	Firstname             string `json:"firstname" validate:"required,alpha"`
	Lastname              string `json:"lastname" validate:"required,alpha"`
	Username              string `json:"username" validate:"required,alphanum"`
	Password              string `json:"password" validate:"required,min=8"`
	Email                 string `json:"email" validate:"required,custom_email"`
	InstitutionName       string `json:"institution_name" validate:"required"`
	CourseName            string `json:"course_name" validate:"required"`
	YearOfStudy           int    `json:"year_of_study" validate:"required,min=1"`
	CompletionDate        string `json:"completion_date" validate:"required"`
	Gender                string `json:"gender" validate:"required,oneof=male female"`
	DateOfBirth           string `json:"date_of_birth"`
	CourseDurationInYears int    `json:"course_duration" validate:"omitempty,min=0"`
}
