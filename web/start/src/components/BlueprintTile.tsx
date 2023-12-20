import React from 'react';
import { Blueprint } from '../types';

interface BlueprintTileProps {
  blueprint: Blueprint;
  onUse: (blueprint: Blueprint) => void;
}

const BlueprintTile: React.FC<BlueprintTileProps> = ({ blueprint, onUse }) => {
  return (
    <div className="bg-gray-900 text-white p-4 rounded-lg shadow-lg">
      <h3 className="text-xl font-bold">{blueprint.spec.name}</h3>
      <p>{blueprint.spec.description}</p>
      <button
        onClick={() => onUse(blueprint)}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-3"
      >
        Use
      </button>
    </div>
  );
};

export default BlueprintTile;
