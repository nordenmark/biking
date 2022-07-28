import { Status, Wrapper } from "@googlemaps/react-wrapper";
import "./App.css";
import { Map } from "./components/Map";
import { ConfigService } from "./config.service";

const render = (status: Status) => {
  switch (status) {
    case Status.LOADING:
      return <p>loading...</p>;
    case Status.FAILURE:
      return <p>error</p>;
    case Status.SUCCESS:
      return <Map />;
  }
};

function App() {
  return (
    <div className="App">
      <Wrapper apiKey={ConfigService.googleMapsApiKey} render={render}>
        <Map />
      </Wrapper>
    </div>
  );
}

export default App;
