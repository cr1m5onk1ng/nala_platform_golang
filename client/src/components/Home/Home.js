import React, { useState } from "react";
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

import { useDispatch } from "react-redux";
import { getPosts } from "../../actions/posts";
import Posts from "../Posts/Posts";

function Home({ darkMode, setDarkMode }) {
  const classes = useStyles();
  const dispatch = useDispatch();
  /*
  useEffect(() => {
    dispatch(getPosts());
  }, [dispatch]);*/
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
        <Container maxWidth="lg">
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
          </Hidden>
          <Box display="flex">
            <Hidden mdDown>
              <HomeDrawer desktopMode={true} />
            </Hidden>
            <Hidden mdUp>
              <HomeDrawer
                visible={drawerVisible}
                onCloseDrawer={closeDrawer}
                desktopMode={false}
              />
            </Hidden>
            <Grow in>
              <Container>
                <Posts />
              </Container>
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
