import axios from "axios";

const serviceAxiosApi = axios.create({
  baseURL: '/elevation-api/v1',
  timeout: 10000,
  headers: { 'content-type': 'application/json' }
});

/**
 * Fetches elevations for an array of [lat, lng] coordinates.
 * The input must be in [latitude, longitude] order.
 */
export async function getElevations(coords: [number, number][]): Promise<number[]|null[]> {
  if (!coords.length) return [];
  // Downsample to 100 points for API limits
  const maxPoints = 100;
  let usedCoords = coords;
  if (coords.length > maxPoints) {
    const step = Math.ceil(coords.length / maxPoints);
    usedCoords = coords.filter((_, i) => i % step === 0);
  }
  const locations = usedCoords.map(([lat, lng]) => `${lat},${lng}`).join("|");
  const url = `/eudem25m?locations=${locations}`;
  try {
    const res = await serviceAxiosApi.get(url);
    const data = res.data;
    if (data.results) {
      return data.results.map((r: any) => r.elevation ?? null);
    }
    return [];
  } catch (e) {
    return usedCoords.map(() => null);
  }
}
