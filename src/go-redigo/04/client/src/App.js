import './App.css';

function App() {
  const getValue = e => {
    const { name, value } = e.target;
    console.log(name, value);
  }
  return (
    <div className="App">
      <h1>Connection</h1>
      <div className='form-wrapper'>
        <input className='id-input' type="text" onChange={getValue} name="ethernetid" placeholder="ethernet id" />
      </div>
      <button className="submit-button">CONNECT</button>
    </div>
  );
}

export default App;
