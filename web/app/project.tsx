
import { Button, Fabric, PrimaryButton, Text, TextField } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';
import { useQueryParam } from './common';
import { DefectTable } from './components';


interface ProjectPageData {
    total: number
    domain: string
    project: string
    defects: Defect[]
    username: string
}

const ProjectPage = (props: { data: ProjectPageData }) => {
    const [qq, setqq] = useQueryParam("query", "")
    return (
        <Fabric className="project-page">
            <div className="canopy">
                <div className="ms-Grid">
                    <div style={{ marginBottom: 8 }}>
                        Welcome: <Text style={{ fontWeight: 'bold' }}>{props.data.username || 'Unknown'}</Text>
                    </div>
                    <h2 className="ms-Grid-col"> Domain: {props.data.domain}, Project: {props.data.project} </h2>
                </div>
                <div className="user-profile-box">
                    <Button as="a" href="/login">Logout</Button>
                </div>
            </div>
            <h2>🐛🐞 Defect list</h2>
            <div style={{ width: "100%", display: "flex" }}>
                <form style={{ flexGrow: 1, display: 'flex' }} method="GET" action={`/domains/${props.data.domain}/projects/${props.data.project}`}>
                    <input type="hidden" name="query" value={qq} />
                    <div style={{flexGrow: 1}}>
                        <TextField value={qq} onChange={(_, v) => setqq(v)} placeholder="Filter query for example owner = 'rungsikorn.r' "></TextField>
                    </div>
                    <PrimaryButton as="input" type="submit" style={{ marginLeft: 8 }}>Submit</PrimaryButton>
                </form>
                <Button disabled={qq === ""} as="a" href={`/domains/${props.data.domain}/projects/${props.data.project}`} style={{ marginLeft: 8 }}>Clear</Button>
            </div>
            <DefectTable data={props.data.defects || []} />
        </Fabric>
    )
}



render(<ProjectPage data={window.__DATA__}></ProjectPage>, document.getElementById("pakkretqc-root"))