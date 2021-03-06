const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  nick_name: state => state.user.nick_name,
  roles: state => state.user.roles,
  perms: state => state.user.perms,
  addRouters: state => state.permission.addRouters,
  permission_routers: state => state.permission.routers
}
export default getters
