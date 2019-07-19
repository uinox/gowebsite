<template>
  <div id="welcome">
    <h1 style="padding-top:25%">{{ msg }}</h1>
    <router-link id="btn" to="/home">Enter</router-link>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome!'
    }
  },
  mounted:function(){
    this.background();
  },
  methods:{
    background(){
      var screenH = window.screen.availHeight;
      var screenW = window.screen.availWidth;
      document.body.style.height = screenH-100 + "px";
      document.getElementById('welcome').style.height= screenH-100 + "px";
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
#welcome{height: 100%;height: 100%;}
#btn{
  width: 120px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  border: 1px solid #aa4705;
  cursor: pointer;
  display: block;
  margin: 40px auto;
  border-radius: 10px;
  font-size: 20px;
  color: #f7b098da;
  font-weight: bold;
  background: url(../assets/btn.jpg) no-repeat center;
  text-decoration: none;
}
</style>
