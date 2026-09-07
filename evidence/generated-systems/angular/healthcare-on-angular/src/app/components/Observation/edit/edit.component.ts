import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ObservationService } from '../../../services/Observation.service';
import { SubBaseComponent } from '../../Observation/sub.base.component';


@Component({
    selector: 'app-edit-observation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditObservationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Observation';

    observationForm: FormGroup;
    observation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ObservationService,
        private fb: FormBuilder
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

    
    updateObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation): void {
        this.route.params.subscribe((params) => {

                        this.service.updateObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexObservation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getObservation(params['id']).subscribe(res => {
                this.observation = res;
            });
        });
    }
}