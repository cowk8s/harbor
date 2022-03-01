export enum StyleMode {
  DARK = 'DARK',
  LIGHT = 'LIGHT'
}

export const HAS_STYLE_MODE: string = 'styleModeLocal';

export interface ThemeInterface {
  showStyle: string;
  mode: string
}

export const THEME_ARRAY: ThemeInterface[] = [
  {
    showStyle: StyleMode.DARK,
    mode: StyleMode.LIGHT
  },
  {
    showStyle: StyleMode.LIGHT,
    mode: StyleMode.DARK
  }
]