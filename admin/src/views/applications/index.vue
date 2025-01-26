<template>
  <div class="p-4 bg-white">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-2xl font-bold">应用管理</h2>
      <n-button type="primary" @click="$router.push({ name: 'application-add' })">
        创建应用
      </n-button>
    </div>

    <mango-table :columns="columns" :data="result?.applications" :loading="queryLoading" @refresh="handleRefresh" />
  </div>
</template>

<script setup lang="ts">
import { h, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { NTag, NButton, useDialog } from 'naive-ui'
import { useQuery, useMutation } from '@vue/apollo-composable'
import { GET_APPLICATIONS, DELETE_APPLICATION } from '@/graphql/application'
import type { Application } from '@/graphql/application'
import MangoTable from '@/components/MangoTable'

const message = useMessage()
const dialog = useDialog()
const router = useRouter()

// 表格列配置
const columns: DataTableColumns<Application> = [
  { title: '名称', key: 'name' },
  { title: '标识', key: 'slug' },
  { title: '描述', key: 'description' },
  {
    title: '状态',
    key: 'status',
    render: (row) => {
      const statusMap = {
        active: '运行',
        inactive: '关闭'
      }
      return h(
        NTag,
        {
          type: row.status === 'active' ? 'success' : 'error',
          bordered: false,
          size: 'small',
        },
        { default: () => statusMap[row.status as keyof typeof statusMap] || '未知' }
      )
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    render: (row) => new Date(row.createdAt).toLocaleString()
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    render: (row) => new Date(row.updatedAt).toLocaleString()
  },
  {
    title: '操作',
    key: 'actions',
    align: 'left',
    render: (row) =>
      h(
        'div',
        { class: 'flex gap-2' },
        [
          h(
            NButton,
            {
              type: 'primary',
              size: 'small',
              onClick: () => {
                router.push({
                  name: 'application-edit',
                  params: {
                    id: row.id
                  }
                })
              }
            },
            { default: () => '编辑' }
          ),
          h(
            NButton,
            {
              type: 'error',
              size: 'small',
              onClick: () => handleDelete(row.id)
            },
            { default: () => '删除' }
          )
        ]
      )
  }
]

// GraphQL 查询和变更
const { result, loading: queryLoading, refetch } = useQuery(GET_APPLICATIONS)
const { mutate: deleteApplication } = useMutation(DELETE_APPLICATION)

// 组件挂载后自动查询数据
onMounted(() => {
  refetch()
})

// 刷新数据
const handleRefresh = () => {
  refetch()
}

// 删除应用
const handleDelete = async (id: string) => {
  dialog.warning({
    title: '确认删除',
    content: '确认删除该应用吗？',
    positiveText: '确定',
    negativeText: '不确定',
    draggable: true,
    onPositiveClick: async () => {
      try {
        await deleteApplication({ id })
        message.success('删除成功')
        handleRefresh()
      } catch (error) {
        message.error('删除失败')
      }
    }
  })
}
</script>
