![Infobae_logo](https://user-images.githubusercontent.com/68082746/163064760-34ec0f9f-0ad1-4c06-81aa-235acaf0b99e.svg)

# 🤖 Infobae-api

## API sencilla para consultar post de ciertos tipos

- Cuenta con varios temas que se llaman por una query de ```topic```
- Los post son solo los links ya que los saque del robots.txt de su pagina al no encontrar una api
- Por si alguien quiere contrubuir dejo el link de los otros sites -> [Aca](https://www.infobae.com/robots.txt)

## Puedo probar la API?

- Si! esta alojada en render -> [Aca](https://infobae-api.onrender.com/)

## 🚧 Rutas

- GET ```api/infobae/``` para consultar la ultima noticia en una categoria random
- GET ```api/infobae/?topic=<topic>``` para traer x topic expecifico
- GET ```api/infobae/?topic=<topic>&size=<size>``` para traer un topic expecifico con n cantidad de posts

## 🧰 Stack

- Typescript
- Nodejs
- Express
- Axios
- Render
