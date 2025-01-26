<script setup lang="ts">
import { h, ref, computed } from 'vue'
import { NDataTable, NButton, NSpace, NIcon, NPopconfirm, NModal, useMessage, NTag } from 'naive-ui'
import { TrashOutline, PencilOutline } from '@vicons/ionicons5'
import { useQuery, useMutation } from '@vue/apollo-composable'
import { GET_USERS, DELETE_USER, CREATE_USER, UPDATE_USER } from '@/graphql/user'
import type { User } from '@/graphql/user'
import UserForm from './components/UserForm.vue'


const message = useMessage()
const showModal = ref(false)
const modalTitle = ref('')
const selectedUser = ref<User | null>(null)
const formLoading = ref(false)

// GraphQL 查询和变更
const { result, loading, refetch } = useQuery(GET_USERS)
const { mutate: createUser } = useMutation(CREATE_USER)
const { mutate: updateUser } = useMutation(UPDATE_USER)
const { mutate: deleteUser } = useMutation(DELETE_USER)

// 用户列表数据
const users = computed(() => result.value?.users || [])

const columns = [
  { title: 'ID', key: 'id' },
  { title: '用户名', key: 'username' },
  { title: '邮箱', key: 'email' },
  { title: '角色', key: 'role' },
  {
    title: '状态', key: 'status',
    render: (row: User) => {
      return h(NTag, {
        type: row.status === 'active' ? 'success' : 'error'
      }, { default: () => row.status === 'active' ? '正常' : '禁用' })
    }
  },
  {
    title: '创建时间', key: 'createdAt',
    render: (row: User) => {
      return new Date(row.createdAt).toLocaleString()
    },
  },
  {
    title: '操作',
    key: 'actions',
    render(row: User) {
      return h(NSpace, null, {
        default: () => [
          h(NButton,
            {
              size: 'small',
              type: 'primary',
              onClick: () => handleEdit(row)
            },
            { default: () => h(NIcon, null, { default: () => h(PencilOutline) }) }
          ),
          h(NPopconfirm,
            {
              onPositiveClick: () => handleDelete(row.id)
            },
            {
              default: () => '确认删除该用户？',
              trigger: () => h(NButton,
                {
                  size: 'small',
                  type: 'error'
                },
                { default: () => h(NIcon, null, { default: () => h(TrashOutline) }) }
              )
            }
          )
        ]
      })
    }
  }
]

// 打开新增用户模态框
const handleAdd = () => {
  modalTitle.value = '新增用户'
  selectedUser.value = null
  showModal.value = true
}

// 打开编辑用户模态框
const handleEdit = (user: User) => {
  modalTitle.value = '编辑用户'
  selectedUser.value = user
  showModal.value = true
}

// 处理用户表单提交
const handleFormSubmit = async (data: Partial<User>) => {
  formLoading.value = true
  try {
    if (selectedUser.value) {
      // 更新用户
      await updateUser({
        variables: {
          id: selectedUser.value.id,
          input: data
        }
      })
      message.success('更新用户成功')
    } else {
      // 创建用户
      await createUser({
        variables: {
          input: data
        }
      })
      message.success('创建用户成功')
    }
    showModal.value = false
    refetch()
  } catch (error) {
    message.error(error instanceof Error ? error.message : '操作失败')
  } finally {
    formLoading.value = false
  }
}

// 处理删除用户
const handleDelete = async (id: number) => {
  try {
    await deleteUser({
      variables: { id }
    })
    message.success('删除用户成功')
    refetch()
  } catch (error) {
    message.error(error instanceof Error ? error.message : '删除失败')
  }
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold">用户管理</h1>
      <n-button type="primary" @click="handleAdd">新增用户</n-button>
    </div>

    <n-data-table :loading="loading" :columns="columns" :data="users" :pagination="{
        pageSize: 10
      }" />

    <n-modal v-model:show="showModal" :title="modalTitle" preset="dialog" :style="{ width: '600px' }">
      <user-form :loading="formLoading" :initial-data="selectedUser" @submit="handleFormSubmit"
        @cancel="showModal = false" />
    </n-modal>
  </div>
</template>
