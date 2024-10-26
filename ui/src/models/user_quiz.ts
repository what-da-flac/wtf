import { Quiz, toQuiz } from './quiz';
import { toUser, User } from './user';
import { UserQuestion } from './user_question';

export enum UserQuizStatus {
  unknown,
  pending,
  started,
  success,
  failed,
}

export class UserQuiz {
  id: string;
  created_at: Date;
  updated_at: Date;
  user: User;
  quiz: Quiz;
  status: string;
  percent_completed: number;

  constructor() {
    this.id = '';
    this.created_at = new Date();
    this.updated_at = new Date();
    this.user = new User();
    this.quiz = new Quiz();
    this.status = '';
    this.percent_completed = 0;
  }
}

export function toUserQuiz(v: any): UserQuiz {
  if (!v) return new UserQuiz();
  return {
    ...v,
    created_at: new Date(v['created_at']),
    updated_at: new Date(v['updated_at']),
    user: toUser(v.user),
    quiz: toQuiz(v.quiz),
  };
}

export type UserQuizSummary = {
  user_questions: UserQuestion[];
  user_quiz: UserQuiz;
};
