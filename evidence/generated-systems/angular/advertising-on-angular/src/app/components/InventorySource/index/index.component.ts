
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InventorySourceService } from '../../../services/InventorySource.service';
import { InventorySource } from '../../../models/InventorySource';

@Component({
    selector: 'app-index-inventorySource',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInventorySourceComponent implements OnInit {

    inventorySources: InventorySource[] = [];

    constructor(
        private router: Router,
        private service: InventorySourceService
) {}

    ngOnInit(): void {
        this.getInventorySources();
}

    getInventorySources(): void {
        this.service.getInventorySources().subscribe((res) => {
        this.inventorySources = res;
    });
}

    deleteInventorySource(id: any): void {
        this.service.deleteInventorySource(id)
            .subscribe(() => {
                this.getInventorySources();
            });
    }
}