import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '仪表盘', icon: 'dashboard' }
    }]
  },
  {
    path: '/iaas',
    component: Layout,
    alwaysShow: true,
    name: 'Iaas',
    meta: { title: 'IAAS', icon: 'layer-group' },
    children: [{
      path: 'annotationTable',
      name: 'AnnotationTable',
      component: () => import('@/views/iaas/annotations/index'),
      meta: { title: '变更事件汇总', icon: 'bug-report', roles: ['ADMIN'] }
    }]
  },
  {
    path: '/manager',
    component: Layout,
    alwaysShow: true,
    name: 'Manager',
    meta: { title: 'API管理', icon: 'layer-group', roles: ['ADMIN'] },
    children: [{
      path: 'apiManagerTable',
      name: 'ApiManagerTable',
      component: () => import('@/views/apimanager/index'),
      meta: { title: '系统API列表', icon: 'bug-report' }
    }]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
        meta: { title: '外部链接', icon: 'international', roles: ['ADMIN'] }
      }
    ]
  },
  {
    path: '/ticket',
    component: Layout,
    alwaysShow: true,
    name: 'Ticket',
    meta: { title: '工单管理', icon: 'message' },
    children: [
      {
        path: 'ticket_list',
        name: 'TicketList',
        component: () => import('@/views/ticket/ticket-table/index'),
        meta: { title: '工单列表', icon: 'news' }
      },
      {
        path: 'ticket_msg',
        name: 'TicketMsg',
        hidden: true,
        component: () => import('@/views/ticket/ticket/msg'),
        meta: { title: '工单信息', icon: 'edit' }
      },
      {
        path: 'ticket_add',
        name: 'TicketAdd',
        component: () => import('@/views/ticket/ticket/create'),
        meta: { title: '创建工单', icon: 'edit' }
      },
      {
        path: 'ticket_edit',
        name: 'TicketEdit',
        hidden: true,
        component: () => import('@/views/ticket/ticket/edit'),
        meta: { title: '修改工单', icon: 'edit' }
      },
      {
        path: 'ticket_type',
        name: 'TicketType',
        component: () => import('@/views/ticket/ticket-config/index'),
        meta: { title: '工单属性', icon: 'resource' }
      }
    ]
  },
  {
    path: '/sys',
    component: Layout,
    alwaysShow: true,
    name: 'Sys',
    meta: { title: '系统管理', icon: 'cog' },
    children: [
      {
        path: 'user',
        name: 'User',
        component: () => import('@/views/account/user/index'),
        meta: { title: '用户管理', icon: 'people' }
      },
      {
        path: 'group',
        name: 'Group',
        component: () => import('@/views/account/group/index'),
        meta: { title: '用户组管理', icon: 'peoples' }
      },
      {
        path: 'perm',
        name: 'Perm',
        component: () => import('@/views/perm/index'),
        meta: { title: '权限管理', icon: 'password' }
      },
      {
        path: 'project',
        name: 'Project',
        component: () => import('@/views/project/index'),
        meta: { title: '项目管理', icon: 'step' }
      }
    ]
  },
    // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

export const asyncRoutes = [
]
const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
