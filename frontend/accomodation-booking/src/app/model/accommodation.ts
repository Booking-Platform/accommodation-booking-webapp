import { Address } from './address';
import { Appointment } from './appointment';

export class Accommodation {
  id: string = '';
  name: string = '';
  min_guest_num: string = '';
  max_guest_num: string = '';
  address!: Address;
  photos: string[] = [];
  benefits: string[] = [];
  automatic_confirmation: boolean = false;
  appointments!: Appointment;
}
