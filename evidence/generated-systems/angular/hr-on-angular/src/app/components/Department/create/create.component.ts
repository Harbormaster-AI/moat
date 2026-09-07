import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DepartmentService } from '../../../services/Department.service';
import { Department } from '../../../models/Department';
import { SubBaseComponent } from '../../Department/sub.base.component';

@Component({
    selector: 'app-create-department',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDepartmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Department';

    departmentForm: FormGroup;
    department: Department;

    constructor( http: HttpClient,
        private departmentService: DepartmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.departmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      code: ['', Validators.required],
      Organization: ['', ],
      Manager: ['', ],
      Positions: ['', ],
      Employees: ['', ],
      CostCenter: ['', ]
        });
    }

    
    addDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter): void {
        this.departmentService
        .addDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter)
            .subscribe(() => {
                this.router.navigate(['/indexDepartment']);
            });
    }

    ngOnInit(): void {
    }
}