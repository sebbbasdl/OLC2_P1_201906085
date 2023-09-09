import React, { useState } from 'react';
import { InputTextarea } from 'primereact/inputtextarea';
import { Button } from 'primereact/button';
import { PostMethod } from '../api/http';

const interpreterAPI = process.env.REACT_APP_API_URL_INTERPRETER;

const Home = () => {
    const [codeText, setCodeText] = useState('');
    const [consoleText, setConsoleText] = useState('');

    const CompileInterpreter = async () => {
        // Limpia la consola antes de hacer la solicitud al backend
        setConsoleText('');

        const resp = await PostMethod(interpreterAPI + 'Interpreter', { Content: codeText });
        
        // Actualiza la consola con la respuesta del backend
        setConsoleText(resp?.Output);
    };

    const handleFileUpload = (e) => {
        const file = e.target.files[0];

        if (file) {
            const reader = new FileReader();

            reader.onload = (e) => {
                const fileContent = e.target.result;
                setCodeText(fileContent);
            };

            reader.readAsText(file);
        }
    };

    return (
        <div>
            <div style={{ textAlign: 'left', marginRight: '2%', marginTop: '10px' }}>
                <label htmlFor="fileInput" style={{ cursor: 'pointer', backgroundImage: 'linear-gradient(to right, #4CAF50, #45a049)', padding: '10px', borderRadius: '5px', color: 'white' }}>
                    Cargar archivo
                    <input
                        type="file"
                        accept=".txt"
                        onChange={handleFileUpload}
                        id="fileInput"
                        style={{ display: 'none' }}
                    />
                </label>
            </div>
            <div>
                <label
                    htmlFor="codeTextarea"
                    style={{
                        background: '#7464ff',
                        padding: '10px',
                        borderRadius: '15px',
                        marginBottom: '10px',
                    }}
                >
                    CÓDIGO:
                </label>
                <InputTextarea
                    id="codeTextarea"
                    value={codeText}
                    rows={30}
                    cols={140}
                    style={{ marginBottom: '5%', marginRight: '2%', backgroundColor: '#f6dfff' }}
                    onChange={(e) => {
                        setCodeText(e.target.value);
                    }}
                    onKeyDown={(e) => {
                        if (e.key === 'Tab') {
                            e.preventDefault(); // Evita el comportamiento predeterminado (cambio de campo)
                            const { selectionStart, selectionEnd } = e.target;
                            const newValue =
                                codeText.substring(0, selectionStart) + '    ' + codeText.substring(selectionEnd);
                            setCodeText(newValue);
                        }
                    }}
                />
            </div>
            <Button label="EJECUTAR" onClick={CompileInterpreter} style={{ width: '600px', height: '50px' }} />
            <br></br>
            <br></br>
            <div>
                <label
                    htmlFor="consoleTextarea"
                    style={{
                        background: 'green',
                        padding: '10px',
                        borderRadius: '15px',
                        marginBottom: '10px',
                    }}
                >
                    CONSOLA:
                </label>
                <InputTextarea
                    id="consoleTextarea"
                    value={consoleText}
                    rows={30}
                    cols={140}
                    style={{ marginBottom: '5%', marginRight: '2%', backgroundColor: '#e2f1db' }}
                    onChange={(e) => {
                        setConsoleText(e.target.value);
                    }}
                />
            </div>
            
            
        </div>
    );
};

export default Home;
