
import { Dropdown, PrimaryButton, TextField } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';


interface LoginPageData {
    username: string
    domains: Array<{ name: string }>
}



export const LoginPage = (props: { data: LoginPageData }) => {
    const [domain, setDomain] = React.useState<string>("")
    return (
        <div>
            <form className="login-form" method="POST" action="/login">
                <TextField disabled={!!props.data.username} defaultValue={props.data.username} name="username" placeholder="username" className="text-field" id="username"></TextField>
                <TextField disabled={!!props.data.username} name="password" placeholder="password" type="password" className="text-field" id="password"></TextField>
                <PrimaryButton disabled={!!props.data.username} as="input" type="submit">Authenticate</PrimaryButton>
                {
                    props.data.username ? (
                        <>
                            <Dropdown
                                label="Domain"
                                selectedKey={domain}
                                onChange={(event, item) => setDomain(item.data)}
                                options={props.data.domains.map(domain => {
                                    return {
                                        key: domain.name,
                                        text: domain.name,
                                        data: domain.name,
                                    }
                                })}
                            ></Dropdown>
                            <input type="hidden" name="currentDomain" value={domain} />
                            <div className="domain-confirm-actions">
                                <PrimaryButton as="input" name="action" value="proceed" type="submit">Proceed</PrimaryButton>
                                <PrimaryButton as="input" name="action" value="cancel" type="submit">Cancel</PrimaryButton>
                            </div>
                        </>
                    ) : null
                }
            </form>
        </div>
    )
}
render(<LoginPage data={window.__DATA__}></LoginPage>, document.getElementById("pakkretqc-root"))