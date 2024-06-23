import { ApiProvider } from "@/services/ApiProvider";
import { I18nProvider } from "@lingui/react";
import { i18n } from "@lingui/core";
import { FC, useEffect } from "react";
import { defaultLocale, dynamicActivate } from "@/i18n";
import { TooltipProvider } from "@/components/ui/tooltip";

export const AppProvider: FC<{ children?: React.ReactNode }> = ({
  children,
}) => {
  useEffect(() => {
    void dynamicActivate(defaultLocale);
  }, []);

  return (
    <I18nProvider i18n={i18n}>
      <TooltipProvider>
        <ApiProvider>{children}</ApiProvider>
      </TooltipProvider>
    </I18nProvider>
  );
};
