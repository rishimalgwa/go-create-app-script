package filedata

func GetTokenData() string {
	return `package utils

	import (
		"fmt"
		"time"
	
		"github.com/golang-jwt/jwt"
		"go.mongodb.org/mongo-driver/bson/primitive"
	)
	
	type TokenPayload struct {
		Id   primitive.ObjectID
		Role string
	}
	
	func GenerateToken(ttl time.Duration, payload TokenPayload, secretJWTKey string) (string, error) {
		token := jwt.New(jwt.SigningMethodHS256)
	
		now := time.Now().UTC()
		claims := token.Claims.(jwt.MapClaims)
	
		claims["sub"] = payload.Id
		claims["role"] = payload.Role
		claims["exp"] = now.Add(ttl).Unix()
		claims["iat"] = now.Unix()
		claims["nbf"] = now.Unix()
	
		tokenString, err := token.SignedString([]byte(secretJWTKey))
	
		if err != nil {
			return "", fmt.Errorf("generating JWT Token failed: %w", err)
		}
	
		return tokenString, nil
	}
	
	func ValidateToken(token string, signedJWTKey string) (TokenPayload, error) {
		tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}
	
			return []byte(signedJWTKey), nil
		})
		if err != nil {
			return TokenPayload{}, fmt.Errorf("invalidate token: %w", err)
		}
	
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok || !tok.Valid {
			return TokenPayload{}, fmt.Errorf("invalid token claim")
		}
		id, _ := primitive.ObjectIDFromHex(claims["sub"].(string))
		res := TokenPayload{
			Id:   id,
			Role: claims["role"].(string),
		}
		return res, nil
	}`

}
func GetValidateData() string {
	return `package utils

	import (
		"github.com/go-playground/validator/v10"
		"golang.org/x/crypto/bcrypt"
	)
	
	var validate = validator.New()
	
	type ErrorResponse struct {
		Field string "json:"field""
		Tag   string "json:"tag""
		Value string "json:"value,omitempty""
	}
	
	func ValidateStruct[T any](payload T) []*ErrorResponse {
		var errors []*ErrorResponse
		err := validate.Struct(payload)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var element ErrorResponse
				element.Field = err.StructNamespace()
				element.Tag = err.Tag()
				element.Value = err.Param()
				errors = append(errors, &element)
			}
		}
		return errors
	}
	
	func HashPassword(password string) (string, error) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	}
	
	func CheckPasswordHash(password, hash string) bool {
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return err == nil
	}
	`

}
