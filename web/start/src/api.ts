import axios from 'axios';

const API_BASE_URL = 'http://localhost:8000';

interface Blueprint {
  type: string;
  spec: any; // Replace 'any' with a more specific type if you have a defined spec structure
}

export const fetchBlueprints = async (): Promise<Blueprint[]> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/v1/blueprints`);
    return response.data.blueprints;
  } catch (error) {
    console.error('Error fetching blueprints:', error);
    throw error;
  }
};

export const generateProject = async (blueprintType: string, options: Record<string, any>): Promise<string> => {
  try {
    const response = await axios.post(`${API_BASE_URL}/v1/generate`, { blueprintType, options });
    return response.data.downloadUrl;
  } catch (error) {
    console.error('Error generating project:', error);
    throw error;
  }
};
