<template>
  <div>
    <el-page-header @back="$router.push('/')" content="统计信息" />
      <div style="margin: 30px;">
      <el-select v-model="round">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <el-table :data="data" v-loading="loading">
        <el-table-column prop="rank" label="排名" />
        <el-table-column prop="school" min-width="120" label="学校" />
        <el-table-column prop="nick_name" min-width="120" label="昵称" />
        <el-table-column prop="game_name" min-width="120" label="雀魂昵称" />
        <el-table-column prop="s0" label="马点1" />
        <el-table-column prop="s1" label="马点2" />
        <el-table-column prop="s2" label="马点3" />
        <el-table-column prop="s3" label="马点4" />
        <el-table-column prop="s4" label="马点5" />
        <el-table-column prop="s5" label="马点6" />
        <el-table-column prop="s" label="总马点" sortable />
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      round: 1,
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
    axios.get("http://47.100.50.175:8088/api/public/all_scores?round=1").then(response => {
      console.log(response)
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
      let rank = 1
      for (let ob of d) {
        ob.rank = rank;
        rank += 1;
      }
      this.data = d;
    })
  }
}
</script>