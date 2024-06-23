import { Trans } from "@lingui/macro";
import { AppShell } from "@/components/app-shell/app-shell";
import { AppProvider } from "./app-provider";

function App() {
  return (
    <AppProvider>
      <AppShell>
        <Trans>This is a translation</Trans>
      </AppShell>
    </AppProvider>
  );
}

export default App;
