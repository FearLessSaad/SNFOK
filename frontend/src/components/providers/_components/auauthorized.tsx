import React from 'react'
import Link from 'next/link'
import { Button } from '@/components/ui/button'
import { ShieldOffIcon } from 'lucide-react'


function UuAuthorized() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-background text-foreground">
      <ShieldOffIcon className="mb-6 h-16 w-16 text-muted-foreground" />
      <h1 className="text-4xl font-bold tracking-tight sm:text-5xl">Access Forbidden!</h1>
      <p className="mt-4 max-w-md text-center text-muted-foreground">
        You don&apos;t have permission to access this page. Please contact your administrator if you believe this is a mistake.
      </p>
      <Button asChild className="mt-8">
        <Link href="/dashboard">Back to Dashboard</Link>
      </Button>
    </div>
  )
}

export default UuAuthorized