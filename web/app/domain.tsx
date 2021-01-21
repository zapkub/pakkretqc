import { Button, Dropdown, Fabric, PrimaryButton } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';


interface DomainPageData {
    projects: Array<{ name: string }>
    domain: string
}


const DomainPage = (props: { data: DomainPageData }) => {
    const [project, setProject] = React.useState<string | undefined>(undefined)
    return (
        <Fabric className="content">
            <div className="project-select-form-container">
                <h2>Domain: {props.data.domain}</h2>
                {
                    props.data.projects.length > 0 ?
                        <>
                            <Dropdown
                                onChange={(evt, val) => setProject(val.data)}
                                defaultValue={props.data.projects[0].name}
                                placeHolder={'Pick one here'}
                                label="🔬Which project you wanna work with today?"
                                options={props.data.projects.map(proj => {
                                    return {
                                        key: proj.name,
                                        data: proj.name,
                                        text: proj.name,
                                    }
                                })}
                            ></Dropdown>
                            <Button  style={{ marginTop: 8, marginRight: 4 }} as="a" href={`/login`}>Back</Button>
                            <PrimaryButton  style={{ marginTop: 4 }} disabled={!project} as="a" href={`./${props.data.domain}/projects/${project}`}>Next</PrimaryButton>
                        </> : "Domain does not have any project"
                }
            </div>
        </Fabric>
    )
}



render(<DomainPage data={window.__DATA__}></DomainPage>, document.getElementById("pakkretqc-root"))