import convert from "xml-js";
import axios from "axios";

class InfobaeService {
  public url: string = "https://www.infobae.com/feeds/sitemap/sites/economia/";

  /**
   * @description Get all posts from infobae
   * @returns {Array} Response - Json with all posts
   */
  async servicePosts() {
    try {
      const dataInfobae = await axios.get(this.url);
      const json = convert.xml2json(dataInfobae.data, {
        compact: true,
        ignoreComment: true,
        trim: true,
        ignoreDeclaration: true,
        ignoreInstruction: true,
        ignoreAttributes: true,
      });
      const jsonParsed = JSON.parse(json);
      return jsonParsed;
    } catch (e) {
      console.log(e);
    }
  }
}

export default new InfobaeService();
