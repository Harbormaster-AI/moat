
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';
import { AuditWorkpaper } from '../../../models/AuditWorkpaper';

@Component({
    selector: 'app-index-auditWorkpaper',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAuditWorkpaperComponent implements OnInit {

    auditWorkpapers: AuditWorkpaper[] = [];

    constructor(
        private router: Router,
        private service: AuditWorkpaperService
) {}

    ngOnInit(): void {
        this.getAuditWorkpapers();
}

    getAuditWorkpapers(): void {
        this.service.getAuditWorkpapers().subscribe((res) => {
        this.auditWorkpapers = res;
    });
}

    deleteAuditWorkpaper(id: any): void {
        this.service.deleteAuditWorkpaper(id)
            .subscribe(() => {
                this.getAuditWorkpapers();
            });
    }
}