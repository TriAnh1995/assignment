package handler

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/friendsofgo/errors"
)

var validTLDs = []string{"com", "org", "net"}

func (u User) validate() error {

	if len(u.Name) == 0 {
		return errors.New("Name cannot be blank")
	}
	firstChar := rune(u.Name[0])
	if !unicode.IsUpper(firstChar) {
		return errors.New("Name Invalid")
	}

	if len(u.Email) == 0 {
		return errors.New("Email cannot be blank")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, u.Email)
	if !(match) {
		return errors.New("Email invalid")
	}
	return nil
}

func (f Friends) validate() error {

	// Check number of input and avoid repeat inputs
	if (len(f.Emails) != 2) || (f.Emails[0] == f.Emails[1]) {
		return errors.New("Please insert two different emails")
	}

	// Check Empty Emails
	if (f.Emails[0] == "") || (f.Emails[1] == "") {
		return errors.New("One of your emails is blank")
	}

	// Check Email structure
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`

	if (!regexp.MustCompile(emailPattern).MatchString(f.Emails[0])) || (!regexp.MustCompile(emailPattern).MatchString(f.Emails[1])) {
		return errors.New("One of your emails is invalid")
	}

	return nil
}

func (e FriendsList) validate() error {
	// Check Email length
	lengthIsValid := 0 < len(e.Email) && len(e.Email) <= 320
	if !(lengthIsValid) {
		return errors.New("Invalid Email Length")
	}

	// Check Email format
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, e.Email)
	if !(match) {
		return errors.New("Invalid Email Format")
	}

	// Check Email TLD
	tldRegex := regexp.MustCompile(fmt.Sprintf("\\.(%s)$", strings.Join(validTLDs, "|")))

	// Find the TLD in the email
	matches := tldRegex.FindStringSubmatch(e.Email)

	// Check if a valid TLD is found
	if len(matches) == 0 {
		return errors.New("Invalid Email TLD")
	}
	return nil
}

func (fu *FullUsers) validate() error {
	return nil
}
