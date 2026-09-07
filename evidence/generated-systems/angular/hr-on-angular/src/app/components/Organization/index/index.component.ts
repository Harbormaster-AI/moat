
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OrganizationService } from '../../../services/Organization.service';
import { Organization } from '../../../models/Organization';

@Component({
    selector: 'app-index-organization',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOrganizationComponent implements OnInit {

    organizations: Organization[] = [];

    constructor(
        private router: Router,
        private service: OrganizationService
) {}

    ngOnInit(): void {
        this.getOrganizations();
}

    getOrganizations(): void {
        this.service.getOrganizations().subscribe((res) => {
        this.organizations = res;
    });
}

    deleteOrganization(id: any): void {
        this.service.deleteOrganization(id)
            .subscribe(() => {
                this.getOrganizations();
            });
    }
}