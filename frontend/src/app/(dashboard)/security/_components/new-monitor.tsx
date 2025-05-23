import { Button } from '@/components/ui/button'
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import { SearchCheckIcon } from 'lucide-react'
import React from 'react'

function NewMonitor() {
  return (
    <div>
        <Dialog>
            <DialogTrigger asChild>
                <Button><SearchCheckIcon/> New Monitor</Button>
            </DialogTrigger>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className="text-xl font-bold">Add New Security Monitor</DialogTitle>
                    <DialogDescription>
                        Configure a new security monitoring rule to detect suspicious activity or policy violations in your Kubernetes cluster. 
                        Specify the criteria and actions for this monitor below.
                    </DialogDescription>
                </DialogHeader>
                
            </DialogContent>
        </Dialog>
    </div>
  )
}

export default NewMonitor
