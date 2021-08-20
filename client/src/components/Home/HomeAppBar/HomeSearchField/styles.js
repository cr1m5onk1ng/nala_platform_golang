import { makeStyles } from "@material-ui/core/styles";

export default makeStyles((theme) => ({
  searchBar: {
    backgroundColor: theme.palette.secondary,
    padding: "2px 4px",
    display: "flex",
    alignItems: "center",
    maxWidth: "460px",
  },
  input: {
    marginLeft: theme.spacing(1),
    flex: 1,
  },
  iconButton: {
    padding: 10,
  },
  divider: {
    height: 28,
    margin: 4,
  },
}));
