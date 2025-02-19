![Infobae_logo](https://user-images.githubusercontent.com/68082746/163064760-34ec0f9f-0ad1-4c06-81aa-235acaf0b99e.svg)

# InfobaeAPI

## 🌐 Description

- It includes various topics that can be accessed via the `topic` parameter.  
- The posts are just links, extracted from the site's `robots.txt` as I couldn't find an official API.
- The `xml/sitemap` endpoint returns all available sitemaps.

## 🚀 Usage 

- You can access the API -> [HERE](https://noticias.jonathan.com.ar/api/infobae/).  

## 🚧 Routes  

- **GET** `xml/sitemap` → Retrieve all sitemaps.  
- **GET** `api/infobae/` → Fetch the latest news from the general sitemap.  
- **GET** `api/infobae/topic/` → Fetch a specific topic.  
- **GET** `api/infobae/topic/?size=x` → Fetch a specific topic with `x` posts.

## ✍️ Examples

```shell
$ curl -Ss "https://noticias.jonathan.com.ar/api/infobae/economia?size=2" | jq   
{
  "news": [
    {
      "changefreq": "hourly",
      "lastmod": "2025-02-18T22:41:31.189Z",
      "url": "https://www.infobae.com/economia/2025/02/18/los-fundamentos-no-cambian-el-gobierno-se-repliega-sobre-su-plan-economico-para-evitar-un-impacto-en-los-mercados/"
    },
    {
      "changefreq": "hourly",
      "lastmod": "2025-02-18T22:35:05.232Z",
      "url": "https://www.infobae.com/economia/2025/02/18/la-bolsa-portena-subio-6-tras-asimilar-el-ruido-por-el-escandalo-de-libra/"
    }
  ]
}
```

## 🧰 Stack  

- Golang  
- Gin  
- GitHub Actions  
- Pre-commit  
- CodeQL  
- Docker  
- Kubernetes  
- Terraform  
- Bruno

## 📁 Folder structure

```md
🌳 src/
┣ 📁 constants/
┃ ┣ 📄 config.go
┃ ┣ 📄 links.go
┃ ┗ 📄 messages.go
┣ 📁 controllers/
┃ ┣ 📄 infobae.controller.go
┃ ┗ 📄 xml.controller.go
┣ 📁 models/
┃ ┗ 📄 response.go
┣ 📁 services/
┃ ┗ 📄 xml.service.go
┗ 📄 main.go
```

## 🤝 Contribute

For more information, check the [CONTRIBUTE](./CONTRIBUTE.md) file

## 📝 License 

[MIT](LICENSE)  
