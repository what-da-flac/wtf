export const HttpStatusBadRequest = 400;

export class AppError extends Error {
  message: string;
  name: string;
  statusCode: number;

  constructor(message: string, name: string = 'no name', statusCode?: number) {
    super(message);
    this.message = message;
    this.name = name;
    this.statusCode = statusCode || 0;
  }
}
