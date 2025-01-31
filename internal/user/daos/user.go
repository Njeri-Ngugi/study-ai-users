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
	// Get 4-letter abbreviation for the institution
	institutionAbbr := getInstitutionAbbr(institutionName)

	// Get 3-letter abbreviation for the course
	courseAbbr := getCourseAbbr(courseName)

	// Combine to form the course code
	return fmt.Sprintf("%s-%s", institutionAbbr, courseAbbr)
}

// Extracts a 4-letter abbreviation from the institution name
func getInstitutionAbbr(name string) string {
	words := splitWords(name)
	abbr := strings.ToUpper(abbreviate(words, 4)) // Get 4-letter abbreviation
	return abbr
}

// Extracts a 3-letter abbreviation from the course name
func getCourseAbbr(name string) string {
	words := splitWords(name)
	abbr := strings.ToUpper(abbreviate(words, 3)) // Get 3-letter abbreviation
	return abbr
}

// Splits a string into words, removing non-alphabetic characters
func splitWords(text string) []string {
	// Remove non-alphabetic characters and split by spaces
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.FindAllString(text, -1)
}

// Generates an abbreviation from a list of words, up to `length` characters
func abbreviate(words []string, length int) string {
	var abbr string

	// Use first letters of words if possible
	for _, word := range words {
		abbr += string(word[0])
		if len(abbr) >= length {
			return abbr[:length]
		}
	}

	// If abbreviation is still short, add second-last letter from the last word
	if len(abbr) < length {
		lastWord := words[len(words)-1] // Get the last word
		secondLastLetter := getSecondLastLetter(lastWord)

		if secondLastLetter != "" {
			abbr += secondLastLetter
		}
	}

	return abbr[:length]
}

// Gets the second last letter from a word
func getSecondLastLetter(word string) string {
	if len(word) < 2 {
		return string(word[len(word)-1])
	}
	return string(word[len(word)-2])
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
