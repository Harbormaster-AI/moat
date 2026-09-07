import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConditionService } from '../../../services/Condition.service';
import { Condition } from '../../../models/Condition';
import { SubBaseComponent } from '../../Condition/sub.base.component';

@Component({
    selector: 'app-create-condition',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateConditionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Condition';

    conditionForm: FormGroup;
    condition: Condition;

    constructor( http: HttpClient,
        private conditionService: ConditionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.conditionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      onsetDate: ['', Validators.required],
      abatementDate: ['', Validators.required],
      Patient: ['', ],
      ClinicalStatus: ['', ],
      VerificationStatus: ['', ]
        });
    }

    
    addCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus): void {
        this.conditionService
        .addCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus)
            .subscribe(() => {
                this.router.navigate(['/indexCondition']);
            });
    }

    ngOnInit(): void {
    }
}