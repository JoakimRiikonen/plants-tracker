import type { Reading } from '../models/reading'
import styles from './PlantsList.module.scss'

type PlantsListProps = {
  readings: Reading[]
  setDetails: (sensorId: string) => void
}

type Status = "ok" | "alert" | "warn" | "lost"

const PlantsList = ({ readings, setDetails }: PlantsListProps) => {
  // TODO: This should be moved to the api, and should probably be configurable per plant
  function getStatus(reading: Reading): Status {
    if (reading.Moisture === -1) {
      return "lost"
    }
    if (reading.Moisture > 30) {
      return "ok"
    }
    if (reading.Moisture > 20) {
      return "warn"
    }

    return "alert"
  }

  return (
    <div>
      <table className={styles.plantsTable}>
        <thead>
          <tr className={styles.plantsTableHeader}>
            <th>Plants</th>
            <th>Status</th>
            <th>Moisture</th>
          </tr>
        </thead>
        <tbody>
          {readings.map(p => (
            <tr className={styles.plantRow} onClick={() => setDetails(p.SensorId)}>
              <td className={styles.plantName}>{p.SensorName}</td>
              <td className={styles.plantStatus + " " + styles[getStatus(p)]}>{getStatus(p).toUpperCase()}</td>
              <td className={styles.plantMoisture}>{p.Moisture}%</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default PlantsList;