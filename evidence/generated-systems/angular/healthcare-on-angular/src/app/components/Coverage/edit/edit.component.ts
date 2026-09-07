import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CoverageService } from '../../../services/Coverage.service';
import { SubBaseComponent } from '../../Coverage/sub.base.component';


@Component({
    selector: 'app-edit-coverage',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCoverageComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Coverage';

    coverageForm: FormGroup;
    coverage: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CoverageService,
        private fb: FormBuilder
) {
        super(http);
        this.coverageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  memberId: ['', Validators.required],
      groupNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      endDate: ['', Validators.required],
      Patient: ['', ],
      Plan: ['', ],
      Claims: ['', ],
      Authorizations: ['', ],
      CoverageType: ['', ]
        });
    }

    
    updateCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCoverage']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCoverage(params['id']).subscribe(res => {
                this.coverage = res;
            });
        });
    }
}