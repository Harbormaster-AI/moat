import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';
import { SubBaseComponent } from '../../PolicyCoverage/sub.base.component';


@Component({
    selector: 'app-edit-policyCoverage',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPolicyCoverageComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PolicyCoverage';

    policyCoverageForm: FormGroup;
    policyCoverage: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PolicyCoverageService,
        private fb: FormBuilder
) {
        super(http);
        this.policyCoverageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  limit: ['', Validators.required],
      deductible: ['', Validators.required],
      premium: ['', Validators.required],
      Policy: ['', ],
      InsuredObjects: ['', ],
      CoverageType: ['', ]
        });
    }

    
    updatePolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPolicyCoverage']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPolicyCoverage(params['id']).subscribe(res => {
                this.policyCoverage = res;
            });
        });
    }
}