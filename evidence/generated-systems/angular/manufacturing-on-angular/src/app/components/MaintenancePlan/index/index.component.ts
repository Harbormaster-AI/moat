
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';
import { MaintenancePlan } from '../../../models/MaintenancePlan';

@Component({
    selector: 'app-index-maintenancePlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMaintenancePlanComponent implements OnInit {

    maintenancePlans: MaintenancePlan[] = [];

    constructor(
        private router: Router,
        private service: MaintenancePlanService
) {}

    ngOnInit(): void {
        this.getMaintenancePlans();
}

    getMaintenancePlans(): void {
        this.service.getMaintenancePlans().subscribe((res) => {
        this.maintenancePlans = res;
    });
}

    deleteMaintenancePlan(id: any): void {
        this.service.deleteMaintenancePlan(id)
            .subscribe(() => {
                this.getMaintenancePlans();
            });
    }
}