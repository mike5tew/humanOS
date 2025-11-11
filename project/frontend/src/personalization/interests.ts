import { MongoClient, Collection } from 'mongodb';
import weaviate, { WeaviateClient } from 'weaviate-ts-client';
import { StudentProfile, Interest, Interaction } from '../database/schemas';

export class InterestDetectionSystem {
  private mongoCollection: Collection<StudentProfile>;
  private weaviateClient: WeaviateClient;
  
  constructor(mongoClient: MongoClient, weaviateClient: WeaviateClient) {
    this.mongoCollection = mongoClient.db('humanos').collection<StudentProfile>('students');
    this.weaviateClient = weaviateClient;
  }
  
  async detectInterests(studentId: string, message: string): Promise<Interest[]> {
    const detectedInterests: Interest[] = [];
    
    const interestPatterns = {
      games: [
        { pattern: /minecraft/i, specific: 'Minecraft' },
        { pattern: /fortnite/i, specific: 'Fortnite' },
        { pattern: /roblox/i, specific: 'Roblox' },
      ],
      sports: [
        { pattern: /football|soccer/i, specific: 'Football' },
        { pattern: /basketball/i, specific: 'Basketball' },
      ],
      hobbies: [
        { pattern: /drawing|art/i, specific: 'Drawing' },
        { pattern: /music|guitar|piano/i, specific: 'Music' },
      ],
    };
    
    for (const [category, patterns] of Object.entries(interestPatterns)) {
      for (const { pattern, specific } of patterns) {
        if (pattern.test(message)) {
          detectedInterests.push({
            category,
            specific,
            confidence: 0.8,
            lastMentioned: new Date(),
            mentionCount: 1
          });
        }
      }
    }
    
    if (detectedInterests.length > 0) {
      await this.updateStudentInterests(studentId, detectedInterests);
    }
    
    await this.storeInterestsInWeaviate(studentId, message, detectedInterests);
    
    return detectedInterests;
  }
  
  private async updateStudentInterests(studentId: string, newInterests: Interest[]): Promise<void> {
    const student = await this.mongoCollection.findOne({ studentId });
    
    if (!student) return;
    
    const existingInterests = student.interests || [];
    
    for (const newInterest of newInterests) {
      const existing = existingInterests.find(
        i => i.category === newInterest.category && i.specific === newInterest.specific
      );
      
      if (existing) {
        existing.mentionCount++;
        existing.lastMentioned = new Date();
        existing.confidence = Math.min(existing.confidence + 0.1, 1.0);
      } else {
        existingInterests.push(newInterest);
      }
    }
    
    await this.mongoCollection.updateOne(
      { studentId },
      { 
        $set: { 
          interests: existingInterests,
          lastInteraction: new Date()
        } 
      }
    );
  }
  
  private async storeInterestsInWeaviate(
    studentId: string, 
    message: string, 
    interests: Interest[]
  ): Promise<void> {
    for (const interest of interests) {
      await this.weaviateClient.data
        .creator()
        .withClassName('StudentInterest')
        .withProperties({
          studentId,
          category: interest.category,
          specific: interest.specific,
          context: message,
          timestamp: new Date().toISOString()
        })
        .do();
    }
  }
  
  async generatePersonalizedResponse(
    studentId: string,
    taskType: string
  ): Promise<string> {
    const student = await this.mongoCollection.findOne({ studentId });
    
    if (!student || !student.interests || student.interests.length === 0) {
      return this.getGenericResponse(taskType);
    }
    
    const recentInterests = student.interests
      .sort((a, b) => b.lastMentioned.getTime() - a.lastMentioned.getTime())
      .slice(0, 5);
    
    const leastRecentInterest = recentInterests[recentInterests.length - 1];
    
    const recentUsage = await this.checkRecentInterestUsage(studentId, leastRecentInterest.specific);
    
    if (recentUsage < 3) {
      return this.personalizeWithInterest(taskType, leastRecentInterest);
    }
    
    return this.getGenericResponse(taskType);
  }
  
  private async checkRecentInterestUsage(
    studentId: string, 
    interest: string
  ): Promise<number> {
    const result = await this.weaviateClient.graphql
      .get()
      .withClassName('Interaction')
      .withFields('studentId content timestamp')
      .withWhere({
        operator: 'And',
        operands: [
          {
            path: ['studentId'],
            operator: 'Equal',
            valueString: studentId
          },
          {
            path: ['content'],
            operator: 'Like',
            valueString: `*${interest}*`
          }
        ]
      })
      .withLimit(10)
      .do();
    
    return result.data?.Get?.Interaction?.length || 0;
  }
  
  private personalizeWithInterest(taskType: string, interest: Interest): string {
    const templates = {
      gameReward: [
        `Complete this and you'll get 5 minutes on ${interest.specific}!`,
        `Let's knock this out so you can play ${interest.specific}.`,
      ],
      taskFraming: [
        `Think of this like ${interest.specific} - you need to figure out the strategy.`,
      ],
    };
    
    const templateCategory = taskType === 'reward' ? 'gameReward' : 'taskFraming';
    
    const options = templates[templateCategory];
    return options[Math.floor(Math.random() * options.length)];
  }
  
  private getGenericResponse(taskType: string): string {
    return "Let's tackle this together. Give it a try and see what you can do!";
  }
}