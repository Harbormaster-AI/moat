
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppointmentService } from '../../../services/Appointment.service';
import { Appointment } from '../../../models/Appointment';

@Component({
    selector: 'app-index-appointment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAppointmentComponent implements OnInit {

    appointments: Appointment[] = [];

    constructor(
        private router: Router,
        private service: AppointmentService
) {}

    ngOnInit(): void {
        this.getAppointments();
}

    getAppointments(): void {
        this.service.getAppointments().subscribe((res) => {
        this.appointments = res;
    });
}

    deleteAppointment(id: any): void {
        this.service.deleteAppointment(id)
            .subscribe(() => {
                this.getAppointments();
            });
    }
}