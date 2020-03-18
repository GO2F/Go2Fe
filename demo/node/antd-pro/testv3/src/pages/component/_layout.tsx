import React from 'react';
import router from 'umi/router';

export default (props: any) => {
  console.log('props.route => ', props.route);
  if (props.location.pathname === '/compontent') {
    router.push('/compontent/list');
    return <div />;
  } else {
    return <div>{props.children}</div>;
  }
};
