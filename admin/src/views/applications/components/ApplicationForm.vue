<template>
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="auto"
    require-mark-placement="right-hanging" size="medium">
    <n-form-item label="应用名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入应用名称" />
    </n-form-item>

    <n-form-item label="应用标识" path="slug">
      <n-input v-model:value="formData.slug" placeholder="请输入应用标识" />
    </n-form-item>

    <n-form-item label="应用描述" path="description">
      <n-input v-model:value="formData.description" type="textarea" placeholder="请输入应用描述" />
    </n-form-item>

    <n-form-item label="应用状态" path="status">
      <n-select v-model:value="formData.status" :options="[
    { label: '启用', value: 'active' },
    { label: '禁用', value: 'inactive' }
  ]" />
    </n-form-item>

    <n-divider>字段配置</n-divider>

    <div v-for="(field, index) in formData.fields" :key="index" class="mb-4">
      <n-card>
        <template #header>
          <div class="flex justify-between items-center">
            <span>字段 #{{ index + 1 }}</span>
            <n-button circle type="error" size="small" @click="removeField(index)">
              <template #icon>
                <n-icon><trash-icon /></n-icon>
              </template>
            </n-button>
          </div>
        </template>

        <n-form-item :path="`fields[${index}].name`" label="字段名称">
          <n-input v-model:value="field.name" placeholder="请输入字段名称" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].slug`" label="字段标识">
          <n-input v-model:value="field.slug" placeholder="请输入字段标识" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].type`" label="字段类型">
          <n-select v-model:value="field.type" :options="[
    { label: '文本', value: 'string' },
    { label: '数字', value: 'number' },
    { label: '布尔', value: 'boolean' },
    { label: '日期', value: 'date' },
    { label: '富文本', value: 'rich_text' },
    { label: '图片', value: 'image' },
    { label: '文件', value: 'file' }
  ]" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].required`" label="是否必填">
          <n-switch v-model:value="field.required" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].description`" label="字段描述">
          <n-input v-model:value="field.description" type="textarea" placeholder="请输入字段描述" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].default`" label="默认值">
          <n-input v-model:value="field.default" placeholder="请输入默认值" />
        </n-form-item>

        <n-form-item :path="`fields[${index}].validation`" label="验证规则">
          <n-input v-model:value="field.validation" type="textarea" placeholder="请输入验证规则" />
        </n-form-item>
      </n-card>
    </div>

    <n-button class="mb-4" block dashed type="primary" @click="addField">
      添加字段
    </n-button>

    <div class="flex justify-end gap-2">
      <n-button @click="$emit('cancel')">取消</n-button>
      <n-button type="primary" :loading="loading" @click="handleSubmit">
        确定
      </n-button>
    </div>
  </n-form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { NButton, NForm, NFormItem, NSwitch, NInput, NSelect } from 'naive-ui'
import type { FormInst } from 'naive-ui'
import type { Application, ApplicationField } from '@/graphql/application'

const props = defineProps<{
  loading?: boolean
  initialData?: Application
}>()

const emit = defineEmits<{
  submit: [data: Partial<Application>]
  cancel: []
}>()

const formRef = ref<FormInst | null>(null)

// 表单数据
const formData = ref<Partial<Application>>({
  name: '',
  slug: '',
  description: '',
  status: 'active',
  fields: []
})

// 如果有初始数据，则填充表单
if (props.initialData) {
  formData.value = { ...props.initialData }
}

// 表单验证规则
const rules = {
  name: {
    required: true,
    message: '请输入应用名称',
    trigger: 'blur'
  },
  slug: {
    required: true,
    message: '请输入应用标识',
    trigger: 'blur'
  },
  'fields.*.name': {
    required: true,
    message: '请输入字段名称',
    trigger: 'blur'
  },
  'fields.*.slug': {
    required: true,
    message: '请输入字段标识',
    trigger: 'blur'
  },
  'fields.*.type': {
    required: true,
    message: '请选择字段类型',
    trigger: 'blur'
  }
}

// 添加字段
const addField = () => {
  if (!formData.value.fields) {
    formData.value.fields = []
  }
  formData.value.fields.push({
    name: '',
    slug: '',
    type: 'string',
    required: false,
    description: '',
    default: '',
    validation: ''
  })
}

// 移除字段
const removeField = (index: number) => {
  formData.value.fields?.splice(index, 1)
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
