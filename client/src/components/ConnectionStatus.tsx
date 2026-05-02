import { Temporal } from "temporal-polyfill";
import styles from './ConnectionStatus.module.scss';

const ConnectionStatus = () => {
  const time = Temporal.Now.zonedDateTimeISO()
  const connected = true;

  return (
    <div className={styles.statusContainer}>
      <div className={styles.currentStatusContainer}>
        <div className={styles.statusLabel}>
          System state:
        </div>
        <div className={styles.status + " " + (connected ? styles.connected : styles.disconnected)}>
          {connected ? "SEEN" : "LOST"}
        </div>
      </div>
      <div>
        <div className={styles.statusLastUpdateLabel}>
          Last update:
        </div>
        <div>
          {time.toString().split('+')[0]}
        </div>
        <div>
        </div>
      </div>
    </div>
  )
}

export default ConnectionStatus;