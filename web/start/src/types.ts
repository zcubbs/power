export interface Blueprint {
  spec: Spec;
  type: string;
}

export interface Spec {
  id: string;
  name: string;
  description: string;
  options: Option[];
}

export interface Option {
  id: string;
  name: string;
  description: string;
  type: string;
  default: string;
  choices?: string[];
}
