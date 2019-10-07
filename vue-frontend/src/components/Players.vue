<template>
  <el-table :data="players" v-loading="loading" :default-sort = "{prop: 'school', order: 'ascending'}">
    <el-table-column prop="school" label="学校" sortable />
    <el-table-column prop="nickname" label="昵称" />
    <el-table-column prop="game_name" label="雀魂昵称" />
  </el-table>
</template>

<script>

import axios from 'axios'

export default {
  name: "players",
  data() {
    return {
      loading: false,
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
        ['BUAA','学院路联合大学']
      ]),
      players: []
    }
  },
  created: function() {
    this.loading = true;
    axios.get('http://47.100.50.175:8088/api/public/find_all').then(response => {
      for (let player of response.data) {
        player.school = this.school_map.get(player.school)
        this.players.push(player)
      }
      this.loading = false;
    })
  }
}
</script>