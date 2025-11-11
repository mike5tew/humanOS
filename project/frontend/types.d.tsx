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
    
    const timestamp = parseInt(parts[1]);
    const now = Date.now();
    const fiveMinutes = 5 * 60 * 1000;
    
    return (now - timestamp) < fiveMinutes;
  }
}