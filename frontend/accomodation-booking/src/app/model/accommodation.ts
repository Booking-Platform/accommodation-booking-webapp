import { Address } from './address';
import { Benefit } from './benefit';

export class Accommodation {
  id: string;
  name: string;
  minGuestNum: string;
  maxGuestNum: string;
  address: Address;
  photo: string[];
  benefits: Benefit[];

  constructor(
    id: string,
    name: string,
    minGuestNum: string,
    maxGuestNum: string,
    address: Address,
    photo: string[],
    benefits: Benefit[]
  ) {
    this.id = id;
    this.name = name;
    this.minGuestNum = minGuestNum;
    this.maxGuestNum = maxGuestNum;
    this.address = address;
    this.photo = photo;
    this.benefits = benefits;
  }
}
