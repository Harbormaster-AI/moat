import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EncounterService } from '../../../services/Encounter.service';
import { SubBaseComponent } from '../../Encounter/sub.base.component';


@Component({
    selector: 'app-edit-encounter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEncounterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Encounter';

    encounterForm: FormGroup;
    encounter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EncounterService,
        private fb: FormBuilder
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

    
    updateEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEncounter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEncounter(params['id']).subscribe(res => {
                this.encounter = res;
            });
        });
    }
}