import { FC, useEffect, useState } from 'react';
import { I18nProvider } from '@lingui/react';
import { i18n } from '@lingui/core';
import { HelmetProvider } from 'react-helmet-async';

import { TooltipProvider } from '@/components/ui/tooltip';
import { defaultLocale, dynamicActivate } from '@/i18n';
import { queryConfig } from '@/lib/react-query';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Toaster } from '@/components/ui/toaster';

export const AppProvider: FC<{ children?: React.ReactNode }> = ({ children }) => {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: queryConfig,
      }),
  );

  useEffect(() => {
    void dynamicActivate(defaultLocale);
  }, []);

  return (
    <I18nProvider i18n={i18n}>
      <HelmetProvider>
        <QueryClientProvider client={queryClient}>
          <TooltipProvider>
            {children}
            <Toaster />
          </TooltipProvider>
        </QueryClientProvider>
      </HelmetProvider>
    </I18nProvider>
  );
};
