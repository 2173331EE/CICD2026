import type { BoatClass } from './boatClasses/Class/types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';

export async function loadBoatClassesFromApi(): Promise<BoatClass[]> {
  const response = await fetch(`${API_BASE_URL}/api/classes`);

  if (!response.ok) {
    throw new Error(`Request failed with status ${response.status}`);
  }

  return (await response.json()) as BoatClass[];
}
