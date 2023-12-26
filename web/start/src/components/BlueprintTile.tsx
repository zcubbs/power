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

const BlueprintTile: React.FC<BlueprintTileProps> = ({blueprint}) => {
  const {toast} = useToast();

  const handleGenerate = async (values: Record<string, string>) => {
    if (!blueprint) return;
    // ensure options are all strings
    Object.keys(values).forEach(key => {
      values[key] = values[key].toString();
    });

    // trigger download using a link
    try {
      // Trigger download using a link
      const link = document.createElement('a');
      link.href = await generateBlueprint(blueprint.spec.id, values);
      link.setAttribute('download', `${blueprint.spec.name}.zip`);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);

      // Show success toast
      toast({
        title: "Blueprint Generated",
        description: `Your '${blueprint.spec.name}' blueprint has been generated and is downloading.`,
        duration: 5000,
      });
    } catch (error) {
      // Show error toast
      toast({
        title: "Blueprint Generation Failed",
        description: `An error occurred while generating the blueprint: ${error}`,
        duration: 5000,
        variant: 'destructive'
      });
    }

  };

  const getColoredBadge = () => {
    if (blueprint.type === 'built-in') {
      return <Badge variant="secondary">Built-in</Badge>;
    } else if (blueprint.type === 'plugin') {
      return <Badge variant="default">Plugin</Badge>;
    } else if (blueprint.type === 'registrar') {
      return <Badge variant="default">Registrar</Badge>;
    } else {
      return <Badge variant="outline">Unknown</Badge>;
    }
  }

  return (
    <Card className="rounded-lg shadow-lg">
      <CardHeader>
        <CardTitle className="text-xl font-bold">
          {blueprint.spec.name}
        </CardTitle>
        <CardDescription>{blueprint.spec.description}</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-2">
          <div className="flex items-center justify-between col-span-2">
            <div className="flex items-center">
              {getColoredBadge()}
            </div>
            <div className="flex items-center justify-end col-span-2">
              <BlueprintCustomizationDialog blueprint={blueprint} onGenerate={handleGenerate}/>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default BlueprintTile;
