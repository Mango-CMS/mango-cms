<script setup lang="ts">
import { h } from 'vue'
import { useRouter, RouterView } from 'vue-router'
import {
  NLayout,
  NLayoutHeader,
  NLayoutContent,
  NLayoutSider,
  NMenu,
  NIcon,
  NSpace,
  NAvatar,
  NDropdown,
  useMessage
} from 'naive-ui'
import { apolloClient } from '@/graphql/client'
import {
  HomeOutline,
  PeopleOutline,
  DocumentTextOutline,
  SettingsOutline,
  LogOutOutline
} from '@vicons/ionicons5'

const router = useRouter()
const message = useMessage()

const menuOptions = [
  {
    label: '仪表盘',
    key: 'dashboard',
    icon: renderIcon(HomeOutline)
  },
  {
    label: '用户管理',
    key: 'users',
    icon: renderIcon(PeopleOutline)
  },
  {
    label: '应用管理',
    key: 'applications',
    icon: renderIcon(DocumentTextOutline)
  },
  {
    label: '系统设置',
    key: 'settings',
    icon: renderIcon(SettingsOutline)
  }
]

const dropdownOptions = [
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogOutOutline)
  }
]

function renderIcon(icon: any) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

function handleMenuClick(key: string) {
  router.push({ name: key })
}

function handleDropdownSelect(key: string) {
  if (key === 'logout') {
    localStorage.removeItem('token')
    apolloClient.clearStore()
    router.push('/login')
    message.success('已成功退出登录')
  }
}
</script>
<template>
  <div class="h-full w-full relative">
    <n-layout class="w-full h-full" position="absolute">
      <n-layout-header class="header" bordered>
        <div class="flex justify-between items-center px-6 h-full">
          <div class="text-xl font-bold">Mango CMS</div>
          <n-space>
            <n-dropdown :options="dropdownOptions" @select="handleDropdownSelect" trigger="click">
              <n-avatar round size="medium" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
            </n-dropdown>
          </n-space>
        </div>
      </n-layout-header>
      <n-layout has-sider class="sider-layout">
        <n-layout-sider collapse-mode="width" :collapsed-width="64" :width="240" show-trigger="arrow-circle" bordered>
          <n-menu :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions" :value="$route.name"
            @update:value="handleMenuClick" />
        </n-layout-sider>

        <n-layout-content embedded content-class="p-4">
          <RouterView />
        </n-layout-content>

      </n-layout>
    </n-layout>
  </div>
</template>
<style scoped>
.header {
  height: 64px;
  padding: 0;
  box-sizing: border-box;
}

.sider-layout {
  height: calc(100vh - 64px);
}
</style>
