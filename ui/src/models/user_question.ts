import { trimAll } from '../helpers/text_utils';
import { toUserQuiz, UserQuiz } from './user_quiz';
import { Question } from './question';
import { UserAnswer } from './user_answer';

export class UserQuestion {
  id: string;
  has_correct_answer: boolean;
  user_quiz: UserQuiz;
  question: Question;

  constructor() {
    this.id = '';
    this.user_quiz = new UserQuiz();
    this.has_correct_answer = false;
    this.question = new Question();
  }
}

export function toUserQuestion(v: any): UserQuestion {
  if (!v) return new UserQuestion();
  return {
    ...v,
    question: toUserQuiz(v.user_quiz),
  };
}

export function questionValidation() {
  return {
    body: (value: string) =>
      trimAll(value).length === 0 ? 'Texto es mandatorio' : null,
    type: (value: string) =>
      trimAll(value).length === 0 ? 'Tipo es mandatorio' : null,
  };
}

export type UserNextQuestion = {
  user_question: UserQuestion | undefined;
  user_answers: UserAnswer[] | undefined;
  user_quiz: UserQuiz;
};
