<template>
  <div>
    <p>准备情况查看</p>
    <el-button @click="getList">刷新</el-button>
    <el-table :data="readies" v-loading="loading">
      <el-table-column prop="status" label="状态" >
        <template slot-scope="scope">
          <a>{{ status_arr[scope.row.status] }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="process" label="进度" >
        <template slot-scope="scope">
          <a>{{ process_arr[scope.row.process] }}</a>
        </template>
      </el-table-column>
      <el-table-column label="玩家1 ID" >
        <template slot-scope="scope">
          <a :style="scope.row.status == 1 ? 'color: blue': scope.row.players[0].ready ? 'color: green' : 'color: red'">{{ scope.row.players[0].team_id + '-' + scope.row.players[0].name }}</a>
        </template>
      </el-table-column>
      <el-table-column label="玩家2 ID" >
        <template slot-scope="scope">
          <a :style="scope.row.status == 1 ? 'color: blue': scope.row.players[1].ready ? 'color: green' : 'color: red'">{{ scope.row.players[1].team_id + '-' + scope.row.players[1].name }}</a>
        </template>
      </el-table-column>
      <el-table-column label="玩家3 ID" >
        <template slot-scope="scope">
          <a :style="scope.row.status == 1 ? 'color: blue': scope.row.players[2].ready ? 'color: green' : 'color: red'">{{ scope.row.players[2].team_id + '-' + scope.row.players[2].name }}</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data () {
    return {
      status_arr: ['等待中', '进行中', '准备完毕', '已结束'],
      // process_arr: ['先锋', '中坚', '大将前半', '大将后半'],
      process_arr: ['先锋前半', '先锋后半', '中坚前半', '中坚后半', '大将前半', '大将后半'],
      readies: [],
      loading: false
    }
  },
  created: function() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true;
      axios
        .get("http://47.100.50.175:8088/api/admin/get_status?round=9", {
          headers: { Authorization: this.$store.state.admin_key }
        }).then(response => {
          if (response.data.busy) {
            this.$message.error('服务器正忙')
          } else {
            this.readies = response.data
          }
          this.loading = false
        })
    }
  }
}
</script>