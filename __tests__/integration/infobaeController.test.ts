import { app } from "#server";
import { describe, it, expect, beforeAll } from 'bun:test';
import request from 'supertest';
import { type Response } from 'supertest';

const baseRoute = '/api/infobae';
const contentTypeKey = 'Content-Type';
const contentTypeValue = /json/;
const httpSuccess = 200;

const newMock = {
    lastmod: expect.any(String),
    link: expect.any(String),
} as const;

type responseNew = {
    body: {
        [key in keyof typeof newMock]: string
    }[];
}

describe('Infobae Module', () => {

    let server: Express.Application;

    beforeAll(async () => {
        server = app;
    })

    it('GET /api/infobae - should return a valid response', async () => {
        await request(app)
            .get(baseRoute)
            .expect(contentTypeKey, contentTypeValue)
            .expect(httpSuccess)
            .then((response) => expectNew(response));
    });

    it('GET /api/infobae - should return 3 posts', async () => {
        await request(app)
            .get(baseRoute)
            .query({ size: 3 })
            .expect(contentTypeKey, contentTypeValue)
            .expect(httpSuccess)
            .then((response: responseNew) => {
                const getCurrentYear = new Date().getFullYear().toString();
                for (const item of response.body) {
                    expect(item.lastmod).toContain(getCurrentYear);
                    expect(item.link).toMatch(/infobae.com/);
                    expect(item).toBeTruthy();
                    expect(item).toMatchObject(newMock);
                }
            });
    });

})

function expectNew(response: Response) {
    expect(response.body).toBeTruthy();
    expect(response.body).toMatchObject(newMock);
}