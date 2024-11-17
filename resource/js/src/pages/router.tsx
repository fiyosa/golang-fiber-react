import { createBrowserRouter } from 'react-router-dom'
import Layout from '../components/molecules/layout/layout'
import Home from './guest/Home'
import Login from './guest/Login'
import NotFound from './guest/404'
import User from './user/User'

const router = createBrowserRouter(
  [
    { path: '/', element: <Home /> },
    { path: '/login', element: <Login /> },
    {
      path: '/',
      element: <Layout />,
      children: [
        { path: 'user', element: <User /> },
        //
      ],
    },
    { path: '*', element: <NotFound /> },
  ],
  {
    future: {
      v7_fetcherPersist: true,
      v7_normalizeFormMethod: true,
      v7_partialHydration: true,
      v7_relativeSplatPath: true,
      v7_skipActionErrorRevalidation: true,
    },
  }
)

export default router
