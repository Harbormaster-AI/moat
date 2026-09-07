
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StockAdjustmentService } from '../../../services/StockAdjustment.service';
import { StockAdjustment } from '../../../models/StockAdjustment';

@Component({
    selector: 'app-index-stockAdjustment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexStockAdjustmentComponent implements OnInit {

    stockAdjustments: StockAdjustment[] = [];

    constructor(
        private router: Router,
        private service: StockAdjustmentService
) {}

    ngOnInit(): void {
        this.getStockAdjustments();
}

    getStockAdjustments(): void {
        this.service.getStockAdjustments().subscribe((res) => {
        this.stockAdjustments = res;
    });
}

    deleteStockAdjustment(id: any): void {
        this.service.deleteStockAdjustment(id)
            .subscribe(() => {
                this.getStockAdjustments();
            });
    }
}