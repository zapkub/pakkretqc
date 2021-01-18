import { DetailsList, IColumn, SelectionMode } from '@fluentui/react'
import * as React from 'react'


const columns: IColumn[] = [
    {
        key: 'col1',
        name: "ID",
        fieldName: "id",
        minWidth: 50,
        maxWidth: 50,
    },
    {
        key: 'col2',
        name: "Severity",
        fieldName: "severity",
        minWidth: 55,
        maxWidth: 55, 
    },
    {
        key: 'col3',
        name: "Name",
        fieldName: "name",
        minWidth: 210,
        isResizable: true,
    },
    {
        key: 'col22',
        name: 'Status',
        fieldName: 'user-46',
        minWidth: 120,
        isResizable: true,
    },
    {
        key: 'col44',
        name: 'Owner',
        fieldName: 'owner',
        minWidth: 100,
        maxWidth: 100,
    },
    {
        key: 'col5',
        name: 'Last Updated Date',
        fieldName: 'last-modified',
        minWidth: 120,
        maxWidth: 120,
    },
    {
        key: 'col4',
        name: 'Created Date',
        fieldName: 'creation-time',
        minWidth: 120,
        maxWidth: 120,
    },
]

export const DeflectTable = (props: { data: Deflect[] }) => {

    return (
        <div>
            <DetailsList
                selectionMode={SelectionMode.none}
                columns={columns}
                items={props.data.map(rec => {
                    rec["creation-time"] = new Date(rec["creation-time"]).toLocaleString()
                    rec["last-modified"] = new Date(rec["last-modified"]).toLocaleString()
                    return rec
                })}
            ></DetailsList>
        </div>
    )
}