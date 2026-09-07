
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RoleAssignmentService } from '../../../services/RoleAssignment.service';
import { RoleAssignment } from '../../../models/RoleAssignment';

@Component({
    selector: 'app-index-roleAssignment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRoleAssignmentComponent implements OnInit {

    roleAssignments: RoleAssignment[] = [];

    constructor(
        private router: Router,
        private service: RoleAssignmentService
) {}

    ngOnInit(): void {
        this.getRoleAssignments();
}

    getRoleAssignments(): void {
        this.service.getRoleAssignments().subscribe((res) => {
        this.roleAssignments = res;
    });
}

    deleteRoleAssignment(id: any): void {
        this.service.deleteRoleAssignment(id)
            .subscribe(() => {
                this.getRoleAssignments();
            });
    }
}