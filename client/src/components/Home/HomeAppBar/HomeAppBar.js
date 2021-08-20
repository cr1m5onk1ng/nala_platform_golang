import React from "react";

import {
  AppBar,
  Toolbar,
  IconButton,
  Button,
  Switch,
  Hidden,
  Typography,
} from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import AccountCircle from "@material-ui/icons/AccountCircle";
import useStyles from "./styles";
import HomeSearchField from "./HomeSearchField/HomeSearchField";
import Add from "@material-ui/icons/Add";

function HomeAppBar({
  darkMode,
  setDarkMode,
  toggleDrawer,
  onOpenForm,
  desktopMode,
}) {
  const classes = useStyles();

  return (
    <AppBar
      color="primary"
      className={desktopMode ? classes.appBar : classes.mAppBar}
    >
      <Toolbar className={classes.toolbar}>
        <Hidden mdUp>
          <IconButton
            edge="start"
            className={classes.menuIcon}
            aria-label="menu"
            onClick={toggleDrawer()}
          >
            <MenuIcon />
          </IconButton>
        </Hidden>
        <Hidden mdDown>
          <Typography variant="h4" className={classes.logo}>
            NaLa
          </Typography>
        </Hidden>
        <div className={classes.grow} />

        {/*<HomeSearchField />*/}

        <div className={classes.grow} />
        <Hidden mdDown>
          <Button
            onClick={onOpenForm}
            className={`${classes.menuIcon} ${classes.addIcon}`}
            startIcon={<Add />}
            variant="contained"
            color="secondary"
          >
            Add Post
          </Button>
        </Hidden>
        <Button
          className={classes.menuIcon}
          startIcon={<AccountCircle />}
          variant="outlined"
          color="inherit"
        >
          Login
        </Button>
        <Hidden smDown>
          <Switch
            value={darkMode}
            onChange={() => setDarkMode(!darkMode)}
            className={classes.icons}
          />
        </Hidden>
      </Toolbar>
    </AppBar>
  );
}

export default HomeAppBar;
