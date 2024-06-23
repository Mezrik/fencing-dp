import { FC } from "react";
import { AsideMenu } from "./aside-menu";

export const AppShell: FC<{ children?: React.ReactNode }> = ({ children }) => {
  return (
    <div className="grid h-screen w-full pl-[53px]">
      <AsideMenu />
      {children}
    </div>
  );
};
