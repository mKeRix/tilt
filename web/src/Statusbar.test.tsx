import React from "react"
import renderer from "react-test-renderer"
import Statusbar, { StatusItem } from "./Statusbar"
import { mount } from "enzyme"
import { twoResourceView } from "./testdata.test"
import { MemoryRouter } from "react-router"
import { TiltBuild } from "./types"

describe("StatusBar", () => {
  let runningVersion: TiltBuild = {
    version: "v0.8.1",
    date: "1970-01-01",
    dev: false,
  }
  it("renders without crashing", () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={[]}
            alertsUrl="/alerts"
            runningVersion={runningVersion}
            latestVersion={null}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })

  it("renders two items both errors", () => {
    let items = twoResourceView().resources.map((res: any) => {
      res.currentBuild = {}
      res.pendingBuildSince = ""
      return new StatusItem(res)
    })
    let statusbar = mount(
      <MemoryRouter>
        <Statusbar
          items={items}
          alertsUrl="/alerts"
          runningVersion={runningVersion}
          latestVersion={null}
        />
      </MemoryRouter>
    )
    expect(
      statusbar.find(".Statusbar-errWarnPanel-count--error").html()
    ).toContain("2")

    statusbar.unmount()
  })

  it("renders two items both errors snapshot", () => {
    let items = twoResourceView().resources.map((res: any) => {
      res.currentBuild = {}
      res.pendingBuildSince = ""
      return new StatusItem(res)
    })
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={items}
            alertsUrl="/alerts"
            runningVersion={runningVersion}
            latestVersion={null}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })

  it("renders two items all ok snapshot", () => {
    let view = twoResourceView()
    view.resources.forEach((res: any) => {
      res.buildHistory[0].error = ""
    })

    let items = view.resources.map((res: any) => new StatusItem(res))
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={items}
            alertsUrl="/alerts"
            runningVersion={runningVersion}
            latestVersion={null}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })

  it("renders two items all ok", () => {
    let view = twoResourceView()
    view.resources.forEach((res: any) => {
      res.buildHistory[0].error = ""
    })
    let items = view.resources.map((res: any) => new StatusItem(res))
    let statusbar = mount(
      <MemoryRouter>
        <Statusbar
          items={items}
          alertsUrl="/alerts"
          runningVersion={runningVersion}
          latestVersion={null}
        />
      </MemoryRouter>
    )
    expect(
      statusbar.find(".Statusbar-errWarnPanel-count--error").html()
    ).toContain("0")

    statusbar.unmount()
  })

  it("renders an upgrade badge when the version is out of date", () => {
    let view = twoResourceView()
    view.resources.forEach((res: any) => {
      res.buildHistory[0].error = ""
    })
    let items = view.resources.map((res: any) => new StatusItem(res))
    let latestVersion = runningVersion
    latestVersion.version = "10.0.0"
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={items}
            alertsUrl="/alerts"
            runningVersion={runningVersion}
            latestVersion={latestVersion}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })

  it("does not render an upgrade badge when there is no latestVersion", () => {
    let view = twoResourceView()
    view.resources.forEach((res: any) => {
      res.buildHistory[0].error = ""
    })
    let items = view.resources.map((res: any) => new StatusItem(res))
    let latestVersion = { version: "", date: "", dev: false }
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={items}
            alertsUrl="/alerts"
            runningVersion={runningVersion}
            latestVersion={latestVersion}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })

  it("does not render an upgrade badge when runningVersion is dev", () => {
    let view = twoResourceView()
    view.resources.forEach((res: any) => {
      res.buildHistory[0].error = ""
    })
    let items = view.resources.map((res: any) => new StatusItem(res))
    let latestVersion = runningVersion
    latestVersion.version = "10.0.0"
    let devRunningVersion = runningVersion
    devRunningVersion.dev = true
    const tree = renderer
      .create(
        <MemoryRouter>
          <Statusbar
            items={items}
            alertsUrl="/alerts"
            runningVersion={devRunningVersion}
            latestVersion={latestVersion}
          />
        </MemoryRouter>
      )
      .toJSON()

    expect(tree).toMatchSnapshot()
  })
})

describe("StatusItem", () => {
  it("can be constructed with no build history", () => {
    let si = new StatusItem({})
    expect(si.hasError).toBe(false)
  })
})
