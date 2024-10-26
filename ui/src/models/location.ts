import { Address } from './address';

export class Location {
  id: string;
  lat: number;
  lng: number;
  position: number;
  address: Address;

  constructor() {
    this.id = '';
    this.lat = 0;
    this.lng = 0;
    this.position = 0;
    this.address = new Address();
  }

  url(): string {
    return `https://maps.google.com/?q=${this.lat},${this.lng}`;
  }
}

export function toLocation(v: any): Location {
  const c = new Location();
  if (!v) return c;
  c.id = v.id;
  c.lng = v.lng;
  c.position = v.position;
  c.lat = v.lat;
  c.address = v.address;
  return c;
}
