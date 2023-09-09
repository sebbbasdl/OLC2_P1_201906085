import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import cst from './imagenes/arbol.svg';
import errores from './imagenes/errores.svg'
import tabla from './imagenes/tabla.svg'

const Reports = () => {
  const [activeTab, setActiveTab] = useState('active'); // Inicialmente, la pestaña "Active" está activa

  const handleTabClick = (tab) => {
    setActiveTab(tab);
  };

  const renderImage = () => {
    const imageSize = {
      width: '1800px', // Ancho deseado (ajusta este valor)
      height: '300px', // Alto deseado (ajusta este valor)
    };
  
    switch (activeTab) {
      case 'active':
        return <img src={cst} alt="Imagen de Active" style={imageSize} />;
      case 'link1':
        return <img src={tabla} alt="Imagen de Link 1" style={imageSize} />;
      case 'link2':
        return <img src={errores} alt="Imagen de Link 2" style={imageSize} />;
      default:
        return null;
    }
  };
  
  

  return (
    <div className="text-center"> {/* Para centrar el contenido */}
      <ul className="nav nav-tabs" style={{ fontSize: '20px' }}> {/* Para hacer las pestañas un poco más grandes */}
        <li className="nav-item">
          <a
            className={`nav-link ${activeTab === 'active' ? 'active' : ''}`}
            href="#"
            onClick={() => handleTabClick('active')}
            style={{
              backgroundColor: activeTab === 'active' ? '#8BC34A' : '#03A9F4', // Color de fondo
              color: activeTab === 'active' ? '#fff' : '', // Color del texto
              padding: '10px 20px', // Añadir espaciado
            }}
          >
            CST
          </a>
        </li>
        <li className="nav-item">
          <a
            className={`nav-link ${activeTab === 'link1' ? 'active' : ''}`}
            href="#"
            onClick={() => handleTabClick('link1')}
            style={{
              backgroundColor: activeTab === 'link1' ? '#8BC34A' : '#03A9F4', // Color de fondo
              color: activeTab === 'link1' ? '#fff' : '', // Color del texto
              padding: '10px 20px', // Añadir espaciado
            }}
          >
            Tabla de Simbolos
          </a>
        </li>
        <li className="nav-item">
          <a
            className={`nav-link ${activeTab === 'link2' ? 'active' : ''}`}
            href="#"
            onClick={() => handleTabClick('link2')}
            style={{
              backgroundColor: activeTab === 'link2' ? '#8BC34A' : '#03A9F4', // Color de fondo
              color: activeTab === 'link2' ? '#fff' : '', // Color del texto
              padding: '10px 20px', // Añadir espaciado
            }}
          >
            Errores
          </a>
        </li>
      </ul>

      <div className="tab-content">
        {renderImage()}
      </div>
    </div>
  );
};

export default Reports;
