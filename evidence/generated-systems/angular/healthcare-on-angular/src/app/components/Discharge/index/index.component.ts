
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DischargeService } from '../../../services/Discharge.service';
import { Discharge } from '../../../models/Discharge';

@Component({
    selector: 'app-index-discharge',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDischargeComponent implements OnInit {

    discharges: Discharge[] = [];

    constructor(
        private router: Router,
        private service: DischargeService
) {}

    ngOnInit(): void {
        this.getDischarges();
}

    getDischarges(): void {
        this.service.getDischarges().subscribe((res) => {
        this.discharges = res;
    });
}

    deleteDischarge(id: any): void {
        this.service.deleteDischarge(id)
            .subscribe(() => {
                this.getDischarges();
            });
    }
}