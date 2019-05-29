// 修改table header的背景色
export function tableHeaderColor({ row, column, rowIndex, columnIndex }) {
  if (rowIndex === 0) {
    return 'background-color: #f5f7fa;color: #000;font-weight: 500;'
  }
}
