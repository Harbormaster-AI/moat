
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';
import { InventoryThresholdAlert } from '../../../models/InventoryThresholdAlert';

@Component({
    selector: 'app-index-inventoryThresholdAlert',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInventoryThresholdAlertComponent implements OnInit {

    inventoryThresholdAlerts: InventoryThresholdAlert[] = [];

    constructor(
        private router: Router,
        private service: InventoryThresholdAlertService
) {}

    ngOnInit(): void {
        this.getInventoryThresholdAlerts();
}

    getInventoryThresholdAlerts(): void {
        this.service.getInventoryThresholdAlerts().subscribe((res) => {
        this.inventoryThresholdAlerts = res;
    });
}

    deleteInventoryThresholdAlert(id: any): void {
        this.service.deleteInventoryThresholdAlert(id)
            .subscribe(() => {
                this.getInventoryThresholdAlerts();
            });
    }
}