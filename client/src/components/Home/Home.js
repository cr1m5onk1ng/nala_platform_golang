import React, { useState } from "react";
import useStyles from "./styles";
import HomeAppBar from "./HomeAppBar/HomeAppBar";
import HomeDrawer from "./HomeDrawer/HomeDrawer";
import { Box } from "@material-ui/core";

function Home({ darkMode, setDarkMode }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <HomeAppBar darkMode={darkMode} setDarkMode={setDarkMode} />
      <Box display="flex">
        <HomeDrawer />
      </Box>
    </div>
  );
}

export default Home;
