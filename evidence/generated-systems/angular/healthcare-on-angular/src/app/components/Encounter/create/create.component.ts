import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EncounterService } from '../../../services/Encounter.service';
import { Encounter } from '../../../models/Encounter';
import { SubBaseComponent } from '../../Encounter/sub.base.component';

@Component({
    selector: 'app-create-encounter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEncounterComponent extends SubBaseComponent implements OnInit {

    title = 'Add Encounter';

    encounterForm: FormGroup;
    encounter: Encounter;

    constructor( http: HttpClient,
        private encounterService: EncounterService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.encounterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  encounterNumber: ['', Validators.required],
      startDateTime: ['', Validators.required],
      endDateTime: ['', Validators.required],
      Patient: ['', ],
      Clinician: ['', ],
      Facility: ['', ],
      Appointment: ['', ],
      Diagnoses: ['', ],
      Procedures: ['', ],
      Observations: ['', ],
      Orders: ['', ],
      Admission: ['', ],
      Discharge: ['', ],
      Status: ['', ],
      EncounterType: ['', ]
        });
    }

    
    addEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType): void {
        this.encounterService
        .addEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType)
            .subscribe(() => {
                this.router.navigate(['/indexEncounter']);
            });
    }

    ngOnInit(): void {
    }
}