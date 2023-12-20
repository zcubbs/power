import React from 'react';

interface GenerateButtonProps {
  onClick: () => void; // Function to be called when the button is clicked
  isGenerating: boolean; // A flag to indicate if the generation process is ongoing
}

const GenerateButton: React.FC<GenerateButtonProps> = ({ onClick, isGenerating }) => {
  return (
    <button
      onClick={onClick}
      disabled={isGenerating}
      className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
    >
      {isGenerating ? 'Generating...' : 'Generate Project'}
    </button>
  );
};

export default GenerateButton;
