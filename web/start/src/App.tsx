import React, {useEffect, useState} from 'react';
import BlueprintSelector from './components/BlueprintSelector';
import ConfigurationForm from './components/ConfigurationForm';
import {fetchBlueprints, generateProject} from './api';
import {Blueprint} from "./types.ts";
import GenerateButton from "./components/GenerateButton.tsx";

const App: React.FC = () => {
  const [blueprints, setBlueprints] = useState<Blueprint[]>([]);
  const [selectedBlueprintType, setSelectedBlueprintType] = useState<string | null>(null);
  const [configOptions, setConfigOptions] = useState<Record<string, any>>({});
  const [isGenerating, setIsGenerating] = useState<boolean>(false);

  useEffect(() => {
    const loadBlueprints = async () => {
      const fetchedBlueprints = await fetchBlueprints();
      setBlueprints(fetchedBlueprints);
      if (fetchedBlueprints.length > 0) {
        setSelectedBlueprintType(fetchedBlueprints[0].type);
      }
    };
    loadBlueprints().catch(console.error);
  }, []);

  const handleBlueprintSelect = (type: string) => {
    setSelectedBlueprintType(type);
  };

  const handleFormSubmit = async () => {
    if (!selectedBlueprintType) return;

    setIsGenerating(true);
    try {
      const downloadUrl = await generateProject(selectedBlueprintType, configOptions);
      window.location.href = downloadUrl; // Trigger download
    } catch (error) {
      console.error('Error generating project:', error);
      // Handle error appropriately
    }
    setIsGenerating(false);
  };

  return (
    <div>
      <BlueprintSelector blueprints={blueprints} onSelect={handleBlueprintSelect} selectedType={selectedBlueprintType}/>
      {selectedBlueprintType && (
        <ConfigurationForm blueprint={blueprints.find(b => b.type === selectedBlueprintType) ?? null}
                           onSubmit={setConfigOptions}/>
      )}
      <GenerateButton onClick={handleFormSubmit} isGenerating={isGenerating}/>
    </div>
  );
};

export default App;
