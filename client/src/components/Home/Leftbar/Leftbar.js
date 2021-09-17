import { Container, Divider, makeStyles, Typography } from "@material-ui/core";
import {
  leftbarMediaIcons,
  leftbarFooterIcons,
  leftbarNavIcons,
} from "../../../constants/mockData";

const useStyles = makeStyles((theme) => ({
  container: {
    height: "100vh",
    color: "white",
    paddingTop: theme.spacing(10),
    backgroundColor: "white",
    position: "sticky",
    top: 0,
    [theme.breakpoints.up("sm")]: {
      backgroundColor: "white",
      color: "#555",
      border: "1px solid #ece7e7",
    },
    [theme.breakpoints.down("xs")]: {
      width: "15vw",
    },
  },
  sidebarContainer: {
    paddingTop: theme.spacing(10),
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    [theme.breakpoints.down("xs")]: {
      alignItems: "center",
    },
  },
  item: {
    display: "flex",
    alignItems: "center",
    marginBottom: theme.spacing(4),
    padding: theme.spacing(2),
    "&:hover": {
      background: theme.palette.hoverColor.main,
    },
    [theme.breakpoints.up("sm")]: {
      marginBottom: theme.spacing(3),
      cursor: "pointer",
    },
  },
  icon: {
    marginRight: theme.spacing(1),
    marginLeft: theme.spacing(5),
    width: "30px",
    height: "30px",
    [theme.breakpoints.up("sm")]: {
      fontSize: "18px",
    },
    [theme.breakpoints.down("sm")]: {
      marginLeft: theme.spacing(0),
      color: theme.palette.secondary.dark,
    },
  },
  text: {
    fontWeight: 500,
    [theme.breakpoints.down("xs")]: {
      display: "none",
    },
  },
  divider: {
    marginBottom: theme.spacing(5),
  },
}));

const Leftbar = () => {
  const classes = useStyles();

  const LeftbarIcons = ({ navIcons, mediaIcons, footerIcons }) => {
    return (
      <Container className={classes.container}>
        <div className={classes.sidebarContainer}>
          {navIcons.map((icon) => (
            <div className={classes.item}>
              {<icon.icon className={classes.icon} />}
              <Typography className={classes.text}>{icon.title}</Typography>
            </div>
          ))}
          <Divider light className={classes.divider} />
          {mediaIcons.map((icon) => (
            <div className={classes.item}>
              {
                <icon.icon
                  className={classes.icon}
                  style={{ color: icon.color }}
                />
              }
              <Typography className={classes.text}>{icon.title}</Typography>
            </div>
          ))}
          <Divider className={classes.divider} />
          {footerIcons.map((icon) => (
            <div className={classes.item}>
              {<icon.icon className={classes.icon} />}
              <Typography className={classes.text}>{icon.title}</Typography>
            </div>
          ))}
        </div>
      </Container>
    );
  };
  return (
    <LeftbarIcons
      navIcons={leftbarNavIcons}
      mediaIcons={leftbarMediaIcons}
      footerIcons={leftbarFooterIcons}
    />
  );
};

export default Leftbar;
