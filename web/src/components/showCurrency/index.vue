<template>
  <section class="showCurrency">
    <div class="icon">
      <img class="mainIcon" :src="mainIcon" alt="">
      <img class="secondaryIcon" :src="secondaryIcon" alt="">
    </div>
    <span class="name">{{ mainName }}-{{ secondaryName }}-{{ walletName }}</span>
  </section>
</template>

<script setup>
// eslint-disable-next-line no-unused-vars
const props = defineProps({
  mainIcon: {
    type: String,
    default: ''
  },
  secondaryIcon: {
    type: String,
    default: ''
  },
  mainName: {
    type: String,
    default: ''
  },
  secondaryName: {
    type: String,
    default: ''
  },
  walletName: {
    type: String,
    default: ''
  }
})

</script>

<style lang="scss" scoped>
.showCurrency {
  display: flex;
  align-items: center;

  .icon {
    position: relative;
    width: 36px;
    height: 36px;

    .mainIcon {
      width: 36px;
      height: 36px;
      border-radius: 36px;
    }

    .secondaryIcon {
      position: absolute;
      right: 0;
      bottom: 0;
      width: 14px;
      height: 14px;
      border-radius: 14px;
    }
  }

  .name {
    font-weight: 700;
    font-size: 16px;
    margin-left: 12px;
  }
}
</style>