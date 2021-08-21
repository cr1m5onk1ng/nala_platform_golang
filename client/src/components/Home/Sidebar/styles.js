import { makeStyles } from "@material-ui/core";
import { blue, green, red } from "@material-ui/core/colors";

export default makeStyles((theme) => ({
  sidebarContainer: {
    width: 240,
    top: "10px",
    zIndex: "10",
    position: "fixed",
    marginTop: "7vh",
    marginLeft: "6vw",
    backgroundColor: theme.palette.background.paper,
    border: "1px solid lightgray",
    borderRadius: "12px",
  },
  sidebarContainerDark: {
    width: 240,
    position: "fixed",
    marginTop: "7vh",
    backgroundColor: "white",
    border: "1px solid lightgray",
    borderRadius: "10px",
    backgroundColor: "black",
  },
  drawerPaper: {
    width: 240,
  },
  listItemText: {
    fontSize: 14,
  },
  listItem: {
    paddingTop: 4,
    paddingBottom: 4,
  },
  listItemFeed: {
    color: blue[400],
  },
  listItemStudy: {
    color: green[400],
  },
  listItemFavorites: {
    color: red[400],
  },
  subheader: {
    textTransform: "uppercase",
  },
}));
