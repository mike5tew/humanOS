import { Collection } from 'mongodb';
import { StudentProfile, TraumaFlag } from '../database/schemas';

export class TraumaDetectionSystem {
  private studentCollection: Collection<StudentProfile>;
  private safeguardingAlertEndpoint: string;
  
  constructor(
    studentCollection: Collection<StudentProfile>,
    safeguardingEndpoint: string
  ) {
    this.studentCollection = studentCollection;
    this.safeguardingAlertEndpoint = safeguardingEndpoint;
  }
  
  async detectTraumaIndicators(
    studentId: string,
    message: string,
    age: number
  ): Promise<{ detected: boolean; severity: number; category: string }> {
    
    // Sexual content detection
    const sexualIndicators = this.detectSexualContent(message, age);
    if (sexualIndicators.detected) {
      await this.handleDetection(studentId, message, sexualIndicators);
      return sexualIndicators;
    }
    
    // Violence indicators
    const violenceIndicators = this.detectViolenceIndicators(message, age);
    if (violenceIndicators.detected) {
      await this.handleDetection(studentId, message, violenceIndicators);
      return violenceIndicators;
    }
    
    return { detected: false, severity: 0, category: 'none' };
  }
  
  private detectSexualContent(
    message: string, 
    age: number
  ): { detected: boolean; severity: number; category: string } {
    
    const highSeverityPatterns = [
      /\b(sexual act|sexual abuse|touched me|made me)\b/i,
    ];
    
    const moderateSeverityPatterns = [
      /\b(inappropriate touch|uncomfortable|scared of)\b/i,
    ];
    
    const ageThreshold = this.getAgeAppropriateThreshold(age);
    
    for (const pattern of highSeverityPatterns) {
      if (pattern.test(message)) {
        return { detected: true, severity: 4, category: 'sexual' };
      }
    }
    
    for (const pattern of moderateSeverityPatterns) {
      if (pattern.test(message)) {
        const severity = age < ageThreshold ? 3 : 2;
        return { detected: true, severity, category: 'sexual' };
      }
    }
    
    return { detected: false, severity: 0, category: 'none' };
  }
  
  private detectViolenceIndicators(
    message: string,
    age: number
  ): { detected: boolean; severity: number; category: string } {
    
    const immediateThreatPatterns = [
      /\b(going to hurt|going to kill|have a plan|get a weapon)\b/i,
      /\b(tonight|tomorrow|after school) .*(hurt|kill|attack)\b/i,
    ];
    
    const severeViolencePatterns = [
      /\b(want to hurt|want to kill|hate .* want .* dead)\b/i,
      /\b(hit|punch|stab|shoot) .*(specific person|name)\b/i,
    ];
    
    for (const pattern of immediateThreatPatterns) {
      if (pattern.test(message)) {
        return { detected: true, severity: 4, category: 'violence' };
      }
    }
    
    for (const pattern of severeViolencePatterns) {
      if (pattern.test(message)) {
        return { detected: true, severity: 3, category: 'violence' };
      }
    }
    
    return { detected: false, severity: 0, category: 'none' };
  }
  
  private async handleDetection(
    studentId: string,
    message: string,
    indicators: { detected: boolean; severity: number; category: string }
  ): Promise<void> {
    
    const traumaFlag: TraumaFlag = {
      timestamp: new Date(),
      severity: indicators.severity as 1 | 2 | 3 | 4,
      category: indicators.category as any,
      content: message,
      aiResponse: this.generateSafeguardingResponse(indicators.severity),
      humanReviewed: false
    };
    
    await this.studentCollection.updateOne(
      { studentId },
      { 
        $push: { traumaIndicators: traumaFlag },
        $set: { 
          safeguardingStatus: indicators.severity >= 3 ? 'escalated' : 'monitoring',
          lastInteraction: new Date()
        }
      }
    );
    
    await this.escalateToSafeguarding(studentId, traumaFlag);
    
    if (indicators.severity === 4) {
      await this.alertEmergencyServices(studentId, traumaFlag);
    }
  }
  
  private generateSafeguardingResponse(severity: number): string {
    if (severity >= 4) {
      return "I'm very concerned about what you've shared. Your safety is the most important thing right now. I'm connecting you with someone who can help immediately. Please stay with me.";
    }
    
    if (severity >= 3) {
      return "Thanks for sharing that with me. I think it would be really helpful for you to talk with someone who specializes in these situations. I'm going to connect you with a support person who can help.";
    }
    
    return "I hear you. Let's take a break from the work for now. I'm going to make sure you get some support.";
  }
  
  private async escalateToSafeguarding(
    studentId: string, 
    traumaFlag: TraumaFlag
  ): Promise<void> {
    const alert = {
      studentId,
      timestamp: traumaFlag.timestamp,
      severity: traumaFlag.severity,
      category: traumaFlag.category,
      content: traumaFlag.content,
      urgent: traumaFlag.severity >= 3
    };
    
    try {
      await fetch(this.safeguardingAlertEndpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(alert)
      });
    } catch (error) {
      console.error('Failed to escalate to safeguarding:', error);
    }
  }
  
  private async alertEmergencyServices(
    studentId: string,
    traumaFlag: TraumaFlag
  ): Promise<void> {
    console.log(`EMERGENCY ALERT: Student ${studentId} - immediate danger detected`);
  }
  
  private getAgeAppropriateThreshold(age: number): number {
    if (age < 10) return 10;
    if (age < 13) return 13;
    return 16;
  }
}