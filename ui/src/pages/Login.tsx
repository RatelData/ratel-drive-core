import React, {useState} from 'react';
import { Form, Button } from 'react-bootstrap';
import './Login.css';

import { loginUrl } from '../util/config';
import * as StorageUtils from '../util/storage'

function Login() {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [storage_device_checked, setStorageDeviceChecked] = useState<boolean>(StorageUtils.isStorageDevice());

    return(
        <div className='Login-Layout'>
            <div className='Login-Card'>
                <div className='Login-Card-Title'>
                    <p>Sign in to Ratel Drive</p>
                </div>
                
                <div className='Login-Card-Body'>
                    <Form>
                        <div className='Login-Card-Body-Cred'>
                            <Form.Group controlId="formEmail">
                                <Form.Control type="email" size='lg' placeholder="example@email.com" value={email} onChange={e=> setEmail(e.target.value)}/>
                            </Form.Group>
                            <Form.Group controlId="formPassword">
                                <Form.Control type="password" size='lg' placeholder="Password" value={password} onChange={e=> setPassword(e.target.value)}/>
                            </Form.Group>
                        </div>
                        

                        <Form.Group controlId="formUtils" className='Login-Card-Body-Utils'>
                            <Form.Check type="checkbox" checked={storage_device_checked} label="Storage Device" onChange={e=> setStorageDeviceChecked(e.target.checked) } />
                            
                            <Form.Text className="text-muted">
                            <a href='/'>Forget your password?</a>
                            </Form.Text>
                        </Form.Group>

                        <Form.Group controlId="formUtils" className='Login-Card-Body-SignIn'>
                            <Button 
                                variant="primary" 
                                type="submit" 
                                size='lg' 
                                className='Login-Card-Body-BtnSignIn'
                                onClick={e => onLogin(e, {email, password, storage_device_checked})}
                            >
                                SIGN IN
                            </Button>
                        </Form.Group>
                    </Form>
                </div>
            </div>
        </div>
    );
}

export default Login;

interface LoginProps {
    email: string,
    password: string,
    storage_device_checked: boolean
}
async function onLogin(e: React.MouseEvent<HTMLElement, MouseEvent>, props: LoginProps) {
    e.preventDefault();

    const req_options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            user: {
                email: props.email,
                password: props.password,
                is_storage_device: props.storage_device_checked
            }
        })
    };

    fetch(loginUrl(), req_options)
        .then(rsp => rsp.json())
        .then(data => {
            StorageUtils.saveUserCredentials(data['user']['user_id'], data['user']['token'], props.storage_device_checked)
        })
        .catch(e => {
            console.log(e)
        });
}