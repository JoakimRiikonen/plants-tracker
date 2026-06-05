import { Temporal } from "temporal-polyfill";
import type { Reading } from "./models/reading";

export async function getLatest(): Promise<{data: Reading[], fetchTime: Temporal.TimeZoneLike}> {
  const url = import.meta.env.VITE_API_BASE_URL + '/latest';
  const response = await fetch(url) 
  const json = await response.json()
  return {data: json, fetchTime: Temporal.Now.zonedDateTimeISO()}
}

export async function setName(sensorId: string, sensorName: string): Promise<boolean> {
  const url = import.meta.env.VITE_API_BASE_URL + '/sensors/' + sensorId + '/setName'
  const response = await fetch(url, {
    method: 'POST',
    body: JSON.stringify({
      sensorName: sensorName
    })
  })

  return response.ok
}