import { Temporal } from "temporal-polyfill";
import type { Reading } from "./models/reading";

export async function getLatest(): Promise<{data: Reading[], fetchTime: Temporal.TimeZoneLike}> {
  const url = import.meta.env.VITE_API_BASE_URL + '/latest';
  const response = await fetch(url) 
  const json = await response.json()
  return {data: json, fetchTime: Temporal.Now.zonedDateTimeISO()}
}