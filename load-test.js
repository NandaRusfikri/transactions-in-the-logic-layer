import http from 'k6/http';
import { check, sleep } from 'k6';


export let options = {
    vus: 400, // Number of virtual users (VUs)
    duration: '180s', // Duration of the test
  };

// URL base API
const BASE_URL = 'http://localhost:9999/v1';
var token = 'isi jewete'




export function NewOrder() {


    const payload = JSON.stringify({
        customer_id: Math.floor(Math.random() * 800) + 10,
        idempotency_key: (Math.random() + 1).toString(36).substring(2, 2 + 20),
        items: [
            {
                product_id: Math.floor(Math.random() * 500) + 10,
                quantity: Math.floor(Math.random() * 2) + 1,
                note: "string note",
            },
            {
                product_id: Math.floor(Math.random() * (999 - 501 + 1)) + 501,
                quantity: Math.floor(Math.random() * 3) + 1,
                note: "string note",
            }
        ]
    });

    const headers = {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
    };

    const res = http.post(`${BASE_URL}/order`, payload, { headers: headers });

    check(res, {
        'NewOrder successfully': (res) => res.status === 200,
    });
    if (res.status !== 200) {
        return { error:res.json("message") }; 
    }

    // Kembalikan data jika berhasil
    return res.json("data");
}


// Main function yang menjalankan semua langkah
export default function () {

    NewOrder();

}
