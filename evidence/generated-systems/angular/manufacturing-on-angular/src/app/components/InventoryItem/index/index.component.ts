
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InventoryItemService } from '../../../services/InventoryItem.service';
import { InventoryItem } from '../../../models/InventoryItem';

@Component({
    selector: 'app-index-inventoryItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInventoryItemComponent implements OnInit {

    inventoryItems: InventoryItem[] = [];

    constructor(
        private router: Router,
        private service: InventoryItemService
) {}

    ngOnInit(): void {
        this.getInventoryItems();
}

    getInventoryItems(): void {
        this.service.getInventoryItems().subscribe((res) => {
        this.inventoryItems = res;
    });
}

    deleteInventoryItem(id: any): void {
        this.service.deleteInventoryItem(id)
            .subscribe(() => {
                this.getInventoryItems();
            });
    }
}