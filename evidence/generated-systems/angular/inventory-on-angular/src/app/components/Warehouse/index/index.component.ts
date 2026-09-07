
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WarehouseService } from '../../../services/Warehouse.service';
import { Warehouse } from '../../../models/Warehouse';

@Component({
    selector: 'app-index-warehouse',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWarehouseComponent implements OnInit {

    warehouses: Warehouse[] = [];

    constructor(
        private router: Router,
        private service: WarehouseService
) {}

    ngOnInit(): void {
        this.getWarehouses();
}

    getWarehouses(): void {
        this.service.getWarehouses().subscribe((res) => {
        this.warehouses = res;
    });
}

    deleteWarehouse(id: any): void {
        this.service.deleteWarehouse(id)
            .subscribe(() => {
                this.getWarehouses();
            });
    }
}