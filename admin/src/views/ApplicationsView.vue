<template>
  <div class="p-4">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">应用模块管理</h1>
      <button
        @click="showCreateDialog = true"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
      >
        创建应用模块
      </button>
    </div>

    <!-- 应用模块列表 -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">名称</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">标识</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">状态</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">操作</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="app in applications" :key="app.id">
            <td class="px-6 py-4 whitespace-nowrap">{{ app.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap">{{ app.slug }}</td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                :class="{
                  'px-2 py-1 text-xs rounded-full':
                    true,
                  'bg-green-100 text-green-800': app.status === 'active',
                  'bg-gray-100 text-gray-800': app.status === 'inactive'
                }"
              >
                {{ app.status === 'active' ? '启用' : '禁用' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">{{ new Date(app.createdAt).toLocaleString() }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <button
                @click="editApplication(app)"
                class="text-blue-600 hover:text-blue-900 mr-2"
              >
                编辑
              </button>
              <button
                @click="showPermissionDialog(app)"
                class="text-green-600 hover:text-green-900 mr-2"
              >
                权限
              </button>
              <button
                @click="deleteApplication(app.id)"
                class="text-red-600 hover:text-red-900"
              >
                删除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 创建/编辑对话框 -->
    <div v-if="showCreateDialog || showEditDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div class="bg-white p-6 rounded-lg w-full max-w-2xl">
        <h2 class="text-xl font-bold mb-4">{{ showEditDialog ? '编辑应用模块' : '创建应用模块' }}</h2>
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">名称</label>
              <input
                v-model="form.name"
                type="text"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                required
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">标识</label>
              <input
                v-model="form.slug"
                type="text"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                required
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">描述</label>
              <textarea
                v-model="form.description"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                rows="3"
              ></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">状态</label>
              <select
                v-model="form.status"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
              >
                <option value="active">启用</option>
                <option value="inactive">禁用</option>
              </select>
            </div>

            <!-- 字段列表 -->
            <div>
              <div class="flex justify-between items-center mb-2">
                <label class="block text-sm font-medium text-gray-700">字段列表</label>
                <button
                  type="button"
                  @click="addField"
                  class="text-blue-600 hover:text-blue-900 text-sm"
                >
                  添加字段
                </button>
              </div>
              <div v-for="(field, index) in form.fields" :key="index" class="border p-4 rounded-md mb-2">
                <div class="grid grid-cols-2 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700">字段名称</label>
                    <input
                      v-model="field.name"
                      type="text"
                      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    >
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700">字段标识</label>
                    <input
                      v-model="field.slug"
                      type="text"
                      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    >
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700">字段类型</label>
                    <select
                      v-model="field.type"
                      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    >
                      <option value="string">文本</option>
                      <option value="number">数字</option>
                      <option value="boolean">布尔值</option>
                      <option value="date">日期</option>
                    </select>
                  </div>
                  <div class="flex items-center">
                    <label class="flex items-center">
                      <input
                        v-model="field.required"
                        type="checkbox"
                        class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      >
                      <span class="ml-2 text-sm text-gray-600">必填</span>
                    </label>
                  </div>
                </div>
                <div class="mt-2">
                  <label class="block text-sm font-medium text-gray-700">字段描述</label>
                  <input
                    v-model="field.description"
                    type="text"
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  >
                </div>
                <button
                  type="button"
                  @click="removeField(index)"
                  class="mt-2 text-red-600 hover:text-red-900 text-sm"
                >
                  删除字段
                </button>
              </div>
            </div>
          </div>

          <div class="mt-6 flex justify-end space-x-3">
            <button
              type="button"
              @click="closeDialog"
              class="px-4 py-2 border rounded-md hover:bg-gray-50"
            >
              取消
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
            >
              确定
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- 权限配置对话框 -->
    <div v-if="showPermissionDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div class="bg-white p-6 rounded-lg w-full max-w-lg">
        <h2 class="text-xl font-bold mb-4">权限配置</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">角色</label>
            <select
              v-model="permissionForm.roleId"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            >
              <option value="admin">管理员</option>
              <option value="editor">编辑</option>
              <option value="viewer">查看者</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">权限</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  type="checkbox"
                  v-model="permissionForm.permissions"
                  value="create"
                  class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                <span class="ml-2 text-sm text-gray-600">创建</span>
              </label>
              <label class="flex items-center">
                <input
                  type="checkbox"
                  v-model="permissionForm.permissions"
                  value="read"
                  class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                <span class="ml-2 text-sm text-gray-600">读取</span>
              </label>
              <label class="flex items-center">
                <input
                  type="checkbox"
                  v-model="permissionForm.permissions"
                  value="update"
                  class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                <span class="ml-2 text-sm text-gray-600">更新</span>
              </label>
              <label class="flex items-center">
                <input
                  type="checkbox"
                  v-model="permissionForm.permissions"
                  value="delete"
                  class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                <span class="ml-2 text-sm text-gray-600">删除</span>
              </label>
            </div>
          </div>
        </div>

        <div class="mt-6 flex justify-end space-x-3">
          <button
            @click="closePermissionDialog"
            class="px-4 py-2 border rounded-md hover:bg-gray-50"
          >
            取消
          </button>
          <button
            @click="savePermission"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
          >
            保存
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useApolloClient } from '@vue/apollo-composable'
import {
  Application,
  ApplicationField,
  GET_APPLICATIONS,
  GET_APPLICATION_PERMISSIONS,
  CREATE_APPLICATION,
  UPDATE_APPLICATION,
  DELETE_APPLICATION,
  SET_APPLICATION_PERMISSION
} from '../graphql/application'

const { client } = useApolloClient()

// 状态变量
const applications = ref<Application[]>([])
const showCreateDialog = ref(false)
const showEditDialog = ref(false)
const showPermissionDialog = ref(false)
const currentApplication = ref<Application | null>(null)

// 表单数据
const form = ref({
  name: '',
  slug: '',
  description: '',
  fields: [] as ApplicationField[],
  status: 'active'
})

const permissionForm = ref({
  roleId: '',
  permissions: [] as string[]
})

// 加载应用模块列表
const loadApplications = async () => {
  try {
    const { data } = await client.query({
      query: GET_APPLICATIONS
    })
    applications.value = data.applications
  } catch (error) {
    console.error('Failed to load applications:', error)
  }
}

// 创建或更新应用模块
const handleSubmit = async () => {
  try {
    if (showEditDialog.value) {
      await client.mutate({
        mutation: UPDATE_APPLICATION,
        variables: {
          id: currentApplication.value?.id,
          input: form.value
        }
      })
    } else {
      await client.mutate({
        mutation: CREATE_APPLICATION,
        variables: {
          input: form.value
        }
      })
    }
    await loadApplications()
    closeDialog()
  } catch (error) {
    console.error('Failed to save application:', error)
  }
}

// 编辑应用模块
const editApplication = (app: Application) => {
  currentApplication.value = app
  form.value = {
    name: app.name,
    slug: app.slug,
    description: app.description,
    fields: [...app.fields],
    status: app.status
  }
  showEditDialog.value = true
}

// 删除应用模块
const deleteApplication = async (id: string) => {
  if (!confirm('确定要删除此应用模块吗？')) return
  try {
    await client.mutate({
      mutation: DELETE_APPLICATION,
      variables: { id }
    })
    await loadApplications()
  } catch (error) {
    console.error('Failed to delete application:', error)
  }
}

// 显示权限配置对话框
const showPermissionDialog = (app: Application) => {
  currentApplication.value = app
  loadApplicationPermissions(app.id)
  permissionForm.value.roleId = ''
  permissionForm.value.permissions = []
  showPermissionDialog.value = true
}

// 加载应用模块权限
const loadApplicationPermissions = async (applicationId: string) => {
  try {
    const { data } = await client.query({
      query: GET_APPLICATION_PERMISSIONS,
      variables: { applicationId }
    })
    if (data.applicationPermissions.length > 0) {
      const permission = data.applicationPermissions[0]
      permissionForm.value.roleId = permission.roleId
      permissionForm.value.permissions = permission.permissions
    }
  } catch (error) {
    console.error('Failed to load application permissions:', error)
  }
}

// 保存权限配置
const savePermission = async () => {
  try {
    await client.mutate({
      mutation: SET_APPLICATION_PERMISSION,
      variables: {
        applicationId: currentApplication.value?.id,
        roleId: permissionForm.value.roleId,
        permissions: permissionForm.value.permissions
      }
    })
    closePermissionDialog()
  } catch (error) {
    console.error('Failed to save application permission:', error)
  }
}

// 添加字段
const addField = () => {
  form.value.fields.push({
    id: '',
    name: '',
    slug: '',
    type: 'string',
    required: false,
    description: '',
    default: null,
    validation: null
  })
}

// 删除字段
const removeField = (index: number) => {
  form.value.fields.splice(index, 1)
}

// 关闭对话框
const closeDialog = () => {
  showCreateDialog.value = false
  showEditDialog.value = false
  form.value = {
    name: '',
    slug: '',
    description: '',
    fields: [],
    status: 'active'
  }
  currentApplication.value = null
}

// 关闭权限配置对话框
const closePermissionDialog = () => {
  showPermissionDialog.value = false
  permissionForm.value = {
    roleId: '',
    permissions: []
  }
  currentApplication.value = null
}

// 初始化加载
onMounted(() => {
  loadApplications()
})
