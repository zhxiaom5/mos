<template>
  <div class="dashboard-editor-container">
    <panel-group
      :unfinishTotal="dataStatus.unfinishTotal"
      :ticketTotal="dataStatus.ticketTotal"
      :myUnfinish="dataStatus.myUnfinish"
      :myTotal="dataStatus.myTotal" />
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
        unfinishTotal: 0,
        ticketTotal: 0,
        myUnfinish: 0,
        myTotal: 0
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
        this.$set(this.dataStatus, 'unfinishTotal', response.datadash.unfinishTotal)
        this.$set(this.dataStatus, 'ticketTotal', response.datadash.ticketTotal)
        this.$set(this.dataStatus, 'myUnfinish', response.datadash.myUnfinish)
        this.$set(this.dataStatus, 'myTotal', response.datadash.myTotal)
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
