// Enums (must use 'as const' for type-only mode)
export const EngagementMode = {
  ComplianceMode: 'compliance',
  EngagementMode: 'engagement',
  ResistanceMode: 'resistance',
} as const;
export type EngagementMode = typeof EngagementMode[keyof typeof EngagementMode];

export const BarrierCategory = {
  AcuteBarrier: 'acute',
  ChronicBarrier: 'chronic',
  StructuralBarrier: 'structural',
} as const;
export type BarrierCategory = typeof BarrierCategory[keyof typeof BarrierCategory];

export const BubbleType = {
  Social: 'social',
  Sensory: 'sensory',
  Autonomy: 'autonomy',
  Status: 'status',
} as const;
export type BubbleType = typeof BubbleType[keyof typeof BubbleType];

// Core Types
export interface ETP {
  name: string;
  category: 'pain' | 'pleasure' | 'social' | 'goal';
  intensity: number; // 0-1
}

export interface BrainState {
  primalLevel: number; // 0-1
  emotionalLevel: number; // 0-1
  rationalLevel: number; // 0-1
  currentMode: 'primal' | 'emotional' | 'rational';
}

export interface RoutineProfile {
  routineDependency: number; // 0-1
  thinkingAtrophy: number; // 0-1
  fenceVoltage: number; // 0-1
}

export interface StudentContext {
  studentId: string;
  age: number;
  brainState: BrainState;
  activatedETPs: ETP[];
  routineProfile: RoutineProfile;
  socialNeed: number; // 0-1
  autonomyResistance: number; // 0-1
  statusSeeking: number; // 0-1
}

export interface InterventionLever {
  name: string;
  description: string;
  steps: string[];
  prerequisites?: string[];
  benefits?: string[];
  etpReduction: string[];
  brainStateTarget: string;
  whenToUse?: string[];
}

export interface StudentBarrier {
  id: string;
  name: string;
  category: string;
  activatedETPs: string[];
  avoidanceTactics: string[];
  effectiveLevers: InterventionLever[];
  underlyingCause: string;
}

export interface Bubble {
  type: string;
  description: string;
  controlValue: number; // 0-1
  socialValue: number; // 0-1
  distractionRisk: number; // 0-1
}

export interface TrustBargain {
  bubbleGranted: Bubble;
  expectedBehavior: string;
  riskMitigation: string[];
  emotionalIntelConvo: string;
}

// Coach Response
export interface CoachResponse {
  message: string;
  intervention: InterventionLever | null;
  detected_barriers: StudentBarrier[];
  safeguarding_alert: boolean;
  reward_earned: boolean;
  reasoning: string[];
  timestamp: string;
}

// Reward System
export interface GameAccessReward {
  unlockCode: string;
  validUntil: Date;
  earnedThrough: string;
  used: boolean;
}

export class RewardSystem {
  generateUnlock(taskCompleted: string): GameAccessReward {
    const code = this.generateTimeLimitedCode();
    const validUntil = new Date(Date.now() + 5 * 60 * 1000); // 5 minutes
    
    return {
      unlockCode: code,
      validUntil,
      earnedThrough: taskCompleted,
      used: false
    };
  }
  
  private generateTimeLimitedCode(): string {
    // Generate secure time-limited code
    const timestamp = Date.now();
    const random = Math.random().toString(36).substring(7);
    return `GAME-${timestamp}-${random}`.toUpperCase();
  }
  
  validateCode(code: string): boolean {
    // Check if code is still valid (within 5 minutes)
    const parts = code.split('-');
    if (parts.length !== 3) return false;
    
    const timestamp = parseInt(parts[1], 10);
    const now = Date.now();
    const fiveMinutes = 5 * 60 * 1000;
    
    return (now - timestamp) < fiveMinutes;
  }
}
