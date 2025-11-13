import { useMutation, useQueryClient } from '@tanstack/react-query';
import { coachAPI } from '../services/api';
import type { StudentContext, CoachResponse } from '../types.d';
import { useChatStore } from '../stores/chatStore';

export function useChat(studentId: string) {
  const queryClient = useQueryClient();
  const { addMessage, setIsTyping } = useChatStore();

  const sendMessageMutation = useMutation({
    mutationFn: ({ message, context }: { message: string; context: StudentContext }) =>
      // Call Go backend API
      coachAPI.sendMessage(studentId, message, context),
    
    onMutate: async ({ message }) => {
      // Optimistically add user message to UI
      addMessage({
        id: Date.now().toString(),
        role: 'user',
        content: message,
        timestamp: new Date().toISOString(),
      });
      setIsTyping(true);
    },

    onSuccess: (data: CoachResponse) => {
      // Add AI response from backend
      addMessage({
        id: Date.now().toString(),
        role: 'assistant',
        content: data.message,
        timestamp: data.timestamp,
        intervention: data.intervention,
        detected_barriers: data.detected_barriers,
        reward_earned: data.reward_earned,
        reasoning: data.reasoning,
      });
      setIsTyping(false);

      // Show reward notification if earned
      if (data.reward_earned) {
        // TODO: Show reward modal or toast
        console.log('🎮 Student earned a reward!');
      }

      // Show safeguarding alert if triggered
      if (data.safeguarding_alert) {
        // TODO: Show safeguarding modal or alert human team
        console.warn('⚠️ Safeguarding alert - human intervention required');
      }

      // Invalidate student profile to refetch updated state
      queryClient.invalidateQueries({ queryKey: ['student', studentId] });
    },

    onError: (error) => {
      console.error('Failed to send message:', error);
      setIsTyping(false);
      // TODO: Show error toast to user
      addMessage({
        id: Date.now().toString(),
        role: 'assistant',
        content: 'Sorry, something went wrong. Please try again.',
        timestamp: new Date().toISOString(),
      });
    },
  });

  return {
    sendMessage: sendMessageMutation.mutate,
    isLoading: sendMessageMutation.isPending,
    error: sendMessageMutation.error,
  };
}
