import React, {useEffect, useState} from 'react';
import {Blueprint, Option} from '../types';
import {Switch} from "@/components/ui/switch.tsx";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle, DialogTrigger
} from "@/components/ui/dialog.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Separator} from "@/components/ui/separator.tsx";
import {Combobox} from "@/components/ui/combobox.tsx";

interface BlueprintCustomizationDialogProps {
  blueprint: Blueprint | null;
  onGenerate: (options: Record<string, any>) => void;
}

const BlueprintCustomizationDialog: React.FC<BlueprintCustomizationDialogProps> = ({blueprint, onGenerate}) => {
  const [options, setOptions] = useState<Record<string, any>>({});

  useEffect(() => {
    // Initialize options with defaults
    if (blueprint) {
      const initialOptions = blueprint.spec.options.reduce((acc: Record<string, any>, option: Option) => {
        acc[option.name] = option.default || '';
        return acc;
      }, {});
      setOptions(initialOptions);
    }
  }, [blueprint]);

  const handleOptionChange = (name: string, value: any) => {
    setOptions({...options, [name]: value});
  };

  if (!blueprint) return null;

  const renderInputField = (option: Option) => {
    switch (option.type) {
      case 'text':
        return <input type="text" value={options[option.name]}
                      onChange={(e) => handleOptionChange(option.name, e.target.value)}/>;
      case 'number':
        return <input type="number" value={options[option.name]}
                      onChange={(e) => handleOptionChange(option.name, e.target.value)}/>;
      case 'select':
        return (
          <Combobox defaultValue={option.default ?? options[option.name]}
                    placeholder="Select an option"
                    options={option.options ?? []}
                    onChange={() => handleOptionChange(option.name, !options[option.name])}
          >
          </Combobox>
        );
      case 'boolean':
        return (
          <div className="flex items-center space-x-2">
            <Switch id="airplane-mode"
                    checked={options[option.name]}
                    onCheckedChange={() => handleOptionChange(option.name, !options[option.name])}
            />
          </div>
        );
      default:
        return null;
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">Use</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle className="text-xl font-bold">
            Blueprint: {blueprint.spec.name}
          </DialogTitle>
          <DialogDescription>
            {blueprint.spec.description}
          </DialogDescription>
        </DialogHeader>
        <Separator className="my-4"/>
        {/* grid with 2 columns and a row per loop iteration */}
        {blueprint.spec.options.map((option) => (
          <div key={option.name} className="grid grid-cols-2 gap-4">
            <div className="text-sm font-bold">{option.name}</div>
            <div>{renderInputField(option)}</div>
          </div>
        ))}
        <Separator className="my-4"/>
        <DialogFooter className="sm:justify-start">
          <Button className="btn btn-primary" onClick={() => onGenerate(options)}>
            Generate
          </Button>
          <DialogClose asChild>
            <Button type="button" variant="secondary">
              Close
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default BlueprintCustomizationDialog;
