<template>
  <div>
    <el-page-header @back="$router.push('/')" content="决赛统计信息" />
    <div style="margin: 30px;">
      <p>Tip: 使用Ctrl+F可以搜索队伍</p>
      <div v-loading="loading" v-if="mode === 0">
        <div v-for="group in groups" :key="group.teams[0].team_id">
          <el-table style="margin: 10px" :border="true" :data="group.teams" :default-sort="{prop: 'scores5', order: 'descending'}">
            <el-table-column prop="team_id" label="队伍ID" width="130px" />
            <el-table-column prop="team_name" label="队名" width="300px">
              <template slot-scope="scope">
                <a>{{ scope.row.team_name }}</a>
                <el-tag type="success" size="small" v-if="map[scope.row.team_id] > 0">+{{map[scope.row.team_id] * 5000}}</el-tag>
                <el-tag type="danger" size="small" v-if="map[scope.row.team_id] < 0">{{map[scope.row.team_id] * 5000}}</el-tag>
              </template>
            </el-table-column>
            <el-table-column v-for="set in sets" :key="set" :prop="'scores' + set" sortable :label="names[set]" />
          </el-table>
        </div>
      </div>
      <div v-loading="loading" v-if="mode === 1">
        <div v-for="group in groups" :key="group.area">
          <p style="font-size: 18px">{{ area_name[group.area] }}</p>
          <el-table style="margin: 10px" :border="true" :data="group.teams" :default-sort="{prop: 'scores3', order: 'descending'}">
            <el-table-column prop="team_id" label="队伍ID" width="130px" />
            <el-table-column prop="team_name" label="队名" width="300px">
              <template slot-scope="scope">
                <a>{{ scope.row.team_name }}</a>
                <el-tag type="success" size="small" v-if="team_ids.indexOf(scope.row.team_id) !== -1">+20000</el-tag>
              </template>
            </el-table-column>
            <el-table-column v-for="set in sets" :key="set" :prop="'scores' + set" sortable :label="names[set]" />
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      sets: [],
      names: ["先锋", "中坚", "大将前半", "大将后半"],
      loading: false,
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
      mode: 0,
      areas: [0, 1, 2, 3],
      area_name: ["东区", "南区", "西区", "北区"],
      groups: [],
      team_ids: [58, 159, 116, 61],
      map: null,
    };
  },
  created: function() {
      this.sets = [0, 1, 2, 3, 4, 5]
      this.names =  ["先锋前半", "先锋后半", "中坚前半", "中坚后半", "大将前半", "大将后半"]
      this.getList()
      this.map = new Map()
      this.map[62] = 0
      this.map[124] = 0
      this.map[159] = 0
      this.map[9] = 0
      this.map[110] = 0
      this.map[7] = 0
      this.map[79] = 0
      this.map[56] = 0
      this.map[90] = 0
      this.map[157] = 0
      this.map[12] = 0
      this.map[155] = 0
      this.map[200] = 0
  },
  watch: {
    round() {
      this.sets = (this.round < 3) ? [0, 1, 2, 3] : [0, 1, 2, 3, 4, 5]
      this.names = (this.round < 3) ? ["先锋", "中坚", "大将前半", "大将后半"] : ["先锋前半", "先锋后半", "中坚前半", "中坚后半", "大将前半", "大将后半"]
      this.getList()
    }
  },
  methods: {
    formatData() {
      for (let group of this.groups) {
        for (let team of group.teams) {
          if (team.scores === null) {
            team.scores = [];
          } else {
            for (let i = 0; i <= 5; i++) {
              if (team.scores[i]) {
                team.scores[i] += this.map[team.team_id] * 5000
                eval("team.scores" + i + "=team.scores[" + i + "]")
              }
            }
            team.sum = this.sum(team.scores)
          }
        }
        const min = Math.min(
          this.sum(group.teams[0].scores),
          this.sum(group.teams[1].scores),
          this.sum(group.teams[2].scores)
        );
        for (let team of group.teams) {
          if (this.sum(team.scores) == min) {
            team.levelup = false;
          } else {
            team.levelup = true;
          }
        }
      }
    },
    getList() {
      this.loading = true
      axios
        .get("http://47.100.50.175:8088/api/public/get_stats?round=9")
        .then(resp => {
          this.groups = resp.data;
          this.formatData();
          this.loading = false
        });
    },
    getByArea() {
      this.loading = true
      axios
        .get("http://47.100.50.175:8088/api/public/get_area_stats?round=9")
        .then(resp => {
          this.groups = resp.data;
          this.formatData();
          this.loading = false
        });
    },
    sum(arr) {
      let ret = 0;
      for (const ele of arr) {
        ret += ele;
      }
      return Math.round(ret / 100) * 100;
    }
  }
};
</script>