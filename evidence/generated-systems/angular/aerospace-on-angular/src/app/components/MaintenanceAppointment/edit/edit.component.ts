import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';
import { SubBaseComponent } from '../../MaintenanceAppointment/sub.base.component';


@Component({
    selector: 'app-edit-maintenanceAppointment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMaintenanceAppointmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MaintenanceAppointment';

    maintenanceAppointmentForm: FormGroup;
    maintenanceAppointment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MaintenanceAppointmentService,
        private fb: FormBuilder
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

    
    updateMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMaintenanceAppointment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMaintenanceAppointment(params['id']).subscribe(res => {
                this.maintenanceAppointment = res;
            });
        });
    }
}