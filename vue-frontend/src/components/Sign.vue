<template>
  <div>
    <el-page-header @back="$router.push('/')" content="队伍报名"></el-page-header>
    <el-button style="margin: 20px 0px 0px 20px" @click="$router.push('/singlesign')">去个人报名</el-button>
    <el-container v-loading="loading">
      <el-form ref="form" :model="form" label-width="120px" style="width: 100%; margin: 40px 40% 0px 20%">
        <el-form-item label="队伍名称">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="队长雀魂ID">
          <el-input v-model="form.leader_id"></el-input>
        </el-form-item>
        <el-form-item label="队长QQ号">
          <el-input v-model="form.leader_qq"></el-input>
        </el-form-item>
        <el-form-item label="队伍密码">
          <el-input type="password" v-model="form.password"></el-input>
        </el-form-item>
         <el-form-item label="重复队伍密码">
          <el-input type="password" v-model="form.retype_password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">立即报名</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-container>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      form: {

      },
      loading: false
    };
  },
  methods: {
    onSubmit() {
      if (!this.form.name || !this.form.leader_id || !this.form.leader_qq || !this.form.password) {
        this.$message.error("字段不能为空")
        return
      }
      if (this.form.password !== this.form.retype_password) {
        this.$message.error("两次输入密码不一致")
        return
      }
      let that = this
      that.loading = true
      axios.post('http://47.100.50.175:8088/api/public/create_team', {
        "name": that.form.name,
        "qq": parseInt(that.form.leader_qq),
        "password": that.form.password,
        "id": parseInt(that.form.leader_id)
      }).then(response => {
        const data = response.data
        if (data.msg === 'duplicate') {
          that.$message.error("队伍名重复了，换一个吧")
          that.loading = false
          return
        } else if (data.msg === "can't find player") {
          that.$message.error("没能找到ID呢……是不是输错了？")
          that.loading = false
          return
        } else if (data.reg) {
          that.$message.error("看起来这个玩家已经报过名了呢…")
          that.loading = false
          return
        }
        that.loading = false
        that.$alert('请牢记管理ID及密码：' + data.team_id + '/' + that.form.password, '您的队伍已报名！', {
          confirmButtonText: '确定',
          callback: () => {
            that.$router.push('/manage')
          }
        });
      }).catch(() => {
        that.$message.error("服务器好像暂时出问题了…")
      })
    }
  }
};
</script>