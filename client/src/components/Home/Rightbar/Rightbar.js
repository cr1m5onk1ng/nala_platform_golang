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
import { communities, members } from "../../../constants/mockData";

const useStyles = makeStyles((theme) => ({
  container: {
    paddingTop: theme.spacing(10),
    position: "sticky",
    top: 0,
  },
  rightbarContainer: {
    paddingTop: theme.spacing(10),
    [theme.breakpoints.up("md")]: {
      paddingRight: theme.spacing(50),
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
  communitiesContainer: {
    paddingBottom: theme.spacing(5),
  },
  badge: {
    margin: theme.spacing(3),
  },
  img: {
    width: theme.spacing(10),
    height: theme.spacing(10),
    border: "1px solid lightgray",
  },
}));

const Rightbar = () => {
  const classes = useStyles();

  const Community = ({ title, thumbnail, members, onClick }) => {
    return (
      <IconButton onClick={onClick}>
        <Badge
          className={classes.badge}
          badgeContent={members}
          max={10000}
          color="secondary"
        >
          <Avatar className={classes.img} alt={title} src={thumbnail} />
        </Badge>
      </IconButton>
    );
  };

  const Communities = ({ communities }) => {
    return (
      <Grid container className={classes.communitiesContainer}>
        {communities.map((com) => (
          <Grid item xs={12} sm={6} md={4}>
            <Community
              title={com.title}
              thumbnail={com.thumbnail}
              members={com.members}
            />
          </Grid>
        ))}
      </Grid>
    );
  };

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
