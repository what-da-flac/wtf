import { popupError, popupWarning } from './Notifier';
import { HttpStatusBadRequest } from '../models/errors';

export function notifyErrResponse(err) {
  let statusCode;
  let message;
  const { response } = err;
  // statusCode may come from axios response or an AppError object
  statusCode = response?.status || err.statusCode;
  message = response?.data?.error || err.message;
  if (!statusCode) {
    window.alert(message);
    return;
  }
  switch (statusCode) {
    case HttpStatusBadRequest:
      return popupWarning({
        title: 'error',
        text: message,
        timer: 3000,
      });
    default:
      return popupError({
        title: 'error',
        text: message,
      });
  }
}
