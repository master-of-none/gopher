export type ShortenRequest = {
  url: string;
};

export type ShortenResponse = {
  short_url: string;
  code: string;
};

export type ErrorResponse = {
  error: string;
};
