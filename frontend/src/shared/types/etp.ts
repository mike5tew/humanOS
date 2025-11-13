export enum EngagementMode {
  ComplianceMode = 'compliance',
  EngagementMode = 'engagement',
  ResistanceMode = 'resistance'
}

export enum BarrierCategory {
  AcuteBarrier = 'acute',
  ChronicBarrier = 'chronic',
  StructuralBarrier = 'structural'
}

export enum BubbleType {
  Social = 'social',
  Sensory = 'sensory',
  Autonomy = 'autonomy',
  Status = 'status'
}

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
  category: BarrierCategory;
  activatedETPs: string[];
  avoidanceTactics: string[];
  effectiveLevers: InterventionLever[];
  underlyingCause: string;
}

export interface Bubble {
  type: BubbleType;
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
