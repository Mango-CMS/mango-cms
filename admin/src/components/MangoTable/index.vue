<template>
  <div class="mango-table">
    <div class="mb-4 flex justify-end gap-x-2">
      <n-tooltip trigger="hover">
        <template #trigger>
          <n-button circle secondary @click="handleRefresh">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
          </n-button>
        </template>
        刷新
      </n-tooltip>

      <n-tooltip trigger="hover">
        <template #trigger>
          <n-button circle secondary @click="handleExport">
            <template #icon>
              <n-icon><download-outline /></n-icon>
            </template>
          </n-button>
        </template>
        导出
      </n-tooltip>
    </div>

    <n-data-table ref="tableRef" :columns="props.columns" :data="props.data" :loading="props.loading"
      :pagination="localPagination" :row-class-name="props.rowClassName" @update:page="handlePageChange"
      @update:page-size="handlePageSizeChange" @row-click="props.rowClick" @row-dblclick="props.rowDblClick"
      @row-contextmenu="props.rowContextmenu" @cell-click="props.cellClick" @cell-dblclick="props.cellDblClick" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { NDataTable, NButton, NTooltip, NIcon } from 'naive-ui'
import { RefreshOutline, DownloadOutline } from '@vicons/ionicons5'

const props = withDefaults(defineProps<{
  columns: any[]
  data: any[]
  loading?: boolean
  rowKey?: string | ((rowData: any) => string)
  pagination?: any
  rowClassName?: string | ((rowData: any) => string)
  rowClick?: (rowData: any) => void
  rowDblClick?: (rowData: any) => void
  rowContextmenu?: (rowData: any) => void
  cellClick?: (rowData: any) => void
  cellDblClick?: (rowData: any) => void
}>(), {
  loading: false,
  rowKey: 'id'
})

const tableRef = ref<InstanceType<typeof NDataTable>>()

const emit = defineEmits<{
  'update:page': [page: number]
  'update:pageSize': [pageSize: number]
  'refresh': []
  'export': []
}>()

// 本地分页配置
const localPagination = computed(() => ({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 30, 40],
  ...props.pagination
}))

// 分页处理
const handlePageChange = (page: number) => {
  emit('update:page', page)
}

const handlePageSizeChange = (pageSize: number) => {
  emit('update:pageSize', pageSize)
}

// 刷新处理
const handleRefresh = () => {
  emit('refresh')
}

// 导出处理
const handleExport = () => {
  tableRef.value?.downloadCsv();
}

// 暴露方法
defineExpose({
  refresh: handleRefresh,
  export: handleExport
})
</script>

<style scoped>
.mango-table {
  width: 100%;
}
</style>
