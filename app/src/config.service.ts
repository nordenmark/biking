export const ConfigService = new (class {
  get googleMapsApiKey(): string {
    const url = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
    if (!url) {
      throw new Error("Missing env variable VITE_GOOGLE_MAPS_API_KEY");
    }
    return url as string;
  }
})();
