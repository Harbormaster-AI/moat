
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ShiftAssignmentService } from '../../../services/ShiftAssignment.service';
import { ShiftAssignment } from '../../../models/ShiftAssignment';

@Component({
    selector: 'app-index-shiftAssignment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexShiftAssignmentComponent implements OnInit {

    shiftAssignments: ShiftAssignment[] = [];

    constructor(
        private router: Router,
        private service: ShiftAssignmentService
) {}

    ngOnInit(): void {
        this.getShiftAssignments();
}

    getShiftAssignments(): void {
        this.service.getShiftAssignments().subscribe((res) => {
        this.shiftAssignments = res;
    });
}

    deleteShiftAssignment(id: any): void {
        this.service.deleteShiftAssignment(id)
            .subscribe(() => {
                this.getShiftAssignments();
            });
    }
}