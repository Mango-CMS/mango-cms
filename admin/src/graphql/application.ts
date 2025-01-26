import { gql } from '@apollo/client/core'

// 应用模块类型定义
export interface ApplicationField {
  id: string
  name: string
  slug: string
  type: string
  required: boolean
  description: string
  default: string
  validation: string
}

export interface ApplicationModel {
  id: string
  name: string
  slug: string
  description: string
  fields: ApplicationField[]
  content?: string
  createdAt: string
  updatedAt: string
}

export interface Application {
  id: string
  name: string
  slug: string
  sign: string
  description: string
  fields: ApplicationField[]
  models: ApplicationModel[]
  permissions: ApplicationPermission[]
  status: string
  createdAt: string
  updatedAt: string
}

export interface ApplicationPermission {
  id: string
  applicationId: string
  roleId: string
  permissions: string[]
  createdAt: string
  updatedAt: string
}

// 查询操作
export const GET_APPLICATIONS = gql`
  query GetApplications {
    applications {
      id
      name
      slug
      description
      fields {
        id
        name
        slug
        type
        required
        description
        default
        validation
      }
      models {
        id
        name
        slug
        description
        fields {
          id
          name
          slug
          type
          required
          description
          default
          validation
        }
        content
        createdAt
        updatedAt
      }
      status
      createdAt
      updatedAt
    }
  }
`

export const GET_APPLICATION = gql`
  query GetApplication($id: String!) {
    application(id: $id) {
      id
      name
      slug
      sign
      description
      fields {
        id
        name
        slug
        type
        required
        description
        default
        validation
      }
      models {
        id
        name
        slug
        description
        fields {
          id
          name
          slug
          type
          required
          description
          default
          validation
        }
        content
        createdAt
        updatedAt
      }
      status
      createdAt
      updatedAt
    }
  }
`

// 变更操作
export const CREATE_APPLICATION = gql`
  mutation CreateApplication(
    $name: String!
    $slug: String!
    $description: String
    $fields: [ApplicationFieldInput]
    $status: String
  ) {
    createApplication(
      name: $name
      slug: $slug
      description: $description
      fields: $fields
      status: $status
    ) {
      id
      name
      slug
      description
      fields {
        id
        name
        slug
        type
        required
        description
        default
        validation
      }
      status
      createdAt
      updatedAt
    }
  }
`

export const UPDATE_APPLICATION = gql`
  mutation UpdateApplication(
    $id: String!
    $name: String
    $slug: String
    $description: String
    $fields: [UpdateApplicationFieldInput]
    $status: String
  ) {
    updateApplication(
      id: $id
      name: $name
      slug: $slug
      description: $description
      fields: $fields
      status: $status
    ) {
      id
      name
      slug
      description
      fields {
        id
        name
        slug
        type
        required
        description
        default
        validation
      }
      status
      createdAt
      updatedAt
    }
  }
`

export const UPDATE_APPLICATION_SIGN = gql`
  mutation UpdateApplicationSign($id: String!) {
    updateApplicationSign(id: $id) {
      id
      sign
    }
  }
`

export const DELETE_APPLICATION = gql`
  mutation DeleteApplication($id: String!) {
    deleteApplication(id: $id)
  }
`

export const SET_APPLICATION_PERMISSION = gql`
  mutation SetApplicationPermission($applicationId: ID!, $roleId: ID!, $permissions: [String!]!) {
    setApplicationPermission(
      applicationId: $applicationId
      roleId: $roleId
      permissions: $permissions
    ) {
      id
      applicationId
      roleId
      permissions
      createdAt
      updatedAt
    }
  }
`
