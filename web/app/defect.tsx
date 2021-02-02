import { Fabric } from '@fluentui/react'
import * as React from 'react'
import { render } from 'react-dom'
import { Attachments, DefectDetail } from './components'

interface DefectPageProps {
    data: {
        defect: Defect
        attachment: Attachment[]
        project: string
        domain: string
    }
}

const DefectPage = (props: DefectPageProps) => {

    return (
        <Fabric>
            <div className="defect-container">
                <div className="defect-detail-container" >
                    <DefectDetail {...props.data.defect} />
                </div>
                <div className="defect-attachments-container">
                    <Attachments domain={props.data.domain} project={props.data.project} attachments={props.data.attachment} />
                </div>
            </div>
        </Fabric>
    )
}


render(<DefectPage data={window.__DATA__}></DefectPage>, document.getElementById("pakkretqc-root"))