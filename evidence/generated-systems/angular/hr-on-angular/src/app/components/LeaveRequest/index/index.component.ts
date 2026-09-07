
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LeaveRequestService } from '../../../services/LeaveRequest.service';
import { LeaveRequest } from '../../../models/LeaveRequest';

@Component({
    selector: 'app-index-leaveRequest',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLeaveRequestComponent implements OnInit {

    leaveRequests: LeaveRequest[] = [];

    constructor(
        private router: Router,
        private service: LeaveRequestService
) {}

    ngOnInit(): void {
        this.getLeaveRequests();
}

    getLeaveRequests(): void {
        this.service.getLeaveRequests().subscribe((res) => {
        this.leaveRequests = res;
    });
}

    deleteLeaveRequest(id: any): void {
        this.service.deleteLeaveRequest(id)
            .subscribe(() => {
                this.getLeaveRequests();
            });
    }
}