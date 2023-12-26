import axios, {AxiosError, AxiosResponse} from 'axios';

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

// Define the expected structure of the error response
interface ErrorResponse {
  message: string;
  // Include other properties that might be in the error response
}

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
  } catch (error: unknown) {
    let errorMessage = 'Failed to generate blueprint';

    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError<ErrorResponse>; // Assert the error type

      if (axiosError.response) {
        const { status, data } = axiosError.response;
        console.error(`Error generating project: ${status} - ${data?.message}`);
        errorMessage = data?.message || 'Error occurred during generation';
      } else {
        console.error(`Error generating project: ${axiosError.message}`);
        errorMessage = axiosError.message;
      }
    } else {
      // Non-Axios errors handling
      console.error('Non-Axios error occurred:', error);
    }

    throw new Error(errorMessage);
  }
};
