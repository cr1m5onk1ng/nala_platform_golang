import React, { useState } from "react";
import { ThemeProvider, createTheme, CssBaseline } from "@material-ui/core";
import Home from "./components/Home/Home";

function App() {
  const [darkMode, setDarkMode] = useState(false);

  const tealTheme = createTheme({
    spacing: 4,
    palette: {
      type: darkMode ? "dark" : "light",
      primary: {
        main: "#80cbc4",
      },
      secondary: {
        main: "#546e7a",
      },
      addButton: {
        main: "#d32f2f",
      },
      background: {
        main: "#e3e3e3",
      },
      contrastThreshold: 3,
      tonalOffset: 0.2,
    },
  });

  const theme = createTheme({
    spacing: 4,
    palette: {
      type: darkMode ? "dark" : "light",
      primary: {
        main: "#f44336",
      },
      secondary: {
        main: "#3EA6FF",
      },
      background: {
        default: darkMode ? "#232323" : "#FFF",
        dark: darkMode ? "#181818" : "#f4f6f8",
        paper: darkMode ? "#232323" : "#FFF",
      },
    },
  });

  return (
    <ThemeProvider theme={tealTheme}>
      <Home darkMode={darkMode} setDarkMode={setDarkMode} />
    </ThemeProvider>
  );
}

export default App;
