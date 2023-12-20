import React, { useState } from 'react';
import {Blueprint} from "../types.ts";

interface ConfigurationFormProps {
  blueprint: Blueprint | null;
  onSubmit: (options: Record<string, any>) => void;
}

const ConfigurationForm: React.FC<ConfigurationFormProps> = ({ blueprint, onSubmit }) => {
  const [options, setOptions] = useState<Record<string, any>>({});

  const handleChange = (key: string, value: any) => {
    setOptions((prev) => ({ ...prev, [key]: value }));
  };

  if (!blueprint) return null;

  return (
    <form onSubmit={(e) => {
      e.preventDefault();
      onSubmit(options);
    }}>
      {Object.keys(blueprint.spec).map((key) => (
        <div key={key}>
          <label>{key}</label>
          <input type="text" onChange={(e) => handleChange(key, e.target.value)} />
        </div>
      ))}
      <button type="submit">Generate Project</button>
    </form>
  );
};

export default ConfigurationForm;
