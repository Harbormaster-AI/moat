import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AppointmentService } from '../../../services/Appointment.service';
import { Appointment } from '../../../models/Appointment';
import { SubBaseComponent } from '../../Appointment/sub.base.component';

@Component({
    selector: 'app-create-appointment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAppointmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Appointment';

    appointmentForm: FormGroup;
    appointment: Appointment;

    constructor( http: HttpClient,
        private appointmentService: AppointmentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority): void {
        this.appointmentService
        .addAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexAppointment']);
            });
    }

    ngOnInit(): void {
    }
}