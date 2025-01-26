package repository

import (
	"context"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const RolePermissionsCollection = "role_permissions"

type PermissionRepository struct{}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

// GetAllPermissions 获取所有权限
func (r *PermissionRepository) GetAllPermissions() ([]model.Permission, error) {
	var permissions []model.Permission
	cursor, err := DB.Collection(PermissionsCollection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}

// GetPermissionByID 根据ID获取权限
func (r *PermissionRepository) GetPermissionByID(id primitive.ObjectID) (*model.Permission, error) {
	var permission model.Permission

	err := DB.Collection(PermissionsCollection).FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&permission)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetPermissionBySlug 根据Slug获取权限
func (r *PermissionRepository) GetPermissionBySlug(slug string) (*model.Permission, error) {
	var permission model.Permission
	err := DB.Collection(PermissionsCollection).FindOne(
		context.Background(),
		bson.M{"slug": slug},
	).Decode(&permission)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetPermissionsByModule 获取指定模块的所有权限
func (r *PermissionRepository) GetPermissionsByModule(module string) ([]model.Permission, error) {
	var permissions []model.Permission
	cursor, err := DB.Collection(PermissionsCollection).Find(
		context.Background(),
		bson.M{"module": module},
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

// GetRolePermissions 获取指定角色的所有权限
func (r *PermissionRepository) GetRolePermissions(role string) ([]model.Permission, error) {
	var rolePermissions []model.RolePermission
	cursor, err := DB.Collection(RolePermissionsCollection).Find(
		context.Background(),
		bson.M{"role": role},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &rolePermissions); err != nil {
		return nil, err
	}

	var permissions []model.Permission
	for _, rp := range rolePermissions {
		permission, err := r.GetPermissionByID(rp.PermissionID)
		if err != nil {
			continue
		}
		permissions = append(permissions, *permission)
	}

	return permissions, nil
}

// CreatePermission 创建新权限
func (r *PermissionRepository) CreatePermission(permission *model.Permission) (*model.Permission, error) {
	permission.ID = primitive.NewObjectID()
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()

	_, err := DB.Collection(PermissionsCollection).InsertOne(context.Background(), permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

// UpdatePermission 更新权限信息
func (r *PermissionRepository) UpdatePermission(permission *model.Permission) error {
	permission.UpdatedAt = time.Now()
	_, err := DB.Collection(PermissionsCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": permission.ID},
		bson.M{"$set": permission},
	)
	return err
}

// DeletePermission 删除权限
func (r *PermissionRepository) DeletePermission(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = DB.Collection(PermissionsCollection).DeleteOne(
		context.Background(),
		bson.M{"_id": objID},
	)
	return err
}

// AssignPermissionToRole 为角色分配权限
func (r *PermissionRepository) AssignPermissionToRole(rolePermission *model.RolePermission) error {
	rolePermission.CreatedAt = time.Now()
	rolePermission.UpdatedAt = time.Now()
	_, err := DB.Collection(RolePermissionsCollection).InsertOne(context.Background(), rolePermission)
	return err
}

// RevokePermissionFromRole 撤销角色的权限
func (r *PermissionRepository) RevokePermissionFromRole(role string, permissionID string) error {
	objID, err := primitive.ObjectIDFromHex(permissionID)
	if err != nil {
		return err
	}

	_, err = DB.Collection(RolePermissionsCollection).DeleteOne(
		context.Background(),
		bson.M{"role": role, "permission_id": objID},
	)
	return err
}
