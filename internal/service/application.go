package service

import (
	"errors"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"github.com/Mango-CMS/mango-cms/internal/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApplicationService struct {
	repo *repository.ApplicationRepository
}

func NewApplicationService() *ApplicationService {
	return &ApplicationService{repo: repository.NewApplicationRepository()}
}

// GetApplicationByID 根据ID获取应用模块
func (s *ApplicationService) GetApplicationByID(id string) (model.Application, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Application{}, err
	}
	application, err := s.repo.GetApplicationByID(objID)
	if err != nil {
		return model.Application{}, err
	}

	return application, nil
}

// GetApplicationBySlug 根据Slug获取应用模块
func (s *ApplicationService) GetApplicationBySlug(slug string) (model.Application, error) {
	return s.repo.GetApplicationBySlug(slug)
}

// GetApplications 获取所有应用模块
func (s *ApplicationService) GetApplications() ([]model.Application, error) {
	applications, err := s.repo.GetApplications()
	if err != nil {
		return nil, err
	}
	return applications, nil
}

// CreateApplication 创建应用模块
func (s *ApplicationService) CreateApplication(application *model.Application) (*model.Application, error) {
	// 检查Slug是否已存在
	exist, err := s.repo.GetApplicationBySlug(application.Slug)
	if err == nil && exist.Sign != "" {
		return nil, errors.New("应用模块标识已存在")
	}

	// 生成应用ID
	application.Sign = tools.GenerateSign(application.Slug)

	// 设置默认状态
	if application.Status == "" {
		application.Status = "active"
	}

	// 为每个字段生成ID
	for i := range application.Fields {
		application.Fields[i].ID = primitive.NewObjectID()
	}

	return application, s.repo.CreateApplication(application)
}

// UpdateApplication 更新应用模块
func (s *ApplicationService) UpdateApplication(application *model.Application) (model.Application, error) {
	// 检查应用模块是否存在
	exist, err := s.repo.GetApplicationByID(application.ID)
	if err != nil {
		return model.Application{}, errors.New("应用模块不存在")
	}

	// 如果修改了Slug，检查新的Slug是否已存在
	if application.Slug != exist.Slug {
		slugExist, err := s.repo.GetApplicationBySlug(application.Slug)
		if err == nil && slugExist.Sign != "" {
			return model.Application{}, errors.New("应用模块标识已存在")
		}
	}

	// 保留创建时间
	application.CreatedAt = exist.CreatedAt
	application.UpdatedAt = time.Now()
	return model.Application{
		ID:          application.ID,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
		Name:        application.Name,
		Slug:        application.Slug,
		Sign:        application.Sign,
		Status:      application.Status,
		Fields:      application.Fields,
		Models:      application.Models,
		Description: application.Description,
	}, s.repo.UpdateApplication(application)
}

// UpdateApplicationSign 更新应用签名
func (s *ApplicationService) UpdateApplicationSign(id string) (model.Application, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Application{}, errors.New("无效的应用ID格式")
	}

	// 获取应用信息
	app, err := s.repo.GetApplicationByID(objID)
	if err != nil {
		return model.Application{}, errors.New("应用不存在")
	}

	// 生成新的签名
	app.Sign = tools.GenerateSign(app.Slug)

	// 更新应用
	err = s.repo.UpdateApplication(&app)
	if err != nil {
		return model.Application{}, err
	}

	return app, nil
}

// DeleteApplication 删除应用模块
func (s *ApplicationService) DeleteApplication(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("无效的应用ID格式")
	}
	return s.repo.DeleteApplication(objID)
}

// GetApplicationPermissions 获取应用模块的权限配置
func (s *ApplicationService) GetApplicationPermissions(applicationID string) ([]model.ApplicationPermission, error) {
	objID, err := primitive.ObjectIDFromHex(applicationID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetApplicationPermissions(objID)
}

// SetApplicationPermission 设置应用模块的权限
func (s *ApplicationService) SetApplicationPermission(applicationID string, roleID string, permissions []string) error {
	objID, err := primitive.ObjectIDFromHex(applicationID)
	if err != nil {
		return err
	}

	// 检查应用模块是否存在
	_, err = s.repo.GetApplicationByID(objID)
	if err != nil {
		return errors.New("应用模块不存在")
	}

	permission := &model.ApplicationPermission{
		ID:            primitive.NewObjectID(),
		ApplicationID: objID,
		RoleID:        roleID,
		Permissions:   permissions,
	}

	return s.repo.SetApplicationPermission(permission)
}
