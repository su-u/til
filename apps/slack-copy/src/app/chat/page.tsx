import { supabase } from '@/util/db';
import { AppLayout } from '@/components/AppLayout';
import React from 'react';
import { ChatLayout } from '@/components/chat/ChatLayout';
import { Messages } from '@/components/chat/Messages';

export default async function Notes() {
  const { data: messages } = await supabase.from('messages').select();

  console.log(JSON.stringify(messages, null, 2));

  return (
    <AppLayout>
      <ChatLayout>
        <Messages messages={messages} />
      </ChatLayout>
    </AppLayout>
  );
}
