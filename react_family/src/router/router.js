import React from 'react';

import {BrowserRouter as Router, Route, Switch, Link} from 'react-router-dom';
/*引用的相对路径*/
import Home from 'pages/Home/Home';
import Page1 from 'pages/Page1/Page1';
import Counter from 'pages/Counter/Counter';
import UserInfo  from 'pages/UserInfo/UserInfo';

/*
 * 1.Router：我们可以把它看成是react路由的一个路由外层盒子，他里面的内容 就是我们单页面应用的路由以及组件。
 * 2.Link:他是react路由中的点击切换到哪一个组件的链接(这里之所以不说是页面，而是说组件，因为切换到另外一个界面只是展示效果，react的本质还是一个单页面
 * 应用--single  page application)。
 * 3.Route:代表了你的路由界面，path代表路径，component代表路径所对应的界面。
 * 特别说明：第一、Router下面只能包含一个盒子标签，类似这里的div。
 * 第二、Link代表一个链接，在html界面中会解析成a标签。作为一个链接，必须有一个to属性，代表链接地址。这个链接地址是一个相对路径。
 */
const getRouter = () => (
    <Router>
        <div>
            <ul>
                <li><Link to="/">首页</Link></li>
                <li><Link to="/page1">Page1</Link></li>
                <li><Link to="/counter">Counter</Link></li>
                <li><Link to="/userinfo">UserInfo</Link></li>
            </ul>
            <Switch>
                <Route exact path="/" component={Home}/>
                <Route path="/page1" component={Page1}/>
                <Route path="/counter" component={Counter}/>
                <Route path="/userinfo" component={UserInfo}/>
            </Switch>
        </div>
    </Router>
);

export default getRouter;