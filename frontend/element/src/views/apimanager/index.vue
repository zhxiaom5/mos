<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入api模糊搜索"/>
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
        :data="tableData.filter(data => !listQuery.search || data.api_uri.toLowerCase().includes(listQuery.search.toLowerCase()))"
        :header-cell-style="tableHeaderColor"
        stripe
        border
        style="width: 100%">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-row v-for="it in props.row.api_key_info" :key="it.id+it.user">
              <el-col :span="8">
                <h4>用户名:{{ it.user }}</h4>
              </el-col>
              <el-col :span="14">
                <h4>密钥:{{ it.auth_key }}</h4>
              </el-col>
              <el-col :span="2">
                <span class="el-tag el-tag--danger pointer" @click="handleKeyDelete(it.id)">
                  <i class="el-icon-delete"></i>
                  <a class="a_underline">删除</a>
                </span>
              </el-col>
            </el-row>      
          </template>
        </el-table-column>
        <el-table-column
          label="API"
          prop="api_uri" />
        <el-table-column
          label="备注"
          prop="content" />
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
            <span class="el-tag pointer" @click="handleAddKey(scope.$index, scope.row)">
              <i class="el-icon-share"></i>
              <a class="a_underline">添加密钥</a>
            </span>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getApiList" />
      <!-- 添加API dialog -->
      <el-dialog :visible.sync="dialogApiFormVisible" title="新增/修改API">
        <el-form ref="apiForm" :model="apiForm" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="apiForm.id" disabled />
          </el-form-item>
          <el-form-item label="API_URI">
            <el-input v-model="apiForm.uri" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="apiForm.content" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="createData()">确定</el-button>
            <el-button @click="dialogApiFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <!-- 添加Key dialog -->
      <el-dialog :visible.sync="dialogKeyFormVisible" title="新增用户KEY">
        <el-form ref="keyForm" :model="keyForm" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="keyForm.apiID" disabled/>
          </el-form-item>
          <el-form-item label="API">
            <el-input v-model="keyForm.apiURI" disabled/>
          </el-form-item>
          <el-form-item label="用户">
            <el-input v-model="keyForm.user" />
          </el-form-item>
          <el-form-item>
            <el-switch
              v-model="isKeyDisable"
              active-text="启用"
              active-color="#13ce66" />
            </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="createKeyData()">确定</el-button>
            <el-button @click="dialogKeyFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'
import { tableHeaderColor } from '../../utils/style.js'
import { getApiList, deleteApi, addApi, addApiKey, deleteApiKey } from '@/api/apimanager'
export default {
  name: 'ApiManagerTable',
  components: { Pagination },
  data() {
    return {
      tableData: [],
      total: 0,
      listQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      isKeyDisable: false,
      // 新增dialog是否显示
      dialogApiFormVisible: false,
      // 添加密钥dialog是否显示
      dialogKeyFormVisible: false,
      apiForm: {
        id: undefined,
        content: "",
        uri: ''
      },
      keyForm: {
        apiID: undefined,
        apiURI: "",
        isKeyDisable: false,
        user: ""
      }
    }
  },
  watch: {
    // 监听密钥启用switch切换
    isKeyDisable(newval, oldval) {
      const This = this
      if (newval === true) {
        This.$set(This.keyForm, 'isKeyDisable', true)
      } else {
        This.$set(This.keyForm, 'isKeyDisable', false)
      }
    }
  },
  created() {
    this.getApiList()
  },
  methods: {
    tableHeaderColor,
    resetTemp() {
      this.apiForm = {
        id: undefined,
        content: "",
        uri: ''
      },
      this.keyForm = {
        apiID: undefined,
        apiURI: "",
        isKeyDisable: false,
        user: ""
      }
    },
    getApiList() {
      getApiList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getApiList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      const This = this
      This.resetTemp()
      This.dialogApiFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      This.$confirm('是否删除?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteApi(row.id).then(response => {
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
        })
      })
    },
    // 删除接口KEY
    handleKeyDelete(id) {
      const This = this
      This.$confirm('是否删除?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteApiKey(id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
            onClose: function refresh() {
              This.getApiList()
            }
          })
        }).catch(() => {
          This.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      })
    },
    handleAddKey(index, row) {
      const This = this
      This.resetTemp()
      This.$set(This.keyForm, 'apiID', row.id)
      This.$set(This.keyForm, 'apiURI', row.api_uri)
      This.dialogKeyFormVisible = true
    },
    createData() {
      const This = this
      addApi(This.apiForm).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 1000,
          onClose: function refresh() {
            This.dialogApiFormVisible = false
            This.getApiList()
          }
        })
      })
    },
    createKeyData() {
      const This = this
      addApiKey(This.keyForm).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 1000,
          onClose: function refresh() {
            This.dialogApiFormVisible = false
            This.getApiList()
          }
        })
      })
    }
  }
}
</script>


