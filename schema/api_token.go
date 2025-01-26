package schema

import (
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// APITokenType 定义API令牌的GraphQL类型
var APITokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "APIToken",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				token := p.Source.(model.APIToken)
				return token.ID.Hex(), nil
			}},
		"userId":      &graphql.Field{Type: graphql.Int},
		"name":        &graphql.Field{Type: graphql.String},
		"token":       &graphql.Field{Type: graphql.String},
		"permissions": &graphql.Field{Type: graphql.String},
		"lastUsedAt":  &graphql.Field{Type: graphql.DateTime},
		"expiresAt":   &graphql.Field{Type: graphql.DateTime},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
		"user":        &graphql.Field{Type: UserType},
	},
})

// APITokenQuery 定义API令牌相关的查询
var APITokenQuery = graphql.Fields{
	"apiToken": &graphql.Field{
		Type: APITokenType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			idObj, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}
			return apiTokenService.GetAPITokenByID(idObj)
		},
	},
	"apiTokens": &graphql.Field{
		Type: graphql.NewList(APITokenType),
		Args: graphql.FieldConfigArgument{
			"userId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userID := p.Args["userId"].(string)
			userIDObj, err := primitive.ObjectIDFromHex(userID)
			if err != nil {
				return nil, err
			}
			return apiTokenService.GetAPITokensByUserID(userIDObj)
		},
	},
}

// APITokenMutation 定义API令牌相关的变更操作
var APITokenMutation = graphql.Fields{
	"createAPIToken": &graphql.Field{
		Type: APITokenType,
		Args: graphql.FieldConfigArgument{
			"userId":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"name":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"permissions": &graphql.ArgumentConfig{Type: graphql.String},
			"expiresIn":   &graphql.ArgumentConfig{Type: graphql.Int}, // 过期时间（秒）
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userID := p.Args["userId"].(string)
			userIDObj, err := primitive.ObjectIDFromHex(userID)
			if err != nil {
				return nil, err
			}
			name := p.Args["name"].(string)
			permissions := ""
			if perms, ok := p.Args["permissions"].(string); ok {
				permissions = perms
			}
			expiresIn := time.Hour * 24 * 30 // 默认30天
			if exp, ok := p.Args["expiresIn"].(int); ok {
				expiresIn = time.Duration(exp) * time.Second
			}
			return apiTokenService.CreateAPIToken(userIDObj, name, permissions, expiresIn)
		},
	},
	"revokeAPIToken": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			idObj, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return false, err
			}
			if err := apiTokenService.RevokeToken(idObj); err != nil {
				return false, err
			}
			return true, nil
		},
	},
}
