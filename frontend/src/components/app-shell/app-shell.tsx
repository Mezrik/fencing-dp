import { FC } from "react";
import { AsideMenu } from "./aside-menu";
import { TooltipProvider } from "@/components/ui/tooltip";

export const AppShell: FC<{ children?: React.ReactNode }> = ({ children }) => {
  return (
    <TooltipProvider>
      <div className="grid h-screen w-full pl-[53px]">
        <AsideMenu />
        {children}
      </div>
    </TooltipProvider>
  );
};
