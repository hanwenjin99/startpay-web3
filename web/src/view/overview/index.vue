<template>
  <main class="content_container">
    <!-- 总览 -->
    <h1 class="title">
      总览
    </h1>
    <section class="assets">
      <h2>资产估值</h2>
      <h1>≈ $2770.753 <img src="@/assets/icons/arrow-right.svg" alt=""> </h1>
    </section>
    <!-- 统计 -->
    <section class="statistics">
      <h1>统计</h1>
      <div class="statisticsActions">
        <span
          v-for="item in statisticsActions"
          :key="item.key"
          :class="['statisticsButton', item.key === statisticsActionActive ? 'active' : '']"
          @click="changestatisticsAction(item.key)"
        >
          {{ item.name }}
        </span>
      </div>
      <div class="statisticsList">
        <span v-for="item in statisticsList" :key="item.key" class="statisticsItem">
          <h1>{{ item.value }} {{ item.unit }}</h1>
          <h2>{{ item.name }}{{ statisticsActionActive === 'recharge' ? '充值' : '提现' }}金额</h2>
        </span>
      </div>
    </section>

    <!-- 公告 -->
    <section class="announcement">
      <h1>公告</h1>
      <el-card class="announcementList" shadow="never">
        <div
          v-for="(item, index) in announcementList"
          :key="item.id"
          class="announcementItem"
          @click="openDetail(index)"
        >
          <h2>{{ item.title }}</h2>
          <h3>{{ dayjs(item.time).format('YYYY-MM-DD HH:mm:ss') }}</h3>
        </div>
      </el-card>
    </section>

    <!-- 公告详情 -->
    <el-dialog
      v-model="dialogVisible"
      width="700"
    >
      <section class="announcementContainer">
        <h1>{{ announcementList[announcementIndex].title }}</h1>
        <h2>{{ announcementList[announcementIndex].time }}</h2>
        <h3>{{ announcementList[announcementIndex].content }}</h3>
      </section>
    </el-dialog>
  </main>
</template>

<script setup>
import  { ref, reactive } from 'vue'
import dayjs from 'dayjs'

const statisticsActions = [{ key: 'recharge', name: '充值' }, { key: 'withdraw', name: '提现' }]

const statisticsActionActive = ref('recharge')

const changestatisticsAction = (key) => {
  statisticsActionActive.value = key
}

// 统计列表
const statisticsList = reactive([
  { key: 'today', name: '今日', value: 0, unit: 'USDT' },
  { key: 'yesterday', name: '昨日', value: 0, unit: 'USDT' },
  { key: 'last7', name: '过去7天', value: 0, unit: 'USDT' },
  { key: 'last30', name: '过去30天', value: 2771.67, unit: 'USDT' },
])

// 公告列表
const announcementList = reactive([
  { id: 1, title: 'FuturesPay 品牌正式升级成为 Satogate', time: '2024-02-23T09:54:20.869+00:00', content: 'x惺惺相惜阿列克' },
  { id: 2, title: 'FuturesPay新增对Polygon公链的支持，继续扩大数字货币服务范围！', time: '2023-06-19T07:31:18.665+00:00', content: 'x惺惺相惜阿列克' },
  { id: 3, title: '全新Futures Pay控制台正式上线！向全球商户提供卓越的数字货币收付款服务', time: '2023-05-24T07:19:13.951+00:00', content: 'x惺惺相惜阿列克' }
])

const dialogVisible = ref(false)
const announcementIndex = ref(0)

const openDetail = (index) => {
  announcementIndex.value = index
  dialogVisible.value = true
}
</script>

<style lang="scss" scoped>
.content_container {
  box-sizing: border-box;
  display: flex;
  flex-direction: column;

  padding: 10px;
  color: #000;

  /* 抗锯齿渲染文字 */
  -webkit-font-smoothing: antialiased;

  .title {
    font-size: 30px;
    margin: 0;
  }

  .assets {
    display: flex;
    flex-direction: column;
    
    margin-top: 32px;
    padding-bottom: 40px;
    border-bottom: 1px solid #e6e6e6;

    h1 {
      display: flex;
      align-items: center;

      margin-bottom: 0;
      margin-top: 8px;

      font-size: 48px;
      cursor: pointer;

      img {
        width: 24px;
        height: 24px;
        margin-left: 12px;
      }
    }

    h2 {
      font-weight: 400;
      margin: 0;
      color: grey;
      font-size: 16px;
    }
  }

  .statistics {
    margin-top: 40px;

    h1 {
      font-size: 24px;
    }

    .statisticsActions {
      display: flex;
      margin-top: 24px;

      .statisticsButton {
        display: flex;
        align-items: center;
        justify-content: center;

        box-sizing: border-box;
        height: 36px;
        margin-right: 8px;
        padding: 0 16px;
        font-size: 14px;
        border-radius: 36px;
        border: 1px solid rgb(230, 230, 230);
        cursor: pointer;
      }

      .active {
        background-color: rgb(230, 230, 230);
        border: 2px solid rgb(0, 0, 0);
      }
    }

    .statisticsList {
      display: flex;
      margin-top: 24px;

      .statisticsItem {
        display: flex;
        flex: 1 1;
        flex-direction: column;
        margin-right: 24px;

        h1 {
          font-size: 24px;
          letter-spacing: -.4px;
          margin: 0;
          line-height: 1.25;
        }

        h2 {
          font-weight: 400;
          margin: 8px 0 0 0;
          color: grey;
          font-size: 16px;
        }
      }
    }
  }

  .announcement {
    margin-top: 48px;

    h1 {
      margin: 0;
      font-size: 24px;
    }
    
    .announcementList {
      background-color: inherit;
      margin-top: 24px;
      line-height: 1.25;

      :deep(.el-card__body) {
        padding-top: 0;
        padding-bottom: 0;
      }
      
      .announcementItem {
        display: flex;
        flex-direction: column;
        padding: 24px 0;

        border-bottom: 1px solid #e6e6e6;
        cursor: pointer;

        h2 {
          font-weight: 400;
          font-size: 16px;
          margin: 0;
        }

        h3 {
          font-weight: 400;
          color: grey;
          font-size: 12px;
          margin-top: 8px;
          margin-bottom: 0;
        }

        &:last-child {
          border-bottom: none;
        }
      }
    }
  }

  .announcementContainer {
    display: flex;
    flex-direction: column;

    h1 {
      margin: 0;
      font-size: 24px;
    }

    h2 {
      margin-top: 16px;
      margin-bottom: 0;
      color: grey;
      font-size: 14px;
      font-weight: 400;
    }

    h3 {
      margin-top: 16px;
      margin-bottom: 0;
      font-size: 16px;
      line-height: 1.5;
      white-space: pre-wrap;
      font-weight: 400;
    }
  }
}
</style>