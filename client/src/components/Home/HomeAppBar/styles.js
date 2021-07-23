import { makeStyles } from "@material-ui/core";
import { DRAWER_WIDTH } from "../../../constants/layout";

export default makeStyles((theme) => ({
  appBar: {
    boxShadow: "none",
    width: `calc(100% - ${DRAWER_WIDTH}px)`,
  },
  mAppBar: {
    boxShadow: "none",
    width: "100%",
  },
  grow: {
    flexGrow: 1,
  },
  icons: {
    paddingRight: theme.spacing(5),
  },
  menuIcon: {
    paddingRight: theme.spacing(5),
    paddingLeft: theme.spacing(6),
    marginLeft: theme.spacing(6),
    marginRight: theme.spacing(6),
  },
  addIcon: {
    borderRadius: 15,
  },
  logo: {
    height: 30,
    paddingLeft: 15,
    paddingBottom: 10,
  },
}));
