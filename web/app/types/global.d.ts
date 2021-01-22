
declare interface Window { __DATA__: any }
declare interface Deflect {
    "user-46": string // this is status
    owner: string
    name: string
    severity: string
    description: string
    "dev-comments": string
    "last-modified": string
    "creation-time": string

    url: string

}

declare interface Attachment {

    type: string;
    "last-modified": string;
    "vc-cur-ver"?: any;
    "vc-user-name"?: any;
    name: string;
    "file-size": number;
    "ref-subtype": number;
    description?: any;
    id: number;
    "ref-type": string;
    entity: {
        id: string
        type: string
    };

}