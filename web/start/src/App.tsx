import React, {useEffect, useState} from 'react';
import BlueprintTile from './components/BlueprintTile';
import {fetchBlueprints} from './api';
import {Blueprint} from './types';
import {ThemeProvider} from "./components/theme-provider.tsx";
import {ToastProvider} from "@/components/ui/toast.tsx";
import {Toaster} from "@/components/ui/toaster.tsx";

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
            <h1 className="text-3xl font-bold mb-5">Blueprints</h1>
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
