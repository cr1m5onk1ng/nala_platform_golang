import {
  Link,
  Avatar,
  Container,
  ImageList,
  ImageListItem,
  makeStyles,
  Typography,
  Divider,
  Badge,
  Grid,
  IconButton,
} from "@material-ui/core";
import { AvatarGroup } from "@material-ui/lab";
import Communities from "./Communities/Communities";
import Discussions from "./Discussions/Discussions";
import {
  communities,
  members,
  communitiesDiscussions,
} from "../../../constants/mockData";

const useStyles = makeStyles((theme) => ({
  container: {
    paddingTop: theme.spacing(10),
    position: "sticky",
    top: 0,
  },
  rightbarContainer: {
    paddingTop: theme.spacing(10),
    [theme.breakpoints.up("md")]: {
      paddingRight: theme.spacing(10),
    },
  },
  title: {
    fontWeight: 500,
    color: "#555",
  },
  link: {
    marginRight: theme.spacing(2),
    color: "#555",
    fontSize: 16,
  },
  membersContainer: {
    marginBottom: 20,
    paddingBottom: theme.spacing(5),
  },
  discussionsTitle: {
    [theme.breakpoints.down("md")]: {
      display: "none",
    },
  },
  badge: {
    margin: theme.spacing(3),
  },
}));

const Rightbar = () => {
  const classes = useStyles();

  const Members = ({ members }) => {
    return (
      <AvatarGroup
        className={classes.membersContainer}
        max={10}
        spacing="small"
      >
        {members.map((member) => (
          <Avatar alt={member.username} src={member.avatar} />
        ))}
      </AvatarGroup>
    );
  };

  return (
    <Container className={classes.container}>
      <div className={classes.rightbarContainer}>
        <Typography variant="h5" className={classes.title} gutterBottom>
          Followed Members
        </Typography>
        <Members members={members} />
        <Typography variant="h5" className={classes.title} gutterBottom>
          Communities
        </Typography>
        <Communities communities={communities} />
        <Typography
          className={classes.discussionsTitle}
          variant="h5"
          className={classes.title}
          gutterBottom
        >
          Top Discussions
        </Typography>
        <Discussions discussions={communitiesDiscussions} />
        <Typography variant="h5" className={classes.title} gutterBottom>
          Categories
        </Typography>
        <Link href="#" className={classes.link} variant="body2">
          Sport
        </Link>
        <Link href="#" className={classes.link} variant="body2">
          Food
        </Link>
        <Link href="#" className={classes.link} variant="body2">
          Music
        </Link>
        <Divider flexItem style={{ marginBottom: 5 }} />
        <Link href="#" className={classes.link} variant="body2">
          Movies
        </Link>
        <Link href="#" className={classes.link} variant="body2">
          Science
        </Link>
        <Link href="#" className={classes.link} variant="body2">
          Life
        </Link>
      </div>
    </Container>
  );
};

export default Rightbar;
