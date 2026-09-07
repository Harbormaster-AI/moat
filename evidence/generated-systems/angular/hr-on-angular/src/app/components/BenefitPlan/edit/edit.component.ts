import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BenefitPlanService } from '../../../services/BenefitPlan.service';
import { SubBaseComponent } from '../../BenefitPlan/sub.base.component';


@Component({
    selector: 'app-edit-benefitPlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBenefitPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BenefitPlan';

    benefitPlanForm: FormGroup;
    benefitPlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BenefitPlanService,
        private fb: FormBuilder
) {
        super(http);
        this.benefitPlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      providerName: ['', Validators.required],
      employeeContributionRate: ['', Validators.required],
      employerContributionRate: ['', Validators.required],
      eligibilityRules: ['', Validators.required],
      Organization: ['', ],
      Enrollments: ['', ],
      BenefitType: ['', ]
        });
    }

    
    updateBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBenefitPlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBenefitPlan(params['id']).subscribe(res => {
                this.benefitPlan = res;
            });
        });
    }
}