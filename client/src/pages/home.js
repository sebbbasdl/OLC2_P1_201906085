import React from 'react';
import { useState } from 'react';
import { InputTextarea } from 'primereact/inputtextarea';
import { Button } from 'primereact/button';
import { PostMethod } from '../api/http';

const interpreterAPI = process.env.REACT_APP_API_URL_INTERPRETER;

const Home = () => {
    const [codeText, setCodeText] = useState('');
    const [consoleText, setConsoleText] = useState('');

    const CompileInterpreter = async () => {
        const resp = await PostMethod(interpreterAPI + 'Interpreter', { Content: codeText });
        await setConsoleText(resp?.Output);
    };

    return (
        <div>
            <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                <div>
                    <label
                        htmlFor="codeTextarea"
                        style={{
                            background: '#7464ff',
                            padding: '10px',
                            borderRadius: '15px',
                            marginBottom: '10px', // Agregado para crear espacio entre la etiqueta y el InputTextarea
                        }}
                    >
                        CODIGO:
                    </label>
                    <InputTextarea
                        id="codeTextarea"
                        value={codeText}
                        rows={30}
                        cols={70}
                        style={{ marginBottom: '5%', marginRight: '2%',backgroundColor: '#f6dfff' }}
                        onChange={(e) => {
                            setCodeText(e.target.value);
                        }}
                    />
                </div>

                <div>
                    <label
                        htmlFor="codeTextarea"
                        style={{
                            background: 'green',
                            padding: '10px',
                            borderRadius: '15px',
                            marginBottom: '10px', // Agregado para crear espacio entre la etiqueta y el InputTextarea
                        }}
                    >
                        CONSOLA:
                    </label>
                    <InputTextarea
                        id="codeTextarea"
                        value={consoleText}
                        rows={30}
                        cols={70}
                        style={{ marginBottom: '5%', marginRight: '2%' , backgroundColor: '#e2f1db'}}
                        onChange={(e) => {
                            setConsoleText(e.target.value);
                        }}
                    />
                </div>
            </div>
            <div style={{ display: 'flex', justifyContent: 'center' }}>
                <Button label="EJECUTAR" onClick={CompileInterpreter} style={{ width: '600px', height: '50px' }} />
            </div>

        </div>
    );
};

export default Home;