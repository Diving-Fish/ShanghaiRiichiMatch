<template>
  <div class="box">
    <div class="hello">
      <h1>第三届魔都日麻高校个人赛选手及负责人须知</h1>
      <h2>0.	写在前面</h2>
      <p style="font-size: 24px">大会室ID：<a style="color: red">864823</a></p>
      <p>赛事网站：<a style="color: red">http://47.100.50.175:8080</a></p> 
      <p>因为服务器性能原因，第一次加载可能会稍慢。请使用PC或分辨率大于720p的横屏显示设备进行访问。</p>
      <h2>1.	概况</h2>
      <p>本次比赛共有<a style="color: red">26</a>所高校和1个赞助机构参加。</p> 
      <p>本次比赛使用<a style="color: red">雀魂平台</a>，请各位选手在参赛前先进行注册。</p> 
      <p>本次比赛赛程从<a style="color: red">10月25日~11月10日</a>，周六周日进行比赛。（预赛从周五开始）</p> 
      <h2>2.	报名</h2>
      <p>各高校队员联系高校负责人进行报名，截止时间为<a style="color: red">10月21日00:00。</a></p> 
      <p>负责人使用负责人账号（稍后将发送到各负责人的邮箱）登录后，进入网站管理界面——添加新队员，即可创建一个空的比赛账号。选手登录比赛账号后，进行信息填写即可正常参与比赛。请您注意修改初始密码，并保管好您的账号密码。</p> 
      <p><a style="color: red">每所高校最多能报名10人，组委会视情况会考虑增加名额</a>（参赛选手须是所在高校的在校生或是毕业生，本次比赛不限制毕业生数量），负责人账号无法参加比赛。若负责人想参加比赛，请自行创建一个新的比赛账号。</p> 
      <p>参赛学校及负责人名单：</p>
      <div style="display: flex">
        <div style="width: 500px">
          <el-table
            :data="schools.slice(0,13)"
            style="margin-left: 20px">
            <el-table-column
              prop="school"
              label="学校">
            </el-table-column>
            <el-table-column
              prop="nickname"
              label="负责人昵称">
            </el-table-column>
          </el-table>
        </div>
        <div style="margin-left: 20px; width: 500px">
          <el-table
            :data="schools.slice(14,27)">
            <el-table-column
              prop="school"
              label="学校">
            </el-table-column>
            <el-table-column
              prop="nickname"
              label="负责人昵称">
            </el-table-column>
          </el-table>
        </div>
      </div>
      <p>关于报名相关疑问请找本次主办：水鱼(QQ：2300756578)进行咨询。</p>
      <h2>3.	赛程及赛制</h2>
      <p>淘汰赛：</p>
      <p>X进128：10月25日19:00 ~ 10月26日18:00</p>
      <p>128进96：10月26日19:00 ~</p>
      <p>96进64：10月27日19:00 ~</p>
      <p>预赛：</p>
      <p>64进32：11月1日19:00 ~</p> 
      <p>32进16：11月2日19:00 ~</p>
      <p>正赛：</p>
      <p>16进8：11月8日19:00 ~</p> 
      <p>8进4：11月9日19:00 ~</p> 
      <p>决赛：11月10日19:00 ~</p> 
      <p>第一阶段于10月25日19:00 ~ 26日18:00进行，选手将在规定时间内自行在大会室进行预约，人满即开，<a style="color: red">每人只能打6个半庄</a>，所有选手通过6个半庄的<a style="color: red">马点之和</a>进行排名，前<a style="color: red">128</a>的选手晋级，未能打满6个半庄的选手视为弃权，淘汰。</p>
      <p>第二阶段于10月26日19:00 ~，10月27日19:00 ~，进行后台随机匹配对战，每一轮将在<a style="color: red">所有选手结束比赛后</a>进行下一轮随机，在进行<a style="color: red">3</a>次随即后决出前96/64进入预赛。</p>
      <p>预赛阶段：第一轮将为随机匹配对战，第二轮将根据第一轮各选手马点顺位排名进行下一轮比赛，依此进行<a style="color: red">4</a>轮比赛后通过<a style="color: red">马点之和</a>进行排名，最终前<a style="color: red">16</a>名选手晋级正赛。</p>
      <p>正赛阶段：比赛开始为<a style="color: red">预赛阶段马点顺位排名</a>进行第一轮比赛，<a style="color: red">所有选手结束比赛后</a>根据当前马点顺位进行第二轮比赛，依此进行<a style="color: red">4</a>轮比赛后，最终取前8进入半决赛，前4进入决赛。决赛通过<a style="color: red">5个半庄的马点之和</a>进行排名，决出冠亚季殿军。</p>
      <p>开始后进行所有晋级选手必须在比赛日<a style="color: red">18:00~18:30</a>之间进行签到，签到后将进行分桌。</p>
      <p>无法签到请提前告知自己的负责人，正赛开始后10分钟内未到将由神秘代打陪打。</p>
      <h2>4.  比赛规则</h2>
      <p>除无便捷提示外，规则和雀魂段位场相同。具体如下：</p>
      <p>原点：25000</p>
      <p>1位必要点数：30000</p>
      <p>返点：25000</p>
      <p>立直供托1000点，场棒300点，不听罚符累计3000点。</p>
      <p>顺位马：+15，+5，-5，-15。</p>
      <p>赤宝牌数量：3</p>
      <p><a style="color: red">无西入</a>，有一发，有杠宝牌、里宝牌，有击飞，有四风连打、四杠散了、四家立直、九种九牌流局，有流局满贯和多倍役满。</p>
      <p>无头跳，无三家和了流局。</p>
      <p>亲家和牌、听牌连庄；满足终局要求后亲家和牌、听牌终局。</p>
      <p>马点计算公式：(结算点数-25000)/1000+顺位马。例如选手A以33400点拿到2位，选手A本局的马点即为(33400-25000)/1000+5=+13.4。</p> 
      <h2>5.	其他</h2>
      <p>本次比赛严禁任何不正打行为，包括但不限于代打、通牌，一经发现，立即取消该选手比赛资格。</p> 
      <p>报名/比赛过程中若有任何疑问，可以联系各高校负责人，或联系主办<a style="color: red">水鱼（QQ:2300756578）</a>。</p>
      <p><a style="color: red">本须知最终解释权归日麻高校个人赛组委会所有。</a></p> 
    </div>
    <div class="sponsor">
      <h2>赞助名单</h2>
      <h3>个人:</h3>
      <p>安然(工具人，东京贫民，日代群：639283754)</p>
      <p>KY(某不知名上海高校团体赛创始人)</p>
      <p>死线姥爷(斑鸠雀庄:132276906)</p>
      <p>卡布(备战TI10，第二届个人赛冠军)</p>
      <p>漏漏(就是微博上的那个哇大风)</p>
      <p>GS(第一届个人赛创始人)</p>
      <p>July(据说是羊脱单前的最后一个徒弟)</p>
      <p>鹤丸国永（华政新一代顶梁柱）</p>
      <p>檀姐姐（超高校级经理人，外贸的女仆长，过气偶像）</p>
      <p><a href="https://live.bilibili.com/21555096/">璃亚</a>(VTuber，FreedomZone动漫协会看板娘，B站直播间: 21555096)</p>
      <p><a href="https://live.bilibili.com/918226/">空瞳画家</a>(音游椰叶，B站直播间：918226)</p>
      <p>......</p>
      <h3>团体:</h3>
      <p>カラス雀庄(约雀：766932393)</p>
      <p>Freedomzone动漫协会(上海交通大学)</p>
      <p>北天神高校(上交FZ家禽动物园分部)</p>
      <p>FZ日语开挂镇守府(上交FZ家禽动物园总部)</p>
      <p>天梦竞技麻将社(北京航空航天大学)</p>
      <p>SEU·东南大学雀魂(群号：980287102 )</p>
      <p><a href="https://paiyouzhijia.taobao.com">牌友之家淘宝店</a></p>
      <p>......</p>
      <h3>解说:</h3>
      <p><a href="https://live.bilibili.com/99349/">肖教练</a>(复旦精神支柱，打牌极限的求道派新晋凤凰，B站直播间：99349)</p>
      <p>......</p>
      <p>以及海选阶段会被淘汰的N名谐星。</p>
      <h2>奖品及奖池列表</h2>
      <p>冠军：订制奖杯+雀魂月卡六个月(可折现150)+奖池自选一 + カラス雀庄30h白打券。(可授权)</p>
      <p>亚军：订制奖杯+雀魂月卡五个月(可折现120)+奖池自选一 + カラス雀庄20h白打券。(可授权)</p>
      <p>季军：订制奖杯+雀魂月卡四个月(可折现100)+奖池自选一 + カラス雀庄10h白打券。(可授权)</p>
      <p>殿军：雀魂月卡三个月(可折现80)+奖池自选一</p>
      <p>5-8：雀魂月卡两个月(可折现50)+奖池自选一</p>
      <p>9-12：雀魂月卡一个月(可折现25)+奖池自选一</p>
      <p>13-16：雀魂月卡一个月(可折现25)</p>
      <h3>奖池列表</h3>
      <p>1. 大洋化学 AMOS MAX 麻将+台布+手册 X 1</p>
      <p>2. 大洋化学 AMOS GABIN 麻将 X 1</p>
      <p>3. 大洋化学 麻将桌布 台布 桌垫 JUNKMAT X 1</p>
      <p>4. 大富翁10 steam游戏 X 1</p>
      <p>5. 剧毒少女 秋季例大祭第六回新刊 X 1</p>
      <p>6. 瓦西子透明麻将 X 2</p>
      <p>7. 雀魂官方吧唧全套 X 2</p>
      <p>8. 天麻小挂画 X 4</p>
      <p>9. 麻将博物馆定制 麻将扑克牌 X N</p>
      <p>10. 大洋化学 白彩色筹码 彩色点棒 X N</p>
      <p>11. BW 徽章set X 1</p>
      <p>12. 麦当劳 200 元礼品卡 + melody 玩具一套 X 1</p>
      <p style="margin-top: 30px;">有赞助想法请联系<a style="color: red">安然（QQ:513696293)</a>。</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      schools: [
        {school: "华东师范大学", nickname: "czh", qq: "1158915441"},
        {school: "上海外国语大学", nickname: "二戸鈴", qq: "785137272"},
        {school: "复旦大学", nickname: "taku", qq: "374095644"},
        {school: "上海大学", nickname: "呆轩", qq: "939918608"},
        {school: "上海海事大学", nickname: "山风", qq: "675946979"},
        {school: "上海对外经贸大学", nickname: "檀檀", qq: "1297244236"},
        {school: "上海应用技术大学", nickname: "勇士", qq: "1211445112"},
        {school: "上海工程技术大学", nickname: "南浦", qq: "396335709"},
        {school: "上海财经大学", nickname: "飞啊飞", qq: "1318250776"},
        {school: "上海理工大学", nickname: "鸡腿", qq: "982919014"},
        {school: "浙江大学", nickname: "Sapphire ", qq: "1035627110"},
        {school: "上海政法学院", nickname: "T妹", qq: "844768831"},
        {school: "学院路联合大学", nickname: "小弓"},
        {school: "斑鸠雀庄", nickname: "死线"},
        {school: "华东政法大学", nickname: "清风沐璃", qq: "2258349120"},
        {school: "上海立信会计金融学院", nickname: "哲学", qq: "1005536376"},
        {school: "上海电力大学", nickname: "qb", qq: "1499227861"},
        {school: "南京大学", nickname: "简单", qq: "514405488"},
        {school: "华东理工大学", nickname: "云", qq: "491445219"},
        {school: "上海海洋大学", nickname: "MrCrystalJoker", qq: "2267926518"},
        {school: "上海师范大学", nickname: "未来", qq: "974020665"},
        {school: "上海交通大学", nickname: "水鱼", qq: "2300756578"},
        {school: "东华大学", nickname: "梦境无恒", qq: "862919793"},
        {school: "东南大学", nickname:"Mew's"},
        {school: "同济大学", nickname:"中张"},
        {school: "上海科技大学", nickname:"叉子"},
        {school: "苏州大学", nickname:"蓝白"}
      ]
    }
  },
  name: "HelloWorld",
  props: {
    msg: String
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.box {
  display: flex;
}

.hello {
  padding: 0px 30px;
}

.sponsor {
  width: 35%;
  padding: 30px;
  border-left: solid 1px #e6e6e6;
}
p {
  margin-left: 20px;
}
h1 {
  font-size: 2em;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
