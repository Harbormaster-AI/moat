
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';
import { StockAdjustmentLine } from '../../../models/StockAdjustmentLine';

@Component({
    selector: 'app-index-stockAdjustmentLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexStockAdjustmentLineComponent implements OnInit {

    stockAdjustmentLines: StockAdjustmentLine[] = [];

    constructor(
        private router: Router,
        private service: StockAdjustmentLineService
) {}

    ngOnInit(): void {
        this.getStockAdjustmentLines();
}

    getStockAdjustmentLines(): void {
        this.service.getStockAdjustmentLines().subscribe((res) => {
        this.stockAdjustmentLines = res;
    });
}

    deleteStockAdjustmentLine(id: any): void {
        this.service.deleteStockAdjustmentLine(id)
            .subscribe(() => {
                this.getStockAdjustmentLines();
            });
    }
}