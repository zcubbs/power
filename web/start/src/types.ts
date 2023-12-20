export interface Blueprint {
  type: string;
  spec: Spec;
}

export interface Spec {
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
