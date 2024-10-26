export class Movie {
  id: string;
  title: string;
  description: string;
  image_url: string;

  constructor() {
    this.id = '';
    this.title = '';
    this.description = '';
    this.image_url = '';
  }
}

export function toMovie(v: any): Movie {
  if (!v) return new Movie();
  return {
    ...v,
  };
}
