<template>
  <div class="dashboard-editor-container">
    <panel-group
      :myMonthTicket="dataStatus.myMonthTicket"
      :myUnfinishTicket="dataStatus.myUnfinishTicket"
      :groupUnfinishTicket="dataStatus.groupUnfinishTicket"
      :allMonthTicket="dataStatus.allMonthTicket" />
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
      <ticket-line-chart :chart-data="lineChartData" :x-tag="x_tag"/>
    </el-row>
  </div>
</template>

<script>

import PanelGroup from './components/PanelGroup'
import { getTicketLineData } from '@/api/ticket'
import TicketLineChart from './components/LineChart'
import { mapGetters } from 'vuex'

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    TicketLineChart
  },
  data() {
    return {
      x_tag: [],
      lineChartData: {},
      dataStatus: {
        myMonthTicket: 0,
        myUnfinishTicket: 0,
        groupUnfinishTicket: 0,
        allMonthTicket: 0
      }
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'roles'
    ])
  },
  created() {
    this.getTicketDashStatus()
  },
  methods: {
    getTicketDashStatus() {
      // getTicketDashData().then(response => {
      //   this.$set(this.dataStatus, 'monthTicket', response.data.monthTicket)
      //   this.$set(this.dataStatus, 'MyUnfinishTicket', response.data.MyUnfinishTicket)
      //   this.$set(this.dataStatus, 'GroupUnfinishTicket', response.data.GroupUnfinishTicket)
      //   this.$set(this.dataStatus, 'allMonthTicket', response.data.allMonthTicket)
      // })
      getTicketLineData().then(response => {
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
