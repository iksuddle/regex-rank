import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { BrowserRouter, Routes } from 'react-router'
import { Route } from 'react-router'
import RegexInput from './components/RegexInput.tsx'
import Dashboard from './components/Dashboard.tsx'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<App />} >
                    <Route index element={<RegexInput />} />
                    <Route path="/login" element={<Dashboard />}/>
                </Route>
            </Routes>
        </BrowserRouter>
    </StrictMode>,
)
