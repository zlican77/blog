<template>
    <div class="markdown-editor">
      <MdEditor
        v-model="content"
        @onChange="handleChange"
        :theme="isDark ? 'dark' : 'light'"
        :toolbars="toolbars"
        :preview="preview"
        @onUploadImg="onUploadImg"
        :showCodeRowNumber="true"
        codeTheme="github"
      />
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { MdEditor } from 'md-editor-v3'
  import 'md-editor-v3/lib/style.css'
  import { useTheme } from '@/composables/useTheme'
  
  const props = defineProps({
    modelValue: {
      type: String,
      default: ''
    },
    preview: {
      type: Boolean,
      default: true
    }
  })
  
  const emit = defineEmits(['update:modelValue', 'change'])
  
  const { isDark } = useTheme()
  const content = ref(props.modelValue)
  
  const toolbars = [
    'bold',
    'underline',
    'italic',
    'strikeThrough',
    '-',
    'title',
    'sub',
    'sup',
    'quote',
    'unorderedList',
    'orderedList',
    'task',
    '-',
    'codeRow',
    'code',
    'link',
    'image',
    'table',
    'mermaid',
    '-',
    'revoke',
    'next',
    'save',
    '=',
    'preview',
    'htmlPreview',
    'catalog'
  ]
  
  const handleChange = (value) => {
    emit('update:modelValue', value)
    emit('change', value)
  }
  
  // 图片上传处理
  const onUploadImg = async (files, callback) => {
    try {
      // 这里实现图片上传逻辑
      const urls = await Promise.all(
        files.map(async (file) => {
          // 示例：调用上传API
          // const formData = new FormData()
          // formData.append('image', file)
          // const res = await uploadImage(formData)
          // return res.data.url
          
          // 临时返回blob URL用于演示
          return URL.createObjectURL(file)
        })
      )
      callback(urls)
    } catch (error) {
      console.error('图片上传失败:', error)
    }
  }
  </script>
  
  <style scoped>
  .markdown-editor {
    width: 100%;
    border-radius: 8px;
    overflow: hidden;
  }
  
  :deep(.md-editor) {
    border-radius: 8px;
  }
  </style> 