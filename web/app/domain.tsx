import { Button, Dropdown } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';


interface DomainPageData {
    projects: Array<{ name: string }>
    domain: string
}


const DomainPage = (props: { data: DomainPageData }) => {
    const [project, setProject] = React.useState<string | undefined>(undefined)
    return (
        <div>
            Domain: {props.data.domain}
            {
                props.data.projects.length > 0 ?
                    <>
                        <Dropdown
                            onChange={(evt, val) => setProject(val.data)}
                            defaultValue={props.data.projects[0].name}
                            options={props.data.projects.map(proj => {
                                return {
                                    key: proj.name,
                                    data: proj.name,
                                    text: proj.name,
                                }
                            })}
                        ></Dropdown>
                        <Button style={{ marginTop: 4 }} disabled={!project} as="a" href={`./${props.data.domain}/projects/${project}`}>Next</Button>
                    </> : "Domain does not have any project"
            }
        </div>
    )
}



render(<DomainPage data={window.__DATA__}></DomainPage>, document.getElementById("pakkretqc-root"))