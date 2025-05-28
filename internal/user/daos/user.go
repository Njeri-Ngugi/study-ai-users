package daos

import (
	"context"
	"errors"
	"fmt"
	"github.com/Njeri-Ngugi/toolbox/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
	"users/internal/user/models"
)

func GenerateCourseCode(institutionName, courseName string) string {
	institutionAbbr := getInstitutionAbbr(institutionName)
	courseAbbr := getCourseAbbr(courseName)
	return fmt.Sprintf("%s-%s", institutionAbbr, courseAbbr)
}

func getInstitutionAbbr(name string) string {
	words := splitWords(name)
	if len(words) == 0 {
		return "UNKN" // fallback abbreviation
	}
	return strings.ToUpper(abbreviate(words, 4))
}

func getCourseAbbr(name string) string {
	words := splitWords(name)
	if len(words) == 0 {
		return "GEN" // fallback abbreviation
	}
	return strings.ToUpper(abbreviate(words, 3))
}

func splitWords(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.FindAllString(text, -1)
}

func abbreviate(words []string, length int) string {
	var abbr string

	// First: try to use initials
	for _, word := range words {
		if len(word) > 0 {
			abbr += string(word[0])
			if len(abbr) >= length {
				return abbr[:length]
			}
		}
	}

	// If only one word and it's short, pad from inside the word
	if len(words) == 1 {
		word := words[0]
		for i := 1; i < len(word) && len(abbr) < length; i++ {
			abbr += string(word[i])
		}
	}

	// If still not long enough, pad with 'X'
	for len(abbr) < length {
		abbr += "X"
	}

	return abbr[:length]
}

func ConvertDateIntoTime(dateString string) (*time.Time, error) {
	if dateString == "" {
		return nil, errors.New("dateString is empty")
	}

	// Define the date format
	layout := "2006-01-02"

	// Parse the start date
	formattedDate, err := time.Parse(layout, dateString)
	if err != nil {
		logrus.Errorf("Error parsing date: %v", err)
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	// Return the formatted dates
	return &formattedDate, nil
}

func CheckUserExistsByUsernameOrEmail(ctx context.Context, username, email string) (*models.User, bool, error) {
	// check if user exists by username first
	condition := &models.User{
		Username: username,
	}

	user, err := postgres.DbService.DAO.Get(ctx, condition)
	if err != nil {
		// check if user record was not found for username
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			// check if user exists by email
			condition = &models.User{
				Email: email,
			}
			emailUser, err := postgres.DbService.DAO.Get(ctx, condition)
			if err != nil {
				if err.Error() == gorm.ErrRecordNotFound.Error() {
					// user doesn't exist at all, so return false
					return nil, false, nil
				} else {
					return nil, false, errors.New("error retrieving user by email")
				}
			}
			if emailUser != nil {
				return emailUser.(*models.User), true, nil
			}
		} else {
			return nil, false, errors.New("error retrieving user by username")
		}
	}

	if user != nil {
		return user.(*models.User), true, nil
	}
	return nil, false, nil
}
