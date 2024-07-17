<template>
  <main class="content_container">
    <header class="header">
      <h1 class="title">每月账单</h1>

      <el-date-picker v-model="selectMonth" type="month" placeholder="选择月份" @change="handleChangeDate" />
    </header>
    <div class="month_bill_chart">
      <div class="left_chart">
        <!-- 总费用 -->
        <div class="left_chart_item">
          <img class="item_icon" :src="totalCostIcon" alt="">
          <span class="item_name">总费用</span>
          <span class="item_value item_value_strong">{{ summary.total }} USDT</span>
        </div>
        <!-- = -->
        <div class="left_chart_item">
          <img class="item_icon" :src="sum" alt="">
        </div>
        <!-- 月费 -->
        <div class="left_chart_item">
          <img class="item_icon" :src="monthCost" alt="">
          <span class="item_name">月费</span>
          <span class="item_value_waived">500</span>
          <span class="item_value">{{ summary.monthly }} USDT</span>
        </div>
        <!-- + -->
        <div class="left_chart_item">
          <img class="item_icon" :src="addCost" alt="">
        </div>
        <!-- 燃料费 -->
        <div class="left_chart_item">
          <img class="item_icon" :src="gasCost" alt="">
          <span class="item_name">链上燃料费</span>
          <span class="item_value">{{ summary.gas }} USDT</span>
        </div>
        <!-- + -->
        <div class="left_chart_item">
          <img class="item_icon" :src="addCost" alt="">
        </div>
        <!-- 手续费 -->
        <div class="left_chart_item">
          <img class="item_icon" :src="feeCost" alt="">
          <span class="item_name">兑换手续费</span>
          <span class="item_value">{{ summary.switchFee }} USDT</span>
        </div>

        <span class="left_chart_percent" />

        <!-- 状态 -->
        <div class="left_chart_status">
          <span class="date">账单周期: {{ billCycle }}</span>
          <div class="status">
            <span :class="['signle', summary.billStatus === 'PENDING' ? 'pending' : 'settled' ]" />
            {{ summary.billStatus === 'PENDING' ? '待处理' : '已结算' }}
          </div>
        </div>
      </div>
      <!-- 分割线 -->
      <span class="line" />
      <div class="right_chart">
        <div class="right_chart_desc">
          <div class="screen_container">
            <span class="name">费用趋势</span>
            <el-select v-model="trendSelect" style="width: 150px;" size="large">
              <el-option
                v-for="item in trendOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </div>
          <span class="average">月均 0.03 USDT</span>
        </div>
        <!-- 柱状图 -->
        <v-chart :option="option" autoresize style="width: 100%; height: 200px; margin-top: 16px;" />
      </div>
    </div>

    <!-- 记录 -->
    <section class="record">
      <div class="title">
        <span>费用明细</span>
        <el-button link icon="upload" @click="handleDownload">导出文件</el-button>
      </div>
    </section>

    <!-- 表格 -->
    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间"></el-table-column>
      <el-table-column label="从"></el-table-column>
      <el-table-column label="到"></el-table-column>
      <el-table-column label="Txid"></el-table-column>
      <el-table-column label="交易金额"></el-table-column>
      <el-table-column label="兑换手续费"></el-table-column>
      <el-table-column label="链上燃料费"></el-table-column>
      <el-table-column label="总费用"></el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="footerPage">
      <el-pagination
        background
        layout="prev, pager, next"
        :current-page="currentPage"
        :total="total"
        @current-change="handleChangePage"
      />
    </div>
  </main>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import dayjs from 'dayjs'

// 使用vue-chart
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { BarChart } from 'echarts/charts'
import { TooltipComponent, GridComponent } from 'echarts/components'
import VChart from 'vue-echarts'

use([
  CanvasRenderer,
  BarChart,
  TooltipComponent,
  GridComponent
])

import { getBillSummary, getBillList, exportBillList } from '@/api/cost'

import totalCostIcon from '@/assets/total_cost_icon.png'
import sum from '@/assets/sum.png'
import monthCost from '@/assets/month_cost.png'
import addCost from '@/assets/add_cost.png'
import gasCost from '@/assets/gas_cost.png'
import feeCost from '@/assets/fee_cost.png'

const selectMonth = ref(dayjs(new Date()).format('YYYY-MM'))
const trendSelect = ref('total')

const recordData = ref([])
const total = ref(0)
const currentPage = ref(1)
const billCycle = computed(() => (
  `${dayjs(selectMonth.value).startOf('month').format('MM-DD')} ~ ${dayjs(selectMonth.value).endOf('month').format('MM-DD')}`
))

const trendOptions = ref([
  { label: '总费用', value: 'total'},
  { label: '月费', value: 'monthly' },
  { label: '链上燃料费', value: 'gas' },
  { label: '兑换手续费', value: 'switchFee' }
])

const summary = ref({})

// 柱状图数据
const trends = ref([])

const option = computed(() => ({
  grid: {
    left: 0,
    top: 0,
    right: 0,
    bottom: 1
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    },
  },
  xAxis: {
    data: trends.value.map(item => item.month).reverse()
  },
  yAxis: {
    show: false
  },
  series: [
    {
      name: '费用',
      type: 'bar',
      barWidth: 150,
      label: {
        show: true,
        position: 'top',
        color: '#bb8300'
      },
      itemStyle: {
        color: '#fab30a',
        borderRadius: [8, 8, 0, 0]
      },
      data: trends.value.map(item => Number(item[trendSelect.value]).toFixed(2)).reverse()
    }
  ]
}));

const queryBillSummary = async (month) => {
  const { code, data = {} } = await getBillSummary(month)
  if (code === 0) {
    summary.value = data
    trends.value = data.trends || []
  }
}

const queryBillList = async (page) => {
  const { code, data = {} } = await getBillList({
    page,
    month: dayjs(selectMonth.value).format('YYYY-MM')
  })
  if (code === 0) {
    recordData.value = data.content || []
    total.value = data.total || 0
  }
}

const handleDownload = () => {
  exportBillList(dayjs(selectMonth.value).format('YYYY-MM'))
}

// 分页查询明细
const handleChangePage = (page) => {
  queryBillList(page)
}

// 选择日期
const handleChangeDate = (val) => {
  currentPage.value = 1
  queryBillSummary(dayjs(val).format('YYYY-MM'))
  queryBillList(1)
}

onMounted(() => {
  queryBillSummary(dayjs(new Date()).format('YYYY-MM'))
  queryBillList(1)
})

</script>

<style lang="scss" scoped>
.content_container {
  box-sizing: border-box;
  display: flex;
  flex-direction: column;

  padding: 10px;
  color: #000;
  
  .header {
    display: flex;
    align-items: center;

    .title {
      font-size: 30px;
      margin: 0 20px 0 0;
    }
  }

  .month_bill_chart {
    display: flex;
    gap: 48px;
    width: 100%;
    padding: 48px 0;
    border-bottom: 1px solid #e6e6e6;

    .left_chart {
      position: relative;
      box-sizing: border-box;
      display: flex;
      flex-direction: column;
      gap: 10px;
      width: 350px;
      min-width: 350px;

      .left_chart_item {
        display: flex;
        align-items: center;
        gap: 8px;
        width: 100%;
        height: 20px;

        .item_icon {
          width: 16px;
          height: 16px;
        }

        .item_name {
          flex: 1 1;
          font-size: 16px;
          font-family: -inter;
        }

        .item_value_waived {
          font-weight: 700;
          font-size: 16px;
          text-decoration: line-through;
        }

        .item_value_strong {
          font-size: 24px !important;
        }

        .item_value {
          font-weight: 700;
          font-size: 16px;
        }
      }

      .left_chart_percent {
        width: 100%;
        height: 40px;
        margin-top: 30px;
        background-color: #f2f2f2;
        border-radius: 8px;
      }

      .left_chart_status {
        box-sizing: border-box;
        position: absolute;
        bottom: 0;
        left: 0;
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        margin-top: 26px;
        
        .date {
          color: grey;
          font-size: 14px;
          font-family: -inter;
        }

        .status {
          display: flex;
          gap: 6px;
          align-items: center;
          padding: 5px 12px;
          background-color: #f2f2f2;
          border-radius: 36px;

          .signle {
            width: 8px;
            height: 8px;
            border-radius: 8px;
          }

          .pending {
            background-color: rgb(255, 204, 51);
          }
          
          .settled {
            background-color: rgb(16, 194, 162);
          }
        }
      }
    }

    .line {
      width: 1px;
      height: 325px;
      background-color: #e6e6e6;
    }

    .right_chart {
      flex: 1 1;

      .right_chart_desc {
        display: flex;
        flex-direction: column;
        width: 100%;

        .screen_container {
          display: flex;
          align-items: center;
          justify-content: space-between;
          width: 100%;

          .name {
            font-weight: 700;
            font-size: 24px;
          }
        }

        .average {
          margin-top: 8px;
          color: grey;
          font-size: 12px;
          font-family: -inter;
        }
      }
    }
  }

  .record {
    display: flex;
    flex-direction: column;

    .title {
      margin: 56px 0 32px 0;
      display: flex;
      align-items: center;
      justify-content: space-between;

      span {
        font-size: 24px;
        font-weight: 700;
      }
    }
  }
}

:deep(.el-input__wrapper) {
  border-radius: 56px;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.2);
}

:deep(.el-select--large .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}

.footerPage {
  width: 100%;
  display: flex;
  justify-content: center;
}

:deep(.el-pagination.is-background .el-pager li),
:deep(.el-pagination.is-background .btn-prev),
:deep(.el-pagination.is-background .btn-next) {
  border-radius: 24px;
}

:deep(.el-pagination.is-background .el-pager li.is-active) {
  background-color: #000;
}

:deep(.el-pagination.is-background .el-pager li):hover {
  color: #000;
}
:deep(.el-pagination.is-background .el-pager li.is-active):hover {
  color: #fff;
}
</style>

<!-- 月份选择器样式 -->
<style lang="scss">
.el-month-table td.current:not(.disabled) .cell {
  background-color: #000 !important;
}

.el-month-table td.today .cell {
  color: #000;
}

.el-month-table td span:hover {
  color: #fff !important;
  background-color: grey;
}
</style>