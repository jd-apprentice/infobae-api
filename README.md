![Infobae_logo](https://user-images.githubusercontent.com/68082746/163064760-34ec0f9f-0ad1-4c06-81aa-235acaf0b99e.svg)

# 🤖 Infobae-api

## API sencilla para consultar post de ciertos tipos

- De momento solo hice el post para economia, si alguna vez necesito los otros los voy a hacer
- Los post son solo los links ya que los saque del robots.txt de su pagina al no encontrar una api
- Por si alguien quiere contrubuir dejo el link de los otros sites -> [Aca](https://www.infobae.com/robots.txt)

## Puedo probar la API?

- Si! esta alojada en heroku -> [Aca](https://infobae-api.herokuapp.com/api/infobae/economia)

## 🚧 Rutas

- GET ```api/infobae/economia``` para consultar la ultima noticia en economia
- GET ```api/infobae/economia?size=3``` para traer n cantidad de noticias, contando desde la ultima

## 🧰 Stack

- Typescript
- Nodejs
- Express
- Axios
