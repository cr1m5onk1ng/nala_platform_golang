import { makeStyles } from "@material-ui/core";
import { DRAWER_WIDTH } from "../../../constants/layout";

export default makeStyles((theme) => ({
  appBar: {
    boxShadow: "none",
  },
  toolbar: {
    display: "flex",
    alignItems: "center",
    justifyItems: "center",
    paddingLeft: "10vw",
    paddingRight: "10vw",
  },
  toolbarSmall: {
    display: "flex",
    alignItems: "center",
    justifyItems: "center",
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
