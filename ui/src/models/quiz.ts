import { Campaign, toCampaign } from './campaign';
import { trimAll } from '../helpers/text_utils';

export enum QuizStatus {
  draft,
  published,
}

export class Quiz {
  id: string;
  name: string;
  video_url: string;
  campaign: Campaign;
  number_of_questions: number;
  status: string;
  youtube_video_id: string;
  thumbnail_url: string;
  reward_amount: number;
  expired: boolean;
  user_count: number;
  max_user_count: number;
  referral_amount: number;
  webhook_url: string;
  webhook_token: string;
  webhook_token_header_name: string;
  reference_id: string;

  constructor() {
    this.id = '';
    this.name = '';
    this.video_url = '';
    this.campaign = new Campaign();
    this.number_of_questions = 1;
    this.status = '';
    this.youtube_video_id = '';
    this.thumbnail_url = '';
    this.reward_amount = 0;
    this.expired = false;
    this.user_count = 0;
    this.max_user_count = 0;
    this.referral_amount = 0;
    this.webhook_url = '';
    this.webhook_token = '';
    this.webhook_token_header_name = 'x-webhook-token';
    this.reference_id = '';
  }
}

export function toQuiz(v: any): Quiz {
  if (!v) return new Quiz();
  return {
    ...v,
    campaign: toCampaign(v.campaign),
  };
}

export function quizValidation() {
  return {
    name: (value: string) =>
      trimAll(value).length === 0 ? 'Nombre es mandatorio' : null,
    video_url: (value: string) =>
      trimAll(value).length === 0 ? 'Video es mandatorio' : null,
    number_of_questions: (value: number) =>
      value <= 0 ? 'Minimo numero de preguntas es 1' : null,
    reward_amount: (value: number) =>
      value <= 0 ? 'El monto de recompensa debe ser mayor a cero' : null,
    max_user_count: (value: number) =>
      value <= 0
        ? 'El numero de usuarios que responden la encuesta debe ser mayor a cero'
        : null,
    referral_amount: (value: number) =>
      value <= 0
        ? 'El monto de recompensa para usuarios referenciados, debe ser mayor a cero'
        : null,
    webhook_url: (value: string) => {
      if (!value) return null;
      try {
        const uri = new URL(value);
        if (uri.protocol !== 'https:') {
          return 'protocol should be https';
        }
        return null;
      } catch (e) {
        return 'invalid url';
      }
    },
    webhook_token: (value: string) => null,
    webhook_token_header_name: (value: string) => null,
  };
}
