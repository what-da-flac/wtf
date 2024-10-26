import { trimAll } from '../helpers/text_utils';
import { Question, toQuestion } from './question';

export class Answer {
  id: string;
  body: string;
  is_valid: boolean;
  question: Question;
  reference_id: string;

  constructor() {
    this.id = '';
    this.body = '';
    this.is_valid = false;
    this.question = new Question();
    this.reference_id = '';
  }
}

export function toAnswer(v: any): Answer {
  if (!v) return new Answer();
  return {
    ...v,
    question: toQuestion(v.question),
  };
}

export function answerValidation() {
  return {
    body: (value: string) =>
      trimAll(value).length === 0 ? 'Texto es mandatorio' : null,
  };
}
