
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SecurityService } from '../../../services/Security.service';
import { Security } from '../../../models/Security';

@Component({
    selector: 'app-index-security',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSecurityComponent implements OnInit {

    securitys: Security[] = [];

    constructor(
        private router: Router,
        private service: SecurityService
) {}

    ngOnInit(): void {
        this.getSecuritys();
}

    getSecuritys(): void {
        this.service.getSecuritys().subscribe((res) => {
        this.securitys = res;
    });
}

    deleteSecurity(id: any): void {
        this.service.deleteSecurity(id)
            .subscribe(() => {
                this.getSecuritys();
            });
    }
}