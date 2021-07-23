import React from "react";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import Chip from "@material-ui/core/Chip";
import useStyles from "./styles";
import ShareIcon from "@material-ui/icons/Share";
import ReactTimeAgo from "react-time-ago";
import LinkIcon from "@material-ui/icons/Link";
import CardActionArea from "@material-ui/core/CardActionArea";

import { Avatar, CardHeader } from "@material-ui/core";
import { DIFFICULTIES } from "../../../constants/categories";

export default function PostCard({
  title,
  mediaType,
  description,
  url,
  creationTime,
  imageUrl,
  username,
  likesCount,
  easyCount,
  mediumCount,
  advancedCount,
  nativeCount,
}) {
  const classes = useStyles();

  const difficultiesRow = (difficulties) => {
    return (
      <div className={classes.difficultiesContainer}>
        <Chip
          className={classes.green}
          variant="outlined"
          className={classes.tagContainer}
          label={"Easy"}
          avatar={<Avatar className={classes.green}>{easyCount}</Avatar>}
        />
        <Chip
          variant="outlined"
          className={classes.tagContainer}
          label={"Medium"}
          avatar={<Avatar className={classes.blue}>{mediumCount}</Avatar>}
        />
        <Chip
          variant="outlined"
          className={classes.tagContainer}
          label={"Advanced"}
          avatar={<Avatar className={classes.pink}>{advancedCount}</Avatar>}
        />
        <Chip
          variant="outlined"
          className={classes.tagContainer}
          label={"Native"}
          avatar={<Avatar className={classes.red}>{nativeCount}</Avatar>}
        />
      </div>
    );
  };

  const tag = (text, count, customClass) => {
    return (
      <div className={`${classes.customTag} ${customClass}`}>
        <Typography className={classes.customTag__text}>{text}</Typography>
        <div className={classes.countBadge}>
          <Typography className={classes.countValue}>{count}</Typography>
        </div>
      </div>
    );
  };

  const difficultiesRowCustom = () => {
    return (
      <div className={classes.customDifficultiesRow}>
        {tag("Easy", 16, classes.green)}
        {tag("Medium", 4, classes.blue)}
        {tag("Advanced", 1, classes.pink)}
        {tag("Native", 0, classes.red)}
      </div>
    );
  };

  const buttonsRow = () => {
    return (
      <div className={classes.buttonsContainer}>
        <IconButton className={classes.button}>
          <LinkIcon />
        </IconButton>
      </div>
    );
  };

  const userRow = (username) => {
    return (
      <CardActionArea>
        <div className={classes.userRow}>
          <Typography
            className={`${classes.userRow__content} ${classes.username}`}
            variant="subtitle2"
          >
            {username}
          </Typography>
          <Typography className={classes.userRow__content} variant="subtitle2">
            {" - "}
          </Typography>
          <ReactTimeAgo date={creationTime} locale="it-IT" />
        </div>
      </CardActionArea>
    );
  };

  return (
    <Card className={classes.cardContainer} elevation={6}>
      {userRow(username)}
      <CardHeader
        action={
          <IconButton>
            <ShareIcon />
          </IconButton>
        }
        title={title}
        subheader={mediaType}
      />
      <CardMedia
        className={classes.image}
        component="img"
        alt="Contemplative Reptile"
        height="140"
        image={imageUrl || "https://source.unsplash.com/random"}
        title={title}
      />
      <CardContent>
        <Typography variant="body1" color="textSecondary">
          {description}
        </Typography>
      </CardContent>
      {difficultiesRowCustom()}
      {buttonsRow()}
    </Card>
  );
}
