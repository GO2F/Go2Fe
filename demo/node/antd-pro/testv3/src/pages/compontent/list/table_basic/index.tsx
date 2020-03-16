import React from 'react';
import styles from './index.less';
import { Table, Divider, Tag } from 'antd';
import Link from 'umi/link';

const columns = [
  {
    title: 'id',
    dataIndex: 'id',
    key: 'id',
    // render: text => <a>{text}</a>
  },
  {
    title: '组件库名称',
    dataIndex: 'display_name',
    key: 'display_name',
  },
  {
    title: '包名',
    dataIndex: 'package_name',
    key: 'package_name',
  },
  // {
  //   title: 'Tags',
  //   key: 'tags',
  //   dataIndex: 'tags',
  //   render: tags => (
  //     <span>
  //       {tags.map(tag => {
  //         let color = tag.length > 5 ? 'geekblue' : 'green';
  //         if (tag === 'loser') {
  //           color = 'volcano';
  //         }
  //         return (
  //           <Tag color={color} key={tag}>
  //             {tag.toUpperCase()}
  //           </Tag>
  //         );
  //       })}
  //     </span>
  //   ),
  // },
  {
    title: '操作',
    key: 'action',
    render: (text: string, record: any) => (
      <span>
        <Link to={`/compontent/detail/${record.id}`}>详情</Link>
        <span>&nbsp;</span>
        <Link to={`/compontent/update/${record.id}`}>修改</Link>
        <Divider type="vertical" />
        <Link to={`/delete/${record.id}`}>删除</Link>
      </span>
    ),
  },
];

const data = [
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
];

export default () => (
  <div className={styles.container}>
    <div id="components-table-demo-basic">
      <Table columns={columns} dataSource={data} />
    </div>
  </div>
);
