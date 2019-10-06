<template>
  <div id="app">
    <title>联机杯雀魂杯赛</title>
    <el-container>
      <el-header class="el-header">
        <a>联机杯雀魂杯赛</a>
        <ul style="float: right; margin: 0px; font-size: 16px">
          <el-button round icon="el-icon-setting" @click="admin">超级管理员</el-button>
        </ul>
      </el-header>
      <el-container>
        <el-aside>
          <el-menu :default-openeds="['1']">
            <el-submenu index="1">
              <template slot="title"><i class="el-icon-menu"></i><a style="font-size:16px">比赛信息</a></template>
              <el-menu-item index="1-1" @click="$router.push('/')"><a>概况</a></el-menu-item>
              <el-menu-item index="1-2" @click="$router.push('/group')"><a>分组</a></el-menu-item>
              <el-menu-item index="1-3" @click="$router.push('/stat')"><a>统计</a></el-menu-item>
            </el-submenu>
            <el-menu-item index="2" @click="$router.push('/login')"><i class="el-icon-s-tools"></i><a style="font-size:16px">队伍管理</a></el-menu-item>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
//import HelloWorld from "./components/HelloWorld.vue";
import axios from 'axios'

export default {
  name: "app",
  components: {
  },
  methods: {
    admin() {
      let that = this
      axios.get('http://47.100.50.175:8088/api/admin/stat', {
        headers: { Authorization: that.$store.state.admin_key }
      }).then(() => {
        that.$router.push('/admin')
      }).catch(() => {
        that.$prompt('请输入管理员密码').then(({ value }) => {
          that.$store.commit('set_admin_key', value)
          axios.get('http://47.100.50.175:8088/api/admin/stat', {
            headers: { Authorization: that.$store.state.admin_key }
          }).then(() => {
            that.$message.success('登录成功')
            that.$router.push('/admin')
          }).catch(() => {
            that.$message.error('密码错误')
          })
        })
      })
    }
  }
};
</script>

<style>
.el-header {
  font-size: 24px;
  background-color: #409eff;
  color: #fff;
  line-height: 60px;
}
</style>
