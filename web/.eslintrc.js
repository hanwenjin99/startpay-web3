module.exports = {
  root: true,
  parserOptions: {
    parser: '@babel/eslint-parser',
    sourceType: 'module'
  },
  env: {
    browser: true,
    node: true,
    es6: true
  },
  extends: ['plugin:vue/recommended', 'eslint:recommended'],
  rules: {
    "vue/max-attributes-per-line" : 0,
    "vue/no-v-model-argument" : 0,
    "vue/multi-word-component-names": 'off', // 组件名称驼峰要求
    "vue/no-unused-vars": 'off', // 未使用的变量报错
    "vue/singleline-html-element-content-newline": 'off' // 内容新行
  }
}
