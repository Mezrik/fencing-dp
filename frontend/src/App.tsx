import { useContext, useEffect, useState } from "react";
import { ApiContext, ApiProvider } from "@/services/ApiProvider";
import { CompetitionResult } from "@/generated/server";

import { I18nProvider } from "@lingui/react";
import { i18n } from "@lingui/core";
import { defaultLocale, dynamicActivate } from "./i18n";
import { Trans } from "@lingui/macro";
import { AppShell } from "./components/app-shell/app-shell";

function App() {
  useEffect(() => {
    dynamicActivate(defaultLocale);
  }, []);

  return (
    <I18nProvider i18n={i18n}>
      <ApiProvider>
        <AppShell>
          <Trans>This is a translation</Trans>
        </AppShell>
      </ApiProvider>
    </I18nProvider>
  );
}

export default App;
