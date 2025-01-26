package service

import (
	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionService struct {
	permissionRepo *repository.PermissionRepository
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		permissionRepo: repository.NewPermissionRepository(),
	}
}

// GetAllPermissions 获取所有权限
func (s *PermissionService) GetAllPermissions() ([]model.Permission, error) {
	return s.permissionRepo.GetAllPermissions()
}

// CreatePermission 创建权限
func (s *PermissionService) CreatePermission(name, slug, description, module, action string) (*model.Permission, error) {
	permission := &model.Permission{
		Name:        name,
		Slug:        slug,
		Description: description,
		Module:      module,
		Action:      action,
	}
	return s.permissionRepo.CreatePermission(permission)
}

// UpdatePermission 更新权限
func (s *PermissionService) UpdatePermission(id primitive.ObjectID, name, slug, description, module, action string) error {
	permission, err := s.GetPermissionByID(id)
	if err != nil {
		return err
	}
	permission.Name = name
	permission.Slug = slug
	permission.Description = description
	permission.Module = module
	permission.Action = action
	return s.permissionRepo.UpdatePermission(permission)
}

// DeletePermission 删除权限
func (s *PermissionService) DeletePermission(id string) error {
	return s.permissionRepo.DeletePermission(id)
}

// GetPermissionsByModule 获取指定模块的所有权限
func (s *PermissionService) GetPermissionsByModule(module string) ([]model.Permission, error) {
	return s.permissionRepo.GetPermissionsByModule(module)
}

// GetPermissionByID 根据ID获取权限
func (s *PermissionService) GetPermissionByID(id primitive.ObjectID) (*model.Permission, error) {
	return s.permissionRepo.GetPermissionByID(id)
}

// GetPermissionBySlug 根据Slug获取权限
func (s *PermissionService) GetPermissionBySlug(slug string) (*model.Permission, error) {
	return s.permissionRepo.GetPermissionBySlug(slug)
}

// GetRolePermissions 获取指定角色的所有权限
func (s *PermissionService) GetRolePermissions(role string) ([]model.RolePermission, error) {
	permissions, err := s.permissionRepo.GetRolePermissions(role)
	if err != nil {
		return nil, err
	}

	rolePermissions := make([]model.RolePermission, len(permissions))
	for i, p := range permissions {
		rolePermissions[i] = model.RolePermission{
			Role:       role,
			Permission: p,
		}
	}

	return rolePermissions, nil
}

// AssignPermissionToRole 为角色分配权限
func (s *PermissionService) AssignPermissionToRole(role string, permissionID primitive.ObjectID) error {
	// 检查权限是否存在
	_, err := s.GetPermissionByID(permissionID)
	if err != nil {
		return err
	}

	rolePermission := &model.RolePermission{
		Role:         role,
		PermissionID: permissionID,
	}

	return s.permissionRepo.AssignPermissionToRole(rolePermission)
}

// RevokePermissionFromRole 撤销角色的权限
func (s *PermissionService) RevokePermissionFromRole(role string, permissionID string) error {
	return s.permissionRepo.RevokePermissionFromRole(role, permissionID)
}
