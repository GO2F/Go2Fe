import { IConfig, IPlugin } from 'umi-types';
import defaultSettings from './defaultSettings'; // https://umijs.org/config/
import slash from 'slash2';
import themePluginConfig from './themePluginConfig';
import proxy from './proxy';
import webpackPlugin from './plugin.config';

import TypePageConfigList from './type_test_config';
import test_config from './test_config';
let pageConfigList: TypePageConfigList = test_config as TypePageConfigList;
type TypeRoute = {
  path: string;
  name?: string;
  icon?: string;
  component?: string;
  // 记录扩展内容
  extends?: any;
  routes?: TypeRoute[];
  exact: boolean;
};
let routeList: TypeRoute[] = [];
const PublicLayoutTemplateUri = './component/_layout.tsx';
const PublicListTemplateUri = './component/list/index.tsx';
const PublicCreateTemplateUri = './component/create/index.tsx';
const PublicDetailTemplateUri = './component/create/index.tsx';
// {
//   path: 'component',
//   name: 'component',
//   icon: 'smile',
//   component: './component/_layout.tsx',
//   // 可以通过props.route获取
//   extends: {
//     hello: 'world',
//   },
//   routes: [
//     {
//       path: '/component/list',
//       component: './component/list/index.tsx',
//     },
//     {
//       path: '/component/create',
//       exact: true,
//       component: './component/create/index.tsx',
//     },
//     {
//       path: '/component/update/:id',
//       exact: true,
//       component: './component/create/index.tsx',
//     },
//     {
//       path: '/component/detail/:id',
//       exact: true,
//       component: './component/create/index.tsx',
//     },
//   ],
// },

for (let pageConfig of pageConfigList) {
  let dataModel = pageConfig.data_model;

  let currentPageRoute: TypeRoute = {
    path: pageConfig.base_url_path,
    component: PublicLayoutTemplateUri,
    name: pageConfig.name,
    routes: [],
    exact: false,
    //   icon: 'smile',
    //   component: './component/_layout.tsx',
  };
  let listRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/list',
    component: PublicListTemplateUri,
    extends: {},
    exact: true,
  };
  let createRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/create',
    component: PublicCreateTemplateUri,
    exact: true,
    extends: {},
  };
  let updateRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/update/:id',
    component: PublicCreateTemplateUri,
    extends: {},
    exact: true,
  };
  let detailRouter: TypeRoute = {
    path: pageConfig.base_url_path + '/detail/:id',
    component: PublicDetailTemplateUri,
    extends: {},
    exact: true,
  };

  let keyConfigList = dataModel.key_list;
  let tableColumnList = [];
  for (let keyConfig of keyConfigList) {
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
    render: (text: string, record: any) => {
      return '123123';
    },
    // (
    // <span>
    //   <Link to={`/component/detail/${record.id}`}>详情</Link>
    //   <span>&nbsp;</span>
    //   <Link to={`/component/update/${record.id}`}>修改</Link>
    //   <Divider type="vertical" />
    //   <Link to={`/delete/${record.id}`}>删除</Link>
    // </span>
    // ),
  });
  listRouter.extends['tableColumnList'] = tableColumnList;
  listRouter.extends['baseApiPath'] = pageConfig.base_api_path;

  let dataModelPageConfig = pageConfig.page_config;
  console.log('dataModelPageConfig => ', dataModelPageConfig);
  console.log('dataModel => ', Object.keys(dataModel));
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

const { pwa } = defaultSettings;

// preview.pro.ant.design only do not use in your production ;
// preview.pro.ant.design 专用环境变量，请不要在你的项目中使用它。
const { ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION, REACT_APP_ENV } = process.env;
const isAntDesignProPreview = ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION === 'site';

const plugins: IPlugin[] = [
  ['umi-plugin-antd-icon-config', {}],
  [
    'umi-plugin-react',
    {
      antd: true,
      dva: {
        hmr: true,
      },
      locale: {
        // default false
        // 不需要启用多语言功能
        enable: false,
        // default zh-CN
        default: 'zh-CN',
        // default true, when it is true, will use `navigator.language` overwrite default
        baseNavigator: true,
      },
      dynamicImport: {
        loadingComponent: './components/PageLoading/index',
        webpackChunkName: true,
        level: 3,
      },
      pwa: pwa
        ? {
            workboxPluginMode: 'InjectManifest',
            workboxOptions: {
              importWorkboxFrom: 'local',
            },
          }
        : false,
      // default close dll, because issue https://github.com/ant-design/ant-design-pro/issues/4665
      // dll features https://webpack.js.org/plugins/dll-plugin/
      // dll: {
      //   include: ['dva', 'dva/router', 'dva/saga', 'dva/fetch'],
      //   exclude: ['@babel/runtime', 'netlify-lambda'],
      // },
    },
  ],
  [
    'umi-plugin-pro-block',
    {
      moveMock: false,
      moveService: false,
      modifyRequest: true,
      autoAddMenu: true,
    },
  ],
];

if (isAntDesignProPreview) {
  // 针对 preview.pro.ant.design 的 GA 统计代码
  plugins.push([
    'umi-plugin-ga',
    {
      code: 'UA-72788897-6',
    },
  ]);
  plugins.push(['umi-plugin-antd-theme', themePluginConfig]);
}
// console.log('routeList =>', JSON.stringify(routeList));
export default {
  plugins,
  hash: true,
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/zh/guide/router.html
  routes: [
    {
      path: '/',
      // component: '../layouts/SecurityLayout',
      component: '../layouts/BasicLayout',
      routes: routeList,
    },
    {
      component: './404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  define: {
    REACT_APP_ENV: REACT_APP_ENV || false,
    ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION:
      ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION || '', // preview.pro.ant.design only do not use in your production ; preview.pro.ant.design 专用环境变量，请不要在你的项目中使用它。
  },
  ignoreMomentLocale: true,
  lessLoaderOptions: {
    javascriptEnabled: true,
  },
  disableRedirectHoist: true,
  cssLoaderOptions: {
    modules: true,
    getLocalIdent: (
      context: {
        resourcePath: string;
      },
      _: string,
      localName: string,
    ) => {
      if (
        context.resourcePath.includes('node_modules') ||
        context.resourcePath.includes('ant.design.pro.less') ||
        context.resourcePath.includes('global.less')
      ) {
        return localName;
      }
      const match = context.resourcePath.match(/src(.*)/);
      if (match && match[1]) {
        const antdProPath = match[1].replace('.less', '');
        const arr = slash(antdProPath)
          .split('/')
          .map((a: string) => a.replace(/([A-Z])/g, '-$1'))
          .map((a: string) => a.toLowerCase());
        return `antd-pro${arr.join('-')}-${localName}`.replace(/--/g, '-');
      }
      return localName;
    },
  },
  manifest: {
    basePath: '/',
  },
  proxy: proxy[REACT_APP_ENV || 'dev'],
  chainWebpack: webpackPlugin,
} as IConfig;
