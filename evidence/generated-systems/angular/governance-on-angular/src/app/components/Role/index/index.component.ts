
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RoleService } from '../../../services/Role.service';
import { Role } from '../../../models/Role';

@Component({
    selector: 'app-index-role',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRoleComponent implements OnInit {

    roles: Role[] = [];

    constructor(
        private router: Router,
        private service: RoleService
) {}

    ngOnInit(): void {
        this.getRoles();
}

    getRoles(): void {
        this.service.getRoles().subscribe((res) => {
        this.roles = res;
    });
}

    deleteRole(id: any): void {
        this.service.deleteRole(id)
            .subscribe(() => {
                this.getRoles();
            });
    }
}