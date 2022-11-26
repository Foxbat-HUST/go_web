package implement

import (
	"fmt"
	"go_web/config"
	"go_web/domain/entity"
	"go_web/domain/repository"
	"go_web/domain/service"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	cfg      *config.Config
	userRepo repository.UserRepo
}

type claim struct {
	UserId    uint32          `json:"userId"`
	UserEmail string          `json:"userEmail"`
	UserType  entity.UserType `json:"userType"`
	jwt.RegisteredClaims
}

func NewAuthService(cfg *config.Config, userRepo repository.UserRepo) service.AuthService {
	return &authServiceImpl{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (a authServiceImpl) ValidateLoginForm(form entity.LoginForm) error {
	return form.Validate()
}
func (a authServiceImpl) AuthWithLoginForm(form entity.LoginForm) error {
	authUser, err := a.userRepo.FindAuthUserByEmail(form.Email)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(form.Password))
}
func (a authServiceImpl) HashPassword(rawPassword string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashPassword), err
}
func (a authServiceImpl) CreateToken(user entity.User) (string, error) {
	expiredAt := time.Now().Add(time.Second * time.Duration(a.cfg.Auth.TokenExpireSeconds))
	cml := claim{
		UserId:    user.ID,
		UserEmail: user.Email,
		UserType:  user.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.cfg.Auth.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, cml)
	tokenStr, err := token.SignedString([]byte(a.cfg.Auth.Secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a authServiceImpl) ParseToken(tokenStr string) (*entity.User, error) {
	clm := &claim{}
	token, err := jwt.ParseWithClaims(tokenStr, clm, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unsupported token")
		}

		return []byte(a.cfg.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if clm, ok := token.Claims.(*claim); ok && token.Valid {
		return &entity.User{
			ID:    clm.UserId,
			Email: clm.UserEmail,
			Type:  clm.UserType,
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
func (a authServiceImpl) InvalidToken() error {
	return nil
}
