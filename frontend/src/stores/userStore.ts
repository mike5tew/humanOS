import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { StudentContext } from '../types.d';

interface UserState {
  studentId: string | null;
  context: StudentContext | null;
  setStudentId: (id: string) => void;
  updateContext: (context: Partial<StudentContext>) => void;
  clearUser: () => void;
}

export const useUserStore = create<UserState>()(
  persist(
    (set) => ({
      studentId: null,
      context: null,

      setStudentId: (id) => set({ studentId: id }),

      updateContext: (contextUpdate) =>
        set((state) => ({
          context: state.context
            ? { ...state.context, ...contextUpdate }
            : null,
        })),

      clearUser: () => set({ studentId: null, context: null }),
    }),
    {
      name: 'user-storage',
    }
  )
);
