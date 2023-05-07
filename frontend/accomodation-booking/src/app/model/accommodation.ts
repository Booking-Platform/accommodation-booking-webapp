import { Address } from './address';

export class Accommodation {
  id: string = '';
  name: string = '';
  // min_guest_num: string = '';
  // max_guest_num: string = '';
  minGuestNum: string = '';
  maxGuestNum: string = '';
  address: Address;
  photos: string[] = [];
  benefits: string[] = [];
  automatic_confirmation: boolean = false;

  constructor(
    id: string,
    name: string,
    // min_guest_num: string,
    // max_guest_num: string,
    minGuestNum: string,
    maxGuestNum: string,
    address: Address,
    photos: string[],
    benefits: string[]
  ) {
    this.id = id;
    this.name = name;
    // this.min_guest_num = min_guest_num;
    // this.max_guest_num = max_guest_num;
    this.minGuestNum = minGuestNum;
    this.maxGuestNum = maxGuestNum;
    this.address = address;
    this.photos = photos;
    this.benefits = benefits;
  }
}
