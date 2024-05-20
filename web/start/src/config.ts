// src/utils/config.js
function getConfig(key: string) {
  // Check if the app is running in production
  if (import.meta.env.PROD) {
    // Production environment: Read from global window object
    return window[key];
  } else {
    // Development environment: Use Vite's import.meta.env
    return import.meta.env[key];
  }
}

export default getConfig;
