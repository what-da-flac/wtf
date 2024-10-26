import { toDate } from '../helpers/text_utils';

export class UserReward {
  id: string;
  created_at: Date;
  amount: number;
  balance: number;
  op: string;

  constructor() {
    this.id = '';
    this.created_at = new Date();
    this.amount = 0;
    this.op = '';
    this.balance = 0;
  }
}

export function toUserReward(v: any): UserReward {
  if (!v) return new UserReward();
  return {
    ...v,
    created_at: toDate(new Date(v['created_at'])),
  };
}
