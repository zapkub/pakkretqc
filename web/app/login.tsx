
import { PrimaryButton, TextField } from '@fluentui/react';
import * as React from 'react';
import { render } from 'react-dom';


export const LoginPage = () => {
    return (
        <div>
            <form className="login-form">
                <TextField id="username"></TextField>
                <TextField id="password"></TextField>
                <PrimaryButton>Login</PrimaryButton>
            </form>
        </div>
    )
}
render(<LoginPage></LoginPage>, document.getElementById("pakkretqc-root"))