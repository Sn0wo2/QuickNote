import React, { useRef, useEffect, useCallback } from 'react'
import './Watermark.scss'

interface WatermarkProps {
  text: string
  fontSize: number
  gapX: number
  gapY: number
}

const Watermark: React.FC<WatermarkProps> = ({ text, fontSize, gapX, gapY }) => {
  const canvasRef = useRef<HTMLCanvasElement>(null)

  const drawWatermark = useCallback(() => {
    const canvas = canvasRef.current
    if (!canvas) return

    const ctx = canvas.getContext('2d')
    if (!ctx) return

    const dpr = window.devicePixelRatio || 1
    const rect = canvas.getBoundingClientRect()
    canvas.width = rect.width * dpr
    canvas.height = rect.height * dpr
    ctx.scale(dpr, dpr)

    ctx.clearRect(0, 0, canvas.width, canvas.height)

    ctx.font = `${fontSize}px 'Helvetica Neue', Arial, sans-serif`
    ctx.fillStyle = 'rgba(0, 0, 0, 0.01)'
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'

    const frequencyX = 0.005
    const frequencyY = 0.004
    const amplitude = 0.3

    for (let y = 0; y < rect.height + gapY; y += gapY) {
      for (let x = 0; x < rect.width + gapX; x += gapX) {
        ctx.save()
        const rotation = (Math.sin(x * frequencyX) + Math.cos(y * frequencyY)) * amplitude
        ctx.translate(x, y)
        ctx.rotate(rotation)
        ctx.fillText(text, 0, 0)
        ctx.restore()
      }
    }
  }, [text, fontSize, gapX, gapY])

  useEffect(() => {
    drawWatermark()
    window.addEventListener('resize', drawWatermark)
    return () => {
      window.removeEventListener('resize', drawWatermark)
    }
  }, [drawWatermark])

  return <canvas ref={canvasRef} className="watermark-canvas" />
}

export default Watermark
