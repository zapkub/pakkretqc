
import { Fabric, PrimaryButton, Text } from '@fluentui/react'
import * as React from 'react'
import { render } from 'react-dom'
import { MadeWithLove } from './components'


const IndexPage = () => {
    return (
        <div className="content">
            <Fabric>
                <div style={{ marginBottom: 16 }}>
                    <h1> Welcome to PakkretQC 💁🏼‍♂️</h1>
                    <Text>Another UI client of Microfocus Application Life Cycle Management. Build on top of ALM RESTful API <br />with modern Javascript UI and Golang.</Text>
                </div>
                <PrimaryButton as="a" href="/login">Getting Start</PrimaryButton>
                <MadeWithLove />

                <h1 id="pakkretqc-release-note">Release Note</h1>
                <h2 id="v0-0-2">v0.0.2 Wednesday, Feb 3, 2021</h2>
                <ul>
                    <li>✨ Support artifact file download.</li>
                    <li>✨ Support defect query filter.</li>
                    <li>🎨 Landing page styling update.</li>
                    <li>🎨 Improve defect detail page UI.</li>
                    <li>🐛 Fix unexpected crash from panic.</li>
                </ul>
                <h2 id="v0-0-1">v0.0.1</h2>
                <p>First release</p>

            </Fabric>
        </div>
    )
}

render(<IndexPage />, document.getElementById("pakkretqc-root"))




