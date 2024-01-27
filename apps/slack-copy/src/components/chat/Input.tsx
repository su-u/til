import React, { FC } from 'react';
import { Button, Input as AntdInput, Select, Space } from 'antd';

export const Input: FC = () => {
  return (
    <Space.Compact style={{ width: '100%' }}>
      <AntdInput defaultValue="" />
      <Button type="primary">Submit</Button>
    </Space.Compact>
  );
};
