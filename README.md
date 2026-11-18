# Please Rewind

Uma página que junta o que eu ando ouvindo e o que eu ando assistindo, em um lugar só.

Os dados vêm da Last.fm e do Letterboxd. A ideia é que qualquer pessoa possa colocar os dois nicks dela e ver a própria página, e também comparar com outra pessoa.

## Status

Nada implementado ainda. Por enquanto só existe o desenho da tela e a decisão de stack.

## Stack

- Go
- Vercel

## De onde vêm os dados

| Fonte | Como | Precisa de key |
|---|---|---|
| Last.fm | API oficial | sim |
| Letterboxd | RSS do perfil (`letterboxd.com/<usuario>/rss/`) | não |

O Letterboxd não tem API pública aberta, então o caminho é o RSS. Ele devolve os 50 itens mais recentes.

## URLs planejadas

```
victorwvm.com.br/rewind/<lastfm>/<letterboxd>
victorwvm.com.br/rewind/<lastfm>/<letterboxd>/vs/<lastfm2>/<letterboxd2>
```

## Janelas de tempo

30 dias, 90 dias, 365 dias e tudo. Batem com os `period` da Last.fm: `1month`, `3month`, `12month` e `overall`.
