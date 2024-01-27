'use client';
import React, { useState } from 'react';
import { Avatar, List, Radio, Space, Typography, Tooltip } from 'antd';
import { dayjs } from '@/util/dayjs';
import { Database } from '@/schema';

const { Paragraph } = Typography;

type PaginationPosition = 'top' | 'bottom' | 'both';

type PaginationAlign = 'start' | 'center' | 'end';

const positionOptions = ['top', 'bottom', 'both'];

const alignOptions = ['start', 'center', 'end'];

type Props = {
  messages: Database['public']['Tables']['messages']['Row'][];
};

export const Messages: React.FC<Props> = ({ messages }) => {
  const [position, setPosition] = useState<PaginationPosition>('bottom');
  const [align, setAlign] = useState<PaginationAlign>('center');

  return (
    <List
      pagination={{ position, align }}
      dataSource={messages}
      renderItem={(item, index) => (
        <List.Item>
          <List.Item.Meta
            avatar={<Avatar src={`https://api.dicebear.com/7.x/miniavs/svg?seed=${index}`} />}
            title={item.text}
            description={item.text}
          />
          <Tooltip placement="top" title={dayjs(item.inserted_at).format('LLL')}>
            <Paragraph>{dayjs(item.inserted_at).format('HH:mm')}</Paragraph>
          </Tooltip>
        </List.Item>
      )}
    />
  );
};
