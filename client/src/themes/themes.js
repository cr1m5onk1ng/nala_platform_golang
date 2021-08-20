import { createTheme } from "@material-ui/core";

export const tealTheme = createTheme({
  spacing: 4,
  palette: {
    type: darkMode ? "dark" : "light",
    primary: {
      main: "#80cbc4",
    },
    secondary: {
      main: "#90a4ae",
    },
    contrastThreshold: 3,
    tonalOffset: 0.2,
  },
});
