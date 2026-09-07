
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FeeScheduleService } from '../../../services/FeeSchedule.service';
import { FeeSchedule } from '../../../models/FeeSchedule';

@Component({
    selector: 'app-index-feeSchedule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFeeScheduleComponent implements OnInit {

    feeSchedules: FeeSchedule[] = [];

    constructor(
        private router: Router,
        private service: FeeScheduleService
) {}

    ngOnInit(): void {
        this.getFeeSchedules();
}

    getFeeSchedules(): void {
        this.service.getFeeSchedules().subscribe((res) => {
        this.feeSchedules = res;
    });
}

    deleteFeeSchedule(id: any): void {
        this.service.deleteFeeSchedule(id)
            .subscribe(() => {
                this.getFeeSchedules();
            });
    }
}