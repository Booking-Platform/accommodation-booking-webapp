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
  {
    path: 'accommodation-details',
    component: ReservationDetailsComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Guest' },
  },
  {
    path: 'confirmReservations',
    component: ReservationConfirmationComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Host' },
  },
  {
    path: 'myReservations',
    component: MyReservationsComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Guest' },
  },
  {
    path: 'createAccommodation',
    component: CreateAccommodationComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Host' },
  },
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
  {
    path: 'addAppointment',
    component: AddAppointmentComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Host' },
  },
  {
    path: 'allAccommodations',
    component: AllAccommodationsComponent,
    canActivate: [AuthGuard],
    data: { requiredRole: 'Host' },
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
