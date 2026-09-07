
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CarePlanService } from '../../../services/CarePlan.service';
import { CarePlan } from '../../../models/CarePlan';

@Component({
    selector: 'app-index-carePlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCarePlanComponent implements OnInit {

    carePlans: CarePlan[] = [];

    constructor(
        private router: Router,
        private service: CarePlanService
) {}

    ngOnInit(): void {
        this.getCarePlans();
}

    getCarePlans(): void {
        this.service.getCarePlans().subscribe((res) => {
        this.carePlans = res;
    });
}

    deleteCarePlan(id: any): void {
        this.service.deleteCarePlan(id)
            .subscribe(() => {
                this.getCarePlans();
            });
    }
}