import React from "react";

import { AppBar, Toolbar, IconButton, Button, Switch } from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import AccountCircle from "@material-ui/icons/AccountCircle";
import useStyles from "./styles";
import HomeSearchField from "./HomeSearchField/HomeSearchField";

function HomeAppBar({ darkMode, setDarkMode }) {
  const classes = useStyles();
  return (
    <AppBar color="inherit" className={classes.appBar}>
      <Toolbar>
        <IconButton edge="start" className={classes.menuIcon} aria-label="menu">
          <MenuIcon />
        </IconButton>
        <div className={classes.grow} />
        <HomeSearchField />
        <div className={classes.grow} />
        <Switch
          value={darkMode}
          onChange={() => setDarkMode(!darkMode)}
          className={classes.icons}
        />

        <Button
          startIcon={<AccountCircle />}
          variant="outlined"
          color="secondary"
        >
          Login
        </Button>
      </Toolbar>
    </AppBar>
  );
}

export default HomeAppBar;
