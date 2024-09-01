import { Request, Response } from "express";
import boom from "@hapi/boom";
import os from "os";

import { DataParse, postQuery } from "@types";
import infobaeService from "../services/infobaeService";

class InfobaeController {
  /**
   * @description Get a post from infobae
   * @param {size} req - Number of posts to return
   * @returns {DataParse} res - Data of the post
   */
  async getPosts(
    req: Request<postQuery>,
    res: Response
  ): Promise<Response> {
    try {
      const hostname = os.hostname();
      const { size, topic } = req.query as unknown as postQuery;
      const data = await infobaeService.servicePosts(topic);
      const allData = data.urlset.url.map((item: DataParse) => {
        return {
          lastmod: item.lastmod._text,
          link: item.loc._text,
          hostname
        };
      });
      const dataParsed = allData.slice(0, size);
      const dataFirst = allData[0];
      return res.json(size === undefined || null ? dataFirst : dataParsed);
    } catch (error: any) {
      return res.json(boom.badImplementation(error));
    }
  }
}

export default new InfobaeController();
