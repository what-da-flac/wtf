import { Question } from './question';
import { UserQuiz } from './user_quiz';
import { Answer } from './answer';

export class UserAnswer {
  id: string;
  is_completed: boolean;
  question: Question;
  user_quiz: UserQuiz;
  answer: Answer;

  constructor() {
    this.id = '';
    this.is_completed = false;
    this.question = new Question();
    this.user_quiz = new UserQuiz();
    this.answer = new Answer();
  }
}
