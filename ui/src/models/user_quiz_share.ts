import { Quiz, toQuiz } from './quiz';
import { toUser, User } from './user';

export class UserQuizShare {
  created_at: Date;
  expires_at: Date;
  id: string;
  quiz: Quiz;
  token: string;
  used_at: Date;
  user_referred: User;
  user_shared: User;

  constructor() {
    this.id = '';
    this.created_at = new Date();
    this.expires_at = new Date();
    this.quiz = new Quiz();
    this.token = '';
    this.used_at = new Date();
    this.user_referred = new User();
    this.user_shared = new User();
  }
}

export function toUserQuizShare(v: any): User {
  if (!v) return new User();
  return {
    ...v,
    created_at: new Date(v['created_at']),
    expires_at: new Date(v['expires_at']),
    used_at: new Date(v['used_at']),
    quiz: toQuiz(v.quiz),
    user_referred: toUser(v.user_referred),
    user_shared: toUser(v.user_shared),
  };
}
