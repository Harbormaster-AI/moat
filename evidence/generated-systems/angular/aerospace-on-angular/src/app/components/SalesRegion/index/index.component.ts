
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SalesRegionService } from '../../../services/SalesRegion.service';
import { SalesRegion } from '../../../models/SalesRegion';

@Component({
    selector: 'app-index-salesRegion',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSalesRegionComponent implements OnInit {

    salesRegions: SalesRegion[] = [];

    constructor(
        private router: Router,
        private service: SalesRegionService
) {}

    ngOnInit(): void {
        this.getSalesRegions();
}

    getSalesRegions(): void {
        this.service.getSalesRegions().subscribe((res) => {
        this.salesRegions = res;
    });
}

    deleteSalesRegion(id: any): void {
        this.service.deleteSalesRegion(id)
            .subscribe(() => {
                this.getSalesRegions();
            });
    }
}