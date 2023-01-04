package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	constant "employee/app/constants"
	form "employee/app/interface"
	db "employee/connect"
	"employee/models"
	"employee/utils"
)

var UserEntity IUser

type userEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IUser interface {
	GetUserByEmail(email string) (*models.User, int, error)
	CreateOne(userForm form.UserLogin) (*models.User, int, error)
	FindUserById(id string) (*models.User, int, error)
	FindAllUser() ([]*models.User, int, error)
	DeleteUserById(id string) (int, error)
	Login(userForm form.User) (*form.LoginResponse, int, error)
}

func NewUserEntity(resource *db.Resource) IUser {
	userRepo := resource.DB.Collection("user")
	UserEntity = &userEntity{resource: resource, repo: userRepo}
	return UserEntity
}

func (entity *userEntity) FindUserById(id string) (*models.User, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var user models.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}
	err = entity.repo.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}
	return &user, http.StatusOK, nil
}

func (entity *userEntity) GetUserByEmail(email string) (*models.User, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var user models.User
	err :=
		entity.repo.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		logrus.Print(err)
		return nil, http.StatusNotFound, errors.New("Email was not found")
	}

	return &user, http.StatusOK, nil
}

func (entity *userEntity) CreateOne(userForm form.UserLogin) (*models.User, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	randomPass := randstr.Hex(6)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(randomPass), bcrypt.DefaultCost)
	user := models.User{
		Id:       primitive.NewObjectID(),
		Email:    userForm.Email,
		Password: string(hashedPassword),
		Roles:    []string{constant.CUSTOMER},
	}
	found, _, _ := entity.GetUserByEmail(user.Email)
	if found != nil {
		return nil, http.StatusBadRequest, errors.New("Email already exists")
	}
	_, err := entity.repo.InsertOne(ctx, user)

	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}
	data := &form.User{
		Email:    user.Email,
		Password: randomPass,
	}

	go SendEmail(data)
	return &user, http.StatusOK, nil
}

func (entity *userEntity) FindAllUser() ([]*models.User, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var users []*models.User
	cursor, err := entity.repo.Find(ctx, bson.M{})
	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	return users, http.StatusOK, nil
}

func (entity *userEntity) DeleteUserById(id string) (int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Print(err)
		return 400, err
	}
	_, err = entity.repo.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		logrus.Print(err)
		return 400, err
	}

	return http.StatusOK, nil
}

func (entity *userEntity) Login(userForm form.User) (*form.LoginResponse, int, error) {
	user, _, _ := entity.GetUserByEmail(userForm.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userForm.Password))
	if err != nil {
		logrus.Print(err)
		return nil, http.StatusBadRequest, errors.New("Wrong email or password")
	}

	token := utils.GenerateToken(&jwt.MapClaims{
		"userId": user.Id.Hex(),
		"email":  user.Email,
		"roles":  user.Roles,
		"exp":    time.Now().Add(720 * time.Hour).Unix(),
	})

	return &form.LoginResponse{
		Email: user.Email,
		Token: token,
	}, http.StatusOK, nil
}
