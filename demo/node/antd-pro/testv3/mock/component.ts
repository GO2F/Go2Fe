// import { Request, Response } from 'express';

// 代码中会兼容本地 service mock 以及部署站点的静态数据
export default {
  // 支持值为 Object 和 Array
  'GET /api/component/list': {
    data: [
      {
        id: 1,
        display_name: 'antd',
        package_name: '@antd',
      },
      {
        id: 2,
        display_name: 'cutter-ui',
        package_name: '@ke/cutter-ui',
      },
      {
        id: 3,
        display_name: '饿了么ui',
        package_name: '@element-ui',
      },
    ],
  },
};
