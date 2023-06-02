package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABC"

func init() {
	rand.Seed(time.Now().UnixNano()) //
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// this is a random comment
// another random comment
func RandomOwner() string {
	return RandomString(6)
}

func RandomCompanyName() string {
	return RandomString(4)
}
func RandomIndustry() string {
	industries := []string{"Education", "Construction", "telecomunication", "Business", "Environment", "Engineering", "Architecture"}
	n := len(industries)

	return industries[rand.Intn(n)]

}
func RandomPosition() string {
	positions := []string{"Management", "CEO", "Director", "MiddleManager", "Sales Manager", "Developer", "CTO", "Marketing Manager"}
	n := len(positions)
	return positions[rand.Intn(n)]

}
func RandomTargetSize() int64 {
	return RandomInt(0, 3000)
}

func RandomEmail() string {
	company := RandomCompanyName()
	domain := RandomString(5) + ".com" // Adjust domain length as needed
	email := company + "@" + domain
	return email
}

func RandomTelephoneNumber() string {
	var sb strings.Builder

	// Generate random area code (3 digits)
	areaCode := strconv.Itoa(rand.Intn(900) + 100)
	sb.WriteString(areaCode)
	sb.WriteString("-")

	// Generate random exchange code (3 digits)
	exchangeCode := strconv.Itoa(rand.Intn(900) + 100)
	sb.WriteString(exchangeCode)
	sb.WriteString("-")

	// Generate random subscriber number (4 digits)
	subscriberNumber := strconv.Itoa(rand.Intn(9000) + 1000)
	sb.WriteString(subscriberNumber)

	return sb.String()
}

func RandomShortText(length int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
