<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition
        mode="out-in"
        name="el-fade-in-linear"
      >
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<script setup>

defineOptions({
  name: 'TransferRecord' // 转账记录
})
</script>