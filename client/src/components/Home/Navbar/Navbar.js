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
import { Cancel, Mail, Notifications, Search, Add } from "@material-ui/icons";
import { useState } from "react";
import MuiAlert from "@material-ui/lab/Alert";
import Form from "./Form/Form";

const useStyles = makeStyles((theme) => ({
  toolbar: {
    display: "flex",
    justifyContent: "space-between",
    [theme.breakpoints.up("sm")]: {
      marginRight: theme.spacing(20),
    },
    [theme.breakpoints.up("md")]: {
      marginRight: theme.spacing(0),
    },
  },
  logoLg: {
    display: "none",
    [theme.breakpoints.up("sm")]: {
      display: "block",
    },
  },
  logoSm: {
    display: "block",
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
    color: "white",
    marginLeft: theme.spacing(1),
  },
  cancel: {
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
  },
  searchButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
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
  },
  badge: {
    marginRight: theme.spacing(3),
  },
  avatar: {
    marginRight: theme.spacing(3),
  },
}));

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const Navbar = () => {
  const [searchOpen, setSearchOpen] = useState(false); // search bar state
  const [formOpen, setFormOpen] = useState(false); // add post form state
  const [openAlert, setOpenAlert] = useState(false); // snackbar state
  const classes = useStyles({ searchOpen });

  // taken straight from material-ui
  const handleSnackbarClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }
    setOpenAlert(false);
  };

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
          <Search />
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
            <Add />
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
      <Form
        isOpen={formOpen}
        setOpen={setFormOpen}
        setOpenAlert={setOpenAlert}
      />
      <Snackbar
        open={openAlert}
        autoHideDuration={4000}
        onClose={handleSnackbarClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "left" }}
      >
        <Alert onClose={handleSnackbarClose} severity="success">
          Post addedd succefully
        </Alert>
      </Snackbar>
    </AppBar>
  );
};

export default Navbar;
