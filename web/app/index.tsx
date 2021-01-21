
import { Button, Fabric } from '@fluentui/react'
import * as React from 'react'
import { render } from 'react-dom'


const IndexPage = () => {
    return (
        <div>
            <Fabric>
                <h2> Welcome to PakkretQC 💁🏼‍♂️</h2>
                <Button as="a" href="/login">Login</Button>
            </Fabric>
        </div>
    )
}

render(<IndexPage />, document.getElementById("pakkretqc-root"))




