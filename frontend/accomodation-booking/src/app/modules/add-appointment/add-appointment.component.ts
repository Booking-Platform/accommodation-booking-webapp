import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { Appointment } from 'src/app/model/appointment';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';

@Component({
  selector: 'app-add-appointment',
  templateUrl: './add-appointment.component.html',
  styleUrls: ['./add-appointment.component.css'],
})
export class AddAppointmentComponent implements OnInit {
  public appointment: Appointment = new Appointment();
  public accommodation: Accommodation = new Accommodation();
  public accommodationID: string = '';

  constructor(
    private route: ActivatedRoute,
    private accommodationService: AccommodationService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.accommodationID = params['accommodationID'];
    });
  }

  addAppointment(): void {
    var newAppointment = {
      id: this.accommodationID,
      from: this.appointment.from,
      to: this.appointment.to,
      price: this.appointment.price,
      per_guest: this.appointment.per_guest.toString(),
      status: 0,
    };

    this.accommodationService
      .addAppointment(newAppointment)
      .subscribe(() => this.router.navigate(['/allAccommodations']));
  }
}
