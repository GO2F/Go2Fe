import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card } from 'antd';

import { TypeExtends } from '../type_extend';

import Form from './form_register';

// import styles from './index.less';

export default class TablePage extends React.Component<any, any> {
  state = {
    item: {},
  };

  render() {
    let extendConfig: TypeExtends = this.props.route.extendConfig;

    let keyConfigList = extendConfig.keyList;
    console.log('keyConfigList => ', keyConfigList);
    return (
      <PageHeaderWrapper>
        <Card>
          <Form keyConfigList={keyConfigList}></Form>
        </Card>
      </PageHeaderWrapper>
    );
  }
}
