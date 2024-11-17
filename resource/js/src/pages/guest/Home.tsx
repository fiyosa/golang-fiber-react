import { useState } from 'react'
import { ImgReact } from '../../assets/img'
import './home.css'
import secret from '../../config/secret'
import { useNavigate } from 'react-router-dom'

function Home() {
  const [count, setCount] = useState(0)
  const navigate = useNavigate()
  console.log(secret)

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={'/vite.svg'} className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={ImgReact.toString()} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button
          onClick={() => {
            setCount((count) => count + 1)
          }}
        >
          count is {count}
        </button>
        <button
          style={{ marginLeft: '5px' }}
          onClick={() => {
            setCount(0)
          }}
        >
          Reset
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <div className="card">
        <button onClick={() => navigate('/login')}>Login</button>
        <button style={{ margin: '0 5px' }} onClick={() => navigate('/user')}>
          User
        </button>
        <button onClick={() => navigate('/unknown')}>404</button>
      </div>
      <p className="read-the-docs">Click on the Vite and React logos to learn more</p>
    </div>
  )
}

export default Home
