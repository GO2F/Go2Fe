import TypePageConfigList from './type_test_config';
import test_config from './go2fe_generate_config';
import { TypeKey } from './type_test_config';
let pageConfigList: TypePageConfigList = test_config as TypePageConfigList;

export type TypeExtends = {
  // 字段配置列表
  keyList: TypeKey[];
  // 基础Api路径
  baseApiPath: string;
  // 基础页面路径
  baseUrlPath: string;
  // 配置启用的页面
  pageConfig: {
    create: boolean;
    update: boolean;
    detail: boolean;
  };
};

export type TypeRoute = {
  path: string;
  name?: string;
  icon?: string;
  component?: string;
  // 记录扩展内容
  extendConfig?: TypeExtends;
  routes?: TypeRoute[];
  exact: boolean;
};
let routeList: TypeRoute[] = [];
const PublicLayoutTemplateUri = './component/_layout.tsx';
const PublicListTemplateUri = './component/list/index.tsx';
const PublicCreateTemplateUri = './component/create/index.tsx';
const PublicDetailTemplateUri = './component/detail/index.tsx';

for (let pageConfig of pageConfigList) {
  let dataModel = pageConfig.data_model;

  let currentPageRoute: TypeRoute = {
    path: pageConfig.base_url_path,
    component: PublicLayoutTemplateUri,
    name: pageConfig.name,
    routes: [],
    exact: false,
  };
  let extendConfig: TypeExtends = {
    keyList: dataModel.key_list,
    baseApiPath: pageConfig.base_api_path,
    pageConfig: pageConfig.page_config,
    baseUrlPath: pageConfig.base_url_path,
  };
  let listRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/list',
    component: PublicListTemplateUri,
    extendConfig: extendConfig,
    exact: true,
  };
  let createRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/create',
    component: PublicCreateTemplateUri,
    exact: true,
    extendConfig: extendConfig,
  };
  let updateRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/update/:id',
    component: PublicCreateTemplateUri,
    extendConfig: extendConfig,
    exact: true,
  };
  let detailRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/detail/:id',
    component: PublicDetailTemplateUri,
    extendConfig: extendConfig,
    exact: true,
  };

  let dataModelPageConfig = pageConfig.page_config;
  currentPageRoute.routes?.push(listRouter);
  if (dataModelPageConfig.create) {
    currentPageRoute.routes?.push(createRouter);
  }
  if (dataModelPageConfig.update) {
    currentPageRoute.routes?.push(updateRouter);
  }
  if (dataModelPageConfig.detail) {
    currentPageRoute.routes?.push(detailRouter);
  }
  routeList.push(currentPageRoute);
}

export default routeList;
