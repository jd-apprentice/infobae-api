import { describe, it, expect, beforeAll } from 'bun:test';
import { config } from '#config';

describe('config object', () => {

    beforeAll(() => {
        process.env.PORT = String(3000);
    });

    it('should have a port property', () => {
        expect(config).toHaveProperty('app');
    });

    it("should have a property port inside app object", () => {
        expect(config.app).toHaveProperty('port');
    });

    it("should have the value of port as 3000", () => {
        expect(config.app.port).toBe(String(3000));
    });

});