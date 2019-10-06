<template>
  <div>
    <el-page-header @back="$router.push('/')" content="队伍管理"></el-page-header>
    <el-container style="width: 80%, margin-top: 40px">
      <el-table :data="players" v-loading="loading">
        <el-table-column prop="id" label="ID" width="150" />
        <el-table-column prop="name" label="昵称" min-width="150" />
        <el-table-column prop="leader" label="身份" width="150" />
      </el-table>
    </el-container>
    <el-footer style="margin-top: 20px">
      <el-button :disabled="$store.state.team_id != 200 && players.length >= 5" type="primary" @click="addVisible = true">添加新队员</el-button>
    </el-footer>
    <p style="font-size: 18px; color: #303133">出场顺序管理</p>
    <el-container v-loading="positionLoading" style="width: 80%, margin-top: 40px">
      <el-form :label-position="'left'" label-width="50px">
        <el-form-item label="先锋">
          <el-select clearable v-model="xf" placeholder="请选择">
            <el-option v-for="player in players" :key="player.id" :label="player.name" :value="player.name" :disabled="player.disabled" />
          </el-select>
        </el-form-item>
        <el-form-item label="中坚">
          <el-select clearable v-model="zj" placeholder="请选择">
            <el-option v-for="player in players" :key="player.id" :label="player.name" :value="player.name" :disabled="player.disabled" />
          </el-select>
        </el-form-item>
        <el-form-item label="大将">
          <el-select clearable v-model="dj" placeholder="请选择">
            <el-option v-for="player in players" :key="player.id" :label="player.name" :value="player.name" :disabled="player.disabled" />
          </el-select>
        </el-form-item>
      </el-form>
      <el-footer>
        <el-button type="primary" @click="submit_position" :disabled="!(xf !== '' && zj !== '' && dj !== '')">提交上场次序</el-button>
      </el-footer>
    </el-container>

    <el-dialog title="添加队员" :visible.sync="addVisible">
      <el-form style="margin: 0px 20px" v-loading="addLoading">
        <el-form-item label="队员ID">
          <el-input type="number" v-model="add_id"></el-input>
        </el-form-item>
        <el-button type="primary" @click="add_players">确定</el-button>
        <el-button @click="addVisible = false">取消</el-button>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      loading: false,
      players: [],
      add_id: "",
      addVisible: false,
      addLoading: false,
      positionLoading: false,
      xf: '',
      zj: '',
      dj: ''
    };
  },
  created: function() {
    this.loading = true;
    this.getList()
  },
  watch: {
    xf: function(val, v) {
      this.setDisabled(val, v)
    },
    zj: function(val, v) {
      this.setDisabled(val, v)
    },
    dj: function(val, v) {
      this.setDisabled(val, v)
    }
  },
  methods: {
    setDisabled(p, k) {
      for (let player of this.players) {
        if (p === player.name) {
          player.disabled = true
        }
        if (k === player.name) {
          player.disabled = false
        }
      }
    },
    getList() {
      axios
      .get(
        "http://47.100.50.175:8088/api/public/get_players?id=" +
          this.$store.state.team_id
      )
      .then(response => {
        this.players = response.data.players;
        this.loading = false;
        for (let player of this.players) {
          if (player.position === 0) {
            this.xf = player.name
          } else if (player.position === 1) {
            this.zj = player.name
          } else if (player.position === 2) {
            this.dj = player.name
          }
          if (player.leader === true) {
            player.leader = "队长";
          } else {
            player.leader = "队员";
          }
        }
      });
    },
    add_players() {
      if (!this.add_id) {
        this.$message.error("字段不能为空")
        return
      }
      this.addLoading = true
      axios.post('http://47.100.50.175:8088/api/team/add_player', {
        "id": parseInt(this.add_id)
      }, { headers: { Authorization: this.$store.state.jwt }}).then(response => {
        if (response.data.reg) {
          this.addLoading = false
          this.$message.error("看起来这个玩家已经报过名了呢…")
          return
        }
        this.addLoading = false
        this.loading = true
        this.getList()
        this.addVisible = false
        this.$message.success("添加成功")
      }).catch(() => {
        this.addLoading = false
        this.$message.error("没能找到ID呢……是不是输错了？")
      })
    },
    submit_position() {
      this.positionLoading = true
      axios.post('http://47.100.50.175:8088/api/team/position', {
        positions: [
          {
            name: this.xf,
            position: 0
          },
          {
            name: this.zj,
            position: 1
          },
          {
            name: this.dj,
            position: 2
          }
        ]
      }, { headers: { Authorization: this.$store.state.jwt }}).then(() => {
        this.positionLoading = false
        this.$message.success('更改成功')
      })
    }
  }
};
</script>
