
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuditFindingService } from '../../../services/AuditFinding.service';
import { AuditFinding } from '../../../models/AuditFinding';

@Component({
    selector: 'app-index-auditFinding',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAuditFindingComponent implements OnInit {

    auditFindings: AuditFinding[] = [];

    constructor(
        private router: Router,
        private service: AuditFindingService
) {}

    ngOnInit(): void {
        this.getAuditFindings();
}

    getAuditFindings(): void {
        this.service.getAuditFindings().subscribe((res) => {
        this.auditFindings = res;
    });
}

    deleteAuditFinding(id: any): void {
        this.service.deleteAuditFinding(id)
            .subscribe(() => {
                this.getAuditFindings();
            });
    }
}