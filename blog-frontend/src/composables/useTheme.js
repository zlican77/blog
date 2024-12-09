import { ref, watch } from 'vue'
import { usePreferredDark } from '@vueuse/core'

export function useTheme() {
  const preferredDark = usePreferredDark()
  const isDark = ref(preferredDark.value)

  const toggleTheme = () => {
    isDark.value = !isDark.value
  }

  watch(isDark, (val) => {
    document.documentElement.classList.toggle('dark', val)
  }, { immediate: true })

  return {
    isDark,
    toggleTheme
  }
} 