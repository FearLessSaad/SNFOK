"use client";
import {
  LayoutDashboard,
  Bell,
  ShieldAlert,
  Code,
  HeartHandshake,
  Cog,
  MenuIcon,
  ShieldCheckIcon
} from "lucide-react";
import { BiLogoKubernetes } from "react-icons/bi";
import React, { useState } from "react";
import Logo from "@/components/Logo";
import Link from "next/link";
import { Button, buttonVariants } from "./ui/button";
import { usePathname } from "next/navigation";
import { SheetContent, SheetTrigger, Sheet } from "./ui/sheet";
import { DialogTitle } from "./ui/dialog";

const routes = [
  {
    label: "Dashboard",
    href: "",
    icon: LayoutDashboard,
  },
  {
    label: "Security",
    href: "/security",
    icon: ShieldCheckIcon,
  },
  {
    label: "Threats",
    href: "/threats",
    icon: ShieldAlert,
  },
  {
    label: "Alerts",
    href: "/alerts",
    icon: Bell,
  },
  {
    label: "Ploicies",
    href: "/policies",
    icon: Code,
  },
  {
    label: "Kubernetes",
    href: "/kubernetes",
    icon: BiLogoKubernetes,
  },
  {
    label: "Settings",
    href: "/settings",
    icon: Cog,
  },
];

function DesktopSideBar() {
  const pathname = usePathname();
  const activeRoute =
    routes.find(
      (route) => route.href.length > 0 && pathname.includes(route.href)
    ) || routes[0];
  return (
    <div className="hidden relative md:block min-w-[280px] max-w-[280px] h-screen overflow-hidden w-full bg-primary/5 dark:bg-secondary/30 dark:text-foreground text-muted-foreground border-r-2 border-separate">
      <div className="flex items-center justify-center gap-2 border-b-[1px] border-separate p-4">
        <Logo />
      </div>
      <div className="flex flex-col p-2 gap-1 mt-10">
        {routes.map((route) => {
          return (
            <Link
              href={route.href === ""? "/" : route.href}
              key={route.label}
              className={buttonVariants({
                variant:
                  activeRoute.href === route.href
                    ? "sideBarActive"
                    : "sideItem",
              })}
            >
              <route.icon size={20} />
              {route.label}
            </Link>
          );
        })}
      </div>
    </div>
  );
}

export function MobileSideBar() {
  const pathname = usePathname();
  const activeRoute =
    routes.find(
      (route) => route.href.length > 0 && pathname.includes(route.href)
    ) || routes[0];
  const [isOpen, setOpen] = useState<boolean>(false);
  return (
    <div className="block border-separate bg-background md:hidden">
      <nav className="container flex items-center justify-between pr-8">
        <Sheet open={isOpen} onOpenChange={setOpen}>
          <SheetTrigger asChild>
            <Button variant={"ghost"} size={"icon"}>
              <MenuIcon />
            </Button>
          </SheetTrigger>
          <SheetContent className="w-[400px] sm:w-[540px] p-4" side="left">
            <DialogTitle>
              <Logo />
            </DialogTitle>
            <div className="flex flex-col p-2 gap-1">
              {routes.map((route) => {
                return (
                  <Link
                    href={route.href}
                    onClick={()=>setOpen(!isOpen)}
                    key={route.label}
                    className={buttonVariants({
                      variant:
                        activeRoute.href === route.href
                          ? "sideBarActive"
                          : "sideItem",
                    })}
                  >
                    <route.icon size={20} />
                    {route.label}
                  </Link>
                );
              })}
            </div>
          </SheetContent>
        </Sheet>
      </nav>
    </div>
  );
}

export default DesktopSideBar;
