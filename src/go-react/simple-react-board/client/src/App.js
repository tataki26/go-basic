import { useEffect, useState } from 'react';
import './App.css';
import { CKEditor } from '@ckeditor/ckeditor5-react';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';
import ReactHtmlParser from 'html-react-parser';
import Axios from 'axios';

function App() {
  const [gameContent, setGameContent] = useState({
    title: '',
    content: ''
  });

  const [viewContent, setViewContent] = useState([]);

  useEffect(()=>{
    Axios.get('http://localhost:8080/api/get').then((response)=>{
      setViewContent(response.data)
    })
  }, [viewContent])

  const submitReview = () => {
    Axios.post('http://localhost:8080/api/insert', {
      title: gameContent.title,
      content: gameContent.content
    }).then(() => {
      alert('등록 완료!');
    })
  };

  const getValue = e => {
    const { name, value } = e.target;
      setGameContent({
        ...gameContent,
        [name]: value
      })
      console.log(gameContent)
  };

  return (
    <div className="App">
      <h1>Game Review</h1>
      <div className='game-container'>
        {viewContent.map(element =>
          <div style={{ border: '1px solid #333' }}>
            <h2>{element.Title}</h2>
            <div>
             {ReactHtmlParser(element.Content)}
            </div>
            </div>
            )}
      </div>
      <div className='form-wrapper'>
        <input className="title-input" type='text' placeholder='제목' onChange={getValue} name='title'/>
        <CKEditor
          editor={ClassicEditor}
          data="<p>Hello from CKEditor 5!</p>"
          onReady={editor => {
            // You can store the "editor" and use when it is needed.
            console.log('Editor is ready to use!', editor);
          }}
          onChange={(event, editor) => {
            const data = editor.getData();
            console.log({ event, editor, data });
            setGameContent({
              ...gameContent,
              content: data
            })
            console.log(gameContent);
          }}
          onBlur={(event, editor) => {
            console.log('Blur.', editor);
          }}
          onFocus={(event, editor) => {
            console.log('Focus.', editor);
          }}
        />
      </div>
      <button className="submit-button"
      onClick={submitReview}
      >입력</button>
    </div>
  );
}

export default App;