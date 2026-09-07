import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';
import { MaintenanceAppointment } from '../../../models/MaintenanceAppointment';
import { SubBaseComponent } from '../../MaintenanceAppointment/sub.base.component';

@Component({
    selector: 'app-create-maintenanceAppointment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMaintenanceAppointmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add MaintenanceAppointment';

    maintenanceAppointmentForm: FormGroup;
    maintenanceAppointment: MaintenanceAppointment;

    constructor( http: HttpClient,
        private maintenanceAppointmentService: MaintenanceAppointmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.maintenanceAppointmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  appointmentDate: ['', Validators.required],
      Aircraft: ['', ],
      MroFacility: ['', ],
      WorkOrder: ['', ],
      Status: ['', ]
        });
    }

    
    addMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status): void {
        this.maintenanceAppointmentService
        .addMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMaintenanceAppointment']);
            });
    }

    ngOnInit(): void {
    }
}