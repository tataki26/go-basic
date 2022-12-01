import './App.css';

function App() {
  return (
    <div className="App">
      <h1>Connection</h1>
      <div className='form-wrapper'>
        <input type="text" name="ethernetid" placeholder="ethernet id" />
      </div>
      <button className="submit-button">CONNECT</button>
    </div>
  );
}

export default App;
