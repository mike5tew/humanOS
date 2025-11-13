import type { StudentContext, CoachResponse } from '../types.d';

export interface StudentProfile {
  student_id: string;
  age: number;
  brain_state: {
    primal_level: number;
    emotional_level: number;
    rational_level: number;
  };
  active_barriers: string[];
  rewards_earned: number;
  play_break_stage: string;
  last_interaction: string;
}

export class CoachAPI {
  private baseURL: string;

  constructor(baseURL = 'http://localhost:8080') {
    this.baseURL = baseURL;
  }

  async sendMessage(
    studentId: string,
    message: string,
    context: StudentContext
  ): Promise<CoachResponse> {
    const response = await fetch(`${this.baseURL}/api/coach/message`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        student_id: studentId,
        message,
        context,
      }),
    });

    if (!response.ok) {
      throw new Error(`API error: ${response.status}`);
    }

    return response.json();
  }

  async getStudentProfile(studentId: string): Promise<StudentProfile> {
    // Call Go backend to get student profile
    const response = await fetch(
      `${this.baseURL}/api/student/${studentId}/profile`
    );

    if (!response.ok) {
      throw new Error(`Failed to fetch student profile: ${response.status}`);
    }

    return response.json();
  }

  async getHealth(): Promise<{ status: string }> {
    const response = await fetch(`${this.baseURL}/api/health`);
    return response.json();
  }
}

export const coachAPI = new CoachAPI();
