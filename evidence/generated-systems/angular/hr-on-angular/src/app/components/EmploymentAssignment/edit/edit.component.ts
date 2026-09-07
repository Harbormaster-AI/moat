import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';
import { SubBaseComponent } from '../../EmploymentAssignment/sub.base.component';


@Component({
    selector: 'app-edit-employmentAssignment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEmploymentAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EmploymentAssignment';

    employmentAssignmentForm: FormGroup;
    employmentAssignment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EmploymentAssignmentService,
        private fb: FormBuilder
) {
        super(http);
        this.employmentAssignmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      primary: ['', Validators.required],
      Employee: ['', ],
      Position: ['', ],
      Supervisor: ['', ],
      AssignmentType: ['', ],
      Status: ['', ]
        });
    }

    
    updateEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEmploymentAssignment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEmploymentAssignment(params['id']).subscribe(res => {
                this.employmentAssignment = res;
            });
        });
    }
}