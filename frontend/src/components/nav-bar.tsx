import React from "react";
import BreadCrumbHeader from "./BreadCrumbHeader";
import { ModeToggle } from "./ThemeModeToggle";
import { AvatarFallback, Avatar, AvatarImage } from "./ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import { LogOutIcon, UserCircle2Icon } from "lucide-react";

function NavBar() {
  return (
    <header className="flex items-center justify-between px-6 py-4 h-[50px]">
      <BreadCrumbHeader />
      <div className="gap-3 flex items-center">
        <ModeToggle />
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Avatar className="border-2 border-primary">
              <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
              <AvatarFallback>MS</AvatarFallback>
            </Avatar>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>Muhammad Salman</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem><UserCircle2Icon /> Profile</DropdownMenuItem>
            <DropdownMenuItem> Rotate Token</DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem><LogOutIcon/> Logout</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </header>
  );
}

export default NavBar;
