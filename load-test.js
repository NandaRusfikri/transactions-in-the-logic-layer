import http from 'k6/http';
import { check, sleep } from 'k6';


export let options = {
    vus: 1, // Number of virtual users (VUs)
    duration: '1s', // Duration of the test
  };

// URL base API
const BASE_URL = 'http://localhost:9999/v1';
var token = 'isi jewete'




export function NewOrder() {


    const payload = JSON.stringify({
        customer_id: Math.floor(Math.random() * 500) + 10,
        items: [
            {
                product_id: Math.floor(Math.random() * 500) + 10,
                quantity: Math.floor(Math.random() * 20) + 1,
                note: "string note",
            },
            {
                product_id: Math.floor(Math.random() * 500) + 10,
                quantity: Math.floor(Math.random() * 20) + 1,
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
