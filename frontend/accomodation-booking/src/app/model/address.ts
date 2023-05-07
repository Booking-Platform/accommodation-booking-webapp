export class Address {
  country: string = '';
  city: string = '';
  street: string = '';
  number: string = '';

  constructor(country: string, city: string, street: string, number: string) {
    this.country = country;
    this.city = city;
    this.street = street;
    this.number = number;
  }
}
