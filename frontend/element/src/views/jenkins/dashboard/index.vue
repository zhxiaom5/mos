<template>
  <div class="dashboard-editor-container">
    <!-- <panel-group
      :my_month_ticket="ticket_status.my_month_ticket"
      :my_unfinish_ticket="ticket_status.my_unfinish_ticket"
      :group_unfinish_ticket="ticket_status.group_unfinish_ticket"
      :all_month_ticket="ticket_status.all_month_ticket" /> -->
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
      <line-chart :chart-data="lineChartData" :x-tag="x_tag"/>
    </el-row>
  </div>
</template>

<script>

// import PanelGroup from './components/PanelGroup'
import { getJenkinsLineData } from '@/api/jenkins'
import LineChart from '@/components/Charts/LineChart'
// import { mapGetters } from 'vuex'
// const lineChartData = {
//   newVisitis: {
//     expectedData: [100, 120, 161, 134, 105, 160, 165],
//     actualData: [120, 82, 91, 154, 162, 140, 145]
//   }
// }
export default {
  name: 'JenkinsDashboard',
  components: {
    //PanelGroup,
    LineChart
  },
  data() {
    return {
      x_tag: [],
      lineChartData: {}
    //   ticket_status: {
    //     my_month_ticket: 0,
    //     my_unfinish_ticket: 0,
    //     group_unfinish_ticket: 0,
    //     all_month_ticket: 0
    //  }
    }
  },
//   computed: {
//     ...mapGetters([
//       'name',
//       'roles'
//     ])
//   },
  created() {
    this.getDashData()
  },
  methods: {
    getDashData() {
    //   getTicketDashData().then(response => {
    //     this.$set(this.ticket_status, 'my_month_ticket', response.data.my_month_ticket)
    //     this.$set(this.ticket_status, 'my_unfinish_ticket', response.data.my_unfinish_ticket)
    //     this.$set(this.ticket_status, 'group_unfinish_ticket', response.data.group_unfinish_ticket)
    //     this.$set(this.ticket_status, 'all_month_ticket', response.data.all_month_ticket)
    //   })
      getJenkinsLineData().then(response => {
        this.lineChartData = response.data
        this.x_tag = response.x_tag
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}
</style>
