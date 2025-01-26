<script setup lang="ts">
import { h } from 'vue'
import { NInput, NSelect, NSwitch, NButton, NIcon, NDataTable } from 'naive-ui'
import { CloseOutline } from '@vicons/ionicons5'

interface Props {
  fields: any[]
  onUpdate: (fields: any[]) => void
}

const props = defineProps<Props>()

const fieldTypes = [
  { label: '文本', value: 'string' },
  { label: '数字', value: 'number' },
  { label: '布尔', value: 'boolean' },
  { label: '日期', value: 'date' },
  { label: '富文本', value: 'rich_text' },
  { label: '图片', value: 'image' },
  { label: '文件', value: 'file' }
]

const columns = [
  {
    title: '字段名称',
    key: 'name',
    width: 150,
    render: (row: any, index: number) => h(NInput, {
      value: row.name,
      onUpdateValue: (v) => handleFieldUpdate(index, 'name', v),
      placeholder: '请输入字段名称'
    })
  },
  {
    title: '字段标识',
    key: 'slug',
    width: 150,
    render: (row: any, index: number) => h(NInput, {
      value: row.slug,
      onUpdateValue: (v) => handleFieldUpdate(index, 'slug', v),
      placeholder: '请输入字段标识'
    })
  },
  {
    title: '字段类型',
    key: 'type',
    width: 150,
    render: (row: any, index: number) => h(NSelect, {
      value: row.type,
      options: fieldTypes,
      onUpdateValue: (v) => handleFieldUpdate(index, 'type', v)
    })
  },
  {
    title: '是否必填',
    key: 'required',
    width: 100,
    render: (row: any, index: number) => h(NSwitch, {
      value: row.required,
      onUpdateValue: (v) => handleFieldUpdate(index, 'required', v)
    })
  },
  {
    title: '字段描述',
    key: 'description',
    width: 200,
    render: (row: any, index: number) => h(NInput, {
      value: row.description,
      type: 'textarea',
      onUpdateValue: (v) => handleFieldUpdate(index, 'description', v),
      placeholder: '请输入字段描述'
    })
  },
  {
    title: '操作',
    key: 'actions',
    width: 100,
    render: (row: any, index: number) => h(NButton, {
      circle: true,
      type: 'error',
      size: 'small',
      onClick: () => handleRemoveField(index)
    }, { icon: () => h(NIcon, null, { default: () => h(CloseOutline) }) })
  }
]

const handleFieldUpdate = (index: number, key: string, value: any) => {
  const updatedFields = [...props.fields]
  updatedFields[index][key] = value
  props.onUpdate(updatedFields)
}

const handleRemoveField = (index: number) => {
  const updatedFields = [...props.fields]
  updatedFields.splice(index, 1)
  props.onUpdate(updatedFields)
}

defineExpose({
  columns
})
</script>

<template>
  <n-data-table :columns="columns" :data="fields" :pagination="false" />
</template>
