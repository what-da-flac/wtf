export class Address {
  address: string;
  city: string;
  country: string;
  country_short: string;
  full: string;
  state: string;
  street: string;
  street_number: string;
  zipcode: string;

  constructor() {
    this.address = '';
    this.city = '';
    this.country = '';
    this.country_short = '';
    this.full = '';
    this.state = '';
    this.street = '';
    this.street_number = '';
    this.zipcode = '';
  }
}

export function toAddress(v: any): Address {
  if (!v) return new Address();
  return {
    ...v,
  };
}
