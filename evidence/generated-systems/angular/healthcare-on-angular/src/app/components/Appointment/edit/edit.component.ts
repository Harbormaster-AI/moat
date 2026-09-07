import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AppointmentService } from '../../../services/Appointment.service';
import { SubBaseComponent } from '../../Appointment/sub.base.component';


@Component({
    selector: 'app-edit-appointment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAppointmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Appointment';

    appointmentForm: FormGroup;
    appointment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AppointmentService,
        private fb: FormBuilder
) {
        super(http);
        this.appointmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  appointmentDate: ['', Validators.required],
      reason: ['', Validators.required],
      Patient: ['', ],
      Clinician: ['', ],
      Facility: ['', ],
      Encounter: ['', ],
      Status: ['', ],
      Priority: ['', ]
        });
    }

    
    updateAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAppointment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAppointment(params['id']).subscribe(res => {
                this.appointment = res;
            });
        });
    }
}