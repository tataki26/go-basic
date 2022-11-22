import { useState } from 'react';
import './App.css';
import { CKEditor } from '@ckeditor/ckeditor5-react';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';
import ReactHtmlParser from 'html-react-parser';

function App() {
  const [gameContent, setGameContent] = useState({
    title: '',
    content: ''
  });

  const [viewContent, setViewContent] = useState([]);

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
      <h1>Movie Review</h1>
      <div className='movie-container'>
        {viewContent.map(element =>
          <div>
            <h2>{element.title}</h2>
            <div>
              {ReactHtmlParser(element.content)}
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
      onClick={() => {
        setViewContent(viewContent.concat({...gameContent}));
      }
      }>입력</button>
    </div>
  );
}

export default App;