<script setup lang="ts">
import { ref } from 'vue';
import { NStep, NSteps, NButton, NForm, NGrid, NFormItemGi, NInput, NSelect, NDivider } from "naive-ui";
import type { FormInst } from "naive-ui";
import type { Application } from '@/graphql/application';
import { useMutation } from '@vue/apollo-composable';
import { CREATE_APPLICATION } from '@/graphql/application';
import { useRouter } from 'vue-router';
import { useMessage } from 'naive-ui';
import FieldsTable from './components/FieldsTable.vue';

const router = useRouter();
const message = useMessage();
const formRef = ref<FormInst | null>(null);
const current = ref(1);

const steps = [
  { title: '应用基本信息', description: '应用的基本信息' },
  { title: '应用模块', description: '设置应用的字段' },
  { title: '应用权限', description: '设置应用的权限' },
];

const formData = ref<Partial<Application>>({
  name: '',
  slug: '',
  description: '',
  status: 'active',
  fields: [],
  permissions: []
});

const rules = {
  name: { required: true, message: '请输入应用名称', trigger: 'blur' },
  slug: { required: true, message: '请输入应用标识', trigger: 'blur' },
  status: { required: true, message: '请选择应用状态', trigger: 'blur' },
  'fields.*.name': { required: true, message: '请输入字段名称', trigger: 'blur' },
  'fields.*.slug': { required: true, message: '请输入字段标识', trigger: 'blur' },
  'fields.*.type': { required: true, message: '请选择字段类型', trigger: 'blur' }
};

const addField = () => {
  formData.value.fields?.push({
    name: '',
    slug: '',
    type: 'string',
    required: false,
    description: '',
    default: '',
    validation: ''
  });
};

const handleFieldsUpdate = (fields: any[]) => {
  formData.value.fields = [...fields];
};

const handleNextStep = (nextStep: number) => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      current.value = nextStep;
    } else {
      message.error('请填写完必填项后再进行下一步');
    }
  }, rule => {
    if (current.value === 1) {
      return ['name', 'slug', 'status'].includes(rule.key as string);
    } else if (current.value === 2) {
      return rule.key?.startsWith('fields.') && ['name', 'slug', 'type'].some(field => rule.key?.endsWith(field));
    }
    return true;
  });
};

const { mutate: createApplication } = useMutation(CREATE_APPLICATION);

const handleSubmit = async () => {
  try {
    await formRef.value?.validate();
    await createApplication({
      ...formData.value
    });
    message.success('应用创建成功');
    router.push({ name: 'applications' });
  } catch (error) {
    message.error('应用创建失败');
    console.error(error);
  }
};
</script>

<template>
  <div class="p-4 bg-white space-y-6">
    <h2 class="text-2xl font-bold">创建应用</h2>

    <n-steps :current="current">
      <n-step v-for="(item, i) in steps" :key="i" :title="item.title" :description="item.description" />
    </n-steps>

    <n-form ref="formRef" :model="formData" :rules="rules" label-width="auto" require-mark-placement="right-hanging"
      size="medium">
      <!-- 基本信息 -->
      <div v-if="current === 1" class="space-y-4">
        <n-grid :cols="12" :x-gap="24">
          <n-form-item-gi :span="12" label="应用名称" path="name">
            <n-input v-model:value="formData.name" placeholder="请输入应用名称" />
          </n-form-item-gi>
          <n-form-item-gi :span="12" label="应用标识" path="slug">
            <n-input v-model:value="formData.slug" placeholder="请输入应用标识" />
          </n-form-item-gi>
          <n-form-item-gi :span="12" label="应用描述" path="description">
            <n-input v-model:value="formData.description" type="textarea" placeholder="请输入应用描述" />
          </n-form-item-gi>
          <n-form-item-gi :span="12" label="应用状态" path="status">
            <n-select v-model:value="formData.status" :options="[
      { label: '启用', value: 'active' },
      { label: '禁用', value: 'inactive' }
    ]" />
          </n-form-item-gi>
        </n-grid>
      </div>

      <!-- 字段配置 -->
      <div v-if="current === 2" class="space-y-4">
        <n-divider>字段配置</n-divider>
        <fields-table :fields="formData.fields || []" :on-update="handleFieldsUpdate" />
        <n-button block dashed type="primary" @click="addField">添加字段</n-button>
      </div>

      <!-- 权限配置 -->
      <div v-if="current === 3" class="space-y-4">
        <n-divider>权限配置</n-divider>
        <n-grid :cols="12" :x-gap="24">
          <n-form-item-gi :span="12" label="基础权限">
            <n-select v-model:value="formData.permissions" multiple :options="[
      { label: '查看列表', value: 'list' },
      { label: '查看详情', value: 'view' },
      { label: '创建', value: 'create' },
      { label: '编辑', value: 'update' },
      { label: '删除', value: 'delete' }
    ]" />
          </n-form-item-gi>
        </n-grid>
      </div>

      <!-- 步骤操作按钮 -->
      <div class="flex justify-end gap-2 mt-4">
        <n-button v-if="current <= 1" @click="$router.back()">取消</n-button>
        <n-button v-if="current > 1" @click="current--">上一步</n-button>
        <n-button v-if="current < 3" type="primary" @click="handleNextStep(current + 1)">下一步</n-button>
        <n-button v-else type="primary" @click="handleSubmit">提交</n-button>
      </div>
    </n-form>
  </div>
</template>
