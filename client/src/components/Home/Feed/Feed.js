import { Container, makeStyles } from "@material-ui/core";
import PostCard from "./Post/PostCard";
import { posts } from "../../../constants/mockData";

const useStyles = makeStyles((theme) => ({
  rootContainer: {
    paddingTop: theme.spacing(10),
    [theme.breakpoints.down("xs")]: {
      paddingTop: theme.spacing(10),
      paddingLeft: theme.spacing(0),
      paddingRight: theme.spacing(0),
    },
  },
  postContainer: {
    paddingTop: theme.spacing(10),
    [theme.breakpoints.up("md")]: {
      marginLeft: "15%",
      marginRight: "15%",
    },
    [theme.breakpoints.up("lg")]: {
      marginLeft: "25%",
      marginRight: "25%",
    },
    [theme.breakpoints.down("xs")]: {
      paddingTop: theme.spacing(10),
      paddingLeft: theme.spacing(0),
      paddingRight: theme.spacing(0),
    },
  },
}));

const Feed = () => {
  const classes = useStyles();
  return (
    <Container className={classes.rootContainer}>
      <div className={classes.postContainer}>
        {posts?.map((post) => (
          <PostCard
            title={post.title}
            mediaType={post.mediaType}
            description={post.description}
            url={post.url}
            creationTime={post.creationTime}
            imageUrl={post.imageUrl}
            username={post.username}
            likesCount={post.likesCount}
            easyCount={post.easyCount}
            mediumCount={post.mediumCount}
            advancedCount={post.advancedCount}
            nativeCount={post.nativeCount}
          />
        ))}
      </div>
    </Container>
  );
};

export default Feed;
