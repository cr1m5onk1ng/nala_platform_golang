import React, { useState } from "react";
import {
  Toolbar,
  Drawer,
  Button,
  List,
  Typography,
  Divider,
  ListItem,
  ListItemIcon,
  ListItemText,
  ListSubheader,
  Box,
  Hidden,
} from "@material-ui/core";

import FavoriteBorderIcon from "@material-ui/icons/FavoriteBorder";
import LocalLibraryIcon from "@material-ui/icons/LocalLibrary";
import RssFeedIcon from "@material-ui/icons/RssFeed";
import AddIcon from "@material-ui/icons/Add";
import { MEDIA_TYPES } from "../../../constants/categories";
import { DIFFICULTIES } from "../../../constants/categories";

import useStyles from "./styles";

function HomeDrawer({ visible, onCloseDrawer, desktopMode }) {
  const classes = useStyles();

  return (
    <Drawer
      open={!desktopMode ? visible : true}
      onClose={onCloseDrawer !== undefined ? onCloseDrawer() : null}
      anchor="left"
      variant={desktopMode ? "permanent" : "temporary"}
      className={classes.drawer}
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.drawerContainer}>
        <Divider />
        <List>
          <ListItem button classes={{ root: classes.listItem }}>
            <ListItemIcon>
              {<RssFeedIcon className={classes.listItemFeed} />}
            </ListItemIcon>
            <ListItemText
              classes={{
                primary: classes.listItemText,
              }}
              primary={"Feed"}
            />
          </ListItem>
          <ListItem button classes={{ root: classes.listItem }}>
            <ListItemIcon>
              {<LocalLibraryIcon className={classes.listItemStudy} />}
            </ListItemIcon>
            <ListItemText
              classes={{
                primary: classes.listItemText,
              }}
              primary={"Study Plans"}
            />
          </ListItem>
          <ListItem button classes={{ root: classes.listItem }}>
            <ListItemIcon>
              {<FavoriteBorderIcon className={classes.listItemFavorites} />}
            </ListItemIcon>
            <ListItemText
              classes={{
                primary: classes.listItemText,
              }}
              primary={"Favorites"}
            />
          </ListItem>
        </List>
        <Divider />

        <List
          component="nav"
          aria-labelledby="nested-list-subheader"
          subheader={
            <ListSubheader
              component="div"
              id="nested-list-subheader"
              className={classes.subheader}
            >
              Difficulty
            </ListSubheader>
          }
        >
          {DIFFICULTIES.map((diff) => (
            <ListItem key={diff.id} button classes={{ root: classes.listItem }}>
              <ListItemIcon>
                <AddIcon />
              </ListItemIcon>
              <ListItemText
                classes={{
                  primary: classes.listItemText,
                }}
                primary={diff.difficulty}
              />
            </ListItem>
          ))}
        </List>
        <Divider />
        <List
          component="nav"
          aria-labelledby="nested-list-subheader"
          subheader={
            <ListSubheader
              component="div"
              id="nested-list-subheader"
              className={classes.subheader}
            >
              Media Type
            </ListSubheader>
          }
        >
          {MEDIA_TYPES.map((media) => (
            <ListItem
              key={MediaKeySession.id}
              button
              classes={{ root: classes.listItem }}
            >
              <ListItemIcon>
                <AddIcon />
              </ListItemIcon>
              <ListItemText
                classes={{
                  primary: classes.listItemText,
                }}
                primary={media.mediaType}
              />
            </ListItem>
          ))}
        </List>
      </div>
    </Drawer>
  );
}

export default HomeDrawer;
