import React from "react";
import { MenuIcon } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";

function PodActions({pod}:{pod: string}) {


    const onDetailsClick = () => {

    }
    const onIsolateClick = () => {
        console.log(pod)
    }

    const onExposeClick = () => {
        console.log(pod)
    }

    const onSecurityClick = () => {

    }
  return (
    <div className="flex gap-2">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant={"outline"} size={"icon"}>
            <MenuIcon />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuItem onClick={onDetailsClick}>Details</DropdownMenuItem>
          <DropdownMenuItem onClick={onIsolateClick}>Isolate</DropdownMenuItem>
          <DropdownMenuItem onClick={onExposeClick}>Expose</DropdownMenuItem>
          <DropdownMenuItem onClick={onSecurityClick}>Security</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}

export default PodActions;
