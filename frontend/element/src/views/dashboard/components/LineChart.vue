<template>
  <div :class="className" :style="{height:height,width:width}"/>
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
export default {
  name: 'TicketLineChart',
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '500px'
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    chartData: {
      type: Object,
      required: true
    },
    xTag: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      chart: null,
      sidebarElm: null,
      line_color: ''

    }
  },
  watch: {
    chartData: {
      deep: true,
      handler(val) {
        this.setOptions(val)
      }
    }
  },
  mounted() {
    this.initChart()
    if (this.autoResize) {
      // this.__resizeHandler = debounce(() => {
      //   if (this.chart) {
      //     this.chart.resize()
      //   }
      // }, 100)
      window.addEventListener('resize', this.resizeHandler)
    }
    // 监听侧边栏的变化
    this.sidebarElm = document.getElementsByClassName('sidebar-container')[0]
    // console.log(this.sidebarElm)
    this.sidebarElm && this.sidebarElm.addEventListener('transitionend', this.sidebarResizeHandler)
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    if (this.autoResize) {
      window.removeEventListener('resize', this.resizeHandler)
    }
    this.sidebarElm && this.sidebarElm.removeEventListener('transitionend', this.sidebarResizeHandler)
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    RandomColor() {
      const r = Math.floor(Math.random() * 256)
      const g = Math.floor(Math.random() * 256)
      const b = Math.floor(Math.random() * 256)
      const col = '#' + r.toString(16) + g.toString(16) + b.toString(16)
      return col
    },
    resizeHandler() {
      if (this.chart) {
        this.chart.resize()
      }
    },
    sidebarResizeHandler(e) {
      if (e.propertyName === 'width') {
        // console.log('resize siderbar')
        this.__resizeHandler()
      }
    },
    setOptions(assData) {
      const This = this
      var series_array = []
      var dataLegend = []
      for (var key in assData) {
        // console.log(key, assData[key])
        var line_color = '#' + Math.floor(Math.random() * 256).toString(16) + Math.floor(Math.random() * 256).toString(16) + Math.floor(Math.random() * 256).toString(16)
        var tmp = {
          name: key, itemStyle: {
            normal: {
              color: line_color,
              lineStyle: {
                color: line_color,
                width: 2
              }
            }
          },
          smooth: true,
          type: 'line',
          data: assData[key],
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        }
        // console.log(line_color)
        dataLegend.push(key)
        series_array.push(tmp)
      }
      // console.log(series_array)
      this.chart.setOption({
        xAxis: {
          data: This.xTag,
          boundaryGap: false,
          axisTick: {
            show: false
          },
          axisLabel: {
            interval: 'auto'
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: {
          axisTick: {
            show: false
          }
        },
        legend: {
          data: dataLegend
        },
        series: series_array,
        title : {
          text: '过去30日工单情况'
        },
        toolbox: {
          show : true,
          feature : {
            mark : {show: true},
            dataView : {show: true, readOnly: false},
            magicType : {show: true, type: ['line', 'bar']},
            restore : {show: true},
            saveAsImage : {show: true}
          }
        },
      })
    },
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      this.setOptions(this.chartData)
    }
  }
}
</script>
