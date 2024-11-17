import { useNavigate } from 'react-router-dom'

export default function NotFound() {
  const navigate = useNavigate()

  return (
    <>
      <h1>404 Not Found</h1>
      <button onClick={() => navigate('/')}>Home</button>
    </>
  )
}
