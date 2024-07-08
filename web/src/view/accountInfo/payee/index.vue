<template>
  <main class="content_container">
    <h1 class="title">收款人</h1>

    <!-- 搜索 -->
    <section class="top">
      <section class="search">
        <el-input
          v-model="searchString"
          placeholder="名称、地址"
          prefix-icon="search"
          size="large"
          style="width: 600px; margin-right: 16px;"
        />

        <el-select
          v-model="selectInteVal"
          clearable
          placeholder="网络"
          style="width: 100px; margin-right: 20px;"
          size="large"
        >
          <el-option
            v-for="item in inteOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>

        <!-- 重置按钮 -->
        <el-button size="large" color="#000" plain type="info" round>重置</el-button>
      </section>
      <el-button color="#000" icon="plus" type="info" round @click="router.push('addPayee')">添加收款人</el-button>
    </section>

    <!-- 收款人列表 -->
    <el-table :data="list" style="width: 100%">
      <el-table-column>
        <template #default="scope">
          <div class="message">
            <span class="char">{{ scope.row.name.slice(0, 1) }}</span>
            <div class="desc">
              <span>{{ scope.row.name }}</span>
              <span>{{ scope.row.address }}</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column>
        <template #default="scope">
          <div class="chain">
            <img :src="scope.row.chainIcon" alt="">
            <span>{{ scope.row.chain }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #default="scope">
          <el-button icon="delete" type="danger" circle />
          <el-button color="#000" plain round @click.stop="selectPayee">转账</el-button>
        </template>
      </el-table-column>
    </el-table>
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const searchString = ref('')
const selectInteVal = ref('')

const inteOptions = reactive([
  { label: 'ETH', value: 'ETH' },
  { label: 'BSC', value: 'BSC' },
  { label: 'TRON', value: 'TRON' },
  { label: 'POLYGON', value: 'POLYGON' },
  { label: 'BTC', value: 'BTC' },
])

const list = [
  {
    address: "0xa0eb5663d9c027cd6dbf0c27e82ee63fb5caaa76",
    chain: "ETH",
    chainIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
    id: "665e9d4ba98d5b5cbe760d52",
    isInternal: false,
    name: "startpay-bison"
  }
]

const selectPayee = () => {
  router.push('/layout/account/transfer/single')
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
    margin: 0 0 30px 0;
  }

  .top {
    margin-bottom: 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .search {
      display: flex;
      align-items: center;
    }
  }

  .message {
    display: flex;
    align-items: center;
    .char {
      box-sizing: border-box;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 36px;
      height: 36px;
      background-color: #ffefcb;
      border-radius: 36px;
      color: #faaf00;
      font-size: 20px;
      font-weight: 700;
      line-height: 1.25;
    }

    .desc {
      display: flex;
      flex: 3 1;
      flex-direction: column;
      margin-left: 12px;
      line-height: 1.25;

      span {
        font-family: -inter;
        &:first-child {
          font-size: 16px;
          color: #000;
          font-weight: 700;
        }

        &:last-child {
          margin-top: 4px;
          color: grey;
          font-size: 14px;
        }
      }
    }
  }

  .chain {
    display: flex;
    flex: 2 1;
    align-items: center;
    margin-left: 12px;
    font-family: -inter;

    img {
      width: 24px;
      height: 24px;
      margin-right: 8px;
      border-radius: 24px;
    }

    span {
      color: #000;
      font-size: 14px;
    }
  }
}

:deep(.el-input--large .el-input__wrapper),
:deep(.el-select--large .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}

// 隐藏表头
:deep(.el-table thead) {
  display: none;
}
</style>