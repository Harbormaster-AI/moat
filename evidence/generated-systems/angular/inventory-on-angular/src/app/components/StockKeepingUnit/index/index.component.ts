
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';
import { StockKeepingUnit } from '../../../models/StockKeepingUnit';

@Component({
    selector: 'app-index-stockKeepingUnit',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexStockKeepingUnitComponent implements OnInit {

    stockKeepingUnits: StockKeepingUnit[] = [];

    constructor(
        private router: Router,
        private service: StockKeepingUnitService
) {}

    ngOnInit(): void {
        this.getStockKeepingUnits();
}

    getStockKeepingUnits(): void {
        this.service.getStockKeepingUnits().subscribe((res) => {
        this.stockKeepingUnits = res;
    });
}

    deleteStockKeepingUnit(id: any): void {
        this.service.deleteStockKeepingUnit(id)
            .subscribe(() => {
                this.getStockKeepingUnits();
            });
    }
}