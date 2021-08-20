import { makeStyles } from "@material-ui/core";
import { blue, green, red } from "@material-ui/core/colors";

export default makeStyles((theme) => ({
  drawer: {
    width: 240,
    flexShrink: 1,
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
  visible: {
    width: 240,
    flexShrink: 0,
    zIndex: 999,
  },
}));
