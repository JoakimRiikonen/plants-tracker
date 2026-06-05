import { useQuery } from "@tanstack/react-query"
import ConnectionStatus from "./components/ConnectionStatus"
import PlantsList from "./components/PlantsList"
import Time from "./components/Time"
import { getLatest } from "./api"
import { useState } from "react"
import PlantDetails from "./components/PlantDetails"

const App = () => {
  const { error, data, refetch } = useQuery({
    queryKey: ['readings'],
    queryFn: getLatest,
    refetchInterval: 5 * 60 * 1000
  })

  const [details, setDetails] = useState<string | undefined>();

  return (
    <main>
      <Time />
      {data &&
        <>
          <ConnectionStatus connected={!error} lastUpdateTime={data.fetchTime} />
          <PlantsList readings={data.data} setDetails={(s) => setDetails(s)} />
          {details &&
            <PlantDetails
              reading={data.data.find(r => r.SensorId == details)}
              onClose={() => setDetails(undefined)}
              refetchData={() => refetch()}
            />
          }
        </>
      }
    </main>
  )
}

export default App
