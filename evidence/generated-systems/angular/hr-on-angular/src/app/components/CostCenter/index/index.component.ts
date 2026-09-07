
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CostCenterService } from '../../../services/CostCenter.service';
import { CostCenter } from '../../../models/CostCenter';

@Component({
    selector: 'app-index-costCenter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCostCenterComponent implements OnInit {

    costCenters: CostCenter[] = [];

    constructor(
        private router: Router,
        private service: CostCenterService
) {}

    ngOnInit(): void {
        this.getCostCenters();
}

    getCostCenters(): void {
        this.service.getCostCenters().subscribe((res) => {
        this.costCenters = res;
    });
}

    deleteCostCenter(id: any): void {
        this.service.deleteCostCenter(id)
            .subscribe(() => {
                this.getCostCenters();
            });
    }
}