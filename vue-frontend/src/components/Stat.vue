<template>
  <div>
    <el-page-header @back="$router.push('/')" content="统计信息" />
    <div style="margin: 30px;">
      <div style="display: flex">
        <el-select style="margin-bottom: 30px" v-model="round">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </div>
      <el-table :border="true" :data="filter_data" v-loading="loading">
        <el-table-column prop="rank" label="排名" />
        <el-table-column prop="school" min-width="120" label="学校" />
        <el-table-column prop="nick_name" min-width="120" label="昵称" />
        <el-table-column prop="game_name" min-width="120" label="雀魂昵称" />
        <el-table-column v-for="a in point_map.get(round)" :key="a" :label="'马点' + a" >
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
      round: 6,
      filt1: false,
      filt2: false,
      data: [],
      point_map: new Map([
        [1, [1,2,3,4,5,6]],
        [2, [1,2,3]],
        [3, [1,2,3]],
        [4, [1,2,3,4]],
        [5, [1,2,3,4]],
        [6, [1,2,3,4]],
        [7, [1,2,3,4]],
        [8, [1,2,3,4,5]]
      ]),
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
    this.getList();
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
          return (a.check_in == 1);
        })
      }
      let rank = 1
      for (let ob of d) {
        ob.rank = rank;
        rank += 1;
      }
      return d
    }
  },
  watch: {
    round() {
      this.getList();
    }
  },
  methods: {
    getList() {
      this.loading = true;
      axios.get("http://47.100.50.175:8088/api/public/all_scores?round=" + this.round).then(response => {
        let d = response.data
        if (d == null) {
          this.data = [];
          this.loading = false;
          return;
        }
        for (let ob of d) {
          ob.school = this.school_map.get(ob.school)
          ob.s = 0
          for (let j of this.point_map.get(this.round)) {
            let i = j - 1;
            eval("ob.s" + i + " = ob.scores[" + i + "] ? ob.scores[" + i + "] : null")
            eval("ob.s += ob.s" + i)
          }
          ob.s = parseFloat(ob.s.toFixed(1))
        }
        d.sort((a, b) => {
          return b.s - a.s
        })
        this.data = d;
        this.loading = false
      })
    }
  }
}
</script>