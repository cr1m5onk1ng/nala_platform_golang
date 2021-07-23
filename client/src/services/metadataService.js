import { parser } from "html-metadata-parser";

export const getMetadataFromUrl = async (url) => {
  try {
    const metadata = await parser(url);
    return {
      title: metadata.meta.title,
      description: metadata.meta.description,
      image: metadata.og.image,
    };
  } catch (error) {
    console.log(error);
  }
};
