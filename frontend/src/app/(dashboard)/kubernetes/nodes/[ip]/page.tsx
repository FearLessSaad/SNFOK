import React from 'react'

function NodeInfoPage({params}:{params: {ip: string}}) {
    const {ip} = params
  return (
    <div>
      <h1 className='text-3xl'>{ip}</h1>
    </div>
  )
}

export default NodeInfoPage
