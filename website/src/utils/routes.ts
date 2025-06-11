/**
 * Parses a WKT LINESTRING and returns an array of [lat, lng] pairs.
 */
export function parseRoute(route: string): [number, number][] {
  if (!route) return [];
  const match = route.match(/\((.*)\)/);
  if (!match) return [];
  return match[1].split(',').map(pair => {
    const [lng, lat] = pair.trim().split(' ').map(Number);
    return [lat, lng];
  });
}
