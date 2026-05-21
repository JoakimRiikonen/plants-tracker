import { useEffect, useState } from "react";
import { Temporal } from "temporal-polyfill";
import styles from './Time.module.scss';

const Time = () => {
  const [time, setTime] = useState<Temporal.ZonedDateTime>();

  useEffect(() => {
    const timer = setInterval(() => {
      let time = Temporal.Now.zonedDateTimeISO();
      setTime(time);
    }, 1000);

    return () => {
      clearInterval(timer);
    };
  }, [])

  const formatPrimary = () => {
    const hours = time?.hour.toString().padStart(2, "0");
    const minutes = time?.minute.toString().padStart(2, "0");
    const seconds = time?.second.toString().padStart(2, "0");
    return `${hours}:${minutes}:${seconds}`
  }

  const requestFullScreen = () => {
    document.body.requestFullscreen()
  }

  const monthToText = (month?: number) => {
    switch (month) {
      case 1:
        return "Jan"
      case 2:
        return "Feb"
      case 3:
        return "Mar"
      case 4:
        return "Apr"
      case 5:
        return "May"
      case 6:
        return "June"
      case 7:
        return "July"
      case 8:
        return "Aug"
      case 9:
        return "Sept"
      case 10:
        return "Oct"
      case 11:
        return "Nov"
      case 12:
        return "Dec"
    }

    return "";
  }

  return (
    <div className={styles.timeContainer}>
      <div className={styles.timePrimary} onClick={() => requestFullScreen()}>
        {formatPrimary()}
      </div>
      <div className={styles.timeSecondary}>
        {monthToText(time?.month)} {time?.day.toString().padStart(2, "0")} {time?.year}
      </div>
    </div>
  )
}

export default Time;
