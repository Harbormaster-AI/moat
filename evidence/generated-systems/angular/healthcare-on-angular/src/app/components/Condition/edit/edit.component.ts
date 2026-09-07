import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ConditionService } from '../../../services/Condition.service';
import { SubBaseComponent } from '../../Condition/sub.base.component';


@Component({
    selector: 'app-edit-condition',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditConditionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Condition';

    conditionForm: FormGroup;
    condition: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ConditionService,
        private fb: FormBuilder
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

    
    updateCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCondition']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCondition(params['id']).subscribe(res => {
                this.condition = res;
            });
        });
    }
}