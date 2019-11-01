<template>
  <div>
    <el-page-header @back="$router.push('/')" content="分组信息" />
      <div style="margin: 30px;">
        <p>已签到的玩家将以绿色标识</p>
      <el-select v-model="round">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <el-select style="margin-left: 20px" v-model="process">
        <el-option v-for="item in bz_map.get(round)" :key="item" :label="'第' + item + '半庄'" :value="item" />
      </el-select>
      <el-table style="margin-top: 30px" :border="true" :data="data" v-loading="loading">
        <el-table-column label="桌号">
          <template slot-scope="scope">
            {{ scope.row.group_id + 1 }}
          </template>
        </el-table-column>
        <el-table-column v-for="i in [1, 2, 3, 4]" :key="i" :label="'玩家' + i">
          <template slot-scope="scope">
            <a :style="scope.row.player_list[i - 1].checkin == 1 ? 'color: green' : 'color: red'">{{ scope.row.player_list[i - 1].name }}</a>
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
  created: function() {
    this.getList()
  },
  watch: {
    round: function() {
      this.getList()
    },
    process: function() {
      this.getList()
    }
  },
  methods: {
    getList() {
      this.loading = true;
      axios.get("http://47.100.50.175:8088/api/public/get_group?round=" + this.round + "&process=" + this.process).then(response => {
        if (response.data == null) {
          this.data = [];
          this.loading = false;
          return;
        }
        this.data = response.data;
        this.loading = false;
      })
    }
  }
}
</script>