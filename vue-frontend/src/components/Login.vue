<template>
  <div>
    <el-page-header @back="$router.push('/')" content="登录"></el-page-header>
    <el-container v-loading="loading">
      <el-form
        ref="form"
        :model="form"
        label-width="120px"
        style="width: 100%; margin: 40px 40% 0px 20%"
      >
        <el-form-item label="用户名">
          <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input type="password" v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">登录</el-button>
        </el-form-item>
      </el-form>
    </el-container>
  </div>
</template>

<script>
/* eslint-disable */
import axios from "axios";
export default {
  data() {
    return {
      form: {},
      loading: false
    };
  },
  methods: {
    onSubmit() {
      axios.post('http://47.100.50.175:8088/api/public/login', {
          "username": this.form.username,
          "password": this.form.password
      }).then(response => {
          this.$store.commit('set_jwt', response.data.token)
          this.$message.success('登录成功，正在为您跳转...')
          this.$router.push('/manage')
          this.$router.push('/login')
      }).catch(() => {
          this.$message.error('ID或密码错误')
      })
    }
  }
};
</script>