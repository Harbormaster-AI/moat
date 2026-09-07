
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApprovalService } from '../../../services/Approval.service';
import { Approval } from '../../../models/Approval';

@Component({
    selector: 'app-index-approval',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexApprovalComponent implements OnInit {

    approvals: Approval[] = [];

    constructor(
        private router: Router,
        private service: ApprovalService
) {}

    ngOnInit(): void {
        this.getApprovals();
}

    getApprovals(): void {
        this.service.getApprovals().subscribe((res) => {
        this.approvals = res;
    });
}

    deleteApproval(id: any): void {
        this.service.deleteApproval(id)
            .subscribe(() => {
                this.getApprovals();
            });
    }
}