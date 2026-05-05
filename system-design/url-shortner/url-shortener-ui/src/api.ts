import type { ShortenRequest, ShortenResponse, ErrorResponse } from "./types";

export async function shortenURl(
  payload: ShortenRequest,
  expiry?: number,
): Promise<ShortenResponse> {
  let endpoint = "http://localhost:8080/shorten";

  if (expiry) endpoint += `?expiry=${expiry}`;

  const res = await fetch(endpoint, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  if (res.status == 429) {
    const retryAfter = res.headers.get("Retry After");
    throw new Error(`Rate Limited. Retry After ${retryAfter}`);
  }

  const data = (await res.json()) as ShortenResponse | ErrorResponse;

  if (!res.ok) {
    const err = data as ErrorResponse;
    throw new Error(err.error || "Request Failed");
  }

  return data as ShortenResponse;
}
