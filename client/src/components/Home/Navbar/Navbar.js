import {
  alpha,
  AppBar,
  Avatar,
  Badge,
  InputBase,
  makeStyles,
  Toolbar,
  Typography,
  Button,
  Snackbar,
  IconButton,
} from "@material-ui/core";
import {
  Cancel,
  Notifications,
  Search,
  Add,
  AddOutlined,
} from "@material-ui/icons";
import { useState } from "react";
import MuiAlert from "@material-ui/lab/Alert";
import Form from "./Form/Form";

const useStyles = makeStyles((theme) => ({
  toolbar: {
    display: "flex",
    justifyContent: "space-between",

    [theme.breakpoints.up("md")]: {
      marginRight: theme.spacing(0),
    },
  },
  logoLg: {
    display: "none",
    marginLeft: theme.spacing(3),
    [theme.breakpoints.up("sm")]: {
      display: "block",
    },
  },
  logoSm: {
    display: "block",
    marginLeft: theme.spacing(3),
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
  },
  search: {
    display: "flex",
    alignItems: "center",
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    "&:hover": {
      backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    borderRadius: theme.shape.borderRadius,
    width: "50%",
    [theme.breakpoints.down("xs")]: {
      display: (props) => (props.searchOpen ? "flex" : "none"),
      width: "70%",
      marginRight: theme.spacing(5),
    },
  },
  input: {
    color: theme.palette.secondary.main,
    marginLeft: theme.spacing(1),
  },
  cancel: {
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
  },
  searchButton: {
    marginRight: theme.spacing(5),

    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
  },
  searchIcon: {
    margin: theme.spacing(2),
  },
  icons: {
    alignItems: "center",
    display: (props) => (props.searchOpen ? "none" : "flex"),
  },
  addButton: {
    borderRadius: 15,
    [theme.breakpoints.down("xs")]: {
      display: "none",
    },
    marginRight: theme.spacing(3),
  },
  addIconButton: {
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
    marginRight: theme.spacing(3),
    height: "30px",
    width: "30px",
    border: "1px solid",
    borderColor: theme.palette.secondary,
  },
  badge: {
    marginRight: theme.spacing(3),
  },
  avatar: {
    marginRight: theme.spacing(3),
  },
}));

const Navbar = () => {
  const [searchOpen, setSearchOpen] = useState(false); // search bar state
  const [formOpen, setFormOpen] = useState(false); // add post form state
  const classes = useStyles({ searchOpen });

  return (
    <AppBar position="fixed">
      <Toolbar className={classes.toolbar}>
        <Typography variant="h6" className={classes.logoLg}>
          NaLa
        </Typography>
        <Typography variant="h6" className={classes.logoSm}>
          NALA
        </Typography>
        <div className={classes.search}>
          <Search className={classes.searchIcon} />
          <InputBase placeholder="Search..." className={classes.input} />
          <Cancel
            className={classes.cancel}
            onClick={() => setSearchOpen(false)}
          />
        </div>
        <div className={classes.icons}>
          <Search
            className={classes.searchButton}
            onClick={() => setSearchOpen(true)}
          />
          <Button
            onClick={() => setFormOpen(true)}
            className={classes.addButton}
            startIcon={<Add />}
            variant="contained"
            color="secondary"
          >
            Add Post
          </Button>
          <IconButton
            className={classes.addIconButton}
            onClick={() => setFormOpen(true)}
            color="inherit"
            aria-label="add post"
          >
            <AddOutlined />
          </IconButton>
          <IconButton onClick={() => {}}>
            <Badge badgeContent={2} color="secondary" className={classes.badge}>
              <Notifications />
            </Badge>
          </IconButton>
          <IconButton onClick={() => {}}>
            <Avatar className={classes.avatar} alt="Satoshi Nakamoto" src="" />
          </IconButton>
        </div>
      </Toolbar>
      <Form isOpen={formOpen} setOpen={setFormOpen} />
    </AppBar>
  );
};

export default Navbar;
