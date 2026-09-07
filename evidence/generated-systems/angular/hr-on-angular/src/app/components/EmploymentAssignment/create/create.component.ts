import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';
import { EmploymentAssignment } from '../../../models/EmploymentAssignment';
import { SubBaseComponent } from '../../EmploymentAssignment/sub.base.component';

@Component({
    selector: 'app-create-employmentAssignment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEmploymentAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add EmploymentAssignment';

    employmentAssignmentForm: FormGroup;
    employmentAssignment: EmploymentAssignment;

    constructor( http: HttpClient,
        private employmentAssignmentService: EmploymentAssignmentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status): void {
        this.employmentAssignmentService
        .addEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexEmploymentAssignment']);
            });
    }

    ngOnInit(): void {
    }
}