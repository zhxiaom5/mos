<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入annotations搜索"/>
        </el-form-item>
        <el-form-item>
          <span class="el-tag el-tag--info pointer" @click="handleSearch()">
            <i class="el-icon-search"></i>
            <a class="a_underline">搜索</a>
          </span>
        </el-form-item>
        <el-form-item>
          <span class="el-tag pointer" @click="handleAdd()">
            <i class="el-icon-plus"></i>
            <a class="a_underline">添加</a>
          </span>
        </el-form-item>
      </el-form>
      <el-table
        :data="tableData.filter(data => !listQuery.search || data.group.toLowerCase().includes(listQuery.search.toLowerCase()))"
        :header-cell-style="tableHeaderColor"
        stripe
        border
        style="width: 100%">
        <el-table-column
          type="selection"
          width="55" />
        <el-table-column
          label="RegionID"
          prop="region_id"
          width="150" />
        <el-table-column
          label="所属项目名(project)"
          prop="project"
          width="200" />
        <el-table-column
          label="内容(text)"
          prop="text" />
        <el-table-column
          label="标签(tag)"
          prop="tag">
          <template slot-scope="scope">
            <span v-for="tags in scope.row.tag" :key="tags+scope.$index" class="el-tag el-tag--success">{{ tags }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="标签时间"
          prop="start_time"
          width="200"
        />
        <el-table-column
          align="center"
          label="操作"
          width="300"
        >
          <template slot-scope="scope">
            <span class="el-tag el-tag--danger pointer" @click="handleDelete(scope.$index, scope.row)">
              <i class="el-icon-delete"></i>
              <a class="a_underline">删除</a>
            </span>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getAnnonList" />
      <el-dialog :visible.sync="dialogAnnonFormVisible" title="新增/修改备注">
        <el-form ref="annonForm" :model="annonForm" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="annonForm.id" disabled />
          </el-form-item>
          <el-form-item label="项目">
            <el-input v-model="annonForm.project" :disabled="isEditable" />
          </el-form-item>
          <el-form-item label="内容">
            <el-input v-model="annonForm.text" />
          </el-form-item>
          <el-form-item label="标签">
            <el-input v-model="annonForm.tags" placeholder="多标签请用','隔开" />
          </el-form-item>
          <el-form-item label="类型">
            <el-switch
              v-model="annonForm.isRegion"
              active-color="green"
              inactive-color="red"
              inactive-text="时间点"
              active-text="时间范围"
            />
          </el-form-item>
          <el-form-item label="时间">
            <el-date-picker
              v-model="timePicker.startTime"
              type="datetime"
              size="small"
              placeholder="开始时间"
              @change="dateTimeChange('start')"
            />
            <el-date-picker
              v-if="isEndTimeVisbale"
              v-model="timePicker.endTime"
              type="datetime"
              size="small"
              placeholder="结束时间"
              @change="dateTimeChange('end')"
            />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
            <el-button @click="dialogAnnonFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'
import { tableHeaderColor } from '../../../utils/style.js'
import { getAnnonList, addAnnon, deleteAnnon, Tst } from '@/api/annotations'
export default {
  name: 'AnnotationTable',
  components: { Pagination },
  data() {
    return {
      timePicker: {
        startTime: undefined,
        endTime: undefined
      },
      tableData: [],
      total: 0,
      listQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      isEndTimeVisbale: true,
      isEditable: false,
      // 新增，更新状态
      dialogStatus: '',
      // 新增dialog是否显示
      dialogAnnonFormVisible: false,
      annonForm: {
        id: undefined,
        isRegion: false,
        project: '',
        text: '',
        tags: '',
        startTime: undefined,
        endTime: undefined
      }
    }
  },
  watch: {
    'annonForm.isRegion'(newval, oldval) {
      if (newval === true) {
        this.isEndTimeVisbale = true
      } else {
        this.isEndTimeVisbale = false
      }
    }
  },
  created() {
    this.getAnnonList()
  },
  methods: {
    tableHeaderColor,
    resetTemp() {
      this.annonForm = {
        id: undefined,
        isRegion: false,
        project: '',
        text: '',
        tags: '',
        startTime: undefined,
        endTime: undefined
      }
    },
    dateTimeChange(val) {
      const This = this
      if (val === 'start') {
        This.$set(This.annonForm, 'startTime', Date.parse(new Date(This.timePicker.startTime)) / 1000)
      } else if (val === 'end') {
        This.$set(This.annonForm, 'endTime', Date.parse(new Date(This.timePicker.endTime)) / 10000)
      } else {
        console.log('selected')
      }
    },
    getAnnonList() {
      getAnnonList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getAnnonList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      const This = this
      This.resetTemp()
      This.dialogStatus = 'create'
      This.isEditable = false
      if (This.annonForm.isRegion === false) {
        This.isEndTimeVisbale = false
      }
      This.dialogAnnonFormVisible = true
    },
    handleEdit(index, row) {
    //   this.group_form = Object.assign({}, row)
    //   this.dialogStatus = 'update'
    //   this.isGroupNameEditAble = true
    //   getAccountMsg().then(response => {
    //     this.userArray = response.data
    //   })
    //   getGroupMsg(this.group_form.id).then(response => {
    //     this.$set(this.group_form, 'user_selected', response.data)
    //     // console.log(this.group_form.userSelected)
    //   })
    //   // console.log(index, row)
    //   this.dialogGroupFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      This.$confirm('是否删除?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // console.log(row.id)
        deleteAnnon(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
            onClose: function refresh() {
              This.getAnnonList()
            }
          })
        }).catch(() => {
          This.$message({
            type: 'info',
            message: '已取消删除'
          })
          // console.log(row)
        })
      })
    },
    createData() {
      const This = this
      // console.log(This.annonForm)
      addAnnon(This.annonForm).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 1000,
          onClose: function refresh() {
            This.dialogAnnonFormVisible = false
            This.getAnnonList()
          }
        })
      })
    },
    updateData() {
    //   const This = this
    //   updateGroup(this.group_form).then(response => {
    //     this.$message({
    //       message: response.message,
    //       type: 'success',
    //       duration: 2000,
    //       onClose: function refresh() {
    //         This.dialogGroupFormVisible = false
    //         This.getGroupList()
    //       }
    //     })
    //   })
    },
    test() {
      Tst()
    }
  }
}
</script>

