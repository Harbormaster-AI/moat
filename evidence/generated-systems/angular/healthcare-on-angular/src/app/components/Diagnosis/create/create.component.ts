import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DiagnosisService } from '../../../services/Diagnosis.service';
import { Diagnosis } from '../../../models/Diagnosis';
import { SubBaseComponent } from '../../Diagnosis/sub.base.component';

@Component({
    selector: 'app-create-diagnosis',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDiagnosisComponent extends SubBaseComponent implements OnInit {

    title = 'Add Diagnosis';

    diagnosisForm: FormGroup;
    diagnosis: Diagnosis;

    constructor( http: HttpClient,
        private diagnosisService: DiagnosisService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty): void {
        this.diagnosisService
        .addDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty)
            .subscribe(() => {
                this.router.navigate(['/indexDiagnosis']);
            });
    }

    ngOnInit(): void {
    }
}