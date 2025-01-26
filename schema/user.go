package schema

import (
	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/graphql-go/graphql"
)

// RoleType 定义角色的GraphQL类型
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"name":        &graphql.Field{Type: graphql.String},
		"displayName": &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
	},
})

// UserType 定义用户的GraphQL类型
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(model.User)
				return user.ID.Hex(), nil
			}},
		"username":  &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"role":      &graphql.Field{Type: graphql.String},
		"status":    &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"updatedAt": &graphql.Field{Type: graphql.DateTime},
	},
})

// UserQuery 定义用户相关的查询
var UserQuery = graphql.Fields{
	"user": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return userService.GetUserByID(id)
		},
	},
	"users": &graphql.Field{
		Type: graphql.NewList(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return userService.GetUsers()
		},
	},
	"roles": &graphql.Field{
		Type: graphql.NewList(RoleType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return []map[string]interface{}{
				{"name": "admin", "displayName": "管理员", "description": "系统管理员"},
				{"name": "editor", "displayName": "编辑", "description": "内容编辑"},
				{"name": "user", "displayName": "用户", "description": "普通用户"},
			}, nil
		},
	},
}

// UserMutation 定义用户相关的变更操作
var UserMutation = graphql.Fields{
	"createUser": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"role":     &graphql.ArgumentConfig{Type: graphql.String},
			"status":   &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := &model.User{
				Username: p.Args["username"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}
			if role, ok := p.Args["role"].(string); ok {
				user.Role = role
			}
			if status, ok := p.Args["status"].(string); ok {
				user.Status = status
			}
			if err := userService.CreateUser(user); err != nil {
				return nil, err
			}
			return user, nil
		},
	},
	"updateUser": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"username": &graphql.ArgumentConfig{Type: graphql.String},
			"email":    &graphql.ArgumentConfig{Type: graphql.String},
			"password": &graphql.ArgumentConfig{Type: graphql.String},
			"role":     &graphql.ArgumentConfig{Type: graphql.String},
			"status":   &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			user, err := userService.GetUserByID(id)
			if err != nil {
				return nil, err
			}
			if username, ok := p.Args["username"].(string); ok {
				user.Username = username
			}
			if email, ok := p.Args["email"].(string); ok {
				user.Email = email
			}
			if password, ok := p.Args["password"].(string); ok {
				user.Password = password
			}
			if role, ok := p.Args["role"].(string); ok {
				user.Role = role
			}
			if status, ok := p.Args["status"].(string); ok {
				user.Status = status
			}
			if err := userService.UpdateUser(user); err != nil {
				return nil, err
			}
			return user, nil
		},
	},
	"deleteUser": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			if err := userService.DeleteUser(id); err != nil {
				return false, err
			}
			return true, nil
		},
	},
}
