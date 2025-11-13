import React, { useEffect, useRef } from 'react';
import { useChatStore } from '../../stores/chatStore';
import { useUserStore } from '../../stores/userStore';
import { useChat } from '../../hooks/useChat';
import MessageBubble from './MessageBubble';
import InputBox from './InputBox';

export default function ChatInterface() {
  const { messages, isTyping } = useChatStore();
  const { studentId, context } = useUserStore();
  const { sendMessage, isLoading } = useChat(studentId!);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  // Auto-scroll to bottom when new message arrives
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const handleSendMessage = (message: string) => {
    if (!context) {
      console.error('No student context available');
      return;
    }
    sendMessage({ message, context });
  };

  return (
    <div className="chat-interface">
      <div className="messages-container">
        {messages.map((msg) => (
          <MessageBubble key={msg.id} message={msg} />
        ))}
        {isTyping && (
          <div className="typing-indicator">
            <span></span>
            <span></span>
            <span></span>
          </div>
        )}
        <div ref={messagesEndRef} />
      </div>

      <InputBox onSend={handleSendMessage} disabled={isLoading} />
    </div>
  );
}
