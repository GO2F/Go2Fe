export type TypeKey = {
  key: string;
  // 字段中文名
  title: string;
  var_type: 'string' | 'int32';
  is_show_in_list: boolean;
  is_unique_key: boolean;
};

export type TypeDataModel = {
  key_list: TypeKey[];
};

export type TypePageConfig = {
  version: number;
  version: 0.1;
  data_model: TypeDataModel;
  // 数据模型在前端展示的模块名
  name: string;
  page_config: {
    create: boolean;
    update: boolean;
    detail: boolean;
  };
  base_url_path: string;
  base_api_path: string;
};

export type TypePageListConfig = TypePageConfig[];

declare const TestConfig: TypePageListConfig;

export default TypePageListConfig;
