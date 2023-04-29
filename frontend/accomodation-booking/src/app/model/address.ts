export class Address {
  id: string;
  country: string;
  city: string;
  street: string;
  number: string;

  constructor(
    id: string,
    country: string,
    city: string,
    street: string,
    number: string
  ) {
    this.id = id;
    this.country = country;
    this.city = city;
    this.street = street;
    this.number = number;
  }
}
