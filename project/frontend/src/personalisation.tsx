import { ObjectId } from 'mongodb';

/**
 * MongoDB Schema for Student Profile and Interaction History
 */
export interface StudentProfile {
  _id: ObjectId;
  studentId: string;
  age: number;
  developmentalStage: 'early_years' | 'middle_years' | 'early_adolescence' | 'late_adolescence';
  
  // Barrier profile
  primaryBarrier?: string;
  barrierHistory: BarrierDetection[];
  
  // Play break stage
  playBreakStage: 1 | 2 | 3 | 4;
  playBreakHistory: PlayBreakTransition[];
  
  // Interests for personalization
  interests: Interest[];
  
  // Trauma/safeguarding
  traumaIndicators: TraumaFlag[];
  safeguardingStatus: 'clear' | 'monitoring' | 'escalated' | 'active_support';
  
  // Progress tracking
  workDurationTolerance: number; // seconds
  selfInitiationRate: number; // 0-1
  rewardDependency: number; // 0-1 (1 = fully dependent)
  
  // Metadata
  createdAt: Date;
  lastInteraction: Date;
  totalInteractions: number;
}

export interface BarrierDetection {
  timestamp: Date;
  barrierType: string;
  confidence: number;
  context: string;
  interventionUsed: string;
  outcome?: string;
}

export interface PlayBreakTransition {
  timestamp: Date;
  fromStage: number;
  toStage: number;
  reason: string;
  success: boolean;
}

export interface Interest {
  category: string; // 'games', 'sports', 'hobbies', 'subjects', 'media'
  specific: string; // e.g., 'Minecraft', 'football', 'drawing', 'space'
  confidence: number; // 0-1 based on frequency of mention
  lastMentioned: Date;
  mentionCount: number;
}

export interface TraumaFlag {
  timestamp: Date;
  severity: 1 | 2 | 3 | 4; // Escalation level
  category: 'sexual' | 'violence' | 'neglect' | 'emotional' | 'physical' | 'other';
  content: string; // Encrypted or secured
  aiResponse: string;
  humanReviewed: boolean;
  reviewedBy?: string;
  reviewedAt?: Date;
  outcome?: string;
}

/**
 * Interaction History for Chat Context
 */
export interface Interaction {
  _id: ObjectId;
  studentId: string;
  timestamp: Date;
  sessionId: string;
  
  // Message content
  studentMessage: string;
  aiResponse: string;
  
  // Context
  topic?: string;
  taskType?: string;
  
  // Metadata
  emotionalTone?: 'positive' | 'negative' | 'neutral' | 'distressed';
  barrierDetected?: string;
  traumaFlagged: boolean;
  rewardGiven: boolean;
  
  // Vector embedding for semantic search
  embedding?: number[]; // For Weaviate
}

/**
 * Weaviate Schema for Semantic Search
 */
export interface WeaviateInteraction {
  studentId: string;
  timestamp: string;
  content: string;
  topic: string;
  emotionalTone: string;
  interestsDetected: string[];
  
  // Relationships
  relatedInteractions?: string[]; // IDs of semantically similar interactions
}

export interface WeaviateInterest {
  studentId: string;
  category: string;
  specific: string;
  context: string; // How it was mentioned
  
  // Semantic similarity for personalization
  relatedTopics?: string[];
}