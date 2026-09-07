
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ShiftService } from '../../../services/Shift.service';
import { Shift } from '../../../models/Shift';

@Component({
    selector: 'app-index-shift',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexShiftComponent implements OnInit {

    shifts: Shift[] = [];

    constructor(
        private router: Router,
        private service: ShiftService
) {}

    ngOnInit(): void {
        this.getShifts();
}

    getShifts(): void {
        this.service.getShifts().subscribe((res) => {
        this.shifts = res;
    });
}

    deleteShift(id: any): void {
        this.service.deleteShift(id)
            .subscribe(() => {
                this.getShifts();
            });
    }
}