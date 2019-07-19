import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/welcome'
import Home from '@/page/home/home'
import Category from '@/page/category/category'
import Article from '@/page/article/article'
import Article_add from '@/page/article/children/article_add'
import Article_edit from '@/page/article/children/article_edit'
import Article_view from '@/page/article/children/article_view'
import Tech from '@/page/tech/tech'
import Login from '@/page/login/login'
import Register from '@/page/login/register'
import Register1 from '@/page/login/register1'

Vue.use(Router)

export default new Router({
  mode:'history',
  routes: [
    // {
    //   path: '/',
    //   name: 'home',
    //   component: Home,
    //   meta:{
    //     needLogin: false
    //   }
    // },
    {
      path: '/home',
      name: 'home',
      component: Home,
      meta: {
        needLogin: false
      }
    },
    {
      path: '/',
      name: 'login',
      component: Login,
      meta:{
        needLogin: false
      }
    },
    {
      path: '/category',
      name: 'category',
      component: Category,
      meta: {
        needLogin: false
      }
    },
    {
      path: '/article',
      name: 'article',
      component: Article,
      meta: {
        needLogin: false
      }
    },
    {
      path: '/article/article_add',
      name: 'article_add',
      component: Article_add
    },
    {
      path: '/article/article_edit',
      name: 'article_edit',
      component: Article_edit
    },
    {
      path: '/article/article_view',
      name: 'article_view',
      component: Article_view
    },
    {
      path: '/tech',
      name: 'tech',
      component: Tech,
      meta: {
        needLogin: false
      }
    },
    {
      path: '/register',
      name: 'Register',
      component: Register,
      meta: {
        needLogin: false
      }
    },
    {
      path: '/register1',
      name: 'Register1',
      component: Register1
    }
  ],
  // linkExactActiveClass:'isActive', //精准匹配
  linkActiveClass:'isActive'  //模糊匹配
})
