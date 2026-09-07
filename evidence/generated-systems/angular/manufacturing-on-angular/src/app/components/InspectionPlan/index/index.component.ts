
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InspectionPlanService } from '../../../services/InspectionPlan.service';
import { InspectionPlan } from '../../../models/InspectionPlan';

@Component({
    selector: 'app-index-inspectionPlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInspectionPlanComponent implements OnInit {

    inspectionPlans: InspectionPlan[] = [];

    constructor(
        private router: Router,
        private service: InspectionPlanService
) {}

    ngOnInit(): void {
        this.getInspectionPlans();
}

    getInspectionPlans(): void {
        this.service.getInspectionPlans().subscribe((res) => {
        this.inspectionPlans = res;
    });
}

    deleteInspectionPlan(id: any): void {
        this.service.deleteInspectionPlan(id)
            .subscribe(() => {
                this.getInspectionPlans();
            });
    }
}