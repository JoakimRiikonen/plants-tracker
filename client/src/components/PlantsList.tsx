import styles from './PlantsList.module.scss'

const PlantsList = () => {
  const plants: Plant[] = [
    {
      id: "123",
      name: "CHILI",
      status: "ok",
      moisture: 99.9
    },
    {
      id: "123",
      name: "SUNFLO",
      status: "warn",
      moisture: 23.3
    },
    {
      id: "123",
      name: "LILY",
      status: "alert",
      moisture: 10.8
    },
    {
      id: "123",
      name: "BASIL",
      status: "lost",
      moisture: 10.8
    },
  ]

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
          {plants.map(p => (
            <tr className={styles.plantRow}>
              <td className={styles.plantName}>{p.name}</td>
              <td className={styles.plantStatus + " " + styles[p.status]}>{p.status.toUpperCase()}</td>
              <td className={styles.plantMoisture}>{p.moisture}%</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

type Plant = {
  id: string,
  name: string,
  status: Status,
  moisture: number,
}

type Status = "ok" | "alert" | "warn" | "lost"

export default PlantsList;