import React from 'react';
import ChatInterface from '../components/chat/ChatInterface';

export default function ChatPage() {
  return (
    <div className="chat-page">
      <header className="chat-header">
        <h1>HumanOS AI Tutor</h1>
        <nav>
          <a href="/dashboard">Dashboard</a>
        </nav>
      </header>

      <main className="chat-main">
        <ChatInterface />
      </main>
    </div>
  );
}
