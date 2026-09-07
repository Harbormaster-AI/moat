
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';
import { PurchaseOrderLine } from '../../../models/PurchaseOrderLine';

@Component({
    selector: 'app-index-purchaseOrderLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPurchaseOrderLineComponent implements OnInit {

    purchaseOrderLines: PurchaseOrderLine[] = [];

    constructor(
        private router: Router,
        private service: PurchaseOrderLineService
) {}

    ngOnInit(): void {
        this.getPurchaseOrderLines();
}

    getPurchaseOrderLines(): void {
        this.service.getPurchaseOrderLines().subscribe((res) => {
        this.purchaseOrderLines = res;
    });
}

    deletePurchaseOrderLine(id: any): void {
        this.service.deletePurchaseOrderLine(id)
            .subscribe(() => {
                this.getPurchaseOrderLines();
            });
    }
}