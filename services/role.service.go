package service

import (
	"employee/app/enum"
	form "employee/app/interface"
	db "employee/connect"
	"employee/models"
	"employee/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var RoleEntity IRole

type roleEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IRole interface {
	CreateRole(roleForm form.Role) (*models.Role, int, error)
	FindRoleByName(name string) (*models.Role, int, error)
	FindRoleById(id string) (*models.Role, int, error)
	FindAllRoles() ([]*models.Role, int, error)
	DeleteRoleById(id string) (int, error)
	UpdatePermissionByAdmin(roleId string, permissionForm form.Permissions) (int, error)
}

func NewRoleEntity(resource *db.Resource) IRole {
	roleRepo := resource.DB.Collection("role")
	RoleEntity = &roleEntity{resource: resource, repo: roleRepo}
	return RoleEntity
}

func (r *roleEntity) CreateRole(roleForm form.Role) (*models.Role, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	found, _, _ := r.FindRoleByName(roleForm.Name)

	if found != nil {
		return nil, 400, errors.New("Role already exists")
	}

	role := &models.Role{
		Id:          primitive.NewObjectID(),
		Name:        roleForm.Name,
		Description: roleForm.Description,
		Permissions: []enum.PermissionEnum{},
	}

	_, err := r.repo.InsertOne(ctx, role)
	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}

	return role, http.StatusOK, nil
}

func (r *roleEntity) FindRoleByName(name string) (*models.Role, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var role models.Role
	err := r.repo.FindOne(ctx, bson.M{"name": name}).Decode(&role)
	if err != nil {
		logrus.Print(err)
		return nil, http.StatusBadRequest, errors.New("Role not found")
	}
	return &role, http.StatusOK, nil
}

func (r *roleEntity) FindRoleById(id string) (*models.Role, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var role models.Role
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Print(err)
		return nil, 400, err
	}
	err = r.repo.FindOne(ctx, bson.M{"_id": objectId}).Decode(&role)
	if err != nil {
		logrus.Print(err)
		return nil, http.StatusNotFound, errors.New("Role not found")
	}
	return &role, http.StatusOK, nil
}

func (r *roleEntity) FindAllRoles() ([]*models.Role, int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var roles []*models.Role
	cursor, err := r.repo.Find(ctx, bson.M{})
	if err != nil {
		logrus.Print(err)
		return nil, http.StatusBadRequest, err
	}
	if err = cursor.All(ctx, &roles); err != nil {
		logrus.Print(err)
		return nil, http.StatusBadRequest, err
	}
	return roles, http.StatusOK, nil
}

func (r *roleEntity) DeleteRoleById(id string) (int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Print(err)
		return 400, err
	}

	result, err := r.repo.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		logrus.Print(err)
		return http.StatusBadRequest, err
	}
	if result.DeletedCount == 0 {
		return http.StatusNotFound, errors.New(fmt.Sprintf("Role with id %s not found", id))
	}
	return http.StatusOK, nil
}

func (r *roleEntity) UpdatePermissionByAdmin(roleId string, permissionForm form.Permissions) (int, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		logrus.Print(err)
		return 400, err
	}
	_, err = r.repo.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"permissions": permissionForm.Permissions}})
	if err != nil {
		logrus.Print(err)
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}
