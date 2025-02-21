<script setup lang="ts">
  import { ref, watch } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useI18n } from "vue-i18n"
  import { InfoBox } from 'bkui-vue/lib'
  import { storeToRefs } from 'pinia'
  import { useGlobalStore } from '../../../../../store/global'
  import { createApp } from "../../../../../api";

  const router = useRouter()
  const { t } = useI18n()

  const props = defineProps<{
    show: boolean
  }>()
  const emits = defineEmits(['update:show', 'reload'])

  const { spaceId } = storeToRefs(useGlobalStore())

  const formData = ref({
    name: '',
    config_type: 'file',
    reload_type: 'file',
    reload_file_path: '/bscp_test', // @todo 待确认
    mode: 'normal',
    deploy_type: 'common',
    memo: '', // @todo 包含换行符后接口会报错
  })
  const rules = {
    name: [
      {
        validator: (value: string) => value.length >= 2,
        message: '最小长度2个字符'
      },
      {
        validator: (value: string) => value.length <= 32,
        message: '最大长度32个字符'
      },
      {
        validator: (value: string) => {
          return /^[a-zA-Z0-9][a-zA-Z0-9_\-]*[a-zA-Z0-9]?$/.test(value)
        },
        message: '服务名称由英文、数字、下划线、中划线组成且以英文、数字开头和结尾'
      }
    ],
    memo: [
      {
        validator: (value: string) => {
          if (value.length > 0) {
            return /^[\u4e00-\u9fa5a-zA-Z0-9][\u4e00-\u9fa5a-zA-Z0-9_\-\s]*[\u4e00-\u9fa5a-zA-Z0-9]?$/.test(value)
          }
          return true
        },
        message: '仅允许使用中文、英文、数字、下划线、中划线、空格，且必须以中文、英文、数字开头和结尾'
      }
    ]
  }
  const formRef = ref()
  const pending = ref(false)

  watch(() => props.show, (val) => {
    if (val) {
      formData.value.name = ''
      formData.value.memo = ''
    }
  })

  const handleCreateConfirm = async () => {
    await formRef.value.validate()
    pending.value = false
    try {
      const resp = await createApp(spaceId.value, formData.value)
      InfoBox({
        type: "success",
        title: "服务新建成功",
        subTitle: "接下来你可以在服务下新增并使用配置项",
        headerAlign: "center",
        footerAlign: "center",
        confirmText: "新增配置项",
        cancelText: "稍后再说",
        onConfirm() {
          router.push({
            name: 'service-config',
            params: {
              spaceId: spaceId.value,
              appId: resp.id
            }
          })
        },
        onClosed() {
          emits('reload')
        }
      } as any);
      handleClose()
    } catch (e) {
      console.error(e)
    } finally {
      pending.value = false
    }
  };

  const handleClose = () => {
    emits('update:show', false)
  }
</script>
<template>
  <bk-sideslider
    width="640"
    :is-show="props.show"
    :title="t('新建服务')"
    :before-close="handleClose">
    <div class="create-app-form">
      <bk-form form-type="vertical" ref="formRef" :model="formData" :rules="rules">
        <bk-form-item :label="t('服务名称')" property="name" required>
          <bk-input
            placeholder="请输入2~32字符，只允许英文、数字、下划线、中划线且必须以英文、数字开头和结尾"
            v-model="formData.name"
          ></bk-input>
        </bk-form-item>
        <bk-form-item :label="t('服务描述')" property="memo">
          <bk-input
            placeholder="请输入"
            type="textarea"
            v-model="formData.memo" />
        </bk-form-item>
      </bk-form>
    </div>
    <template #footer>
      <div class="create-app-footer">
        <bk-button
          theme="primary"
          :loading="pending"
          @click="handleCreateConfirm">
          {{ t("提交") }}
        </bk-button>
        <bk-button @click="handleClose">{{ t("取消") }}</bk-button>
      </div>
    </template>
  </bk-sideslider>
</template>
<style lang="scss" scoped>
.create-app-form {
  padding: 20px 24px;
  height: calc(100vh - 108px);
}
.create-app-footer {
  padding: 8px 24px;
  height: 48px;
  width: 100%;
  background: #fafbfd;
  box-shadow: 0 -1px 0 0 #dcdee5;
  button {
    margin-right: 8px;
    min-width: 88px;
  }
}
</style>
