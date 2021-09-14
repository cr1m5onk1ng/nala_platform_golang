import axios from "axios";
import { decode } from "paseto";

// MAIN API
const api = axios.create({ baseURL: "http://localhost:5000/nala/v1" });
const apiProtected = axios.create({ baseURL: "http://localhost:5000/nala/v1" });

// Verify Token Validity
const refreshToken = async (user) => {
  try {
    const { error, message, data } = await api.post("/users/tokens/paseto", {
      token: user.accessToken,
    });
    if (error) {
      console.log(
        "something went wrong with refresh token request: " + message
      );
      return null;
    }
    return data;
  } catch (err) {
    console.log(err);
  }
};

apiProtected.interceptors.request.use(
  async (config) => {
    let currentDate = new Date();
    const user = JSON.parse(localStorage.getItem("user"));
    console.log("User from interceptor: ", user);
    if (user) {
      const decodedToken = decode(user.accessToken);
      const expDate = new Date(decodedToken.payload.expired_at).getTime();
      console.log("User Token Expiration Date: ", expDate);
      if (expDate < currentDate.getTime()) {
        const data = await refreshToken();
        console.log("Data from refreshToken() call: ", data);
        if (data) {
          config.headers["authorization"] = "Bearer " + data.accessToken;
        }
      }
      return config;
    } else {
      return Promise.reject("User not authenticated");
    }
  },
  (error) => {
    return Promise.reject(error);
  }
);

// LOGIN
export const login = (data) => api.post("/users/login", data);
export const register = (data) => api.post("/users/register", data);
export const logout = (data) => api.post("/users/logout", data);

// POSTS
export const fetchPost = (id) => api.get(`/posts/${id}`);

export const fetchPostsByLanguage = (lang, page) =>
  api.get(`/posts/langs/${lang}?page=${page}`);

export const fetchPostsByLanguageAndMedia = (lang, media, page) =>
  api.get(`/posts/q?lang=${lang}&media=${media}&page=${page}`);

export const fetchPostsByLanguageAndCategory = (lang, cat, page) =>
  api.get(`/posts/q?lang=${lang}&cat=${cat}&page=${page}`);

export const fetchPostsByLanguageAndMediaAndCategory = (
  lang,
  media,
  cat,
  page
) => api.get(`/posts/q?lang=${lang}&media=${media}&cat=${cat}?&age=${page}`);

export const fetchPostsByUser = (user) => api.get(`/posts/users/${user}`);

export const fetchPostsBySearch = (query) =>
  api.get(`/posts/search?query=${query.search || "none"}&tags=${query.tags}`);

export const createPost = (post) => apiProtected.post("/posts", post);

export const likePost = (postId) => apiProtected.patch(`/posts/${postId}/like`);

export const updatePost = (id, updatedPost) =>
  apiProtected.patch(`/posts/${id}`, updatedPost);

export const deletePost = (id) => apiProtected.delete(`/posts/${id}`);

// DISCUSSIONS

export const addDiscussion = (discussionData) => {
  apiProtected.post(`/posts/discussions`, discussionData);
};

export const removeDiscussion = (discussionData) => {
  apiProtected.delete(`/posts/discussions`, discussionData);
};

export const updateDiscussion = (discussionData) => {
  apiProtected.put(`/posts/discussions`, discussionData);
};

export const getPostDiscussions = (postId) => {
  apiProtected.get(`/posts/discussions/${postId}`);
};

// COMMENTS

export const addComment = (commentData) => {
  apiProtected.post(`/posts/discussions/comments`, commentData);
};

export const removeComment = (commentData) => {
  apiProtected.delete(`/posts/discussions/comments`, commentData);
};

export const updateComment = (commentData) => {
  apiProtected.put(`/posts/discussions/comments`, commentData);
};

export const likeComment = (likeData) =>
  apiProtected.post(`/posts/discussions/comments/like`, likeData);

export const unlikeComment = (likeData) => {
  apiProtected.delete(`/posts/discussions/comments/like`, likeData);
};

export const getDiscussionComments = (discussionId) => {
  apiProtected.get(`/posts/discussions/comments/${discussionId}`);
};

export const getCommentsLikes = (commentId) => {
  apiProtected.get(`/posts/discussions/comments/likes/${commentId}`);
};

export const getCommentsLikesCount = (commentId) => {
  apiProtected.get(`/posts/discussions/comments/likes/count/${commentId}`);
};
