import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface Message {
  id: string;
  role: 'user' | 'assistant';
  content: string;
  timestamp: string;
  intervention?: any;
  detected_barriers?: any[];
  reward_earned?: boolean;
  reasoning?: string[];
}

interface ChatState {
  messages: Message[];
  isTyping: boolean;
  addMessage: (message: Message) => void;
  setIsTyping: (isTyping: boolean) => void;
  clearMessages: () => void;
}

export const useChatStore = create<ChatState>()(
  persist(
    (set) => ({
      messages: [],
      isTyping: false,

      addMessage: (message) =>
        set((state) => ({
          messages: [...state.messages, message],
        })),

      setIsTyping: (isTyping) => set({ isTyping }),

      clearMessages: () => set({ messages: [] }),
    }),
    {
      name: 'chat-storage', // LocalStorage key
    }
  )
);
