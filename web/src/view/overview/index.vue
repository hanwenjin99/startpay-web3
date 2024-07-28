<template>
  <main class="content_container">
    <!-- 总览 -->
    <h1 class="title">
      总览
    </h1>
    <section class="assets">
      <h2>资产估值</h2>
      <h1 @click="router.push('/layout/account/property')">≈ ${{ Number(dashboardData.assetValuationAmount ?? 0).toFixed(2) }} <img src="@/assets/icons/arrow-right.svg" alt=""> </h1>
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
          <h1>{{ dashboardData[`total${statisticsActionActive}Amount${item.value}`] ?? 0 }} {{ dashboardData.assetValuationCurrency ?? 'USDT' }}</h1>
          <h2>{{ item.name }}{{ statisticsActionActive === 'Deposit' ? '充值' : '提现' }}金额</h2>
        </span>
      </div>
    </section>

    <!-- 公告 -->
    <section v-if="announcementList.length > 0" class="announcement">
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
import  { ref, reactive, onMounted } from 'vue'
import dayjs from 'dayjs'
import { useRouter } from 'vue-router'

import { getDashboard, getAnnouncementList } from '@/api/overview'

const router = useRouter()

const statisticsActions = [{ key: 'Deposit', name: '充值' }, { key: 'Withdraw', name: '提现' }]

const statisticsActionActive = ref('Deposit')

const changestatisticsAction = (key) => {
  statisticsActionActive.value = key
}

const dashboardData = ref({})

// 统计列表
const statisticsList = reactive([
  { key: 'today', name: '今日', value: 'Today', unit: 'USDT' },
  { key: 'yesterday', name: '昨日', value: 'Yesterday', unit: 'USDT' },
  { key: 'last7days', name: '过去7天', value: '7Days', unit: 'USDT' },
  { key: 'last30days', name: '过去30天', value: '30Days', unit: 'USDT' },
])

// 公告列表
const announcementList = ref([])

const dialogVisible = ref(false)
const announcementIndex = ref(0)

const openDetail = (index) => {
  announcementIndex.value = index
  dialogVisible.value = true
}

const initDashboard = async () => {
  const { code, data } = await getDashboard()
  if (code === 0) {
    dashboardData.value = data
  }
}

const initAnnouncementList = async () => {
  const { code, data } = await getAnnouncementList()
  if (code === 0) {
    announcementList.value = data?.announcementList || []
  }
}

onMounted(() => {
  // 初始化获取估值统计数据
  initDashboard()
  // 初始化获取公告列表
  initAnnouncementList()
})
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