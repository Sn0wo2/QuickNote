import * as React from 'react'
import { useCallback, useEffect, useRef, useState } from 'react'
import MDEditor from '@uiw/react-md-editor'
import './Note.scss'
import { DarkModeToggle } from '../components/DarkModeToggle.tsx'
import { toast, Toaster } from 'react-hot-toast'
import { getNote, saveNote } from '../services/noteAPI'
import { useParams } from 'react-router-dom'
import Watermark from '../components/Watermark.tsx'

function useCtrlS(callback: () => void) {
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 's') {
        e.preventDefault()
        callback()
      }
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [callback])
}

export function Note() {
  const { id } = useParams<{ id: string }>()
  const [title, setTitle] = useState<string>('')
  const [content, setContent] = useState<string>('')
  const [mode, setMode] = useState<'edit' | 'preview' | 'both'>('both')
  const [visible, setVisible] = useState<boolean>(false)
  const prevTitleRef = useRef<string>('')
  const prevContentRef = useRef<string>('')
  const lastSaveTimeRef = useRef<number>(0)
  const saveTimerRef = useRef<NodeJS.Timeout | null>(null)

  const loadNote = useCallback(async () => {
    if (!id) {
      console.error('Note ID is not provided')
      return
    }
    try {
      const noteData = await getNote(id)
      if (noteData) {
        setTitle(noteData.title)
        setContent(noteData.content)
        prevTitleRef.current = noteData.title
        prevContentRef.current = noteData.content
      }
    } catch (e) {
      toast.error('Failed to load note.')
      console.error(e)
    }
    setTimeout(() => {
      setVisible(true)
    }, 100)
  }, [id])

  useEffect(() => {
    void loadNote()

    const handleBeforeUnload = (e: BeforeUnloadEvent) => {
      if (title !== prevTitleRef.current || content !== prevContentRef.current) {
        e.preventDefault()
        return ''
      }
    }

    window.addEventListener('beforeunload', handleBeforeUnload)
    return () => {
      window.removeEventListener('beforeunload', handleBeforeUnload)
    }
  }, [])

  const now = Date.now()

  const save = useCallback(async () => {
    if (!id) {
      console.error('Note ID is undefined')
      return
    }
    try {
      const success = await saveNote(id, title, content)
      if (success) {
        toast.success('Note saved successfully')
        prevTitleRef.current = title
        prevContentRef.current = content
        lastSaveTimeRef.current = now
      }
    } catch (e) {
      toast.error('Failed to save note.')
      console.error(e)
    }
  }, [id, title, content, now])

  const throttledSave = useCallback(async () => {
    const hasChanges = title !== prevTitleRef.current || content !== prevContentRef.current

    if (!hasChanges) return

    if (now - lastSaveTimeRef.current >= 5000) {
      await save()
    } else {
      if (saveTimerRef.current) clearTimeout(saveTimerRef.current)
      saveTimerRef.current = setTimeout(
        () => throttledSave(),
        5000 - (now - lastSaveTimeRef.current),
      )
    }
  }, [title, content, now, save])

  useEffect(() => {
    void throttledSave()
    return () => {
      if (saveTimerRef.current) {
        clearTimeout(saveTimerRef.current)
      }
    }
  }, [throttledSave])

  const handleModeChange = (value: 'edit' | 'preview' | 'both') => {
    setMode(value)
    setVisible(false)
    setTimeout(() => {
      setVisible(true)
    }, 100)
  }

  const handleContentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setContent(e.target.value)
  }

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value)
  }

  useCtrlS(() => {
    void save()
  })

  return (
    <>
      <Watermark text={id || ''} fontSize={10} gapX={150} gapY={150} />
      <div className="content">
        <DarkModeToggle />
        <div className={`note-container ${visible ? 'visible' : ''}`}>
          <div className="note-mode-toggle">
            <button onClick={() => handleModeChange('edit')}>Edit Only</button>
            <button onClick={() => handleModeChange('preview')}>Preview Only</button>
            <button onClick={() => handleModeChange('both')}>Both</button>
            <div className="note-logo" />
          </div>
          <div className="note-header">
            <input
              type="text"
              className="note-title"
              value={title}
              onChange={handleTitleChange}
              placeholder="Note title"
            />
          </div>
          <div className="note-content">
            {(mode === 'edit' || mode === 'both') && (
              <textarea
                className="note-editor"
                value={content}
                onChange={handleContentChange}
                placeholder="Write your note here (Markdown)..."
              />
            )}
            {(mode === 'preview' || mode === 'both') && (
              <div className="note-preview" data-color-mode="light">
                <h1>{title}</h1>
                <MDEditor.Markdown source={content} />
              </div>
            )}
          </div>
        </div>
        <Toaster position="top-right" />
        <button
          onClick={() => {
            navigator.clipboard
              .writeText(window.location.href)
              .then(() => {
                toast.success('Copied to clipboard!')
              })
              .catch((e) => {
                console.error(e)
                toast.error(e instanceof Error ? e.message : String(e))
              })
          }}
          className="url"
        >
          {window.location.host + window.location.pathname}
        </button>
      </div>
    </>
  )
}
