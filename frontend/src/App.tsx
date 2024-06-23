import { useContext, useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { ApiContext, ApiProvider } from "@/services/ApiProvider";
import { CompetitionResult } from "@/generated/server";

import { I18nProvider } from "@lingui/react";
import { i18n } from "@lingui/core";
import { defaultLocale, dynamicActivate } from "./i18n";
import { Trans } from "@lingui/macro";

const TestComponent = () => {
  const api = useContext(ApiContext);
  const [competitions, setCompetitions] = useState<CompetitionResult[]>([]);

  const getCompetitions = async () => {
    const competitions = await api?.GetCompetitions();
    competitions && setCompetitions(competitions);
  };

  useEffect(() => {
    getCompetitions();
  }, []);

  return (
    <div>
      {competitions.map((competition, i) => (
        <div key={i}>
          <h2>{competition.name}</h2>
        </div>
      ))}
    </div>
  );
};

function App() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    dynamicActivate(defaultLocale);
  }, []);

  return (
    <I18nProvider i18n={i18n}>
      <ApiProvider>
        <div>
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <h1>Vite + React</h1>
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
        <TestComponent />
        <Trans>This is a translation</Trans>
      </ApiProvider>
    </I18nProvider>
  );
}

export default App;
