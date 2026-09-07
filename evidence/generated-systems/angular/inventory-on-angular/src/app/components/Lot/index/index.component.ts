
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LotService } from '../../../services/Lot.service';
import { Lot } from '../../../models/Lot';

@Component({
    selector: 'app-index-lot',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLotComponent implements OnInit {

    lots: Lot[] = [];

    constructor(
        private router: Router,
        private service: LotService
) {}

    ngOnInit(): void {
        this.getLots();
}

    getLots(): void {
        this.service.getLots().subscribe((res) => {
        this.lots = res;
    });
}

    deleteLot(id: any): void {
        this.service.deleteLot(id)
            .subscribe(() => {
                this.getLots();
            });
    }
}