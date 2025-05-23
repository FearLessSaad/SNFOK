import { Badge } from '@/components/ui/badge';
import React from 'react'

function NameSpacePage({params}: {params: {name: string}}) {
    const {name} = params;
  return (
    <div className="p-6 space-y-4">
      <h1 className="text-3xl font-semibold">Namespace Info</h1>
        <p className='text-muted-foreground -mt-3'><Badge variant={"default"}>{name}</Badge> <Badge variant={"outline"}>192.168.1.1</Badge></p>

    </div>
  )
}

export default NameSpacePage
