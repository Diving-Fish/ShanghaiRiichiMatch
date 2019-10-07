<template>
  <div>
    <el-page-header @back="$router.push('/')" content="管理"></el-page-header>
    <el-container v-if="$store.state.sid == 0" style="width: 80%, margin-top: 40px">
      <el-table :data="players" v-loading="loading">
        <el-table-column prop="sid" label="ID" width="150" />
        <el-table-column prop="nickname" label="昵称"/>
        <el-table-column prop="game_id" label="雀魂ID"/>
        <el-table-column prop="game_name" label="雀魂昵称"/>
        <el-table-column label="绑定情况">
          <template slot-scope="scope">
            <a v-if="scope.row.sid == 0">工具人是不用绑定的（大雾）</a>
            <a v-if="scope.row.sid != 0 && scope.row.bound" style="cursor: pointer" @click="reset(scope.row.sid)">已绑定，点击重置</a>
            <a v-if="scope.row.sid != 0 && !scope.row.bound">未绑定，密码为{{scope.row.password}}</a>
          </template>
        </el-table-column>
      </el-table>
      <el-footer style="margin-top: 20px">
        <el-button @click="open" :disabled="players.length >= 11" type="primary">添加新队员</el-button>
      </el-footer>
    </el-container>
    <div v-if="$store.state.sid != 0 && stat.game_id == 0">
      <p style="font-size: 16px; margin-left: 40px; margin-top: 40px">请先填写个人信息</p>
      <el-form
          ref="form"
          :model="form"
          label-width="120px"
          style="width: 40%; margin: 40px 0% 0px 0%"
        >
        <el-form-item label="昵称">
          <el-input v-model="form.nickname"></el-input>
        </el-form-item>
        <el-form-item label="雀魂ID" v-loading="loading">
          <div style="display: flex">
            <el-input v-model="form.game_id" placeholder="请填写雀魂好友界面右上角的8位数字ID"></el-input>
            <el-button style="margin-left: 20px" @click="find">查找玩家</el-button>
          </div>
        </el-form-item>
        <el-form-item label="雀魂昵称">
          <el-input v-model="form.game_name" :disabled="true" />
        </el-form-item>
        <el-button :disabled="!form.game_name || !form.nickname" style="margin-left: 120px;" type="primary" @click="submit">提交</el-button>
      </el-form>
    </div>
    <div v-if="$store.state.sid != 0 && stat.game_id != 0">
      <p style="font-size: 20px; margin-left: 20px; margin-top: 40px">个人信息</p>
      <el-form label-position="left"
          ref="form"
          :model="form"
          label-width="80px"
          style="width: 40%; margin: 0px 0px 0px 20px"
        >
        <el-form-item style="margin-bottom: 0px" label="学校">
          <span>{{ school_map.get(stat.school) }}</span>
        </el-form-item>
        <el-form-item style="margin-bottom: 0px" label="昵称">
          <span>{{ stat.nickname }}</span>
        </el-form-item>
        <el-form-item style="margin-bottom: 0px" label="雀魂ID">
          <span>{{ stat.game_id }}</span>
        </el-form-item>
        <el-form-item style="margin-bottom: 0px" label="雀魂昵称">
          <span>{{ stat.game_name }}</span>
        </el-form-item>
      </el-form>
    </div>
    <el-footer v-if="stat.nickname != 0" style="margin-top: 20px">
      <el-button type="warning" @click="cpwdVisible = true">修改密码</el-button>
    </el-footer>

    <el-dialog title="修改密码" :visible.sync="cpwdVisible">
      <el-form label-position="right" :model="cpwd" label-width="120px" style="margin-right: 120px">
        <el-form-item label="新密码">
          <el-input type="password" v-model="cpwd.password" />
        </el-form-item>
        <el-form-item label="确认新密码">
          <el-input type="password" v-model="cpwd.retype_password" />
        </el-form-item>
        <p style="margin-left: 120px; color: red" v-if="cpwd.password && cpwd.password.length < 6">密码长度最少为6</p>
        <p style="margin-left: 120px; color: red" v-if="cpwd.password && cpwd.retype_password && cpwd.password != cpwd.retype_password">两次输入密码不一致</p>
        <el-button :disabled="!cpwd.password || !cpwd.retype_password || cpwd.retype_password != cpwd.password || cpwd.password.length < 6" type="primary" style="margin-left: 120px" @click="cpwdSubmit">确认修改</el-button>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
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
        ['DHU','东华大学']
      ]),
      loading: false,
      players: [],
      stat: {},
      form: {},
      cpwd: {},
      cpwdVisible: false,
    };
  },
  created: function() {
    this.status();
    if (this.$store.state.sid == 0) {
      this.loading = true;
      this.getList()
    }
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
    status() {
      axios.get('http://47.100.50.175:8088/api/player/status', {
      headers: { Authorization: this.$store.state.jwt }
      }).then(response => {
        this.stat = response.data
      })
    },
    find() {
      this.loading = true;
      axios.get('http://47.100.50.175:8088/api/public/search_player?id=' + this.form.game_id).then(response => {
        this.form.game_name = response.data.name
        this.loading = false;
      }).catch(() => {
        this.$message.error('未找到该玩家')
      })
    },
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
    submit() {
      axios.post('http://47.100.50.175:8088/api/player/bind', {
        "id": parseInt(this.form.game_id),
        "game_name": this.form.game_name,
        "nickname": this.form.nickname,
      }, { headers: { Authorization: this.$store.state.jwt } }).then(() => {
        this.$message.success('提交成功！')
        this.$router.push('/login')
      }).catch(() => {
        this.$message.error('未知错误')
      })
    },
    reset(sid) {
      this.$confirm('确定要重置该队员吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }).then(() => {
        axios.get('http://47.100.50.175:8088/api/admin/reset?sid=' + sid, {
          headers: { Authorization: this.$store.state.jwt }
        }).then(response => {
          this.$alert('用户名：' + response.data.username + '   密码：' + response.data.password, '重置完毕',
          {
            confirmButtonText: '确定'
          })
          this.getList();
        }).catch(() => {

        });
      });
    },
    cpwdSubmit() {
      axios.post('http://47.100.50.175:8088/api/player/change_pwd', {
        "new_pwd": this.cpwd.password,
      }, { headers: { Authorization: this.$store.state.jwt } }).then(() => {
        this.$message.success('更改成功！')
        this.cpwdVisible = false;
      }).catch(() => {
        this.$message.error('未知错误')
      })
    },
    getList() {
      axios.get('http://47.100.50.175:8088/api/admin/get', {
        headers: { Authorization: this.$store.state.jwt }
      }).then(response => {
        this.players = response.data;
        this.loading = false;
      }).catch(() => {

      });
    },
  }
};
</script>
