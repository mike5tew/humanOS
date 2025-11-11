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
  
  /**
   * Detect interests from student message
   */
  async detectInterests(studentId: string, message: string): Promise<Interest[]> {
    const detectedInterests: Interest[] = [];
    
    // Interest patterns (expandable)
    const interestPatterns = {
      games: [
        { pattern: /minecraft/i, specific: 'Minecraft' },
        { pattern: /fortnite/i, specific: 'Fortnite' },
        { pattern: /roblox/i, specific: 'Roblox' },
        { pattern: /among us/i, specific: 'Among Us' },
        // Add more gaming patterns
      ],
      sports: [
        { pattern: /football|soccer/i, specific: 'Football' },
        { pattern: /basketball/i, specific: 'Basketball' },
        { pattern: /swimming/i, specific: 'Swimming' },
        // Add more sports
      ],
      hobbies: [
        { pattern: /drawing|art/i, specific: 'Drawing' },
        { pattern: /music|guitar|piano/i, specific: 'Music' },
        { pattern: /reading|books/i, specific: 'Reading' },
        // Add more hobbies
      ],
      subjects: [
        { pattern: /space|astronomy|planets/i, specific: 'Space' },
        { pattern: /dinosaur/i, specific: 'Dinosaurs' },
        { pattern: /animals|wildlife/i, specific: 'Animals' },
        // Add more subjects
      ],
      media: [
        { pattern: /youtube/i, specific: 'YouTube' },
        { pattern: /tiktok/i, specific: 'TikTok' },
        { pattern: /netflix|series|show/i, specific: 'TV Shows' },
        // Add more media
      ]
    };
    
    // Check each pattern
    for (const [category, patterns] of Object.entries(interestPatterns)) {
      for (const { pattern, specific } of patterns) {
        if (pattern.test(message)) {
          detectedInterests.push({
            category,
            specific,
            confidence: 0.8, // Adjust based on context
            lastMentioned: new Date(),
            mentionCount: 1
          });
        }
      }
    }
    
    // Update MongoDB
    if (detectedInterests.length > 0) {
      await this.updateStudentInterests(studentId, detectedInterests);
    }
    
    // Store in Weaviate for semantic search
    await this.storeInterestsInWeaviate(studentId, message, detectedInterests);
    
    return detectedInterests;
  }
  
  /**
   * Update student interests in MongoDB
   */
  private async updateStudentInterests(studentId: string, newInterests: Interest[]): Promise<void> {
    const student = await this.mongoCollection.findOne({ studentId });
    
    if (!student) return;
    
    const existingInterests = student.interests || [];
    
    for (const newInterest of newInterests) {
      const existing = existingInterests.find(
        i => i.category === newInterest.category && i.specific === newInterest.specific
      );
      
      if (existing) {
        // Update existing interest
        existing.mentionCount++;
        existing.lastMentioned = new Date();
        existing.confidence = Math.min(existing.confidence + 0.1, 1.0);
      } else {
        // Add new interest
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
  
  /**
   * Store interests in Weaviate for semantic search
   */
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
  
  /**
   * Generate personalized responses using interests
   */
  async generatePersonalizedResponse(
    studentId: string,
    taskType: string
  ): Promise<string> {
    const student = await this.mongoCollection.findOne({ studentId });
    
    if (!student || !student.interests || student.interests.length === 0) {
      return this.getGenericResponse(taskType);
    }
    
    // Get recent interests (avoid repetition)
    const recentInterests = student.interests
      .sort((a, b) => b.lastMentioned.getTime() - a.lastMentioned.getTime())
      .slice(0, 5);
    
    // Find least recently used interest
    const leastRecentInterest = recentInterests[recentInterests.length - 1];
    
    // Check if we've used this recently
    const recentUsage = await this.checkRecentInterestUsage(studentId, leastRecentInterest.specific);
    
    if (recentUsage < 3) { // Used less than 3 times recently
      return this.personalizeWithInterest(taskType, leastRecentInterest);
    }
    
    // Fall back to generic if overused
    return this.getGenericResponse(taskType);
  }
  
  /**
   * Check how many times an interest was used in recent interactions
   */
  private async checkRecentInterestUsage(
    studentId: string, 
    interest: string
  ): Promise<number> {
    // Query Weaviate for recent interactions mentioning this interest
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
  
  /**
   * Personalize response with interest
   */
  private personalizeWithInterest(taskType: string, interest: Interest): string {
    const templates = {
      gameReward: [
        `Complete this and you'll get 5 minutes on ${interest.specific}!`,
        `Let's knock this out so you can play ${interest.specific}.`,
        `Finish this task = ${interest.specific} time. Deal?`
      ],
      taskFraming: [
        `Think of this like ${interest.specific} - you need to figure out the strategy.`,
        `This is like leveling up in ${interest.specific} - just need to complete this challenge.`,
        `Remember how you solved that puzzle in ${interest.specific}? Same thinking here.`
      ],
      encouragement: [
        `You've got this - you tackle way harder stuff in ${interest.specific}!`,
        `If you can master ${interest.specific}, you can handle this.`,
        `Apply that ${interest.specific} focus here and you'll crush it.`
      ]
    };
    
    // Select appropriate template based on task type
    const templateCategory = taskType === 'reward' ? 'gameReward' : 
                             taskType === 'challenge' ? 'taskFraming' : 'encouragement';
    
    const options = templates[templateCategory];
    return options[Math.floor(Math.random() * options.length)];
  }
  
  private getGenericResponse(taskType: string): string {
    // Generic responses when no interests available
    return "Let's tackle this together. Give it a try and see what you can do!";
  }
}