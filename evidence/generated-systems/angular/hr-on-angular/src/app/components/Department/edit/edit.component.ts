import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DepartmentService } from '../../../services/Department.service';
import { SubBaseComponent } from '../../Department/sub.base.component';


@Component({
    selector: 'app-edit-department',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDepartmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Department';

    departmentForm: FormGroup;
    department: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DepartmentService,
        private fb: FormBuilder
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

    
    updateDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDepartment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDepartment(params['id']).subscribe(res => {
                this.department = res;
            });
        });
    }
}