<template>
  <div class="app-container">
    <div class="filter-container">
      <el-tabs style="margin-top:15px;">
        <!-- Annotations -->
        <el-tab-pane>
          <span slot="label"><i class="el-icon-date"></i> Annontation</span>
          <el-form :inline="true" class="demo-form-inline">
            <el-form-item>
              <el-input v-model="listQuery.search" size="small" style="" placeholder="输入annotations搜索"/>
            </el-form-item>
            <el-form-item>
              <span class="el-tag el-tag--info pointer" @click="getAnnonList()">
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
            :data="tableData.filter(data => !listQuery.search || data.text.toLowerCase().includes(listQuery.search.toLowerCase()))"
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
        </el-tab-pane>
        <!-- Jenkins -->
        <el-tab-pane>
          <span slot="label"><i class="el-icon-star-off"></i> Jenkins</span>
          <el-form :inline="true" class="demo-form-inline">
            <el-form-item>
              <el-input v-model="listQueryJenkins.search" size="small" style="" placeholder="输入Jenkins搜索"/>
            </el-form-item>
            <el-form-item>
              <span class="el-tag el-tag--info pointer" @click="getJenkinsList()">
                <i class="el-icon-search"></i>
                <a class="a_underline">搜索</a>
              </span>
            </el-form-item>
          </el-form>
          <el-table
            :data="tableJenkinsData.filter(data => !listQueryJenkins.search || data.project.toLowerCase().includes(listQueryJenkins.search.toLowerCase()))"
            :header-cell-style="tableHeaderColor"
            stripe
            border
            style="width: 100%">
            <el-table-column
              type="selection"
              width="55" />
            <el-table-column
              label="项目"
              prop="project"
              width="150" />
            <el-table-column
              label="标题"
              prop="title"
              width="350" />
            <el-table-column
              label="构建者"
              prop="build_user" />
            <el-table-column
              label="状态"
              prop="build_status" />  
            <el-table-column
              label="主机"
              prop="hosts" />
            <el-table-column
              label="标签时间"
              prop="create_time"
              width="200" />
            <el-table-column
              align="center"
              label="操作"
              width="300">
              <template>
                <span class="el-tag el-tag--danger pointer" @click="handleJenkinsDelete()">
                  <i class="el-icon-delete"></i>
                  <a class="a_underline">删除</a>
                </span>
              </template>
            </el-table-column>
          </el-table>
          <pagination v-show="total>0" :total="totalJenkins" :current_page.sync="listQueryJenkins.current_page" :page_size.sync="listQueryJenkins.page_size" align="right" @pagination="getJenkinsList" />
        </el-tab-pane>
      </el-tabs>  
    </div>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'
import { tableHeaderColor } from '../../../utils/style.js'
import { getAnnonList, addAnnon, deleteAnnon, getJenkinsList } from '@/api/annotations'
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
      tableJenkinsData: [],
      totalJenkins: 0,
      total: 0,
      listQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      listQueryJenkins: {
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
    this.getJenkinsList()
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
    getJenkinsList() {
      getJenkinsList(this.listQueryJenkins).then(response => {
        this.tableJenkinsData = response.data
        this.totalJenkins = response.total
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
    handleJenkinsDelete() {
      this.$message({
        message: "禁止删除",
        type: 'error',
        duration: 1000
      })
    }
  }
}
</script>

