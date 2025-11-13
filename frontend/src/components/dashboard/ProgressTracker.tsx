import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { coachAPI } from '../../services/api';

interface ProgressTrackerProps {
  studentId: string;
}

export default function ProgressTracker({ studentId }: ProgressTrackerProps) {
  const { data: profile, isLoading, error } = useQuery({
    queryKey: ['student', studentId, 'profile'],
    queryFn: () => coachAPI.getStudentProfile(studentId),
    refetchInterval: 5000, // Refresh every 5 seconds
  });

  if (isLoading) {
    return <div className="progress-tracker loading">Loading profile...</div>;
  }

  if (error) {
    return (
      <div className="progress-tracker error">
        Failed to load profile: {error instanceof Error ? error.message : 'Unknown error'}
      </div>
    );
  }

  if (!profile) {
    return <div className="progress-tracker">No profile data</div>;
  }

  return (
    <div className="progress-tracker">
      <h2>Your Progress</h2>

      <div className="brain-state">
        <h3>Current State</h3>
        <div className="gauge">
          <label>Emotional Level</label>
          <progress value={profile.brain_state.emotional_level} max={1} />
          <span className="percentage">
            {Math.round(profile.brain_state.emotional_level * 100)}%
          </span>
        </div>
        <div className="gauge">
          <label>Rational Capacity</label>
          <progress value={profile.brain_state.rational_level} max={1} />
          <span className="percentage">
            {Math.round(profile.brain_state.rational_level * 100)}%
          </span>
        </div>
      </div>

      {profile.active_barriers.length > 0 && (
        <div className="barriers">
          <h3>Active Barriers</h3>
          <ul>
            {profile.active_barriers.map((barrier: string) => (
              <li key={barrier}>{barrier}</li>
            ))}
          </ul>
        </div>
      )}

      <div className="rewards">
        <h3>Play Break Progress</h3>
        <p>
          <strong>Current Stage:</strong> {profile.play_break_stage}
        </p>
        <p>
          <strong>Rewards Earned:</strong> {profile.rewards_earned}
        </p>
      </div>

      <div className="last-interaction">
        <p className="timestamp">
          Last activity: {new Date(profile.last_interaction).toLocaleTimeString()}
        </p>
      </div>
    </div>
  );
}
