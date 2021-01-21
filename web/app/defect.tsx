import * as React from 'react'
import { render } from 'react-dom'

interface DefectPageProps {
    data: {
        defect: Deflect
    }
}

const DefectPage = (props: DefectPageProps) => {

    console.log(props)
    return (
        <div>
            {props.data.defect.name}

            <div dangerouslySetInnerHTML={{
                __html: props.data.defect.description,
            }}></div>
            <div dangerouslySetInnerHTML={{
                __html: props.data.defect["dev-comments"],
            }}></div>
        </div>
    )
}


render(<DefectPage data={window.__DATA__}></DefectPage>, document.getElementById("pakkretqc-root"))