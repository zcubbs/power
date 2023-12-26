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
    fetchBlueprints().then(setBlueprints);
  }, []);

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <ToastProvider>
        <div className="text-white min-h-screen">
          <div className="container mx-auto p-5">
            <img src={Logo} alt="logo" className="h-12 mb-5"/>
            <Separator/>
            <h2 className="text-2xl font-bold mt-10 mb-5">Blueprints</h2>
            <div className="grid grid-cols-3 gap-4">
              {blueprints.map((blueprint) => (
                <BlueprintTile key={blueprint.spec.id} blueprint={blueprint}/>
              ))}
            </div>
          </div>
        </div>
        <Toaster />
      </ToastProvider>
    </ThemeProvider>
  );
};

export default App;
