<template>
  <div>
    <el-page-header @back="$router.push('/')" content="统计信息" />
      <div style="margin: 30px;">
      <el-select v-model="round">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p>当前暂无统计信息……</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      round: 1,
      options: [{
        value: 1,
        label: 'X 进 64'
      },
      {
        value: 2,
        label: '64 进 32'
      },
      {
        value: 3,
        label: '32 进 16'
      },
      {
        value: 4,
        label: '16 进 8'
      },
      {
        value: 5,
        label: '半决赛（8 进 4）'
      },
      {
        value: 6,
        label: '决赛'
      }],
    }
  },
}
</script>