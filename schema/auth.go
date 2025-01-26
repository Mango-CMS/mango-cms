package schema

import (
	"github.com/Mango-CMS/mango-cms/internal/auth"
	"github.com/graphql-go/graphql"
)

// AuthPayloadType 定义认证响应的GraphQL类型
var AuthPayloadType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthPayload",
	Fields: graphql.Fields{
		"token": &graphql.Field{Type: graphql.String},
		"user":  &graphql.Field{Type: UserType},
	},
})

// AuthMutation 定义认证相关的变更操作
var AuthMutation = graphql.Fields{
	"login": &graphql.Field{
		Type: AuthPayloadType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username := p.Args["username"].(string)
			password := p.Args["password"].(string)

			// 使用UserService的Login方法进行用户认证
			token, err := userService.Login(username, password)
			if err != nil {
				return nil, err
			}

			// 获取用户信息
			user, err := userService.GetUserByUsername(username)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"token": token,
				"user":  user,
			}, nil
		},
	},
	"refreshToken": &graphql.Field{
		Type: AuthPayloadType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 从上下文中获取用户信息
			claims, ok := p.Context.Value("user").(*auth.Claims)
			if !ok {
				return nil, auth.ErrUnauthorized
			}

			// 获取用户信息
			user, err := userService.GetUserByID(claims.UserID)
			if err != nil {
				return nil, err
			}

			// 生成新的JWT令牌
			token, err := auth.GenerateToken(user)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"token": token,
				"user":  user,
			}, nil
		},
	},
}
