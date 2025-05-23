import React from 'react'
import Logo from './Logo'
import { LoaderCircle } from 'lucide-react'

function Loading() {
  return (
    <div className='min-h-screen w-full flex items-center justify-center flex-col gap-1'>
        <Logo/>
        <p className='text-2xl font-bold mb-3'>Loading please wait.</p>
        <LoaderCircle className='animate-spin h-8 w-8'/>
    </div>
  )
}

export default Loading


