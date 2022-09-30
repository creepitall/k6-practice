import http from 'k6/http';
import {SharedArray} from 'k6/data';
import {check} from "k6";

export let options = {
    discardResponseBodies: true,
    scenarios: {
        query_scenario: {
            executor: 'ramping-arrival-rate',
            startRate: 0,
            timeUnit: '1s',
            preAllocatedVUs: 150,
            maxVUs: 1000,
            stages: [
                {
                    target: 1000, // 100 рпс
                    duration: '3m' // в течении 1 минут
                },
            ]
        }
},
}

const data = new SharedArray('data', function () {
    const f = JSON.parse(open('./ammo.json'));
    return f; // должен быть массив на выходе
});

export default function () {
    const element = data[Math.floor(Math.random() * data.length)];

    const url = '';

    const params = {
        headers: {
            'Content-Type': 'application/json',
            'x-aer-test-type': 'load',
            'x-aer-app-name': 'k6-loadtest'
        },
    };

    let resp = http.post(url, JSON.stringify(element), params);

    check(resp, {
        'is status 200': (t) => t.status === 200
    })
}