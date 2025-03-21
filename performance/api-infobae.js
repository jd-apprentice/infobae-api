// Usage: k6 cloud run performance/api-infobae.js  --env TARGET_URL=<target-url>
// https://grafana.com/docs/k6/latest/using-k6/environment-variables/
// https://grafana.com/docs/k6/latest/get-started/results-output/

import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: __ENV.VUS || 5,
    duration: __ENV.DURATION || '60s',
    thresholds: {
        http_req_duration: ['p(99)<2000'], // 99% of requests must complete below 2 seconds
        http_req_failed: ['rate<0.1'], // http errors should be less than 10%
    },
};

/**
 * @usage k6 cloud run performance/api-infobae.js --env TARGET_URL=<target-url>
 * @environment VUS (default: 5)
 * @environment TARGET_URL (default: ...)
 * @environment DURATION (default: 60s)
 * @description K6 performance test
 */
export default function () {
    const url = __ENV.TARGET_URL || "https://news.jonathan.com.ar/api/infobae";

    // {"changefreq":"string","lastmod":"string","message":"string","url":"https://www.infobae.com/..."}

    let response = http.get(url);

    check(response, {
        'is status 200': (r) => r.status === 200,
    });

    sleep(1);
};