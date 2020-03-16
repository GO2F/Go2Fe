import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card } from 'antd';

import TableBase from './table_basic';
// import styles from './index.less';

export default (): React.ReactNode => (
  <PageHeaderWrapper>
    <Card>
      <p>123123</p>
      <TableBase />
    </Card>
  </PageHeaderWrapper>
);
