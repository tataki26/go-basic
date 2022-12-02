import './App.css';
import { useState } from 'react';
import Axios from 'axios';

function App() {
  const [connId, setConnId] = useState('')

  const getValue = e => {
    const { name, value } = e.target;
    console.log(name, value);
    setConnId(value);
  }

  const submitId = () => {
    Axios.post('http://localhost:8080/api/insert', {
      id: connId
    }).then(()=>{
      alert('등록 완료');
    })
  };

  return (
    <div className="App">
      <h1>Connection</h1>
      <div className='form-wrapper'>
        <input className='id-input' type="text" onChange={getValue} name="ethernetid" placeholder="ethernet id" />
      </div>
      <button className="submit-button" onClick={submitId}>CONNECT</button>
    </div>
  );
}

export default App;
