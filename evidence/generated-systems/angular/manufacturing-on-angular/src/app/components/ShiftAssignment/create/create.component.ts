import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ShiftAssignmentService } from '../../../services/ShiftAssignment.service';
import { ShiftAssignment } from '../../../models/ShiftAssignment';
import { SubBaseComponent } from '../../ShiftAssignment/sub.base.component';

@Component({
    selector: 'app-create-shiftAssignment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateShiftAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add ShiftAssignment';

    shiftAssignmentForm: FormGroup;
    shiftAssignment: ShiftAssignment;

    constructor( http: HttpClient,
        private shiftAssignmentService: ShiftAssignmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.shiftAssignmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  assignmentDate: ['', Validators.required],
      Shift: ['', ],
      Employee: ['', ],
      WorkCenter: ['', ]
        });
    }

    
    addShiftAssignment(assignmentDate, Shift, Employee, WorkCenter): void {
        this.shiftAssignmentService
        .addShiftAssignment(assignmentDate, Shift, Employee, WorkCenter)
            .subscribe(() => {
                this.router.navigate(['/indexShiftAssignment']);
            });
    }

    ngOnInit(): void {
    }
}