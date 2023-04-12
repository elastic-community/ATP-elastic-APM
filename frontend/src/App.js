import React from 'react';

import { Layout } from 'antd';

import './App.css';

import {  Menu, Breadcrumb } from 'antd';
import Ranking from './ranking';

const { Header, Content, Footer } = Layout;


function App() {
  return (
    <div className="App">
    <>
    <Layout className="layout">
    <Header>
      <div className="logo" />
      <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['1']}>
        <Menu.Item key="1">Ranking Unholster 2023 </Menu.Item>
      </Menu>
    </Header>
    <Content style={{ padding: '0 50px' }}>
      <Breadcrumb style={{ margin: '16px 0' }}>
        <Breadcrumb.Item>Inicio</Breadcrumb.Item>
        <Breadcrumb.Item>Listar</Breadcrumb.Item>
      </Breadcrumb>
      <div className="site-layout-content">
          <Ranking /> 

      </div>
    </Content>
    <Footer style={{ textAlign: 'center' }}>Sistema de Ranking ATP de miercoles  Unholster </Footer>
  </Layout>
  </>
  </div>
  );
}

export default App;
