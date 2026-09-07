
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BenefitPlanService } from '../../../services/BenefitPlan.service';
import { BenefitPlan } from '../../../models/BenefitPlan';

@Component({
    selector: 'app-index-benefitPlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBenefitPlanComponent implements OnInit {

    benefitPlans: BenefitPlan[] = [];

    constructor(
        private router: Router,
        private service: BenefitPlanService
) {}

    ngOnInit(): void {
        this.getBenefitPlans();
}

    getBenefitPlans(): void {
        this.service.getBenefitPlans().subscribe((res) => {
        this.benefitPlans = res;
    });
}

    deleteBenefitPlan(id: any): void {
        this.service.deleteBenefitPlan(id)
            .subscribe(() => {
                this.getBenefitPlans();
            });
    }
}