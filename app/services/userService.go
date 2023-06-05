package services

import (
	"context"
	"errors"
	"fmt"
	"go-folder-sample/app/helpers"
	"go-folder-sample/app/models"
	"go-folder-sample/app/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

//	func NewUserService(userRepo repositories.UserRepository, rabbitmq *messaging.RabbitMQ) *UserService {
//		return &UserService{
//			UserRepo: userRepo,
//			RabbitMQ: rabbitmq,
//		}
//	}
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: *userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, register *models.Register) (string, error) {

	filter := bson.M{
		"$or": []bson.M{
			{"username": register.Username},
			{"logins.email": register.Email},
		},
	}
	checkUniqueEmailAndUsername, err := s.UserRepo.GetOne(ctx, filter)
	if err != nil {
		return "", errors.New("errors in database")
	}
	if len(checkUniqueEmailAndUsername) != 0 {
		return "", errors.New("user Already Exist With that Username or Email")
	}
	identifier := primitive.NewObjectID().Hex()
	hashedPassword, err := helpers.EncryptPassword(register.Password)
	if err != nil {
		return "", errors.New("errors in password encryption")
	}
	err = s.UserRepo.Create(ctx, bson.M{
		"_id":             identifier,
		"identifier":      identifier,
		"username":        register.Username,
		"firstName":       register.FirstName,
		"lastName":        register.LastName,
		"status":          0,
		"follow":          []string{},
		"followedBy":      []string{},
		"followCount":     0,
		"followedByCount": 0,
		"logins": []bson.M{
			{
				"email":     register.Email,
				"password":  hashedPassword,
				"type":      "local",
				"createdAt": time.Now().UTC(),
			},
		},
	},
	)
	if err != nil {
		return "", errors.New("errors in database")
	}
	// Send Email Confirmation Logic
	return "User Account Created Successfully And Email Confirmation Send", nil
}

func (s *UserService) Login(ctx context.Context, login *models.Login) (string, error) {

	filter := bson.M{
		"status": bson.M{"$lt": 3},
		"logins": bson.M{
			"$elemMatch": bson.M{
				"email": login.Email,
				"type":  "local",
			},
		},
	}
	projection := bson.M{
		"_id":        0,
		"status":     1,
		"identifier": 1,
		"logins.$":   1,
	}

	options := options.FindOne().SetProjection(projection)
	result, err := s.UserRepo.GetOneWithOptions(ctx, filter, options)

	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("no user found with this given email")
	}
	status, ok := result["status"].(int32)
	if !ok {
		return "", errors.New("error in data")
	}
	if status == 0 {
		return "", errors.New("user Email Verification is Pending")
	}
	if status == 2 {
		return "", errors.New("account blocked")
	}

	logins, ok := result["logins"].(primitive.A)
	if !ok {
		return "", errors.New("error in data")
	}

	loginData, ok := logins[0].(primitive.M)
	if !ok {
		return "", errors.New("error in data")
	}

	hashedPassword := loginData["password"].(string)
	userPassword := login.Password
	err = helpers.VerifyPassword(hashedPassword, userPassword)
	if err != nil {
		return "", errors.New("password didn't match")
	}

	identifier := result["identifier"].(string)
	accessToken, refreshToken, err := helpers.GenerateAllTokens(identifier)
	if err != nil {
		return "", errors.New("token generation error")
	}
	token := fmt.Sprintf("%s;%s", accessToken, refreshToken)
	return token, nil

}

func (s *UserService) UserExist(ctx context.Context, userExist *models.UserExist) (string, error) {
	filter := bson.M{
		"username": userExist.Username,
	}
	checkUniqueUsername, err := s.UserRepo.GetOne(ctx, filter)
	if err != nil {
		return "", errors.New("errors in database")
	}
	fmt.Println(checkUniqueUsername)
	if len(checkUniqueUsername) != 0 {
		return "", errors.New("user Already Exist With that Username")
	}
	return "Unique Username Exist", nil
}

func (s *UserService) RefreshToken(ctx context.Context, refreshToken *models.RefreshToken) (string, string, error) {
	claims, err := helpers.ValidateRefreshToken(refreshToken.RefreshToken)
	if err != nil {
		return "", "", err
	}
	accessToken, refreshNewToken, err := helpers.GenerateAllTokens(claims.Id)
	if err != nil {
		return "", "", errors.New("token generation error")
	}
	return accessToken, refreshNewToken, nil
}
