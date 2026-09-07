
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DepartmentService } from '../../../services/Department.service';
import { Department } from '../../../models/Department';

@Component({
    selector: 'app-index-department',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDepartmentComponent implements OnInit {

    departments: Department[] = [];

    constructor(
        private router: Router,
        private service: DepartmentService
) {}

    ngOnInit(): void {
        this.getDepartments();
}

    getDepartments(): void {
        this.service.getDepartments().subscribe((res) => {
        this.departments = res;
    });
}

    deleteDepartment(id: any): void {
        this.service.deleteDepartment(id)
            .subscribe(() => {
                this.getDepartments();
            });
    }
}