import { FC, useEffect, useRef } from "react";

interface MapProps {
  center?: google.maps.LatLngLiteral;
  zoom?: number;
}

export const Map: FC<MapProps> = ({ center, zoom }) => {
  const ref = useRef(null);

  useEffect(() => {
    if (ref.current) {
      new window.google.maps.Map(ref.current, {
        center,
        zoom,
      });
    }
  }, []);

  return <div ref={ref} style={{ width: "100vw", height: "100vh" }} />;
};

Map.defaultProps = {
  center: { lat: 59.8332748, lng: 17.623428 },
  zoom: 13,
};
