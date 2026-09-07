import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';
import { PolicyCoverage } from '../../../models/PolicyCoverage';
import { SubBaseComponent } from '../../PolicyCoverage/sub.base.component';

@Component({
    selector: 'app-create-policyCoverage',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePolicyCoverageComponent extends SubBaseComponent implements OnInit {

    title = 'Add PolicyCoverage';

    policyCoverageForm: FormGroup;
    policyCoverage: PolicyCoverage;

    constructor( http: HttpClient,
        private policyCoverageService: PolicyCoverageService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType): void {
        this.policyCoverageService
        .addPolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType)
            .subscribe(() => {
                this.router.navigate(['/indexPolicyCoverage']);
            });
    }

    ngOnInit(): void {
    }
}