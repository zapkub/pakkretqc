
import { Button, Fabric, Text } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';
import { DefectTable } from './components';


interface ProjectPageData {
    total: number
    domain: string
    project: string
    defects: Deflect[]
    username: string
}

const ProjectPage = (props: { data: ProjectPageData }) => {
    return (
        <Fabric className="project-page">
            <div className="canopy">
                <div className="ms-Grid">
                    <div style={{marginBottom: 8}}>
                        Welcome: <Text style={{ fontWeight: 'bold' }}>{props.data.username || 'Unknown'}</Text>
                    </div>
                    <h2 className="ms-Grid-col"> Domain: {props.data.domain}, Project: {props.data.project} </h2>
                </div>
                <div className="user-profile-box">
                    <Button as="a" href="/login">Logout</Button>
                </div>
            </div>
            <h2>🐛🐞 Defect list</h2>
            <DefectTable data={props.data.defects} />
        </Fabric>
    )
}



render(<ProjectPage data={window.__DATA__}></ProjectPage>, document.getElementById("pakkretqc-root"))