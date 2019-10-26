<template>
  <div>
    <el-page-header @back="$router.push('/')" content="统计信息" />
    <div style="margin: 30px;">
      <div style="display: flex">
        <el-select style="margin-bottom: 30px" v-model="round">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <el-checkbox style="margin-left: 30px; line-height: 40px" v-model="filt1" v-if="round == 1">仅查看打满6场的玩家</el-checkbox>
        <el-checkbox style="margin-left: 30px; line-height: 40px" v-model="filt2">仅查看已签到的玩家</el-checkbox>
      </div>
      <el-table :border="true" :data="filter_data" v-loading="loading">
        <el-table-column prop="rank" label="排名" />
        <el-table-column prop="school" min-width="120" label="学校" />
        <el-table-column prop="nick_name" min-width="120" label="昵称" />
        <el-table-column prop="game_name" min-width="120" label="雀魂昵称" />
        <el-table-column v-for="a in [1,2,3,4,5,6]" :key="a" :label="'马点' + a" >
          <template slot-scope="scope">
            <a v-if="scope.row.scores[a - 1] > 0" style="color: #00aa00">+{{scope.row.scores[a - 1]}}</a>
            <a v-if="scope.row.scores[a - 1] < 0" style="color: #ff5555">{{scope.row.scores[a - 1]}}</a>
            <a v-if="scope.row.scores[a - 1] == 0">{{scope.row.scores[a - 1]}}</a>
          </template>
        </el-table-column>
        <el-table-column prop="s" label="总马点">
          <template slot-scope="scope">
            <a v-if="scope.row.s > 0" style="color: green">+{{scope.row.s}}</a>
            <a v-if="scope.row.s < 0" style="color: red">{{scope.row.s}}</a>
            <a v-if="scope.row.s == 0">{{scope.row.s}}</a>
          </template>
        </el-table-column>
        <el-table-column prop="check_in" label="签到情况" />
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      loading: false,
      round: 1,
      filt1: false,
      filt2: false,
      data: [],
      school_map: new Map([
        ['ECNU', '华东师范大学'],
        ['SISU', '上海外国语大学'],
        ['FDU', '复旦大学'],
        ['SHU', '上海大学'],
        ['SHMTU','上海海事大学'],
        ['SUIBE','上海对外经贸大学'],
        ['SIT','上海应用技术大学'],
        ['SUES','上海工程技术大学'],
        ['SUFE','上海财经大学'],
        ['USST','上海理工大学'],
        ['ZJU','浙江大学'],
        ['SHUPL','上海政法学院'],
        ['ECUPL','华东政法大学'],
        ['LIXIN','上海立信会计金融学院'],
        ['SHUEP','上海电力大学'],
        ['NJU','南京大学'],
        ['ECUST','华东理工大学'],
        ['SHOU','上海海洋大学'],
        ['SHNU','上海师范大学'],
        ['SJTU','上海交通大学'],
        ['SDJU','上海电机学院'],
        ['DHU','东华大学'],
        ['SEU','东南大学'],
        ['BUAA','学院路联合大学'],
        ['TJU','同济大学'],
        ['STU','上海科技大学'],
        ['SUDA','苏州大学']
      ]),
      options: [{
        value: 1,
        label: 'X 进 128'
      },
      {
        value: 2,
        label: '128 进 96'
      },
      {
        value: 3,
        label: '96 进 64'
      },
      {
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
    this.loading = true;
    axios.get("http://47.100.50.175:8088/api/public/all_scores?round=1").then(response => {
      let d = response.data
      for (let ob of d) {
        ob.school = this.school_map.get(ob.school)
        ob.s0 = ob.scores[0] ? ob.scores[0] : null
        ob.s1 = ob.scores[1] ? ob.scores[1] : null
        ob.s2 = ob.scores[2] ? ob.scores[2] : null
        ob.s3 = ob.scores[3] ? ob.scores[3] : null
        ob.s4 = ob.scores[4] ? ob.scores[4] : null
        ob.s5 = ob.scores[5] ? ob.scores[5] : null
        ob.s = ob.s0 + ob.s1 + ob.s2 + ob.s3 + ob.s4 + ob.s5
        ob.s = parseFloat(ob.s.toFixed(1))
      }
      d.sort((a, b) => {
        return b.s - a.s
      })
      this.data = d;
      this.loading = false
    })
  },
  computed: {
    filter_data: function() {
      let d = this.data
      if (this.filt1) {
        d = d.filter(a => {
          return a.s0 && a.s1 && a.s2 && a.s3 && a.s4 && a.s5;
        })
      }
      if (this.filt2) {
        d = d.filter(a => {
          return a.check_in;
        })
      }
      let rank = 1
      for (let ob of d) {
        ob.rank = rank;
        rank += 1;
      }
      return d
    }
  }
}
</script>