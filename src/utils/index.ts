const economia = "https://www.infobae.com/feeds/sitemap/sites/economia/";
const ciencia = "https://www.infobae.com/feeds/sitemap/sites/salud/ciencia/";
const fotos = "https://www.infobae.com/feeds/sitemap/sites/america/fotos/";
const mundo = "https://www.infobae.com/feeds/sitemap/sites/america/mundo/";
const politica = "https://www.infobae.com/feeds/sitemap/sites/politica/";
const init = "https://www.infobae.com/sitemap.xml";

function selectService(topic: string) {
  switch (topic) {
    case "economia":
      return economia;
    case "ciencia":
      return ciencia;
    case "fotos":
      return fotos;
    case "mundo":
      return mundo;
    case "politica":
      return politica;
    case undefined:
      return init;
    default:
      return init;
  }
}

export default selectService;
