package user

import (
	"fmt"
	"net/http"

	"github.com/codezeron/apigo/config"
	"github.com/codezeron/apigo/services/auth"
	"github.com/codezeron/apigo/types"
	"github.com/codezeron/apigo/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store 		types.UserStore 
	
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
		//pega o conteudo do JSON
		var payload types.LoginUser
		if err := utils.ParseJSON(r, &payload); err != nil{
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
			return
		}
		// validate email no db
		u, err := h.store.GetUserByEmail(payload.Email)
		if err!= nil {
      utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
      return
    }
		//compara o password 
		if !auth.ComparePassword(u.Password, []byte(payload.Password)) {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
			return
		}
		secret := []byte(config.Envies.JWTSecret)
		token, err := auth.CreateJWT(secret, u.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
      return
    }
		utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//pega o conteudo do JSON
	var user types.RegisterUser
	if err := utils.ParseJSON(r, &user); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate user
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}


	//checar se o user ja existe 
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest , fmt.Errorf("user with this email %s already exists", user.Email))
		return
	}
	//encrypt password
	hashedPassword, err:= auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	//se nao, cria um novo
	err = h.store.CreateUser(types.User{
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	utils.WriteJSON(w, http.StatusCreated, nil)
}
