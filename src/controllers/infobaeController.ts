import { Request, Response } from "express";
import infobaeService from "../services/infobaeService";
import boom from "@hapi/boom";
import { DataParse, ParsedQs, SizeProps } from "../models/types";
class InfobaeController {
  /**
   * @description Get a post from infobae
   * @param {size} req - Number of posts to return
   * @returns {DataParse} res - Data of the post
   */

  async getPosts(
    req: Request<ParsedQs, SizeProps>,
    res: Response
  ): Promise<Response> {
    try {
      const { topic }: ParsedQs = req.query;
      const { size } = req.query as unknown as SizeProps;
      const data = await infobaeService.servicePosts(topic?.toString()!);
      const allData = data.urlset.url.map((item: DataParse) => {
        return {
          lastmod: item.lastmod._text,
          link: item.loc._text,
        };
      });
      const dataParsed = allData.slice(0, size);
      return res.json(size === undefined || null ? allData[0] : dataParsed);
    } catch (error: any) {
      return res.json(boom.badImplementation(error));
    }
  }
}

export default new InfobaeController();
