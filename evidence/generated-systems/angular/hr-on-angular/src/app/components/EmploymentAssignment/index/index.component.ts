
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';
import { EmploymentAssignment } from '../../../models/EmploymentAssignment';

@Component({
    selector: 'app-index-employmentAssignment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEmploymentAssignmentComponent implements OnInit {

    employmentAssignments: EmploymentAssignment[] = [];

    constructor(
        private router: Router,
        private service: EmploymentAssignmentService
) {}

    ngOnInit(): void {
        this.getEmploymentAssignments();
}

    getEmploymentAssignments(): void {
        this.service.getEmploymentAssignments().subscribe((res) => {
        this.employmentAssignments = res;
    });
}

    deleteEmploymentAssignment(id: any): void {
        this.service.deleteEmploymentAssignment(id)
            .subscribe(() => {
                this.getEmploymentAssignments();
            });
    }
}