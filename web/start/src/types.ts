export interface Blueprint {
  spec: Spec;
}

export interface Spec {
  id: string;
  name: string;
  description: string;
  options: Option[];
}

export interface Option {
  name: string;
  description: string;
  type: string;
  default: string;
  options?: string[];
}
