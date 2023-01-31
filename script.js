import http from 'k6/http';
import {SharedArray} from 'k6/data';
import {check} from "k6";

let env = JSON.parse(open('./env.json'));

export let options = {
    discardResponseBodies: true,
    scenarios: {
        // scn_ramping: {
        //     executor: 'ramping-arrival-rate',
        //     startRate: 0,
        //     timeUnit: '1s',
        //     preAllocatedVUs: 150,
        //     maxVUs: 2000,
        //     stages: [
        //         {
        //             target: 30000, // 100 рпс
        //             duration: '5m' // в течении 1 минут
        //         },
        //     ]
        // },
        scn_constant: {
            executor: 'constant-arrival-rate',
            timeUnit: '1s',
            preAllocatedVUs: 200,
            maxVUs: 2000,
            duration: '10m',
            rate: 30000, // 30к это около 20к реального // 20к-18к
        }
},
}

const data = new SharedArray('data', function () {
    const f = JSON.parse(open('./ammo.json'));
    return f; // должен быть массив на выходе
});

export default function () {
    const element = data[Math.floor(Math.random() * data.length)];

    let urls = env.urls;

    const url = urls[Math.floor(Math.random() * urls.length)];

    let h = {}
    for (let hdrs of env.headers) {
        h[hdrs.key] = hdrs.value
    }

    const params = {
        headers: h,
    //    timeout: '1s',
    };

    let resp = http.post(url, JSON.stringify(element), params);

    check(resp, {
        'http timeout': (t) => params.timeout,
        'is status 200': (t) => t.status === 200,
        'is status 404': (t) => t.status === 404,
        'is status 504': (t) => t.status === 504,
        'is error ': (t) => t.error.length
    })
}