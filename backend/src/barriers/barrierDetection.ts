import { StudentBarrier, StudentContext, ETP } from '../../../shared/types/etp';
import barriersData from '../../../shared/schemas/barriers.json';

export interface DetectedBarrier {
  barrier: StudentBarrier;
  confidence: number; // 0-1
  reasoning: string[];
}

export class BarrierDetector {
  private barriers: StudentBarrier[];

  constructor() {
    this.barriers = barriersData.barriers as any;
  }

  detectBarriers(input: string, context: StudentContext): DetectedBarrier[] {
    const detected: DetectedBarrier[] = [];

    // Check for "I don't know" - primary avoidance indicator
    if (this.containsIDontKnow(input)) {
      const lackOfMotivation = this.barriers.find(b => b.id === 'lack_of_motivation');
      if (lackOfMotivation) {
        detected.push({
          barrier: lackOfMotivation,
          confidence: 0.8,
          reasoning: ['Student used "I don\'t know" - primary avoidance tactic']
        });
      }
    }

    // Check for confrontational language
    if (this.isConfrontational(input)) {
      const confrontational = this.barriers.find(b => b.id === 'confrontational_showoff');
      if (confrontational) {
        detected.push({
          barrier: confrontational,
          confidence: 0.7,
          reasoning: ['Confrontational or dismissive language detected']
        });
      }
    }

    // Check for minimal response (potential silent avoider)
    if (this.isMinimalResponse(input)) {
      const silentAvoider = this.barriers.find(b => b.id === 'silent_avoider');
      if (silentAvoider) {
        detected.push({
          barrier: silentAvoider,
          confidence: 0.6,
          reasoning: ['Minimal engagement, very short response']
        });
      }
    }

    // ...existing code for other barrier detection patterns...

    return detected.sort((a, b) => b.confidence - a.confidence);
  }

  private containsIDontKnow(input: string): boolean {
    const variations = [
      /i don'?t know/i,
      /idk/i,
      /dunno/i,
      /no idea/i
    ];
    return variations.some(pattern => pattern.test(input));
  }

  private isConfrontational(input: string): boolean {
    const patterns = [
      /this is (stupid|dumb|boring)/i,
      /why (do|should) i/i,
      /i don'?t (care|want to)/i,
      /whatever/i,
      /so what/i
    ];
    return patterns.some(pattern => pattern.test(input));
  }

  private isMinimalResponse(input: string): boolean {
    const trimmed = input.trim();
    return trimmed.length < 10 && !this.containsIDontKnow(input);
  }
}
