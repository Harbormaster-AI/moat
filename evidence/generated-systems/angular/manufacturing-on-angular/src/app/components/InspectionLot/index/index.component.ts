
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InspectionLotService } from '../../../services/InspectionLot.service';
import { InspectionLot } from '../../../models/InspectionLot';

@Component({
    selector: 'app-index-inspectionLot',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInspectionLotComponent implements OnInit {

    inspectionLots: InspectionLot[] = [];

    constructor(
        private router: Router,
        private service: InspectionLotService
) {}

    ngOnInit(): void {
        this.getInspectionLots();
}

    getInspectionLots(): void {
        this.service.getInspectionLots().subscribe((res) => {
        this.inspectionLots = res;
    });
}

    deleteInspectionLot(id: any): void {
        this.service.deleteInspectionLot(id)
            .subscribe(() => {
                this.getInspectionLots();
            });
    }
}