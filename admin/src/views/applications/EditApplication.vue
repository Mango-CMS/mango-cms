<script setup lang="ts">
import { ref, watch, onMounted, h } from 'vue'
import { useRoute } from 'vue-router'
import { useQuery, useMutation } from '@vue/apollo-composable'
import { GET_APPLICATION, UPDATE_APPLICATION, UPDATE_APPLICATION_SIGN } from '@/graphql/application'
import type { Application } from '@/graphql/application'
import { NCard, NButton, NDescriptions, NDescriptionsItem, NSpace, NDataTable, NTag, NModal, useMessage, useDialog } from 'naive-ui'
import ApplicationForm from './components/ApplicationForm.vue'

const route = useRoute()
const message = useMessage()
const dialog = useDialog()

const application = ref<Application | null>(null)

// 获取应用详情
const { result, loading, refetch } = useQuery(GET_APPLICATION, {
  id: route.params.id as string
})

// 编辑表单
const showEditModal = ref(false)

onMounted(() => {
  refetch()
})

// 监听查询结果
watch(result, value => {
  application.value = value.application
})

// 字段表格列配置
const fieldColumns = [
  { title: '字段名称', key: 'name' },
  { title: '字段标识', key: 'slug' },
  {
    title: '字段类型', key: 'type',
    render: (row: any) => h(NTag, { type: 'info' }, { default: () => row.type })
  },
  {
    title: '是否必填', key: 'required',
    render: (row: any) => h(NTag, { type: row.required ? 'success' : 'warning' },
      { default: () => row.required ? '是' : '否' }
    )
  },
  { title: '描述', key: 'description' }
]

// 模块表格列配置
const modelColumns = [
  { title: '模块名称', key: 'name' },
  { title: '模块标识', key: 'slug' },
  { title: '描述', key: 'description' }
]

// 处理编辑提交
const { mutate: updateApplication, loading: updating } = useMutation(UPDATE_APPLICATION)

// 处理更新签名
const { mutate: updateSign, loading: updatingSign } = useMutation(UPDATE_APPLICATION_SIGN)
const handleUpdateSign = async () => {
  dialog.warning({
    title: '提示',
    content: '更新签名后，所有使用该应用的接口都需要重新签名，确认更新签名吗？',
    positiveText: '确认',
    negativeText: '取消',
    onPositiveClick: async () => {
      try {
        await updateSign({
          id: route.params.id as string
        })
        message.success('签名更新成功')
        refetch()
      } catch (error) {
        message.error('签名更新失败')
      }
    }
  })

}
const handleEditSubmit = async (data: Partial<Application>) => {
  try {
    await updateApplication({
      id: route.params.id as string,
      ...data
    })
    message.success('更新成功')
    showEditModal.value = false
  } catch (error) {
    message.error('更新失败')
  }
}
</script>
<template>
  <div class="p-4  space-y-4">
    <n-card title="基本信息">
      <template #header-extra>
        <n-space>
          <n-button size="small" type="primary" :loading="updating" @click="showEditModal = true">编辑</n-button>
          <n-button size="small" type="info" :loading="updatingSign" @click="handleUpdateSign">更新签名</n-button>
        </n-space>
      </template>

      <n-descriptions v-if="application">
        <n-descriptions-item label="应用名称">
          {{ application.name }}
        </n-descriptions-item>
        <n-descriptions-item label="应用标识">
          {{ application.slug }}
        </n-descriptions-item>
        <n-descriptions-item label="应用签名">
          {{ application.sign }}
        </n-descriptions-item>
        <n-descriptions-item label="应用状态">
          <n-tag :type="application.status === 'active' ? 'success' : 'warning'">{{ application.status ===
            'active'
            ?
            '启用'
            : '禁用' }}</n-tag>
        </n-descriptions-item>
        <n-descriptions-item label="创建时间">
          {{ new Date(application.createdAt).toLocaleString() }}
        </n-descriptions-item>
        <n-descriptions-item label="更新时间">
          {{ new Date(application.updatedAt).toLocaleString() }}
        </n-descriptions-item>
      </n-descriptions>
    </n-card>

    <n-card title="应用字段">
      <n-data-table :columns="fieldColumns" :data="application?.fields || []" :loading="loading" />
    </n-card>

    <n-card title="应用模块">
      <n-data-table :columns="modelColumns" :data="application?.models || []" :loading="loading" />
    </n-card>

    <!-- 编辑弹窗 -->
    <n-modal v-model:show="showEditModal" preset="card" title="编辑应用" style="width: 800px">
      <application-form v-if="application" :initial-data="application" :loading="updating" @submit="handleEditSubmit"
        @cancel="showEditModal = false" />
    </n-modal>
  </div>
</template>
