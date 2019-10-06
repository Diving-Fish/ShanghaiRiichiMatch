<template>
  <div>
    <el-page-header @back="$router.push('/')" content="决赛分组信息" />
    <div style="margin: 30px;">
      <div v-for="i in areas" :key="i">
        <div v-for="group in groups.filter(a => a.area === i)" :key="group.teams[0].team_id">
          <el-table style="margin: 10px" :border="true" :data="arrange(group)">
            <el-table-column prop="team_id" label="队伍ID" width="130px"/>
            <el-table-column prop="team_name" label="队名" />
            <el-table-column prop="xf" label="先锋" />
            <el-table-column prop="zj" label="中坚" />
            <el-table-column prop="dj" label="大将" />
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      areas: [0, 1, 2, 3],
      area_name: ['东区', '南区', '西区', '北区'],
      groups: [],
      round: 3,
      options: [{
        value: 0,
        label: '第一轮'
      }, {
        value: 1,
        label: '第二轮'
      }, {
        value: 2,
        label: '第三轮'
      }, {
        value: 3,
        label: '第四轮'
      }],
    }
  },
  created: function() {
    axios.get('http://47.100.50.175:8088/api/public/get_groups?round=9').then(resp => {
      this.groups = resp.data
    })
  },
  watch: {
    round() {
      axios.get('http://47.100.50.175:8088/api/public/get_groups?round=9').then(resp => {
        this.groups = resp.data
      })
    }
  },
  methods: {
    arrange(group) {
      let ret = []
      for (const team of group.teams) {
        let xf, zj, dj
        let temp = team.players.filter(p => p.position === 0)
        xf = (temp.length === 0) ? '' : temp[0].name
        temp = team.players.filter(p => p.position === 1)
        zj = (temp.length === 0) ? '' : temp[0].name
        temp = team.players.filter(p => p.position === 2)
        dj = (temp.length === 0) ? '' : temp[0].name
        ret.push({
          team_id: team.team_id,
          team_name: team.team_name,
          xf: xf,
          zj: zj,
          dj: dj,
        })
      }
      return ret
    },
  }
}
</script>