import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DiagnosisService } from '../../../services/Diagnosis.service';
import { SubBaseComponent } from '../../Diagnosis/sub.base.component';


@Component({
    selector: 'app-edit-diagnosis',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDiagnosisComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Diagnosis';

    diagnosisForm: FormGroup;
    diagnosis: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DiagnosisService,
        private fb: FormBuilder
) {
        super(http);
        this.diagnosisForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      description: ['', Validators.required],
      onsetDate: ['', Validators.required],
      Encounter: ['', ],
      Patient: ['', ],
      Certainty: ['', ]
        });
    }

    
    updateDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDiagnosis']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDiagnosis(params['id']).subscribe(res => {
                this.diagnosis = res;
            });
        });
    }
}