import { ObjectUrls } from "@types";

const URL: ObjectUrls = {
  autos: "https://www.infobae.com/feeds/sitemap/sites/autos/",
  economia: "https://www.infobae.com/feeds/sitemap/sites/economia/",
  horoscopo: "https://www.infobae.com/feeds/sitemap/sites/horoscopo/",
  sociedad: "https://www.infobae.com/feeds/sitemap/sites/sociedad/",
  salud: "https://www.infobae.com/feeds/sitemap/sites/salud/",
  ciencia: "https://www.infobae.com/feeds/sitemap/sites/salud/ciencia/",
  fotos: "https://www.infobae.com/feeds/sitemap/sites/america/fotos/",
  mundo: "https://www.infobae.com/feeds/sitemap/sites/america/mundo/",
  politica: "https://www.infobae.com/feeds/sitemap/sites/politica/",
  init: "https://www.infobae.com/sitemap.xml",
};

export const selectService = (topic: string): string => URL[topic] || URL.init;
