
import { Button, Fabric, Text } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';
import { DeflectTable } from './components';


interface ProjectPageData {

    total: number
    domain: string
    project: string
    defects: Deflect[]

    username: string

}

const ProjectPage = (props: { data: ProjectPageData }) => {
    return (
        <Fabric>
            <div>
                <Button as="a" href="/login">Logout</Button>
                <div className="ms-Grid">
                    <h2 className="ms-Grid-col"> Domain: {props.data.domain}, Project: {props.data.project} </h2>
                    
                    Welcome <Text style={{fontWeight: 'bold'}}>{props.data.username}</Text>
                </div>
            </div>
            <DeflectTable data={props.data.defects} />
        </Fabric>
    )
}



render(<ProjectPage data={window.__DATA__}></ProjectPage>, document.getElementById("pakkretqc-root"))