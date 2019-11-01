<template>
  <div>
    <el-page-header @back="$router.push('/')" content="分组信息" />
      <div style="margin: 30px;">
      <el-select v-model="round">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <el-select style="margin-left: 20px" v-model="process">
        <el-option v-for="item in bz_map.get(round)" :key="item" :label="'第' + item + '半庄'" :value="item" />
      </el-select>
      <el-table style="margin-top: 30px" :border="true" :data="data" v-loading="loading">
        <el-table-column label="玩家1">
          <template slot-scope="scope">

          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      loading: false,
      round: 4,
      process: 1,
      data: [],
      bz_map: new Map([
        [4, [1, 2, 3, 4]],
        [5, [1]],
        [6, [1]],
        [7, [1]],
        [8, [1]]
      ]),
      options: [{
        value: 4,
        label: '64 进 32'
      },
      {
        value: 5,
        label: '32 进 16'
      },
      {
        value: 6,
        label: '16 进 8'
      },
      {
        value: 7,
        label: '半决赛（8 进 4）'
      },
      {
        value: 8,
        label: '决赛'
      }],
    }
  },
  methods: {
    getList() {
      this.loading = true;
      axios.get("http://47.100.50.175:8088/api/public/get_group?round=" + this.round + "&process=" + this.process).then(response => {
        this.data = response.data;
        this.loading = false;
      })
    }
  }
}
</script>