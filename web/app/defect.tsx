import { Fabric, Text } from '@fluentui/react'
import * as React from 'react'
import { render } from 'react-dom'

interface DefectPageProps {
    data: {
        defect: Deflect
        attachment: Attachment[]
    }
}

const DefectPage = (props: DefectPageProps) => {

    return (
        <Fabric>
            {props.data.defect.name}
            <div>
                <div dangerouslySetInnerHTML={{
                    __html: props.data.defect.description,
                }}></div>
                <div dangerouslySetInnerHTML={{
                    __html: props.data.defect["dev-comments"],
                }}></div>
            </div>

            <div>
                <h2>Attachments</h2>
                {
                    props.data.attachment.map(attach => {
                        return (
                            <div key={attach.id}>
                                <Text>
                                    {attach.id} <a href=""> {attach.name} </a>
                                </Text>
                            </div>
                        )
                    })
                }
            </div>
        </Fabric>
    )
}


render(<DefectPage data={window.__DATA__}></DefectPage>, document.getElementById("pakkretqc-root"))