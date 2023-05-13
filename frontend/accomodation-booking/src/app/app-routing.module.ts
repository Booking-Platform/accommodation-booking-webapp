import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ViewComponent } from './modules/view/view.component';
import { ReservationDetailsComponent } from './modules/reservation-details/reservation-details.component';
import { ReservationConfirmationComponent } from './modules/reservation-confirmation/reservation-confirmation.component';
import { MyReservationsComponent } from './modules/my-reservations/my-reservations.component';
import { CreateAccommodationComponent } from './modules/create-accommodation/create-accommodation.component';
import { RegisterComponent } from './modules/register/register.component';
import { LoginComponent } from './modules/login/login.component';
import { AuthGuard } from './guard/auth.guard';
import { AddAppointmentComponent } from './modules/add-appointment/add-appointment.component';
import { AllAccommodationsComponent } from './modules/all-accommodations/all-accommodations.component';

const routes: Routes = [
  { path: 'view', component: ViewComponent },
  { path: 'accommodation-details', component: ReservationDetailsComponent },
  {
    path: 'confirmReservations',
    component: ReservationConfirmationComponent,
    canActivate: [AuthGuard],
  },
  {
    path: 'myReservations',
    component: MyReservationsComponent,
    canActivate: [AuthGuard],
  },
  {
    path: 'createAccommodation',
    component: CreateAccommodationComponent,
    canActivate: [AuthGuard],
  },
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
  { path: 'addAppointment', component: AddAppointmentComponent, canActivate: [AuthGuard], },
  { path: 'allAccommodations', component: AllAccommodationsComponent },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
