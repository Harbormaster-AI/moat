
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuditProgramService } from '../../../services/AuditProgram.service';
import { AuditProgram } from '../../../models/AuditProgram';

@Component({
    selector: 'app-index-auditProgram',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAuditProgramComponent implements OnInit {

    auditPrograms: AuditProgram[] = [];

    constructor(
        private router: Router,
        private service: AuditProgramService
) {}

    ngOnInit(): void {
        this.getAuditPrograms();
}

    getAuditPrograms(): void {
        this.service.getAuditPrograms().subscribe((res) => {
        this.auditPrograms = res;
    });
}

    deleteAuditProgram(id: any): void {
        this.service.deleteAuditProgram(id)
            .subscribe(() => {
                this.getAuditPrograms();
            });
    }
}