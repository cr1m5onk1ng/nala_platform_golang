import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import useStyles from "./styles";
import HomeAppBar from "./HomeAppBar/HomeAppBar";
import HomeDrawer from "./HomeDrawer/HomeDrawer";
import {
  Box,
  Hidden,
  Container,
  Grow,
  Grid,
  CssBaseline,
} from "@material-ui/core";
import FormPopUp from "./HomeAppBar/Form/FormPopUp";
import PostForm from "./HomeAppBar/Form/PostForm";

import { useDispatch, useSelector } from "react-redux";
import { getPosts } from "../../actions/posts";
import Posts from "../Posts/Posts";
import Sidebar from "./Sidebar/Sidebar";
import axios from "axios";
import { decode } from "paseto";

function Home({ darkMode, setDarkMode }) {
  const classes = useStyles();
  const dispatch = useDispatch();
  const history = useHistory();

  // USER STATE
  const userState = useSelector((state) => state.user);

  useEffect(() => {
    localStorage.setItem("user", JSON.stringify(userState));
  }, [userState]);
  /*
  useEffect(() => {
    dispatch(getPosts());
  }, [dispatch]);*/

  // MENUS STATE
  const [currentId, setCurrentId] = useState(0);

  const [drawerVisible, setDrawerVisibile] = useState(false);

  const [formOpen, setFormOpen] = useState(false);

  const openFormDialog = () => {
    setFormOpen(true);
  };

  const closeFormDialog = () => {
    setFormOpen(false);
  };

  const toggleDrawer = () => (event) => {
    setDrawerVisibile(!drawerVisible);
  };

  const closeDrawer = () => (event) => {
    setDrawerVisibile(false);
  };

  return (
    <>
      <CssBaseline />
      <div className={classes.home}>
        <Container maxWidth="xl">
          <Hidden mdUp>
            <HomeAppBar
              darkMode={darkMode}
              setDarkMode={setDarkMode}
              toggleDrawer={toggleDrawer}
              onOpenForm={openFormDialog}
              onCloseForm={closeFormDialog}
              desktopMode={false}
            />
          </Hidden>
          <Hidden mdDown>
            <HomeAppBar
              darkMode={darkMode}
              setDarkMode={setDarkMode}
              toggleDrawer={toggleDrawer}
              onOpenForm={openFormDialog}
              onCloseForm={closeFormDialog}
              desktopMode={true}
            />
            <Sidebar />
          </Hidden>
          <Box className={classes.bodyContainer}>
            <Hidden mdUp>
              <HomeDrawer
                visible={drawerVisible}
                onCloseDrawer={closeDrawer}
                desktopMode={false}
              />
            </Hidden>
            <Grow in>
              <Posts />
            </Grow>
          </Box>
        </Container>
        <FormPopUp
          title={"Add a resource by pasting a valid url of the content"}
          open={formOpen}
          onClose={closeFormDialog}
        >
          <PostForm currentId={currentId} setCurrentId={setCurrentId} />
        </FormPopUp>
      </div>
    </>
  );
}

export default Home;
