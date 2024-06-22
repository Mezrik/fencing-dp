import React, { createContext } from "react";
import { Api } from "./Api";
import { RestApi } from "./rest/RestApi";
import { DesktopApi } from "./desktop/DesktopApi";

export const ApiContext = createContext<Api | null>(null);

export const ApiProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const api =
    import.meta.env.MODE === "desktop" ? new DesktopApi() : new RestApi();

  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
};
