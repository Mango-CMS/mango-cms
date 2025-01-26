<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { NCard, NForm, NFormItem, NInput, NButton, useMessage } from 'naive-ui'
import { useMutation } from '@vue/apollo-composable'
import { LOGIN } from '@/graphql/auth'

const router = useRouter()
const message = useMessage()
const formData = ref({
  username: '',
  password: ''
})
const formRef = ref()

const rules = {
  username: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur'
  },
  password: {
    required: true,
    message: '请输入密码',
    trigger: 'blur'
  }
}

const { mutate: login, loading } = useMutation(LOGIN)

const handleLogin = async () => {
  formRef.value?.validate((errors: any) => {
    if (errors) {
      message.error(errors)
      return;
    }
  })
  try {
    const { data } = await login({
      username: formData.value.username,
      password: formData.value.password
    })

    if (data?.login?.token) {
      localStorage.setItem('token', data.login.token)
      message.success('登录成功')
      router.push('/')
    } else {
      throw new Error('登录失败：未获取到有效token')
    }
  } catch (error) {
    console.error('登录失败:', error)
    message.error(error instanceof Error ? error.message : '登录失败，请重试')
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="w-96">
      <n-card>
        <div class="text-center mb-8">
          Mangoo CMS
          <h1 class="text-2xl font-bold">登录</h1>
        </div>
        <n-form ref="formRef" :rules="rules">
          <n-form-item label="用户名">
            <n-input v-model:value="formData.username" placeholder="请输入用户名" />
          </n-form-item>
          <n-form-item label="密码">
            <n-input v-model:value="formData.password" type="password" placeholder="请输入密码" @keyup.enter="handleLogin" />
          </n-form-item>
          <n-button type="primary" block :loading="loading" @click="handleLogin">
            登录
          </n-button>
        </n-form>
      </n-card>
    </div>
  </div>
</template>
