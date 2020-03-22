import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card } from 'antd';
import Link from 'umi/link';
import { Divider } from 'antd';

import request from 'umi-request';
import { TypeExtends } from '../type_extend';
import TableBase from './table_basic';
// import styles from './index.less';

export default class TablePage extends React.Component<any, any> {
  state = {
    itemList: [],
  };

  async fetchData() {
    const extendConfig: TypeExtends = this.props.route.extendConfig;
    const baseApiPath = extendConfig?.baseApiPath;
    const api = `${baseApiPath}/list`;
    let response = await request.get(api).catch(() => {
      return {};
    });
    let itemList = response?.data || [];
    this.setState({
      itemList: itemList,
    });
  }

  componentDidMount() {
    // 初始化数据
    this.fetchData();
    console.log('数据初始化完毕 this.state.itemList => ', this.state.itemList);
  }

  render(): React.ReactNode {
    let extendConfig: TypeExtends = this.props.route.extendConfig;

    let keyConfigList = extendConfig.keyList;
    let tableColumnList = [];
    for (let keyConfig of keyConfigList) {
      // 目前支持直接展示int/string等文本类型

      // 处理字段隐藏逻辑
      if (keyConfig.is_show_in_list === false) {
        continue;
      }

      let columnItem = {
        title: keyConfig.title,
        dataIndex: keyConfig.key,
        key: keyConfig.key,
      };
      tableColumnList.push(columnItem);
    }
    // 添加操作栏
    tableColumnList.push({
      title: '操作',
      key: 'action',
      fixed: 'right',
      render: (text: string, record: any) => {
        return (
          <div>
            <span>
              {extendConfig?.pageConfig?.detail ? (
                <Link to={`/component/detail/${record.id}`}>详情</Link>
              ) : null}
              <span>&nbsp;</span>
              {extendConfig?.pageConfig?.update ? (
                <Link to={`/component/update/${record.id}`}>修改</Link>
              ) : null}
              <Divider type="vertical" />
              <Link to={`/delete/${record.id}`}>删除</Link>
            </span>
          </div>
        );
      },
    });

    return (
      <PageHeaderWrapper>
        <Card>
          <TableBase columns={tableColumnList} data={this.state.itemList} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}
