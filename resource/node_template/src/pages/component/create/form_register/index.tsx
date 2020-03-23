import React from 'react';
import styles from './index.less';
import { Form, Input, Button, Divider } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import { TypeExtends } from '../../type_extend';
import router from 'umi/router';

interface CustomerFormProps extends FormComponentProps {
  // 初始值
  initData: {
    [key: string]: any;
  };
  // 列表页地址
  homeUri: string;
  // 回调函数
  handleSubmit: Function;
  // 配置列表, 用于生成Form表单
  keyConfigList: TypeExtends['keyList'];
}

class RegistrationForm extends React.Component<
  CustomerFormProps,
  {
    confirmDirty: boolean;
    autoCompleteResult: any[];
  }
> {
  state = {
    confirmDirty: false,
    autoCompleteResult: [],
  };

  handleSubmit = async (e: any) => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
      }
    });
    let result = this.props.form.getFieldsValue();
    await this.props.handleSubmit(result);
  };

  render() {
    const { getFieldDecorator } = this.props.form;

    const keyConfigList = this.props.keyConfigList;

    let formItemList: React.ReactElement[] = [];
    for (let keyConfig of keyConfigList) {
      // 对key特殊处理
      if (keyConfig.is_unique_key) {
        let formItem = (
          <Form.Item label={keyConfig.title}>
            <span>{this.props.initData[keyConfig.key]}</span>
          </Form.Item>
        );
        formItemList = [formItem, ...formItemList];
        continue;
      }

      if (keyConfig.var_type === 'string') {
        let formItem = (
          <Form.Item label={keyConfig.title || ' '}>
            {getFieldDecorator(keyConfig.key, {
              initialValue: this.props.initData[keyConfig.key],
            })(<Input />)}
          </Form.Item>
        );
        formItemList.push(formItem);
        continue;
      }
    }

    const formItemLayout = {
      labelCol: { span: 6 },
      wrapperCol: { span: 14 },
    };
    const tailFormItemLayout = {
      wrapperCol: {
        offset: 6,
        span: 14,
      },
    };

    return (
      <Form {...formItemLayout} onSubmit={this.handleSubmit}>
        {formItemList}
        <Form.Item {...tailFormItemLayout}>
          <Button
            type="default"
            onClick={() => {
              router.push(this.props.homeUri);
            }}
          >
            返回
          </Button>
          &nbsp; &nbsp; &nbsp;
          <Button type="primary" htmlType="submit">
            提交
          </Button>
        </Form.Item>
      </Form>
    );
  }
}

const WrappedRegistrationForm = Form.create({ name: 'register' })(RegistrationForm);

export default (props: any) => (
  <div className={styles.container}>
    <div id="components-form-demo-register">
      <WrappedRegistrationForm
        homeUri={props.homeUri}
        initData={props.initData}
        keyConfigList={props.keyConfigList}
        handleSubmit={props.handleSubmit}
      />
    </div>
  </div>
);
