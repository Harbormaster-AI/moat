import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ObservationService } from '../../../services/Observation.service';
import { Observation } from '../../../models/Observation';
import { SubBaseComponent } from '../../Observation/sub.base.component';

@Component({
    selector: 'app-create-observation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateObservationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Observation';

    observationForm: FormGroup;
    observation: Observation;

    constructor( http: HttpClient,
        private observationService: ObservationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.observationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      value: ['', Validators.required],
      unit: ['', Validators.required],
      effectiveDateTime: ['', Validators.required],
      Encounter: ['', ],
      Patient: ['', ],
      Device: ['', ],
      LabResult: ['', ],
      Interpretation: ['', ]
        });
    }

    
    addObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation): void {
        this.observationService
        .addObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation)
            .subscribe(() => {
                this.router.navigate(['/indexObservation']);
            });
    }

    ngOnInit(): void {
    }
}