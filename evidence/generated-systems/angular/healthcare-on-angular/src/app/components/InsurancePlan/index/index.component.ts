
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsurancePlanService } from '../../../services/InsurancePlan.service';
import { InsurancePlan } from '../../../models/InsurancePlan';

@Component({
    selector: 'app-index-insurancePlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsurancePlanComponent implements OnInit {

    insurancePlans: InsurancePlan[] = [];

    constructor(
        private router: Router,
        private service: InsurancePlanService
) {}

    ngOnInit(): void {
        this.getInsurancePlans();
}

    getInsurancePlans(): void {
        this.service.getInsurancePlans().subscribe((res) => {
        this.insurancePlans = res;
    });
}

    deleteInsurancePlan(id: any): void {
        this.service.deleteInsurancePlan(id)
            .subscribe(() => {
                this.getInsurancePlans();
            });
    }
}