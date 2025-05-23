"use client";

import { TooltipProvider } from "@radix-ui/react-tooltip";
import React from "react";
import { Tooltip, TooltipContent, TooltipTrigger } from "./ui/tooltip";

function TooltipWrapper({
  content,
  children,
  side,
}: {
  content: string;
  children: React.ReactNode;
  side?: "top" | "bottom" | "left" | "right";
}) {
  return (
    <TooltipProvider delayDuration={0}>
        <Tooltip>
            <TooltipTrigger asChild>
                {children}
            </TooltipTrigger>
            <TooltipContent side={side}>
                {content}
            </TooltipContent>
        </Tooltip>
    </TooltipProvider>
  );
}

export default TooltipWrapper;
