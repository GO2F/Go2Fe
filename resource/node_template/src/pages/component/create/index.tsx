import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card, message } from 'antd';

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
    if (this.getPageType() === 'update') {
      this.asyncFetchData();
    }
  }

  getPageType(): 'create' | 'update' {
    // 判断当前页面是create页面还是update页面
    let extendConfig: TypeExtends = this.props.route.extendConfig;
    if (location.href.includes(`${extendConfig.baseUrlPath}/create`)) {
      return 'create';
    }
    return 'update';
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
    let item = response?.data || {};
    this.setState({
      item: item,
    });
  };

  handleSubmit = async (item: any) => {
    let extendConfig: TypeExtends = this.props.route.extendConfig;

    let baseApi = extendConfig.baseApiPath;
    if (this.getPageType() === 'create') {
      // 对应create
      let api = baseApi + '/create';
      console.log('create: add item =>', api, item);
      let response = await request
        .post(api, {
          params: {
            ...item,
          },
        })
        .catch(() => {
          return {};
        });
      if (response.code !== 0) {
        message.error(`提交失败, 错误信息:${response?.msg}`);
        console.log('接口返回值异常, 错误信息:', response?.msg);
        return;
      }
      message.info('创建成功', 1, () => {
        // 提交完毕, 回到列表页
        router.push(extendConfig.baseUrlPath + '/list');
      });
    } else {
      // 对应update
      let api = baseApi + '/update';
      console.log('update: submit item =>', api, item);
      let response = await request
        .post(api, {
          params: {
            // 更新信息中, 一定要带id
            ...item,
          },
        })
        .catch(() => {
          return {};
        });
      if (response.code !== 0) {
        message.error(`更新失败, 错误信息:${response?.msg}`);
        console.log('接口返回值异常, 错误信息:', response?.msg);
        return;
      }
      message.info('更新成功', 1, () => {
        // 提交完毕, 回到列表页
        router.push(extendConfig.baseUrlPath + '/list');
      });
    }
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
            homeUri={extendConfig.baseUrlPath + '/list'}
          ></Form>
        </Card>
      </PageHeaderWrapper>
    );
  }
}
