import React, {useEffect, useState} from 'react';
import BlueprintTile from './components/BlueprintTile';
import {fetchBlueprints} from './api';
import {Blueprint} from './types';
import {ThemeProvider} from "./components/theme-provider.tsx";
import {ToastProvider} from "@/components/ui/toast.tsx";
import {Toaster} from "@/components/ui/toaster.tsx";
import Logo from './assets/logo.png';
import {Separator} from "@/components/ui/separator.tsx";

const App: React.FC = () => {
  const [blueprints, setBlueprints] = useState<Blueprint[]>([]);
  useEffect(() => {
    document.title = window.VITE_APP_TITLE || 'Power - Starters';
    fetchBlueprints().then(setBlueprints);
  }, []);

  const getNoBlueprints = () => {
    return (
      <div className="text-center text-gray-400 mt-10 p-10">
        <p className="text-2xl font-bold">No blueprints found</p>
        <p className="text-lg">Load plugins or enable built-in blueprints to get started.</p>
      </div>
    );
  };

  const getBlueprints = () => {
    return (
      <div className="grid grid-cols-3 gap-4">
        {blueprints?.map((blueprint) => (
          <BlueprintTile key={blueprint.spec.id} blueprint={blueprint}/>
        ))}
      </div>
    );
  }

  const getLogo = () => {
    const logoUrl = window.VITE_APP_LOGO_URL;
    return (
      <img src={logoUrl || Logo} alt="logo" className="h-12" />
    );
  };

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <ToastProvider>
        <div className="text-white min-h-screen">
          <div className="container mx-auto p-5">
            <div className="mb-5 flex">
              {getLogo()}
              <span className="flex-auto-leading-none ml-2">
                {/* add version */}
              </span>
            </div>
            <Separator/>
            <h2 className="text-3xl font-bold mt-10 mb-10">Plugins</h2>
            {blueprints && blueprints.length > 0 ? getBlueprints() : getNoBlueprints()}
          </div>
        </div>
        <Toaster />
      </ToastProvider>
    </ThemeProvider>
  );
};

export default App;
