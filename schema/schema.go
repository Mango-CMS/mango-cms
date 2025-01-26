package schema

import (
	"github.com/Mango-CMS/mango-cms/internal/service"
	"github.com/graphql-go/graphql"
)

var userService *service.UserService
var apiTokenService *service.APITokenService
var permissionService *service.PermissionService
var applicationService *service.ApplicationService

// InitServices 初始化服务
func InitServices(
	us *service.UserService,
	ats *service.APITokenService,
	ps *service.PermissionService,
	as *service.ApplicationService,
) {
	userService = us
	apiTokenService = ats
	permissionService = ps
	applicationService = as
}

// Schema 定义GraphQL根Schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: mergeFields(
			UserQuery,
			APITokenQuery,
			PermissionQuery,
			ApplicationQuery,
		),
	}),
	Mutation: graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: mergeFields(
			UserMutation,
			APITokenMutation,
			AuthMutation,
			PermissionMutation,
			ApplicationMutation,
		),
	}),
})

// mergeFields 合并多个Fields
func mergeFields(fields ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}
	for _, f := range fields {
		for k, v := range f {
			merged[k] = v
		}
	}
	return merged
}
