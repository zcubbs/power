import React, { useEffect, useState } from 'react';
import { Blueprint, Option } from '../types';
import { Switch } from "@/components/ui/switch";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Combobox } from "@/components/ui/combobox";
import { Input } from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";

interface BlueprintCustomizationDialogProps {
  blueprint: Blueprint | null;
  onGenerate: (options: Record<string, string>) => void;
}

const BlueprintCustomizationDialog: React.FC<BlueprintCustomizationDialogProps> = ({ blueprint, onGenerate }) => {
  const [options, setOptions] = useState<Record<string, string>>({});
  const [nameToIdMap, setNameToIdMap] = useState<Record<string, string>>({});
  const [open, setOpen] = useState<boolean>(false);

  useEffect(() => {
    if (blueprint) {
      const initialOptions: Record<string, string> = {};
      const newNameToIdMap: Record<string, string> = {};
      blueprint.spec.options.forEach((option: Option) => {
        initialOptions[option.id] = option.default || '';
        newNameToIdMap[option.name] = option.id;
      });
      setOptions(initialOptions);
      setNameToIdMap(newNameToIdMap);
    }
  }, [blueprint]);

  const handleOptionChange = (name: string, value: string) => {
    const optionId = nameToIdMap[name];
    setOptions(prevOptions => ({ ...prevOptions, [optionId]: value }));
  };

  const handleGenerate = () => {
    setOpen(false);
    onGenerate(options);
  };

  const renderInputField = (option: Option) => {
    const handleComboboxChange = (selectedValue: string) => {
      handleOptionChange(option.name, selectedValue);
    };

    const handleSelectChange = (selectedValue: string) => {
      handleOptionChange(option.name, selectedValue);
    }

    switch (option.type) {
      case 'text':
      case 'number':
        return (
          <Input
            type={option.type}
            value={options[option.id] || ''}
            onChange={(e) => handleOptionChange(option.name, e.target.value)}
          />
        );
      case 'combobox':
        return (
          <Combobox
            defaultValue={option.default}
            placeholder="Select an option"
            options={option.choices ?? []} // Ensure this array is populated
            onChange={handleComboboxChange}
          />
        );
      case 'select':
        return (
          <Select defaultValue={option.default} onValueChange={handleSelectChange}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder={option.default} />
            </SelectTrigger>
            <SelectContent>
              {option.choices?.map((choice: string) => (
                <SelectItem key={choice} value={choice}>
                  {choice}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        );
      case 'boolean':
        return (
          <Switch
            id={`switch-${option.id}`}
            checked={options[option.id] === 'true'}
            onCheckedChange={(isChecked) => handleOptionChange(option.name, isChecked ? 'true' : 'false')}
          />
        );
      default:
        return null;
    }
  };

  if (!blueprint) {
    return null;
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="secondary">
          Use
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Blueprint: {blueprint.spec.name}</DialogTitle>
          <DialogDescription>{blueprint.spec.description}</DialogDescription>
        </DialogHeader>
        <Separator />
        {blueprint.spec.options.map((option: Option) => (
          <div key={option.id} className="grid grid-cols-2 gap-4">
            <div>{option.name}</div>
            <div>{renderInputField(option)}</div>
          </div>
        ))}
        <Separator />
        <DialogFooter>
          <Button onClick={handleGenerate}>Generate</Button>
          <DialogClose asChild>
            <Button variant="secondary">Close</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default BlueprintCustomizationDialog;
