package repository

import (
	"context"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ApplicationsCollection = "applications"
const ApplicationPermissionsCollection = "application_permissions"

type ApplicationRepository struct{}

func NewApplicationRepository() *ApplicationRepository {
	return &ApplicationRepository{}
}

// GetApplicationByID 根据ID获取应用模块
func (r *ApplicationRepository) GetApplicationByID(id primitive.ObjectID) (model.Application, error) {
	var application model.Application
	err := DB.Collection(ApplicationsCollection).FindOne(
		context.Background(),
		bson.M{"_id": id, "deleted_at": nil},
	).Decode(&application)
	if err != nil {
		return model.Application{}, err
	}
	return application, nil
}

// GetApplicationBySlug 根据Slug获取应用模块
func (r *ApplicationRepository) GetApplicationBySlug(slug string) (model.Application, error) {
	var application model.Application
	err := DB.Collection(ApplicationsCollection).FindOne(
		context.Background(),
		bson.M{"slug": slug, "deleted_at": nil},
	).Decode(&application)
	if err != nil {
		return model.Application{}, err
	}
	return application, nil
}

// GetApplications 获取所有应用模块
func (r *ApplicationRepository) GetApplications() ([]model.Application, error) {
	var applications []model.Application
	cursor, err := DB.Collection(ApplicationsCollection).Find(
		context.Background(),
		bson.M{"deleted_at": nil},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &applications); err != nil {
		return nil, err
	}
	return applications, nil
}

// CreateApplication 创建应用模块
func (r *ApplicationRepository) CreateApplication(application *model.Application) error {
	application.ID = primitive.NewObjectID()
	application.CreatedAt = time.Now()
	application.UpdatedAt = time.Now()

	_, err := DB.Collection(ApplicationsCollection).InsertOne(context.Background(), application)
	return err
}

// UpdateApplication 更新应用模块
func (r *ApplicationRepository) UpdateApplication(application *model.Application) error {
	application.UpdatedAt = time.Now()
	_, err := DB.Collection(ApplicationsCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": application.ID},
		bson.M{"$set": application},
	)
	return err
}

// DeleteApplication 删除应用模块
func (r *ApplicationRepository) DeleteApplication(id primitive.ObjectID) error {
	_, err := DB.Collection(ApplicationsCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"deleted_at": time.Now()}},
	)
	return err
}

// GetApplicationPermissions 获取应用模块的权限配置
func (r *ApplicationRepository) GetApplicationPermissions(applicationID primitive.ObjectID) ([]model.ApplicationPermission, error) {
	var permissions []model.ApplicationPermission
	cursor, err := DB.Collection(ApplicationPermissionsCollection).Find(
		context.Background(),
		bson.M{"application_id": applicationID},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}

// SetApplicationPermission 设置应用模块的权限
func (r *ApplicationRepository) SetApplicationPermission(permission *model.ApplicationPermission) error {
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()

	_, err := DB.Collection(ApplicationPermissionsCollection).UpdateOne(
		context.Background(),
		bson.M{
			"application_id": permission.ApplicationID,
			"role_id":        permission.RoleID,
		},
		bson.M{"$set": permission},
		options.Update().SetUpsert(true),
	)
	return err
}
