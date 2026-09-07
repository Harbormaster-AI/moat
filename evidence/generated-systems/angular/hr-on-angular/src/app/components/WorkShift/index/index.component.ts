
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WorkShiftService } from '../../../services/WorkShift.service';
import { WorkShift } from '../../../models/WorkShift';

@Component({
    selector: 'app-index-workShift',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWorkShiftComponent implements OnInit {

    workShifts: WorkShift[] = [];

    constructor(
        private router: Router,
        private service: WorkShiftService
) {}

    ngOnInit(): void {
        this.getWorkShifts();
}

    getWorkShifts(): void {
        this.service.getWorkShifts().subscribe((res) => {
        this.workShifts = res;
    });
}

    deleteWorkShift(id: any): void {
        this.service.deleteWorkShift(id)
            .subscribe(() => {
                this.getWorkShifts();
            });
    }
}