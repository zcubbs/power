import React from 'react';
import {Blueprint} from '../types';
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import BlueprintCustomizationDialog from "@/components/BlueprintCustomizationDialog.tsx";
import {generateProject} from "@/api.ts";
import {useToast} from "@/components/ui/use-toast.ts";

interface BlueprintTileProps {
  blueprint: Blueprint;
}

const BlueprintTile: React.FC<BlueprintTileProps> = ({ blueprint }) => {
  const { toast } = useToast();

  const handleGenerate = async (options: Record<string, string>) => {
    if (!blueprint) return;
    // ensure options are all strings
    Object.keys(options).forEach(key => {
      options[key] = options[key].toString();
    });

    // trigger download using a link
    const link = document.createElement('a');
    link.href = await generateProject(blueprint.spec.id, options);
    link.setAttribute('download', `${blueprint.spec.name}.zip`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    // show toast
    toast({
      title: "Blueprint Generated",
      description: `Your '${blueprint.spec.name}' blueprint has been generated and is downloading.`,
      duration: 5000,
      variant: "success",
    });
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
