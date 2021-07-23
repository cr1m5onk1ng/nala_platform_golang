import React, { useState, useEffect } from "react";
import { TextField, Button, Typography, Paper } from "@material-ui/core";
import useStyles from "./styles";
import FileBase from "react-file-base64";
import { useDispatch, useSelector } from "react-redux";
import { createPost, updatePost } from "../../../../actions/posts";

import DifficultyCheckbox from "./DifficultyCheckbox";

function PostForm({ currentId, setCurrentId }) {
  const classes = useStyles();
  const dispatch = useDispatch();

  const [postData, setPostData] = useState({
    title: "",
    description: "",
    url: "",
    difficulty: "",
    tags: "",
    selectedFile: "",
  });

  const [checkBoxState, setCheckBoxState] = useState({
    easyChecked: false,
    mediumChecked: false,
    advancedChecked: false,
    nativeChecked: false,
  });

  const handleCheckBoxChange = (event) => {
    setCheckBoxState({
      ...checkBoxState,
      [event.target.name]: event.target.checked,
    });
  };

  const post = useSelector((state) =>
    currentId ? state.posts.find((message) => message._id === currentId) : null
  );

  useEffect(() => {
    if (post) setPostData(post);
  }, [post]);

  const onSubmit = async (e) => {
    e.preventDefault();

    if (currentId === 0) {
      dispatch(createPost(postData));
      clear();
    } else {
      dispatch(updatePost(currentId, postData));
      clear();
    }
  };

  const clear = () => {
    setCurrentId(0);
    setPostData({
      title: "",
      description: "",
      url: "",
      difficulty: "",
      media: "",
      tags: "",
      selectedFile: "",
    });
  };

  return (
    <Paper className={classes.paper}>
      <form
        autoComplete="off"
        noValidate
        className={`${classes.root} ${classes.form}`}
        onSubmit={onSubmit}
      >
        <Typography variant="h6">
          {currentId ? `Editing "${post.title}"` : "Uploading a Resource"}
        </Typography>
        <TextField
          name="url"
          variant="outlined"
          label="Url"
          fullWidth
          value={postData.creator}
          onChange={(e) =>
            setPostData({ ...postData, creator: e.target.value })
          }
        />
        <TextField
          name="title"
          variant="outlined"
          label="Title"
          fullWidth
          value={postData.title}
          onChange={(e) => setPostData({ ...postData, title: e.target.value })}
        />
        <TextField
          name="description"
          variant="outlined"
          label="Description"
          fullWidth
          multiline
          rows={4}
          value={postData.message}
          onChange={(e) =>
            setPostData({ ...postData, message: e.target.value })
          }
        />
        <TextField
          name="media"
          variant="outlined"
          label="Media type"
          fullWidth
          value={postData.tags}
          onChange={(e) =>
            setPostData({ ...postData, tags: e.target.value.split(",") })
          }
        />
        <TextField
          name="tags"
          variant="outlined"
          label="Tags"
          fullWidth
          value={postData.tags}
          onChange={(e) =>
            setPostData({ ...postData, tags: e.target.value.split(",") })
          }
        />

        <DifficultyCheckbox
          easyChecked={checkBoxState.easyChecked}
          mediumChecked={checkBoxState.mediumChecked}
          advancedChecked={checkBoxState.advancedChecked}
          nativeChecked={checkBoxState.nativeChecked}
          handleChange={handleCheckBoxChange}
        />
        <div className={classes.fileInput}>
          <FileBase
            type="file"
            multiple={false}
            onDone={({ base64 }) =>
              setPostData({ ...postData, selectedFile: base64 })
            }
          />
        </div>
        <Button
          className={classes.buttonSubmit}
          variant="contained"
          color="primary"
          size="large"
          type="submit"
          fullWidth
        >
          Submit
        </Button>
        <Button
          variant="contained"
          color="secondary"
          size="small"
          onClick={clear}
          fullWidth
        >
          Clear
        </Button>
      </form>
    </Paper>
  );
}

export default PostForm;
