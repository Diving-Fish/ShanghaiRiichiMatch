<template>
  <div>
    <el-page-header @back="$router.push('/')" content="管理"></el-page-header>
    <el-container v-if="$store.state.sid == 0" style="width: 80%, margin-top: 40px">
      <el-table :data="players" v-loading="loading">
        <el-table-column prop="sid" label="ID" width="150" />
        <el-table-column prop="nickname" label="昵称"/>
        <el-table-column prop="game_id" label="雀魂ID"/>
        <el-table-column prop="game_name" label="雀魂昵称"/>
      </el-table>
      <el-footer style="margin-top: 20px">
        <el-button @click="open" :disabled="players.length >= 11" type="primary">添加新队员</el-button>
      </el-footer>
    </el-container>
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
    open() {
       this.$confirm('确定要添加新队员吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }).then(() => {
        axios.get('http://47.100.50.175:8088/api/admin/apply', {
        headers: { Authorization: this.$store.state.jwt }
        }).then(response => {
          this.$alert('用户名：' + response.data.username + '   密码：' + response.data.password, '添加完毕',
          {
            confirmButtonText: '确定'
          })
          this.getList();
        })
      }).catch(() => {

      });
    },
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
      axios.get('http://47.100.50.175:8088/api/admin/get', {
        headers: { Authorization: this.$store.state.jwt }
      }).then(response => {
        console.log(response)
        this.players = response.data;
        this.loading = false;
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
