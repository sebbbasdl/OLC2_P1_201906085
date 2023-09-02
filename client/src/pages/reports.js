import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './imagenes/arbol.svg';

const Reports = () => {
  const [activeTab, setActiveTab] = useState('active'); // Inicialmente, la pestaña "Active" está activa

  const handleTabClick = (tab) => {
    setActiveTab(tab);
  };

  const renderImage = () => {
    switch (activeTab) {
      case 'active':
        return <img src={logo} alt="Imagen de Active" />;
      case 'link1':
        return <img src="imagen_link1.jpg" alt="Imagen de Link 1" />;
      case 'link2':
        return <img src="imagen_link2.jpg" alt="Imagen de Link 2" />;
      default:
        return null; // En caso de pestaña deshabilitada o no reconocida
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
            Tabla de Simbolos
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
            CST
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
