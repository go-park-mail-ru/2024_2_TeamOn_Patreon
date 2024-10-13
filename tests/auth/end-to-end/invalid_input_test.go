package end_to_end

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

// Тест на невалидные данные
func TestInvalidInput(t *testing.T) {
	var testInvalidInput = []struct {
		name            string
		username        string
		password        string
		forOnlyRegister bool
	}{
		{name: "too small username",
			username: "tes",
			password: "Test@1234",
		},

		{name: "too long username",
			username: "test_test_test_test_test",
			password: "Test@1234",
		},

		{name: "username with space",
			username: "Test t",
			password: "Test@1234",
		},

		{name: "username with special characters",
			username: "Test!",
			password: "Test@1234",
		},

		{name: "password is too small",
			username: "test",
			password: "Test@14",
		},

		{name: "password is too long",
			username: "test",
			password: strings.Repeat("A", 65),
		},

		{name: "password with space",
			username: "test",
			password: "Test@ 1234",
		},

		{name: "password without special characters",
			username:        "test",
			password:        "Test01234",
			forOnlyRegister: true,
		},

		{name: "password without digital",
			username:        "test",
			password:        "Test_user",
			forOnlyRegister: true,
		},

		{name: "password without big char",
			username:        "test",
			password:        "test@1234",
			forOnlyRegister: true,
		},

		{name: "password without small char",
			username:        "test",
			password:        "TEST@1234",
			forOnlyRegister: true,
		},
	}

	// test end-to-end
	ts := SetupTestServer()
	defer ts.TearDown()

	for _, tempTest := range testInvalidInput {
		t.Run(tempTest.name, func(t *testing.T) {
			// Тело запроса с невалидными данными
			reqBody := map[string]string{
				"username": tempTest.username,
				"password": tempTest.password,
			}

			// Создание POST request к /auth/register
			resp := ts.MakeRequest(t, "POST", "/auth/register", reqBody)
			defer resp.Body.Close()

			// Статус 400
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400")

			if !tempTest.forOnlyRegister {
				// Создание POST request к /auth/login
				resp := ts.MakeRequest(t, "POST", "/auth/login", reqBody)
				defer resp.Body.Close()

				// Статус 400
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400")
			}
		})
	}
}
