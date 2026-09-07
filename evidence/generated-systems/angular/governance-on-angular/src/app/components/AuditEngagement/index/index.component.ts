
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuditEngagementService } from '../../../services/AuditEngagement.service';
import { AuditEngagement } from '../../../models/AuditEngagement';

@Component({
    selector: 'app-index-auditEngagement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAuditEngagementComponent implements OnInit {

    auditEngagements: AuditEngagement[] = [];

    constructor(
        private router: Router,
        private service: AuditEngagementService
) {}

    ngOnInit(): void {
        this.getAuditEngagements();
}

    getAuditEngagements(): void {
        this.service.getAuditEngagements().subscribe((res) => {
        this.auditEngagements = res;
    });
}

    deleteAuditEngagement(id: any): void {
        this.service.deleteAuditEngagement(id)
            .subscribe(() => {
                this.getAuditEngagements();
            });
    }
}