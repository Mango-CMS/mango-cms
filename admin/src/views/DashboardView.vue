<script setup lang="ts">
import { NCard, NGrid, NGridItem, NStatistic, NIcon } from 'naive-ui'
import { h } from 'vue'
import { PersonOutline, BookOutline, EyeOutline } from '@vicons/ionicons5'

const stats = [
  {
    label: '总用户数',
    value: 1234,
    icon: PersonOutline
  },
  {
    label: '文章数量',
    value: 56,
    icon: BookOutline
  },
  {
    label: '今日访问',
    value: 789,
    icon: EyeOutline
  }
]
</script>

<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold">仪表盘</h1>

    <n-grid :cols="3" :x-gap="16">
      <n-grid-item v-for="stat in stats" :key="stat.label">
        <n-card>
          <n-statistic>
            <template #prefix>
              <n-icon size="24">
                <component :is="stat.icon" />
              </n-icon>
            </template>
            <template #label>{{ stat.label }}</template>
            {{ stat.value }}
          </n-statistic>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>
