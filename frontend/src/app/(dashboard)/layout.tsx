import BreadCrumbHeader from '@/components/BreadCrumbHeader'
import DashboardLoadingProvider from '@/components/providers/dashboard-loading-provider'
import DesktopSideBar from '@/components/SideBar'
import { ModeToggle } from '@/components/ThemeModeToggle'
import { Separator } from '@/components/ui/separator'
import React from 'react'

function layout({children}: {children: React.ReactNode}) {
  return (
    <div className='flex h-screen'>
        <DesktopSideBar/>
        <div className='flex flex-col flex-1 min-h-screen'>
            <header className='flex items-center justify-between px-6 py-4 h-[50px]'>
                <BreadCrumbHeader/>
                <div className='gap-3 flex items-center'>
                    <ModeToggle/>
                </div>
            </header>
            <Separator />
            <div className='overflow-auto'>
                <div className='flex-1 px-8 w-full py-4 text-accent-foreground'>
                    <DashboardLoadingProvider>
                        {children}
                    </DashboardLoadingProvider>
                </div>
            </div>
        </div>
    </div>
  )
}

export default layout