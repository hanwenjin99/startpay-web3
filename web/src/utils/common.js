import { ElMessage } from 'element-plus'

export const copyMessage = (text) => {
  let oInput = {}
  oInput =  document.querySelector('#copyInput')
  if (!oInput) {
    oInput = document.createElement('textarea') // 多行数据复制时分行展示
    oInput.id = 'copyInput'
    document.body.appendChild(oInput)
  } else {
    oInput.style.display=''
  }
  oInput.value = text
  oInput.select() // 选择对象
  document.execCommand('Copy') // 执行浏览器复制命令
  oInput.style.display = 'none'
  ElMessage.success('复制成功')
}