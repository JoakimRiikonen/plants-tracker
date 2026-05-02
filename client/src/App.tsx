import ConnectionStatus from "./components/ConnectionStatus"
import PlantsList from "./components/PlantsList"
import Time from "./components/Time"

const App = () => {
  return (
    <main>
      <Time />
      <ConnectionStatus />
      <PlantsList />
    </main>
  )
}

export default App
