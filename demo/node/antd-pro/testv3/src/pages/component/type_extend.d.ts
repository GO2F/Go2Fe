// extends字段类型定义

export type TypeKey = {
  key: string;
  // 字段中文名
  title: string;
  var_type: 'string' | 'int32';
  is_show_in_list: boolean;
  is_unique_key: boolean;
};

export type TypeExtends = {
  // 字段配置列表
  keyList: TypeKey[];
  // 基础Api路径
  baseApiPath: string;
  // 配置启用的页面
  pageConfig: {
    create: boolean;
    update: boolean;
    detail: boolean;
  };
};
