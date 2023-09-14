import http from 'k6/http';

const url = 'http://app.default.svc.cluster.local:8080';
import * as faker from 'https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js';

const menuPath = "/api/v1/menu/"
const transactionPath = "/api/v1/transaction/"

export default function () {
  let image = faker.image.food();
  let title = faker.lorem.sentence(3);
  let price = faker.random.number(3);
  let description = faker.lorem.sentence(3);
  let day = faker.random.word(1);
  let menu = {
    day_menus: [
      {
        day: day,
        image: image,
      }
    ],
    description: description,
    price: price,
    title: title
  }

  // Using a JSON string as body
  let res = http.post(`${url}${menuPath}`, JSON.stringify(menu), {
    headers: { 'Content-Type': 'application/json' },
  });
  console.log(res.json());
  let menu_id = res.json().data.id;

  let transaction ={
    address: faker.random.word(1),
    amount: faker.random.number(3),
    count: faker.random.number(1),
    is_afternoon: true,
    is_morning: true,
    is_noon: true,
    menu_id: menu_id
  }

  res = http.post(`${url}${transactionPath}`, JSON.stringify(transaction), {
    headers: { 'Content-Type': 'application/json' },
  });
  let transaction_id = res.json().data.id;
  let status = {
    id : transaction_id,
    status : "ongoing"
  }
  res = http.patch(`${url}${transactionPath}`, JSON.stringify(status), {
    headers: { 'Content-Type': 'application/json' },
  });

}
