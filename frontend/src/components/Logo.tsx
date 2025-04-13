import { cn } from "@/lib/utils";
import { ShieldCheck } from "lucide-react";
import Link from "next/link";
import React from "react";

function Logo({
  fontSize = "text-2xl",
  iconSize = 20,
}: {
  fontSize?: string;
  iconSize?: number;
}) {
  return (
    <Link
      href={"/"}
      className={cn("font-extrabold flex items-center gap-2", fontSize)}
    >
      <div className="flex items-center justify-center">
        <span className="bg-gradient-to-r from-emerald-500 to-emerald-600 bg-clip-text text-transparent ">
          SNF
        </span>
        <ShieldCheck className="font-extrabold text-stone-700 dark:text-stone-300" />
        <span className="bg-gradient-to-r from-emerald-500 to-emerald-600 bg-clip-text text-transparent">
          K
        </span>
      </div>
    </Link>
  );
}

export default Logo;
