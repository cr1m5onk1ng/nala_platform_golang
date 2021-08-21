import React from "react";

import {
  List,
  Divider,
  ListItem,
  ListItemIcon,
  ListItemText,
  ListSubheader,
  Toolbar,
  Paper,
} from "@material-ui/core";

import FavoriteBorderIcon from "@material-ui/icons/FavoriteBorder";
import LocalLibraryIcon from "@material-ui/icons/LocalLibrary";
import RssFeedIcon from "@material-ui/icons/RssFeed";
import AddIcon from "@material-ui/icons/Add";
import { MEDIA_TYPES } from "../../../constants/categories";
import { DIFFICULTIES } from "../../../constants/categories";

import useStyles from "./styles";

function Sidebar({ darkMode }) {
  const classes = useStyles();
  return (
    <div>
      <Paper elevation={12}>
        <div className={classes.sidebarContainer}>
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
              <ListItem
                key={diff.id}
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
      </Paper>
    </div>
  );
}

export default Sidebar;
