import { Status, Wrapper } from "@googlemaps/react-wrapper";
import { useQuery } from "@tanstack/react-query";
import { FC } from "react";
import { Loading } from "./components/Loading";
import { Map } from "./components/Map";
import { ConfigService } from "./config.service";
import { Session } from "./types";

const render = (sessions: Session[]) => (status: Status) => {
  switch (status) {
    case Status.LOADING:
      return <Loading />;
    case Status.FAILURE:
      return <p>error</p>;
    case Status.SUCCESS:
      return <Map sessions={sessions} />;
  }
};

const fetchSessions = () =>
  fetch(`${ConfigService.apiUrl}/sessions`).then((response) => response.json());

export const Home: FC = () => {
  const { isLoading, data } = useQuery(["sessions"], fetchSessions);

  if (isLoading) {
    return <Loading />;
  }

  return (
    <Wrapper apiKey={ConfigService.googleMapsApiKey} render={render(data)} />
  );
};
