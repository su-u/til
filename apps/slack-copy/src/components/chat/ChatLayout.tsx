'use client';
import { Input } from '@/components/chat/Input';
import React from 'react';
import { Breadcrumb, Layout, Menu, theme, Typography } from 'antd';

const { Header, Content, Footer, Sider } = Layout;

type Props = {
  children?: React.ReactNode;
};

export const ChatLayout: React.FC<Props> = ({ children }) => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <Layout>
      <Header style={{ background: colorBgContainer, display: 'flex', alignItems: 'center' }}>
        <Typography.Title level={2}>h2. Ant Design</Typography.Title>
      </Header>
      <Content style={{ margin: '16px' }}>
        <div
          style={{
            padding: 24,
            minHeight: 360,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          {children}
        </div>
      </Content>
      <Footer style={{ textAlign: 'center' }}>
        <Input />
      </Footer>
    </Layout>
  );
};
