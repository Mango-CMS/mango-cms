package schema

import (
	"encoding/json"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ApplicationFieldType 定义应用字段的GraphQL类型
var ApplicationFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ApplicationField",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				field := p.Source.(model.ApplicationField)
				return field.ID.Hex(), nil
			}},
		"name":        &graphql.Field{Type: graphql.String},
		"slug":        &graphql.Field{Type: graphql.String},
		"type":        &graphql.Field{Type: graphql.String},
		"required":    &graphql.Field{Type: graphql.Boolean},
		"description": &graphql.Field{Type: graphql.String},
		"default":     &graphql.Field{Type: graphql.String},
		"validation":  &graphql.Field{Type: graphql.String},
	},
})

// ApplicationModelType 定义应用模型的GraphQL类型
var ApplicationModelType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ApplicationModel",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				model := p.Source.(model.ApplicationModel)
				return model.ID.Hex(), nil
			}},
		"name":        &graphql.Field{Type: graphql.String},
		"slug":        &graphql.Field{Type: graphql.String},
		"sign":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"fields":      &graphql.Field{Type: graphql.NewList(ApplicationFieldType)},
		"content":     &graphql.Field{Type: JSONScalar},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
	},
})

// ApplicationType 定义应用的GraphQL类型
var ApplicationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Application",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				app := p.Source.(model.Application)
				return app.ID.Hex(), nil
			}},
		"name":        &graphql.Field{Type: graphql.String},
		"slug":        &graphql.Field{Type: graphql.String},
		"sign":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"fields":      &graphql.Field{Type: graphql.NewList(ApplicationFieldType)},
		"models":      &graphql.Field{Type: graphql.NewList(ApplicationModelType)},
		"status":      &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
	},
})

// JSONScalar 定义JSON类型的GraphQL标量
var JSONScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "The `JSON` scalar type represents JSON values as specified by [ECMA-404]",
	// 序列化：将 Go 值转换为 JSON
	Serialize: func(value interface{}) interface{} {
		return value
	},
	// 反序列化：将输入值解析为 Go 值
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	// 解析字面值
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			var result interface{}
			err := json.Unmarshal([]byte(valueAST.Value), &result)
			if err != nil {
				return nil
			}
			return result
		}
		return nil
	},
})

// ApplicationQuery 定义应用相关的查询
var ApplicationQuery = graphql.Fields{
	"application": &graphql.Field{
		Type: ApplicationType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return applicationService.GetApplicationByID(id)
		},
	},
	"applications": &graphql.Field{
		Type: graphql.NewList(ApplicationType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return applicationService.GetApplications()
		},
	},
}

// ApplicationMutation 定义应用相关的变更操作
var ApplicationMutation = graphql.Fields{
	"createApplication": &graphql.Field{
		Type: ApplicationType,
		Args: graphql.FieldConfigArgument{
			"name":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"slug":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"description": &graphql.ArgumentConfig{Type: graphql.String},
			"fields": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "ApplicationFieldInput",
				Fields: graphql.InputObjectConfigFieldMap{
					"name":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
					"slug":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
					"type":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
					"required":    &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
					"default":     &graphql.InputObjectFieldConfig{Type: graphql.String},
					"validation":  &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
			}))},
			"status": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			app := model.Application{
				Name: p.Args["name"].(string),
				Slug: p.Args["slug"].(string),
			}

			if description, ok := p.Args["description"].(string); ok {
				app.Description = description
			}

			if fields, ok := p.Args["fields"].([]interface{}); ok {
				appFields := make([]model.ApplicationField, len(fields))
				for i, field := range fields {
					if fieldMap, ok := field.(map[string]interface{}); ok {
						appField := model.ApplicationField{}
						if name, ok := fieldMap["name"].(string); ok {
							appField.Name = name
						}
						if slug, ok := fieldMap["slug"].(string); ok {
							appField.Slug = slug
						}
						if fieldType, ok := fieldMap["type"].(string); ok {
							appField.Type = fieldType
						}
						if required, ok := fieldMap["required"].(bool); ok {
							appField.Required = required
						}
						if description, ok := fieldMap["description"].(string); ok {
							appField.Description = description
						}
						if defaultVal, ok := fieldMap["default"].(string); ok {
							appField.Default = defaultVal
						}
						if validation, ok := fieldMap["validation"].(string); ok {
							appField.Validation = validation
						}
						appFields[i] = appField
					}
				}
				app.Fields = appFields
			}

			if status, ok := p.Args["status"].(string); ok {
				app.Status = status
			}

			return applicationService.CreateApplication(&app)
		},
	},
	"updateApplication": &graphql.Field{
		Type: ApplicationType,
		Args: graphql.FieldConfigArgument{
			"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			"name":        &graphql.ArgumentConfig{Type: graphql.String},
			"slug":        &graphql.ArgumentConfig{Type: graphql.String},
			"description": &graphql.ArgumentConfig{Type: graphql.String},
			"fields": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "UpdateApplicationFieldInput",
				Fields: graphql.InputObjectConfigFieldMap{
					"id":          &graphql.InputObjectFieldConfig{Type: graphql.ID},
					"name":        &graphql.InputObjectFieldConfig{Type: graphql.String},
					"slug":        &graphql.InputObjectFieldConfig{Type: graphql.String},
					"type":        &graphql.InputObjectFieldConfig{Type: graphql.String},
					"required":    &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
					"default":     &graphql.InputObjectFieldConfig{Type: graphql.String},
					"validation":  &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
			}))},
			"status": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			idObj, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}

			app := model.Application{
				ID: idObj,
			}

			if val, ok := p.Args["name"].(string); ok {
				app.Name = val
			}
			if val, ok := p.Args["slug"].(string); ok {
				app.Slug = val
			}
			if val, ok := p.Args["description"].(string); ok {
				app.Description = val
			}
			if val, ok := p.Args["status"].(string); ok {
				app.Status = val
			}

			if fieldsArg, ok := p.Args["fields"].([]interface{}); ok {
				fields := make([]model.ApplicationField, len(fieldsArg))
				for i, field := range fieldsArg {
					if fieldMap, ok := field.(map[string]interface{}); ok {
						appField := model.ApplicationField{}
						if id, ok := fieldMap["id"].(string); ok {
							objID, err := primitive.ObjectIDFromHex(id)
							if err == nil {
								appField.ID = objID
							}
						}
						if name, ok := fieldMap["name"].(string); ok {
							appField.Name = name
						}
						if slug, ok := fieldMap["slug"].(string); ok {
							appField.Slug = slug
						}
						if fieldType, ok := fieldMap["type"].(string); ok {
							appField.Type = fieldType
						}
						if required, ok := fieldMap["required"].(bool); ok {
							appField.Required = required
						}
						if description, ok := fieldMap["description"].(string); ok {
							appField.Description = description
						}
						if defaultVal, ok := fieldMap["default"].(string); ok {
							appField.Default = defaultVal
						}
						if validation, ok := fieldMap["validation"].(string); ok {
							appField.Validation = validation
						}
						fields[i] = appField
					}
				}
				app.Fields = fields
			}

			return applicationService.UpdateApplication(&app)
		},
	},
	"updateApplicationSign": &graphql.Field{
		Type: ApplicationType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return applicationService.UpdateApplicationSign(id)
		},
	},
	"deleteApplication": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			if err := applicationService.DeleteApplication(id); err != nil {
				return false, err
			}
			return true, nil
		},
	},
}
