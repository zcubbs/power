import React from 'react';
import {Blueprint} from "../types.ts";

interface BlueprintSelectorProps {
  blueprints: Blueprint[];
  onSelect: (blueprintType: string) => void;
  selectedType: string | null;
}

const BlueprintSelector: React.FC<BlueprintSelectorProps> = ({ blueprints, onSelect, selectedType }) => {
  return (
    <select value={selectedType || ''} onChange={(e) => onSelect(e.target.value)}>
      {blueprints.map((blueprint) => (
        <option key={blueprint.type} value={blueprint.type}>
          {blueprint.type}
        </option>
      ))}
    </select>
  );
};

export default BlueprintSelector;
