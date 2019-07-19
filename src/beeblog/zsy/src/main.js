import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false

/* eslint-disable no-new */

router.beforeEach(function(to, from, next){
  let needLogin=to.matched.some(record => record.meta.needLogin);
  let token=window.sessionStorage.token;
  if (needLogin){
    if(!token){
      next({path:'/'})
    }else{
      next()
    }
  }else{
    next()
  }
});

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
