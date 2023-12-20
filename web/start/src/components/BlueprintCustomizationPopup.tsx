import React, { useState, useEffect } from 'react';
import { Blueprint, Option } from '../types';

interface BlueprintCustomizationPopupProps {
  blueprint: Blueprint | null;
  onClose: () => void;
  onGenerate: (options: Record<string, any>) => void;
}

const BlueprintCustomizationPopup: React.FC<BlueprintCustomizationPopupProps> = ({ blueprint, onClose, onGenerate }) => {
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
    setOptions({ ...options, [name]: value });
  };

  if (!blueprint) return null;

  const renderInputField = (option: Option) => {
    switch (option.type) {
      case 'text':
        return <input type="text" value={options[option.name]} onChange={(e) => handleOptionChange(option.name, e.target.value)} />;
      case 'number':
        return <input type="number" value={options[option.name]} onChange={(e) => handleOptionChange(option.name, e.target.value)} />;
      case 'select':
        return (
          <select value={options[option.name]} onChange={(e) => handleOptionChange(option.name, e.target.value)}>
            {option.options?.map((opt, idx) => (
              <option key={idx} value={opt}>{opt}</option>
            ))}
          </select>
        );
      // Add other cases for different types of inputs
      default:
        return null;
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full">
      <div className="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800">
        <h3 className="text-lg font-bold">{blueprint.spec.name}</h3>
        <p>{blueprint.spec.description}</p>
        {blueprint.spec.options.map((option, idx) => (
          <div key={idx} className="mb-4">
            <label className="block mb-2">{option.name}</label>
            {renderInputField(option)}
          </div>
        ))}
        <button onClick={() => onGenerate(options)} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-3">
          Generate
        </button>
        <button onClick={onClose} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mt-3">
          Close
        </button>
      </div>
    </div>
  );
};

export default BlueprintCustomizationPopup;
