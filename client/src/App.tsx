import { useQuery } from "@tanstack/react-query"
import ConnectionStatus from "./components/ConnectionStatus"
import PlantsList from "./components/PlantsList"
import Time from "./components/Time"
import { getLatest } from "./api"

const App = () => {
  const { isPending, error, data, isFetching } = useQuery({
    queryKey: ['readings'],
    queryFn: getLatest,
    refetchInterval: 5 * 60 * 1000
  })

  return (
    <main>
      <Time />
      {data &&
        <>
          <ConnectionStatus connected={!error} lastUpdateTime={data.fetchTime} />
          <PlantsList readings={data.data} />
        </>
      }
    </main>
  )
}

export default App
