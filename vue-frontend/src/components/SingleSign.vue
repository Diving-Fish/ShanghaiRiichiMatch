<template>
  <div>
    <el-page-header @back="$router.push('/')" content="个人报名"></el-page-header>
    <el-button style="margin: 20px 0px 0px 20px" @click="$router.push('/sign')">回到队伍报名</el-button>
    <el-container v-loading="loading">
      <el-form ref="form" :model="form" label-width="120px" style="width: 100%; margin: 40px 40% 0px 20%">
        <el-form-item label="雀魂ID">
          <el-input v-model="form.id"></el-input>
        </el-form-item>
        <el-form-item label="QQ号">
          <el-input v-model="form.qq"></el-input>
        </el-form-item>
        <el-form-item>
            <el-checkbox v-model="form.checked">我已知悉个人报名会由队伍自动分配队友，且无法添加替补</el-checkbox>
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
/* eslint-disable */
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
      if (!this.form.id || !this.form.qq ) {
        this.$message.error("字段不能为空")
        return
      }
      if (!this.form.checked) {
          this.$message.error("请勾选须知")
          return
      }
      let that = this
      that.loading = true
      axios.post('http://47.100.50.175:8088/api/public/single_sign', {
        "qq": parseInt(that.form.qq),
        "id": parseInt(that.form.id)
      }).then(response => {
          if (response.data.reg) {
              that.$message.error("看起来这个玩家已经报过名了呢…")
              that.loading = false
          } else {
              that.$message.success("报名成功！")
              that.$router.push('/')
          }
      }).catch(error => {
          that.$message.error("没能找到ID呢……是不是输错了？")
          that.loading = false
      })
    }
  }
};
</script>