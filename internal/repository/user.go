package repository

import (
	"context"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetUserByID 根据ID获取用户
func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = DB.Collection(UsersCollection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Collection(UsersCollection).FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(user *model.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := DB.Collection(UsersCollection).InsertOne(context.Background(), user)
	return err
}

// UpdateUser 更新用户信息
func (r *UserRepository) UpdateUser(user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := DB.Collection(UsersCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

// DeleteUser 删除用户
func (r *UserRepository) DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = DB.Collection(UsersCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"deleted_at": time.Now()}},
	)
	return err
}

// GetUsers 获取所有用户列表
func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	cursor, err := DB.Collection(UsersCollection).Find(context.Background(), bson.M{"deleted_at": nil})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByRole 根据角色获取用户列表
func (r *UserRepository) GetUsersByRole(role string) ([]model.User, error) {
	var users []model.User
	cursor, err := DB.Collection(UsersCollection).Find(
		context.Background(),
		bson.M{"role": role, "deleted_at": nil},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
