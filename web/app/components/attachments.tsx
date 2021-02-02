import { DetailsList, IColumn, Link, SelectionMode } from '@fluentui/react'
import * as React from 'react'

const columns: (domain: string, project: string) => IColumn[] = (domain, project) => [
    {
        key: 'col1',
        name: "ID",
        fieldName: "id",
        minWidth: 50,
        maxWidth: 50,
    },
    {
        key: 'col2',
        name: "Name",
        fieldName: "name",
        minWidth: 255,
        maxWidth: 455,
        data: 'string',
        isRowHeader: true,
        onRender: (item, idx, col) => {
            return <Link target="popup" download={item['name']} href={`/domains/${domain}/projects/${project}/attachments/${item.id}`}>{item['name']}</Link>
        }
    },
    {
        key: 'col22',
        name: 'Size',
        fieldName: 'file-size',
        minWidth: 120,
        isResizable: true,
    },

]

export const Attachments = (props: { domain: string, project: string, attachments: Attachment[] }) => {
    return (
        <>
            <h3>📑 Attachments</h3>
            <div >
                <DetailsList
                    selectionMode={SelectionMode.none}
                    columns={columns(props.domain, props.project)}
                    items={props.attachments}
                />
            </div>
        </>
    )
}