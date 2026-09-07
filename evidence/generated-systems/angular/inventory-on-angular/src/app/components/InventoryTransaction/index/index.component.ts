
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InventoryTransactionService } from '../../../services/InventoryTransaction.service';
import { InventoryTransaction } from '../../../models/InventoryTransaction';

@Component({
    selector: 'app-index-inventoryTransaction',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInventoryTransactionComponent implements OnInit {

    inventoryTransactions: InventoryTransaction[] = [];

    constructor(
        private router: Router,
        private service: InventoryTransactionService
) {}

    ngOnInit(): void {
        this.getInventoryTransactions();
}

    getInventoryTransactions(): void {
        this.service.getInventoryTransactions().subscribe((res) => {
        this.inventoryTransactions = res;
    });
}

    deleteInventoryTransaction(id: any): void {
        this.service.deleteInventoryTransaction(id)
            .subscribe(() => {
                this.getInventoryTransactions();
            });
    }
}