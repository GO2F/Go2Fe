import React from 'react';
import Link from 'umi/link';
import { Divider } from 'antd';

import TypePageConfigList from './type_test_config';
import test_config from './test_config';
let pageConfigList: TypePageConfigList = test_config as TypePageConfigList;
type TypeRoute = {
  path: string;
  name?: string;
  icon?: string;
  component?: string;
  // 记录扩展内容
  extendConfig?: any;
  routes?: TypeRoute[];
  exact: boolean;
};
let routeList: TypeRoute[] = [];
const PublicLayoutTemplateUri = './component/_layout.tsx';
const PublicListTemplateUri = './component/list/index.tsx';
const PublicCreateTemplateUri = './component/create/index.tsx';
const PublicDetailTemplateUri = './component/create/index.tsx';

for (let pageConfig of pageConfigList) {
  let dataModel = pageConfig.data_model;

  let currentPageRoute: TypeRoute = {
    path: pageConfig.base_url_path,
    component: PublicLayoutTemplateUri,
    name: pageConfig.name,
    routes: [],
    exact: false,
  };
  let listRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/list',
    component: PublicListTemplateUri,
    extendConfig: {},
    exact: true,
  };
  let createRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/create',
    component: PublicCreateTemplateUri,
    exact: true,
    extendConfig: {},
  };
  let updateRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/update/:id',
    component: PublicCreateTemplateUri,
    extendConfig: {},
    exact: true,
  };
  let detailRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/detail/:id',
    component: PublicDetailTemplateUri,
    extendConfig: {},
    exact: true,
  };

  listRouter.extendConfig['keyList'] = dataModel.key_list;
  listRouter.extendConfig['baseApiPath'] = pageConfig.base_api_path;
  listRouter.extendConfig['pageConfig'] = pageConfig.page_config;

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
