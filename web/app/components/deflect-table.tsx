import { DetailsList, IColumn, Link, mergeStyleSets, SelectionMode, Text } from '@fluentui/react';
import * as React from 'react';
import * as timeago from 'timeago.js';

const classNames = mergeStyleSets({
    severifyCell: {
      textAlign: 'center'
    },
    nameCell: {
    }
})
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
        data: 'string',
        className: classNames.severifyCell,
        isRowHeader: true,
        onRender: (item) => {
            return severityMap(item['severity'])
                    // rec.severity = severityMap(rec.severity)
        }
    },
    {
        key: 'col3',
        name: "Name",
        fieldName: "name",
        minWidth: 210,
        isResizable: true,
        isRowHeader: true,
        onRender: (item) =>{
            return <Link target="popup" href={item['url']}>{item['name']}</Link>
        }
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
        minWidth: 100,
        maxWidth: 100,
    },
]

export const DefectTable = (props: { data: Defect[] }) => {

    return (
        <div>
            <DetailsList
                selectionMode={SelectionMode.none}
                columns={columns}
                items={props.data.map(rec => {
                    rec["creation-time"] = new Date(rec["creation-time"]).toLocaleDateString()
                    rec["last-modified"] = timeago.format(rec["last-modified"])
                    return rec
                })}
            ></DetailsList>
            {
                props.data.length === 0 ? <div style={{textAlign: 'center'}}><Text>No data</Text></div> : null
            }
        </div>
    )
}

const severityMap = (v: string) => {
    switch (v) {
        case "1":
            return "🤬"
        case "2":
            return "😱"
        case "3":
            return "😕"
        case "4":
            return "😒"
    }
}