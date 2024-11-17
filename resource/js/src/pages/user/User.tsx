import { useNavigate } from 'react-router-dom'

export default function User() {
  const navigate = useNavigate()

  return (
    <div>
      <h1>User</h1>
      <button onClick={() => navigate('/')}>Home</button>
    </div>
  )
}
