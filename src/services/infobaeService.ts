import convert from "xml-js";
import axios from "axios";
import selectService from "../utils/index";
import { JsonParsedProps } from "src/models/types";

class InfobaeService {
  /**
   * @description Get all posts from infobae
   * @returns {Array} Response - Json with all posts
   */
  async servicePosts(topic: string): Promise<JsonParsedProps> {
    try {
      const dataInfobae = await axios.get(selectService(topic)!);
      const json = convert.xml2json(dataInfobae.data, {
        compact: true,
        ignoreComment: true,
        trim: true,
        ignoreDeclaration: true,
        ignoreInstruction: true,
        ignoreAttributes: true,
      });
      const jsonParsed: JsonParsedProps = JSON.parse(json);
      return jsonParsed;
    } catch (error: any) {
      throw new Error(error);
    }
  }
}

export default new InfobaeService();
