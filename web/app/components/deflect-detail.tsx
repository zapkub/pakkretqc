import * as React from 'react'






export const DefectDetail = (props: Defect) => {
    return (
        <>
            <h3>🐞 {props.name} </h3>
            <h4>Detected By: {props["detected-by"]}</h4>
            <div dangerouslySetInnerHTML={{ __html: props.description }}></div>

            <h3>🗣 Comments</h3>
            <div style={{}} dangerouslySetInnerHTML={{ __html: props["dev-comments"] }}></div>
        </>
    )
}