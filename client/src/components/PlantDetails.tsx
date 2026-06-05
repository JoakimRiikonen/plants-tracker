import { useState } from 'react'
import type { Reading } from '../models/reading'
import styles from './PlantDetails.module.scss'
import { setName as setNameCall } from '../api'

type PlantDetailsProps = {
  reading: Reading | undefined
  onClose: () => void
  refetchData: () => void
}

const PlantDetails = ({ reading, onClose, refetchData }: PlantDetailsProps) => {
  const [changeName, setChangeName] = useState<boolean>(false);
  const [name, setName] = useState<string>("")

  if (!reading) {
    return <></>
  }

  function onNameChange() {
    if (!reading) {
      return
    }

    setNameCall(reading.SensorId, name)
      .then(() => {
        setChangeName(false);
        refetchData()
      })
  }

  return (
    <div className={styles.plantDetails}>
      <div className={styles.header}>
        <div>{reading.SensorName}</div>
        <button className={styles.closeButton} onClick={() => onClose()}>✕</button>
      </div>
      {changeName
        ? <>
          <div className={styles.changeNameContainer}>
            <input type='text' value={name} onChange={e => setName(e.target.value)} />
            <button type='button' onClick={() => onNameChange()}>Set name</button>
          </div>
        </>
        : <>
          <button className={styles.changeNameButton} onClick={() => setChangeName(true)}>CHANGE NAME</button>
        </>}
    </div>
  )
}

export default PlantDetails