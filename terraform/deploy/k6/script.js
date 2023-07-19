import http from 'k6/http';

const url = 'http://a354dd963cf914b93a81d88888509b69-603239191.us-west-2.elb.amazonaws.com:8080';
import * as faker from 'https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js';
// import {faker} from "@faker-js/faker"

const menuPath = "/api/v1/menu/"
const transactionPath = "/api/v1/transaction/"

export default function () {
  let image = faker.image.food();
  let title = faker.lorem.sentence(3);
  let price = faker.random.number(3);
  let description = faker.lorem.sentence(3);
  let day = faker.random.word(1);
  // let image = "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png";
  // let title = "test";
  // let price = 1000;
  // let description = "test";
  // let day = "senin";
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
  // let transaction ={
  //   address: "string",
  //   amount: 0,
  //   count: 0,
  //   is_afternoon: true,
  //   is_morning: true,
  //   is_noon: true,
  //   menu_id: menu_id
  // }

  res = http.post(`${url}${transactionPath}`, JSON.stringify(transaction), {
    headers: { 'Content-Type': 'application/json' },
  });
  console.log(res.json());
  let transaction_id = res.json().data.id;
  let status = {
    id : transaction_id,
    status : "ongoing"
  }
  res = http.patch(`${url}${transactionPath}`, JSON.stringify(status), {
    headers: { 'Content-Type': 'application/json' },
  });
  console.log(res.json());


}
