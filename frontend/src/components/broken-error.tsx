import React from 'react'
import {ServerOffIcon} from "lucide-react";
import { Button } from './ui/button';

function BrokenError() {
  return (
    <div className='flex justify-center items-center gap-4 flex-col w-full'>
        <ServerOffIcon className='text-primary-foreground h-24 w-24'/>
        <p className='w-full md:w-[30%] text-center text-muted-foreground font-semibold'>You are facing error in your system services please fix it. Contact with administrator</p>
        <Button>Contact Us</Button>
    </div>
  )
}

export default BrokenError
