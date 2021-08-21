import { CircularProgress, Grid } from "@material-ui/core";
import React, { useState, useEffect } from "react";
import { useSelector } from "react-redux";
import useStyles from "./styles";
import { posts } from "../../constants/mockData";
import PostCard from "./Post/PostCard";

function Posts() {
  const classes = useStyles();
  //const { posts, isLoading } = useSelector((state) => state.posts);
  const isLoading = !posts.length;

  return (
    <Grid
      className={classes.container}
      container
      spacing={6}
      justifyContent="center"
    >
      {posts?.map((post) => (
        <Grid item key={post._id}>
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
        </Grid>
      ))}
    </Grid>
  );
}

export default Posts;
