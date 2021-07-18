import { makeStyles } from "@material-ui/core";

export default makeStyles((theme) => ({
  appBar: {
    boxShadow: "none",
    zIndex: theme.zIndex.drawer + 1,
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
  },
  logo: {
    height: 30,
  },
}));
