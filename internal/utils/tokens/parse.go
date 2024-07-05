package tokens

// func parseAccessToken(accessToken string) *UserClaims {
// 	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})

// 	return parsedAccessToken.Claims.(*UserClaims)
// }

// func parseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
// 	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})

// 	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
// }

// func VerifyTokens(tokens Tokens) error {
// 	var err error

// 	userClaims := parseAccessToken(tokens.AccessToken)
// 	refreshClaims := parseRefreshToken(tokens.RefreshToken)

// 	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 	// 	return secretKey, nil
// 	// })

// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// if !token.Valid {
// 	// 	return fmt.Errorf("invalid token")
// 	// }

// 	// refresh token is expired
// 	if refreshClaims.Valid() != nil {
// 		tokens.RefreshToken, err = newRefreshToken(*refreshClaims)
// 		if err != nil {
// 			log.Fatal("error creating refresh token")
// 		}
// 	}

// 	// access token is expired
// 	if userClaims.RegisteredClaims.Valid() != nil && refreshClaims.Valid() == nil {
// 		tokens.AccessToken, err = newToken(*userClaims)
// 		if err != nil {
// 			log.Fatal("error creating access token")
// 		}
// 	}

// 	return nil
// }
