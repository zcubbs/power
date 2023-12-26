import React from 'react';
import {Blueprint} from '../types';
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import BlueprintCustomizationDialog from "@/components/BlueprintCustomizationDialog.tsx";
import {generateBlueprint} from "@/api.ts";
import {useToast} from "@/components/ui/use-toast.ts";
import {Badge} from "@/components/ui/badge.tsx";

interface BlueprintTileProps {
  blueprint: Blueprint;
}

const BlueprintTile: React.FC<BlueprintTileProps> = ({ blueprint }) => {
  const { toast } = useToast();

  const handleGenerate = async (values: Record<string, string>) => {
    if (!blueprint) return;
    // ensure options are all strings
    Object.keys(values).forEach(key => {
      values[key] = values[key].toString();
    });

    // trigger download using a link
    const link = document.createElement('a');
    link.href = await generateBlueprint(blueprint.spec.id, values);
    link.setAttribute('download', `${blueprint.spec.name}.zip`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    // show toast
    toast({
      title: "Blueprint Generated",
      description: `Your '${blueprint.spec.name}' blueprint has been generated and is downloading.`,
      duration: 5000,
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
        <Badge className="mr-2">{blueprint.type}</Badge>
        <BlueprintCustomizationDialog blueprint={blueprint} onGenerate={handleGenerate}/>
      </CardContent>
    </Card>
  );
};

export default BlueprintTile;
