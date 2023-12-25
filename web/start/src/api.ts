import axios, {AxiosResponse} from 'axios';

export const API_BASE_URL = import.meta.env.VITE_API_URL;

interface Blueprint {
  type: string;
  spec: any; // Replace 'any' with a more specific type if you have a defined spec structure
}

type FetchBlueprintsResponse = {
  blueprints: Blueprint[];
};

export const fetchBlueprints = async (): Promise<Blueprint[]> => {
  try {
    const response: AxiosResponse<FetchBlueprintsResponse> = await axios.get(`${API_BASE_URL}/v1/blueprints`);
    return response.data.blueprints;
  } catch (error) {
    console.error('Error fetching blueprints:', error);
    throw error;
  }
};

type GenerateResponse = {
  downloadUrl: string;
};

export const generateBlueprint = async (blueprintId: string, values: Record<string, any>): Promise<string> => {
  try {
    const response: AxiosResponse<GenerateResponse> = await axios.post(`${API_BASE_URL}/v1/generate`, {
      blueprintId,
      values
    });
    return response.data.downloadUrl;
  } catch (error) {
    console.error('Error generating project:', error);
    throw error;
  }
};
