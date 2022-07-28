import { FC, useEffect, useRef } from "react";
import { Session } from "../types";

interface MapProps {
  sessions: Session[];
  center?: google.maps.LatLngLiteral;
  zoom?: number;
}

export const Map: FC<MapProps> = ({ sessions, center, zoom }) => {
  const ref = useRef(null);

  useEffect(() => {
    if (ref.current) {
      const map = new window.google.maps.Map(ref.current, {
        center,
        zoom,
      });

      const points: google.maps.LatLng[] = [];
      const bounds = new google.maps.LatLngBounds();

      sessions.forEach((session) => {
        session.points.forEach(({ lat, lng }) => {
          const point = new google.maps.LatLng(lat, lng);
          points.push(point);
          bounds.extend(point);
        });
      });

      const polyline = new google.maps.Polyline({
        path: points,
        strokeColor: "#FF0000",
        strokeOpacity: 0.7,
        strokeWeight: 3,
      });
      polyline.setMap(map);

      map.fitBounds(bounds);
    }
  }, []);

  return <div ref={ref} style={{ width: "100vw", height: "100vh" }} />;
};

Map.defaultProps = {
  center: { lat: 59.8332748, lng: 17.623428 },
  zoom: 2,
};
