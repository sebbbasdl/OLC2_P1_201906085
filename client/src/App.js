import './App.css';

//theme
import "primereact/resources/themes/lara-light-indigo/theme.css";     
//core
import "primereact/resources/primereact.min.css";                                               

import Home from './pages/home';
import Reports from './pages/reports';

function App() {
  return (
    <div className="App">
      <div className="App-header2">
        <div className="text-center">
          <h1 className="display-4 title-animation">
          T-Swift
          </h1>
        </div>
        <br></br>
        <header className="App-header">
        <Home/>
      </header>
      <br></br>
      <br></br>
      <div style={{ position: 'relative' }}>
        <div style={{
          position: 'absolute',
          top: '50%',
          left: '0',
          width: '100%',
          height: '8px', // Grosor de la línea
          backgroundColor: '#007bff' // Color celeste
        }}></div>
      </div>
      <div>
        <h2 style={{
          textAlign: 'center',
          backgroundColor: '#005b9f', // Color celeste más oscuro
          color: 'white', // Color del texto
          padding: '10px' // Espacio interior alrededor del texto
        }}>Reportes:</h2>
      </div>
      
      <br></br>
      <br></br>
      <div className='App-Header reports-background'>
        <Reports/>
      </div>
      </div>  
      
      

      
      
    </div>


  );
}

export default App;
