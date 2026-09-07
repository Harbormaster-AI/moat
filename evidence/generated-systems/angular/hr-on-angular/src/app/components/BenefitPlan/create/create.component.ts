import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BenefitPlanService } from '../../../services/BenefitPlan.service';
import { BenefitPlan } from '../../../models/BenefitPlan';
import { SubBaseComponent } from '../../BenefitPlan/sub.base.component';

@Component({
    selector: 'app-create-benefitPlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBenefitPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add BenefitPlan';

    benefitPlanForm: FormGroup;
    benefitPlan: BenefitPlan;

    constructor( http: HttpClient,
        private benefitPlanService: BenefitPlanService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType): void {
        this.benefitPlanService
        .addBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType)
            .subscribe(() => {
                this.router.navigate(['/indexBenefitPlan']);
            });
    }

    ngOnInit(): void {
    }
}