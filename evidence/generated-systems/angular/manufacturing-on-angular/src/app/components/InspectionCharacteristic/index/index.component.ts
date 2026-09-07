
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';
import { InspectionCharacteristic } from '../../../models/InspectionCharacteristic';

@Component({
    selector: 'app-index-inspectionCharacteristic',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInspectionCharacteristicComponent implements OnInit {

    inspectionCharacteristics: InspectionCharacteristic[] = [];

    constructor(
        private router: Router,
        private service: InspectionCharacteristicService
) {}

    ngOnInit(): void {
        this.getInspectionCharacteristics();
}

    getInspectionCharacteristics(): void {
        this.service.getInspectionCharacteristics().subscribe((res) => {
        this.inspectionCharacteristics = res;
    });
}

    deleteInspectionCharacteristic(id: any): void {
        this.service.deleteInspectionCharacteristic(id)
            .subscribe(() => {
                this.getInspectionCharacteristics();
            });
    }
}