import { Request, Response } from "express";
import infobaeService from "../services/infobaeService";
import boom from "@hapi/boom";

class InfobaeController {
  /**
   * @description Get a post from infobae
   * @param {size} req - Number of posts to return
   * @returns {Array} res - Array of objects with the values of infobae
   */

  async getPost(req: Request, res: Response) {
    try {
      const { size } = req.query;
      const data = await infobaeService.servicePosts();
      const allData = data.urlset.url.map((item: any) => {
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
