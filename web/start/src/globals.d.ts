// Extends the Window interface to include your custom environment variables
declare global {
  interface Window {
    VITE_APP_LOGO_URL?: string;
    VITE_API_URL?: string;
    VITE_APP_TITLE?: string;
  }
}

// If this file has no imports/exports, ensure it's still treated as a module
export {};
