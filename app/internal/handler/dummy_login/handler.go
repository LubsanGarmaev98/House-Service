package dummy_login

import (
	"encoding/json"
	"github.com/timurzdev/mentorship-test-task/internal/generated"
	"github.com/timurzdev/mentorship-test-task/internal/handler"
	"github.com/timurzdev/mentorship-test-task/internal/service/token"
	"net/http"
)

type Handler struct {
	tokenService *token.TokenService
}

func NewHandler(
	tokenService *token.TokenService,
) *Handler {
	return &Handler{
		tokenService: tokenService,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request, params generated.GetDummyLoginParams) {
	entityToken, err := h.tokenService.GenerateTokenForDummy(params.UserType)
	if err != nil {
		handler.ErrorResponse(w, err)
		return
	}

	generateToken := generated.Token(entityToken.Token)

	bytes, err := json.Marshal(generateToken)
	if err != nil {
		handler.ErrorResponse(w, err)
		return
	}

	handler.SuccessResponse(w, bytes)
}
