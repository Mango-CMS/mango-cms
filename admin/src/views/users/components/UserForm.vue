<template>
  <n-form ref="formRef" :model="formData" :rules="rules">
    <n-form-item label="用户名" path="username">
      <n-input v-model:value="formData.username" placeholder="请输入用户名" />
    </n-form-item>

    <n-form-item label="邮箱" path="email">
      <n-input v-model:value="formData.email" placeholder="请输入邮箱" />
    </n-form-item>

    <n-form-item label="密码" path="password" v-if="!initialData">
      <n-input v-model:value="formData.password" type="password" placeholder="请输入密码" />
    </n-form-item>

    <n-form-item label="角色" path="role">
      <n-select v-model:value="formData.role" :options="roleOptions" placeholder="请选择角色" />
    </n-form-item>

    <n-form-item label="状态" path="status">
      <n-select v-model:value="formData.status" :options="statusOptions" placeholder="请选择状态" />
    </n-form-item>

    <div class="flex justify-end gap-2">
      <n-button @click="$emit('cancel')">取消</n-button>
      <n-button type="primary" :loading="loading" @click="handleSubmit">确定</n-button>
    </div>
  </n-form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { NForm, NFormItem, NInput, NSelect, NButton } from 'naive-ui'
import type { FormInst } from 'naive-ui'
import type { User } from '@/graphql/user'

const props = defineProps<{
  loading?: boolean
  initialData?: User
}>()

const emit = defineEmits<{
  submit: [data: Partial<User>]
  cancel: []
}>()

const formRef = ref<FormInst | null>(null)

// 表单数据
const formData = ref<Partial<User>>({
  username: '',
  email: '',
  password: '',
  role: 'editor',
  status: 'active'
})

// 如果有初始数据，则填充表单
if (props.initialData) {
  const { password, ...rest } = props.initialData
  formData.value = { ...rest }
}

// 角色选项
const roleOptions = [
  { label: '管理员', value: 'admin' },
  { label: '编辑', value: 'editor' },
  { label: '普通用户', value: 'user' }
]

// 状态选项
const statusOptions = [
  { label: '正常', value: 'active' },
  { label: '禁用', value: 'inactive' }
]

// 表单验证规则
const rules = {
  username: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur'
  },
  email: [
    {
      required: true,
      message: '请输入邮箱',
      trigger: 'blur'
    },
    {
      type: 'email',
      message: '请输入正确的邮箱格式',
      trigger: ['blur', 'input']
    }
  ],
  password: {
    required: !props.initialData,
    message: '请输入密码',
    trigger: 'blur'
  },
  role: {
    required: true,
    message: '请选择角色',
    trigger: 'blur'
  },
  status: {
    required: true,
    message: '请选择状态',
    trigger: 'blur'
  }
}

// 提交表单
const handleSubmit = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      emit('submit', formData.value)
    }
  })
}
</script>
