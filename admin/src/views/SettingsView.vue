<script setup lang="ts">
import { ref } from 'vue'
import { NCard, NTabs, NTabPane, NForm, NFormItem, NInput, NButton, NSpace, NSwitch } from 'naive-ui'

const activeTab = ref('system')
const loading = ref(false)

const systemSettings = ref({
  siteName: 'Mango CMS',
  siteDescription: '一个现代的内容管理系统',
  allowRegistration: true,
  maintenanceMode: false
})

const personalSettings = ref({
  nickname: '',
  email: '',
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const handleSaveSystemSettings = async () => {
  loading.value = true
  try {
    // TODO: 实现保存系统设置的API调用
    console.log('保存系统设置:', systemSettings.value)
  } catch (error) {
    console.error('保存系统设置失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSavePersonalSettings = async () => {
  loading.value = true
  try {
    // TODO: 实现保存个人设置的API调用
    console.log('保存个人设置:', personalSettings.value)
  } catch (error) {
    console.error('保存个人设置失败:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold">系统设置</h1>

    <n-tabs v-model:value="activeTab" type="segment">
      <n-tab-pane name="system" tab="系统设置">
        <n-card>
          <n-form>
            <n-form-item label="站点名称">
              <n-input v-model:value="systemSettings.siteName" />
            </n-form-item>
            <n-form-item label="站点描述">
              <n-input
                v-model:value="systemSettings.siteDescription"
                type="textarea"
                :rows="3"
              />
            </n-form-item>
            <n-form-item label="允许注册">
              <n-switch v-model:value="systemSettings.allowRegistration" />
            </n-form-item>
            <n-form-item label="维护模式">
              <n-switch v-model:value="systemSettings.maintenanceMode" />
            </n-form-item>
            <n-space justify="end">
              <n-button
                type="primary"
                :loading="loading"
                @click="handleSaveSystemSettings"
              >
                保存设置
              </n-button>
            </n-space>
          </n-form>
        </n-card>
      </n-tab-pane>

      <n-tab-pane name="personal" tab="个人设置">
        <n-card>
          <n-form>
            <n-form-item label="昵称">
              <n-input v-model:value="personalSettings.nickname" />
            </n-form-item>
            <n-form-item label="邮箱">
              <n-input v-model:value="personalSettings.email" />
            </n-form-item>
            <n-form-item label="当前密码">
              <n-input
                v-model:value="personalSettings.currentPassword"
                type="password"
              />
            </n-form-item>
            <n-form-item label="新密码">
              <n-input
                v-model:value="personalSettings.newPassword"
                type="password"
              />
            </n-form-item>
            <n-form-item label="确认新密码">
              <n-input
                v-model:value="personalSettings.confirmPassword"
                type="password"
              />
            </n-form-item>
            <n-space justify="end">
              <n-button
                type="primary"
                :loading="loading"
                @click="handleSavePersonalSettings"
              >
                保存设置
              </n-button>
            </n-space>
          </n-form>
        </n-card>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>
