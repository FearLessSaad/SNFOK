import React from 'react'

const page = ({params}:{params:{id:string}}) => {
  const {id} = params
    return (
    <div>
      <div>{id}</div>
    </div>
  )
}

export default page
