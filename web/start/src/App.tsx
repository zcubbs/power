import React, { useState, useEffect } from 'react';
import BlueprintTile from './components/BlueprintTile';
import BlueprintCustomizationPopup from './components/BlueprintCustomizationPopup';
import { fetchBlueprints, generateProject } from './api';
import { Blueprint } from './types';

const App: React.FC = () => {
  const [blueprints, setBlueprints] = useState<Blueprint[]>([]);
  const [selectedBlueprint, setSelectedBlueprint] = useState<Blueprint | null>(null);
  const [isPopupOpen, setIsPopupOpen] = useState<boolean>(false);

  useEffect(() => {
    fetchBlueprints().then(setBlueprints);
  }, []);

  const handleUseBlueprint = (blueprint: Blueprint) => {
    setSelectedBlueprint(blueprint);
    setIsPopupOpen(true);
  };

  const handleGenerate = async (options: Record<string, any>) => {
    if (!selectedBlueprint) return;
    const downloadUrl = await generateProject(selectedBlueprint.spec.id, options);
    window.location.href = downloadUrl; // Trigger download
    setIsPopupOpen(false);
  };

  return (
    <div className="bg-gray-900 text-white min-h-screen">
      <div className="container mx-auto p-4">
        <h1 className="text-3xl font-bold">Blueprints</h1>
        <div className="grid grid-cols-3 gap-4">
          {blueprints.map((blueprint) => (
            <BlueprintTile key={blueprint.spec.id} blueprint={blueprint} onUse={handleUseBlueprint} />
          ))}
        </div>
      </div>
      {isPopupOpen && (
        <BlueprintCustomizationPopup
          blueprint={selectedBlueprint}
          onClose={() => setIsPopupOpen(false)}
          onGenerate={handleGenerate}
        />
      )}
    </div>
  );
};

export default App;
