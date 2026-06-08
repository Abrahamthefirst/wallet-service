package authservice

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// func (s *AuthService) SignUp(dto dtos.SignupRequestDto, baseURL string) (*entities.User, error) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

// 	user := entities.User{
// 		Email:    dto.Email,
// 		Password: string(hashedPassword),
// 		Username: dto.Username,
// 	}

// 	newUser, err := s.repo.Create(user)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: user with email %s already exists", entities.ErrConflict, dto.Email)
// 	}

// 	go func() {
// 		if err := s.VerifyEmail(context.Background(), newUser.Email, baseURL); err != nil {
// 			fmt.Println("Verification email error:", err)
// 		}
// 	}()

// 	return newUser, nil
// }

// func (a *AuthService) Login(ctx context.Context, dto dtos.LoginRequestDto) (*dtos.LoginServiceOkResponse, error) {
// 	user, err := a.repo.FindByEmail(ctx, dto.Email)

// 	if err != nil {
// 		if errors.Is(err, entities.RecordNotFound) {
// 			return nil, entities.ErrInvalidCredentials

// 		}
// 		return nil, err
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
// 		return nil, entities.ErrInvalidCredentials
// 	}

// 	atClaims := entities.TokenClaims{
// 		UserID: uint(user.ID),
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Issuer:    "my-auth-server",
// 		},
// 	}

// 	rtClaims := entities.TokenClaims{
// 		UserID: uint(user.ID),
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Issuer:    "my-auth-server",
// 		},
// 	}

// 	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

// 	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
// 	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

// 	if err != nil {
// 		return nil, entities.ErrInternal
// 	}

// 	return &dtos.LoginServiceOkResponse{User: user, AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil

// }
