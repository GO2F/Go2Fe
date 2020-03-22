import React from 'react';
import styles from './index.less';
import { Table } from 'antd';

export default (props: any) => {
  console.log('props.columns ', props.columns);
  return (
    <div className={styles.container}>
      <div id="components-table-demo-basic">
        <Table columns={props.columns} dataSource={props.data} />
      </div>
    </div>
  );
};
