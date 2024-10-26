import dayjs from 'dayjs';

export function trimAll(s: any): any {
  switch (typeof s) {
    case 'string':
      return s.trim();
    case 'object':
      for (const [key] of Object.entries(s)) {
        s[key] = trimAll(s[key]);
      }
  }
}

export function truncateTime(utcDate: Date): Date {
  return new Date(utcDate.getFullYear(), utcDate.getMonth(), utcDate.getDate());
}

export function toUTCDate(utcDate: Date): Date {
  return new Date(
    utcDate.getUTCFullYear(),
    utcDate.getUTCMonth(),
    utcDate.getUTCDate(),
    utcDate.getUTCHours(),
    utcDate.getUTCMinutes(),
    utcDate.getUTCSeconds()
  );
}

export function toDate(utcDate: Date): Date {
  return new Date(
    utcDate.getFullYear(),
    utcDate.getMonth(),
    utcDate.getDate(),
    utcDate.getHours(),
    utcDate.getMinutes(),
    utcDate.getSeconds()
  );
}

export function formatDate(d: Date): string {
  return dayjs(d).format('YYYY-MM-DD');
}

export function trimBigText(d: string, maxLength: number = 50): string {
  return d.length >= maxLength ? `${d.substring(0, maxLength - 1)}...` : d;
}

export function capitalizeAllWords(text: string): string {
  return text.replace(/\b\w/g, char => char.toUpperCase());
}

export function trimAndCapitalize(d: string, maxLength: number = 50): string {
  const trimmedText =
    d.length >= maxLength ? `${d.substring(0, maxLength - 1)}...` : d;
  return trimmedText.replace(/\b\w/g, char => char.toUpperCase()); // Capitalize first letter of each word
}

export function formatCurrency(number: number, currency = 'USD') {
  const formatter = new Intl.NumberFormat(navigator.language, {
    style: 'currency',
    currency: currency,
  });

  return formatter.format(number);
}
