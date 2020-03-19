import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card } from 'antd';

import { TypeExtends } from '../type_extend';

import request from 'umi-request';
import router from 'umi/router';

import Form from './form_register';

// import styles from './index.less';

export default class TablePage extends React.Component<any, any> {
  state = {
    item: {},
  };

  componentDidMount() {
    this.asyncFetchData();
  }

  asyncFetchData = async () => {
    const extendConfig: TypeExtends = this.props.route.extendConfig;

    console.log('hello world =>', this.props.route);

    const baseApiPath = extendConfig?.baseApiPath;
    const api = `${baseApiPath}/get`;
    let response = await request
      .get(api, {
        params: {
          id: this.props.match.params.id,
        },
      })
      .catch(() => {
        return {};
      });
    console.log('response =>', response);
    let item = response.data || {};
    this.setState({
      item: item,
    });
  };

  handleSubmit = async (item: any) => {
    console.log('submit item =>', item);
    // 提交完毕, 回到列表页
    let extendConfig: TypeExtends = this.props.route.extendConfig;
    router.push(extendConfig.baseUrlPath + '/list');
  };

  render() {
    let extendConfig: TypeExtends = this.props.route.extendConfig;
    console.log('hello world =>', this.props);

    let keyConfigList = extendConfig.keyList;
    console.log('keyConfigList => ', keyConfigList);
    return (
      <PageHeaderWrapper>
        <Card>
          <Form
            keyConfigList={keyConfigList}
            initData={this.state.item}
            handleSubmit={this.handleSubmit}
          ></Form>
        </Card>
      </PageHeaderWrapper>
    );
  }
}
