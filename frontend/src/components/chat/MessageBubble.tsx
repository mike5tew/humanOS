import React from 'react';

interface Message {
  id: string;
  role: 'user' | 'assistant';
  content: string;
  timestamp: string;
  reward_earned?: boolean;
  detected_barriers?: any[];
}

interface MessageBubbleProps {
  message: Message;
}

export default function MessageBubble({ message }: MessageBubbleProps) {
  const isUser = message.role === 'user';

  return (
    <div className={`message-bubble ${isUser ? 'user' : 'assistant'}`}>
      <div className="message-content">
        {message.content}
      </div>

      {message.reward_earned && (
        <div className="reward-badge">
          🎮 Earned 5-minute play break!
        </div>
      )}

      {message.detected_barriers && message.detected_barriers.length > 0 && (
        <div className="barrier-info">
          <small>Detected: {message.detected_barriers[0].name}</small>
        </div>
      )}

      <div className="message-timestamp">
        {new Date(message.timestamp).toLocaleTimeString()}
      </div>
    </div>
  );
}
