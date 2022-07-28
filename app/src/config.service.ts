export const ConfigService = new (class {
  get apiUrl(): string {
    const url = import.meta.env.VITE_API_URL;
    if (!url) {
      throw new Error("Missing env variable VITE_API_URL");
    }
    return url! as string;
  }

  get googleMapsApiKey(): string {
    const url = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
    if (!url) {
      throw new Error("Missing env variable VITE_GOOGLE_MAPS_API_KEY");
    }
    return url as string;
  }
})();
