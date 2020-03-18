import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card } from 'antd';

import TableBase from './table_basic';
// import styles from './index.less';

export default (props: any): React.ReactNode => {
  console.log('table page, props =>', props);
  return (
    <PageHeaderWrapper>
      <Card>
        <TableBase />
      </Card>
    </PageHeaderWrapper>
  );
};
