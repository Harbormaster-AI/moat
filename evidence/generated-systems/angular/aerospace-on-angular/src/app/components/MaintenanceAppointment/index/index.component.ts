
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';
import { MaintenanceAppointment } from '../../../models/MaintenanceAppointment';

@Component({
    selector: 'app-index-maintenanceAppointment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMaintenanceAppointmentComponent implements OnInit {

    maintenanceAppointments: MaintenanceAppointment[] = [];

    constructor(
        private router: Router,
        private service: MaintenanceAppointmentService
) {}

    ngOnInit(): void {
        this.getMaintenanceAppointments();
}

    getMaintenanceAppointments(): void {
        this.service.getMaintenanceAppointments().subscribe((res) => {
        this.maintenanceAppointments = res;
    });
}

    deleteMaintenanceAppointment(id: any): void {
        this.service.deleteMaintenanceAppointment(id)
            .subscribe(() => {
                this.getMaintenanceAppointments();
            });
    }
}