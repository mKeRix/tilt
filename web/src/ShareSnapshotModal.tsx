import React from "react"
import { Snapshot } from "./types"
import "./ShareSnapshotModal.scss"

type props = {
  handleSendSnapshot: (snapshot: Snapshot) => void
  handleClose: () => void
  show: boolean
  state: Snapshot
  snapshotUrl: string
  registerTokenUrl: string
}
export default function ShareSnapshotModal(props: props) {
  const showHideClassName = props.show
    ? "modal display-block"
    : "modal display-none"
  const hasLink = props.snapshotUrl !== ""
  if (!hasLink) {
    return (
      <div className={showHideClassName}>
        <section className="modal-main">
          <section className="modal-snapshotUrlWrap">
            <button onClick={props.handleClose}>close</button>
            <button onClick={() => props.handleSendSnapshot(props.state)}>
              Share Snapshot
            </button>
          </section>
          <hr />
          {loggedOutTiltCloudCopy(props.registerTokenUrl)}
        </section>
      </div>
    )
  }

  return (
    <div className={showHideClassName}>
      <section className="modal-main">
        <section className="modal-snapshotUrlWrap">
          <button onClick={props.handleClose}>close</button>
          <button onClick={() => window.open(props.snapshotUrl)}>
            Open link in new tab
          </button>
        </section>
        <hr />
        {loggedOutTiltCloudCopy("")}
      </section>
    </div>
  )
}

const loggedOutTiltCloudCopy = (registerTokenURL: string) => (
  <section className="modal-cloud">
    <p>
      Go to <a href={registerTokenURL}>TiltCloud</a> to manage your snapshots.
      <br />
      (You'll just need to link your GitHub Account)
    </p>
  </section>
)