import React from "react";
import {
  makeStyles,
  Card,
  IconButton,
  Typography,
  Avatar,
  CardContent,
  CardHeader,
  CardActionArea,
} from "@material-ui/core";
import ReactTimeAgo from "react-time-ago";
import {
  ChatBubbleOutline,
  ShareOutlined,
  BookmarkBorder,
  MoreVert,
} from "@material-ui/icons";
import Community from "../Communities/Community";
import { communitiesDiscussions } from "../../../../constants/mockData";

const useStyles = makeStyles((theme) => ({
  cardContainer: {
    display: "flex",
    flexDirection: "column",
    borderRadius: 12,
    padding: theme.spacing(3),
    marginBottom: theme.spacing(5),
  },
  buttonsContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-start",
  },
  button: {
    margin: theme.spacing(1),
  },
  userRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-start",
    marginTop: theme.spacing(1),
    "&:hover": {
      background: theme.palette.hoverColor.main,
      cursor: "pointer",
    },
  },
  communityRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-start",
    marginTop: theme.spacing(1),
  },
  buttonRow: {
    display: "flex",
    alignItems: "center",
    margin: theme.spacing(1),
  },
  userRow__content: {
    margin: theme.spacing(2),
  },
  username: {
    fontSize: "16px",
  },
}));

const Discussion = ({
  community,
  username,
  creationTime,
  title,
  commentsCount,
  onClickCommunity,
  onClickUser,
}) => {
  const classes = useStyles();

  const communitySection = (title, thumbnail, onClick) => {
    return (
      <IconButton onClick={onClick}>
        <Avatar className={classes.img} alt={title} src={thumbnail} />
      </IconButton>
    );
  };

  const communityRow = (comm, onClick) => {
    {
      communitySection(comm.title, comm.thumbnail, onClick);
    }
  };

  const userRow = (username, onClickUser) => {
    return (
      <div className={classes.userRow}>
        <Typography
          className={`${classes.userRow__content} ${classes.username}`}
          variant="subtitle2"
          onClick={onClickUser}
        >
          {username}
        </Typography>
      </div>
    );
  };

  const buttonsRow = (commentsCount) => {
    return (
      <div className={classes.buttonsContainer}>
        <div className={classes.buttonRow}>
          <IconButton>
            <ChatBubbleOutline />
          </IconButton>
          <Typography variant="body2">{commentsCount} comments</Typography>
        </div>
        <IconButton className={classes.button}>
          <ShareOutlined />
        </IconButton>
        <IconButton className={classes.button}>
          <BookmarkBorder />
        </IconButton>
      </div>
    );
  };

  return (
    <Card className={classes.cardContainer} elevation={6}>
      <CardHeader
        avatar={communitySection(
          community.title,
          community.thumbnail,
          onClickCommunity
        )}
        action={
          <IconButton aria-label="settings">
            <MoreVert />
          </IconButton>
        }
        title={userRow(username, onClickUser)}
        subheader={<ReactTimeAgo date={creationTime} locale="it-IT" />}
      />
      <CardContent>
        <Typography
          variant="body1"
          component="p"
          color="textSecondary"
          className={classes.description}
        >
          {title}
        </Typography>
      </CardContent>
      {buttonsRow(commentsCount)}
    </Card>
  );
};

export default Discussion;
