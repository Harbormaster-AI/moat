
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LeavePolicyService } from '../../../services/LeavePolicy.service';
import { LeavePolicy } from '../../../models/LeavePolicy';

@Component({
    selector: 'app-index-leavePolicy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLeavePolicyComponent implements OnInit {

    leavePolicys: LeavePolicy[] = [];

    constructor(
        private router: Router,
        private service: LeavePolicyService
) {}

    ngOnInit(): void {
        this.getLeavePolicys();
}

    getLeavePolicys(): void {
        this.service.getLeavePolicys().subscribe((res) => {
        this.leavePolicys = res;
    });
}

    deleteLeavePolicy(id: any): void {
        this.service.deleteLeavePolicy(id)
            .subscribe(() => {
                this.getLeavePolicys();
            });
    }
}