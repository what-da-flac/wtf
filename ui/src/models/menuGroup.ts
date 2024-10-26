import { MenuItem } from './menuItem';
import React from 'react';

export type MenuGroup = {
  label: string;
  icon: React.FC<any>;
  initiallyOpened: boolean;
  items: MenuItem[];
};
