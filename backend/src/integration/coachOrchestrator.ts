import { StudentContext, InterventionLever, StudentBarrier } from '../../../shared/types/etp';
import { BarrierDetector } from '../barriers/barrierDetection';
import { AgeAppropriateness } from '../barriers/ageAppropriate';
import { InterestPersonalizer } from '../personalization/interests';
import { PlayBreakManager } from '../personalization/playBreakGraduation';
import { TraumaDetector } from '../safeguarding/traumaDetection';

export interface CoachResponse {
  message: string;
  intervention: InterventionLever | null;
  detectedBarriers: StudentBarrier[];
  safeguardingAlert: boolean;
  rewardEarned: boolean;
  reasoning: string[];
}

export class CoachOrchestrator {
  private barrierDetector: BarrierDetector;
  private ageChecker: AgeAppropriateness;
  private interestPersonalizer: InterestPersonalizer;
  private playBreakManager: PlayBreakManager;
  private traumaDetector: TraumaDetector;

  constructor() {
    this.barrierDetector = new BarrierDetector();
    this.ageChecker = new AgeAppropriateness();
    this.interestPersonalizer = new InterestPersonalizer();
    this.playBreakManager = new PlayBreakManager();
    this.traumaDetector = new TraumaDetector();
  }

  async processStudentMessage(
    studentId: string,
    message: string,
    context: StudentContext
  ): Promise<CoachResponse> {
    const reasoning: string[] = [];

    // 1. Check for trauma/safeguarding issues (highest priority)
    const traumaResult = await this.traumaDetector.scan(message, context.age);
    if (traumaResult.severity >= 3) {
      reasoning.push('Safeguarding concern detected - escalating to human team');
      return {
        message: this.ageChecker.safeguardingResponse(context.age),
        intervention: null,
        detectedBarriers: [],
        safeguardingAlert: true,
        rewardEarned: false,
        reasoning
      };
    }

    // 2. Detect barriers
    const detectedBarriers = this.barrierDetector.detectBarriers(message, context);
    if (detectedBarriers.length > 0) {
      reasoning.push(`Detected barrier: ${detectedBarriers[0].barrier.name}`);
    }

    // 3. Select appropriate intervention
    const intervention = this.selectIntervention(detectedBarriers, context);
    reasoning.push(`Selected intervention: ${intervention?.name || 'none'}`);

    // 4. Generate age-appropriate response
    let responseMessage = this.generateResponse(intervention, context, detectedBarriers);
    responseMessage = this.ageChecker.adjustLanguage(responseMessage, context.age);
    
    // 5. Personalize with interests if available
    const interests = await this.interestPersonalizer.getInterests(studentId);
    if (interests.length > 0 && Math.random() > 0.7) { // 30% of responses
      responseMessage = this.interestPersonalizer.personalize(responseMessage, interests);
      reasoning.push('Personalized with student interests');
    }

    // 6. Check if reward earned
    const rewardEarned = this.checkRewardEarned(message, context, intervention);
    if (rewardEarned) {
      reasoning.push('Student earned play break reward');
    }

    return {
      message: responseMessage,
      intervention,
      detectedBarriers: detectedBarriers.map(d => d.barrier),
      safeguardingAlert: traumaResult.severity > 0,
      rewardEarned,
      reasoning
    };
  }

  private selectIntervention(
    detectedBarriers: any[],
    context: StudentContext
  ): InterventionLever | null {
    if (detectedBarriers.length === 0) return null;

    const topBarrier = detectedBarriers[0].barrier;
    
    // Check brain state - if emotional voltage too high, prioritize calming
    if (context.brainState.emotionalLevel > 0.7) {
      // Find calming intervention
      return topBarrier.effectiveLevers.find(l => 
        l.brainStateTarget.includes('lower') || 
        l.brainStateTarget.includes('calm')
      ) || topBarrier.effectiveLevers[0];
    }

    // Return first effective lever for the barrier
    return topBarrier.effectiveLevers[0] || null;
  }

  private generateResponse(
    intervention: InterventionLever | null,
    context: StudentContext,
    barriers: any[]
  ): string {
    if (!intervention) {
      return "I'm here to help. What would you like to work on?";
    }

    // Generate response based on intervention steps
    const steps = intervention.steps;
    return steps[0] || intervention.description;
  }

  private checkRewardEarned(
    message: string,
    context: StudentContext,
    intervention: InterventionLever | null
  ): boolean {
    // Simple heuristic: longer messages = more engagement = reward
    // In real implementation, track task completion
    return message.length > 50 && !message.toLowerCase().includes("i don't know");
  }
}
