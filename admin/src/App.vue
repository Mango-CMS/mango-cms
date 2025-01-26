<script setup lang="ts">
import { RouterView } from 'vue-router'
import { NMessageProvider, NConfigProvider, NModalProvider, NNotificationProvider, NDialogProvider, dateZhCN, zhCN } from 'naive-ui'
</script>

<template>
  <n-config-provider abstract :theme="null" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <n-modal-provider>
        <n-dialog-provider>
          <n-notification-provider>
            <RouterView />
          </n-notification-provider>
        </n-dialog-provider>
      </n-modal-provider>
    </n-message-provider>
  </n-config-provider>
</template>
