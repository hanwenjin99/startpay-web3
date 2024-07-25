<template>
  <div class="gva-form-box">
    <div class="grid grid-cols-12 w-full gap-2">
      <div class="col-span-3 h-full">
        <div class="w-full h-full bg-white dark:bg-slate-900 px-4 py-8 rounded-lg shadow-lg box-border">
          <div class="user-card px-6 text-center bg-white dark:bg-slate-900 shrink-0">
            <div class="flex justify-center">
              <SelectImage
                v-model="userStore.userInfo.headerImg"
                file-type="image"
              />
            </div>
            <div class="py-6 text-center">
              <p
                v-if="!editFlag"
                class="text-3xl flex justify-center items-center gap-4"
              >
                {{ userStore.userInfo.nickName }}
                <el-icon
                  class="cursor-pointer text-sm"
                  color="#66b1ff"
                  @click="openEdit"
                >
                  <edit />
                </el-icon>
              </p>
              <p
                v-if="editFlag"
                class="flex justify-center items-center gap-4"
              >
                <el-input v-model="nickName" />
                <el-icon
                  class="cursor-pointer"
                  color="#67c23a"
                  @click="enterEdit"
                >
                  <check />
                </el-icon>
                <el-icon
                  class="cursor-pointer"
                  color="#f23c3c"
                  @click="closeEdit"
                >
                  <close />
                </el-icon>
              </p>
              <p class="text-gray-500 mt-2 text-md">这个家伙很懒，什么都没有留下</p>
            </div>
            <div class="w-full h-full text-left">
              <ul class="inline-block h-full w-full">
                <li class="info-list">
                  <el-icon>
                    <user />
                  </el-icon>
                  {{ userStore.userInfo.nickName }}
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="香港斯特派有限公司
                           Start Pay Technology Limited"
                  placement="top"
                >
                  <li class="info-list">
                    <el-icon>
                      <data-analysis />
                    </el-icon>
                   香港斯特派有限公司
                   Start Pay Technology Limited
                  </li>
                </el-tooltip>
                <li class="info-list">
                  <el-icon>
                    <video-camera />
                  </el-icon>
                  中国·上海市·闵行区
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="GoLang/JavaScript/Vue/Gorm"
                  placement="top"
                >
                  <li class="info-list">
                    <el-icon>
                      <medal />
                    </el-icon>
                    GoLang/JavaScript/Vue/Gorm
                  </li>
                </el-tooltip>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-9 ">
        <div class="bg-white dark:bg-slate-900 h-full px-4 py-8 rounded-lg shadow-lg box-border">
          <el-tabs
            v-model="activeName"
            @tab-click="handleClick"
          >
            <el-tab-pane
              label="账号绑定"
              name="second"
            >
              <ul>
                <li class="borderd">
                  <p class="pb-2.5 text-xl text-gray-600">密保手机</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    已绑定手机:{{ userStore.userInfo.phone }}
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="changePhoneFlag = true"
                    >立即修改</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">密保邮箱</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    已绑定邮箱：{{ userStore.userInfo.email }}
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="changeEmailFlag = true"
                    >立即修改</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">密保问题</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    未设置密保问题
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                    >去设置</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">修改密码</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    修改个人密码
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="showPassword = true"
                    >修改密码</a>
                  </p>
                </li>
                <!-- 新增虚拟MFA绑定/解绑 -->
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">虚拟MFA</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    绑定虚拟MFA后，您可以在登录时通过他来进行二次校验，为了方便您获取虚拟MFA动态码，绑定前请确保您已下载MFA设备
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="showMfa = true"
                    >{{ isBindMfa ? '前往解绑' : '前往绑定' }}</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>

    <el-dialog
      v-model="showPassword"
      title="修改密码"
      width="360px"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          :minlength="6"
          label="原密码"
          prop="password"
        >
          <el-input
            v-model="pwdModify.password"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          label="新密码"
          prop="newPassword"
        >
          <el-input
            v-model="pwdModify.newPassword"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          label="确认密码"
          prop="confirmPassword"
        >
          <el-input
            v-model="pwdModify.confirmPassword"
            show-password
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button

            @click="showPassword = false"
          >取 消</el-button>
          <el-button

            type="primary"
            @click="savePassword"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changePhoneFlag"
      title="绑定手机"
      width="600px"
    >
      <el-form :model="phoneForm">
        <el-form-item
          label="手机号"
          label-width="120px"
        >
          <el-input
            v-model="phoneForm.phone"
            placeholder="请输入手机号"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="验证码"
          label-width="120px"
        >
          <div class="flex w-full gap-4">
            <el-input
              v-model="phoneForm.code"
              class="flex-1"
              autocomplete="off"
              placeholder="请自行设计短信服务，此处为模拟随便写"
              style="width:300px"
            />
            <el-button
              type="primary"
              :disabled="time>0"
              @click="getCode"
            >{{ time>0?`(${time}s)后重新获取`:'获取验证码' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button

            @click="closeChangePhone"
          >取消</el-button>
          <el-button
            type="primary"

            @click="changePhone"
          >更改</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changeEmailFlag"
      title="绑定邮箱"
      width="600px"
    >
      <el-form :model="emailForm">
        <el-form-item
          label="邮箱"
          label-width="120px"
        >
          <el-input
            v-model="emailForm.email"
            placeholder="请输入邮箱"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="验证码"
          label-width="120px"
        >
          <div class="flex w-full gap-4">
            <el-input
              v-model="emailForm.code"
              class="flex-1"
              placeholder="请自行设计邮件服务，此处为模拟随便写"
              autocomplete="off"
              style="width:300px"
            />
            <el-button
              type="primary"
              :disabled="emailTime>0"
              @click="getEmailCode"
            >{{ emailTime>0?`(${emailTime}s)后重新获取`:'获取验证码' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button

            @click="closeChangeEmail"
          >取消</el-button>
          <el-button
            type="primary"

            @click="changeEmail"
          >更改</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 绑定MFA -->
    <el-dialog
      v-model="showMfa"
      :title="`${ isBindMfa ? '需要验证您的身份' : '绑定虚拟MFA' }`"
      width="600px"
    >
      <template v-if="isBindMfa">
        <div class="unbind">
          <span>为确认是您本人操作，请进行身份验证。</span>
          <span style="margin: 20px 0;">请输入从已绑定的虚拟 MFA 应用中获取的 6 位数字验证码</span>
          <div class="input-box">
            <div
              v-for="item in [0, 1, 2, 3, 4, 5]"
              :key="item"
              :class="['code-item', codeValue.length === item && inputFocus ? 'code-item-active' : '']"
            >
              {{ codeValue[item] }}
            </div>
            <el-input
              class="input-code"
              :model-value="codeValue"
              maxlength="6"
              @input="codeInputChange"
              @blur="inputFocus = fasle"
              @focus="inputFocus = true"
            />
          </div>
        </div>
      </template>
      <template v-else>
        <div style="height: 450px">
          <el-steps direction="vertical">
            <el-step title="第一步：下载应用" description="在手机上安装支持虚拟 MFA 的应用。" />
            <el-step title="第二步：获取动态码">
              <template #description>
                <div class="mfaCode">
                  打开应用，通过扫描下方二维码或手动输入下方密钥信息添加帐号，获取虚拟 MFA 动态码。
                  <div class="qrcode-box">
                    <img class="qrcode-img" src="" alt="">
                    <div class="qrcode-desc">
                      <div class="columnFlex">
                        <span>帐号名</span>
                        <span>haohaohuaihuai</span>
                      </div>
                      
                      <div class="columnFlex">
                        <span>密钥</span>
                        <span>UJAXPTVWRS4OVA26Y77ZIIXLS4Q7PAFXEZJYDGGKG6PUCMYUNUPR2KKDEV56B3VU</span>
                      </div>
                    </div>
                  </div>
                </div>
              </template>
            </el-step>
            <el-step title="第三步：输入动态码">
              <template #description>
                请输入从虚拟 MFA 应用中获取的连续两组 6 位数字验证码。
                <el-form
                  label-position="right"
                  label-width="auto"
                  :model="formLabelAlign"
                  style="max-width: 400px; margin-top: 10px;"
                >
                  <el-form-item label="第一组验证码">
                    <el-input v-model="formLabelAlign.firstIdentityCode" placeholder="6位数字验证码" />
                  </el-form-item>
                  <el-form-item label="第二组验证码">
                    <el-input v-model="formLabelAlign.secondIdentityCode" placeholder="6位数字验证码" />
                  </el-form-item>
                </el-form>
              </template>
            </el-step>
          </el-steps>
        </div>
      </template>
      <template #footer>
        <el-button plain color="#000">取消</el-button>
        <el-button color="#000">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { setSelfInfo, changePassword } from '@/api/user.js'
import { reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import SelectImage from '@/components/selectImage/selectImage.vue'

defineOptions({
  name: 'Person',
})

const isBindMfa = ref(true)
const showMfa = ref(false)
const formLabelAlign = reactive({
  firstIdentityCode: '',
  secondIdentityCode: ''
})
const codeValue = ref('')
const inputFocus = ref(false)
// 解绑时输入验证码
const codeInputChange = (e) => {
  if (e) {
    if (/^\+?[0-9][0-9]*$/.test(e)) {
      codeValue.value = e
    }
  } else {
    codeValue.value = ''
  }
}

const activeName = ref('second')
const rules = reactive({
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请输入确认密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error('两次密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const modifyPwdForm = ref(null)
const showPassword = ref(false)
const pwdModify = ref({})
const nickName = ref('')
const editFlag = ref(false)
const savePassword = async() => {
  modifyPwdForm.value.validate((valid) => {
    if (valid) {
      changePassword({
        password: pwdModify.value.password,
        newPassword: pwdModify.value.newPassword,
      }).then((res) => {
        if (res.code === 0) {
          ElMessage.success('修改密码成功！')
        }
        showPassword.value = false
      })
    } else {
      return false
    }
  })
}

const clearPassword = () => {
  pwdModify.value = {
    password: '',
    newPassword: '',
    confirmPassword: '',
  }
  modifyPwdForm.value.clearValidate()
}

watch(() => userStore.userInfo.headerImg, async(val) => {
  const res = await setSelfInfo({ headerImg: val })
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: val })
    ElMessage({
      type: 'success',
      message: '设置成功',
    })
  }
})

const openEdit = () => {
  nickName.value = userStore.userInfo.nickName
  editFlag.value = true
}

const closeEdit = () => {
  nickName.value = ''
  editFlag.value = false
}

const enterEdit = async() => {
  const res = await setSelfInfo({
    nickName: nickName.value
  })
  if (res.code === 0) {
    userStore.ResetUserInfo({ nickName: nickName.value })
    ElMessage({
      type: 'success',
      message: '设置成功',
    })
  }
  nickName.value = ''
  editFlag.value = false
}

const handleClick = (tab, event) => {
  console.log(tab, event)
}

const changePhoneFlag = ref(false)
const time = ref(0)
const phoneForm = reactive({
  phone: '',
  code: ''
})

const getCode = async() => {
  time.value = 60
  let timer = setInterval(() => {
    time.value--
    if (time.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangePhone = () => {
  changePhoneFlag.value = false
  phoneForm.phone = ''
  phoneForm.code = ''
}

const changePhone = async() => {
  const res = await setSelfInfo({ phone: phoneForm.phone })
  if (res.code === 0) {
    ElMessage.success('修改成功')
    userStore.ResetUserInfo({ phone: phoneForm.phone })
    closeChangePhone()
  }
}

const changeEmailFlag = ref(false)
const emailTime = ref(0)
const emailForm = reactive({
  email: '',
  code: ''
})

const getEmailCode = async() => {
  emailTime.value = 60
  let timer = setInterval(() => {
    emailTime.value--
    if (emailTime.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangeEmail = () => {
  changeEmailFlag.value = false
  emailForm.email = ''
  emailForm.code = ''
}

const changeEmail = async() => {
  const res = await setSelfInfo({ email: emailForm.email })
  if (res.code === 0) {
    ElMessage.success('修改成功')
    userStore.ResetUserInfo({ email: emailForm.email })
    closeChangeEmail()
  }
}

</script>

<style lang="scss" scoped>
.mfaCode {
  display: flex;
  flex-direction: column;
  font-size: 12px;

  .qrcode-box {
    margin: 10px 0;
    background: #f2f2f2;
    padding: 15px;
    border-radius: 8px;
    display: flex;

    .qrcode-img {
      margin-right: 10px;
      width: 115px;
      height: 115px;
    }

    .qrcode-desc {
      display: flex;
      flex-direction: column;

      .columnFlex {
        display: flex;
        flex-direction: column;
        max-width: 330px;
        font-size: 12px;

        &:first-child {
          margin-bottom: 15px
        }
      }
    }
  }
}

.unbind {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 12px
}

.input-box {
  position: relative;
  display: flex;
  margin-bottom: 10px;
  .input-code {
    position: absolute;
  }
  .code-item {
    width: 64px;
    height: 64px;
    border-radius: 8px;
    margin-right: 10px;
    line-height: 64px;
    text-align: center;
    font-size: 24px;
    color: #000;
    background-color: rgba(0, 0, 0, 0.05);
    border: 1px solid transparent;
  }

  .code-item-active {
    border: 1px solid rgba(0, 0, 0, 0.2);
  }

  :deep(.el-input__wrapper) {
    width: 444px;
    height: 64px;
    background-color: transparent;
    border: none;
    box-shadow: none;
  }

  :deep(.el-input__inner) {
    color: transparent;
  }
}

:deep(.el-step__head.is-wait) {
  color: #000 !important;
  border-color: #000 !important;
}

:deep(.el-step__line) {
  background-color: #000 !important;
}

:deep(.el-step__title.is-wait) {
  color: #000 !important;
  font-weight: bold;
}

:deep(.el-step__description.is-wait) {
  color: #000 !important;
}
</style>

<style lang="scss">
.borderd {
  @apply border-b-2 border-solid border-gray-100 dark:border-gray-500 border-t-0 border-r-0 border-l-0;
    &:last-child{
      @apply border-b-0;
    }
 }

.info-list{
  @apply w-full whitespace-nowrap overflow-hidden text-ellipsis py-3 text-lg text-gray-700
}

</style>
