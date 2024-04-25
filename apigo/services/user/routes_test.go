package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codezeron/apigo/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if user payload is invalid", func(t *testing.T){
				payload := types.RegisterUser{
					FirstName: "user",
					LastName: "123",
					Email: "invalid",
					Password: "aaaaaa",
				}
				marshalled, _ := json.Marshal(payload)
			
				req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				router := mux.NewRouter()
			
				router.HandleFunc("/register", handler.handleRegister)
				router.ServeHTTP(rr, req)
			
				if rr.Code != http.StatusBadRequest {
					t.Errorf("expected status code %d, status %d", http.StatusBadRequest,rr.Code)
				}
		})
	t.Run("should correctly register the user", func(t *testing.T){
		payload := types.RegisterUser{
			FirstName: "user",
			LastName: "123",
			Email: "valid@gmail.com",
			Password: "aaaaaa",
		}
		marshalled, _ := json.Marshal(payload)
	
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
	
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
	
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, status %d", http.StatusCreated,rr.Code)
		}
	})
}
type mockUserStore struct{}

// CreateUser implements types.UserStore.
func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

// GetUserByID implements types.UserStore.
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}


