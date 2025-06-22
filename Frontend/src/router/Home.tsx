import { useEffect, useState } from 'react'
import { DarkModeToggle } from '../components/DarkModeToggle'
import './Home.scss'
import { useNavigate } from 'react-router-dom'
import * as React from 'react'
import Watermark from '../components/Watermark.tsx'

export function Home() {
  const [visible, setVisible] = useState(false)
  const [uuid, setUUID] = useState<string>(crypto.randomUUID())
  const navigate = useNavigate()

  useEffect(() => {
    const timer = setTimeout(() => setVisible(true), 100)
    return () => clearTimeout(timer)
  }, [])

  const handleNavigation = () => {
    setVisible(false)
    setTimeout(() => {
      navigate(`/note/${uuid}`)
    }, 500)
  }

  const handleKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleNavigation()
    }
  }

  return (
    <>
      <Watermark text={uuid} fontSize={20} gapX={150} gapY={150} />
      <DarkModeToggle />
      <div className="content">
        <div className={`background ${visible ? 'visible' : ''}`}>
          <div className="title">
            <div className="logo" />
            <div
              className="github"
              onClick={() => window.open('https://github.com/Sn0wo2/QuickNote', '_blank')}
            />
          </div>

          <p className="subtitle">
            <span className="highlight">QuickNote</span>
            <span className="note">Create and share notes quickly and easily.</span>
            <span className="warning">
              Do not upload any content that violates laws and regulations.
            </span>
          </p>

          <div className="input-container">
            <input
              className="uuid-input"
              type="text"
              value={uuid}
              onChange={(e) => setUUID(e.target.value)}
              onKeyDown={handleKeyPress}
            />
            <button className="submit-btn" onClick={handleNavigation}>
              &rarr;
            </button>
          </div>
        </div>
      </div>
    </>
  )
}
