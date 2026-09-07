
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PricingPlanService } from '../../../services/PricingPlan.service';
import { PricingPlan } from '../../../models/PricingPlan';

@Component({
    selector: 'app-index-pricingPlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPricingPlanComponent implements OnInit {

    pricingPlans: PricingPlan[] = [];

    constructor(
        private router: Router,
        private service: PricingPlanService
) {}

    ngOnInit(): void {
        this.getPricingPlans();
}

    getPricingPlans(): void {
        this.service.getPricingPlans().subscribe((res) => {
        this.pricingPlans = res;
    });
}

    deletePricingPlan(id: any): void {
        this.service.deletePricingPlan(id)
            .subscribe(() => {
                this.getPricingPlans();
            });
    }
}