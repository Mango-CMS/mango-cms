package schema

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PermissionType 定义权限的GraphQL类型
var PermissionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Permission",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.String},
		"name":        &graphql.Field{Type: graphql.String},
		"slug":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"module":      &graphql.Field{Type: graphql.String},
		"action":      &graphql.Field{Type: graphql.String},
	},
})

// RolePermissionType 定义角色权限关联的GraphQL类型
var RolePermissionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RolePermission",
	Fields: graphql.Fields{
		"role":       &graphql.Field{Type: graphql.String},
		"permission": &graphql.Field{Type: PermissionType},
	},
})

// PermissionQuery 定义权限相关的查询
var PermissionQuery = graphql.Fields{
	"permissions": &graphql.Field{
		Type: graphql.NewList(PermissionType),
		Args: graphql.FieldConfigArgument{
			"module": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if module, ok := p.Args["module"].(string); ok {
				return permissionService.GetPermissionsByModule(module)
			}
			return permissionService.GetAllPermissions()
		},
	},
	"rolePermissions": &graphql.Field{
		Type: graphql.NewList(RolePermissionType),
		Args: graphql.FieldConfigArgument{
			"role": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			role := p.Args["role"].(string)
			return permissionService.GetRolePermissions(role)
		},
	},
}

// PermissionMutation 定义权限相关的变更操作
var PermissionMutation = graphql.Fields{
	"createPermission": &graphql.Field{
		Type: PermissionType,
		Args: graphql.FieldConfigArgument{
			"name":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"slug":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"description": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"module":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"action":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name := p.Args["name"].(string)
			slug := p.Args["slug"].(string)
			description := p.Args["description"].(string)
			module := p.Args["module"].(string)
			action := p.Args["action"].(string)
			return permissionService.CreatePermission(name, slug, description, module, action)
		},
	},
	"assignPermissionToRole": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"role":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"permissionId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			role := p.Args["role"].(string)
			permissionId := p.Args["permissionId"].(string)
			permissionObjID, err := primitive.ObjectIDFromHex(permissionId)
			if err != nil {
				return false, err
			}
			return true, permissionService.AssignPermissionToRole(role, permissionObjID)
		},
	},
	"revokePermissionFromRole": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"role":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"permissionId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			role := p.Args["role"].(string)
			permissionId := p.Args["permissionId"].(string)
			return true, permissionService.RevokePermissionFromRole(role, permissionId)
		},
	},
}
