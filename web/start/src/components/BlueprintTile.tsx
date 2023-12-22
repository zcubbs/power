import React from 'react';
import {Blueprint} from '../types';
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import BlueprintCustomizationDialog from "@/components/BlueprintCustomizationDialog.tsx";
import {generateProject} from "@/api.ts";

interface BlueprintTileProps {
  blueprint: Blueprint;
}

const BlueprintTile: React.FC<BlueprintTileProps> = ({ blueprint }) => {

  const handleGenerate = async (options: Record<string, any>) => {
    if (!blueprint) return;
    window.location.href = await generateProject(blueprint.spec.id, options); // Trigger download
  };

  return (
    <Card className="rounded-lg shadow-lg">
      <CardHeader>
        <CardTitle className="text-xl font-bold">
          {blueprint.spec.name}
        </CardTitle>
        <CardDescription>{blueprint.spec.description}</CardDescription>
      </CardHeader>
      <CardContent>
        <BlueprintCustomizationDialog blueprint={blueprint} onGenerate={handleGenerate}/>
      </CardContent>
    </Card>
  );
};

export default BlueprintTile;
