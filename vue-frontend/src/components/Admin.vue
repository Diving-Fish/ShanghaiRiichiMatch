<template>
  <div>
    <el-button @click="$router.push('/admin2')">查看准备情况</el-button>
    <p style="margin-left: 20px">已报名队伍数量：{{ teams.length }}</p>
    <el-container>
      <el-collapse v-loading="loading" v-model="active_teams" style="width: 100%">
        <el-collapse-item v-for="team in teams" :key="team.team_id" :name="team.team_id">
          <template slot="title" v-if="team.team_id === 0">
            <div
              style="width: 100%; font-size: 16px"
            >
              <ul style="float: left">
								<a style="margin-right: 20px">个人报名玩家</a>
							</ul>
            </div>
          </template>
          <template slot="title" v-if="team.team_id !== 0">
            <div
              style="width: 100%; font-size: 16px"
              :style="team.players === null || team.players.length < 3 ? 'color: red' : 'color: black'"
            >
              <ul style="float: left">
								<a style="margin-right: 20px">队伍ID：{{team.team_id}}</a>
								<a style="margin-right: 20px">队伍密码：{{team.password}}</a>
                <a>QQ：{{ team.qq }}</a>
							</ul>
              <ul style="float: right; margin-right: 20px">
                <a style="margin-right: 20px">{{team.team_name}}</a>
                <el-button
                  size="small"
                  @click="del_team(team.team_id)"
                  type="danger"
                  icon="el-icon-delete"
                />
              </ul>
            </div>
          </template>
          <el-table :data="team.players" style="margin: 0px 40px">
            <el-table-column prop="id" label="ID" width="150" />
            <el-table-column prop="name" label="昵称" min-width="150" />
            <el-table-column prop="leader" label="身份" width="150" />
          </el-table>
        </el-collapse-item>
      </el-collapse>
    </el-container>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      loading: false,
      teams: [],
      active_teams: []
    };
  },
  created: function() {
    this.getList();
  },
  methods: {
    getList() {
      this.loading = true;
      axios
        .get("http://47.100.50.175:8088/api/admin/fetch_all_players", {
          headers: { Authorization: this.$store.state.admin_key }
        })
        .then(response => {
          this.teams = response.data;
          for (let team of this.teams) {
            if (team.players === null) continue;
            for (let player of team.players) {
              player.leader = player.leader === true ? "队长" : "队员";
            }
          }
          this.loading = false;
        });
    },
    del_team(team_id) {
      this.$confirm("此操作将删除队伍，是否继续？", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "danger"
      })
        .then(() => {
          this.loading = true;
          axios
            .get(
              "http://47.100.50.175:8088/api/admin/delete_team?team_id=" +
                team_id,
              {
                headers: { Authorization: this.$store.state.admin_key }
              }
            )
            .then(() => {
              this.$message.success("删除成功");
              this.getList();
            });
        })
        .catch(() => {});
    }
  }
};
</script>
