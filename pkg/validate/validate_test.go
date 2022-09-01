package validate

import "testing"

func TestValidateEmail(t *testing.T) {

	type Testable struct {
		email string
		res   bool
	}

	tt := []Testable{
		{email: "tenderlis@mail.com", res: true},
		{email: "funcable@", res: false},
	}

	for _, test := range tt {
		if test.res != ValidateEmail(test.email) {
			t.Errorf("failed on %s", test.email)
		}
	}

}