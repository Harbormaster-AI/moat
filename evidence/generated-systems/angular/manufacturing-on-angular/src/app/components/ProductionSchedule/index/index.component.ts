
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductionScheduleService } from '../../../services/ProductionSchedule.service';
import { ProductionSchedule } from '../../../models/ProductionSchedule';

@Component({
    selector: 'app-index-productionSchedule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductionScheduleComponent implements OnInit {

    productionSchedules: ProductionSchedule[] = [];

    constructor(
        private router: Router,
        private service: ProductionScheduleService
) {}

    ngOnInit(): void {
        this.getProductionSchedules();
}

    getProductionSchedules(): void {
        this.service.getProductionSchedules().subscribe((res) => {
        this.productionSchedules = res;
    });
}

    deleteProductionSchedule(id: any): void {
        this.service.deleteProductionSchedule(id)
            .subscribe(() => {
                this.getProductionSchedules();
            });
    }
}